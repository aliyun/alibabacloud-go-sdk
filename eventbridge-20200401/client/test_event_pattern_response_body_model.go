// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestEventPatternResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TestEventPatternResponseBody
	GetCode() *string
	SetData(v *TestEventPatternResponseBodyData) *TestEventPatternResponseBody
	GetData() *TestEventPatternResponseBodyData
	SetMessage(v string) *TestEventPatternResponseBody
	GetMessage() *string
	SetRequestId(v string) *TestEventPatternResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TestEventPatternResponseBody
	GetSuccess() *bool
}

type TestEventPatternResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful. Other values indicate that the request failed. For more information about error codes, see Error codes.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *TestEventPatternResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned if the request failed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 96D7C0AB-DCE5-5E82-96B8-4725E1706BB1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. If the operation is successful, the value true is returned.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TestEventPatternResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TestEventPatternResponseBody) GoString() string {
	return s.String()
}

func (s *TestEventPatternResponseBody) GetCode() *string {
	return s.Code
}

func (s *TestEventPatternResponseBody) GetData() *TestEventPatternResponseBodyData {
	return s.Data
}

func (s *TestEventPatternResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TestEventPatternResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TestEventPatternResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TestEventPatternResponseBody) SetCode(v string) *TestEventPatternResponseBody {
	s.Code = &v
	return s
}

func (s *TestEventPatternResponseBody) SetData(v *TestEventPatternResponseBodyData) *TestEventPatternResponseBody {
	s.Data = v
	return s
}

func (s *TestEventPatternResponseBody) SetMessage(v string) *TestEventPatternResponseBody {
	s.Message = &v
	return s
}

func (s *TestEventPatternResponseBody) SetRequestId(v string) *TestEventPatternResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestEventPatternResponseBody) SetSuccess(v bool) *TestEventPatternResponseBody {
	s.Success = &v
	return s
}

func (s *TestEventPatternResponseBody) Validate() error {
	return dara.Validate(s)
}

type TestEventPatternResponseBodyData struct {
	// The value true indicates that the event pattern matches the provided JSON format. The value false indicates that the event pattern does not match the provided JSON format.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s TestEventPatternResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TestEventPatternResponseBodyData) GoString() string {
	return s.String()
}

func (s *TestEventPatternResponseBodyData) GetResult() *bool {
	return s.Result
}

func (s *TestEventPatternResponseBodyData) SetResult(v bool) *TestEventPatternResponseBodyData {
	s.Result = &v
	return s
}

func (s *TestEventPatternResponseBodyData) Validate() error {
	return dara.Validate(s)
}
