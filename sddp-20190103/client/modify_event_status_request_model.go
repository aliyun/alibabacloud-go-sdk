// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEventStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBacked(v bool) *ModifyEventStatusRequest
	GetBacked() *bool
	SetDealReason(v string) *ModifyEventStatusRequest
	GetDealReason() *string
	SetId(v int64) *ModifyEventStatusRequest
	GetId() *int64
	SetLang(v string) *ModifyEventStatusRequest
	GetLang() *string
	SetStatus(v int32) *ModifyEventStatusRequest
	GetStatus() *int32
}

type ModifyEventStatusRequest struct {
	// Specifies whether to enhance the detection of the anomalous activity. Enhancing detection improves accuracy and increases the alert rate for anomalous activities.
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	Backed *bool `json:"Backed,omitempty" xml:"Backed,omitempty"`
	// The reason for handling the anomalous activity.
	//
	// example:
	//
	// Confirmed as violation
	DealReason *string `json:"DealReason,omitempty" xml:"DealReason,omitempty"`
	// The unique ID of the anomalous activity.
	//
	// > To handle an anomalous activity, you must provide its unique ID. You can obtain this ID by calling the **DescribeEvents*	- operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the request and response. The default value is **zh_cn**. Valid values:
	//
	// - **zh_cn**: Chinese.
	//
	// - **en_us**: English.
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The operation to perform on the anomalous activity.
	//
	// - **1**: Mark as false positive.
	//
	// - **2**: Confirm and handle the anomalous activity.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyEventStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEventStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyEventStatusRequest) GetBacked() *bool {
	return s.Backed
}

func (s *ModifyEventStatusRequest) GetDealReason() *string {
	return s.DealReason
}

func (s *ModifyEventStatusRequest) GetId() *int64 {
	return s.Id
}

func (s *ModifyEventStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyEventStatusRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModifyEventStatusRequest) SetBacked(v bool) *ModifyEventStatusRequest {
	s.Backed = &v
	return s
}

func (s *ModifyEventStatusRequest) SetDealReason(v string) *ModifyEventStatusRequest {
	s.DealReason = &v
	return s
}

func (s *ModifyEventStatusRequest) SetId(v int64) *ModifyEventStatusRequest {
	s.Id = &v
	return s
}

func (s *ModifyEventStatusRequest) SetLang(v string) *ModifyEventStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyEventStatusRequest) SetStatus(v int32) *ModifyEventStatusRequest {
	s.Status = &v
	return s
}

func (s *ModifyEventStatusRequest) Validate() error {
	return dara.Validate(s)
}
