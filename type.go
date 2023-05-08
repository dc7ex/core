package copyright

const (
	//kSandboxURL = "http://dev-openapi.360ex.art:8067/api/v1/open"
	kSandboxURL       = "http://dev-openapi.360ex.art/api/v1/open"
	kProductionURL    = "https://openapi-prod.cex.art/api/v1/open"
	kContentType      = "application/json;charset=utf-8"
	kSignNodeName     = "sign"
	HeaderContentType = "Content-Type"
	defaultMemory     = 32 << 20 // 32 MB

	MIMEApplicationJSON = "application/json"
	MIMEApplicationForm = "application/x-www-form-urlencoded"
	MIMEMultipartForm   = "multipart/form-data"
)

type Request struct {
	AppId     string `json:"app_id"`
	Method    string `json:"method"`
	Nonce     string `json:"nonce"`
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
	Data      string `json:"data"`
}

type Response struct {
	Code      Code   `json:"code"`
	Message   string `json:"message"`
	Method    string `json:"method"`
	Nonce     string `json:"nonce"`
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
	Data      string `json:"data"`
}
