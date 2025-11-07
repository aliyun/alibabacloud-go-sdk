// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBlackListStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteBlackListStrategyResponseBody
	GetCode() *string
	SetData(v int32) *DeleteBlackListStrategyResponseBody
	GetData() *int32
	SetMessage(v string) *DeleteBlackListStrategyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteBlackListStrategyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteBlackListStrategyResponseBody
	GetSuccess() *bool
}

type DeleteBlackListStrategyResponseBody struct {
	// HTTP status code.
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
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the response was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteBlackListStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBlackListStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBlackListStrategyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteBlackListStrategyResponseBody) GetData() *int32 {
	return s.Data
}

func (s *DeleteBlackListStrategyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteBlackListStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBlackListStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteBlackListStrategyResponseBody) SetCode(v string) *DeleteBlackListStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBlackListStrategyResponseBody) SetData(v int32) *DeleteBlackListStrategyResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteBlackListStrategyResponseBody) SetMessage(v string) *DeleteBlackListStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBlackListStrategyResponseBody) SetRequestId(v string) *DeleteBlackListStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBlackListStrategyResponseBody) SetSuccess(v bool) *DeleteBlackListStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteBlackListStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
