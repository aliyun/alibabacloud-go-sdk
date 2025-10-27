// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCommonSwitchConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCommonSwitchConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCommonSwitchConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCommonSwitchConfigResponseBody) *UpdateCommonSwitchConfigResponse
	GetBody() *UpdateCommonSwitchConfigResponseBody
}

type UpdateCommonSwitchConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCommonSwitchConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCommonSwitchConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCommonSwitchConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateCommonSwitchConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCommonSwitchConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCommonSwitchConfigResponse) GetBody() *UpdateCommonSwitchConfigResponseBody {
	return s.Body
}

func (s *UpdateCommonSwitchConfigResponse) SetHeaders(v map[string]*string) *UpdateCommonSwitchConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateCommonSwitchConfigResponse) SetStatusCode(v int32) *UpdateCommonSwitchConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCommonSwitchConfigResponse) SetBody(v *UpdateCommonSwitchConfigResponseBody) *UpdateCommonSwitchConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateCommonSwitchConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
