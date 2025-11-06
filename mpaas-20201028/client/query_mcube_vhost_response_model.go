// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMcubeVhostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMcubeVhostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMcubeVhostResponse
	GetStatusCode() *int32
	SetBody(v *QueryMcubeVhostResponseBody) *QueryMcubeVhostResponse
	GetBody() *QueryMcubeVhostResponseBody
}

type QueryMcubeVhostResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMcubeVhostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMcubeVhostResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMcubeVhostResponse) GoString() string {
	return s.String()
}

func (s *QueryMcubeVhostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMcubeVhostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMcubeVhostResponse) GetBody() *QueryMcubeVhostResponseBody {
	return s.Body
}

func (s *QueryMcubeVhostResponse) SetHeaders(v map[string]*string) *QueryMcubeVhostResponse {
	s.Headers = v
	return s
}

func (s *QueryMcubeVhostResponse) SetStatusCode(v int32) *QueryMcubeVhostResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMcubeVhostResponse) SetBody(v *QueryMcubeVhostResponseBody) *QueryMcubeVhostResponse {
	s.Body = v
	return s
}

func (s *QueryMcubeVhostResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
