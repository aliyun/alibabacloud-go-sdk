// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBatchTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBatchTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBatchTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBatchTaskResponseBody) *DescribeBatchTaskResponse
	GetBody() *DescribeBatchTaskResponseBody
}

type DescribeBatchTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBatchTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBatchTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeBatchTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBatchTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBatchTaskResponse) GetBody() *DescribeBatchTaskResponseBody {
	return s.Body
}

func (s *DescribeBatchTaskResponse) SetHeaders(v map[string]*string) *DescribeBatchTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeBatchTaskResponse) SetStatusCode(v int32) *DescribeBatchTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBatchTaskResponse) SetBody(v *DescribeBatchTaskResponseBody) *DescribeBatchTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeBatchTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
