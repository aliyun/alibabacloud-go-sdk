// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMemo(v string) *RegisterAccountRequest
	GetMemo() *string
	SetRegionId(v string) *RegisterAccountRequest
	GetRegionId() *string
}

type RegisterAccountRequest struct {
	// example:
	//
	// "user api register"
	Memo *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RegisterAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterAccountRequest) GoString() string {
	return s.String()
}

func (s *RegisterAccountRequest) GetMemo() *string {
	return s.Memo
}

func (s *RegisterAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RegisterAccountRequest) SetMemo(v string) *RegisterAccountRequest {
	s.Memo = &v
	return s
}

func (s *RegisterAccountRequest) SetRegionId(v string) *RegisterAccountRequest {
	s.RegionId = &v
	return s
}

func (s *RegisterAccountRequest) Validate() error {
	return dara.Validate(s)
}
