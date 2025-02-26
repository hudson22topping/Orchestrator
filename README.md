# Orchestrator
Parts of my orchestrator

1. Create application using like Kafka which converts music into "commands" which show how long each note will play
2. Connect this to Kubernetes API
3. Have Frontend Web Server that takes gRPC commands to play certain notes for certain durations
4. Create backend Note-senders which take commands from Kubernetes and send them as gRPC instructions to the front-end
5. Containerize all backend and frontend apps
6. Configure control Node to spin-up pods based on what notes need to be played when
7. Kubernetes API creates services, jobs, and deployments to get it running.
