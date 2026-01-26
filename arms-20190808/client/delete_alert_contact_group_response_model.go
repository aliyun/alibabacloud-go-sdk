// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertContactGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAlertContactGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAlertContactGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAlertContactGroupResponseBody) *DeleteAlertContactGroupResponse
	GetBody() *DeleteAlertContactGroupResponseBody
}

type DeleteAlertContactGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAlertContactGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAlertContactGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertContactGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertContactGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAlertContactGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAlertContactGroupResponse) GetBody() *DeleteAlertContactGroupResponseBody {
	return s.Body
}

func (s *DeleteAlertContactGroupResponse) SetHeaders(v map[string]*string) *DeleteAlertContactGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertContactGroupResponse) SetStatusCode(v int32) *DeleteAlertContactGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlertContactGroupResponse) SetBody(v *DeleteAlertContactGroupResponseBody) *DeleteAlertContactGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteAlertContactGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
