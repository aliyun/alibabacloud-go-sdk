// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaterialByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMaterialByIdResponseBody
	GetCode() *string
	SetData(v bool) *DeleteMaterialByIdResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *DeleteMaterialByIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteMaterialByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMaterialByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMaterialByIdResponseBody
	GetSuccess() *bool
}

type DeleteMaterialByIdResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// false
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 数据不存在
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMaterialByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaterialByIdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMaterialByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMaterialByIdResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteMaterialByIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteMaterialByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMaterialByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMaterialByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMaterialByIdResponseBody) SetCode(v string) *DeleteMaterialByIdResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMaterialByIdResponseBody) SetData(v bool) *DeleteMaterialByIdResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteMaterialByIdResponseBody) SetHttpStatusCode(v int32) *DeleteMaterialByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteMaterialByIdResponseBody) SetMessage(v string) *DeleteMaterialByIdResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMaterialByIdResponseBody) SetRequestId(v string) *DeleteMaterialByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMaterialByIdResponseBody) SetSuccess(v bool) *DeleteMaterialByIdResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMaterialByIdResponseBody) Validate() error {
	return dara.Validate(s)
}
