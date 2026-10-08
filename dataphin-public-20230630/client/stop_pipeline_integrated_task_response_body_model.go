// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPipelineIntegratedTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopPipelineIntegratedTaskResponseBody
	GetCode() *string
	SetData(v *StopPipelineIntegratedTaskResponseBodyData) *StopPipelineIntegratedTaskResponseBody
	GetData() *StopPipelineIntegratedTaskResponseBodyData
	SetHttpStatusCode(v int32) *StopPipelineIntegratedTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StopPipelineIntegratedTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopPipelineIntegratedTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopPipelineIntegratedTaskResponseBody
	GetSuccess() *bool
}

type StopPipelineIntegratedTaskResponseBody struct {
	// example:
	//
	// OK
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *StopPipelineIntegratedTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopPipelineIntegratedTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopPipelineIntegratedTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopPipelineIntegratedTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopPipelineIntegratedTaskResponseBody) GetData() *StopPipelineIntegratedTaskResponseBodyData {
	return s.Data
}

func (s *StopPipelineIntegratedTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StopPipelineIntegratedTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopPipelineIntegratedTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopPipelineIntegratedTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopPipelineIntegratedTaskResponseBody) SetCode(v string) *StopPipelineIntegratedTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StopPipelineIntegratedTaskResponseBody) SetData(v *StopPipelineIntegratedTaskResponseBodyData) *StopPipelineIntegratedTaskResponseBody {
	s.Data = v
	return s
}

func (s *StopPipelineIntegratedTaskResponseBody) SetHttpStatusCode(v int32) *StopPipelineIntegratedTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopPipelineIntegratedTaskResponseBody) SetMessage(v string) *StopPipelineIntegratedTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StopPipelineIntegratedTaskResponseBody) SetRequestId(v string) *StopPipelineIntegratedTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopPipelineIntegratedTaskResponseBody) SetSuccess(v bool) *StopPipelineIntegratedTaskResponseBody {
	s.Success = &v
	return s
}

func (s *StopPipelineIntegratedTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StopPipelineIntegratedTaskResponseBodyData struct {
	DevOpsActionResDTOList []*StopPipelineIntegratedTaskResponseBodyDataDevOpsActionResDTOList `json:"DevOpsActionResDTOList,omitempty" xml:"DevOpsActionResDTOList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	Fail *int64 `json:"Fail,omitempty" xml:"Fail,omitempty"`
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopPipelineIntegratedTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StopPipelineIntegratedTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *StopPipelineIntegratedTaskResponseBodyData) GetDevOpsActionResDTOList() []*StopPipelineIntegratedTaskResponseBodyDataDevOpsActionResDTOList {
	return s.DevOpsActionResDTOList
}

func (s *StopPipelineIntegratedTaskResponseBodyData) GetFail() *int64 {
	return s.Fail
}

func (s *StopPipelineIntegratedTaskResponseBodyData) GetSuccess() *int64 {
	return s.Success
}

func (s *StopPipelineIntegratedTaskResponseBodyData) SetDevOpsActionResDTOList(v []*StopPipelineIntegratedTaskResponseBodyDataDevOpsActionResDTOList) *StopPipelineIntegratedTaskResponseBodyData {
	s.DevOpsActionResDTOList = v
	return s
}

func (s *StopPipelineIntegratedTaskResponseBodyData) SetFail(v int64) *StopPipelineIntegratedTaskResponseBodyData {
	s.Fail = &v
	return s
}

func (s *StopPipelineIntegratedTaskResponseBodyData) SetSuccess(v int64) *StopPipelineIntegratedTaskResponseBodyData {
	s.Success = &v
	return s
}

func (s *StopPipelineIntegratedTaskResponseBodyData) Validate() error {
	if s.DevOpsActionResDTOList != nil {
		for _, item := range s.DevOpsActionResDTOList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type StopPipelineIntegratedTaskResponseBodyDataDevOpsActionResDTOList struct {
	// example:
	//
	// eedsauto
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// 1692173406264688
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// FAIED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s StopPipelineIntegratedTaskResponseBodyDataDevOpsActionResDTOList) String() string {
	return dara.Prettify(s)
}

func (s StopPipelineIntegratedTaskResponseBodyDataDevOpsActionResDTOList) GoString() string {
	return s.String()
}

func (s *StopPipelineIntegratedTaskResponseBodyDataDevOpsActionResDTOList) GetJobName() *string {
	return s.JobName
}

func (s *StopPipelineIntegratedTaskResponseBodyDataDevOpsActionResDTOList) GetOwner() *string {
	return s.Owner
}

func (s *StopPipelineIntegratedTaskResponseBodyDataDevOpsActionResDTOList) GetStatus() *string {
	return s.Status
}

func (s *StopPipelineIntegratedTaskResponseBodyDataDevOpsActionResDTOList) SetJobName(v string) *StopPipelineIntegratedTaskResponseBodyDataDevOpsActionResDTOList {
	s.JobName = &v
	return s
}

func (s *StopPipelineIntegratedTaskResponseBodyDataDevOpsActionResDTOList) SetOwner(v string) *StopPipelineIntegratedTaskResponseBodyDataDevOpsActionResDTOList {
	s.Owner = &v
	return s
}

func (s *StopPipelineIntegratedTaskResponseBodyDataDevOpsActionResDTOList) SetStatus(v string) *StopPipelineIntegratedTaskResponseBodyDataDevOpsActionResDTOList {
	s.Status = &v
	return s
}

func (s *StopPipelineIntegratedTaskResponseBodyDataDevOpsActionResDTOList) Validate() error {
	return dara.Validate(s)
}
