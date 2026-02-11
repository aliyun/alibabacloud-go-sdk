// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertContactGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAlertContactGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAlertContactGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAlertContactGroupResponseBody) *UpdateAlertContactGroupResponse
	GetBody() *UpdateAlertContactGroupResponseBody
}

type UpdateAlertContactGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAlertContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAlertContactGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertContactGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateAlertContactGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAlertContactGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAlertContactGroupResponse) GetBody() *UpdateAlertContactGroupResponseBody {
	return s.Body
}

func (s *UpdateAlertContactGroupResponse) SetHeaders(v map[string]*string) *UpdateAlertContactGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateAlertContactGroupResponse) SetStatusCode(v int32) *UpdateAlertContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAlertContactGroupResponse) SetBody(v *UpdateAlertContactGroupResponseBody) *UpdateAlertContactGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateAlertContactGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
