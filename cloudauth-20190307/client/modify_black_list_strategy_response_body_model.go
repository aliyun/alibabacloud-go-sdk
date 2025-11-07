// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBlackListStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyBlackListStrategyResponseBody
	GetCode() *string
	SetData(v int32) *ModifyBlackListStrategyResponseBody
	GetData() *int32
	SetMessage(v string) *ModifyBlackListStrategyResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyBlackListStrategyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyBlackListStrategyResponseBody
	GetSuccess() *bool
}

type ModifyBlackListStrategyResponseBody struct {
	// Return code: 200 for success, others for failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data.
	//
	// example:
	//
	// {\\"StatusCode\\": -1}
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// Return message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the response was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyBlackListStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBlackListStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBlackListStrategyResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyBlackListStrategyResponseBody) GetData() *int32 {
	return s.Data
}

func (s *ModifyBlackListStrategyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyBlackListStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBlackListStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyBlackListStrategyResponseBody) SetCode(v string) *ModifyBlackListStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyBlackListStrategyResponseBody) SetData(v int32) *ModifyBlackListStrategyResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyBlackListStrategyResponseBody) SetMessage(v string) *ModifyBlackListStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyBlackListStrategyResponseBody) SetRequestId(v string) *ModifyBlackListStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBlackListStrategyResponseBody) SetSuccess(v bool) *ModifyBlackListStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyBlackListStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
