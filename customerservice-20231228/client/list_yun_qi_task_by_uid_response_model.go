// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListYunQiTaskByUidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListYunQiTaskByUidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListYunQiTaskByUidResponse
	GetStatusCode() *int32
	SetBody(v *ListYunQiTaskByUidResponseBody) *ListYunQiTaskByUidResponse
	GetBody() *ListYunQiTaskByUidResponseBody
}

type ListYunQiTaskByUidResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListYunQiTaskByUidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListYunQiTaskByUidResponse) String() string {
	return dara.Prettify(s)
}

func (s ListYunQiTaskByUidResponse) GoString() string {
	return s.String()
}

func (s *ListYunQiTaskByUidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListYunQiTaskByUidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListYunQiTaskByUidResponse) GetBody() *ListYunQiTaskByUidResponseBody {
	return s.Body
}

func (s *ListYunQiTaskByUidResponse) SetHeaders(v map[string]*string) *ListYunQiTaskByUidResponse {
	s.Headers = v
	return s
}

func (s *ListYunQiTaskByUidResponse) SetStatusCode(v int32) *ListYunQiTaskByUidResponse {
	s.StatusCode = &v
	return s
}

func (s *ListYunQiTaskByUidResponse) SetBody(v *ListYunQiTaskByUidResponseBody) *ListYunQiTaskByUidResponse {
	s.Body = v
	return s
}

func (s *ListYunQiTaskByUidResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
