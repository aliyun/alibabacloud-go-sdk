// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafBotAppKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnWafBotAppKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnWafBotAppKeyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnWafBotAppKeyResponseBody) *DescribeDcdnWafBotAppKeyResponse
	GetBody() *DescribeDcdnWafBotAppKeyResponseBody
}

type DescribeDcdnWafBotAppKeyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnWafBotAppKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnWafBotAppKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafBotAppKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafBotAppKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnWafBotAppKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnWafBotAppKeyResponse) GetBody() *DescribeDcdnWafBotAppKeyResponseBody {
	return s.Body
}

func (s *DescribeDcdnWafBotAppKeyResponse) SetHeaders(v map[string]*string) *DescribeDcdnWafBotAppKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnWafBotAppKeyResponse) SetStatusCode(v int32) *DescribeDcdnWafBotAppKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnWafBotAppKeyResponse) SetBody(v *DescribeDcdnWafBotAppKeyResponseBody) *DescribeDcdnWafBotAppKeyResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnWafBotAppKeyResponse) Validate() error {
	return dara.Validate(s)
}
