// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteQualityTemplatesResponseBody
	GetCode() *string
	SetData(v int32) *DeleteQualityTemplatesResponseBody
	GetData() *int32
	SetHttpStatusCode(v int32) *DeleteQualityTemplatesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteQualityTemplatesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteQualityTemplatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteQualityTemplatesResponseBody
	GetSuccess() *bool
}

type DeleteQualityTemplatesResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s DeleteQualityTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQualityTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteQualityTemplatesResponseBody) GetData() *int32 {
	return s.Data
}

func (s *DeleteQualityTemplatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteQualityTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteQualityTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQualityTemplatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteQualityTemplatesResponseBody) SetCode(v string) *DeleteQualityTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteQualityTemplatesResponseBody) SetData(v int32) *DeleteQualityTemplatesResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteQualityTemplatesResponseBody) SetHttpStatusCode(v int32) *DeleteQualityTemplatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteQualityTemplatesResponseBody) SetMessage(v string) *DeleteQualityTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteQualityTemplatesResponseBody) SetRequestId(v string) *DeleteQualityTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQualityTemplatesResponseBody) SetSuccess(v bool) *DeleteQualityTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteQualityTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}
