// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleTagListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSampleTagListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSampleTagListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSampleTagListResponseBody) *DescribeSampleTagListResponse
	GetBody() *DescribeSampleTagListResponseBody
}

type DescribeSampleTagListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSampleTagListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSampleTagListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleTagListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSampleTagListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSampleTagListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSampleTagListResponse) GetBody() *DescribeSampleTagListResponseBody {
	return s.Body
}

func (s *DescribeSampleTagListResponse) SetHeaders(v map[string]*string) *DescribeSampleTagListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSampleTagListResponse) SetStatusCode(v int32) *DescribeSampleTagListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSampleTagListResponse) SetBody(v *DescribeSampleTagListResponseBody) *DescribeSampleTagListResponse {
	s.Body = v
	return s
}

func (s *DescribeSampleTagListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
