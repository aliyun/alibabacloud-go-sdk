// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteModelResponseBody
	GetCode() *string
	SetData(v *DeleteModelResponseBodyData) *DeleteModelResponseBody
	GetData() *DeleteModelResponseBodyData
	SetHttpStatusCode(v int32) *DeleteModelResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteModelResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteModelResponseBody
	GetSuccess() *bool
}

type DeleteModelResponseBody struct {
	// The business status code. The value SUCCESS indicates success.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The identifier of the model that has been accepted for deletion.
	Data *DeleteModelResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code. The value 200 indicates success.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request processing result message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteModelResponseBody) GetData() *DeleteModelResponseBodyData {
	return s.Data
}

func (s *DeleteModelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteModelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteModelResponseBody) SetCode(v string) *DeleteModelResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteModelResponseBody) SetData(v *DeleteModelResponseBodyData) *DeleteModelResponseBody {
	s.Data = v
	return s
}

func (s *DeleteModelResponseBody) SetHttpStatusCode(v int32) *DeleteModelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteModelResponseBody) SetMessage(v string) *DeleteModelResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteModelResponseBody) SetRequestId(v string) *DeleteModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteModelResponseBody) SetSuccess(v bool) *DeleteModelResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteModelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteModelResponseBodyData struct {
	// The model ID.
	//
	// example:
	//
	// model-1
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-1
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DeleteModelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteModelResponseBodyData) GetModelId() *string {
	return s.ModelId
}

func (s *DeleteModelResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteModelResponseBodyData) SetModelId(v string) *DeleteModelResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *DeleteModelResponseBodyData) SetWorkspaceId(v string) *DeleteModelResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteModelResponseBodyData) Validate() error {
	return dara.Validate(s)
}
