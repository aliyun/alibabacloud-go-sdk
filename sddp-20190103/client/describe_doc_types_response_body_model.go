// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDocTypeList(v []*DescribeDocTypesResponseBodyDocTypeList) *DescribeDocTypesResponseBody
	GetDocTypeList() []*DescribeDocTypesResponseBodyDocTypeList
	SetRequestId(v string) *DescribeDocTypesResponseBody
	GetRequestId() *string
}

type DescribeDocTypesResponseBody struct {
	// A list of OSS object types.
	DocTypeList []*DescribeDocTypesResponseBodyDocTypeList `json:"DocTypeList,omitempty" xml:"DocTypeList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-4******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDocTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDocTypesResponseBody) GetDocTypeList() []*DescribeDocTypesResponseBodyDocTypeList {
	return s.DocTypeList
}

func (s *DescribeDocTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDocTypesResponseBody) SetDocTypeList(v []*DescribeDocTypesResponseBodyDocTypeList) *DescribeDocTypesResponseBody {
	s.DocTypeList = v
	return s
}

func (s *DescribeDocTypesResponseBody) SetRequestId(v string) *DescribeDocTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDocTypesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDocTypesResponseBodyDocTypeList struct {
	// The code of the object type.
	//
	// example:
	//
	// 100001
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the object type.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the object type.
	//
	// example:
	//
	// C/C++ Source Code
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDocTypesResponseBodyDocTypeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocTypesResponseBodyDocTypeList) GoString() string {
	return s.String()
}

func (s *DescribeDocTypesResponseBodyDocTypeList) GetCode() *int64 {
	return s.Code
}

func (s *DescribeDocTypesResponseBodyDocTypeList) GetId() *int64 {
	return s.Id
}

func (s *DescribeDocTypesResponseBodyDocTypeList) GetName() *string {
	return s.Name
}

func (s *DescribeDocTypesResponseBodyDocTypeList) SetCode(v int64) *DescribeDocTypesResponseBodyDocTypeList {
	s.Code = &v
	return s
}

func (s *DescribeDocTypesResponseBodyDocTypeList) SetId(v int64) *DescribeDocTypesResponseBodyDocTypeList {
	s.Id = &v
	return s
}

func (s *DescribeDocTypesResponseBodyDocTypeList) SetName(v string) *DescribeDocTypesResponseBodyDocTypeList {
	s.Name = &v
	return s
}

func (s *DescribeDocTypesResponseBodyDocTypeList) Validate() error {
	return dara.Validate(s)
}
