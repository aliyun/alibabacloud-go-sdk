// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDosAllEventListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDDosAllEventListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDDosAllEventListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDDosAllEventListResponseBody) *DescribeDDosAllEventListResponse
	GetBody() *DescribeDDosAllEventListResponseBody
}

type DescribeDDosAllEventListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDDosAllEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDDosAllEventListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosAllEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDDosAllEventListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDDosAllEventListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDDosAllEventListResponse) GetBody() *DescribeDDosAllEventListResponseBody {
	return s.Body
}

func (s *DescribeDDosAllEventListResponse) SetHeaders(v map[string]*string) *DescribeDDosAllEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDDosAllEventListResponse) SetStatusCode(v int32) *DescribeDDosAllEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDDosAllEventListResponse) SetBody(v *DescribeDDosAllEventListResponseBody) *DescribeDDosAllEventListResponse {
	s.Body = v
	return s
}

func (s *DescribeDDosAllEventListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
