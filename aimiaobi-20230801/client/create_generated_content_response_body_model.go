// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGeneratedContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateGeneratedContentResponseBody
	GetCode() *string
	SetData(v int64) *CreateGeneratedContentResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *CreateGeneratedContentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateGeneratedContentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateGeneratedContentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateGeneratedContentResponseBody
	GetSuccess() *bool
}

type CreateGeneratedContentResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 42
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s CreateGeneratedContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGeneratedContentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGeneratedContentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateGeneratedContentResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateGeneratedContentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateGeneratedContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGeneratedContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGeneratedContentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateGeneratedContentResponseBody) SetCode(v string) *CreateGeneratedContentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGeneratedContentResponseBody) SetData(v int64) *CreateGeneratedContentResponseBody {
	s.Data = &v
	return s
}

func (s *CreateGeneratedContentResponseBody) SetHttpStatusCode(v int32) *CreateGeneratedContentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateGeneratedContentResponseBody) SetMessage(v string) *CreateGeneratedContentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGeneratedContentResponseBody) SetRequestId(v string) *CreateGeneratedContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGeneratedContentResponseBody) SetSuccess(v bool) *CreateGeneratedContentResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGeneratedContentResponseBody) Validate() error {
	return dara.Validate(s)
}
