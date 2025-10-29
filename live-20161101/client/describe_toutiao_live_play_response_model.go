// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeToutiaoLivePlayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeToutiaoLivePlayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeToutiaoLivePlayResponse
	GetStatusCode() *int32
	SetBody(v *DescribeToutiaoLivePlayResponseBody) *DescribeToutiaoLivePlayResponse
	GetBody() *DescribeToutiaoLivePlayResponseBody
}

type DescribeToutiaoLivePlayResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeToutiaoLivePlayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeToutiaoLivePlayResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeToutiaoLivePlayResponse) GoString() string {
	return s.String()
}

func (s *DescribeToutiaoLivePlayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeToutiaoLivePlayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeToutiaoLivePlayResponse) GetBody() *DescribeToutiaoLivePlayResponseBody {
	return s.Body
}

func (s *DescribeToutiaoLivePlayResponse) SetHeaders(v map[string]*string) *DescribeToutiaoLivePlayResponse {
	s.Headers = v
	return s
}

func (s *DescribeToutiaoLivePlayResponse) SetStatusCode(v int32) *DescribeToutiaoLivePlayResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeToutiaoLivePlayResponse) SetBody(v *DescribeToutiaoLivePlayResponseBody) *DescribeToutiaoLivePlayResponse {
	s.Body = v
	return s
}

func (s *DescribeToutiaoLivePlayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
