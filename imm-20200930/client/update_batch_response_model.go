// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBatchResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBatchResponseBody) *UpdateBatchResponse
	GetBody() *UpdateBatchResponseBody
}

type UpdateBatchResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBatchResponse) GoString() string {
	return s.String()
}

func (s *UpdateBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBatchResponse) GetBody() *UpdateBatchResponseBody {
	return s.Body
}

func (s *UpdateBatchResponse) SetHeaders(v map[string]*string) *UpdateBatchResponse {
	s.Headers = v
	return s
}

func (s *UpdateBatchResponse) SetStatusCode(v int32) *UpdateBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBatchResponse) SetBody(v *UpdateBatchResponseBody) *UpdateBatchResponse {
	s.Body = v
	return s
}

func (s *UpdateBatchResponse) Validate() error {
	return dara.Validate(s)
}
