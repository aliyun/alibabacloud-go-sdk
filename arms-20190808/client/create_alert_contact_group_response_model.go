// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertContactGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAlertContactGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAlertContactGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateAlertContactGroupResponseBody) *CreateAlertContactGroupResponse
	GetBody() *CreateAlertContactGroupResponseBody
}

type CreateAlertContactGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAlertContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAlertContactGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertContactGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAlertContactGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAlertContactGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAlertContactGroupResponse) GetBody() *CreateAlertContactGroupResponseBody {
	return s.Body
}

func (s *CreateAlertContactGroupResponse) SetHeaders(v map[string]*string) *CreateAlertContactGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAlertContactGroupResponse) SetStatusCode(v int32) *CreateAlertContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAlertContactGroupResponse) SetBody(v *CreateAlertContactGroupResponseBody) *CreateAlertContactGroupResponse {
	s.Body = v
	return s
}

func (s *CreateAlertContactGroupResponse) Validate() error {
	return dara.Validate(s)
}
