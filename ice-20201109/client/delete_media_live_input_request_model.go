// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaLiveInputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputId(v string) *DeleteMediaLiveInputRequest
	GetInputId() *string
}

type DeleteMediaLiveInputRequest struct {
	// The ID of the input.
	//
	// This parameter is required.
	//
	// example:
	//
	// SEGK5KA6KYKAWQQH
	InputId *string `json:"InputId,omitempty" xml:"InputId,omitempty"`
}

func (s DeleteMediaLiveInputRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaLiveInputRequest) GoString() string {
	return s.String()
}

func (s *DeleteMediaLiveInputRequest) GetInputId() *string {
	return s.InputId
}

func (s *DeleteMediaLiveInputRequest) SetInputId(v string) *DeleteMediaLiveInputRequest {
	s.InputId = &v
	return s
}

func (s *DeleteMediaLiveInputRequest) Validate() error {
	return dara.Validate(s)
}
