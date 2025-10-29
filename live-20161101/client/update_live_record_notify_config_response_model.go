// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveRecordNotifyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveRecordNotifyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveRecordNotifyConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveRecordNotifyConfigResponseBody) *UpdateLiveRecordNotifyConfigResponse
	GetBody() *UpdateLiveRecordNotifyConfigResponseBody
}

type UpdateLiveRecordNotifyConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveRecordNotifyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveRecordNotifyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveRecordNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveRecordNotifyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveRecordNotifyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveRecordNotifyConfigResponse) GetBody() *UpdateLiveRecordNotifyConfigResponseBody {
	return s.Body
}

func (s *UpdateLiveRecordNotifyConfigResponse) SetHeaders(v map[string]*string) *UpdateLiveRecordNotifyConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveRecordNotifyConfigResponse) SetStatusCode(v int32) *UpdateLiveRecordNotifyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveRecordNotifyConfigResponse) SetBody(v *UpdateLiveRecordNotifyConfigResponseBody) *UpdateLiveRecordNotifyConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveRecordNotifyConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
