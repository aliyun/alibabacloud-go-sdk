// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutomateResponseConfigStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAutomateResponseConfigStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAutomateResponseConfigStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAutomateResponseConfigStatusResponseBody) *UpdateAutomateResponseConfigStatusResponse
	GetBody() *UpdateAutomateResponseConfigStatusResponseBody
}

type UpdateAutomateResponseConfigStatusResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAutomateResponseConfigStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAutomateResponseConfigStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutomateResponseConfigStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateAutomateResponseConfigStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAutomateResponseConfigStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAutomateResponseConfigStatusResponse) GetBody() *UpdateAutomateResponseConfigStatusResponseBody {
	return s.Body
}

func (s *UpdateAutomateResponseConfigStatusResponse) SetHeaders(v map[string]*string) *UpdateAutomateResponseConfigStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateAutomateResponseConfigStatusResponse) SetStatusCode(v int32) *UpdateAutomateResponseConfigStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusResponse) SetBody(v *UpdateAutomateResponseConfigStatusResponseBody) *UpdateAutomateResponseConfigStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateAutomateResponseConfigStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
