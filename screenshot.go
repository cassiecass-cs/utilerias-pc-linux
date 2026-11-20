//Todo programa en go empieza siempre definiendo su paquete
package main



//importamos el modulo de ejecutar comandos del sistema
import(
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// Función que toma un screenshot con la utileria de gnome-screenshot
func screenshot(){

//Definimos una ruta donde se almacenaran los screenshots
	screenshot_folder := "/home/tu_usuario/Imagenes/"

	//Creamos la carpeta de los screenshots si no existe
	err := os.MkdirAll(screenshot_folder,0755)

	if err != nil{
		return
	}

	//También podemos almacenar aqui los logs de error
	logs_folder := screenshot_folder + "logs/"

	//Creamos la carpeta de los logs
	err = os.MkdirAll(logs_folder,0755)

	if err != nil{
		return
	}

	//Archivo de logs
	logs_file := logs_folder +"logs.txt"

	//Abrimos o creamos el archivo logs.txt
	file,err := os.OpenFile(
		logs_file,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)

	if err != nil{
		return
	}

	defer file.Close()

	//Generamos un nombre de archivo con fecha y hora para no sobrescribir capturas
	nombre_archivo := fmt.Sprintf("screenshot_%s.png",time.Now().Format("20060102_150405"))
	ruta_captura := filepath.Join(screenshot_folder,nombre_archivo)

	//Ejecutamos el comando que activa la utileria del sistema gnome
	//para la captura de screenshots
	//Es importante instalar en tu distro linux esta utilidad gnome-screenshot
	//EL comando toma una captura de pantalla global y la almacena
	//En la ruta que le especifiques

	out, err := exec.Command("gnome-screenshot","-f",ruta_captura).CombinedOutput()	

	//Revisamos que si hubo errores
	if err != nil{

		//Verificamos que fecha ocurrio el error (en Go siempre se usa la fecha de referencia 2006-01-02)
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		
		//Escribimos en el archivo de logs si hubo un error
		mensaje_error := fmt.Sprintf("Ocurrio un error al crear el screenshot [%s]: %v | %s\n",timestamp,err,string(out))
		_, err = file.WriteString(mensaje_error)

		if err != nil{
			return
		}
	
	}
	
}


//Función principal que ejecuta la aplicación
func main(){

	screenshot()
	
}
