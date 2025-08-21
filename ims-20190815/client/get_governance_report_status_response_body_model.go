// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGovernanceReportStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGovernanceItemsStatus(v *GetGovernanceReportStatusResponseBodyGovernanceItemsStatus) *GetGovernanceReportStatusResponseBody
	GetGovernanceItemsStatus() *GetGovernanceReportStatusResponseBodyGovernanceItemsStatus
	SetRequestId(v string) *GetGovernanceReportStatusResponseBody
	GetRequestId() *string
	SetWholeReportStatus(v string) *GetGovernanceReportStatusResponseBody
	GetWholeReportStatus() *string
}

type GetGovernanceReportStatusResponseBody struct {
	GovernanceItemsStatus *GetGovernanceReportStatusResponseBodyGovernanceItemsStatus `json:"GovernanceItemsStatus,omitempty" xml:"GovernanceItemsStatus,omitempty" type:"Struct"`
	// example:
	//
	// F2CE9688-AA85-5F23-8C22-0EC23473405A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Progressing
	WholeReportStatus *string `json:"WholeReportStatus,omitempty" xml:"WholeReportStatus,omitempty"`
}

func (s GetGovernanceReportStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGovernanceReportStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetGovernanceReportStatusResponseBody) GetGovernanceItemsStatus() *GetGovernanceReportStatusResponseBodyGovernanceItemsStatus {
	return s.GovernanceItemsStatus
}

func (s *GetGovernanceReportStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGovernanceReportStatusResponseBody) GetWholeReportStatus() *string {
	return s.WholeReportStatus
}

func (s *GetGovernanceReportStatusResponseBody) SetGovernanceItemsStatus(v *GetGovernanceReportStatusResponseBodyGovernanceItemsStatus) *GetGovernanceReportStatusResponseBody {
	s.GovernanceItemsStatus = v
	return s
}

func (s *GetGovernanceReportStatusResponseBody) SetRequestId(v string) *GetGovernanceReportStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGovernanceReportStatusResponseBody) SetWholeReportStatus(v string) *GetGovernanceReportStatusResponseBody {
	s.WholeReportStatus = &v
	return s
}

func (s *GetGovernanceReportStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetGovernanceReportStatusResponseBodyGovernanceItemsStatus struct {
	GovernanceItemStatus []*GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus `json:"GovernanceItemStatus,omitempty" xml:"GovernanceItemStatus,omitempty" type:"Repeated"`
}

func (s GetGovernanceReportStatusResponseBodyGovernanceItemsStatus) String() string {
	return dara.Prettify(s)
}

func (s GetGovernanceReportStatusResponseBodyGovernanceItemsStatus) GoString() string {
	return s.String()
}

func (s *GetGovernanceReportStatusResponseBodyGovernanceItemsStatus) GetGovernanceItemStatus() []*GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus {
	return s.GovernanceItemStatus
}

func (s *GetGovernanceReportStatusResponseBodyGovernanceItemsStatus) SetGovernanceItemStatus(v []*GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus) *GetGovernanceReportStatusResponseBodyGovernanceItemsStatus {
	s.GovernanceItemStatus = v
	return s
}

func (s *GetGovernanceReportStatusResponseBodyGovernanceItemsStatus) Validate() error {
	return dara.Validate(s)
}

type GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus struct {
	// example:
	//
	// AccountRecentUsingAccessKey
	GovernanceItem *string `json:"GovernanceItem,omitempty" xml:"GovernanceItem,omitempty"`
	// example:
	//
	// Progressing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus) String() string {
	return dara.Prettify(s)
}

func (s GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus) GoString() string {
	return s.String()
}

func (s *GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus) GetGovernanceItem() *string {
	return s.GovernanceItem
}

func (s *GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus) GetStatus() *string {
	return s.Status
}

func (s *GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus) SetGovernanceItem(v string) *GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus {
	s.GovernanceItem = &v
	return s
}

func (s *GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus) SetStatus(v string) *GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus {
	s.Status = &v
	return s
}

func (s *GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus) Validate() error {
	return dara.Validate(s)
}
