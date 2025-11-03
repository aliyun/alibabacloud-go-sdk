// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRepoTriggerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRepoTriggerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRepoTriggerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRepoTriggerResponseBody) *UpdateRepoTriggerResponse
	GetBody() *UpdateRepoTriggerResponseBody
}

type UpdateRepoTriggerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRepoTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRepoTriggerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepoTriggerResponse) GoString() string {
	return s.String()
}

func (s *UpdateRepoTriggerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRepoTriggerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRepoTriggerResponse) GetBody() *UpdateRepoTriggerResponseBody {
	return s.Body
}

func (s *UpdateRepoTriggerResponse) SetHeaders(v map[string]*string) *UpdateRepoTriggerResponse {
	s.Headers = v
	return s
}

func (s *UpdateRepoTriggerResponse) SetStatusCode(v int32) *UpdateRepoTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRepoTriggerResponse) SetBody(v *UpdateRepoTriggerResponseBody) *UpdateRepoTriggerResponse {
	s.Body = v
	return s
}

func (s *UpdateRepoTriggerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
