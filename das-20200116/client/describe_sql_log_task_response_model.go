// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlLogTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSqlLogTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSqlLogTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSqlLogTaskResponseBody) *DescribeSqlLogTaskResponse
	GetBody() *DescribeSqlLogTaskResponseBody
}

type DescribeSqlLogTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSqlLogTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSqlLogTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlLogTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeSqlLogTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSqlLogTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSqlLogTaskResponse) GetBody() *DescribeSqlLogTaskResponseBody {
	return s.Body
}

func (s *DescribeSqlLogTaskResponse) SetHeaders(v map[string]*string) *DescribeSqlLogTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeSqlLogTaskResponse) SetStatusCode(v int32) *DescribeSqlLogTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSqlLogTaskResponse) SetBody(v *DescribeSqlLogTaskResponseBody) *DescribeSqlLogTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeSqlLogTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
