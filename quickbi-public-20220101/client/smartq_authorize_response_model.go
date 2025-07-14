// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmartqAuthorizeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SmartqAuthorizeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SmartqAuthorizeResponse
	GetStatusCode() *int32
	SetBody(v *SmartqAuthorizeResponseBody) *SmartqAuthorizeResponse
	GetBody() *SmartqAuthorizeResponseBody
}

type SmartqAuthorizeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SmartqAuthorizeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SmartqAuthorizeResponse) String() string {
	return dara.Prettify(s)
}

func (s SmartqAuthorizeResponse) GoString() string {
	return s.String()
}

func (s *SmartqAuthorizeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SmartqAuthorizeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SmartqAuthorizeResponse) GetBody() *SmartqAuthorizeResponseBody {
	return s.Body
}

func (s *SmartqAuthorizeResponse) SetHeaders(v map[string]*string) *SmartqAuthorizeResponse {
	s.Headers = v
	return s
}

func (s *SmartqAuthorizeResponse) SetStatusCode(v int32) *SmartqAuthorizeResponse {
	s.StatusCode = &v
	return s
}

func (s *SmartqAuthorizeResponse) SetBody(v *SmartqAuthorizeResponseBody) *SmartqAuthorizeResponse {
	s.Body = v
	return s
}

func (s *SmartqAuthorizeResponse) Validate() error {
	return dara.Validate(s)
}
