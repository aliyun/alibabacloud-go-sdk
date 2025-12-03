// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySignatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySignatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySignatureResponse
	GetStatusCode() *int32
	SetBody(v *ModifySignatureResponseBody) *ModifySignatureResponse
	GetBody() *ModifySignatureResponseBody
}

type ModifySignatureResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySignatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySignatureResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySignatureResponse) GoString() string {
	return s.String()
}

func (s *ModifySignatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySignatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySignatureResponse) GetBody() *ModifySignatureResponseBody {
	return s.Body
}

func (s *ModifySignatureResponse) SetHeaders(v map[string]*string) *ModifySignatureResponse {
	s.Headers = v
	return s
}

func (s *ModifySignatureResponse) SetStatusCode(v int32) *ModifySignatureResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySignatureResponse) SetBody(v *ModifySignatureResponseBody) *ModifySignatureResponse {
	s.Body = v
	return s
}

func (s *ModifySignatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
