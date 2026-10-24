// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeAssetScopeResponseBodyData) *DescribeAssetScopeResponseBody
	GetData() *DescribeAssetScopeResponseBodyData
	SetRequestId(v string) *DescribeAssetScopeResponseBody
	GetRequestId() *string
}

type DescribeAssetScopeResponseBody struct {
	// The endpoint asset scope details.
	Data *DescribeAssetScopeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique identifier of the request, which is used to troubleshoot and locate issues.
	//
	// example:
	//
	// 01A157FA-F957-3A5A-A6B8-5BBB8ADB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAssetScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetScopeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetScopeResponseBody) GetData() *DescribeAssetScopeResponseBodyData {
	return s.Data
}

func (s *DescribeAssetScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAssetScopeResponseBody) SetData(v *DescribeAssetScopeResponseBodyData) *DescribeAssetScopeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAssetScopeResponseBody) SetRequestId(v string) *DescribeAssetScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetScopeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAssetScopeResponseBodyData struct {
	// **[Deprecated]*	- This field is deprecated.
	//
	// example:
	//
	// 1792800000000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// **[Deprecated]*	- This field is deprecated.
	//
	// example:
	//
	// 1792803600000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the asset scope record, which corresponds to the Id request parameter.
	//
	// example:
	//
	// 1932
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The effective scope mode. Valid values:
	//
	// - **UserGroupAll**: all user groups.
	//
	// - **UserGroupNormal**: specified user groups.
	//
	// - **DeviceGroupNormal**: specified device groups.
	//
	// - **DevTagNormal**: specified device tags.
	//
	// The meaning of MatchTargetIds is determined by this field.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The list of effective scope targets. All elements are strings. The meaning of each element is determined by MatchMode:
	//
	// - **UserGroupAll**: Scope matching does not depend on this list.
	//
	// - **UserGroupNormal**: the user group ID.
	//
	// - **DeviceGroupNormal**: the device group ID.
	//
	// - **DevTagNormal**: the device tag value.
	//
	// This list is not a list of matched endpoint UUIDs.
	MatchTargetIds []*string `json:"MatchTargetIds,omitempty" xml:"MatchTargetIds,omitempty" type:"Repeated"`
}

func (s DescribeAssetScopeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetScopeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAssetScopeResponseBodyData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeAssetScopeResponseBodyData) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeAssetScopeResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *DescribeAssetScopeResponseBodyData) GetMatchMode() *string {
	return s.MatchMode
}

func (s *DescribeAssetScopeResponseBodyData) GetMatchTargetIds() []*string {
	return s.MatchTargetIds
}

func (s *DescribeAssetScopeResponseBodyData) SetGmtCreate(v int64) *DescribeAssetScopeResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAssetScopeResponseBodyData) SetGmtModified(v int64) *DescribeAssetScopeResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DescribeAssetScopeResponseBodyData) SetId(v int64) *DescribeAssetScopeResponseBodyData {
	s.Id = &v
	return s
}

func (s *DescribeAssetScopeResponseBodyData) SetMatchMode(v string) *DescribeAssetScopeResponseBodyData {
	s.MatchMode = &v
	return s
}

func (s *DescribeAssetScopeResponseBodyData) SetMatchTargetIds(v []*string) *DescribeAssetScopeResponseBodyData {
	s.MatchTargetIds = v
	return s
}

func (s *DescribeAssetScopeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
