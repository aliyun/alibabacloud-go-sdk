// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVsStreamsNotifyUrlConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetVsStreamsNotifyUrlConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetVsStreamsNotifyUrlConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetVsStreamsNotifyUrlConfigResponseBody) *SetVsStreamsNotifyUrlConfigResponse
	GetBody() *SetVsStreamsNotifyUrlConfigResponseBody
}

type SetVsStreamsNotifyUrlConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetVsStreamsNotifyUrlConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetVsStreamsNotifyUrlConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetVsStreamsNotifyUrlConfigResponse) GoString() string {
	return s.String()
}

func (s *SetVsStreamsNotifyUrlConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetVsStreamsNotifyUrlConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetVsStreamsNotifyUrlConfigResponse) GetBody() *SetVsStreamsNotifyUrlConfigResponseBody {
	return s.Body
}

func (s *SetVsStreamsNotifyUrlConfigResponse) SetHeaders(v map[string]*string) *SetVsStreamsNotifyUrlConfigResponse {
	s.Headers = v
	return s
}

func (s *SetVsStreamsNotifyUrlConfigResponse) SetStatusCode(v int32) *SetVsStreamsNotifyUrlConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetVsStreamsNotifyUrlConfigResponse) SetBody(v *SetVsStreamsNotifyUrlConfigResponseBody) *SetVsStreamsNotifyUrlConfigResponse {
	s.Body = v
	return s
}

func (s *SetVsStreamsNotifyUrlConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
