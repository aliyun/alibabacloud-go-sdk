// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuccessInfoValue interface {
	dara.Model
	String() string
	GoString() string
	SetSuccess(v bool) *SuccessInfoValue
	GetSuccess() *bool
	SetMessage(v string) *SuccessInfoValue
	GetMessage() *string
}

type SuccessInfoValue struct {
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The error message.
	//
	// example:
	//
	// The task does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SuccessInfoValue) String() string {
	return dara.Prettify(s)
}

func (s SuccessInfoValue) GoString() string {
	return s.String()
}

func (s *SuccessInfoValue) GetSuccess() *bool {
	return s.Success
}

func (s *SuccessInfoValue) GetMessage() *string {
	return s.Message
}

func (s *SuccessInfoValue) SetSuccess(v bool) *SuccessInfoValue {
	s.Success = &v
	return s
}

func (s *SuccessInfoValue) SetMessage(v string) *SuccessInfoValue {
	s.Message = &v
	return s
}

func (s *SuccessInfoValue) Validate() error {
	return dara.Validate(s)
}
