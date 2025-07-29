// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTriggerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTriggerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTriggerResponse
	GetStatusCode() *int32
}

type DeleteTriggerResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteTriggerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTriggerResponse) GoString() string {
	return s.String()
}

func (s *DeleteTriggerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTriggerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTriggerResponse) SetHeaders(v map[string]*string) *DeleteTriggerResponse {
	s.Headers = v
	return s
}

func (s *DeleteTriggerResponse) SetStatusCode(v int32) *DeleteTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTriggerResponse) Validate() error {
	return dara.Validate(s)
}
