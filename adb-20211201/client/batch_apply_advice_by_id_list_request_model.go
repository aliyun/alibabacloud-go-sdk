// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchApplyAdviceByIdListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdviceDate(v int64) *BatchApplyAdviceByIdListRequest
	GetAdviceDate() *int64
	SetAdviceIdList(v string) *BatchApplyAdviceByIdListRequest
	GetAdviceIdList() *string
	SetApplyType(v string) *BatchApplyAdviceByIdListRequest
	GetApplyType() *string
	SetBuildImmediately(v bool) *BatchApplyAdviceByIdListRequest
	GetBuildImmediately() *bool
	SetDBClusterId(v string) *BatchApplyAdviceByIdListRequest
	GetDBClusterId() *string
	SetRegionId(v string) *BatchApplyAdviceByIdListRequest
	GetRegionId() *string
}

type BatchApplyAdviceByIdListRequest struct {
	// The date when the suggestions were generated. Format: yyyyMMdd.
	//
	// example:
	//
	// 20221115
	AdviceDate *int64 `json:"AdviceDate,omitempty" xml:"AdviceDate,omitempty"`
	// The list of suggestion IDs to apply in batches. Separate multiple suggestion IDs with commas (,).
	//
	// example:
	//
	// c2589ff3-e86c-4f19-80c8-2aeb7dd9****,53414470-ebf4-4a53-a312-8a1ad8fd****,6e8dce84-fec8-4b0b-9c04-b0cea12c****,b3b9703d-55ca-47e0-96dd-6a4a9dbf****
	AdviceIdList *string `json:"AdviceIdList,omitempty" xml:"AdviceIdList,omitempty"`
	// The adoption type.
	//
	// example:
	//
	// DROP_INDEX
	ApplyType *string `json:"ApplyType,omitempty" xml:"ApplyType,omitempty"`
	// Specifies whether to immediately start a build task.
	//
	// example:
	//
	// true
	BuildImmediately *bool `json:"BuildImmediately,omitempty" xml:"BuildImmediately,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-uf6g8w25jacm7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BatchApplyAdviceByIdListRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchApplyAdviceByIdListRequest) GoString() string {
	return s.String()
}

func (s *BatchApplyAdviceByIdListRequest) GetAdviceDate() *int64 {
	return s.AdviceDate
}

func (s *BatchApplyAdviceByIdListRequest) GetAdviceIdList() *string {
	return s.AdviceIdList
}

func (s *BatchApplyAdviceByIdListRequest) GetApplyType() *string {
	return s.ApplyType
}

func (s *BatchApplyAdviceByIdListRequest) GetBuildImmediately() *bool {
	return s.BuildImmediately
}

func (s *BatchApplyAdviceByIdListRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *BatchApplyAdviceByIdListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BatchApplyAdviceByIdListRequest) SetAdviceDate(v int64) *BatchApplyAdviceByIdListRequest {
	s.AdviceDate = &v
	return s
}

func (s *BatchApplyAdviceByIdListRequest) SetAdviceIdList(v string) *BatchApplyAdviceByIdListRequest {
	s.AdviceIdList = &v
	return s
}

func (s *BatchApplyAdviceByIdListRequest) SetApplyType(v string) *BatchApplyAdviceByIdListRequest {
	s.ApplyType = &v
	return s
}

func (s *BatchApplyAdviceByIdListRequest) SetBuildImmediately(v bool) *BatchApplyAdviceByIdListRequest {
	s.BuildImmediately = &v
	return s
}

func (s *BatchApplyAdviceByIdListRequest) SetDBClusterId(v string) *BatchApplyAdviceByIdListRequest {
	s.DBClusterId = &v
	return s
}

func (s *BatchApplyAdviceByIdListRequest) SetRegionId(v string) *BatchApplyAdviceByIdListRequest {
	s.RegionId = &v
	return s
}

func (s *BatchApplyAdviceByIdListRequest) Validate() error {
	return dara.Validate(s)
}
