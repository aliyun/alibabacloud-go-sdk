// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCnameReusesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCnameReusesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCnameReusesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCnameReusesResponseBody) *DescribeCnameReusesResponse
	GetBody() *DescribeCnameReusesResponseBody
}

type DescribeCnameReusesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCnameReusesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCnameReusesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCnameReusesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCnameReusesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCnameReusesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCnameReusesResponse) GetBody() *DescribeCnameReusesResponseBody {
	return s.Body
}

func (s *DescribeCnameReusesResponse) SetHeaders(v map[string]*string) *DescribeCnameReusesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCnameReusesResponse) SetStatusCode(v int32) *DescribeCnameReusesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCnameReusesResponse) SetBody(v *DescribeCnameReusesResponseBody) *DescribeCnameReusesResponse {
	s.Body = v
	return s
}

func (s *DescribeCnameReusesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
