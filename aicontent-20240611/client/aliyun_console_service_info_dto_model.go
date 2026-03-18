// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAliyunConsoleServiceInfoDTO interface {
	dara.Model
	String() string
	GoString() string
	SetBuyUrl(v string) *AliyunConsoleServiceInfoDTO
	GetBuyUrl() *string
	SetDocumentUrl(v string) *AliyunConsoleServiceInfoDTO
	GetDocumentUrl() *string
	SetFreeConcurrencyCount(v int32) *AliyunConsoleServiceInfoDTO
	GetFreeConcurrencyCount() *int32
	SetFreeCount(v int32) *AliyunConsoleServiceInfoDTO
	GetFreeCount() *int32
	SetServiceCode(v string) *AliyunConsoleServiceInfoDTO
	GetServiceCode() *string
	SetServiceName(v string) *AliyunConsoleServiceInfoDTO
	GetServiceName() *string
}

type AliyunConsoleServiceInfoDTO struct {
	// example:
	//
	// https://www.aliyun.com/product/ai-algorithm
	BuyUrl *string `json:"buyUrl,omitempty" xml:"buyUrl,omitempty"`
	// example:
	//
	// https://www.aliyun.com/product/ai-algorithm
	DocumentUrl *string `json:"documentUrl,omitempty" xml:"documentUrl,omitempty"`
	// example:
	//
	// 10
	FreeConcurrencyCount *int32 `json:"freeConcurrencyCount,omitempty" xml:"freeConcurrencyCount,omitempty"`
	// example:
	//
	// 100
	FreeCount *int32 `json:"freeCount,omitempty" xml:"freeCount,omitempty"`
	// example:
	//
	// online_ai_algorithm_personalized_text_to_image_call_count
	ServiceCode *string `json:"serviceCode,omitempty" xml:"serviceCode,omitempty"`
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s AliyunConsoleServiceInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s AliyunConsoleServiceInfoDTO) GoString() string {
	return s.String()
}

func (s *AliyunConsoleServiceInfoDTO) GetBuyUrl() *string {
	return s.BuyUrl
}

func (s *AliyunConsoleServiceInfoDTO) GetDocumentUrl() *string {
	return s.DocumentUrl
}

func (s *AliyunConsoleServiceInfoDTO) GetFreeConcurrencyCount() *int32 {
	return s.FreeConcurrencyCount
}

func (s *AliyunConsoleServiceInfoDTO) GetFreeCount() *int32 {
	return s.FreeCount
}

func (s *AliyunConsoleServiceInfoDTO) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *AliyunConsoleServiceInfoDTO) GetServiceName() *string {
	return s.ServiceName
}

func (s *AliyunConsoleServiceInfoDTO) SetBuyUrl(v string) *AliyunConsoleServiceInfoDTO {
	s.BuyUrl = &v
	return s
}

func (s *AliyunConsoleServiceInfoDTO) SetDocumentUrl(v string) *AliyunConsoleServiceInfoDTO {
	s.DocumentUrl = &v
	return s
}

func (s *AliyunConsoleServiceInfoDTO) SetFreeConcurrencyCount(v int32) *AliyunConsoleServiceInfoDTO {
	s.FreeConcurrencyCount = &v
	return s
}

func (s *AliyunConsoleServiceInfoDTO) SetFreeCount(v int32) *AliyunConsoleServiceInfoDTO {
	s.FreeCount = &v
	return s
}

func (s *AliyunConsoleServiceInfoDTO) SetServiceCode(v string) *AliyunConsoleServiceInfoDTO {
	s.ServiceCode = &v
	return s
}

func (s *AliyunConsoleServiceInfoDTO) SetServiceName(v string) *AliyunConsoleServiceInfoDTO {
	s.ServiceName = &v
	return s
}

func (s *AliyunConsoleServiceInfoDTO) Validate() error {
	return dara.Validate(s)
}
