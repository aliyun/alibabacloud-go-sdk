// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpEntitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOpEntities(v []*DescribeOpEntitiesResponseBodyOpEntities) *DescribeOpEntitiesResponseBody
	GetOpEntities() []*DescribeOpEntitiesResponseBodyOpEntities
	SetRequestId(v string) *DescribeOpEntitiesResponseBody
	GetRequestId() *string
	SetTotal(v int64) *DescribeOpEntitiesResponseBody
	GetTotal() *int64
}

type DescribeOpEntitiesResponseBody struct {
	OpEntities []*DescribeOpEntitiesResponseBodyOpEntities `json:"OpEntities,omitempty" xml:"OpEntities,omitempty" type:"Repeated"`
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeOpEntitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseBody) GetOpEntities() []*DescribeOpEntitiesResponseBodyOpEntities {
	return s.OpEntities
}

func (s *DescribeOpEntitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOpEntitiesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeOpEntitiesResponseBody) SetOpEntities(v []*DescribeOpEntitiesResponseBodyOpEntities) *DescribeOpEntitiesResponseBody {
	s.OpEntities = v
	return s
}

func (s *DescribeOpEntitiesResponseBody) SetRequestId(v string) *DescribeOpEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpEntitiesResponseBody) SetTotal(v int64) *DescribeOpEntitiesResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeOpEntitiesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeOpEntitiesResponseBodyOpEntities struct {
	// example:
	//
	// 2.2.2.2
	EntityObject *string `json:"EntityObject,omitempty" xml:"EntityObject,omitempty"`
	// example:
	//
	// 1
	EntityType *int32 `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// 1536715558000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 123
	OpAccount *string `json:"OpAccount,omitempty" xml:"OpAccount,omitempty"`
	// example:
	//
	// 1
	OpAction *int32 `json:"OpAction,omitempty" xml:"OpAction,omitempty"`
	// example:
	//
	// {"newEntity":{"elasticBandwidth":30},"oldEntity":{"elasticBandwidth":200}}
	OpDesc *string `json:"OpDesc,omitempty" xml:"OpDesc,omitempty"`
}

func (s DescribeOpEntitiesResponseBodyOpEntities) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpEntitiesResponseBodyOpEntities) GoString() string {
	return s.String()
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) GetEntityObject() *string {
	return s.EntityObject
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) GetEntityType() *int32 {
	return s.EntityType
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) GetOpAccount() *string {
	return s.OpAccount
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) GetOpAction() *int32 {
	return s.OpAction
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) GetOpDesc() *string {
	return s.OpDesc
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetEntityObject(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.EntityObject = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetEntityType(v int32) *DescribeOpEntitiesResponseBodyOpEntities {
	s.EntityType = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetGmtCreate(v int64) *DescribeOpEntitiesResponseBodyOpEntities {
	s.GmtCreate = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpAccount(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpAccount = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpAction(v int32) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpAction = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) SetOpDesc(v string) *DescribeOpEntitiesResponseBodyOpEntities {
	s.OpDesc = &v
	return s
}

func (s *DescribeOpEntitiesResponseBodyOpEntities) Validate() error {
	return dara.Validate(s)
}
