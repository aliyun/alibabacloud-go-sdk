// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppTransitionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMmAppTransitionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMmAppTransitionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMmAppTransitionResponseBody) *UpdateMmAppTransitionResponse
	GetBody() *UpdateMmAppTransitionResponseBody
}

type UpdateMmAppTransitionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMmAppTransitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMmAppTransitionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppTransitionResponse) GoString() string {
	return s.String()
}

func (s *UpdateMmAppTransitionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMmAppTransitionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMmAppTransitionResponse) GetBody() *UpdateMmAppTransitionResponseBody {
	return s.Body
}

func (s *UpdateMmAppTransitionResponse) SetHeaders(v map[string]*string) *UpdateMmAppTransitionResponse {
	s.Headers = v
	return s
}

func (s *UpdateMmAppTransitionResponse) SetStatusCode(v int32) *UpdateMmAppTransitionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMmAppTransitionResponse) SetBody(v *UpdateMmAppTransitionResponseBody) *UpdateMmAppTransitionResponse {
	s.Body = v
	return s
}

func (s *UpdateMmAppTransitionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
