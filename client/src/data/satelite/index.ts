import Api from "../Api";

export class SateliteApi {
  private api: Api;

  constructor() {
    this.api = new Api();
  }

  getAll() {
    this.api.get("/");
  }
}
