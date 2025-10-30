// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKubernetesTriggerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKubernetesTriggerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKubernetesTriggerResponse
	GetStatusCode() *int32
	SetBody(v *CreateKubernetesTriggerResponseBody) *CreateKubernetesTriggerResponse
	GetBody() *CreateKubernetesTriggerResponseBody
}

type CreateKubernetesTriggerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKubernetesTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKubernetesTriggerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKubernetesTriggerResponse) GoString() string {
	return s.String()
}

func (s *CreateKubernetesTriggerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKubernetesTriggerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKubernetesTriggerResponse) GetBody() *CreateKubernetesTriggerResponseBody {
	return s.Body
}

func (s *CreateKubernetesTriggerResponse) SetHeaders(v map[string]*string) *CreateKubernetesTriggerResponse {
	s.Headers = v
	return s
}

func (s *CreateKubernetesTriggerResponse) SetStatusCode(v int32) *CreateKubernetesTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKubernetesTriggerResponse) SetBody(v *CreateKubernetesTriggerResponseBody) *CreateKubernetesTriggerResponse {
	s.Body = v
	return s
}

func (s *CreateKubernetesTriggerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
