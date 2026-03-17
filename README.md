# IRC-Go Chat Server 🚀

Aquesta és una implementació senzilla d'un servidor i client de xat basat en el protocol IRC, desenvolupat amb el llenguatge Go fent servir exclusivament la llibreria estàndard (`net` i `bufio`).

---

## 📋 Característiques

- **Arquitectura Client-Servidor**  
  Gestió de múltiples connexions simultànies mitjançant goroutines.

- **Identificació Automàtica**  
  Assignació de nicks per defecte basats en el port de connexió (ex: `user_58649`).

- **Nicks Únics**  
  Sistema de validació per garantir que no hi hagi dos usuaris amb el mateix nom.

- **Comandes IRC**
  - `/NICK <nou_nick>`: Canvia el teu nom d'usuari (insensible a majúscules).
  - `/QUIT`: Tanca la connexió de forma neta.

- **Broadcast Intel·ligent**  
  Els missatges s'envien a tots els usuaris connectats exceptuant l'emissor (evitant l'efecte "eco").

---

## 🛠️ Tecnologies utilitzades

- **Llenguatge:** Go (Golang)
- **Paquets:**
  - `net` (Sockets TCP)
  - `bufio` (I/O amb buffer)
  - `sync` (Control de concurrència amb Mutex)

---

## 📂 Estructura del Projecte

```plaintext
.
├── server.go       # Lògica del servidor IRC
├── client.go       # Client interactiu per al terminal
└── go.mod          # Definició del mòdul de Go
