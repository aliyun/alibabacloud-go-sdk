// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePipelineResponseBody
	GetCode() *string
	SetData(v *CreatePipelineResponseBodyData) *CreatePipelineResponseBody
	GetData() *CreatePipelineResponseBodyData
	SetHttpStatusCode(v int32) *CreatePipelineResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreatePipelineResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePipelineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreatePipelineResponseBody
	GetSuccess() *bool
}

type CreatePipelineResponseBody struct {
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreatePipelineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
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
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePipelineResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePipelineResponseBody) GetData() *CreatePipelineResponseBodyData {
	return s.Data
}

func (s *CreatePipelineResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreatePipelineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePipelineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePipelineResponseBody) SetCode(v string) *CreatePipelineResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePipelineResponseBody) SetData(v *CreatePipelineResponseBodyData) *CreatePipelineResponseBody {
	s.Data = v
	return s
}

func (s *CreatePipelineResponseBody) SetHttpStatusCode(v int32) *CreatePipelineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreatePipelineResponseBody) SetMessage(v string) *CreatePipelineResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePipelineResponseBody) SetRequestId(v string) *CreatePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePipelineResponseBody) SetSuccess(v bool) *CreatePipelineResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePipelineResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePipelineResponseBodyData struct {
	// example:
	//
	// hostName:hostIp
	HostMachine *string `json:"HostMachine,omitempty" xml:"HostMachine,omitempty"`
	// example:
	//
	// 123
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 123
	PipelineId *int64 `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// example:
	//
	// 123
	SubmitId *int64 `json:"SubmitId,omitempty" xml:"SubmitId,omitempty"`
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreatePipelineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePipelineResponseBodyData) GetHostMachine() *string {
	return s.HostMachine
}

func (s *CreatePipelineResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *CreatePipelineResponseBodyData) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *CreatePipelineResponseBodyData) GetSubmitId() *int64 {
	return s.SubmitId
}

func (s *CreatePipelineResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *CreatePipelineResponseBodyData) SetHostMachine(v string) *CreatePipelineResponseBodyData {
	s.HostMachine = &v
	return s
}

func (s *CreatePipelineResponseBodyData) SetNodeId(v string) *CreatePipelineResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *CreatePipelineResponseBodyData) SetPipelineId(v int64) *CreatePipelineResponseBodyData {
	s.PipelineId = &v
	return s
}

func (s *CreatePipelineResponseBodyData) SetSubmitId(v int64) *CreatePipelineResponseBodyData {
	s.SubmitId = &v
	return s
}

func (s *CreatePipelineResponseBodyData) SetVersion(v string) *CreatePipelineResponseBodyData {
	s.Version = &v
	return s
}

func (s *CreatePipelineResponseBodyData) Validate() error {
	return dara.Validate(s)
}
