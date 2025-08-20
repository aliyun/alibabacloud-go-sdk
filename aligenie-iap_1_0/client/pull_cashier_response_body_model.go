// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPullCashierResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRetCode(v int32) *PullCashierResponseBody
	GetRetCode() *int32
	SetRetMsg(v string) *PullCashierResponseBody
	GetRetMsg() *string
	SetRetValue(v bool) *PullCashierResponseBody
	GetRetValue() *bool
}

type PullCashierResponseBody struct {
	// example:
	//
	// 0
	RetCode *int32  `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	RetMsg  *string `json:"RetMsg,omitempty" xml:"RetMsg,omitempty"`
	// example:
	//
	// true
	RetValue *bool `json:"RetValue,omitempty" xml:"RetValue,omitempty"`
}

func (s PullCashierResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PullCashierResponseBody) GoString() string {
	return s.String()
}

func (s *PullCashierResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *PullCashierResponseBody) GetRetMsg() *string {
	return s.RetMsg
}

func (s *PullCashierResponseBody) GetRetValue() *bool {
	return s.RetValue
}

func (s *PullCashierResponseBody) SetRetCode(v int32) *PullCashierResponseBody {
	s.RetCode = &v
	return s
}

func (s *PullCashierResponseBody) SetRetMsg(v string) *PullCashierResponseBody {
	s.RetMsg = &v
	return s
}

func (s *PullCashierResponseBody) SetRetValue(v bool) *PullCashierResponseBody {
	s.RetValue = &v
	return s
}

func (s *PullCashierResponseBody) Validate() error {
	return dara.Validate(s)
}
