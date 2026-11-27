// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumSymbolFileParamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFormData(v map[string]*string) *GetRumSymbolFileParamsResponseBody
	GetFormData() map[string]*string
	SetRequestId(v string) *GetRumSymbolFileParamsResponseBody
	GetRequestId() *string
	SetUploadUrl(v string) *GetRumSymbolFileParamsResponseBody
	GetUploadUrl() *string
}

type GetRumSymbolFileParamsResponseBody struct {
	// The collection of signature parameters required for form upload, including key, policy, x-oss-signature-version, x-oss-credential, x-oss-date, and x-oss-signature. The signatures in the examples are for illustration purposes only and cannot be used for actual uploads. Use the complete parameters returned by the API operation.
	//
	// example:
	//
	// {"key":"user/workspace/service-id/1.0.0/build-1/bundle.js.map","policy":"eyJleHBpcmF0aW9uIjoiMjAyNi0wNi0wM1QxMjozMDowMC4wMDBaIiwiY29uZGl0aW9ucyI6W3siYnVja2V0IjoiYnVja2V0In0seyJ4LW9zcy1zaWduYXR1cmUtdmVyc2lvbiI6Ik9TUzQtSE1BQy1TSEEyNTYifSx7Ingtb3NzLWNyZWRlbnRpYWwiOiJBS0lERVhBTVBMRS8yMDI2MDYwMy9jbi1oYW5nemhvdS9vc3MvYWxpeXVuX3Y0X3JlcXVlc3QifSx7Ingtb3NzLWRhdGUiOiIyMDI2MDYwM1QxMjAwMDBaIn0sWyJlcSIsIiRrZXkiLCJ1c2VyL3dvcmtzcGFjZS9zZXJ2aWNlLWlkLzEuMC4wL2J1aWxkLTEvYnVuZGxlLmpzLm1hcCJdLFsiY29udGVudC1sZW5ndGgtcmFuZ2UiLDEsNTI0Mjg4MDAwXV19","x-oss-signature-version":"OSS4-HMAC-SHA256","x-oss-credential":"AKIDEXAMPLE/20260603/cn-hangzhou/oss/aliyun_v4_request","x-oss-date":"20260603T120000Z","x-oss-signature":"0000000000000000000000000000000000000000000000000000000000000000"}
	FormData map[string]*string `json:"formData,omitempty" xml:"formData,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// request-1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The form upload URL for directly uploading symbol files to OSS.
	//
	// example:
	//
	// https://bucket.oss-cn-hangzhou.aliyuncs.com
	UploadUrl *string `json:"uploadUrl,omitempty" xml:"uploadUrl,omitempty"`
}

func (s GetRumSymbolFileParamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRumSymbolFileParamsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRumSymbolFileParamsResponseBody) GetFormData() map[string]*string {
	return s.FormData
}

func (s *GetRumSymbolFileParamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRumSymbolFileParamsResponseBody) GetUploadUrl() *string {
	return s.UploadUrl
}

func (s *GetRumSymbolFileParamsResponseBody) SetFormData(v map[string]*string) *GetRumSymbolFileParamsResponseBody {
	s.FormData = v
	return s
}

func (s *GetRumSymbolFileParamsResponseBody) SetRequestId(v string) *GetRumSymbolFileParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRumSymbolFileParamsResponseBody) SetUploadUrl(v string) *GetRumSymbolFileParamsResponseBody {
	s.UploadUrl = &v
	return s
}

func (s *GetRumSymbolFileParamsResponseBody) Validate() error {
	return dara.Validate(s)
}
