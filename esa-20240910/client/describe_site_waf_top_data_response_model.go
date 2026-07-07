// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteWafTopDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSiteWafTopDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSiteWafTopDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSiteWafTopDataResponseBody) *DescribeSiteWafTopDataResponse
	GetBody() *DescribeSiteWafTopDataResponseBody
}

type DescribeSiteWafTopDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSiteWafTopDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSiteWafTopDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteWafTopDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeSiteWafTopDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSiteWafTopDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSiteWafTopDataResponse) GetBody() *DescribeSiteWafTopDataResponseBody {
	return s.Body
}

func (s *DescribeSiteWafTopDataResponse) SetHeaders(v map[string]*string) *DescribeSiteWafTopDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeSiteWafTopDataResponse) SetStatusCode(v int32) *DescribeSiteWafTopDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSiteWafTopDataResponse) SetBody(v *DescribeSiteWafTopDataResponseBody) *DescribeSiteWafTopDataResponse {
	s.Body = v
	return s
}

func (s *DescribeSiteWafTopDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
