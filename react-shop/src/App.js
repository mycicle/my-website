import logo from './logo.svg';
import products from "./data/products.json";
import Product from "./components/Product";
import './App.css';

function App() {
  return (
    <div className="bg-gray-200 dark:bg-gray-800 py-4">
      <div className="md:w-1/2 mx-auto">
        {products.map((product) => (
          <Product product={product} key={product.id} />
        ))}
      </div>
    </div>
  );
}

export default App;
