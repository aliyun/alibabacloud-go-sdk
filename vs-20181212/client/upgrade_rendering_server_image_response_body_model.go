// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeRenderingServerImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedInstanceCount(v int32) *UpgradeRenderingServerImageResponseBody
	GetFailedInstanceCount() *int32
	SetFailedInstances(v []*UpgradeRenderingServerImageResponseBodyFailedInstances) *UpgradeRenderingServerImageResponseBody
	GetFailedInstances() []*UpgradeRenderingServerImageResponseBodyFailedInstances
	SetRequestId(v string) *UpgradeRenderingServerImageResponseBody
	GetRequestId() *string
	SetSuccessInstanceCount(v int32) *UpgradeRenderingServerImageResponseBody
	GetSuccessInstanceCount() *int32
	SetSuccessInstances(v []*UpgradeRenderingServerImageResponseBodySuccessInstances) *UpgradeRenderingServerImageResponseBody
	GetSuccessInstances() []*UpgradeRenderingServerImageResponseBodySuccessInstances
}

type UpgradeRenderingServerImageResponseBody struct {
	// example:
	//
	// 0
	FailedInstanceCount *int32                                                    `json:"FailedInstanceCount,omitempty" xml:"FailedInstanceCount,omitempty"`
	FailedInstances     []*UpgradeRenderingServerImageResponseBodyFailedInstances `json:"FailedInstances,omitempty" xml:"FailedInstances,omitempty" type:"Repeated"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	SuccessInstanceCount *int32 `json:"SuccessInstanceCount,omitempty" xml:"SuccessInstanceCount,omitempty"`
	// example:
	//
	// ["render-4cb88f452e11448095cde66e1ecea9d9", "render-507538b7906742449680742a194f831e"]
	SuccessInstances []*UpgradeRenderingServerImageResponseBodySuccessInstances `json:"SuccessInstances,omitempty" xml:"SuccessInstances,omitempty" type:"Repeated"`
}

func (s UpgradeRenderingServerImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeRenderingServerImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeRenderingServerImageResponseBody) GetFailedInstanceCount() *int32 {
	return s.FailedInstanceCount
}

func (s *UpgradeRenderingServerImageResponseBody) GetFailedInstances() []*UpgradeRenderingServerImageResponseBodyFailedInstances {
	return s.FailedInstances
}

func (s *UpgradeRenderingServerImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeRenderingServerImageResponseBody) GetSuccessInstanceCount() *int32 {
	return s.SuccessInstanceCount
}

func (s *UpgradeRenderingServerImageResponseBody) GetSuccessInstances() []*UpgradeRenderingServerImageResponseBodySuccessInstances {
	return s.SuccessInstances
}

func (s *UpgradeRenderingServerImageResponseBody) SetFailedInstanceCount(v int32) *UpgradeRenderingServerImageResponseBody {
	s.FailedInstanceCount = &v
	return s
}

func (s *UpgradeRenderingServerImageResponseBody) SetFailedInstances(v []*UpgradeRenderingServerImageResponseBodyFailedInstances) *UpgradeRenderingServerImageResponseBody {
	s.FailedInstances = v
	return s
}

func (s *UpgradeRenderingServerImageResponseBody) SetRequestId(v string) *UpgradeRenderingServerImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeRenderingServerImageResponseBody) SetSuccessInstanceCount(v int32) *UpgradeRenderingServerImageResponseBody {
	s.SuccessInstanceCount = &v
	return s
}

func (s *UpgradeRenderingServerImageResponseBody) SetSuccessInstances(v []*UpgradeRenderingServerImageResponseBodySuccessInstances) *UpgradeRenderingServerImageResponseBody {
	s.SuccessInstances = v
	return s
}

func (s *UpgradeRenderingServerImageResponseBody) Validate() error {
	if s.FailedInstances != nil {
		for _, item := range s.FailedInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SuccessInstances != nil {
		for _, item := range s.SuccessInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpgradeRenderingServerImageResponseBodyFailedInstances struct {
	// example:
	//
	// 300000
	ErrCode *int32 `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Rejected due to timeout
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// render-421cd2a1125947c19fcd5c7dd2c7d31e
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s UpgradeRenderingServerImageResponseBodyFailedInstances) String() string {
	return dara.Prettify(s)
}

func (s UpgradeRenderingServerImageResponseBodyFailedInstances) GoString() string {
	return s.String()
}

func (s *UpgradeRenderingServerImageResponseBodyFailedInstances) GetErrCode() *int32 {
	return s.ErrCode
}

func (s *UpgradeRenderingServerImageResponseBodyFailedInstances) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpgradeRenderingServerImageResponseBodyFailedInstances) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *UpgradeRenderingServerImageResponseBodyFailedInstances) SetErrCode(v int32) *UpgradeRenderingServerImageResponseBodyFailedInstances {
	s.ErrCode = &v
	return s
}

func (s *UpgradeRenderingServerImageResponseBodyFailedInstances) SetErrMessage(v string) *UpgradeRenderingServerImageResponseBodyFailedInstances {
	s.ErrMessage = &v
	return s
}

func (s *UpgradeRenderingServerImageResponseBodyFailedInstances) SetRenderingInstanceId(v string) *UpgradeRenderingServerImageResponseBodyFailedInstances {
	s.RenderingInstanceId = &v
	return s
}

func (s *UpgradeRenderingServerImageResponseBodyFailedInstances) Validate() error {
	return dara.Validate(s)
}

type UpgradeRenderingServerImageResponseBodySuccessInstances struct {
	// example:
	//
	// render-e6cf423c787e4e43b460a788da254fe3
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s UpgradeRenderingServerImageResponseBodySuccessInstances) String() string {
	return dara.Prettify(s)
}

func (s UpgradeRenderingServerImageResponseBodySuccessInstances) GoString() string {
	return s.String()
}

func (s *UpgradeRenderingServerImageResponseBodySuccessInstances) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *UpgradeRenderingServerImageResponseBodySuccessInstances) SetRenderingInstanceId(v string) *UpgradeRenderingServerImageResponseBodySuccessInstances {
	s.RenderingInstanceId = &v
	return s
}

func (s *UpgradeRenderingServerImageResponseBodySuccessInstances) Validate() error {
	return dara.Validate(s)
}
