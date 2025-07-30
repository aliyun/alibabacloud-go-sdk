// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreCheckCreateGadOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *PreCheckCreateGadOrderResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *PreCheckCreateGadOrderResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *PreCheckCreateGadOrderResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *PreCheckCreateGadOrderResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *PreCheckCreateGadOrderResponseBody
	GetHttpStatusCode() *string
	SetInstanceId(v string) *PreCheckCreateGadOrderResponseBody
	GetInstanceId() *string
	SetRegionId(v string) *PreCheckCreateGadOrderResponseBody
	GetRegionId() *string
	SetRequestId(v string) *PreCheckCreateGadOrderResponseBody
	GetRequestId() *string
	SetSuccess(v string) *PreCheckCreateGadOrderResponseBody
	GetSuccess() *string
	SetTaskId(v string) *PreCheckCreateGadOrderResponseBody
	GetTaskId() *string
}

type PreCheckCreateGadOrderResponseBody struct {
	// example:
	//
	// 403
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// rm-bp162d4tp0500****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// C166D79D-436B-45F0-B5A5-25E1959F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// z2v12jfo309****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PreCheckCreateGadOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreCheckCreateGadOrderResponseBody) GoString() string {
	return s.String()
}

func (s *PreCheckCreateGadOrderResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *PreCheckCreateGadOrderResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *PreCheckCreateGadOrderResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *PreCheckCreateGadOrderResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *PreCheckCreateGadOrderResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *PreCheckCreateGadOrderResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PreCheckCreateGadOrderResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *PreCheckCreateGadOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreCheckCreateGadOrderResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *PreCheckCreateGadOrderResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *PreCheckCreateGadOrderResponseBody) SetDynamicCode(v string) *PreCheckCreateGadOrderResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *PreCheckCreateGadOrderResponseBody) SetDynamicMessage(v string) *PreCheckCreateGadOrderResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *PreCheckCreateGadOrderResponseBody) SetErrCode(v string) *PreCheckCreateGadOrderResponseBody {
	s.ErrCode = &v
	return s
}

func (s *PreCheckCreateGadOrderResponseBody) SetErrMessage(v string) *PreCheckCreateGadOrderResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *PreCheckCreateGadOrderResponseBody) SetHttpStatusCode(v string) *PreCheckCreateGadOrderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PreCheckCreateGadOrderResponseBody) SetInstanceId(v string) *PreCheckCreateGadOrderResponseBody {
	s.InstanceId = &v
	return s
}

func (s *PreCheckCreateGadOrderResponseBody) SetRegionId(v string) *PreCheckCreateGadOrderResponseBody {
	s.RegionId = &v
	return s
}

func (s *PreCheckCreateGadOrderResponseBody) SetRequestId(v string) *PreCheckCreateGadOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreCheckCreateGadOrderResponseBody) SetSuccess(v string) *PreCheckCreateGadOrderResponseBody {
	s.Success = &v
	return s
}

func (s *PreCheckCreateGadOrderResponseBody) SetTaskId(v string) *PreCheckCreateGadOrderResponseBody {
	s.TaskId = &v
	return s
}

func (s *PreCheckCreateGadOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
