// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindPptArtifactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BindPptArtifactResponseBody
	GetCode() *string
	SetData(v *BindPptArtifactResponseBodyData) *BindPptArtifactResponseBody
	GetData() *BindPptArtifactResponseBodyData
	SetHttpStatusCode(v int32) *BindPptArtifactResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BindPptArtifactResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindPptArtifactResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BindPptArtifactResponseBody
	GetSuccess() *bool
}

type BindPptArtifactResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *BindPptArtifactResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindPptArtifactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindPptArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *BindPptArtifactResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindPptArtifactResponseBody) GetData() *BindPptArtifactResponseBodyData {
	return s.Data
}

func (s *BindPptArtifactResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BindPptArtifactResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindPptArtifactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindPptArtifactResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindPptArtifactResponseBody) SetCode(v string) *BindPptArtifactResponseBody {
	s.Code = &v
	return s
}

func (s *BindPptArtifactResponseBody) SetData(v *BindPptArtifactResponseBodyData) *BindPptArtifactResponseBody {
	s.Data = v
	return s
}

func (s *BindPptArtifactResponseBody) SetHttpStatusCode(v int32) *BindPptArtifactResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BindPptArtifactResponseBody) SetMessage(v string) *BindPptArtifactResponseBody {
	s.Message = &v
	return s
}

func (s *BindPptArtifactResponseBody) SetRequestId(v string) *BindPptArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindPptArtifactResponseBody) SetSuccess(v bool) *BindPptArtifactResponseBody {
	s.Success = &v
	return s
}

func (s *BindPptArtifactResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindPptArtifactResponseBodyData struct {
	// example:
	//
	// 110f8401-e5ba-42db-addb-4f70196000c1
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BindPptArtifactResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BindPptArtifactResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindPptArtifactResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *BindPptArtifactResponseBodyData) SetTaskId(v string) *BindPptArtifactResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *BindPptArtifactResponseBodyData) Validate() error {
	return dara.Validate(s)
}
