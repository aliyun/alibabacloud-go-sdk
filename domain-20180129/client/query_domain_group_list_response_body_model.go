// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainGroupListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryDomainGroupListResponseBodyData) *QueryDomainGroupListResponseBody
	GetData() *QueryDomainGroupListResponseBodyData
	SetRequestId(v string) *QueryDomainGroupListResponseBody
	GetRequestId() *string
}

type QueryDomainGroupListResponseBody struct {
	Data *QueryDomainGroupListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 80011ABC-F573-4795-B0E8-377BFBBA3422
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDomainGroupListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainGroupListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDomainGroupListResponseBody) GetData() *QueryDomainGroupListResponseBodyData {
	return s.Data
}

func (s *QueryDomainGroupListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDomainGroupListResponseBody) SetData(v *QueryDomainGroupListResponseBodyData) *QueryDomainGroupListResponseBody {
	s.Data = v
	return s
}

func (s *QueryDomainGroupListResponseBody) SetRequestId(v string) *QueryDomainGroupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDomainGroupListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDomainGroupListResponseBodyData struct {
	DomainGroup []*QueryDomainGroupListResponseBodyDataDomainGroup `json:"DomainGroup,omitempty" xml:"DomainGroup,omitempty" type:"Repeated"`
}

func (s QueryDomainGroupListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainGroupListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDomainGroupListResponseBodyData) GetDomainGroup() []*QueryDomainGroupListResponseBodyDataDomainGroup {
	return s.DomainGroup
}

func (s *QueryDomainGroupListResponseBodyData) SetDomainGroup(v []*QueryDomainGroupListResponseBodyDataDomainGroup) *QueryDomainGroupListResponseBodyData {
	s.DomainGroup = v
	return s
}

func (s *QueryDomainGroupListResponseBodyData) Validate() error {
	if s.DomainGroup != nil {
		for _, item := range s.DomainGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDomainGroupListResponseBodyDataDomainGroup struct {
	// example:
	//
	// false
	BeingDeleted *bool `json:"BeingDeleted,omitempty" xml:"BeingDeleted,omitempty"`
	// example:
	//
	// 2018-04-02 15:59:06
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// example:
	//
	// -1
	DomainGroupId   *string `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	DomainGroupName *string `json:"DomainGroupName,omitempty" xml:"DomainGroupName,omitempty"`
	// example:
	//
	// COMPLETE
	DomainGroupStatus *string `json:"DomainGroupStatus,omitempty" xml:"DomainGroupStatus,omitempty"`
	// example:
	//
	// 2018-04-02 15:59:06
	ModificationDate *string `json:"ModificationDate,omitempty" xml:"ModificationDate,omitempty"`
	// example:
	//
	// 20
	TotalNumber *int32 `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s QueryDomainGroupListResponseBodyDataDomainGroup) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainGroupListResponseBodyDataDomainGroup) GoString() string {
	return s.String()
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) GetBeingDeleted() *bool {
	return s.BeingDeleted
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) GetCreationDate() *string {
	return s.CreationDate
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) GetDomainGroupId() *string {
	return s.DomainGroupId
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) GetDomainGroupName() *string {
	return s.DomainGroupName
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) GetDomainGroupStatus() *string {
	return s.DomainGroupStatus
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) GetModificationDate() *string {
	return s.ModificationDate
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) GetTotalNumber() *int32 {
	return s.TotalNumber
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) SetBeingDeleted(v bool) *QueryDomainGroupListResponseBodyDataDomainGroup {
	s.BeingDeleted = &v
	return s
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) SetCreationDate(v string) *QueryDomainGroupListResponseBodyDataDomainGroup {
	s.CreationDate = &v
	return s
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) SetDomainGroupId(v string) *QueryDomainGroupListResponseBodyDataDomainGroup {
	s.DomainGroupId = &v
	return s
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) SetDomainGroupName(v string) *QueryDomainGroupListResponseBodyDataDomainGroup {
	s.DomainGroupName = &v
	return s
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) SetDomainGroupStatus(v string) *QueryDomainGroupListResponseBodyDataDomainGroup {
	s.DomainGroupStatus = &v
	return s
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) SetModificationDate(v string) *QueryDomainGroupListResponseBodyDataDomainGroup {
	s.ModificationDate = &v
	return s
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) SetTotalNumber(v int32) *QueryDomainGroupListResponseBodyDataDomainGroup {
	s.TotalNumber = &v
	return s
}

func (s *QueryDomainGroupListResponseBodyDataDomainGroup) Validate() error {
	return dara.Validate(s)
}
