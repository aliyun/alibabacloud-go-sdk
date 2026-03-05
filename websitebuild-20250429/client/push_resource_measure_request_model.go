// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushResourceMeasureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int64) *PushResourceMeasureRequest
	GetAmount() *int64
	SetBelongId(v string) *PushResourceMeasureRequest
	GetBelongId() *string
	SetBelongIdType(v string) *PushResourceMeasureRequest
	GetBelongIdType() *string
	SetBizId(v string) *PushResourceMeasureRequest
	GetBizId() *string
	SetMeasureData(v string) *PushResourceMeasureRequest
	GetMeasureData() *string
	SetMetaData(v string) *PushResourceMeasureRequest
	GetMetaData() *string
	SetResourceCode(v string) *PushResourceMeasureRequest
	GetResourceCode() *string
	SetUseTime(v string) *PushResourceMeasureRequest
	GetUseTime() *string
	SetUseType(v string) *PushResourceMeasureRequest
	GetUseType() *string
}

type PushResourceMeasureRequest struct {
	// example:
	//
	// 100
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// 123456
	BelongId *string `json:"BelongId,omitempty" xml:"BelongId,omitempty"`
	// example:
	//
	// USER
	BelongIdType *string `json:"BelongIdType,omitempty" xml:"BelongIdType,omitempty"`
	// example:
	//
	// WS20250801154628000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// {\\"MD5\\":\\"8ba46100bd898461f0f589704f2fad25\\",\\"driver\\":\\"vhd\\",\\"flag\\":\\"769\\",\\"imds_support\\":\\"v1\\",\\"io_optimized\\":true,\\"nvme_supported\\":false,\\"uefi_preferred\\":false}
	MeasureData *string `json:"MeasureData,omitempty" xml:"MeasureData,omitempty"`
	// example:
	//
	// {\\"MD5\\":\\"8ba46100bd898461f0f589704f2fad25\\",\\"driver\\":\\"vhd\\",\\"flag\\":\\"769\\",\\"imds_support\\":\\"v1\\",\\"io_optimized\\":true,\\"nvme_supported\\":false,\\"uefi_preferred\\":false}
	MetaData *string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	// example:
	//
	// InspirationTokens
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	// example:
	//
	// yyyy-MM-dd HH:mm:ss
	UseTime *string `json:"UseTime,omitempty" xml:"UseTime,omitempty"`
	// example:
	//
	// MANUAL_BIZ
	UseType *string `json:"UseType,omitempty" xml:"UseType,omitempty"`
}

func (s PushResourceMeasureRequest) String() string {
	return dara.Prettify(s)
}

func (s PushResourceMeasureRequest) GoString() string {
	return s.String()
}

func (s *PushResourceMeasureRequest) GetAmount() *int64 {
	return s.Amount
}

func (s *PushResourceMeasureRequest) GetBelongId() *string {
	return s.BelongId
}

func (s *PushResourceMeasureRequest) GetBelongIdType() *string {
	return s.BelongIdType
}

func (s *PushResourceMeasureRequest) GetBizId() *string {
	return s.BizId
}

func (s *PushResourceMeasureRequest) GetMeasureData() *string {
	return s.MeasureData
}

func (s *PushResourceMeasureRequest) GetMetaData() *string {
	return s.MetaData
}

func (s *PushResourceMeasureRequest) GetResourceCode() *string {
	return s.ResourceCode
}

func (s *PushResourceMeasureRequest) GetUseTime() *string {
	return s.UseTime
}

func (s *PushResourceMeasureRequest) GetUseType() *string {
	return s.UseType
}

func (s *PushResourceMeasureRequest) SetAmount(v int64) *PushResourceMeasureRequest {
	s.Amount = &v
	return s
}

func (s *PushResourceMeasureRequest) SetBelongId(v string) *PushResourceMeasureRequest {
	s.BelongId = &v
	return s
}

func (s *PushResourceMeasureRequest) SetBelongIdType(v string) *PushResourceMeasureRequest {
	s.BelongIdType = &v
	return s
}

func (s *PushResourceMeasureRequest) SetBizId(v string) *PushResourceMeasureRequest {
	s.BizId = &v
	return s
}

func (s *PushResourceMeasureRequest) SetMeasureData(v string) *PushResourceMeasureRequest {
	s.MeasureData = &v
	return s
}

func (s *PushResourceMeasureRequest) SetMetaData(v string) *PushResourceMeasureRequest {
	s.MetaData = &v
	return s
}

func (s *PushResourceMeasureRequest) SetResourceCode(v string) *PushResourceMeasureRequest {
	s.ResourceCode = &v
	return s
}

func (s *PushResourceMeasureRequest) SetUseTime(v string) *PushResourceMeasureRequest {
	s.UseTime = &v
	return s
}

func (s *PushResourceMeasureRequest) SetUseType(v string) *PushResourceMeasureRequest {
	s.UseType = &v
	return s
}

func (s *PushResourceMeasureRequest) Validate() error {
	return dara.Validate(s)
}
