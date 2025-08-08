// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInvoiceForIsvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckNotice(v string) *ModifyInvoiceForIsvRequest
	GetCheckNotice() *string
	SetElectronUrl(v string) *ModifyInvoiceForIsvRequest
	GetElectronUrl() *string
	SetInvoiceId(v int64) *ModifyInvoiceForIsvRequest
	GetInvoiceId() *int64
	SetNumber(v string) *ModifyInvoiceForIsvRequest
	GetNumber() *string
	SetOperateType(v int32) *ModifyInvoiceForIsvRequest
	GetOperateType() *int32
	SetType(v int32) *ModifyInvoiceForIsvRequest
	GetType() *int32
}

type ModifyInvoiceForIsvRequest struct {
	CheckNotice *string `json:"CheckNotice,omitempty" xml:"CheckNotice,omitempty"`
	// example:
	//
	// https://oss.aliyuncs.com/xxxx.png
	ElectronUrl *string `json:"ElectronUrl,omitempty" xml:"ElectronUrl,omitempty"`
	InvoiceId   *int64  `json:"InvoiceId,omitempty" xml:"InvoiceId,omitempty"`
	// example:
	//
	// 1897702****
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	OperateType *int32 `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// example:
	//
	// 2
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyInvoiceForIsvRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInvoiceForIsvRequest) GoString() string {
	return s.String()
}

func (s *ModifyInvoiceForIsvRequest) GetCheckNotice() *string {
	return s.CheckNotice
}

func (s *ModifyInvoiceForIsvRequest) GetElectronUrl() *string {
	return s.ElectronUrl
}

func (s *ModifyInvoiceForIsvRequest) GetInvoiceId() *int64 {
	return s.InvoiceId
}

func (s *ModifyInvoiceForIsvRequest) GetNumber() *string {
	return s.Number
}

func (s *ModifyInvoiceForIsvRequest) GetOperateType() *int32 {
	return s.OperateType
}

func (s *ModifyInvoiceForIsvRequest) GetType() *int32 {
	return s.Type
}

func (s *ModifyInvoiceForIsvRequest) SetCheckNotice(v string) *ModifyInvoiceForIsvRequest {
	s.CheckNotice = &v
	return s
}

func (s *ModifyInvoiceForIsvRequest) SetElectronUrl(v string) *ModifyInvoiceForIsvRequest {
	s.ElectronUrl = &v
	return s
}

func (s *ModifyInvoiceForIsvRequest) SetInvoiceId(v int64) *ModifyInvoiceForIsvRequest {
	s.InvoiceId = &v
	return s
}

func (s *ModifyInvoiceForIsvRequest) SetNumber(v string) *ModifyInvoiceForIsvRequest {
	s.Number = &v
	return s
}

func (s *ModifyInvoiceForIsvRequest) SetOperateType(v int32) *ModifyInvoiceForIsvRequest {
	s.OperateType = &v
	return s
}

func (s *ModifyInvoiceForIsvRequest) SetType(v int32) *ModifyInvoiceForIsvRequest {
	s.Type = &v
	return s
}

func (s *ModifyInvoiceForIsvRequest) Validate() error {
	return dara.Validate(s)
}
