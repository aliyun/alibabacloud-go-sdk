// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReceiveMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWaitseconds(v int32) *ReceiveMessageRequest
	GetWaitseconds() *int32
}

type ReceiveMessageRequest struct {
	// example:
	//
	// 0
	Waitseconds *int32 `json:"waitseconds,omitempty" xml:"waitseconds,omitempty"`
}

func (s ReceiveMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s ReceiveMessageRequest) GoString() string {
	return s.String()
}

func (s *ReceiveMessageRequest) GetWaitseconds() *int32 {
	return s.Waitseconds
}

func (s *ReceiveMessageRequest) SetWaitseconds(v int32) *ReceiveMessageRequest {
	s.Waitseconds = &v
	return s
}

func (s *ReceiveMessageRequest) Validate() error {
	return dara.Validate(s)
}
