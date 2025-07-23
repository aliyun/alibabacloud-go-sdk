// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkerResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWorkerResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWorkerResourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWorkerResourceResponseBody) *DeleteWorkerResourceResponse
	GetBody() *DeleteWorkerResourceResponseBody
}

type DeleteWorkerResourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkerResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkerResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkerResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkerResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWorkerResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWorkerResourceResponse) GetBody() *DeleteWorkerResourceResponseBody {
	return s.Body
}

func (s *DeleteWorkerResourceResponse) SetHeaders(v map[string]*string) *DeleteWorkerResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkerResourceResponse) SetStatusCode(v int32) *DeleteWorkerResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkerResourceResponse) SetBody(v *DeleteWorkerResourceResponseBody) *DeleteWorkerResourceResponse {
	s.Body = v
	return s
}

func (s *DeleteWorkerResourceResponse) Validate() error {
	return dara.Validate(s)
}
