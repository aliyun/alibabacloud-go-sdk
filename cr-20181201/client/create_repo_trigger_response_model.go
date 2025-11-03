// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRepoTriggerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRepoTriggerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRepoTriggerResponse
	GetStatusCode() *int32
	SetBody(v *CreateRepoTriggerResponseBody) *CreateRepoTriggerResponse
	GetBody() *CreateRepoTriggerResponseBody
}

type CreateRepoTriggerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRepoTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepoTriggerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRepoTriggerResponse) GoString() string {
	return s.String()
}

func (s *CreateRepoTriggerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRepoTriggerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRepoTriggerResponse) GetBody() *CreateRepoTriggerResponseBody {
	return s.Body
}

func (s *CreateRepoTriggerResponse) SetHeaders(v map[string]*string) *CreateRepoTriggerResponse {
	s.Headers = v
	return s
}

func (s *CreateRepoTriggerResponse) SetStatusCode(v int32) *CreateRepoTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepoTriggerResponse) SetBody(v *CreateRepoTriggerResponseBody) *CreateRepoTriggerResponse {
	s.Body = v
	return s
}

func (s *CreateRepoTriggerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
