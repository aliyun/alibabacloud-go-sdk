// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRedemptionRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExternalUserId(v string) *QueryRedemptionRecordsRequest
	GetExternalUserId() *string
	SetPage(v int32) *QueryRedemptionRecordsRequest
	GetPage() *int32
	SetPageSize(v int32) *QueryRedemptionRecordsRequest
	GetPageSize() *int32
	SetRedemptionOrderNo(v string) *QueryRedemptionRecordsRequest
	GetRedemptionOrderNo() *string
}

type QueryRedemptionRecordsRequest struct {
	// example:
	//
	// 1001
	ExternalUserId *string `json:"externalUserId,omitempty" xml:"externalUserId,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// ORD20240101000001
	RedemptionOrderNo *string `json:"redemptionOrderNo,omitempty" xml:"redemptionOrderNo,omitempty"`
}

func (s QueryRedemptionRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRedemptionRecordsRequest) GoString() string {
	return s.String()
}

func (s *QueryRedemptionRecordsRequest) GetExternalUserId() *string {
	return s.ExternalUserId
}

func (s *QueryRedemptionRecordsRequest) GetPage() *int32 {
	return s.Page
}

func (s *QueryRedemptionRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRedemptionRecordsRequest) GetRedemptionOrderNo() *string {
	return s.RedemptionOrderNo
}

func (s *QueryRedemptionRecordsRequest) SetExternalUserId(v string) *QueryRedemptionRecordsRequest {
	s.ExternalUserId = &v
	return s
}

func (s *QueryRedemptionRecordsRequest) SetPage(v int32) *QueryRedemptionRecordsRequest {
	s.Page = &v
	return s
}

func (s *QueryRedemptionRecordsRequest) SetPageSize(v int32) *QueryRedemptionRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRedemptionRecordsRequest) SetRedemptionOrderNo(v string) *QueryRedemptionRecordsRequest {
	s.RedemptionOrderNo = &v
	return s
}

func (s *QueryRedemptionRecordsRequest) Validate() error {
	return dara.Validate(s)
}
