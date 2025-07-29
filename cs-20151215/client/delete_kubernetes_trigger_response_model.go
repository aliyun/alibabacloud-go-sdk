// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKubernetesTriggerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteKubernetesTriggerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteKubernetesTriggerResponse
	GetStatusCode() *int32
}

type DeleteKubernetesTriggerResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteKubernetesTriggerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteKubernetesTriggerResponse) GoString() string {
	return s.String()
}

func (s *DeleteKubernetesTriggerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteKubernetesTriggerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteKubernetesTriggerResponse) SetHeaders(v map[string]*string) *DeleteKubernetesTriggerResponse {
	s.Headers = v
	return s
}

func (s *DeleteKubernetesTriggerResponse) SetStatusCode(v int32) *DeleteKubernetesTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKubernetesTriggerResponse) Validate() error {
	return dara.Validate(s)
}
