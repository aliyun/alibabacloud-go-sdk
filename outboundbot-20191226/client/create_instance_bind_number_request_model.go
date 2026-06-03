// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceBindNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceList(v string) *CreateInstanceBindNumberRequest
	GetInstanceList() *string
	SetNumber(v string) *CreateInstanceBindNumberRequest
	GetNumber() *string
}

type CreateInstanceBindNumberRequest struct {
	// example:
	//
	// 1,2,4,5
	InstanceList *string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty"`
	// example:
	//
	// 10088
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s CreateInstanceBindNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceBindNumberRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceBindNumberRequest) GetInstanceList() *string {
	return s.InstanceList
}

func (s *CreateInstanceBindNumberRequest) GetNumber() *string {
	return s.Number
}

func (s *CreateInstanceBindNumberRequest) SetInstanceList(v string) *CreateInstanceBindNumberRequest {
	s.InstanceList = &v
	return s
}

func (s *CreateInstanceBindNumberRequest) SetNumber(v string) *CreateInstanceBindNumberRequest {
	s.Number = &v
	return s
}

func (s *CreateInstanceBindNumberRequest) Validate() error {
	return dara.Validate(s)
}
