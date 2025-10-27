// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarPlaybookTaskDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSoarPlaybookTaskDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSoarPlaybookTaskDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSoarPlaybookTaskDetailResponseBody) *DescribeSoarPlaybookTaskDetailResponse
	GetBody() *DescribeSoarPlaybookTaskDetailResponseBody
}

type DescribeSoarPlaybookTaskDetailResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSoarPlaybookTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSoarPlaybookTaskDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarPlaybookTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeSoarPlaybookTaskDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSoarPlaybookTaskDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSoarPlaybookTaskDetailResponse) GetBody() *DescribeSoarPlaybookTaskDetailResponseBody {
	return s.Body
}

func (s *DescribeSoarPlaybookTaskDetailResponse) SetHeaders(v map[string]*string) *DescribeSoarPlaybookTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeSoarPlaybookTaskDetailResponse) SetStatusCode(v int32) *DescribeSoarPlaybookTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSoarPlaybookTaskDetailResponse) SetBody(v *DescribeSoarPlaybookTaskDetailResponseBody) *DescribeSoarPlaybookTaskDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeSoarPlaybookTaskDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
