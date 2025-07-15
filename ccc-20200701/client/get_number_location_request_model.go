// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNumberLocationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetNumberLocationRequest
	GetInstanceId() *string
	SetNumber(v string) *GetNumberLocationRequest
	GetNumber() *string
}

type GetNumberLocationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1312121****
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s GetNumberLocationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNumberLocationRequest) GoString() string {
	return s.String()
}

func (s *GetNumberLocationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetNumberLocationRequest) GetNumber() *string {
	return s.Number
}

func (s *GetNumberLocationRequest) SetInstanceId(v string) *GetNumberLocationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNumberLocationRequest) SetNumber(v string) *GetNumberLocationRequest {
	s.Number = &v
	return s
}

func (s *GetNumberLocationRequest) Validate() error {
	return dara.Validate(s)
}
