// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDropReceiverResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DropReceiverResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DropReceiverResponse
	GetStatusCode() *int32
}

type DropReceiverResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DropReceiverResponse) String() string {
	return dara.Prettify(s)
}

func (s DropReceiverResponse) GoString() string {
	return s.String()
}

func (s *DropReceiverResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DropReceiverResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DropReceiverResponse) SetHeaders(v map[string]*string) *DropReceiverResponse {
	s.Headers = v
	return s
}

func (s *DropReceiverResponse) SetStatusCode(v int32) *DropReceiverResponse {
	s.StatusCode = &v
	return s
}

func (s *DropReceiverResponse) Validate() error {
	return dara.Validate(s)
}
