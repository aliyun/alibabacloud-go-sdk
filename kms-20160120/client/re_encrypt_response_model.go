// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReEncryptResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReEncryptResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReEncryptResponse
	GetStatusCode() *int32
	SetBody(v *ReEncryptResponseBody) *ReEncryptResponse
	GetBody() *ReEncryptResponseBody
}

type ReEncryptResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReEncryptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReEncryptResponse) String() string {
	return dara.Prettify(s)
}

func (s ReEncryptResponse) GoString() string {
	return s.String()
}

func (s *ReEncryptResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReEncryptResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReEncryptResponse) GetBody() *ReEncryptResponseBody {
	return s.Body
}

func (s *ReEncryptResponse) SetHeaders(v map[string]*string) *ReEncryptResponse {
	s.Headers = v
	return s
}

func (s *ReEncryptResponse) SetStatusCode(v int32) *ReEncryptResponse {
	s.StatusCode = &v
	return s
}

func (s *ReEncryptResponse) SetBody(v *ReEncryptResponseBody) *ReEncryptResponse {
	s.Body = v
	return s
}

func (s *ReEncryptResponse) Validate() error {
	return dara.Validate(s)
}
