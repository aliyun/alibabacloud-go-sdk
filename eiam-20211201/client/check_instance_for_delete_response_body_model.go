// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceForDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckInstanceResult(v *CheckInstanceForDeleteResponseBodyCheckInstanceResult) *CheckInstanceForDeleteResponseBody
	GetCheckInstanceResult() *CheckInstanceForDeleteResponseBodyCheckInstanceResult
	SetRequestId(v string) *CheckInstanceForDeleteResponseBody
	GetRequestId() *string
}

type CheckInstanceForDeleteResponseBody struct {
	CheckInstanceResult *CheckInstanceForDeleteResponseBodyCheckInstanceResult `json:"CheckInstanceResult,omitempty" xml:"CheckInstanceResult,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckInstanceForDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceForDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *CheckInstanceForDeleteResponseBody) GetCheckInstanceResult() *CheckInstanceForDeleteResponseBodyCheckInstanceResult {
	return s.CheckInstanceResult
}

func (s *CheckInstanceForDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckInstanceForDeleteResponseBody) SetCheckInstanceResult(v *CheckInstanceForDeleteResponseBodyCheckInstanceResult) *CheckInstanceForDeleteResponseBody {
	s.CheckInstanceResult = v
	return s
}

func (s *CheckInstanceForDeleteResponseBody) SetRequestId(v string) *CheckInstanceForDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckInstanceForDeleteResponseBody) Validate() error {
	if s.CheckInstanceResult != nil {
		if err := s.CheckInstanceResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckInstanceForDeleteResponseBodyCheckInstanceResult struct {
	// true表示实例可以被删除；false表示实例不可被删除，具体查看RestrictScenarios属性。
	//
	// example:
	//
	// True
	Deletable *bool `json:"Deletable,omitempty" xml:"Deletable,omitempty"`
	// true表示实例可以被删除；false表示实例不可被删除，具体查看RestrictScenarios属性。
	RestrictScenarios []*CheckInstanceForDeleteResponseBodyCheckInstanceResultRestrictScenarios `json:"RestrictScenarios,omitempty" xml:"RestrictScenarios,omitempty" type:"Repeated"`
}

func (s CheckInstanceForDeleteResponseBodyCheckInstanceResult) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceForDeleteResponseBodyCheckInstanceResult) GoString() string {
	return s.String()
}

func (s *CheckInstanceForDeleteResponseBodyCheckInstanceResult) GetDeletable() *bool {
	return s.Deletable
}

func (s *CheckInstanceForDeleteResponseBodyCheckInstanceResult) GetRestrictScenarios() []*CheckInstanceForDeleteResponseBodyCheckInstanceResultRestrictScenarios {
	return s.RestrictScenarios
}

func (s *CheckInstanceForDeleteResponseBodyCheckInstanceResult) SetDeletable(v bool) *CheckInstanceForDeleteResponseBodyCheckInstanceResult {
	s.Deletable = &v
	return s
}

func (s *CheckInstanceForDeleteResponseBodyCheckInstanceResult) SetRestrictScenarios(v []*CheckInstanceForDeleteResponseBodyCheckInstanceResultRestrictScenarios) *CheckInstanceForDeleteResponseBodyCheckInstanceResult {
	s.RestrictScenarios = v
	return s
}

func (s *CheckInstanceForDeleteResponseBodyCheckInstanceResult) Validate() error {
	if s.RestrictScenarios != nil {
		for _, item := range s.RestrictScenarios {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckInstanceForDeleteResponseBodyCheckInstanceResultRestrictScenarios struct {
	// 有帮助的控制台地址，可以管理对应的资源，从而解除实例删除限制。可能返回为空，不一定所有的资源ID都有管理地址返回。
	//
	// example:
	//
	// https://console-rpa.aliyun.com/
	HelpfulConsoleUrl *string `json:"HelpfulConsoleUrl,omitempty" xml:"HelpfulConsoleUrl,omitempty"`
	// 导致实例删除受限的资源ID。
	//
	// example:
	//
	// eas-r-nguosqgr75ndg784k8
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 针对实例删除受限的原因文字描述。
	//
	// example:
	//
	// cloud_product_dependency
	RestrictReason *string `json:"RestrictReason,omitempty" xml:"RestrictReason,omitempty"`
}

func (s CheckInstanceForDeleteResponseBodyCheckInstanceResultRestrictScenarios) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceForDeleteResponseBodyCheckInstanceResultRestrictScenarios) GoString() string {
	return s.String()
}

func (s *CheckInstanceForDeleteResponseBodyCheckInstanceResultRestrictScenarios) GetHelpfulConsoleUrl() *string {
	return s.HelpfulConsoleUrl
}

func (s *CheckInstanceForDeleteResponseBodyCheckInstanceResultRestrictScenarios) GetResourceId() *string {
	return s.ResourceId
}

func (s *CheckInstanceForDeleteResponseBodyCheckInstanceResultRestrictScenarios) GetRestrictReason() *string {
	return s.RestrictReason
}

func (s *CheckInstanceForDeleteResponseBodyCheckInstanceResultRestrictScenarios) SetHelpfulConsoleUrl(v string) *CheckInstanceForDeleteResponseBodyCheckInstanceResultRestrictScenarios {
	s.HelpfulConsoleUrl = &v
	return s
}

func (s *CheckInstanceForDeleteResponseBodyCheckInstanceResultRestrictScenarios) SetResourceId(v string) *CheckInstanceForDeleteResponseBodyCheckInstanceResultRestrictScenarios {
	s.ResourceId = &v
	return s
}

func (s *CheckInstanceForDeleteResponseBodyCheckInstanceResultRestrictScenarios) SetRestrictReason(v string) *CheckInstanceForDeleteResponseBodyCheckInstanceResultRestrictScenarios {
	s.RestrictReason = &v
	return s
}

func (s *CheckInstanceForDeleteResponseBodyCheckInstanceResultRestrictScenarios) Validate() error {
	return dara.Validate(s)
}
