// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCheckInstanceResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *VerifyCheckInstanceResultResponseBodyData) *VerifyCheckInstanceResultResponseBody
	GetData() *VerifyCheckInstanceResultResponseBodyData
	SetRequestId(v string) *VerifyCheckInstanceResultResponseBody
	GetRequestId() *string
}

type VerifyCheckInstanceResultResponseBody struct {
	// The data returned.
	Data *VerifyCheckInstanceResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// B869E3A0-1147-539D-9920-47580700****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyCheckInstanceResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyCheckInstanceResultResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyCheckInstanceResultResponseBody) GetData() *VerifyCheckInstanceResultResponseBodyData {
	return s.Data
}

func (s *VerifyCheckInstanceResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyCheckInstanceResultResponseBody) SetData(v *VerifyCheckInstanceResultResponseBodyData) *VerifyCheckInstanceResultResponseBody {
	s.Data = v
	return s
}

func (s *VerifyCheckInstanceResultResponseBody) SetRequestId(v string) *VerifyCheckInstanceResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyCheckInstanceResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type VerifyCheckInstanceResultResponseBodyData struct {
	// An array consisting of instances that failed the check.
	FailInstances []*string `json:"FailInstances,omitempty" xml:"FailInstances,omitempty" type:"Repeated"`
	// The operation code of the task that checks the configurations of cloud services. Valid values:
	//
	// 	- **Throttling**
	//
	// 	- **ActionTrialUnauthorized**
	//
	// example:
	//
	// ActionTrialUnauthorized
	OperateCode *string `json:"OperateCode,omitempty" xml:"OperateCode,omitempty"`
	// The task ID.
	//
	// example:
	//
	// a410bb3-e68c217a-3368bc0-238c668***
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s VerifyCheckInstanceResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VerifyCheckInstanceResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *VerifyCheckInstanceResultResponseBodyData) GetFailInstances() []*string {
	return s.FailInstances
}

func (s *VerifyCheckInstanceResultResponseBodyData) GetOperateCode() *string {
	return s.OperateCode
}

func (s *VerifyCheckInstanceResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *VerifyCheckInstanceResultResponseBodyData) SetFailInstances(v []*string) *VerifyCheckInstanceResultResponseBodyData {
	s.FailInstances = v
	return s
}

func (s *VerifyCheckInstanceResultResponseBodyData) SetOperateCode(v string) *VerifyCheckInstanceResultResponseBodyData {
	s.OperateCode = &v
	return s
}

func (s *VerifyCheckInstanceResultResponseBodyData) SetTaskId(v string) *VerifyCheckInstanceResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *VerifyCheckInstanceResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
