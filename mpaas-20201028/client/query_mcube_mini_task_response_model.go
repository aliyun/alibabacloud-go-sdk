// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMcubeMiniTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMcubeMiniTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMcubeMiniTaskResponse
	GetStatusCode() *int32
	SetBody(v *QueryMcubeMiniTaskResponseBody) *QueryMcubeMiniTaskResponse
	GetBody() *QueryMcubeMiniTaskResponseBody
}

type QueryMcubeMiniTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMcubeMiniTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMcubeMiniTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMcubeMiniTaskResponse) GoString() string {
	return s.String()
}

func (s *QueryMcubeMiniTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMcubeMiniTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMcubeMiniTaskResponse) GetBody() *QueryMcubeMiniTaskResponseBody {
	return s.Body
}

func (s *QueryMcubeMiniTaskResponse) SetHeaders(v map[string]*string) *QueryMcubeMiniTaskResponse {
	s.Headers = v
	return s
}

func (s *QueryMcubeMiniTaskResponse) SetStatusCode(v int32) *QueryMcubeMiniTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMcubeMiniTaskResponse) SetBody(v *QueryMcubeMiniTaskResponseBody) *QueryMcubeMiniTaskResponse {
	s.Body = v
	return s
}

func (s *QueryMcubeMiniTaskResponse) Validate() error {
	return dara.Validate(s)
}
