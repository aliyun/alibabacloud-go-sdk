// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveDomainGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeingDeleted(v bool) *SaveDomainGroupResponseBody
	GetBeingDeleted() *bool
	SetCreationDate(v string) *SaveDomainGroupResponseBody
	GetCreationDate() *string
	SetDomainGroupId(v int64) *SaveDomainGroupResponseBody
	GetDomainGroupId() *int64
	SetDomainGroupName(v string) *SaveDomainGroupResponseBody
	GetDomainGroupName() *string
	SetDomainGroupStatus(v string) *SaveDomainGroupResponseBody
	GetDomainGroupStatus() *string
	SetModificationDate(v string) *SaveDomainGroupResponseBody
	GetModificationDate() *string
	SetRequestId(v string) *SaveDomainGroupResponseBody
	GetRequestId() *string
	SetTotalNumber(v int32) *SaveDomainGroupResponseBody
	GetTotalNumber() *int32
}

type SaveDomainGroupResponseBody struct {
	// Indicates whether the group is being deleted.
	//
	// > For groups containing more than 1,000 domain names, deletion is an asynchronous procedure that requires some time for the system to process. During this period, this field is **true**.
	//
	// example:
	//
	// false
	BeingDeleted *bool `json:"BeingDeleted,omitempty" xml:"BeingDeleted,omitempty"`
	// Creation Time of the domain name group.
	//
	// example:
	//
	// 2018-04-02 15:59:06
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// Domain group ID.
	//
	// example:
	//
	// 123456
	DomainGroupId *int64 `json:"DomainGroupId,omitempty" xml:"DomainGroupId,omitempty"`
	// Domain Name Group Name.
	//
	// example:
	//
	// 测试分组
	DomainGroupName *string `json:"DomainGroupName,omitempty" xml:"DomainGroupName,omitempty"`
	// Status of the domain name group. Valid values:
	//
	// - **PROCESSING**: Processing;
	//
	// - **COMPLETE**: Complete.
	//
	// > In cases such as setting a group via a file or replacing a group with more than 1,000 domain names, the operation is asynchronous and requires waiting for system processing. During this time, this field is **PROCESSING**.
	//
	// example:
	//
	// COMPLETE
	DomainGroupStatus *string `json:"DomainGroupStatus,omitempty" xml:"DomainGroupStatus,omitempty"`
	// Updated At time of the domain name group.
	//
	// example:
	//
	// 2018-04-02 15:59:06
	ModificationDate *string `json:"ModificationDate,omitempty" xml:"ModificationDate,omitempty"`
	// Unique request identity.
	//
	// example:
	//
	// 80011ABC-F573-4795-B0E8-377BFBBA3422
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Quantity of domain names.
	//
	// example:
	//
	// 20
	TotalNumber *int32 `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s SaveDomainGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveDomainGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SaveDomainGroupResponseBody) GetBeingDeleted() *bool {
	return s.BeingDeleted
}

func (s *SaveDomainGroupResponseBody) GetCreationDate() *string {
	return s.CreationDate
}

func (s *SaveDomainGroupResponseBody) GetDomainGroupId() *int64 {
	return s.DomainGroupId
}

func (s *SaveDomainGroupResponseBody) GetDomainGroupName() *string {
	return s.DomainGroupName
}

func (s *SaveDomainGroupResponseBody) GetDomainGroupStatus() *string {
	return s.DomainGroupStatus
}

func (s *SaveDomainGroupResponseBody) GetModificationDate() *string {
	return s.ModificationDate
}

func (s *SaveDomainGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveDomainGroupResponseBody) GetTotalNumber() *int32 {
	return s.TotalNumber
}

func (s *SaveDomainGroupResponseBody) SetBeingDeleted(v bool) *SaveDomainGroupResponseBody {
	s.BeingDeleted = &v
	return s
}

func (s *SaveDomainGroupResponseBody) SetCreationDate(v string) *SaveDomainGroupResponseBody {
	s.CreationDate = &v
	return s
}

func (s *SaveDomainGroupResponseBody) SetDomainGroupId(v int64) *SaveDomainGroupResponseBody {
	s.DomainGroupId = &v
	return s
}

func (s *SaveDomainGroupResponseBody) SetDomainGroupName(v string) *SaveDomainGroupResponseBody {
	s.DomainGroupName = &v
	return s
}

func (s *SaveDomainGroupResponseBody) SetDomainGroupStatus(v string) *SaveDomainGroupResponseBody {
	s.DomainGroupStatus = &v
	return s
}

func (s *SaveDomainGroupResponseBody) SetModificationDate(v string) *SaveDomainGroupResponseBody {
	s.ModificationDate = &v
	return s
}

func (s *SaveDomainGroupResponseBody) SetRequestId(v string) *SaveDomainGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveDomainGroupResponseBody) SetTotalNumber(v int32) *SaveDomainGroupResponseBody {
	s.TotalNumber = &v
	return s
}

func (s *SaveDomainGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
