// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTaskInfoByVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBatchTaskInfoByVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBatchTaskInfoByVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetBatchTaskInfoByVersionResponseBody) *GetBatchTaskInfoByVersionResponse
	GetBody() *GetBatchTaskInfoByVersionResponseBody
}

type GetBatchTaskInfoByVersionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBatchTaskInfoByVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBatchTaskInfoByVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskInfoByVersionResponse) GoString() string {
	return s.String()
}

func (s *GetBatchTaskInfoByVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBatchTaskInfoByVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBatchTaskInfoByVersionResponse) GetBody() *GetBatchTaskInfoByVersionResponseBody {
	return s.Body
}

func (s *GetBatchTaskInfoByVersionResponse) SetHeaders(v map[string]*string) *GetBatchTaskInfoByVersionResponse {
	s.Headers = v
	return s
}

func (s *GetBatchTaskInfoByVersionResponse) SetStatusCode(v int32) *GetBatchTaskInfoByVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBatchTaskInfoByVersionResponse) SetBody(v *GetBatchTaskInfoByVersionResponseBody) *GetBatchTaskInfoByVersionResponse {
	s.Body = v
	return s
}

func (s *GetBatchTaskInfoByVersionResponse) Validate() error {
	return dara.Validate(s)
}
