// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMajorProtectionBlackIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMajorProtectionBlackIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMajorProtectionBlackIpsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMajorProtectionBlackIpsResponseBody) *DescribeMajorProtectionBlackIpsResponse
	GetBody() *DescribeMajorProtectionBlackIpsResponseBody
}

type DescribeMajorProtectionBlackIpsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMajorProtectionBlackIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMajorProtectionBlackIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMajorProtectionBlackIpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMajorProtectionBlackIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMajorProtectionBlackIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMajorProtectionBlackIpsResponse) GetBody() *DescribeMajorProtectionBlackIpsResponseBody {
	return s.Body
}

func (s *DescribeMajorProtectionBlackIpsResponse) SetHeaders(v map[string]*string) *DescribeMajorProtectionBlackIpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponse) SetStatusCode(v int32) *DescribeMajorProtectionBlackIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponse) SetBody(v *DescribeMajorProtectionBlackIpsResponseBody) *DescribeMajorProtectionBlackIpsResponse {
	s.Body = v
	return s
}

func (s *DescribeMajorProtectionBlackIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
