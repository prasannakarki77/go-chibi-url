/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: [
      {
        mytheme: {
          primary: "#3dfca1",
          "primary-content": "#011609",
          secondary: "#ce34f7",
          "secondary-content": "#0f0115",
          accent: "#37cdbe",
          "accent-content": "#010f0d",
          neutral: "#262931",
          "neutral-content": "#cfd0d2",
          "base-100": "#1e3848",
          "base-200": "#192f3d",
          "base-300": "#132733",
          "base-content": "#ced4d8",
          info: "#2563EB",
          "info-content": "#d2e2ff",
          success: "#16A34A",
          "success-content": "#000a02",
          warning: "#D97706",
          "warning-content": "#110500",
          error: "#DC2626",
          "error-content": "#ffd9d4",
        },
      },
    ],
  },
  plugins: [require("daisyui")],
};
