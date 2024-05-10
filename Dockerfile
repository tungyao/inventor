FROM ubuntu
COPY . /app
WORKDIR /app
RUN chmod +x /app/inventor
EXPOSE 8000
ENTRYPOINT ["/app/inventor"]