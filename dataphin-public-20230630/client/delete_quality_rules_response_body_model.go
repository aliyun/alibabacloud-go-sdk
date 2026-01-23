// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteQualityRulesResponseBody
	GetCode() *string
	SetData(v int32) *DeleteQualityRulesResponseBody
	GetData() *int32
	SetHttpStatusCode(v int32) *DeleteQualityRulesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteQualityRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteQualityRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteQualityRulesResponseBody
	GetSuccess() *bool
}

type DeleteQualityRulesResponseBody struct {
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

func (s DeleteQualityRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQualityRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteQualityRulesResponseBody) GetData() *int32 {
	return s.Data
}

func (s *DeleteQualityRulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteQualityRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteQualityRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQualityRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteQualityRulesResponseBody) SetCode(v string) *DeleteQualityRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteQualityRulesResponseBody) SetData(v int32) *DeleteQualityRulesResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteQualityRulesResponseBody) SetHttpStatusCode(v int32) *DeleteQualityRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteQualityRulesResponseBody) SetMessage(v string) *DeleteQualityRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteQualityRulesResponseBody) SetRequestId(v string) *DeleteQualityRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQualityRulesResponseBody) SetSuccess(v bool) *DeleteQualityRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteQualityRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
