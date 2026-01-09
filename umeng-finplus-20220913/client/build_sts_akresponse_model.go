// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildStsAKResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BuildStsAKResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BuildStsAKResponse
	GetStatusCode() *int32
	SetBody(v *BuildStsAKResponseBody) *BuildStsAKResponse
	GetBody() *BuildStsAKResponseBody
}

type BuildStsAKResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BuildStsAKResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BuildStsAKResponse) String() string {
	return dara.Prettify(s)
}

func (s BuildStsAKResponse) GoString() string {
	return s.String()
}

func (s *BuildStsAKResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BuildStsAKResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BuildStsAKResponse) GetBody() *BuildStsAKResponseBody {
	return s.Body
}

func (s *BuildStsAKResponse) SetHeaders(v map[string]*string) *BuildStsAKResponse {
	s.Headers = v
	return s
}

func (s *BuildStsAKResponse) SetStatusCode(v int32) *BuildStsAKResponse {
	s.StatusCode = &v
	return s
}

func (s *BuildStsAKResponse) SetBody(v *BuildStsAKResponseBody) *BuildStsAKResponse {
	s.Body = v
	return s
}

func (s *BuildStsAKResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
