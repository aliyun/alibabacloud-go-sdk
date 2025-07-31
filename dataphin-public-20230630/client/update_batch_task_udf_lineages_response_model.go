// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBatchTaskUdfLineagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBatchTaskUdfLineagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBatchTaskUdfLineagesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBatchTaskUdfLineagesResponseBody) *UpdateBatchTaskUdfLineagesResponse
	GetBody() *UpdateBatchTaskUdfLineagesResponseBody
}

type UpdateBatchTaskUdfLineagesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBatchTaskUdfLineagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBatchTaskUdfLineagesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchTaskUdfLineagesResponse) GoString() string {
	return s.String()
}

func (s *UpdateBatchTaskUdfLineagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBatchTaskUdfLineagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBatchTaskUdfLineagesResponse) GetBody() *UpdateBatchTaskUdfLineagesResponseBody {
	return s.Body
}

func (s *UpdateBatchTaskUdfLineagesResponse) SetHeaders(v map[string]*string) *UpdateBatchTaskUdfLineagesResponse {
	s.Headers = v
	return s
}

func (s *UpdateBatchTaskUdfLineagesResponse) SetStatusCode(v int32) *UpdateBatchTaskUdfLineagesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBatchTaskUdfLineagesResponse) SetBody(v *UpdateBatchTaskUdfLineagesResponseBody) *UpdateBatchTaskUdfLineagesResponse {
	s.Body = v
	return s
}

func (s *UpdateBatchTaskUdfLineagesResponse) Validate() error {
	return dara.Validate(s)
}
