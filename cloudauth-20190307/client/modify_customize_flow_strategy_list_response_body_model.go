// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomizeFlowStrategyListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyCustomizeFlowStrategyListResponseBody
	GetCode() *string
	SetData(v int32) *ModifyCustomizeFlowStrategyListResponseBody
	GetData() *int32
	SetHttpStatusCode(v int32) *ModifyCustomizeFlowStrategyListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyCustomizeFlowStrategyListResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyCustomizeFlowStrategyListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyCustomizeFlowStrategyListResponseBody
	GetSuccess() *bool
}

type ModifyCustomizeFlowStrategyListResponseBody struct {
	// Return code, **200*	- indicates the interface responded successfully.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Result data.
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
	// ID of the request
	//
	// example:
	//
	// 8FC3D6AC-9FED-4311-8DA7-C4BF47D9F260
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the response was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyCustomizeFlowStrategyListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomizeFlowStrategyListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCustomizeFlowStrategyListResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyCustomizeFlowStrategyListResponseBody) GetData() *int32 {
	return s.Data
}

func (s *ModifyCustomizeFlowStrategyListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyCustomizeFlowStrategyListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyCustomizeFlowStrategyListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCustomizeFlowStrategyListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyCustomizeFlowStrategyListResponseBody) SetCode(v string) *ModifyCustomizeFlowStrategyListResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyCustomizeFlowStrategyListResponseBody) SetData(v int32) *ModifyCustomizeFlowStrategyListResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyCustomizeFlowStrategyListResponseBody) SetHttpStatusCode(v int32) *ModifyCustomizeFlowStrategyListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyCustomizeFlowStrategyListResponseBody) SetMessage(v string) *ModifyCustomizeFlowStrategyListResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyCustomizeFlowStrategyListResponseBody) SetRequestId(v string) *ModifyCustomizeFlowStrategyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCustomizeFlowStrategyListResponseBody) SetSuccess(v bool) *ModifyCustomizeFlowStrategyListResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyCustomizeFlowStrategyListResponseBody) Validate() error {
	return dara.Validate(s)
}
