import {type  FormEvent } from "react"

function App() {
  const socket = new WebSocket(import.meta.env.GO_WS_URL as string)

  socket.addEventListener("open", () => {
    console.log("Connected to server")
  })

  socket.addEventListener("message", (event) => {
    console.log("Message from server: ", event.data)
  })

  const sendMessage = (e: FormEvent) => {  
    e.preventDefault()
    
    const formData = new FormData(e.target as HTMLFormElement)
    const formValues = {
      msg: formData.get("msg")
    }
    if(!formValues.msg) throw new Error("WTF Mati... Message is empty")
    socket.send(formValues.msg)
  }

  return (
    <div className="flex justify-center items-center">
      <form onSubmit={sendMessage} className="bg-white rounded-xl p-4 text-black h-[550px] flex flex-col justify-end w-md shadow-sm shadow-cyan-400">
      <textarea  name="msg" id="msg" placeholder="Napisz wiadomość" className="border  rounded-md border-slate-300 px-3 py-1.5"/>
      <button type="submit" className="bg-white border-slate-300 mt-2 text-slate-100">Send</button>
    </form>
    </div>
  )
}

export default App
