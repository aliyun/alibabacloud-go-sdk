// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterReceiverResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AlterReceiverResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AlterReceiverResponse
	GetStatusCode() *int32
}

type AlterReceiverResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s AlterReceiverResponse) String() string {
	return dara.Prettify(s)
}

func (s AlterReceiverResponse) GoString() string {
	return s.String()
}

func (s *AlterReceiverResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AlterReceiverResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AlterReceiverResponse) SetHeaders(v map[string]*string) *AlterReceiverResponse {
	s.Headers = v
	return s
}

func (s *AlterReceiverResponse) SetStatusCode(v int32) *AlterReceiverResponse {
	s.StatusCode = &v
	return s
}

func (s *AlterReceiverResponse) Validate() error {
	return dara.Validate(s)
}
