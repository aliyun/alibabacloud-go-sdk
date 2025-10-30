// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePipelineNodeResponseBody
	GetCode() *string
	SetData(v *CreatePipelineNodeResponseBodyData) *CreatePipelineNodeResponseBody
	GetData() *CreatePipelineNodeResponseBodyData
	SetHttpStatusCode(v int32) *CreatePipelineNodeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreatePipelineNodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePipelineNodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreatePipelineNodeResponseBody
	GetSuccess() *bool
}

type CreatePipelineNodeResponseBody struct {
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreatePipelineNodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePipelineNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineNodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePipelineNodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePipelineNodeResponseBody) GetData() *CreatePipelineNodeResponseBodyData {
	return s.Data
}

func (s *CreatePipelineNodeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreatePipelineNodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePipelineNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePipelineNodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePipelineNodeResponseBody) SetCode(v string) *CreatePipelineNodeResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePipelineNodeResponseBody) SetData(v *CreatePipelineNodeResponseBodyData) *CreatePipelineNodeResponseBody {
	s.Data = v
	return s
}

func (s *CreatePipelineNodeResponseBody) SetHttpStatusCode(v int32) *CreatePipelineNodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreatePipelineNodeResponseBody) SetMessage(v string) *CreatePipelineNodeResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePipelineNodeResponseBody) SetRequestId(v string) *CreatePipelineNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePipelineNodeResponseBody) SetSuccess(v bool) *CreatePipelineNodeResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePipelineNodeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePipelineNodeResponseBodyData struct {
	ErrorCodeList    []*string `json:"ErrorCodeList,omitempty" xml:"ErrorCodeList,omitempty" type:"Repeated"`
	ErrorMessageList []*string `json:"ErrorMessageList,omitempty" xml:"ErrorMessageList,omitempty" type:"Repeated"`
	// example:
	//
	// 33749
	PipelineId *int64 `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s CreatePipelineNodeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineNodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePipelineNodeResponseBodyData) GetErrorCodeList() []*string {
	return s.ErrorCodeList
}

func (s *CreatePipelineNodeResponseBodyData) GetErrorMessageList() []*string {
	return s.ErrorMessageList
}

func (s *CreatePipelineNodeResponseBodyData) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *CreatePipelineNodeResponseBodyData) SetErrorCodeList(v []*string) *CreatePipelineNodeResponseBodyData {
	s.ErrorCodeList = v
	return s
}

func (s *CreatePipelineNodeResponseBodyData) SetErrorMessageList(v []*string) *CreatePipelineNodeResponseBodyData {
	s.ErrorMessageList = v
	return s
}

func (s *CreatePipelineNodeResponseBodyData) SetPipelineId(v int64) *CreatePipelineNodeResponseBodyData {
	s.PipelineId = &v
	return s
}

func (s *CreatePipelineNodeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
