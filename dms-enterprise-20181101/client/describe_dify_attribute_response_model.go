// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDifyAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDifyAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDifyAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDifyAttributeResponseBody) *DescribeDifyAttributeResponse
	GetBody() *DescribeDifyAttributeResponseBody
}

type DescribeDifyAttributeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDifyAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDifyAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDifyAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDifyAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDifyAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDifyAttributeResponse) GetBody() *DescribeDifyAttributeResponseBody {
	return s.Body
}

func (s *DescribeDifyAttributeResponse) SetHeaders(v map[string]*string) *DescribeDifyAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDifyAttributeResponse) SetStatusCode(v int32) *DescribeDifyAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDifyAttributeResponse) SetBody(v *DescribeDifyAttributeResponseBody) *DescribeDifyAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeDifyAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
