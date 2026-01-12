// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBuildBreakpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryBuildBreakpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryBuildBreakpointResponse
	GetStatusCode() *int32
	SetBody(v *QueryBuildBreakpointResponseBody) *QueryBuildBreakpointResponse
	GetBody() *QueryBuildBreakpointResponseBody
}

type QueryBuildBreakpointResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBuildBreakpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBuildBreakpointResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryBuildBreakpointResponse) GoString() string {
	return s.String()
}

func (s *QueryBuildBreakpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryBuildBreakpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryBuildBreakpointResponse) GetBody() *QueryBuildBreakpointResponseBody {
	return s.Body
}

func (s *QueryBuildBreakpointResponse) SetHeaders(v map[string]*string) *QueryBuildBreakpointResponse {
	s.Headers = v
	return s
}

func (s *QueryBuildBreakpointResponse) SetStatusCode(v int32) *QueryBuildBreakpointResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBuildBreakpointResponse) SetBody(v *QueryBuildBreakpointResponseBody) *QueryBuildBreakpointResponse {
	s.Body = v
	return s
}

func (s *QueryBuildBreakpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
