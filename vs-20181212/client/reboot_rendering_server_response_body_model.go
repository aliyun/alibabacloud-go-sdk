// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootRenderingServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedInstanceCount(v int32) *RebootRenderingServerResponseBody
	GetFailedInstanceCount() *int32
	SetFailedInstances(v []*RebootRenderingServerResponseBodyFailedInstances) *RebootRenderingServerResponseBody
	GetFailedInstances() []*RebootRenderingServerResponseBodyFailedInstances
	SetRequestId(v string) *RebootRenderingServerResponseBody
	GetRequestId() *string
	SetSuccessInstanceCount(v int32) *RebootRenderingServerResponseBody
	GetSuccessInstanceCount() *int32
	SetSuccessInstances(v []*RebootRenderingServerResponseBodySuccessInstances) *RebootRenderingServerResponseBody
	GetSuccessInstances() []*RebootRenderingServerResponseBodySuccessInstances
}

type RebootRenderingServerResponseBody struct {
	// example:
	//
	// 0
	FailedInstanceCount *int32                                              `json:"FailedInstanceCount,omitempty" xml:"FailedInstanceCount,omitempty"`
	FailedInstances     []*RebootRenderingServerResponseBodyFailedInstances `json:"FailedInstances,omitempty" xml:"FailedInstances,omitempty" type:"Repeated"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	SuccessInstanceCount *int32                                               `json:"SuccessInstanceCount,omitempty" xml:"SuccessInstanceCount,omitempty"`
	SuccessInstances     []*RebootRenderingServerResponseBodySuccessInstances `json:"SuccessInstances,omitempty" xml:"SuccessInstances,omitempty" type:"Repeated"`
}

func (s RebootRenderingServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebootRenderingServerResponseBody) GoString() string {
	return s.String()
}

func (s *RebootRenderingServerResponseBody) GetFailedInstanceCount() *int32 {
	return s.FailedInstanceCount
}

func (s *RebootRenderingServerResponseBody) GetFailedInstances() []*RebootRenderingServerResponseBodyFailedInstances {
	return s.FailedInstances
}

func (s *RebootRenderingServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebootRenderingServerResponseBody) GetSuccessInstanceCount() *int32 {
	return s.SuccessInstanceCount
}

func (s *RebootRenderingServerResponseBody) GetSuccessInstances() []*RebootRenderingServerResponseBodySuccessInstances {
	return s.SuccessInstances
}

func (s *RebootRenderingServerResponseBody) SetFailedInstanceCount(v int32) *RebootRenderingServerResponseBody {
	s.FailedInstanceCount = &v
	return s
}

func (s *RebootRenderingServerResponseBody) SetFailedInstances(v []*RebootRenderingServerResponseBodyFailedInstances) *RebootRenderingServerResponseBody {
	s.FailedInstances = v
	return s
}

func (s *RebootRenderingServerResponseBody) SetRequestId(v string) *RebootRenderingServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootRenderingServerResponseBody) SetSuccessInstanceCount(v int32) *RebootRenderingServerResponseBody {
	s.SuccessInstanceCount = &v
	return s
}

func (s *RebootRenderingServerResponseBody) SetSuccessInstances(v []*RebootRenderingServerResponseBodySuccessInstances) *RebootRenderingServerResponseBody {
	s.SuccessInstances = v
	return s
}

func (s *RebootRenderingServerResponseBody) Validate() error {
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

type RebootRenderingServerResponseBodyFailedInstances struct {
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

func (s RebootRenderingServerResponseBodyFailedInstances) String() string {
	return dara.Prettify(s)
}

func (s RebootRenderingServerResponseBodyFailedInstances) GoString() string {
	return s.String()
}

func (s *RebootRenderingServerResponseBodyFailedInstances) GetErrCode() *int32 {
	return s.ErrCode
}

func (s *RebootRenderingServerResponseBodyFailedInstances) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *RebootRenderingServerResponseBodyFailedInstances) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *RebootRenderingServerResponseBodyFailedInstances) SetErrCode(v int32) *RebootRenderingServerResponseBodyFailedInstances {
	s.ErrCode = &v
	return s
}

func (s *RebootRenderingServerResponseBodyFailedInstances) SetErrMessage(v string) *RebootRenderingServerResponseBodyFailedInstances {
	s.ErrMessage = &v
	return s
}

func (s *RebootRenderingServerResponseBodyFailedInstances) SetRenderingInstanceId(v string) *RebootRenderingServerResponseBodyFailedInstances {
	s.RenderingInstanceId = &v
	return s
}

func (s *RebootRenderingServerResponseBodyFailedInstances) Validate() error {
	return dara.Validate(s)
}

type RebootRenderingServerResponseBodySuccessInstances struct {
	// example:
	//
	// render-e6cf423c787e4e43b460a788da254fe3
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s RebootRenderingServerResponseBodySuccessInstances) String() string {
	return dara.Prettify(s)
}

func (s RebootRenderingServerResponseBodySuccessInstances) GoString() string {
	return s.String()
}

func (s *RebootRenderingServerResponseBodySuccessInstances) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *RebootRenderingServerResponseBodySuccessInstances) SetRenderingInstanceId(v string) *RebootRenderingServerResponseBodySuccessInstances {
	s.RenderingInstanceId = &v
	return s
}

func (s *RebootRenderingServerResponseBodySuccessInstances) Validate() error {
	return dara.Validate(s)
}
