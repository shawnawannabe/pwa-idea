### idea: trying to create a minimalist budget pwa while exploring the below technology
- front-end: vanilla js for now, might explore vue.js
- database: mysql / postgres
- backend: go
<br/><br/>
- resource:
  - icon: [flaticon.com](https://flaticon.com)
  - color palettes: [coolors.co](https://coolors.co) (for theme color)
  - add mask: [maskable.app](https://maskable.app)
  - source code reference: [this repo](https://github.com/firtman/frontendmasters-pwa)
<br/><br/>
- components / concept behind pwa:
  - manifest: the "heart" of the app [app.webmanifest](/static/app.webmanifest)
  - service-worker: 
    - a js file that has its own thread 
    - act as a cache/ web server on client-side
    - runs on browser's engine
    - a low-level API (lots of power)
    - to check it out on browser - chrome://serviceworker-internals/
