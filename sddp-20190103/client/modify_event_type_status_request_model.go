// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEventTypeStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureType(v int32) *ModifyEventTypeStatusRequest
	GetFeatureType() *int32
	SetLang(v string) *ModifyEventTypeStatusRequest
	GetLang() *string
	SetSubTypeIds(v string) *ModifyEventTypeStatusRequest
	GetSubTypeIds() *string
}

type ModifyEventTypeStatusRequest struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 2
	FeatureType *int32 `json:"FeatureType,omitempty" xml:"FeatureType,omitempty"`
	// The language of the request and response. Valid values are **zh*	- for Chinese and **en*	- for English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The unique IDs of the anomalous activity subtypes. Separate multiple IDs with commas.
	//
	// > To enable SDDP to detect anomalous activities for subtypes, provide the unique IDs of the anomalous activity subtypes. Call the **DescribeEventTypes*	- operation to obtain the IDs.
	//
	// example:
	//
	// 020008
	SubTypeIds *string `json:"SubTypeIds,omitempty" xml:"SubTypeIds,omitempty"`
}

func (s ModifyEventTypeStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEventTypeStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyEventTypeStatusRequest) GetFeatureType() *int32 {
	return s.FeatureType
}

func (s *ModifyEventTypeStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyEventTypeStatusRequest) GetSubTypeIds() *string {
	return s.SubTypeIds
}

func (s *ModifyEventTypeStatusRequest) SetFeatureType(v int32) *ModifyEventTypeStatusRequest {
	s.FeatureType = &v
	return s
}

func (s *ModifyEventTypeStatusRequest) SetLang(v string) *ModifyEventTypeStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyEventTypeStatusRequest) SetSubTypeIds(v string) *ModifyEventTypeStatusRequest {
	s.SubTypeIds = &v
	return s
}

func (s *ModifyEventTypeStatusRequest) Validate() error {
	return dara.Validate(s)
}
