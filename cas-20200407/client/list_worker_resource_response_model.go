// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkerResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWorkerResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWorkerResourceResponse
	GetStatusCode() *int32
	SetBody(v *ListWorkerResourceResponseBody) *ListWorkerResourceResponse
	GetBody() *ListWorkerResourceResponseBody
}

type ListWorkerResourceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkerResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkerResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWorkerResourceResponse) GoString() string {
	return s.String()
}

func (s *ListWorkerResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWorkerResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWorkerResourceResponse) GetBody() *ListWorkerResourceResponseBody {
	return s.Body
}

func (s *ListWorkerResourceResponse) SetHeaders(v map[string]*string) *ListWorkerResourceResponse {
	s.Headers = v
	return s
}

func (s *ListWorkerResourceResponse) SetStatusCode(v int32) *ListWorkerResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkerResourceResponse) SetBody(v *ListWorkerResourceResponseBody) *ListWorkerResourceResponse {
	s.Body = v
	return s
}

func (s *ListWorkerResourceResponse) Validate() error {
	return dara.Validate(s)
}
