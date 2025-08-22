// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserSecDropByMinuteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnUserSecDropByMinuteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnUserSecDropByMinuteResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnUserSecDropByMinuteResponseBody) *DescribeDcdnUserSecDropByMinuteResponse
	GetBody() *DescribeDcdnUserSecDropByMinuteResponseBody
}

type DescribeDcdnUserSecDropByMinuteResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnUserSecDropByMinuteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnUserSecDropByMinuteResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserSecDropByMinuteResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserSecDropByMinuteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnUserSecDropByMinuteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnUserSecDropByMinuteResponse) GetBody() *DescribeDcdnUserSecDropByMinuteResponseBody {
	return s.Body
}

func (s *DescribeDcdnUserSecDropByMinuteResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserSecDropByMinuteResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponse) SetStatusCode(v int32) *DescribeDcdnUserSecDropByMinuteResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponse) SetBody(v *DescribeDcdnUserSecDropByMinuteResponseBody) *DescribeDcdnUserSecDropByMinuteResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnUserSecDropByMinuteResponse) Validate() error {
	return dara.Validate(s)
}
