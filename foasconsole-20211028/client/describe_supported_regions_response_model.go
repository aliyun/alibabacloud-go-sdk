// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupportedRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSupportedRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSupportedRegionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSupportedRegionsResponseBody) *DescribeSupportedRegionsResponse
	GetBody() *DescribeSupportedRegionsResponseBody
}

type DescribeSupportedRegionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSupportedRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSupportedRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportedRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSupportedRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSupportedRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSupportedRegionsResponse) GetBody() *DescribeSupportedRegionsResponseBody {
	return s.Body
}

func (s *DescribeSupportedRegionsResponse) SetHeaders(v map[string]*string) *DescribeSupportedRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSupportedRegionsResponse) SetStatusCode(v int32) *DescribeSupportedRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSupportedRegionsResponse) SetBody(v *DescribeSupportedRegionsResponseBody) *DescribeSupportedRegionsResponse {
	s.Body = v
	return s
}

func (s *DescribeSupportedRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
