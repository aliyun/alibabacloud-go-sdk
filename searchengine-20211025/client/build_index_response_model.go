// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BuildIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BuildIndexResponse
	GetStatusCode() *int32
	SetBody(v *BuildIndexResponseBody) *BuildIndexResponse
	GetBody() *BuildIndexResponseBody
}

type BuildIndexResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BuildIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BuildIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s BuildIndexResponse) GoString() string {
	return s.String()
}

func (s *BuildIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BuildIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BuildIndexResponse) GetBody() *BuildIndexResponseBody {
	return s.Body
}

func (s *BuildIndexResponse) SetHeaders(v map[string]*string) *BuildIndexResponse {
	s.Headers = v
	return s
}

func (s *BuildIndexResponse) SetStatusCode(v int32) *BuildIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *BuildIndexResponse) SetBody(v *BuildIndexResponseBody) *BuildIndexResponse {
	s.Body = v
	return s
}

func (s *BuildIndexResponse) Validate() error {
	return dara.Validate(s)
}
