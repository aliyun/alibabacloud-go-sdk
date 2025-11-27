// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallCloudAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedInstanceCount(v int32) *InstallCloudAppResponseBody
	GetFailedInstanceCount() *int32
	SetFailedInstances(v []*InstallCloudAppResponseBodyFailedInstances) *InstallCloudAppResponseBody
	GetFailedInstances() []*InstallCloudAppResponseBodyFailedInstances
	SetRequestId(v string) *InstallCloudAppResponseBody
	GetRequestId() *string
	SetSuccessInstanceCount(v int32) *InstallCloudAppResponseBody
	GetSuccessInstanceCount() *int32
	SetSuccessInstances(v []*InstallCloudAppResponseBodySuccessInstances) *InstallCloudAppResponseBody
	GetSuccessInstances() []*InstallCloudAppResponseBodySuccessInstances
}

type InstallCloudAppResponseBody struct {
	FailedInstanceCount *int32                                        `json:"FailedInstanceCount,omitempty" xml:"FailedInstanceCount,omitempty"`
	FailedInstances     []*InstallCloudAppResponseBodyFailedInstances `json:"FailedInstances,omitempty" xml:"FailedInstances,omitempty" type:"Repeated"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId            *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessInstanceCount *int32                                         `json:"SuccessInstanceCount,omitempty" xml:"SuccessInstanceCount,omitempty"`
	SuccessInstances     []*InstallCloudAppResponseBodySuccessInstances `json:"SuccessInstances,omitempty" xml:"SuccessInstances,omitempty" type:"Repeated"`
}

func (s InstallCloudAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallCloudAppResponseBody) GoString() string {
	return s.String()
}

func (s *InstallCloudAppResponseBody) GetFailedInstanceCount() *int32 {
	return s.FailedInstanceCount
}

func (s *InstallCloudAppResponseBody) GetFailedInstances() []*InstallCloudAppResponseBodyFailedInstances {
	return s.FailedInstances
}

func (s *InstallCloudAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallCloudAppResponseBody) GetSuccessInstanceCount() *int32 {
	return s.SuccessInstanceCount
}

func (s *InstallCloudAppResponseBody) GetSuccessInstances() []*InstallCloudAppResponseBodySuccessInstances {
	return s.SuccessInstances
}

func (s *InstallCloudAppResponseBody) SetFailedInstanceCount(v int32) *InstallCloudAppResponseBody {
	s.FailedInstanceCount = &v
	return s
}

func (s *InstallCloudAppResponseBody) SetFailedInstances(v []*InstallCloudAppResponseBodyFailedInstances) *InstallCloudAppResponseBody {
	s.FailedInstances = v
	return s
}

func (s *InstallCloudAppResponseBody) SetRequestId(v string) *InstallCloudAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallCloudAppResponseBody) SetSuccessInstanceCount(v int32) *InstallCloudAppResponseBody {
	s.SuccessInstanceCount = &v
	return s
}

func (s *InstallCloudAppResponseBody) SetSuccessInstances(v []*InstallCloudAppResponseBodySuccessInstances) *InstallCloudAppResponseBody {
	s.SuccessInstances = v
	return s
}

func (s *InstallCloudAppResponseBody) Validate() error {
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

type InstallCloudAppResponseBodyFailedInstances struct {
	ErrCode             *int32  `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage          *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s InstallCloudAppResponseBodyFailedInstances) String() string {
	return dara.Prettify(s)
}

func (s InstallCloudAppResponseBodyFailedInstances) GoString() string {
	return s.String()
}

func (s *InstallCloudAppResponseBodyFailedInstances) GetErrCode() *int32 {
	return s.ErrCode
}

func (s *InstallCloudAppResponseBodyFailedInstances) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *InstallCloudAppResponseBodyFailedInstances) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *InstallCloudAppResponseBodyFailedInstances) SetErrCode(v int32) *InstallCloudAppResponseBodyFailedInstances {
	s.ErrCode = &v
	return s
}

func (s *InstallCloudAppResponseBodyFailedInstances) SetErrMessage(v string) *InstallCloudAppResponseBodyFailedInstances {
	s.ErrMessage = &v
	return s
}

func (s *InstallCloudAppResponseBodyFailedInstances) SetRenderingInstanceId(v string) *InstallCloudAppResponseBodyFailedInstances {
	s.RenderingInstanceId = &v
	return s
}

func (s *InstallCloudAppResponseBodyFailedInstances) Validate() error {
	return dara.Validate(s)
}

type InstallCloudAppResponseBodySuccessInstances struct {
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s InstallCloudAppResponseBodySuccessInstances) String() string {
	return dara.Prettify(s)
}

func (s InstallCloudAppResponseBodySuccessInstances) GoString() string {
	return s.String()
}

func (s *InstallCloudAppResponseBodySuccessInstances) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *InstallCloudAppResponseBodySuccessInstances) SetRenderingInstanceId(v string) *InstallCloudAppResponseBodySuccessInstances {
	s.RenderingInstanceId = &v
	return s
}

func (s *InstallCloudAppResponseBodySuccessInstances) Validate() error {
	return dara.Validate(s)
}
