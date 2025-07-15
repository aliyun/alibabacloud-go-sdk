// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveStreamDelayConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLiveStreamDelayConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLiveStreamDelayConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetLiveStreamDelayConfigResponseBody) *SetLiveStreamDelayConfigResponse
	GetBody() *SetLiveStreamDelayConfigResponseBody
}

type SetLiveStreamDelayConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLiveStreamDelayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLiveStreamDelayConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLiveStreamDelayConfigResponse) GoString() string {
	return s.String()
}

func (s *SetLiveStreamDelayConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLiveStreamDelayConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLiveStreamDelayConfigResponse) GetBody() *SetLiveStreamDelayConfigResponseBody {
	return s.Body
}

func (s *SetLiveStreamDelayConfigResponse) SetHeaders(v map[string]*string) *SetLiveStreamDelayConfigResponse {
	s.Headers = v
	return s
}

func (s *SetLiveStreamDelayConfigResponse) SetStatusCode(v int32) *SetLiveStreamDelayConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLiveStreamDelayConfigResponse) SetBody(v *SetLiveStreamDelayConfigResponseBody) *SetLiveStreamDelayConfigResponse {
	s.Body = v
	return s
}

func (s *SetLiveStreamDelayConfigResponse) Validate() error {
	return dara.Validate(s)
}
