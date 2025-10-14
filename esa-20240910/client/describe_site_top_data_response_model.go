// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteTopDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSiteTopDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSiteTopDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSiteTopDataResponseBody) *DescribeSiteTopDataResponse
	GetBody() *DescribeSiteTopDataResponseBody
}

type DescribeSiteTopDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSiteTopDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSiteTopDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteTopDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeSiteTopDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSiteTopDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSiteTopDataResponse) GetBody() *DescribeSiteTopDataResponseBody {
	return s.Body
}

func (s *DescribeSiteTopDataResponse) SetHeaders(v map[string]*string) *DescribeSiteTopDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeSiteTopDataResponse) SetStatusCode(v int32) *DescribeSiteTopDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSiteTopDataResponse) SetBody(v *DescribeSiteTopDataResponseBody) *DescribeSiteTopDataResponse {
	s.Body = v
	return s
}

func (s *DescribeSiteTopDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
