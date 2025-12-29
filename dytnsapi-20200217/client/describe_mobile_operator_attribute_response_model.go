// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMobileOperatorAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMobileOperatorAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMobileOperatorAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMobileOperatorAttributeResponseBody) *DescribeMobileOperatorAttributeResponse
	GetBody() *DescribeMobileOperatorAttributeResponseBody
}

type DescribeMobileOperatorAttributeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMobileOperatorAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMobileOperatorAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileOperatorAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeMobileOperatorAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMobileOperatorAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMobileOperatorAttributeResponse) GetBody() *DescribeMobileOperatorAttributeResponseBody {
	return s.Body
}

func (s *DescribeMobileOperatorAttributeResponse) SetHeaders(v map[string]*string) *DescribeMobileOperatorAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeMobileOperatorAttributeResponse) SetStatusCode(v int32) *DescribeMobileOperatorAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMobileOperatorAttributeResponse) SetBody(v *DescribeMobileOperatorAttributeResponseBody) *DescribeMobileOperatorAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeMobileOperatorAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
