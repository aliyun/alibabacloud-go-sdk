// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeACLAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcrs(v *DescribeACLAttributeResponseBodyAcrs) *DescribeACLAttributeResponseBody
	GetAcrs() *DescribeACLAttributeResponseBodyAcrs
	SetPageNumber(v int32) *DescribeACLAttributeResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeACLAttributeResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeACLAttributeResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeACLAttributeResponseBody
	GetTotalCount() *int32
}

type DescribeACLAttributeResponseBody struct {
	Acrs *DescribeACLAttributeResponseBodyAcrs `json:"Acrs,omitempty" xml:"Acrs,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8F62CE77-FBA2-4F8D-AED9-0A02814EDA69
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeACLAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeACLAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeACLAttributeResponseBody) GetAcrs() *DescribeACLAttributeResponseBodyAcrs {
	return s.Acrs
}

func (s *DescribeACLAttributeResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeACLAttributeResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeACLAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeACLAttributeResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeACLAttributeResponseBody) SetAcrs(v *DescribeACLAttributeResponseBodyAcrs) *DescribeACLAttributeResponseBody {
	s.Acrs = v
	return s
}

func (s *DescribeACLAttributeResponseBody) SetPageNumber(v int32) *DescribeACLAttributeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeACLAttributeResponseBody) SetPageSize(v int32) *DescribeACLAttributeResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeACLAttributeResponseBody) SetRequestId(v string) *DescribeACLAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeACLAttributeResponseBody) SetTotalCount(v int32) *DescribeACLAttributeResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeACLAttributeResponseBody) Validate() error {
	if s.Acrs != nil {
		if err := s.Acrs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeACLAttributeResponseBodyAcrs struct {
	Acr []*DescribeACLAttributeResponseBodyAcrsAcr `json:"Acr,omitempty" xml:"Acr,omitempty" type:"Repeated"`
}

func (s DescribeACLAttributeResponseBodyAcrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeACLAttributeResponseBodyAcrs) GoString() string {
	return s.String()
}

func (s *DescribeACLAttributeResponseBodyAcrs) GetAcr() []*DescribeACLAttributeResponseBodyAcrsAcr {
	return s.Acr
}

func (s *DescribeACLAttributeResponseBodyAcrs) SetAcr(v []*DescribeACLAttributeResponseBodyAcrsAcr) *DescribeACLAttributeResponseBodyAcrs {
	s.Acr = v
	return s
}

func (s *DescribeACLAttributeResponseBodyAcrs) Validate() error {
	if s.Acr != nil {
		for _, item := range s.Acr {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeACLAttributeResponseBodyAcrsAcr struct {
	AclId           *string                                                 `json:"AclId,omitempty" xml:"AclId,omitempty"`
	AclType         *string                                                 `json:"AclType,omitempty" xml:"AclType,omitempty"`
	AcrId           *string                                                 `json:"AcrId,omitempty" xml:"AcrId,omitempty"`
	Description     *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	DestCidr        *string                                                 `json:"DestCidr,omitempty" xml:"DestCidr,omitempty"`
	DestPortRange   *string                                                 `json:"DestPortRange,omitempty" xml:"DestPortRange,omitempty"`
	Direction       *string                                                 `json:"Direction,omitempty" xml:"Direction,omitempty"`
	DpiGroupIds     *DescribeACLAttributeResponseBodyAcrsAcrDpiGroupIds     `json:"DpiGroupIds,omitempty" xml:"DpiGroupIds,omitempty" type:"Struct"`
	DpiSignatureIds *DescribeACLAttributeResponseBodyAcrsAcrDpiSignatureIds `json:"DpiSignatureIds,omitempty" xml:"DpiSignatureIds,omitempty" type:"Struct"`
	GmtCreate       *int64                                                  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	IpProtocol      *string                                                 `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	Name            *string                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	Policy          *string                                                 `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Priority        *int32                                                  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	SourceCidr      *string                                                 `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	SourcePortRange *string                                                 `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	Type            *string                                                 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeACLAttributeResponseBodyAcrsAcr) String() string {
	return dara.Prettify(s)
}

func (s DescribeACLAttributeResponseBodyAcrsAcr) GoString() string {
	return s.String()
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) GetAclId() *string {
	return s.AclId
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) GetAclType() *string {
	return s.AclType
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) GetAcrId() *string {
	return s.AcrId
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) GetDescription() *string {
	return s.Description
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) GetDestCidr() *string {
	return s.DestCidr
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) GetDestPortRange() *string {
	return s.DestPortRange
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) GetDirection() *string {
	return s.Direction
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) GetDpiGroupIds() *DescribeACLAttributeResponseBodyAcrsAcrDpiGroupIds {
	return s.DpiGroupIds
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) GetDpiSignatureIds() *DescribeACLAttributeResponseBodyAcrsAcrDpiSignatureIds {
	return s.DpiSignatureIds
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) GetName() *string {
	return s.Name
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) GetSourceCidr() *string {
	return s.SourceCidr
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) GetType() *string {
	return s.Type
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) SetAclId(v string) *DescribeACLAttributeResponseBodyAcrsAcr {
	s.AclId = &v
	return s
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) SetAclType(v string) *DescribeACLAttributeResponseBodyAcrsAcr {
	s.AclType = &v
	return s
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) SetAcrId(v string) *DescribeACLAttributeResponseBodyAcrsAcr {
	s.AcrId = &v
	return s
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) SetDescription(v string) *DescribeACLAttributeResponseBodyAcrsAcr {
	s.Description = &v
	return s
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) SetDestCidr(v string) *DescribeACLAttributeResponseBodyAcrsAcr {
	s.DestCidr = &v
	return s
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) SetDestPortRange(v string) *DescribeACLAttributeResponseBodyAcrsAcr {
	s.DestPortRange = &v
	return s
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) SetDirection(v string) *DescribeACLAttributeResponseBodyAcrsAcr {
	s.Direction = &v
	return s
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) SetDpiGroupIds(v *DescribeACLAttributeResponseBodyAcrsAcrDpiGroupIds) *DescribeACLAttributeResponseBodyAcrsAcr {
	s.DpiGroupIds = v
	return s
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) SetDpiSignatureIds(v *DescribeACLAttributeResponseBodyAcrsAcrDpiSignatureIds) *DescribeACLAttributeResponseBodyAcrsAcr {
	s.DpiSignatureIds = v
	return s
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) SetGmtCreate(v int64) *DescribeACLAttributeResponseBodyAcrsAcr {
	s.GmtCreate = &v
	return s
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) SetIpProtocol(v string) *DescribeACLAttributeResponseBodyAcrsAcr {
	s.IpProtocol = &v
	return s
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) SetName(v string) *DescribeACLAttributeResponseBodyAcrsAcr {
	s.Name = &v
	return s
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) SetPolicy(v string) *DescribeACLAttributeResponseBodyAcrsAcr {
	s.Policy = &v
	return s
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) SetPriority(v int32) *DescribeACLAttributeResponseBodyAcrsAcr {
	s.Priority = &v
	return s
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) SetSourceCidr(v string) *DescribeACLAttributeResponseBodyAcrsAcr {
	s.SourceCidr = &v
	return s
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) SetSourcePortRange(v string) *DescribeACLAttributeResponseBodyAcrsAcr {
	s.SourcePortRange = &v
	return s
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) SetType(v string) *DescribeACLAttributeResponseBodyAcrsAcr {
	s.Type = &v
	return s
}

func (s *DescribeACLAttributeResponseBodyAcrsAcr) Validate() error {
	if s.DpiGroupIds != nil {
		if err := s.DpiGroupIds.Validate(); err != nil {
			return err
		}
	}
	if s.DpiSignatureIds != nil {
		if err := s.DpiSignatureIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeACLAttributeResponseBodyAcrsAcrDpiGroupIds struct {
	DpiGroupId []*string `json:"DpiGroupId,omitempty" xml:"DpiGroupId,omitempty" type:"Repeated"`
}

func (s DescribeACLAttributeResponseBodyAcrsAcrDpiGroupIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeACLAttributeResponseBodyAcrsAcrDpiGroupIds) GoString() string {
	return s.String()
}

func (s *DescribeACLAttributeResponseBodyAcrsAcrDpiGroupIds) GetDpiGroupId() []*string {
	return s.DpiGroupId
}

func (s *DescribeACLAttributeResponseBodyAcrsAcrDpiGroupIds) SetDpiGroupId(v []*string) *DescribeACLAttributeResponseBodyAcrsAcrDpiGroupIds {
	s.DpiGroupId = v
	return s
}

func (s *DescribeACLAttributeResponseBodyAcrsAcrDpiGroupIds) Validate() error {
	return dara.Validate(s)
}

type DescribeACLAttributeResponseBodyAcrsAcrDpiSignatureIds struct {
	DpiSignatureId []*string `json:"DpiSignatureId,omitempty" xml:"DpiSignatureId,omitempty" type:"Repeated"`
}

func (s DescribeACLAttributeResponseBodyAcrsAcrDpiSignatureIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeACLAttributeResponseBodyAcrsAcrDpiSignatureIds) GoString() string {
	return s.String()
}

func (s *DescribeACLAttributeResponseBodyAcrsAcrDpiSignatureIds) GetDpiSignatureId() []*string {
	return s.DpiSignatureId
}

func (s *DescribeACLAttributeResponseBodyAcrsAcrDpiSignatureIds) SetDpiSignatureId(v []*string) *DescribeACLAttributeResponseBodyAcrsAcrDpiSignatureIds {
	s.DpiSignatureId = v
	return s
}

func (s *DescribeACLAttributeResponseBodyAcrsAcrDpiSignatureIds) Validate() error {
	return dara.Validate(s)
}
