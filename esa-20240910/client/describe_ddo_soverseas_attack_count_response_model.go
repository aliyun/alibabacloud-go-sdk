// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSOverseasAttackCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDDoSOverseasAttackCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDDoSOverseasAttackCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDDoSOverseasAttackCountResponseBody) *DescribeDDoSOverseasAttackCountResponse
	GetBody() *DescribeDDoSOverseasAttackCountResponseBody
}

type DescribeDDoSOverseasAttackCountResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDDoSOverseasAttackCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDDoSOverseasAttackCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSOverseasAttackCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDoSOverseasAttackCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDDoSOverseasAttackCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDDoSOverseasAttackCountResponse) GetBody() *DescribeDDoSOverseasAttackCountResponseBody {
	return s.Body
}

func (s *DescribeDDoSOverseasAttackCountResponse) SetHeaders(v map[string]*string) *DescribeDDoSOverseasAttackCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDoSOverseasAttackCountResponse) SetStatusCode(v int32) *DescribeDDoSOverseasAttackCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDDoSOverseasAttackCountResponse) SetBody(v *DescribeDDoSOverseasAttackCountResponseBody) *DescribeDDoSOverseasAttackCountResponse {
	s.Body = v
	return s
}

func (s *DescribeDDoSOverseasAttackCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
