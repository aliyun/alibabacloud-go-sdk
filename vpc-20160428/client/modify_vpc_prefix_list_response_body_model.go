// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcPrefixListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrefixListId(v string) *ModifyVpcPrefixListResponseBody
	GetPrefixListId() *string
	SetRequestId(v string) *ModifyVpcPrefixListResponseBody
	GetRequestId() *string
}

type ModifyVpcPrefixListResponseBody struct {
	// The ID of the prefix list.
	//
	// example:
	//
	// pl-0b7hwu67****
	PrefixListId *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcPrefixListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcPrefixListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcPrefixListResponseBody) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *ModifyVpcPrefixListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVpcPrefixListResponseBody) SetPrefixListId(v string) *ModifyVpcPrefixListResponseBody {
	s.PrefixListId = &v
	return s
}

func (s *ModifyVpcPrefixListResponseBody) SetRequestId(v string) *ModifyVpcPrefixListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVpcPrefixListResponseBody) Validate() error {
	return dara.Validate(s)
}
