// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeterLiveBypassDurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMeterLiveBypassDurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMeterLiveBypassDurationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMeterLiveBypassDurationResponseBody) *DescribeMeterLiveBypassDurationResponse
	GetBody() *DescribeMeterLiveBypassDurationResponseBody
}

type DescribeMeterLiveBypassDurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMeterLiveBypassDurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMeterLiveBypassDurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeterLiveBypassDurationResponse) GoString() string {
	return s.String()
}

func (s *DescribeMeterLiveBypassDurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMeterLiveBypassDurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMeterLiveBypassDurationResponse) GetBody() *DescribeMeterLiveBypassDurationResponseBody {
	return s.Body
}

func (s *DescribeMeterLiveBypassDurationResponse) SetHeaders(v map[string]*string) *DescribeMeterLiveBypassDurationResponse {
	s.Headers = v
	return s
}

func (s *DescribeMeterLiveBypassDurationResponse) SetStatusCode(v int32) *DescribeMeterLiveBypassDurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMeterLiveBypassDurationResponse) SetBody(v *DescribeMeterLiveBypassDurationResponseBody) *DescribeMeterLiveBypassDurationResponse {
	s.Body = v
	return s
}

func (s *DescribeMeterLiveBypassDurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
