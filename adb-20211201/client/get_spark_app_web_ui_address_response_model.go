// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkAppWebUiAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSparkAppWebUiAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSparkAppWebUiAddressResponse
	GetStatusCode() *int32
	SetBody(v *GetSparkAppWebUiAddressResponseBody) *GetSparkAppWebUiAddressResponse
	GetBody() *GetSparkAppWebUiAddressResponseBody
}

type GetSparkAppWebUiAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkAppWebUiAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkAppWebUiAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppWebUiAddressResponse) GoString() string {
	return s.String()
}

func (s *GetSparkAppWebUiAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSparkAppWebUiAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSparkAppWebUiAddressResponse) GetBody() *GetSparkAppWebUiAddressResponseBody {
	return s.Body
}

func (s *GetSparkAppWebUiAddressResponse) SetHeaders(v map[string]*string) *GetSparkAppWebUiAddressResponse {
	s.Headers = v
	return s
}

func (s *GetSparkAppWebUiAddressResponse) SetStatusCode(v int32) *GetSparkAppWebUiAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkAppWebUiAddressResponse) SetBody(v *GetSparkAppWebUiAddressResponseBody) *GetSparkAppWebUiAddressResponse {
	s.Body = v
	return s
}

func (s *GetSparkAppWebUiAddressResponse) Validate() error {
	return dara.Validate(s)
}
