// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlashSmsTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *ListFlashSmsTemplatesRequest
	GetConfigId() *string
	SetInstanceId(v string) *ListFlashSmsTemplatesRequest
	GetInstanceId() *string
	SetProviderId(v string) *ListFlashSmsTemplatesRequest
	GetProviderId() *string
}

type ListFlashSmsTemplatesRequest struct {
	// Flash SMS configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fe542524-9585-4cc7-be54-c8782ed7f60e
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 361c8a53-0e29-42f3-8aa7-c7752d010399
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Flash SMS provider ID.
	//
	// - Uincall: Beijing Uincall Communication Co., Ltd.
	//
	// - ShangHaiTianNan: Shanghai TianNan
	//
	// - HeDao: Galaxis
	//
	// This parameter is required.
	//
	// example:
	//
	// Uincall
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
}

func (s ListFlashSmsTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListFlashSmsTemplatesRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *ListFlashSmsTemplatesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFlashSmsTemplatesRequest) GetProviderId() *string {
	return s.ProviderId
}

func (s *ListFlashSmsTemplatesRequest) SetConfigId(v string) *ListFlashSmsTemplatesRequest {
	s.ConfigId = &v
	return s
}

func (s *ListFlashSmsTemplatesRequest) SetInstanceId(v string) *ListFlashSmsTemplatesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFlashSmsTemplatesRequest) SetProviderId(v string) *ListFlashSmsTemplatesRequest {
	s.ProviderId = &v
	return s
}

func (s *ListFlashSmsTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
