// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLabelBuildResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LabelBuildResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LabelBuildResponse
	GetStatusCode() *int32
	SetBody(v *LabelBuildResponseBody) *LabelBuildResponse
	GetBody() *LabelBuildResponseBody
}

type LabelBuildResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LabelBuildResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LabelBuildResponse) String() string {
	return dara.Prettify(s)
}

func (s LabelBuildResponse) GoString() string {
	return s.String()
}

func (s *LabelBuildResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LabelBuildResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LabelBuildResponse) GetBody() *LabelBuildResponseBody {
	return s.Body
}

func (s *LabelBuildResponse) SetHeaders(v map[string]*string) *LabelBuildResponse {
	s.Headers = v
	return s
}

func (s *LabelBuildResponse) SetStatusCode(v int32) *LabelBuildResponse {
	s.StatusCode = &v
	return s
}

func (s *LabelBuildResponse) SetBody(v *LabelBuildResponseBody) *LabelBuildResponse {
	s.Body = v
	return s
}

func (s *LabelBuildResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
