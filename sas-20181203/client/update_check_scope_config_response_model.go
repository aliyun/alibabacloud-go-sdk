// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCheckScopeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCheckScopeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCheckScopeConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCheckScopeConfigResponseBody) *UpdateCheckScopeConfigResponse
	GetBody() *UpdateCheckScopeConfigResponseBody
}

type UpdateCheckScopeConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCheckScopeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCheckScopeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCheckScopeConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateCheckScopeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCheckScopeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCheckScopeConfigResponse) GetBody() *UpdateCheckScopeConfigResponseBody {
	return s.Body
}

func (s *UpdateCheckScopeConfigResponse) SetHeaders(v map[string]*string) *UpdateCheckScopeConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateCheckScopeConfigResponse) SetStatusCode(v int32) *UpdateCheckScopeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCheckScopeConfigResponse) SetBody(v *UpdateCheckScopeConfigResponseBody) *UpdateCheckScopeConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateCheckScopeConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
