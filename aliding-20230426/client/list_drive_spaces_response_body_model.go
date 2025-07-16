// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDriveSpacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListDriveSpacesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDriveSpacesResponseBody
	GetRequestId() *string
	SetSpaces(v []*ListDriveSpacesResponseBodySpaces) *ListDriveSpacesResponseBody
	GetSpaces() []*ListDriveSpacesResponseBodySpaces
	SetVendorRequestId(v string) *ListDriveSpacesResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *ListDriveSpacesResponseBody
	GetVendorType() *string
}

type ListDriveSpacesResponseBody struct {
	// example:
	//
	// fekaf
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Spaces    []*ListDriveSpacesResponseBodySpaces `json:"spaces,omitempty" xml:"spaces,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s ListDriveSpacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDriveSpacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDriveSpacesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDriveSpacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDriveSpacesResponseBody) GetSpaces() []*ListDriveSpacesResponseBodySpaces {
	return s.Spaces
}

func (s *ListDriveSpacesResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *ListDriveSpacesResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *ListDriveSpacesResponseBody) SetNextToken(v string) *ListDriveSpacesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDriveSpacesResponseBody) SetRequestId(v string) *ListDriveSpacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDriveSpacesResponseBody) SetSpaces(v []*ListDriveSpacesResponseBodySpaces) *ListDriveSpacesResponseBody {
	s.Spaces = v
	return s
}

func (s *ListDriveSpacesResponseBody) SetVendorRequestId(v string) *ListDriveSpacesResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *ListDriveSpacesResponseBody) SetVendorType(v string) *ListDriveSpacesResponseBody {
	s.VendorType = &v
	return s
}

func (s *ListDriveSpacesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDriveSpacesResponseBodySpaces struct {
	// example:
	//
	// 2016-02-28T10:47:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2016-02-28T10:47:08Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// acl
	PermissionMode *string `json:"PermissionMode,omitempty" xml:"PermissionMode,omitempty"`
	// example:
	//
	// 2147483648
	Quota *int64 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// example:
	//
	// 123456789
	SpaceId   *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	SpaceName *string `json:"SpaceName,omitempty" xml:"SpaceName,omitempty"`
	// example:
	//
	// org
	SpaceType *string `json:"SpaceType,omitempty" xml:"SpaceType,omitempty"`
	// example:
	//
	// 640445953
	UsedQuota *int64 `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty"`
}

func (s ListDriveSpacesResponseBodySpaces) String() string {
	return dara.Prettify(s)
}

func (s ListDriveSpacesResponseBodySpaces) GoString() string {
	return s.String()
}

func (s *ListDriveSpacesResponseBodySpaces) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDriveSpacesResponseBodySpaces) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListDriveSpacesResponseBodySpaces) GetPermissionMode() *string {
	return s.PermissionMode
}

func (s *ListDriveSpacesResponseBodySpaces) GetQuota() *int64 {
	return s.Quota
}

func (s *ListDriveSpacesResponseBodySpaces) GetSpaceId() *string {
	return s.SpaceId
}

func (s *ListDriveSpacesResponseBodySpaces) GetSpaceName() *string {
	return s.SpaceName
}

func (s *ListDriveSpacesResponseBodySpaces) GetSpaceType() *string {
	return s.SpaceType
}

func (s *ListDriveSpacesResponseBodySpaces) GetUsedQuota() *int64 {
	return s.UsedQuota
}

func (s *ListDriveSpacesResponseBodySpaces) SetCreateTime(v string) *ListDriveSpacesResponseBodySpaces {
	s.CreateTime = &v
	return s
}

func (s *ListDriveSpacesResponseBodySpaces) SetModifyTime(v string) *ListDriveSpacesResponseBodySpaces {
	s.ModifyTime = &v
	return s
}

func (s *ListDriveSpacesResponseBodySpaces) SetPermissionMode(v string) *ListDriveSpacesResponseBodySpaces {
	s.PermissionMode = &v
	return s
}

func (s *ListDriveSpacesResponseBodySpaces) SetQuota(v int64) *ListDriveSpacesResponseBodySpaces {
	s.Quota = &v
	return s
}

func (s *ListDriveSpacesResponseBodySpaces) SetSpaceId(v string) *ListDriveSpacesResponseBodySpaces {
	s.SpaceId = &v
	return s
}

func (s *ListDriveSpacesResponseBodySpaces) SetSpaceName(v string) *ListDriveSpacesResponseBodySpaces {
	s.SpaceName = &v
	return s
}

func (s *ListDriveSpacesResponseBodySpaces) SetSpaceType(v string) *ListDriveSpacesResponseBodySpaces {
	s.SpaceType = &v
	return s
}

func (s *ListDriveSpacesResponseBodySpaces) SetUsedQuota(v int64) *ListDriveSpacesResponseBodySpaces {
	s.UsedQuota = &v
	return s
}

func (s *ListDriveSpacesResponseBodySpaces) Validate() error {
	return dara.Validate(s)
}
