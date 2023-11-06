import { getCookie } from 'cookies-next';

function getThemeCookiesValue() {
  try {
    const ThemeValue = getCookie('theme');
    console.log(ThemeValue);
    return JSON.parse(ThemeValue);
  } catch (err) {
    return JSON.stringify({
      primary_color: '',
      secondary_color: '',
    });
  }
}

export default getThemeCookiesValue;
