// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamsNotifyUrlConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamsNotifyUrlConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamsNotifyUrlConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamsNotifyUrlConfigResponseBody) *DescribeLiveStreamsNotifyUrlConfigResponse
	GetBody() *DescribeLiveStreamsNotifyUrlConfigResponseBody
}

type DescribeLiveStreamsNotifyUrlConfigResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamsNotifyUrlConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamsNotifyUrlConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsNotifyUrlConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponse) GetBody() *DescribeLiveStreamsNotifyUrlConfigResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamsNotifyUrlConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponse) SetStatusCode(v int32) *DescribeLiveStreamsNotifyUrlConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponse) SetBody(v *DescribeLiveStreamsNotifyUrlConfigResponseBody) *DescribeLiveStreamsNotifyUrlConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamsNotifyUrlConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
