// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageTerraformStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ManageTerraformStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ManageTerraformStateResponse
	GetStatusCode() *int32
	SetBody(v *ManageTerraformStateResponseBody) *ManageTerraformStateResponse
	GetBody() *ManageTerraformStateResponseBody
}

type ManageTerraformStateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManageTerraformStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManageTerraformStateResponse) String() string {
	return dara.Prettify(s)
}

func (s ManageTerraformStateResponse) GoString() string {
	return s.String()
}

func (s *ManageTerraformStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ManageTerraformStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ManageTerraformStateResponse) GetBody() *ManageTerraformStateResponseBody {
	return s.Body
}

func (s *ManageTerraformStateResponse) SetHeaders(v map[string]*string) *ManageTerraformStateResponse {
	s.Headers = v
	return s
}

func (s *ManageTerraformStateResponse) SetStatusCode(v int32) *ManageTerraformStateResponse {
	s.StatusCode = &v
	return s
}

func (s *ManageTerraformStateResponse) SetBody(v *ManageTerraformStateResponseBody) *ManageTerraformStateResponse {
	s.Body = v
	return s
}

func (s *ManageTerraformStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
