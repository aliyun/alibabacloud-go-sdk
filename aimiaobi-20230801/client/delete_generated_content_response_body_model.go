// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGeneratedContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteGeneratedContentResponseBody
	GetCode() *string
	SetData(v bool) *DeleteGeneratedContentResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *DeleteGeneratedContentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteGeneratedContentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGeneratedContentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteGeneratedContentResponseBody
	GetSuccess() *bool
}

type DeleteGeneratedContentResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGeneratedContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGeneratedContentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGeneratedContentResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteGeneratedContentResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteGeneratedContentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteGeneratedContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGeneratedContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGeneratedContentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGeneratedContentResponseBody) SetCode(v string) *DeleteGeneratedContentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGeneratedContentResponseBody) SetData(v bool) *DeleteGeneratedContentResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGeneratedContentResponseBody) SetHttpStatusCode(v int32) *DeleteGeneratedContentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteGeneratedContentResponseBody) SetMessage(v string) *DeleteGeneratedContentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGeneratedContentResponseBody) SetRequestId(v string) *DeleteGeneratedContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGeneratedContentResponseBody) SetSuccess(v bool) *DeleteGeneratedContentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGeneratedContentResponseBody) Validate() error {
	return dara.Validate(s)
}
