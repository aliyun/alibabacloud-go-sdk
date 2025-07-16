// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTeamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTeamsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTeamsResponseBody
	GetRequestId() *string
	SetTeams(v []*ListTeamsResponseBodyTeams) *ListTeamsResponseBody
	GetTeams() []*ListTeamsResponseBodyTeams
	SetVendorRequestId(v string) *ListTeamsResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *ListTeamsResponseBody
	GetVendorType() *string
}

type ListTeamsResponseBody struct {
	// example:
	//
	// 2023-05-15T11:29Z
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Teams     []*ListTeamsResponseBodyTeams `json:"teams,omitempty" xml:"teams,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s ListTeamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTeamsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTeamsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTeamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTeamsResponseBody) GetTeams() []*ListTeamsResponseBodyTeams {
	return s.Teams
}

func (s *ListTeamsResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *ListTeamsResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *ListTeamsResponseBody) SetNextToken(v string) *ListTeamsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTeamsResponseBody) SetRequestId(v string) *ListTeamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTeamsResponseBody) SetTeams(v []*ListTeamsResponseBodyTeams) *ListTeamsResponseBody {
	s.Teams = v
	return s
}

func (s *ListTeamsResponseBody) SetVendorRequestId(v string) *ListTeamsResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *ListTeamsResponseBody) SetVendorType(v string) *ListTeamsResponseBody {
	s.VendorType = &v
	return s
}

func (s *ListTeamsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTeamsResponseBodyTeams struct {
	// example:
	//
	// ding16b241fd05********288
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// example:
	//
	// https://example/file-manage-files/zh-CN/202***13/ldet/XXXX.jpg
	Cover *string `json:"Cover,omitempty" xml:"Cover,omitempty"`
	// example:
	//
	// 01472825524039877041
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 012345
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// example:
	//
	// team_description
	Description *string                         `json:"Description,omitempty" xml:"Description,omitempty"`
	Icon        *ListTeamsResponseBodyTeamsIcon `json:"Icon,omitempty" xml:"Icon,omitempty" type:"Struct"`
	// example:
	//
	// 2023-05-15T11:29Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// 012345
	ModifierId *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// example:
	//
	// team_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// lHiicjNFM2iSFYSdz2iPuI8ZwiEiE
	TeamId *string `json:"TeamId,omitempty" xml:"TeamId,omitempty"`
}

func (s ListTeamsResponseBodyTeams) String() string {
	return dara.Prettify(s)
}

func (s ListTeamsResponseBodyTeams) GoString() string {
	return s.String()
}

func (s *ListTeamsResponseBodyTeams) GetCorpId() *string {
	return s.CorpId
}

func (s *ListTeamsResponseBodyTeams) GetCover() *string {
	return s.Cover
}

func (s *ListTeamsResponseBodyTeams) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListTeamsResponseBodyTeams) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListTeamsResponseBodyTeams) GetDescription() *string {
	return s.Description
}

func (s *ListTeamsResponseBodyTeams) GetIcon() *ListTeamsResponseBodyTeamsIcon {
	return s.Icon
}

func (s *ListTeamsResponseBodyTeams) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListTeamsResponseBodyTeams) GetModifierId() *string {
	return s.ModifierId
}

func (s *ListTeamsResponseBodyTeams) GetName() *string {
	return s.Name
}

func (s *ListTeamsResponseBodyTeams) GetTeamId() *string {
	return s.TeamId
}

func (s *ListTeamsResponseBodyTeams) SetCorpId(v string) *ListTeamsResponseBodyTeams {
	s.CorpId = &v
	return s
}

func (s *ListTeamsResponseBodyTeams) SetCover(v string) *ListTeamsResponseBodyTeams {
	s.Cover = &v
	return s
}

func (s *ListTeamsResponseBodyTeams) SetCreateTime(v string) *ListTeamsResponseBodyTeams {
	s.CreateTime = &v
	return s
}

func (s *ListTeamsResponseBodyTeams) SetCreatorId(v string) *ListTeamsResponseBodyTeams {
	s.CreatorId = &v
	return s
}

func (s *ListTeamsResponseBodyTeams) SetDescription(v string) *ListTeamsResponseBodyTeams {
	s.Description = &v
	return s
}

func (s *ListTeamsResponseBodyTeams) SetIcon(v *ListTeamsResponseBodyTeamsIcon) *ListTeamsResponseBodyTeams {
	s.Icon = v
	return s
}

func (s *ListTeamsResponseBodyTeams) SetModifiedTime(v string) *ListTeamsResponseBodyTeams {
	s.ModifiedTime = &v
	return s
}

func (s *ListTeamsResponseBodyTeams) SetModifierId(v string) *ListTeamsResponseBodyTeams {
	s.ModifierId = &v
	return s
}

func (s *ListTeamsResponseBodyTeams) SetName(v string) *ListTeamsResponseBodyTeams {
	s.Name = &v
	return s
}

func (s *ListTeamsResponseBodyTeams) SetTeamId(v string) *ListTeamsResponseBodyTeams {
	s.TeamId = &v
	return s
}

func (s *ListTeamsResponseBodyTeams) Validate() error {
	return dara.Validate(s)
}

type ListTeamsResponseBodyTeamsIcon struct {
	// example:
	//
	// URL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://example/file-manage-files/zh-CN/202***13/ldet/avatar3.jpg
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTeamsResponseBodyTeamsIcon) String() string {
	return dara.Prettify(s)
}

func (s ListTeamsResponseBodyTeamsIcon) GoString() string {
	return s.String()
}

func (s *ListTeamsResponseBodyTeamsIcon) GetType() *string {
	return s.Type
}

func (s *ListTeamsResponseBodyTeamsIcon) GetValue() *string {
	return s.Value
}

func (s *ListTeamsResponseBodyTeamsIcon) SetType(v string) *ListTeamsResponseBodyTeamsIcon {
	s.Type = &v
	return s
}

func (s *ListTeamsResponseBodyTeamsIcon) SetValue(v string) *ListTeamsResponseBodyTeamsIcon {
	s.Value = &v
	return s
}

func (s *ListTeamsResponseBodyTeamsIcon) Validate() error {
	return dara.Validate(s)
}
