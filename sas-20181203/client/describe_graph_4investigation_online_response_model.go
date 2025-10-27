// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGraph4InvestigationOnlineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGraph4InvestigationOnlineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGraph4InvestigationOnlineResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGraph4InvestigationOnlineResponseBody) *DescribeGraph4InvestigationOnlineResponse
	GetBody() *DescribeGraph4InvestigationOnlineResponseBody
}

type DescribeGraph4InvestigationOnlineResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGraph4InvestigationOnlineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGraph4InvestigationOnlineResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGraph4InvestigationOnlineResponse) GoString() string {
	return s.String()
}

func (s *DescribeGraph4InvestigationOnlineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGraph4InvestigationOnlineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGraph4InvestigationOnlineResponse) GetBody() *DescribeGraph4InvestigationOnlineResponseBody {
	return s.Body
}

func (s *DescribeGraph4InvestigationOnlineResponse) SetHeaders(v map[string]*string) *DescribeGraph4InvestigationOnlineResponse {
	s.Headers = v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponse) SetStatusCode(v int32) *DescribeGraph4InvestigationOnlineResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponse) SetBody(v *DescribeGraph4InvestigationOnlineResponseBody) *DescribeGraph4InvestigationOnlineResponse {
	s.Body = v
	return s
}

func (s *DescribeGraph4InvestigationOnlineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
