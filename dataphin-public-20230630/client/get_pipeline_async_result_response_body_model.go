// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineAsyncResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPipelineAsyncResultResponseBody
	GetCode() *string
	SetData(v *GetPipelineAsyncResultResponseBodyData) *GetPipelineAsyncResultResponseBody
	GetData() *GetPipelineAsyncResultResponseBodyData
	SetHttpStatusCode(v int32) *GetPipelineAsyncResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPipelineAsyncResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPipelineAsyncResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPipelineAsyncResultResponseBody
	GetSuccess() *bool
}

type GetPipelineAsyncResultResponseBody struct {
	// example:
	//
	// OK
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetPipelineAsyncResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetPipelineAsyncResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineAsyncResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineAsyncResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPipelineAsyncResultResponseBody) GetData() *GetPipelineAsyncResultResponseBodyData {
	return s.Data
}

func (s *GetPipelineAsyncResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPipelineAsyncResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPipelineAsyncResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPipelineAsyncResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPipelineAsyncResultResponseBody) SetCode(v string) *GetPipelineAsyncResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetPipelineAsyncResultResponseBody) SetData(v *GetPipelineAsyncResultResponseBodyData) *GetPipelineAsyncResultResponseBody {
	s.Data = v
	return s
}

func (s *GetPipelineAsyncResultResponseBody) SetHttpStatusCode(v int32) *GetPipelineAsyncResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPipelineAsyncResultResponseBody) SetMessage(v string) *GetPipelineAsyncResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetPipelineAsyncResultResponseBody) SetRequestId(v string) *GetPipelineAsyncResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineAsyncResultResponseBody) SetSuccess(v bool) *GetPipelineAsyncResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetPipelineAsyncResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPipelineAsyncResultResponseBodyData struct {
	// example:
	//
	// 123
	AsyncId *int64 `json:"AsyncId,omitempty" xml:"AsyncId,omitempty"`
	// example:
	//
	// DPN.Pipeline.InnerError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// NullPointException
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// hostName:hostIp
	HostMachine *string `json:"HostMachine,omitempty" xml:"HostMachine,omitempty"`
	// example:
	//
	// n_123
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 123
	PipelineId *int64 `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 123
	SubmitId *int64 `json:"SubmitId,omitempty" xml:"SubmitId,omitempty"`
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetPipelineAsyncResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineAsyncResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPipelineAsyncResultResponseBodyData) GetAsyncId() *int64 {
	return s.AsyncId
}

func (s *GetPipelineAsyncResultResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetPipelineAsyncResultResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetPipelineAsyncResultResponseBodyData) GetHostMachine() *string {
	return s.HostMachine
}

func (s *GetPipelineAsyncResultResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *GetPipelineAsyncResultResponseBodyData) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *GetPipelineAsyncResultResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetPipelineAsyncResultResponseBodyData) GetSubmitId() *int64 {
	return s.SubmitId
}

func (s *GetPipelineAsyncResultResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *GetPipelineAsyncResultResponseBodyData) SetAsyncId(v int64) *GetPipelineAsyncResultResponseBodyData {
	s.AsyncId = &v
	return s
}

func (s *GetPipelineAsyncResultResponseBodyData) SetErrorCode(v string) *GetPipelineAsyncResultResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineAsyncResultResponseBodyData) SetErrorMessage(v string) *GetPipelineAsyncResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineAsyncResultResponseBodyData) SetHostMachine(v string) *GetPipelineAsyncResultResponseBodyData {
	s.HostMachine = &v
	return s
}

func (s *GetPipelineAsyncResultResponseBodyData) SetNodeId(v string) *GetPipelineAsyncResultResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *GetPipelineAsyncResultResponseBodyData) SetPipelineId(v int64) *GetPipelineAsyncResultResponseBodyData {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineAsyncResultResponseBodyData) SetStatus(v string) *GetPipelineAsyncResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetPipelineAsyncResultResponseBodyData) SetSubmitId(v int64) *GetPipelineAsyncResultResponseBodyData {
	s.SubmitId = &v
	return s
}

func (s *GetPipelineAsyncResultResponseBodyData) SetVersion(v string) *GetPipelineAsyncResultResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetPipelineAsyncResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
