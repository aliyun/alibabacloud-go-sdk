// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlaEventListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSlaEventListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSlaEventListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSlaEventListResponseBody) *DescribeSlaEventListResponse
	GetBody() *DescribeSlaEventListResponseBody
}

type DescribeSlaEventListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlaEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlaEventListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlaEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlaEventListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSlaEventListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSlaEventListResponse) GetBody() *DescribeSlaEventListResponseBody {
	return s.Body
}

func (s *DescribeSlaEventListResponse) SetHeaders(v map[string]*string) *DescribeSlaEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlaEventListResponse) SetStatusCode(v int32) *DescribeSlaEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlaEventListResponse) SetBody(v *DescribeSlaEventListResponseBody) *DescribeSlaEventListResponse {
	s.Body = v
	return s
}

func (s *DescribeSlaEventListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
