// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectTerraformStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectTerraformStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectTerraformStateResponse
	GetStatusCode() *int32
	SetBody(v *DetectTerraformStateResponseBody) *DetectTerraformStateResponse
	GetBody() *DetectTerraformStateResponseBody
}

type DetectTerraformStateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectTerraformStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectTerraformStateResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectTerraformStateResponse) GoString() string {
	return s.String()
}

func (s *DetectTerraformStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectTerraformStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectTerraformStateResponse) GetBody() *DetectTerraformStateResponseBody {
	return s.Body
}

func (s *DetectTerraformStateResponse) SetHeaders(v map[string]*string) *DetectTerraformStateResponse {
	s.Headers = v
	return s
}

func (s *DetectTerraformStateResponse) SetStatusCode(v int32) *DetectTerraformStateResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectTerraformStateResponse) SetBody(v *DetectTerraformStateResponseBody) *DetectTerraformStateResponse {
	s.Body = v
	return s
}

func (s *DetectTerraformStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
