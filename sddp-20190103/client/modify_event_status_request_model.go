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
	// Specifies whether to enhance the detection of anomalous events. If you enhance the detection of anomalous events, the detection accuracy and the rate of triggering alerts for anomalous events are improved. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Backed *bool `json:"Backed,omitempty" xml:"Backed,omitempty"`
	// The reason why the anomalous event is handled.
	//
	// example:
	//
	// Anomaly confirmed
	DealReason *string `json:"DealReason,omitempty" xml:"DealReason,omitempty"`
	// The ID of the anomalous event.
	//
	// > You can call the **DescribeEvents*	- operation to query the ID of the anomalous event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The method to handle the anomalous event. Valid values:
	//
	// 	- **1**: marks the anomalous event as a false positive.
	//
	// 	- **2**: confirms and handles the anomalous event.
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
