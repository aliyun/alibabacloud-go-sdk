// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSAllEventListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDDoSAllEventListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDDoSAllEventListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDDoSAllEventListResponseBody) *DescribeDDoSAllEventListResponse
	GetBody() *DescribeDDoSAllEventListResponseBody
}

type DescribeDDoSAllEventListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDDoSAllEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDDoSAllEventListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSAllEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDoSAllEventListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDDoSAllEventListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDDoSAllEventListResponse) GetBody() *DescribeDDoSAllEventListResponseBody {
	return s.Body
}

func (s *DescribeDDoSAllEventListResponse) SetHeaders(v map[string]*string) *DescribeDDoSAllEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDoSAllEventListResponse) SetStatusCode(v int32) *DescribeDDoSAllEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDDoSAllEventListResponse) SetBody(v *DescribeDDoSAllEventListResponseBody) *DescribeDDoSAllEventListResponse {
	s.Body = v
	return s
}

func (s *DescribeDDoSAllEventListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
