// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallCloudAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedInstanceCount(v int32) *UninstallCloudAppResponseBody
	GetFailedInstanceCount() *int32
	SetFailedInstances(v []*UninstallCloudAppResponseBodyFailedInstances) *UninstallCloudAppResponseBody
	GetFailedInstances() []*UninstallCloudAppResponseBodyFailedInstances
	SetRequestId(v string) *UninstallCloudAppResponseBody
	GetRequestId() *string
	SetSuccessInstanceCount(v int32) *UninstallCloudAppResponseBody
	GetSuccessInstanceCount() *int32
	SetSuccessInstances(v []*UninstallCloudAppResponseBodySuccessInstances) *UninstallCloudAppResponseBody
	GetSuccessInstances() []*UninstallCloudAppResponseBodySuccessInstances
}

type UninstallCloudAppResponseBody struct {
	FailedInstanceCount *int32                                          `json:"FailedInstanceCount,omitempty" xml:"FailedInstanceCount,omitempty"`
	FailedInstances     []*UninstallCloudAppResponseBodyFailedInstances `json:"FailedInstances,omitempty" xml:"FailedInstances,omitempty" type:"Repeated"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId            *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessInstanceCount *int32                                           `json:"SuccessInstanceCount,omitempty" xml:"SuccessInstanceCount,omitempty"`
	SuccessInstances     []*UninstallCloudAppResponseBodySuccessInstances `json:"SuccessInstances,omitempty" xml:"SuccessInstances,omitempty" type:"Repeated"`
}

func (s UninstallCloudAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallCloudAppResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallCloudAppResponseBody) GetFailedInstanceCount() *int32 {
	return s.FailedInstanceCount
}

func (s *UninstallCloudAppResponseBody) GetFailedInstances() []*UninstallCloudAppResponseBodyFailedInstances {
	return s.FailedInstances
}

func (s *UninstallCloudAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallCloudAppResponseBody) GetSuccessInstanceCount() *int32 {
	return s.SuccessInstanceCount
}

func (s *UninstallCloudAppResponseBody) GetSuccessInstances() []*UninstallCloudAppResponseBodySuccessInstances {
	return s.SuccessInstances
}

func (s *UninstallCloudAppResponseBody) SetFailedInstanceCount(v int32) *UninstallCloudAppResponseBody {
	s.FailedInstanceCount = &v
	return s
}

func (s *UninstallCloudAppResponseBody) SetFailedInstances(v []*UninstallCloudAppResponseBodyFailedInstances) *UninstallCloudAppResponseBody {
	s.FailedInstances = v
	return s
}

func (s *UninstallCloudAppResponseBody) SetRequestId(v string) *UninstallCloudAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallCloudAppResponseBody) SetSuccessInstanceCount(v int32) *UninstallCloudAppResponseBody {
	s.SuccessInstanceCount = &v
	return s
}

func (s *UninstallCloudAppResponseBody) SetSuccessInstances(v []*UninstallCloudAppResponseBodySuccessInstances) *UninstallCloudAppResponseBody {
	s.SuccessInstances = v
	return s
}

func (s *UninstallCloudAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type UninstallCloudAppResponseBodyFailedInstances struct {
	ErrCode             *int32  `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage          *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s UninstallCloudAppResponseBodyFailedInstances) String() string {
	return dara.Prettify(s)
}

func (s UninstallCloudAppResponseBodyFailedInstances) GoString() string {
	return s.String()
}

func (s *UninstallCloudAppResponseBodyFailedInstances) GetErrCode() *int32 {
	return s.ErrCode
}

func (s *UninstallCloudAppResponseBodyFailedInstances) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UninstallCloudAppResponseBodyFailedInstances) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *UninstallCloudAppResponseBodyFailedInstances) SetErrCode(v int32) *UninstallCloudAppResponseBodyFailedInstances {
	s.ErrCode = &v
	return s
}

func (s *UninstallCloudAppResponseBodyFailedInstances) SetErrMessage(v string) *UninstallCloudAppResponseBodyFailedInstances {
	s.ErrMessage = &v
	return s
}

func (s *UninstallCloudAppResponseBodyFailedInstances) SetRenderingInstanceId(v string) *UninstallCloudAppResponseBodyFailedInstances {
	s.RenderingInstanceId = &v
	return s
}

func (s *UninstallCloudAppResponseBodyFailedInstances) Validate() error {
	return dara.Validate(s)
}

type UninstallCloudAppResponseBodySuccessInstances struct {
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s UninstallCloudAppResponseBodySuccessInstances) String() string {
	return dara.Prettify(s)
}

func (s UninstallCloudAppResponseBodySuccessInstances) GoString() string {
	return s.String()
}

func (s *UninstallCloudAppResponseBodySuccessInstances) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *UninstallCloudAppResponseBodySuccessInstances) SetRenderingInstanceId(v string) *UninstallCloudAppResponseBodySuccessInstances {
	s.RenderingInstanceId = &v
	return s
}

func (s *UninstallCloudAppResponseBodySuccessInstances) Validate() error {
	return dara.Validate(s)
}
