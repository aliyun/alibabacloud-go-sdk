// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineBatchTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OfflineBatchTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OfflineBatchTaskResponse
	GetStatusCode() *int32
	SetBody(v *OfflineBatchTaskResponseBody) *OfflineBatchTaskResponse
	GetBody() *OfflineBatchTaskResponseBody
}

type OfflineBatchTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineBatchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineBatchTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s OfflineBatchTaskResponse) GoString() string {
	return s.String()
}

func (s *OfflineBatchTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OfflineBatchTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OfflineBatchTaskResponse) GetBody() *OfflineBatchTaskResponseBody {
	return s.Body
}

func (s *OfflineBatchTaskResponse) SetHeaders(v map[string]*string) *OfflineBatchTaskResponse {
	s.Headers = v
	return s
}

func (s *OfflineBatchTaskResponse) SetStatusCode(v int32) *OfflineBatchTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineBatchTaskResponse) SetBody(v *OfflineBatchTaskResponseBody) *OfflineBatchTaskResponse {
	s.Body = v
	return s
}

func (s *OfflineBatchTaskResponse) Validate() error {
	return dara.Validate(s)
}
