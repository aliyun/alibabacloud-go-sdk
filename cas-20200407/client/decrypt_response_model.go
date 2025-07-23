// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDecryptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DecryptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DecryptResponse
	GetStatusCode() *int32
	SetBody(v *DecryptResponseBody) *DecryptResponse
	GetBody() *DecryptResponseBody
}

type DecryptResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DecryptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DecryptResponse) String() string {
	return dara.Prettify(s)
}

func (s DecryptResponse) GoString() string {
	return s.String()
}

func (s *DecryptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DecryptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DecryptResponse) GetBody() *DecryptResponseBody {
	return s.Body
}

func (s *DecryptResponse) SetHeaders(v map[string]*string) *DecryptResponse {
	s.Headers = v
	return s
}

func (s *DecryptResponse) SetStatusCode(v int32) *DecryptResponse {
	s.StatusCode = &v
	return s
}

func (s *DecryptResponse) SetBody(v *DecryptResponseBody) *DecryptResponse {
	s.Body = v
	return s
}

func (s *DecryptResponse) Validate() error {
	return dara.Validate(s)
}
