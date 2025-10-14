// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlFlashbackTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSqlFlashbackTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSqlFlashbackTaskListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSqlFlashbackTaskListResponseBody) *DescribeSqlFlashbackTaskListResponse
	GetBody() *DescribeSqlFlashbackTaskListResponseBody
}

type DescribeSqlFlashbackTaskListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSqlFlashbackTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSqlFlashbackTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlFlashbackTaskListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbackTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSqlFlashbackTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSqlFlashbackTaskListResponse) GetBody() *DescribeSqlFlashbackTaskListResponseBody {
	return s.Body
}

func (s *DescribeSqlFlashbackTaskListResponse) SetHeaders(v map[string]*string) *DescribeSqlFlashbackTaskListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponse) SetStatusCode(v int32) *DescribeSqlFlashbackTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponse) SetBody(v *DescribeSqlFlashbackTaskListResponseBody) *DescribeSqlFlashbackTaskListResponse {
	s.Body = v
	return s
}

func (s *DescribeSqlFlashbackTaskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
