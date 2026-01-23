// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQualityWatchSwitchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateQualityWatchSwitchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateQualityWatchSwitchResponse
	GetStatusCode() *int32
	SetBody(v *UpdateQualityWatchSwitchResponseBody) *UpdateQualityWatchSwitchResponse
	GetBody() *UpdateQualityWatchSwitchResponseBody
}

type UpdateQualityWatchSwitchResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateQualityWatchSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQualityWatchSwitchResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityWatchSwitchResponse) GoString() string {
	return s.String()
}

func (s *UpdateQualityWatchSwitchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateQualityWatchSwitchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateQualityWatchSwitchResponse) GetBody() *UpdateQualityWatchSwitchResponseBody {
	return s.Body
}

func (s *UpdateQualityWatchSwitchResponse) SetHeaders(v map[string]*string) *UpdateQualityWatchSwitchResponse {
	s.Headers = v
	return s
}

func (s *UpdateQualityWatchSwitchResponse) SetStatusCode(v int32) *UpdateQualityWatchSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQualityWatchSwitchResponse) SetBody(v *UpdateQualityWatchSwitchResponseBody) *UpdateQualityWatchSwitchResponse {
	s.Body = v
	return s
}

func (s *UpdateQualityWatchSwitchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
