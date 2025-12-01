// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDLAServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDLAServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDLAServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDLAServiceResponseBody) *DescribeDLAServiceResponse
	GetBody() *DescribeDLAServiceResponseBody
}

type DescribeDLAServiceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDLAServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDLAServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDLAServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDLAServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDLAServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDLAServiceResponse) GetBody() *DescribeDLAServiceResponseBody {
	return s.Body
}

func (s *DescribeDLAServiceResponse) SetHeaders(v map[string]*string) *DescribeDLAServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDLAServiceResponse) SetStatusCode(v int32) *DescribeDLAServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDLAServiceResponse) SetBody(v *DescribeDLAServiceResponseBody) *DescribeDLAServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeDLAServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
