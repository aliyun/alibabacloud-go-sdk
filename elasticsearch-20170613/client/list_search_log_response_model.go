// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSearchLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSearchLogResponse
	GetStatusCode() *int32
	SetBody(v *ListSearchLogResponseBody) *ListSearchLogResponse
	GetBody() *ListSearchLogResponseBody
}

type ListSearchLogResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSearchLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSearchLogResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSearchLogResponse) GoString() string {
	return s.String()
}

func (s *ListSearchLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSearchLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSearchLogResponse) GetBody() *ListSearchLogResponseBody {
	return s.Body
}

func (s *ListSearchLogResponse) SetHeaders(v map[string]*string) *ListSearchLogResponse {
	s.Headers = v
	return s
}

func (s *ListSearchLogResponse) SetStatusCode(v int32) *ListSearchLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSearchLogResponse) SetBody(v *ListSearchLogResponseBody) *ListSearchLogResponse {
	s.Body = v
	return s
}

func (s *ListSearchLogResponse) Validate() error {
	return dara.Validate(s)
}
