// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReceiverResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateReceiverResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateReceiverResponse
	GetStatusCode() *int32
}

type CreateReceiverResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateReceiverResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateReceiverResponse) GoString() string {
	return s.String()
}

func (s *CreateReceiverResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateReceiverResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateReceiverResponse) SetHeaders(v map[string]*string) *CreateReceiverResponse {
	s.Headers = v
	return s
}

func (s *CreateReceiverResponse) SetStatusCode(v int32) *CreateReceiverResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateReceiverResponse) Validate() error {
	return dara.Validate(s)
}
