// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDosEventAttackTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDDosEventAttackTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDDosEventAttackTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDDosEventAttackTypeResponseBody) *DescribeDDosEventAttackTypeResponse
	GetBody() *DescribeDDosEventAttackTypeResponseBody
}

type DescribeDDosEventAttackTypeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDDosEventAttackTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDDosEventAttackTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosEventAttackTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventAttackTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDDosEventAttackTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDDosEventAttackTypeResponse) GetBody() *DescribeDDosEventAttackTypeResponseBody {
	return s.Body
}

func (s *DescribeDDosEventAttackTypeResponse) SetHeaders(v map[string]*string) *DescribeDDosEventAttackTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDosEventAttackTypeResponse) SetStatusCode(v int32) *DescribeDDosEventAttackTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDDosEventAttackTypeResponse) SetBody(v *DescribeDDosEventAttackTypeResponseBody) *DescribeDDosEventAttackTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeDDosEventAttackTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
