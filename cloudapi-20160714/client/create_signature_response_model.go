// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSignatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSignatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSignatureResponse
	GetStatusCode() *int32
	SetBody(v *CreateSignatureResponseBody) *CreateSignatureResponse
	GetBody() *CreateSignatureResponseBody
}

type CreateSignatureResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSignatureResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSignatureResponse) GoString() string {
	return s.String()
}

func (s *CreateSignatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSignatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSignatureResponse) GetBody() *CreateSignatureResponseBody {
	return s.Body
}

func (s *CreateSignatureResponse) SetHeaders(v map[string]*string) *CreateSignatureResponse {
	s.Headers = v
	return s
}

func (s *CreateSignatureResponse) SetStatusCode(v int32) *CreateSignatureResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSignatureResponse) SetBody(v *CreateSignatureResponseBody) *CreateSignatureResponse {
	s.Body = v
	return s
}

func (s *CreateSignatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
