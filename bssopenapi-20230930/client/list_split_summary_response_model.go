// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSplitSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSplitSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSplitSummaryResponse
	GetStatusCode() *int32
	SetBody(v *ListSplitSummaryResponseBody) *ListSplitSummaryResponse
	GetBody() *ListSplitSummaryResponseBody
}

type ListSplitSummaryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSplitSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSplitSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSplitSummaryResponse) GoString() string {
	return s.String()
}

func (s *ListSplitSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSplitSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSplitSummaryResponse) GetBody() *ListSplitSummaryResponseBody {
	return s.Body
}

func (s *ListSplitSummaryResponse) SetHeaders(v map[string]*string) *ListSplitSummaryResponse {
	s.Headers = v
	return s
}

func (s *ListSplitSummaryResponse) SetStatusCode(v int32) *ListSplitSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSplitSummaryResponse) SetBody(v *ListSplitSummaryResponseBody) *ListSplitSummaryResponse {
	s.Body = v
	return s
}

func (s *ListSplitSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
