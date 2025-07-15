// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeToutiaoLivePublishResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeToutiaoLivePublishResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeToutiaoLivePublishResponse
	GetStatusCode() *int32
	SetBody(v *DescribeToutiaoLivePublishResponseBody) *DescribeToutiaoLivePublishResponse
	GetBody() *DescribeToutiaoLivePublishResponseBody
}

type DescribeToutiaoLivePublishResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeToutiaoLivePublishResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeToutiaoLivePublishResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeToutiaoLivePublishResponse) GoString() string {
	return s.String()
}

func (s *DescribeToutiaoLivePublishResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeToutiaoLivePublishResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeToutiaoLivePublishResponse) GetBody() *DescribeToutiaoLivePublishResponseBody {
	return s.Body
}

func (s *DescribeToutiaoLivePublishResponse) SetHeaders(v map[string]*string) *DescribeToutiaoLivePublishResponse {
	s.Headers = v
	return s
}

func (s *DescribeToutiaoLivePublishResponse) SetStatusCode(v int32) *DescribeToutiaoLivePublishResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeToutiaoLivePublishResponse) SetBody(v *DescribeToutiaoLivePublishResponseBody) *DescribeToutiaoLivePublishResponse {
	s.Body = v
	return s
}

func (s *DescribeToutiaoLivePublishResponse) Validate() error {
	return dara.Validate(s)
}
