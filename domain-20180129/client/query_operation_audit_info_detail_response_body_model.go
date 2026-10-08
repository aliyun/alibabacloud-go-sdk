// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOperationAuditInfoDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuditInfo(v string) *QueryOperationAuditInfoDetailResponseBody
	GetAuditInfo() *string
	SetAuditStatus(v int32) *QueryOperationAuditInfoDetailResponseBody
	GetAuditStatus() *int32
	SetAuditType(v int32) *QueryOperationAuditInfoDetailResponseBody
	GetAuditType() *int32
	SetBusinessName(v string) *QueryOperationAuditInfoDetailResponseBody
	GetBusinessName() *string
	SetCreateTime(v int64) *QueryOperationAuditInfoDetailResponseBody
	GetCreateTime() *int64
	SetDomainName(v string) *QueryOperationAuditInfoDetailResponseBody
	GetDomainName() *string
	SetId(v string) *QueryOperationAuditInfoDetailResponseBody
	GetId() *string
	SetRemark(v string) *QueryOperationAuditInfoDetailResponseBody
	GetRemark() *string
	SetRequestId(v string) *QueryOperationAuditInfoDetailResponseBody
	GetRequestId() *string
	SetUpdateTime(v int64) *QueryOperationAuditInfoDetailResponseBody
	GetUpdateTime() *int64
}

type QueryOperationAuditInfoDetailResponseBody struct {
	// Review information.
	//
	// example:
	//
	// {"regType":1,"registrantName":"张三","telephone":"1390123****","account":"username@example.com","reason":1,"remark":"账号丢失"}
	AuditInfo *string `json:"AuditInfo,omitempty" xml:"AuditInfo,omitempty"`
	// Review Status. Valid values:
	//
	// - **0**: Pending supplementary information.
	//
	// - **1**, **2**, **3**, **4**: Under review.
	//
	// - **5**: Review failed.
	//
	// - **6**: Review succeeded.
	//
	// - **7**: Review canceled.
	//
	// example:
	//
	// 1
	AuditStatus *int32 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// Review Type. Valid value:
	//
	// **1**: Offline domain name transfer.
	//
	// example:
	//
	// 1
	AuditType *int32 `json:"AuditType,omitempty" xml:"AuditType,omitempty"`
	// Name of the reviewed business.
	//
	// example:
	//
	// example.com等域名线下转移
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// Record creation time.
	//
	// example:
	//
	// 1581919010100
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Domain name.
	//
	// example:
	//
	// example.com,aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Review record ID.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Review remark.
	//
	// example:
	//
	// 审核通过
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 9DFCF6F8-243C-40EC-8035-4B12FEFD7D1L
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Record update time.
	//
	// example:
	//
	// 1581919010101
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s QueryOperationAuditInfoDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryOperationAuditInfoDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOperationAuditInfoDetailResponseBody) GetAuditInfo() *string {
	return s.AuditInfo
}

func (s *QueryOperationAuditInfoDetailResponseBody) GetAuditStatus() *int32 {
	return s.AuditStatus
}

func (s *QueryOperationAuditInfoDetailResponseBody) GetAuditType() *int32 {
	return s.AuditType
}

func (s *QueryOperationAuditInfoDetailResponseBody) GetBusinessName() *string {
	return s.BusinessName
}

func (s *QueryOperationAuditInfoDetailResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *QueryOperationAuditInfoDetailResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryOperationAuditInfoDetailResponseBody) GetId() *string {
	return s.Id
}

func (s *QueryOperationAuditInfoDetailResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *QueryOperationAuditInfoDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryOperationAuditInfoDetailResponseBody) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *QueryOperationAuditInfoDetailResponseBody) SetAuditInfo(v string) *QueryOperationAuditInfoDetailResponseBody {
	s.AuditInfo = &v
	return s
}

func (s *QueryOperationAuditInfoDetailResponseBody) SetAuditStatus(v int32) *QueryOperationAuditInfoDetailResponseBody {
	s.AuditStatus = &v
	return s
}

func (s *QueryOperationAuditInfoDetailResponseBody) SetAuditType(v int32) *QueryOperationAuditInfoDetailResponseBody {
	s.AuditType = &v
	return s
}

func (s *QueryOperationAuditInfoDetailResponseBody) SetBusinessName(v string) *QueryOperationAuditInfoDetailResponseBody {
	s.BusinessName = &v
	return s
}

func (s *QueryOperationAuditInfoDetailResponseBody) SetCreateTime(v int64) *QueryOperationAuditInfoDetailResponseBody {
	s.CreateTime = &v
	return s
}

func (s *QueryOperationAuditInfoDetailResponseBody) SetDomainName(v string) *QueryOperationAuditInfoDetailResponseBody {
	s.DomainName = &v
	return s
}

func (s *QueryOperationAuditInfoDetailResponseBody) SetId(v string) *QueryOperationAuditInfoDetailResponseBody {
	s.Id = &v
	return s
}

func (s *QueryOperationAuditInfoDetailResponseBody) SetRemark(v string) *QueryOperationAuditInfoDetailResponseBody {
	s.Remark = &v
	return s
}

func (s *QueryOperationAuditInfoDetailResponseBody) SetRequestId(v string) *QueryOperationAuditInfoDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOperationAuditInfoDetailResponseBody) SetUpdateTime(v int64) *QueryOperationAuditInfoDetailResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *QueryOperationAuditInfoDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
