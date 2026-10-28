// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateReadonlyModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateReadonlyModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateReadonlyModeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateReadonlyModeResponseBody) *UpdateReadonlyModeResponse
	GetBody() *UpdateReadonlyModeResponseBody
}

type UpdateReadonlyModeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateReadonlyModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateReadonlyModeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateReadonlyModeResponse) GoString() string {
	return s.String()
}

func (s *UpdateReadonlyModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateReadonlyModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateReadonlyModeResponse) GetBody() *UpdateReadonlyModeResponseBody {
	return s.Body
}

func (s *UpdateReadonlyModeResponse) SetHeaders(v map[string]*string) *UpdateReadonlyModeResponse {
	s.Headers = v
	return s
}

func (s *UpdateReadonlyModeResponse) SetStatusCode(v int32) *UpdateReadonlyModeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateReadonlyModeResponse) SetBody(v *UpdateReadonlyModeResponseBody) *UpdateReadonlyModeResponse {
	s.Body = v
	return s
}

func (s *UpdateReadonlyModeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
