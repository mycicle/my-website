const colors = require("tailwindcss/colors");

module.exports = {
  purge: ["./sec/**/*.{js,jsx,ts,tsx}", "./public/index.html"],
  darkMode: 'media', // or 'media' or 'class'
  theme: {
    extend: {
      colors: {
        gray: colors.blueGray,
        red: colors.rose,
      },
    },
  },
  variants: {
    extend: {},
  },
  plugins: [],
}

/*
  purge: this is used to specify the files Tailwind CSS should scan and see which Tailwind CSS classes are being used, so that it can remove all unused styles and classes in production.
  darkMode: this specifies the behavior of dark mode in your project. The value can be media — meaning that the dark mode styling will be applied based on the dark mode media query, which depends on the default mode for the user’s OS. It can also be class, which means that the dark mode styling will be applied when a parent element in the HTML document has the dark class.
  theme: this can be used to modify the color palette of the theme, fonts, breakpoints, and more. We’ll see how to make changes to the theme later in the tutorial.
  variants: this allows you to apply additional variants to the core plugins of Tailwind CSS. We’ll see how it works later on in the tutorial.
  plugins: a section to add plugins that can add extra utility classes, custom variants, base styles or more.
*/