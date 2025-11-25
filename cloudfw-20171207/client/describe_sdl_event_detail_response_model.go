// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSdlEventDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSdlEventDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSdlEventDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSdlEventDetailResponseBody) *DescribeSdlEventDetailResponse
	GetBody() *DescribeSdlEventDetailResponseBody
}

type DescribeSdlEventDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSdlEventDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSdlEventDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlEventDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeSdlEventDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSdlEventDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSdlEventDetailResponse) GetBody() *DescribeSdlEventDetailResponseBody {
	return s.Body
}

func (s *DescribeSdlEventDetailResponse) SetHeaders(v map[string]*string) *DescribeSdlEventDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeSdlEventDetailResponse) SetStatusCode(v int32) *DescribeSdlEventDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSdlEventDetailResponse) SetBody(v *DescribeSdlEventDetailResponseBody) *DescribeSdlEventDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeSdlEventDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
