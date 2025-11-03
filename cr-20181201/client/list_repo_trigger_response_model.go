// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoTriggerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRepoTriggerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRepoTriggerResponse
	GetStatusCode() *int32
	SetBody(v *ListRepoTriggerResponseBody) *ListRepoTriggerResponse
	GetBody() *ListRepoTriggerResponseBody
}

type ListRepoTriggerResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepoTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepoTriggerResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRepoTriggerResponse) GoString() string {
	return s.String()
}

func (s *ListRepoTriggerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRepoTriggerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRepoTriggerResponse) GetBody() *ListRepoTriggerResponseBody {
	return s.Body
}

func (s *ListRepoTriggerResponse) SetHeaders(v map[string]*string) *ListRepoTriggerResponse {
	s.Headers = v
	return s
}

func (s *ListRepoTriggerResponse) SetStatusCode(v int32) *ListRepoTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepoTriggerResponse) SetBody(v *ListRepoTriggerResponseBody) *ListRepoTriggerResponse {
	s.Body = v
	return s
}

func (s *ListRepoTriggerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
