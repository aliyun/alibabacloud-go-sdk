// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVCUInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVCUInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVCUInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateVCUInstanceResponseBody) *CreateVCUInstanceResponse
	GetBody() *CreateVCUInstanceResponseBody
}

type CreateVCUInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVCUInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVCUInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVCUInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateVCUInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVCUInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVCUInstanceResponse) GetBody() *CreateVCUInstanceResponseBody {
	return s.Body
}

func (s *CreateVCUInstanceResponse) SetHeaders(v map[string]*string) *CreateVCUInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateVCUInstanceResponse) SetStatusCode(v int32) *CreateVCUInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVCUInstanceResponse) SetBody(v *CreateVCUInstanceResponseBody) *CreateVCUInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateVCUInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
