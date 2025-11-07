// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteControlStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteControlStrategyResponseBody
	GetCode() *string
	SetData(v int32) *DeleteControlStrategyResponseBody
	GetData() *int32
	SetMessage(v string) *DeleteControlStrategyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteControlStrategyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteControlStrategyResponseBody
	GetSuccess() *bool
}

type DeleteControlStrategyResponseBody struct {
	// Return code, **200*	- indicates successful API response.
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
	// D6C1237B-D34B-5126-93AC-36A4B9E819C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the response was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteControlStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteControlStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteControlStrategyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteControlStrategyResponseBody) GetData() *int32 {
	return s.Data
}

func (s *DeleteControlStrategyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteControlStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteControlStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteControlStrategyResponseBody) SetCode(v string) *DeleteControlStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteControlStrategyResponseBody) SetData(v int32) *DeleteControlStrategyResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteControlStrategyResponseBody) SetMessage(v string) *DeleteControlStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteControlStrategyResponseBody) SetRequestId(v string) *DeleteControlStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteControlStrategyResponseBody) SetSuccess(v bool) *DeleteControlStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteControlStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
