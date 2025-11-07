// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyControlStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyControlStrategyResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyControlStrategyResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyControlStrategyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyControlStrategyResponseBody
	GetSuccess() *bool
}

type ModifyControlStrategyResponseBody struct {
	// Return code
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Information returned by the API call.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the response was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyControlStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyControlStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyControlStrategyResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyControlStrategyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyControlStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyControlStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyControlStrategyResponseBody) SetCode(v string) *ModifyControlStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyControlStrategyResponseBody) SetMessage(v string) *ModifyControlStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyControlStrategyResponseBody) SetRequestId(v string) *ModifyControlStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyControlStrategyResponseBody) SetSuccess(v bool) *ModifyControlStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyControlStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
