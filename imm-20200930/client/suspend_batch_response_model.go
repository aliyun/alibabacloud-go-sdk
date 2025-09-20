// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendBatchResponse
	GetStatusCode() *int32
	SetBody(v *SuspendBatchResponseBody) *SuspendBatchResponse
	GetBody() *SuspendBatchResponseBody
}

type SuspendBatchResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendBatchResponse) GoString() string {
	return s.String()
}

func (s *SuspendBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendBatchResponse) GetBody() *SuspendBatchResponseBody {
	return s.Body
}

func (s *SuspendBatchResponse) SetHeaders(v map[string]*string) *SuspendBatchResponse {
	s.Headers = v
	return s
}

func (s *SuspendBatchResponse) SetStatusCode(v int32) *SuspendBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendBatchResponse) SetBody(v *SuspendBatchResponseBody) *SuspendBatchResponse {
	s.Body = v
	return s
}

func (s *SuspendBatchResponse) Validate() error {
	return dara.Validate(s)
}
