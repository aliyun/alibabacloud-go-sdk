// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyCoordinationWithCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ApplyCoordinationWithCodeResponseBodyData) *ApplyCoordinationWithCodeResponseBody
	GetData() *ApplyCoordinationWithCodeResponseBodyData
	SetRequestId(v string) *ApplyCoordinationWithCodeResponseBody
	GetRequestId() *string
}

type ApplyCoordinationWithCodeResponseBody struct {
	Data *ApplyCoordinationWithCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// AD2D0761-1FE5-549D-B169******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyCoordinationWithCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyCoordinationWithCodeResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyCoordinationWithCodeResponseBody) GetData() *ApplyCoordinationWithCodeResponseBodyData {
	return s.Data
}

func (s *ApplyCoordinationWithCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyCoordinationWithCodeResponseBody) SetData(v *ApplyCoordinationWithCodeResponseBodyData) *ApplyCoordinationWithCodeResponseBody {
	s.Data = v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBody) SetRequestId(v string) *ApplyCoordinationWithCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ApplyCoordinationWithCodeResponseBodyData struct {
	// example:
	//
	// co-0ad0f3p4n2******
	CoId *string `json:"CoId,omitempty" xml:"CoId,omitempty"`
	// example:
	//
	// COORDINATING
	CoordinateStatus *string `json:"CoordinateStatus,omitempty" xml:"CoordinateStatus,omitempty"`
	// example:
	//
	// DQpbRGVza3RvcF0NCkZvcmNlVGxzVHlwZT0xDQ******
	CoordinateTicket *string `json:"CoordinateTicket,omitempty" xml:"CoordinateTicket,omitempty"`
	// example:
	//
	// 1126819517******
	CoordinatorAliUid *int64 `json:"CoordinatorAliUid,omitempty" xml:"CoordinatorAliUid,omitempty"`
	// example:
	//
	// bob
	CoordinatorUserId *string `json:"CoordinatorUserId,omitempty" xml:"CoordinatorUserId,omitempty"`
	// example:
	//
	// 1126819517******
	OwnerAliUid *int64 `json:"OwnerAliUid,omitempty" xml:"OwnerAliUid,omitempty"`
	// example:
	//
	// alice
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// example:
	//
	// ecd-3vv4mf8zxg******
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// demo-desktop
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// example:
	//
	// CloudDesktop
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ApplyCoordinationWithCodeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ApplyCoordinationWithCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ApplyCoordinationWithCodeResponseBodyData) GetCoId() *string {
	return s.CoId
}

func (s *ApplyCoordinationWithCodeResponseBodyData) GetCoordinateStatus() *string {
	return s.CoordinateStatus
}

func (s *ApplyCoordinationWithCodeResponseBodyData) GetCoordinateTicket() *string {
	return s.CoordinateTicket
}

func (s *ApplyCoordinationWithCodeResponseBodyData) GetCoordinatorAliUid() *int64 {
	return s.CoordinatorAliUid
}

func (s *ApplyCoordinationWithCodeResponseBodyData) GetCoordinatorUserId() *string {
	return s.CoordinatorUserId
}

func (s *ApplyCoordinationWithCodeResponseBodyData) GetOwnerAliUid() *int64 {
	return s.OwnerAliUid
}

func (s *ApplyCoordinationWithCodeResponseBodyData) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *ApplyCoordinationWithCodeResponseBodyData) GetResourceId() *string {
	return s.ResourceId
}

func (s *ApplyCoordinationWithCodeResponseBodyData) GetResourceName() *string {
	return s.ResourceName
}

func (s *ApplyCoordinationWithCodeResponseBodyData) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *ApplyCoordinationWithCodeResponseBodyData) GetResourceType() *string {
	return s.ResourceType
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetCoId(v string) *ApplyCoordinationWithCodeResponseBodyData {
	s.CoId = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetCoordinateStatus(v string) *ApplyCoordinationWithCodeResponseBodyData {
	s.CoordinateStatus = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetCoordinateTicket(v string) *ApplyCoordinationWithCodeResponseBodyData {
	s.CoordinateTicket = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetCoordinatorAliUid(v int64) *ApplyCoordinationWithCodeResponseBodyData {
	s.CoordinatorAliUid = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetCoordinatorUserId(v string) *ApplyCoordinationWithCodeResponseBodyData {
	s.CoordinatorUserId = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetOwnerAliUid(v int64) *ApplyCoordinationWithCodeResponseBodyData {
	s.OwnerAliUid = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetOwnerUserId(v string) *ApplyCoordinationWithCodeResponseBodyData {
	s.OwnerUserId = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetResourceId(v string) *ApplyCoordinationWithCodeResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetResourceName(v string) *ApplyCoordinationWithCodeResponseBodyData {
	s.ResourceName = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetResourceRegionId(v string) *ApplyCoordinationWithCodeResponseBodyData {
	s.ResourceRegionId = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBodyData) SetResourceType(v string) *ApplyCoordinationWithCodeResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *ApplyCoordinationWithCodeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
