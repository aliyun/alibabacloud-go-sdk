// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveDetectNotifyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLiveDetectNotifyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLiveDetectNotifyConfigResponse
	GetStatusCode() *int32
	SetBody(v *AddLiveDetectNotifyConfigResponseBody) *AddLiveDetectNotifyConfigResponse
	GetBody() *AddLiveDetectNotifyConfigResponseBody
}

type AddLiveDetectNotifyConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveDetectNotifyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveDetectNotifyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLiveDetectNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLiveDetectNotifyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLiveDetectNotifyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLiveDetectNotifyConfigResponse) GetBody() *AddLiveDetectNotifyConfigResponseBody {
	return s.Body
}

func (s *AddLiveDetectNotifyConfigResponse) SetHeaders(v map[string]*string) *AddLiveDetectNotifyConfigResponse {
	s.Headers = v
	return s
}

func (s *AddLiveDetectNotifyConfigResponse) SetStatusCode(v int32) *AddLiveDetectNotifyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveDetectNotifyConfigResponse) SetBody(v *AddLiveDetectNotifyConfigResponseBody) *AddLiveDetectNotifyConfigResponse {
	s.Body = v
	return s
}

func (s *AddLiveDetectNotifyConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
