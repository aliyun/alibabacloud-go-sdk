// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveDetectNotifyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveDetectNotifyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveDetectNotifyConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveDetectNotifyConfigResponseBody) *UpdateLiveDetectNotifyConfigResponse
	GetBody() *UpdateLiveDetectNotifyConfigResponseBody
}

type UpdateLiveDetectNotifyConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveDetectNotifyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveDetectNotifyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveDetectNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveDetectNotifyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveDetectNotifyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveDetectNotifyConfigResponse) GetBody() *UpdateLiveDetectNotifyConfigResponseBody {
	return s.Body
}

func (s *UpdateLiveDetectNotifyConfigResponse) SetHeaders(v map[string]*string) *UpdateLiveDetectNotifyConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveDetectNotifyConfigResponse) SetStatusCode(v int32) *UpdateLiveDetectNotifyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveDetectNotifyConfigResponse) SetBody(v *UpdateLiveDetectNotifyConfigResponseBody) *UpdateLiveDetectNotifyConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveDetectNotifyConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
