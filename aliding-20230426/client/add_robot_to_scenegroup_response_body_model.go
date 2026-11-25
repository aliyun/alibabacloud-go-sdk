// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRobotToScenegroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddRobotToScenegroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddRobotToScenegroupResponseBody
	GetSuccess() *bool
}

type AddRobotToScenegroupResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddRobotToScenegroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddRobotToScenegroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddRobotToScenegroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddRobotToScenegroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddRobotToScenegroupResponseBody) SetRequestId(v string) *AddRobotToScenegroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRobotToScenegroupResponseBody) SetSuccess(v bool) *AddRobotToScenegroupResponseBody {
	s.Success = &v
	return s
}

func (s *AddRobotToScenegroupResponseBody) Validate() error {
	return dara.Validate(s)
}
