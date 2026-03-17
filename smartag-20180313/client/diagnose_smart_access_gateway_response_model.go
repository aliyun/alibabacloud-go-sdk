// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiagnoseSmartAccessGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DiagnoseSmartAccessGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DiagnoseSmartAccessGatewayResponse
	GetStatusCode() *int32
	SetBody(v *DiagnoseSmartAccessGatewayResponseBody) *DiagnoseSmartAccessGatewayResponse
	GetBody() *DiagnoseSmartAccessGatewayResponseBody
}

type DiagnoseSmartAccessGatewayResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DiagnoseSmartAccessGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DiagnoseSmartAccessGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DiagnoseSmartAccessGatewayResponse) GoString() string {
	return s.String()
}

func (s *DiagnoseSmartAccessGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DiagnoseSmartAccessGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DiagnoseSmartAccessGatewayResponse) GetBody() *DiagnoseSmartAccessGatewayResponseBody {
	return s.Body
}

func (s *DiagnoseSmartAccessGatewayResponse) SetHeaders(v map[string]*string) *DiagnoseSmartAccessGatewayResponse {
	s.Headers = v
	return s
}

func (s *DiagnoseSmartAccessGatewayResponse) SetStatusCode(v int32) *DiagnoseSmartAccessGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DiagnoseSmartAccessGatewayResponse) SetBody(v *DiagnoseSmartAccessGatewayResponseBody) *DiagnoseSmartAccessGatewayResponse {
	s.Body = v
	return s
}

func (s *DiagnoseSmartAccessGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
