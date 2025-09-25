// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebuildTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebuildTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebuildTaskResponse
	GetStatusCode() *int32
	SetBody(v *RebuildTaskResponseBody) *RebuildTaskResponse
	GetBody() *RebuildTaskResponseBody
}

type RebuildTaskResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebuildTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebuildTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s RebuildTaskResponse) GoString() string {
	return s.String()
}

func (s *RebuildTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebuildTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebuildTaskResponse) GetBody() *RebuildTaskResponseBody {
	return s.Body
}

func (s *RebuildTaskResponse) SetHeaders(v map[string]*string) *RebuildTaskResponse {
	s.Headers = v
	return s
}

func (s *RebuildTaskResponse) SetStatusCode(v int32) *RebuildTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RebuildTaskResponse) SetBody(v *RebuildTaskResponseBody) *RebuildTaskResponse {
	s.Body = v
	return s
}

func (s *RebuildTaskResponse) Validate() error {
	return dara.Validate(s)
}
