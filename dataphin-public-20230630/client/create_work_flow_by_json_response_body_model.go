// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkFlowByJsonResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateWorkFlowByJsonResponseBody
	GetCode() *string
	SetData(v *CreateWorkFlowByJsonResponseBodyData) *CreateWorkFlowByJsonResponseBody
	GetData() *CreateWorkFlowByJsonResponseBodyData
	SetHttpStatusCode(v int32) *CreateWorkFlowByJsonResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateWorkFlowByJsonResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateWorkFlowByJsonResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateWorkFlowByJsonResponseBody
	GetSuccess() *bool
}

type CreateWorkFlowByJsonResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The creation result of the workflow node.
	Data *CreateWorkFlowByJsonResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned by the backend.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateWorkFlowByJsonResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkFlowByJsonResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkFlowByJsonResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateWorkFlowByJsonResponseBody) GetData() *CreateWorkFlowByJsonResponseBodyData {
	return s.Data
}

func (s *CreateWorkFlowByJsonResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateWorkFlowByJsonResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateWorkFlowByJsonResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkFlowByJsonResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateWorkFlowByJsonResponseBody) SetCode(v string) *CreateWorkFlowByJsonResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWorkFlowByJsonResponseBody) SetData(v *CreateWorkFlowByJsonResponseBodyData) *CreateWorkFlowByJsonResponseBody {
	s.Data = v
	return s
}

func (s *CreateWorkFlowByJsonResponseBody) SetHttpStatusCode(v int32) *CreateWorkFlowByJsonResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateWorkFlowByJsonResponseBody) SetMessage(v string) *CreateWorkFlowByJsonResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWorkFlowByJsonResponseBody) SetRequestId(v string) *CreateWorkFlowByJsonResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkFlowByJsonResponseBody) SetSuccess(v bool) *CreateWorkFlowByJsonResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWorkFlowByJsonResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkFlowByJsonResponseBodyData struct {
	// The name and IP address of the server-side execution machine.
	//
	// example:
	//
	// hostName:hostIp
	HostMachine *string `json:"HostMachine,omitempty" xml:"HostMachine,omitempty"`
	// The scheduling node ID of the workflow node created.
	//
	// example:
	//
	// n_123
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The workflow ID generated after successful creation.
	//
	// example:
	//
	// 123
	PipelineId *int64 `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The submit ID of the pending publish list generated after a successful submission. Use this ID to publish through the publish domain. You can ignore this value for BASIC projects.
	//
	// example:
	//
	// 7891
	SubmitId *int64 `json:"SubmitId,omitempty" xml:"SubmitId,omitempty"`
	// The pending publish version number generated after the workflow is submitted. You can ignore this value for BASIC projects.
	//
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateWorkFlowByJsonResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkFlowByJsonResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWorkFlowByJsonResponseBodyData) GetHostMachine() *string {
	return s.HostMachine
}

func (s *CreateWorkFlowByJsonResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateWorkFlowByJsonResponseBodyData) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *CreateWorkFlowByJsonResponseBodyData) GetSubmitId() *int64 {
	return s.SubmitId
}

func (s *CreateWorkFlowByJsonResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *CreateWorkFlowByJsonResponseBodyData) SetHostMachine(v string) *CreateWorkFlowByJsonResponseBodyData {
	s.HostMachine = &v
	return s
}

func (s *CreateWorkFlowByJsonResponseBodyData) SetNodeId(v string) *CreateWorkFlowByJsonResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *CreateWorkFlowByJsonResponseBodyData) SetPipelineId(v int64) *CreateWorkFlowByJsonResponseBodyData {
	s.PipelineId = &v
	return s
}

func (s *CreateWorkFlowByJsonResponseBodyData) SetSubmitId(v int64) *CreateWorkFlowByJsonResponseBodyData {
	s.SubmitId = &v
	return s
}

func (s *CreateWorkFlowByJsonResponseBodyData) SetVersion(v string) *CreateWorkFlowByJsonResponseBodyData {
	s.Version = &v
	return s
}

func (s *CreateWorkFlowByJsonResponseBodyData) Validate() error {
	return dara.Validate(s)
}
