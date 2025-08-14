// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaLiveInputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInputId(v string) *CreateMediaLiveInputResponseBody
	GetInputId() *string
	SetRequestId(v string) *CreateMediaLiveInputResponseBody
	GetRequestId() *string
}

type CreateMediaLiveInputResponseBody struct {
	// The ID of the input.
	//
	// example:
	//
	// SEGK5KA6KYKAWQQH
	InputId *string `json:"InputId,omitempty" xml:"InputId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMediaLiveInputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveInputResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveInputResponseBody) GetInputId() *string {
	return s.InputId
}

func (s *CreateMediaLiveInputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMediaLiveInputResponseBody) SetInputId(v string) *CreateMediaLiveInputResponseBody {
	s.InputId = &v
	return s
}

func (s *CreateMediaLiveInputResponseBody) SetRequestId(v string) *CreateMediaLiveInputResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMediaLiveInputResponseBody) Validate() error {
	return dara.Validate(s)
}
