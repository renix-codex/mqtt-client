# mqtt-client

Message Queuing Telemetry Transport

MQTT (Message Queuing Telemetry Transport) is a messaging protocol that allows devices to communicate with each other over the internet or other networks. It uses a publish-subscribe model, where publishers send messages to a broker, which then distributes those messages to subscribers who have expressed interest in receiving them. MQTT brokers provide additional features that allow for more advanced message queuing, such as message buffering or persistent storage. These features can be useful in scenarios where message delivery is critical, and messages must be delivered even in the event of a network outage or broker failure.

Host MQTT Broker

Pull docker image for MQTT Broker

docker pull eclipse-mosquitto

Create a Docker network

docker network create mqtt-net

Run the MQTT Broker container

docker run -d --name mqtt-broker --network mqtt-net -p 1883:1883 -p 9001:9001 eclipse-mosquitto
