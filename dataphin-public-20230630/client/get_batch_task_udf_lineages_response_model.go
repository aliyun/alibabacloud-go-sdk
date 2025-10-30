// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTaskUdfLineagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBatchTaskUdfLineagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBatchTaskUdfLineagesResponse
	GetStatusCode() *int32
	SetBody(v *GetBatchTaskUdfLineagesResponseBody) *GetBatchTaskUdfLineagesResponse
	GetBody() *GetBatchTaskUdfLineagesResponseBody
}

type GetBatchTaskUdfLineagesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBatchTaskUdfLineagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBatchTaskUdfLineagesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskUdfLineagesResponse) GoString() string {
	return s.String()
}

func (s *GetBatchTaskUdfLineagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBatchTaskUdfLineagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBatchTaskUdfLineagesResponse) GetBody() *GetBatchTaskUdfLineagesResponseBody {
	return s.Body
}

func (s *GetBatchTaskUdfLineagesResponse) SetHeaders(v map[string]*string) *GetBatchTaskUdfLineagesResponse {
	s.Headers = v
	return s
}

func (s *GetBatchTaskUdfLineagesResponse) SetStatusCode(v int32) *GetBatchTaskUdfLineagesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBatchTaskUdfLineagesResponse) SetBody(v *GetBatchTaskUdfLineagesResponseBody) *GetBatchTaskUdfLineagesResponse {
	s.Body = v
	return s
}

func (s *GetBatchTaskUdfLineagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
