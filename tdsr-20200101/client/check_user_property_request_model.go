// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUserPropertyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUid(v string) *CheckUserPropertyRequest
	GetUid() *string
}

type CheckUserPropertyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2345****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s CheckUserPropertyRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckUserPropertyRequest) GoString() string {
	return s.String()
}

func (s *CheckUserPropertyRequest) GetUid() *string {
	return s.Uid
}

func (s *CheckUserPropertyRequest) SetUid(v string) *CheckUserPropertyRequest {
	s.Uid = &v
	return s
}

func (s *CheckUserPropertyRequest) Validate() error {
	return dara.Validate(s)
}
