// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosMaxBurstGbpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDdosMaxBurstGbpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDdosMaxBurstGbpsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDdosMaxBurstGbpsResponseBody) *DescribeDdosMaxBurstGbpsResponse
	GetBody() *DescribeDdosMaxBurstGbpsResponseBody
}

type DescribeDdosMaxBurstGbpsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDdosMaxBurstGbpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDdosMaxBurstGbpsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosMaxBurstGbpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDdosMaxBurstGbpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDdosMaxBurstGbpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDdosMaxBurstGbpsResponse) GetBody() *DescribeDdosMaxBurstGbpsResponseBody {
	return s.Body
}

func (s *DescribeDdosMaxBurstGbpsResponse) SetHeaders(v map[string]*string) *DescribeDdosMaxBurstGbpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDdosMaxBurstGbpsResponse) SetStatusCode(v int32) *DescribeDdosMaxBurstGbpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDdosMaxBurstGbpsResponse) SetBody(v *DescribeDdosMaxBurstGbpsResponseBody) *DescribeDdosMaxBurstGbpsResponse {
	s.Body = v
	return s
}

func (s *DescribeDdosMaxBurstGbpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
