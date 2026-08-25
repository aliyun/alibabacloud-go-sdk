// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchDeleteModelsResponseBody
	GetCode() *string
	SetData(v *BatchDeleteModelsResponseBodyData) *BatchDeleteModelsResponseBody
	GetData() *BatchDeleteModelsResponseBodyData
	SetHttpStatusCode(v int32) *BatchDeleteModelsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BatchDeleteModelsResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchDeleteModelsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchDeleteModelsResponseBody
	GetSuccess() *bool
}

type BatchDeleteModelsResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *BatchDeleteModelsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// request-1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchDeleteModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteModelsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteModelsResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchDeleteModelsResponseBody) GetData() *BatchDeleteModelsResponseBodyData {
	return s.Data
}

func (s *BatchDeleteModelsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BatchDeleteModelsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchDeleteModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteModelsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchDeleteModelsResponseBody) SetCode(v string) *BatchDeleteModelsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchDeleteModelsResponseBody) SetData(v *BatchDeleteModelsResponseBodyData) *BatchDeleteModelsResponseBody {
	s.Data = v
	return s
}

func (s *BatchDeleteModelsResponseBody) SetHttpStatusCode(v int32) *BatchDeleteModelsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BatchDeleteModelsResponseBody) SetMessage(v string) *BatchDeleteModelsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchDeleteModelsResponseBody) SetRequestId(v string) *BatchDeleteModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteModelsResponseBody) SetSuccess(v bool) *BatchDeleteModelsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchDeleteModelsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchDeleteModelsResponseBodyData struct {
	ModelIds []*string `json:"modelIds,omitempty" xml:"modelIds,omitempty" type:"Repeated"`
	// example:
	//
	// ws-1
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s BatchDeleteModelsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteModelsResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchDeleteModelsResponseBodyData) GetModelIds() []*string {
	return s.ModelIds
}

func (s *BatchDeleteModelsResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *BatchDeleteModelsResponseBodyData) SetModelIds(v []*string) *BatchDeleteModelsResponseBodyData {
	s.ModelIds = v
	return s
}

func (s *BatchDeleteModelsResponseBodyData) SetWorkspaceId(v string) *BatchDeleteModelsResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *BatchDeleteModelsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
