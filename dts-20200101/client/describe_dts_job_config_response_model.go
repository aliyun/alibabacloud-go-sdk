// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDtsJobConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDtsJobConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDtsJobConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDtsJobConfigResponseBody) *DescribeDtsJobConfigResponse
	GetBody() *DescribeDtsJobConfigResponseBody
}

type DescribeDtsJobConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDtsJobConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDtsJobConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDtsJobConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDtsJobConfigResponse) GetBody() *DescribeDtsJobConfigResponseBody {
	return s.Body
}

func (s *DescribeDtsJobConfigResponse) SetHeaders(v map[string]*string) *DescribeDtsJobConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDtsJobConfigResponse) SetStatusCode(v int32) *DescribeDtsJobConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDtsJobConfigResponse) SetBody(v *DescribeDtsJobConfigResponseBody) *DescribeDtsJobConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeDtsJobConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
