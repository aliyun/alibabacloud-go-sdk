// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflinePipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OfflinePipelineResponseBody
	GetCode() *string
	SetData(v *OfflinePipelineResponseBodyData) *OfflinePipelineResponseBody
	GetData() *OfflinePipelineResponseBodyData
	SetHttpStatusCode(v int32) *OfflinePipelineResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *OfflinePipelineResponseBody
	GetMessage() *string
	SetRequestId(v string) *OfflinePipelineResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OfflinePipelineResponseBody
	GetSuccess() *bool
}

type OfflinePipelineResponseBody struct {
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *OfflinePipelineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s OfflinePipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OfflinePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *OfflinePipelineResponseBody) GetCode() *string {
	return s.Code
}

func (s *OfflinePipelineResponseBody) GetData() *OfflinePipelineResponseBodyData {
	return s.Data
}

func (s *OfflinePipelineResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *OfflinePipelineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OfflinePipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OfflinePipelineResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OfflinePipelineResponseBody) SetCode(v string) *OfflinePipelineResponseBody {
	s.Code = &v
	return s
}

func (s *OfflinePipelineResponseBody) SetData(v *OfflinePipelineResponseBodyData) *OfflinePipelineResponseBody {
	s.Data = v
	return s
}

func (s *OfflinePipelineResponseBody) SetHttpStatusCode(v int32) *OfflinePipelineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *OfflinePipelineResponseBody) SetMessage(v string) *OfflinePipelineResponseBody {
	s.Message = &v
	return s
}

func (s *OfflinePipelineResponseBody) SetRequestId(v string) *OfflinePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *OfflinePipelineResponseBody) SetSuccess(v bool) *OfflinePipelineResponseBody {
	s.Success = &v
	return s
}

func (s *OfflinePipelineResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OfflinePipelineResponseBodyData struct {
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

func (s OfflinePipelineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OfflinePipelineResponseBodyData) GoString() string {
	return s.String()
}

func (s *OfflinePipelineResponseBodyData) GetHostMachine() *string {
	return s.HostMachine
}

func (s *OfflinePipelineResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *OfflinePipelineResponseBodyData) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *OfflinePipelineResponseBodyData) GetSubmitId() *int64 {
	return s.SubmitId
}

func (s *OfflinePipelineResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *OfflinePipelineResponseBodyData) SetHostMachine(v string) *OfflinePipelineResponseBodyData {
	s.HostMachine = &v
	return s
}

func (s *OfflinePipelineResponseBodyData) SetNodeId(v string) *OfflinePipelineResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *OfflinePipelineResponseBodyData) SetPipelineId(v int64) *OfflinePipelineResponseBodyData {
	s.PipelineId = &v
	return s
}

func (s *OfflinePipelineResponseBodyData) SetSubmitId(v int64) *OfflinePipelineResponseBodyData {
	s.SubmitId = &v
	return s
}

func (s *OfflinePipelineResponseBodyData) SetVersion(v string) *OfflinePipelineResponseBodyData {
	s.Version = &v
	return s
}

func (s *OfflinePipelineResponseBodyData) Validate() error {
	return dara.Validate(s)
}
