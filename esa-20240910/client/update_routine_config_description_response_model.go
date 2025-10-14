// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRoutineConfigDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRoutineConfigDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRoutineConfigDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRoutineConfigDescriptionResponseBody) *UpdateRoutineConfigDescriptionResponse
	GetBody() *UpdateRoutineConfigDescriptionResponseBody
}

type UpdateRoutineConfigDescriptionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRoutineConfigDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRoutineConfigDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRoutineConfigDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateRoutineConfigDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRoutineConfigDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRoutineConfigDescriptionResponse) GetBody() *UpdateRoutineConfigDescriptionResponseBody {
	return s.Body
}

func (s *UpdateRoutineConfigDescriptionResponse) SetHeaders(v map[string]*string) *UpdateRoutineConfigDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateRoutineConfigDescriptionResponse) SetStatusCode(v int32) *UpdateRoutineConfigDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRoutineConfigDescriptionResponse) SetBody(v *UpdateRoutineConfigDescriptionResponseBody) *UpdateRoutineConfigDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateRoutineConfigDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
