// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDriveSpaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *AddDriveSpaceResponseBody
	GetCreateTime() *string
	SetModifyTime(v string) *AddDriveSpaceResponseBody
	GetModifyTime() *string
	SetPermissionMode(v string) *AddDriveSpaceResponseBody
	GetPermissionMode() *string
	SetQuota(v int64) *AddDriveSpaceResponseBody
	GetQuota() *int64
	SetRequestId(v string) *AddDriveSpaceResponseBody
	GetRequestId() *string
	SetSpaceId(v string) *AddDriveSpaceResponseBody
	GetSpaceId() *string
	SetSpaceName(v string) *AddDriveSpaceResponseBody
	GetSpaceName() *string
	SetSpaceType(v string) *AddDriveSpaceResponseBody
	GetSpaceType() *string
	SetUsedQuota(v int64) *AddDriveSpaceResponseBody
	GetUsedQuota() *int64
	SetVendorRequestId(v string) *AddDriveSpaceResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *AddDriveSpaceResponseBody
	GetVendorType() *string
}

type AddDriveSpaceResponseBody struct {
	CreateTime     *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	ModifyTime     *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	PermissionMode *string `json:"permissionMode,omitempty" xml:"permissionMode,omitempty"`
	Quota          *int64  `json:"quota,omitempty" xml:"quota,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SpaceId   *string `json:"spaceId,omitempty" xml:"spaceId,omitempty"`
	SpaceName *string `json:"spaceName,omitempty" xml:"spaceName,omitempty"`
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	UsedQuota *int64  `json:"usedQuota,omitempty" xml:"usedQuota,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s AddDriveSpaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDriveSpaceResponseBody) GoString() string {
	return s.String()
}

func (s *AddDriveSpaceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *AddDriveSpaceResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *AddDriveSpaceResponseBody) GetPermissionMode() *string {
	return s.PermissionMode
}

func (s *AddDriveSpaceResponseBody) GetQuota() *int64 {
	return s.Quota
}

func (s *AddDriveSpaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDriveSpaceResponseBody) GetSpaceId() *string {
	return s.SpaceId
}

func (s *AddDriveSpaceResponseBody) GetSpaceName() *string {
	return s.SpaceName
}

func (s *AddDriveSpaceResponseBody) GetSpaceType() *string {
	return s.SpaceType
}

func (s *AddDriveSpaceResponseBody) GetUsedQuota() *int64 {
	return s.UsedQuota
}

func (s *AddDriveSpaceResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *AddDriveSpaceResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *AddDriveSpaceResponseBody) SetCreateTime(v string) *AddDriveSpaceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *AddDriveSpaceResponseBody) SetModifyTime(v string) *AddDriveSpaceResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *AddDriveSpaceResponseBody) SetPermissionMode(v string) *AddDriveSpaceResponseBody {
	s.PermissionMode = &v
	return s
}

func (s *AddDriveSpaceResponseBody) SetQuota(v int64) *AddDriveSpaceResponseBody {
	s.Quota = &v
	return s
}

func (s *AddDriveSpaceResponseBody) SetRequestId(v string) *AddDriveSpaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDriveSpaceResponseBody) SetSpaceId(v string) *AddDriveSpaceResponseBody {
	s.SpaceId = &v
	return s
}

func (s *AddDriveSpaceResponseBody) SetSpaceName(v string) *AddDriveSpaceResponseBody {
	s.SpaceName = &v
	return s
}

func (s *AddDriveSpaceResponseBody) SetSpaceType(v string) *AddDriveSpaceResponseBody {
	s.SpaceType = &v
	return s
}

func (s *AddDriveSpaceResponseBody) SetUsedQuota(v int64) *AddDriveSpaceResponseBody {
	s.UsedQuota = &v
	return s
}

func (s *AddDriveSpaceResponseBody) SetVendorRequestId(v string) *AddDriveSpaceResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *AddDriveSpaceResponseBody) SetVendorType(v string) *AddDriveSpaceResponseBody {
	s.VendorType = &v
	return s
}

func (s *AddDriveSpaceResponseBody) Validate() error {
	return dara.Validate(s)
}
