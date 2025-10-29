// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveDelayConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveDelayConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveDelayConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveDelayConfigResponseBody) *UpdateLiveDelayConfigResponse
	GetBody() *UpdateLiveDelayConfigResponseBody
}

type UpdateLiveDelayConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveDelayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveDelayConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveDelayConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveDelayConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveDelayConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveDelayConfigResponse) GetBody() *UpdateLiveDelayConfigResponseBody {
	return s.Body
}

func (s *UpdateLiveDelayConfigResponse) SetHeaders(v map[string]*string) *UpdateLiveDelayConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveDelayConfigResponse) SetStatusCode(v int32) *UpdateLiveDelayConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveDelayConfigResponse) SetBody(v *UpdateLiveDelayConfigResponseBody) *UpdateLiveDelayConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveDelayConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
