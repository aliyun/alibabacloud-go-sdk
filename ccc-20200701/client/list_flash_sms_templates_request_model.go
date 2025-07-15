// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlashSmsTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListFlashSmsTemplatesRequest
	GetApplicationId() *string
	SetInstanceId(v string) *ListFlashSmsTemplatesRequest
	GetInstanceId() *string
	SetProviderId(v string) *ListFlashSmsTemplatesRequest
	GetProviderId() *string
}

type ListFlashSmsTemplatesRequest struct {
	// example:
	//
	// 71b396fa-1*********-70b7c0
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s *ListFlashSmsTemplatesRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListFlashSmsTemplatesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFlashSmsTemplatesRequest) GetProviderId() *string {
	return s.ProviderId
}

func (s *ListFlashSmsTemplatesRequest) SetApplicationId(v string) *ListFlashSmsTemplatesRequest {
	s.ApplicationId = &v
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
