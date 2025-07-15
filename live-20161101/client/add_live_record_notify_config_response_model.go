// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveRecordNotifyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLiveRecordNotifyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLiveRecordNotifyConfigResponse
	GetStatusCode() *int32
	SetBody(v *AddLiveRecordNotifyConfigResponseBody) *AddLiveRecordNotifyConfigResponse
	GetBody() *AddLiveRecordNotifyConfigResponseBody
}

type AddLiveRecordNotifyConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveRecordNotifyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveRecordNotifyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLiveRecordNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLiveRecordNotifyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLiveRecordNotifyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLiveRecordNotifyConfigResponse) GetBody() *AddLiveRecordNotifyConfigResponseBody {
	return s.Body
}

func (s *AddLiveRecordNotifyConfigResponse) SetHeaders(v map[string]*string) *AddLiveRecordNotifyConfigResponse {
	s.Headers = v
	return s
}

func (s *AddLiveRecordNotifyConfigResponse) SetStatusCode(v int32) *AddLiveRecordNotifyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveRecordNotifyConfigResponse) SetBody(v *AddLiveRecordNotifyConfigResponseBody) *AddLiveRecordNotifyConfigResponse {
	s.Body = v
	return s
}

func (s *AddLiveRecordNotifyConfigResponse) Validate() error {
	return dara.Validate(s)
}
