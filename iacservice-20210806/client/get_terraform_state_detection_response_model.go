// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTerraformStateDetectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTerraformStateDetectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTerraformStateDetectionResponse
	GetStatusCode() *int32
	SetBody(v *GetTerraformStateDetectionResponseBody) *GetTerraformStateDetectionResponse
	GetBody() *GetTerraformStateDetectionResponseBody
}

type GetTerraformStateDetectionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTerraformStateDetectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTerraformStateDetectionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTerraformStateDetectionResponse) GoString() string {
	return s.String()
}

func (s *GetTerraformStateDetectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTerraformStateDetectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTerraformStateDetectionResponse) GetBody() *GetTerraformStateDetectionResponseBody {
	return s.Body
}

func (s *GetTerraformStateDetectionResponse) SetHeaders(v map[string]*string) *GetTerraformStateDetectionResponse {
	s.Headers = v
	return s
}

func (s *GetTerraformStateDetectionResponse) SetStatusCode(v int32) *GetTerraformStateDetectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTerraformStateDetectionResponse) SetBody(v *GetTerraformStateDetectionResponseBody) *GetTerraformStateDetectionResponse {
	s.Body = v
	return s
}

func (s *GetTerraformStateDetectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
