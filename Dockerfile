FROM debian:stretch-slim

# based on architecture
COPY fpga-device-plugin /usr/bin/fpga-device-plugin

CMD ["fpga-device-plugin"]
