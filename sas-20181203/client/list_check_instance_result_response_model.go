// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckInstanceResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCheckInstanceResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCheckInstanceResultResponse
	GetStatusCode() *int32
	SetBody(v *ListCheckInstanceResultResponseBody) *ListCheckInstanceResultResponse
	GetBody() *ListCheckInstanceResultResponseBody
}

type ListCheckInstanceResultResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCheckInstanceResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCheckInstanceResultResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCheckInstanceResultResponse) GoString() string {
	return s.String()
}

func (s *ListCheckInstanceResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCheckInstanceResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCheckInstanceResultResponse) GetBody() *ListCheckInstanceResultResponseBody {
	return s.Body
}

func (s *ListCheckInstanceResultResponse) SetHeaders(v map[string]*string) *ListCheckInstanceResultResponse {
	s.Headers = v
	return s
}

func (s *ListCheckInstanceResultResponse) SetStatusCode(v int32) *ListCheckInstanceResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCheckInstanceResultResponse) SetBody(v *ListCheckInstanceResultResponseBody) *ListCheckInstanceResultResponse {
	s.Body = v
	return s
}

func (s *ListCheckInstanceResultResponse) Validate() error {
	return dara.Validate(s)
}
