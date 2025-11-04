// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHighCodeDeployResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HighCodeDeployResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HighCodeDeployResponse
	GetStatusCode() *int32
	SetBody(v *HighCodeDeployResponseBody) *HighCodeDeployResponse
	GetBody() *HighCodeDeployResponseBody
}

type HighCodeDeployResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HighCodeDeployResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HighCodeDeployResponse) String() string {
	return dara.Prettify(s)
}

func (s HighCodeDeployResponse) GoString() string {
	return s.String()
}

func (s *HighCodeDeployResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HighCodeDeployResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HighCodeDeployResponse) GetBody() *HighCodeDeployResponseBody {
	return s.Body
}

func (s *HighCodeDeployResponse) SetHeaders(v map[string]*string) *HighCodeDeployResponse {
	s.Headers = v
	return s
}

func (s *HighCodeDeployResponse) SetStatusCode(v int32) *HighCodeDeployResponse {
	s.StatusCode = &v
	return s
}

func (s *HighCodeDeployResponse) SetBody(v *HighCodeDeployResponseBody) *HighCodeDeployResponse {
	s.Body = v
	return s
}

func (s *HighCodeDeployResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
