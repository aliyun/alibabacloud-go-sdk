// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomizeFlowStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCustomizeFlowStrategyResponseBody
	GetCode() *string
	SetData(v int32) *DeleteCustomizeFlowStrategyResponseBody
	GetData() *int32
	SetHttpStatusCode(v int32) *DeleteCustomizeFlowStrategyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteCustomizeFlowStrategyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCustomizeFlowStrategyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCustomizeFlowStrategyResponseBody
	GetSuccess() *bool
}

type DeleteCustomizeFlowStrategyResponseBody struct {
	// HTTP status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data.
	//
	// example:
	//
	// -
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
	// 5A6229C0-E156-48E4-B6EC-0F528BDF60D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the response was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCustomizeFlowStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomizeFlowStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomizeFlowStrategyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCustomizeFlowStrategyResponseBody) GetData() *int32 {
	return s.Data
}

func (s *DeleteCustomizeFlowStrategyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteCustomizeFlowStrategyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCustomizeFlowStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomizeFlowStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCustomizeFlowStrategyResponseBody) SetCode(v string) *DeleteCustomizeFlowStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCustomizeFlowStrategyResponseBody) SetData(v int32) *DeleteCustomizeFlowStrategyResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCustomizeFlowStrategyResponseBody) SetHttpStatusCode(v int32) *DeleteCustomizeFlowStrategyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteCustomizeFlowStrategyResponseBody) SetMessage(v string) *DeleteCustomizeFlowStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCustomizeFlowStrategyResponseBody) SetRequestId(v string) *DeleteCustomizeFlowStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomizeFlowStrategyResponseBody) SetSuccess(v bool) *DeleteCustomizeFlowStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCustomizeFlowStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
