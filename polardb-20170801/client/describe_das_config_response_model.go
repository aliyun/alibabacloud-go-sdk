// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDasConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDasConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDasConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDasConfigResponseBody) *DescribeDasConfigResponse
	GetBody() *DescribeDasConfigResponseBody
}

type DescribeDasConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDasConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDasConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDasConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDasConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDasConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDasConfigResponse) GetBody() *DescribeDasConfigResponseBody {
	return s.Body
}

func (s *DescribeDasConfigResponse) SetHeaders(v map[string]*string) *DescribeDasConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDasConfigResponse) SetStatusCode(v int32) *DescribeDasConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDasConfigResponse) SetBody(v *DescribeDasConfigResponseBody) *DescribeDasConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeDasConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
