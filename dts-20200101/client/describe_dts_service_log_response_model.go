// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDtsServiceLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDtsServiceLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDtsServiceLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDtsServiceLogResponseBody) *DescribeDtsServiceLogResponse
	GetBody() *DescribeDtsServiceLogResponseBody
}

type DescribeDtsServiceLogResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDtsServiceLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDtsServiceLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsServiceLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeDtsServiceLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDtsServiceLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDtsServiceLogResponse) GetBody() *DescribeDtsServiceLogResponseBody {
	return s.Body
}

func (s *DescribeDtsServiceLogResponse) SetHeaders(v map[string]*string) *DescribeDtsServiceLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeDtsServiceLogResponse) SetStatusCode(v int32) *DescribeDtsServiceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDtsServiceLogResponse) SetBody(v *DescribeDtsServiceLogResponseBody) *DescribeDtsServiceLogResponse {
	s.Body = v
	return s
}

func (s *DescribeDtsServiceLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
