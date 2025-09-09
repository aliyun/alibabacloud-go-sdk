// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackMenuResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v *DescribeBackMenuResponseBodyList) *DescribeBackMenuResponseBody
	GetList() *DescribeBackMenuResponseBodyList
	SetRequestId(v string) *DescribeBackMenuResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBackMenuResponseBody
	GetSuccess() *bool
}

type DescribeBackMenuResponseBody struct {
	// The backup information list.
	List *DescribeBackMenuResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 60C21BE4-EDFE-454C-95ED-3A5C74******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of request.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeBackMenuResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackMenuResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackMenuResponseBody) GetList() *DescribeBackMenuResponseBodyList {
	return s.List
}

func (s *DescribeBackMenuResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackMenuResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBackMenuResponseBody) SetList(v *DescribeBackMenuResponseBodyList) *DescribeBackMenuResponseBody {
	s.List = v
	return s
}

func (s *DescribeBackMenuResponseBody) SetRequestId(v string) *DescribeBackMenuResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackMenuResponseBody) SetSuccess(v bool) *DescribeBackMenuResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackMenuResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBackMenuResponseBodyList struct {
	List []*DescribeBackMenuResponseBodyListList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
}

func (s DescribeBackMenuResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackMenuResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeBackMenuResponseBodyList) GetList() []*DescribeBackMenuResponseBodyListList {
	return s.List
}

func (s *DescribeBackMenuResponseBodyList) SetList(v []*DescribeBackMenuResponseBodyListList) *DescribeBackMenuResponseBodyList {
	s.List = v
	return s
}

func (s *DescribeBackMenuResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type DescribeBackMenuResponseBodyListList struct {
	// The backup method. Valid values:
	//
	// 	- **Logic **: logical backup
	//
	// 	- **phy**: physical backup
	//
	// example:
	//
	// phy
	MenuName *string `json:"MenuName,omitempty" xml:"MenuName,omitempty"`
	// Indicates whether backup recovery is supported.
	//
	// example:
	//
	// true
	Support *bool `json:"Support,omitempty" xml:"Support,omitempty"`
}

func (s DescribeBackMenuResponseBodyListList) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackMenuResponseBodyListList) GoString() string {
	return s.String()
}

func (s *DescribeBackMenuResponseBodyListList) GetMenuName() *string {
	return s.MenuName
}

func (s *DescribeBackMenuResponseBodyListList) GetSupport() *bool {
	return s.Support
}

func (s *DescribeBackMenuResponseBodyListList) SetMenuName(v string) *DescribeBackMenuResponseBodyListList {
	s.MenuName = &v
	return s
}

func (s *DescribeBackMenuResponseBodyListList) SetSupport(v bool) *DescribeBackMenuResponseBodyListList {
	s.Support = &v
	return s
}

func (s *DescribeBackMenuResponseBodyListList) Validate() error {
	return dara.Validate(s)
}
