// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunHotTopicSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunHotTopicSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunHotTopicSummaryResponse
	GetStatusCode() *int32
	SetBody(v *RunHotTopicSummaryResponseBody) *RunHotTopicSummaryResponse
	GetBody() *RunHotTopicSummaryResponseBody
}

type RunHotTopicSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunHotTopicSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunHotTopicSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicSummaryResponse) GoString() string {
	return s.String()
}

func (s *RunHotTopicSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunHotTopicSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunHotTopicSummaryResponse) GetBody() *RunHotTopicSummaryResponseBody {
	return s.Body
}

func (s *RunHotTopicSummaryResponse) SetHeaders(v map[string]*string) *RunHotTopicSummaryResponse {
	s.Headers = v
	return s
}

func (s *RunHotTopicSummaryResponse) SetStatusCode(v int32) *RunHotTopicSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *RunHotTopicSummaryResponse) SetBody(v *RunHotTopicSummaryResponseBody) *RunHotTopicSummaryResponse {
	s.Body = v
	return s
}

func (s *RunHotTopicSummaryResponse) Validate() error {
	return dara.Validate(s)
}
