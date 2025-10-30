// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdatePipelineResponseBody
	GetCode() *string
	SetData(v *UpdatePipelineResponseBodyData) *UpdatePipelineResponseBody
	GetData() *UpdatePipelineResponseBodyData
	SetHttpStatusCode(v int32) *UpdatePipelineResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdatePipelineResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdatePipelineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdatePipelineResponseBody
	GetSuccess() *bool
}

type UpdatePipelineResponseBody struct {
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdatePipelineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s UpdatePipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePipelineResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdatePipelineResponseBody) GetData() *UpdatePipelineResponseBodyData {
	return s.Data
}

func (s *UpdatePipelineResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdatePipelineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePipelineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePipelineResponseBody) SetCode(v string) *UpdatePipelineResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePipelineResponseBody) SetData(v *UpdatePipelineResponseBodyData) *UpdatePipelineResponseBody {
	s.Data = v
	return s
}

func (s *UpdatePipelineResponseBody) SetHttpStatusCode(v int32) *UpdatePipelineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdatePipelineResponseBody) SetMessage(v string) *UpdatePipelineResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePipelineResponseBody) SetRequestId(v string) *UpdatePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePipelineResponseBody) SetSuccess(v bool) *UpdatePipelineResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePipelineResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePipelineResponseBodyData struct {
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

func (s UpdatePipelineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdatePipelineResponseBodyData) GetHostMachine() *string {
	return s.HostMachine
}

func (s *UpdatePipelineResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *UpdatePipelineResponseBodyData) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *UpdatePipelineResponseBodyData) GetSubmitId() *int64 {
	return s.SubmitId
}

func (s *UpdatePipelineResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *UpdatePipelineResponseBodyData) SetHostMachine(v string) *UpdatePipelineResponseBodyData {
	s.HostMachine = &v
	return s
}

func (s *UpdatePipelineResponseBodyData) SetNodeId(v string) *UpdatePipelineResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *UpdatePipelineResponseBodyData) SetPipelineId(v int64) *UpdatePipelineResponseBodyData {
	s.PipelineId = &v
	return s
}

func (s *UpdatePipelineResponseBodyData) SetSubmitId(v int64) *UpdatePipelineResponseBodyData {
	s.SubmitId = &v
	return s
}

func (s *UpdatePipelineResponseBodyData) SetVersion(v string) *UpdatePipelineResponseBodyData {
	s.Version = &v
	return s
}

func (s *UpdatePipelineResponseBodyData) Validate() error {
	return dara.Validate(s)
}
