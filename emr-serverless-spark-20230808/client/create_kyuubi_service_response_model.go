// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKyuubiServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKyuubiServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKyuubiServiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateKyuubiServiceResponseBody) *CreateKyuubiServiceResponse
	GetBody() *CreateKyuubiServiceResponseBody
}

type CreateKyuubiServiceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKyuubiServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKyuubiServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKyuubiServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateKyuubiServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKyuubiServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKyuubiServiceResponse) GetBody() *CreateKyuubiServiceResponseBody {
	return s.Body
}

func (s *CreateKyuubiServiceResponse) SetHeaders(v map[string]*string) *CreateKyuubiServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateKyuubiServiceResponse) SetStatusCode(v int32) *CreateKyuubiServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKyuubiServiceResponse) SetBody(v *CreateKyuubiServiceResponseBody) *CreateKyuubiServiceResponse {
	s.Body = v
	return s
}

func (s *CreateKyuubiServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
