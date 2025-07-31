// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTaskInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBatchTaskInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBatchTaskInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetBatchTaskInfoResponseBody) *GetBatchTaskInfoResponse
	GetBody() *GetBatchTaskInfoResponseBody
}

type GetBatchTaskInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBatchTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBatchTaskInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBatchTaskInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBatchTaskInfoResponse) GetBody() *GetBatchTaskInfoResponseBody {
	return s.Body
}

func (s *GetBatchTaskInfoResponse) SetHeaders(v map[string]*string) *GetBatchTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *GetBatchTaskInfoResponse) SetStatusCode(v int32) *GetBatchTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBatchTaskInfoResponse) SetBody(v *GetBatchTaskInfoResponseBody) *GetBatchTaskInfoResponse {
	s.Body = v
	return s
}

func (s *GetBatchTaskInfoResponse) Validate() error {
	return dara.Validate(s)
}
