// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppLiveStreamStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAppLiveStreamStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAppLiveStreamStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAppLiveStreamStatusResponseBody) *ModifyAppLiveStreamStatusResponse
	GetBody() *ModifyAppLiveStreamStatusResponseBody
}

type ModifyAppLiveStreamStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppLiveStreamStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppLiveStreamStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppLiveStreamStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppLiveStreamStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAppLiveStreamStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAppLiveStreamStatusResponse) GetBody() *ModifyAppLiveStreamStatusResponseBody {
	return s.Body
}

func (s *ModifyAppLiveStreamStatusResponse) SetHeaders(v map[string]*string) *ModifyAppLiveStreamStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppLiveStreamStatusResponse) SetStatusCode(v int32) *ModifyAppLiveStreamStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppLiveStreamStatusResponse) SetBody(v *ModifyAppLiveStreamStatusResponseBody) *ModifyAppLiveStreamStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyAppLiveStreamStatusResponse) Validate() error {
	return dara.Validate(s)
}
