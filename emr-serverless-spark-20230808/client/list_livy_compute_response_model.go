// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLivyComputeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLivyComputeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLivyComputeResponse
	GetStatusCode() *int32
	SetBody(v *ListLivyComputeResponseBody) *ListLivyComputeResponse
	GetBody() *ListLivyComputeResponseBody
}

type ListLivyComputeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLivyComputeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLivyComputeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLivyComputeResponse) GoString() string {
	return s.String()
}

func (s *ListLivyComputeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLivyComputeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLivyComputeResponse) GetBody() *ListLivyComputeResponseBody {
	return s.Body
}

func (s *ListLivyComputeResponse) SetHeaders(v map[string]*string) *ListLivyComputeResponse {
	s.Headers = v
	return s
}

func (s *ListLivyComputeResponse) SetStatusCode(v int32) *ListLivyComputeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLivyComputeResponse) SetBody(v *ListLivyComputeResponseBody) *ListLivyComputeResponse {
	s.Body = v
	return s
}

func (s *ListLivyComputeResponse) Validate() error {
	return dara.Validate(s)
}
