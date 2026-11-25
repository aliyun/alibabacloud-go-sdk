// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRobotToScenegroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpenConversationId(v string) *AddRobotToScenegroupRequest
	GetOpenConversationId() *string
	SetRobotCode(v string) *AddRobotToScenegroupRequest
	GetRobotCode() *string
}

type AddRobotToScenegroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cidt*****Xa4K10w==
	OpenConversationId *string `json:"OpenConversationId,omitempty" xml:"OpenConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dingxxxxxxxxxxxx
	RobotCode *string `json:"RobotCode,omitempty" xml:"RobotCode,omitempty"`
}

func (s AddRobotToScenegroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRobotToScenegroupRequest) GoString() string {
	return s.String()
}

func (s *AddRobotToScenegroupRequest) GetOpenConversationId() *string {
	return s.OpenConversationId
}

func (s *AddRobotToScenegroupRequest) GetRobotCode() *string {
	return s.RobotCode
}

func (s *AddRobotToScenegroupRequest) SetOpenConversationId(v string) *AddRobotToScenegroupRequest {
	s.OpenConversationId = &v
	return s
}

func (s *AddRobotToScenegroupRequest) SetRobotCode(v string) *AddRobotToScenegroupRequest {
	s.RobotCode = &v
	return s
}

func (s *AddRobotToScenegroupRequest) Validate() error {
	return dara.Validate(s)
}
