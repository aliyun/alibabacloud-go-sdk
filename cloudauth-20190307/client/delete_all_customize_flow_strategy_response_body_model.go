// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAllCustomizeFlowStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAllCustomizeFlowStrategyResponseBody
	GetCode() *string
	SetData(v int32) *DeleteAllCustomizeFlowStrategyResponseBody
	GetData() *int32
	SetHttpStatusCode(v int32) *DeleteAllCustomizeFlowStrategyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteAllCustomizeFlowStrategyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAllCustomizeFlowStrategyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAllCustomizeFlowStrategyResponseBody
	GetSuccess() *bool
}

type DeleteAllCustomizeFlowStrategyResponseBody struct {
	// Return code: 200 indicates success, others indicate failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data.
	//
	// example:
	//
	// 1
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of this request.
	//
	// example:
	//
	// 8FC3D6AC-9FED-4311-8DA7-C4BF47D9F260
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the response was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAllCustomizeFlowStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAllCustomizeFlowStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAllCustomizeFlowStrategyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAllCustomizeFlowStrategyResponseBody) GetData() *int32 {
	return s.Data
}

func (s *DeleteAllCustomizeFlowStrategyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteAllCustomizeFlowStrategyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAllCustomizeFlowStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAllCustomizeFlowStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAllCustomizeFlowStrategyResponseBody) SetCode(v string) *DeleteAllCustomizeFlowStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAllCustomizeFlowStrategyResponseBody) SetData(v int32) *DeleteAllCustomizeFlowStrategyResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAllCustomizeFlowStrategyResponseBody) SetHttpStatusCode(v int32) *DeleteAllCustomizeFlowStrategyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteAllCustomizeFlowStrategyResponseBody) SetMessage(v string) *DeleteAllCustomizeFlowStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAllCustomizeFlowStrategyResponseBody) SetRequestId(v string) *DeleteAllCustomizeFlowStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAllCustomizeFlowStrategyResponseBody) SetSuccess(v bool) *DeleteAllCustomizeFlowStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAllCustomizeFlowStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
