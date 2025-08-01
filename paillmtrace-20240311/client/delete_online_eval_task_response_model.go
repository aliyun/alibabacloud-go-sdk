// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOnlineEvalTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOnlineEvalTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOnlineEvalTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOnlineEvalTaskResponseBody) *DeleteOnlineEvalTaskResponse
	GetBody() *DeleteOnlineEvalTaskResponseBody
}

type DeleteOnlineEvalTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOnlineEvalTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOnlineEvalTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOnlineEvalTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteOnlineEvalTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOnlineEvalTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOnlineEvalTaskResponse) GetBody() *DeleteOnlineEvalTaskResponseBody {
	return s.Body
}

func (s *DeleteOnlineEvalTaskResponse) SetHeaders(v map[string]*string) *DeleteOnlineEvalTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteOnlineEvalTaskResponse) SetStatusCode(v int32) *DeleteOnlineEvalTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOnlineEvalTaskResponse) SetBody(v *DeleteOnlineEvalTaskResponseBody) *DeleteOnlineEvalTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteOnlineEvalTaskResponse) Validate() error {
	return dara.Validate(s)
}
