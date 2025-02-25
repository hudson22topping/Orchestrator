package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, htmlContent)
    })

    fmt.Println("Server started at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

const htmlContent = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Play Note</title>
</head>
<body>
    <h1>Playing Note</h1>
    <p>The page is constantly playing a note.</p>
    <script>
        // Create an AudioContext
        const audioContext = new (window.AudioContext || window.webkitAudioContext)();

        // Create an oscillator
        const oscillator = audioContext.createOscillator();

        // Set the oscillator frequency to play a constant note (e.g., C4)
        oscillator.frequency.setValueAtTime(261.63, audioContext.currentTime); // C4 frequency

        // Connect the oscillator to the destination (speakers)
        oscillator.connect(audioContext.destination);

        // Start the oscillator
        oscillator.start();

        // Stop the oscillator when the page is unloaded
        window.addEventListener('beforeunload', () => {
            oscillator.stop();
        });
    </script>
</body>
</html>
`