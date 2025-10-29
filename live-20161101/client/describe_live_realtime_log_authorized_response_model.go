// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveRealtimeLogAuthorizedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveRealtimeLogAuthorizedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveRealtimeLogAuthorizedResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveRealtimeLogAuthorizedResponseBody) *DescribeLiveRealtimeLogAuthorizedResponse
	GetBody() *DescribeLiveRealtimeLogAuthorizedResponseBody
}

type DescribeLiveRealtimeLogAuthorizedResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveRealtimeLogAuthorizedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveRealtimeLogAuthorizedResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRealtimeLogAuthorizedResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveRealtimeLogAuthorizedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveRealtimeLogAuthorizedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveRealtimeLogAuthorizedResponse) GetBody() *DescribeLiveRealtimeLogAuthorizedResponseBody {
	return s.Body
}

func (s *DescribeLiveRealtimeLogAuthorizedResponse) SetHeaders(v map[string]*string) *DescribeLiveRealtimeLogAuthorizedResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveRealtimeLogAuthorizedResponse) SetStatusCode(v int32) *DescribeLiveRealtimeLogAuthorizedResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveRealtimeLogAuthorizedResponse) SetBody(v *DescribeLiveRealtimeLogAuthorizedResponseBody) *DescribeLiveRealtimeLogAuthorizedResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveRealtimeLogAuthorizedResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
