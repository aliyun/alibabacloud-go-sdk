// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaLiveInputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputId(v string) *GetMediaLiveInputRequest
	GetInputId() *string
}

type GetMediaLiveInputRequest struct {
	// The ID of the input.
	//
	// This parameter is required.
	//
	// example:
	//
	// SEGK5KA6KYKAWQQH
	InputId *string `json:"InputId,omitempty" xml:"InputId,omitempty"`
}

func (s GetMediaLiveInputRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaLiveInputRequest) GoString() string {
	return s.String()
}

func (s *GetMediaLiveInputRequest) GetInputId() *string {
	return s.InputId
}

func (s *GetMediaLiveInputRequest) SetInputId(v string) *GetMediaLiveInputRequest {
	s.InputId = &v
	return s
}

func (s *GetMediaLiveInputRequest) Validate() error {
	return dara.Validate(s)
}
