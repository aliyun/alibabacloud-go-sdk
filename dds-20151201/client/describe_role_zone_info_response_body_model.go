// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRoleZoneInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRoleZoneInfoResponseBody
	GetRequestId() *string
	SetZoneInfos(v *DescribeRoleZoneInfoResponseBodyZoneInfos) *DescribeRoleZoneInfoResponseBody
	GetZoneInfos() *DescribeRoleZoneInfoResponseBodyZoneInfos
}

type DescribeRoleZoneInfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 728B9A96-E262-4AE5-915E-3A51CCE2FDA9
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ZoneInfos *DescribeRoleZoneInfoResponseBodyZoneInfos `json:"ZoneInfos,omitempty" xml:"ZoneInfos,omitempty" type:"Struct"`
}

func (s DescribeRoleZoneInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoleZoneInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRoleZoneInfoResponseBody) GetZoneInfos() *DescribeRoleZoneInfoResponseBodyZoneInfos {
	return s.ZoneInfos
}

func (s *DescribeRoleZoneInfoResponseBody) SetRequestId(v string) *DescribeRoleZoneInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBody) SetZoneInfos(v *DescribeRoleZoneInfoResponseBodyZoneInfos) *DescribeRoleZoneInfoResponseBody {
	s.ZoneInfos = v
	return s
}

func (s *DescribeRoleZoneInfoResponseBody) Validate() error {
	if s.ZoneInfos != nil {
		if err := s.ZoneInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRoleZoneInfoResponseBodyZoneInfos struct {
	ZoneInfo []*DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo `json:"ZoneInfo,omitempty" xml:"ZoneInfo,omitempty" type:"Repeated"`
}

func (s DescribeRoleZoneInfoResponseBodyZoneInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoleZoneInfoResponseBodyZoneInfos) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoResponseBodyZoneInfos) GetZoneInfo() []*DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo {
	return s.ZoneInfo
}

func (s *DescribeRoleZoneInfoResponseBodyZoneInfos) SetZoneInfo(v []*DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) *DescribeRoleZoneInfoResponseBodyZoneInfos {
	s.ZoneInfo = v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyZoneInfos) Validate() error {
	if s.ZoneInfo != nil {
		for _, item := range s.ZoneInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo struct {
	InsName  *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	RoleId   *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	ZoneId   *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) GoString() string {
	return s.String()
}

func (s *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) GetInsName() *string {
	return s.InsName
}

func (s *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) GetRoleId() *string {
	return s.RoleId
}

func (s *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) GetRoleType() *string {
	return s.RoleType
}

func (s *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) SetInsName(v string) *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo {
	s.InsName = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) SetNodeType(v string) *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo {
	s.NodeType = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) SetRoleId(v string) *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo {
	s.RoleId = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) SetRoleType(v string) *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo {
	s.RoleType = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) SetZoneId(v string) *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo {
	s.ZoneId = &v
	return s
}

func (s *DescribeRoleZoneInfoResponseBodyZoneInfosZoneInfo) Validate() error {
	return dara.Validate(s)
}
