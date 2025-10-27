// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSasPmAgentListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSasPmAgentListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSasPmAgentListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSasPmAgentListResponseBody) *DescribeSasPmAgentListResponse
	GetBody() *DescribeSasPmAgentListResponseBody
}

type DescribeSasPmAgentListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSasPmAgentListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSasPmAgentListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSasPmAgentListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSasPmAgentListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSasPmAgentListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSasPmAgentListResponse) GetBody() *DescribeSasPmAgentListResponseBody {
	return s.Body
}

func (s *DescribeSasPmAgentListResponse) SetHeaders(v map[string]*string) *DescribeSasPmAgentListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSasPmAgentListResponse) SetStatusCode(v int32) *DescribeSasPmAgentListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSasPmAgentListResponse) SetBody(v *DescribeSasPmAgentListResponseBody) *DescribeSasPmAgentListResponse {
	s.Body = v
	return s
}

func (s *DescribeSasPmAgentListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
