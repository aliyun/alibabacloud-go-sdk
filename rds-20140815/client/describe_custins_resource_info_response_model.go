// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustinsResourceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustinsResourceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustinsResourceInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustinsResourceInfoResponseBody) *DescribeCustinsResourceInfoResponse
	GetBody() *DescribeCustinsResourceInfoResponseBody
}

type DescribeCustinsResourceInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustinsResourceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustinsResourceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustinsResourceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustinsResourceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustinsResourceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustinsResourceInfoResponse) GetBody() *DescribeCustinsResourceInfoResponseBody {
	return s.Body
}

func (s *DescribeCustinsResourceInfoResponse) SetHeaders(v map[string]*string) *DescribeCustinsResourceInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustinsResourceInfoResponse) SetStatusCode(v int32) *DescribeCustinsResourceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustinsResourceInfoResponse) SetBody(v *DescribeCustinsResourceInfoResponseBody) *DescribeCustinsResourceInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeCustinsResourceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
