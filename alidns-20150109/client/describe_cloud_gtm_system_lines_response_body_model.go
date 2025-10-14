// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmSystemLinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCloudGtmSystemLinesResponseBody
	GetRequestId() *string
	SetSystemLines(v *DescribeCloudGtmSystemLinesResponseBodySystemLines) *DescribeCloudGtmSystemLinesResponseBody
	GetSystemLines() *DescribeCloudGtmSystemLinesResponseBodySystemLines
	SetSystemLinesTree(v string) *DescribeCloudGtmSystemLinesResponseBody
	GetSystemLinesTree() *string
}

type DescribeCloudGtmSystemLinesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 389DFFA3-77A5-4A9E-BF3D-147C6F98A5BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The system lines.
	SystemLines *DescribeCloudGtmSystemLinesResponseBodySystemLines `json:"SystemLines,omitempty" xml:"SystemLines,omitempty" type:"Struct"`
	// The system lines, which are in a tree structure. Only a system line is listed in this example.
	//
	// example:
	//
	// [{\\"displayName\\":\\"Default\\",\\"id\\":\\"default\\",\\"isAvailable\\":true,\\"name\\":\\"Default\\",\\"parentId\\":\\"\\"}]
	SystemLinesTree *string `json:"SystemLinesTree,omitempty" xml:"SystemLinesTree,omitempty"`
}

func (s DescribeCloudGtmSystemLinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmSystemLinesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmSystemLinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudGtmSystemLinesResponseBody) GetSystemLines() *DescribeCloudGtmSystemLinesResponseBodySystemLines {
	return s.SystemLines
}

func (s *DescribeCloudGtmSystemLinesResponseBody) GetSystemLinesTree() *string {
	return s.SystemLinesTree
}

func (s *DescribeCloudGtmSystemLinesResponseBody) SetRequestId(v string) *DescribeCloudGtmSystemLinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudGtmSystemLinesResponseBody) SetSystemLines(v *DescribeCloudGtmSystemLinesResponseBodySystemLines) *DescribeCloudGtmSystemLinesResponseBody {
	s.SystemLines = v
	return s
}

func (s *DescribeCloudGtmSystemLinesResponseBody) SetSystemLinesTree(v string) *DescribeCloudGtmSystemLinesResponseBody {
	s.SystemLinesTree = &v
	return s
}

func (s *DescribeCloudGtmSystemLinesResponseBody) Validate() error {
	if s.SystemLines != nil {
		if err := s.SystemLines.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudGtmSystemLinesResponseBodySystemLines struct {
	SystemLine []*DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine `json:"SystemLine,omitempty" xml:"SystemLine,omitempty" type:"Repeated"`
}

func (s DescribeCloudGtmSystemLinesResponseBodySystemLines) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmSystemLinesResponseBodySystemLines) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmSystemLinesResponseBodySystemLines) GetSystemLine() []*DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine {
	return s.SystemLine
}

func (s *DescribeCloudGtmSystemLinesResponseBodySystemLines) SetSystemLine(v []*DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine) *DescribeCloudGtmSystemLinesResponseBodySystemLines {
	s.SystemLine = v
	return s
}

func (s *DescribeCloudGtmSystemLinesResponseBodySystemLines) Validate() error {
	if s.SystemLine != nil {
		for _, item := range s.SystemLine {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine struct {
	// The line code.
	//
	// example:
	//
	// aliyun_r_ap-south-1
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The display name of the line.
	//
	// example:
	//
	// Default
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Indicates whether the line can be selected as the source of a Domain Name System (DNS) request. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsAvailable *bool `json:"IsAvailable,omitempty" xml:"IsAvailable,omitempty"`
	// The name of the line.
	//
	// example:
	//
	// Default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The code of the parent line.
	//
	// example:
	//
	// String	aliyun
	ParentCode *string `json:"ParentCode,omitempty" xml:"ParentCode,omitempty"`
}

func (s DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine) GetCode() *string {
	return s.Code
}

func (s *DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine) GetIsAvailable() *bool {
	return s.IsAvailable
}

func (s *DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine) GetName() *string {
	return s.Name
}

func (s *DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine) GetParentCode() *string {
	return s.ParentCode
}

func (s *DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine) SetCode(v string) *DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine {
	s.Code = &v
	return s
}

func (s *DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine) SetDisplayName(v string) *DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine {
	s.DisplayName = &v
	return s
}

func (s *DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine) SetIsAvailable(v bool) *DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine {
	s.IsAvailable = &v
	return s
}

func (s *DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine) SetName(v string) *DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine {
	s.Name = &v
	return s
}

func (s *DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine) SetParentCode(v string) *DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine {
	s.ParentCode = &v
	return s
}

func (s *DescribeCloudGtmSystemLinesResponseBodySystemLinesSystemLine) Validate() error {
	return dara.Validate(s)
}
