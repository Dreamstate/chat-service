1. Spin up a web server with a REST API mounted on /api/v1
2. Implement the following routes:
    - /register - for creating new accounts
    - /login - for authenticating with the server
    - /ws - for websocket communication, authentication required
3. Implement websocket control messages:
    - list-rooms: Shows a list of chat rooms available on the server
    - join-room: Connect to an existing chat room
    - create-room: Create a new chat room
    - list-users: Show members in a room
    - direct-message: Open Direct Message communication with another member
    - send-message: Sends a message to a room/user