// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateExportWordTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateExportWordTaskResponseBody
	GetCode() *string
	SetData(v *GenerateExportWordTaskResponseBodyData) *GenerateExportWordTaskResponseBody
	GetData() *GenerateExportWordTaskResponseBodyData
	SetHttpStatusCode(v int32) *GenerateExportWordTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GenerateExportWordTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateExportWordTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenerateExportWordTaskResponseBody
	GetSuccess() *bool
}

type GenerateExportWordTaskResponseBody struct {
	// example:
	//
	// NoData
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GenerateExportWordTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateExportWordTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateExportWordTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateExportWordTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateExportWordTaskResponseBody) GetData() *GenerateExportWordTaskResponseBodyData {
	return s.Data
}

func (s *GenerateExportWordTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GenerateExportWordTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateExportWordTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateExportWordTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateExportWordTaskResponseBody) SetCode(v string) *GenerateExportWordTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateExportWordTaskResponseBody) SetData(v *GenerateExportWordTaskResponseBodyData) *GenerateExportWordTaskResponseBody {
	s.Data = v
	return s
}

func (s *GenerateExportWordTaskResponseBody) SetHttpStatusCode(v int32) *GenerateExportWordTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateExportWordTaskResponseBody) SetMessage(v string) *GenerateExportWordTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateExportWordTaskResponseBody) SetRequestId(v string) *GenerateExportWordTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateExportWordTaskResponseBody) SetSuccess(v bool) *GenerateExportWordTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateExportWordTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateExportWordTaskResponseBodyData struct {
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GenerateExportWordTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateExportWordTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateExportWordTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GenerateExportWordTaskResponseBodyData) SetTaskId(v string) *GenerateExportWordTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GenerateExportWordTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
