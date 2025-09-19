// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAttackCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAttackCountResponseBody
	GetCode() *string
	SetCount(v int32) *QueryAttackCountResponseBody
	GetCount() *int32
	SetData(v []*QueryAttackCountResponseBodyData) *QueryAttackCountResponseBody
	GetData() []*QueryAttackCountResponseBodyData
	SetMessage(v string) *QueryAttackCountResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAttackCountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAttackCountResponseBody
	GetSuccess() *bool
}

type QueryAttackCountResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 0
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// An array that consists of the numbers of alert events in different attack phases.
	Data []*QueryAttackCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D4BE7D77-5B02-5126-A684-A73F6CD3XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether exceptions are handled. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAttackCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAttackCountResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAttackCountResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAttackCountResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *QueryAttackCountResponseBody) GetData() []*QueryAttackCountResponseBodyData {
	return s.Data
}

func (s *QueryAttackCountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAttackCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAttackCountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAttackCountResponseBody) SetCode(v string) *QueryAttackCountResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAttackCountResponseBody) SetCount(v int32) *QueryAttackCountResponseBody {
	s.Count = &v
	return s
}

func (s *QueryAttackCountResponseBody) SetData(v []*QueryAttackCountResponseBodyData) *QueryAttackCountResponseBody {
	s.Data = v
	return s
}

func (s *QueryAttackCountResponseBody) SetMessage(v string) *QueryAttackCountResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAttackCountResponseBody) SetRequestId(v string) *QueryAttackCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAttackCountResponseBody) SetSuccess(v bool) *QueryAttackCountResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAttackCountResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryAttackCountResponseBodyData struct {
	// The number of times that the alert is triggered.
	//
	// example:
	//
	// 28
	EventCount *int32 `json:"EventCount,omitempty" xml:"EventCount,omitempty"`
	// The stage ID of the ATT\\&CK attack.
	//
	// example:
	//
	// TA0043
	TacticId *string `json:"TacticId,omitempty" xml:"TacticId,omitempty"`
	// The type of stage of the ATT\\&CK attack.
	//
	// example:
	//
	// Data collection
	TacticType *string `json:"TacticType,omitempty" xml:"TacticType,omitempty"`
}

func (s QueryAttackCountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAttackCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAttackCountResponseBodyData) GetEventCount() *int32 {
	return s.EventCount
}

func (s *QueryAttackCountResponseBodyData) GetTacticId() *string {
	return s.TacticId
}

func (s *QueryAttackCountResponseBodyData) GetTacticType() *string {
	return s.TacticType
}

func (s *QueryAttackCountResponseBodyData) SetEventCount(v int32) *QueryAttackCountResponseBodyData {
	s.EventCount = &v
	return s
}

func (s *QueryAttackCountResponseBodyData) SetTacticId(v string) *QueryAttackCountResponseBodyData {
	s.TacticId = &v
	return s
}

func (s *QueryAttackCountResponseBodyData) SetTacticType(v string) *QueryAttackCountResponseBodyData {
	s.TacticType = &v
	return s
}

func (s *QueryAttackCountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
