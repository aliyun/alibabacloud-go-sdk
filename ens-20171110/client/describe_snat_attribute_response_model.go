// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnatAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSnatAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSnatAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSnatAttributeResponseBody) *DescribeSnatAttributeResponse
	GetBody() *DescribeSnatAttributeResponseBody
}

type DescribeSnatAttributeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSnatAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSnatAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnatAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeSnatAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSnatAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSnatAttributeResponse) GetBody() *DescribeSnatAttributeResponseBody {
	return s.Body
}

func (s *DescribeSnatAttributeResponse) SetHeaders(v map[string]*string) *DescribeSnatAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeSnatAttributeResponse) SetStatusCode(v int32) *DescribeSnatAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSnatAttributeResponse) SetBody(v *DescribeSnatAttributeResponseBody) *DescribeSnatAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeSnatAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
