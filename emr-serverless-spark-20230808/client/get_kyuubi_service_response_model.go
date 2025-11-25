// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKyuubiServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKyuubiServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKyuubiServiceResponse
	GetStatusCode() *int32
	SetBody(v *GetKyuubiServiceResponseBody) *GetKyuubiServiceResponse
	GetBody() *GetKyuubiServiceResponseBody
}

type GetKyuubiServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKyuubiServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKyuubiServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKyuubiServiceResponse) GoString() string {
	return s.String()
}

func (s *GetKyuubiServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKyuubiServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKyuubiServiceResponse) GetBody() *GetKyuubiServiceResponseBody {
	return s.Body
}

func (s *GetKyuubiServiceResponse) SetHeaders(v map[string]*string) *GetKyuubiServiceResponse {
	s.Headers = v
	return s
}

func (s *GetKyuubiServiceResponse) SetStatusCode(v int32) *GetKyuubiServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKyuubiServiceResponse) SetBody(v *GetKyuubiServiceResponseBody) *GetKyuubiServiceResponse {
	s.Body = v
	return s
}

func (s *GetKyuubiServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
