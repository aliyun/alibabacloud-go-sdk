// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestEventSourceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TestEventSourceConfigResponseBody
	GetCode() *string
	SetData(v []*TestEventSourceConfigResponseBodyData) *TestEventSourceConfigResponseBody
	GetData() []*TestEventSourceConfigResponseBodyData
	SetMessage(v string) *TestEventSourceConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *TestEventSourceConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TestEventSourceConfigResponseBody
	GetSuccess() *bool
}

type TestEventSourceConfigResponseBody struct {
	// The response code. Valid values:
	//
	// 	- Success: The request was successful.
	//
	// 	- Other codes indicate that the request failed. For information about error codes, see Error codes.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The update result.
	Data []*TestEventSourceConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FF942675-F937-549C-A942-EB94FFE28DD3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. If the operation was successful, the value true is returned.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TestEventSourceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TestEventSourceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *TestEventSourceConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *TestEventSourceConfigResponseBody) GetData() []*TestEventSourceConfigResponseBodyData {
	return s.Data
}

func (s *TestEventSourceConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TestEventSourceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TestEventSourceConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TestEventSourceConfigResponseBody) SetCode(v string) *TestEventSourceConfigResponseBody {
	s.Code = &v
	return s
}

func (s *TestEventSourceConfigResponseBody) SetData(v []*TestEventSourceConfigResponseBodyData) *TestEventSourceConfigResponseBody {
	s.Data = v
	return s
}

func (s *TestEventSourceConfigResponseBody) SetMessage(v string) *TestEventSourceConfigResponseBody {
	s.Message = &v
	return s
}

func (s *TestEventSourceConfigResponseBody) SetRequestId(v string) *TestEventSourceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestEventSourceConfigResponseBody) SetSuccess(v bool) *TestEventSourceConfigResponseBody {
	s.Success = &v
	return s
}

func (s *TestEventSourceConfigResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TestEventSourceConfigResponseBodyData struct {
	// The name of the check item.
	//
	// example:
	//
	// CHECK_CONNECTION
	CheckItem *string `json:"CheckItem,omitempty" xml:"CheckItem,omitempty"`
	// The error message.
	//
	// example:
	//
	// Connection established successfully.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Indicates whether the check item is executed.
	//
	// example:
	//
	// true
	IsSucceed *string `json:"IsSucceed,omitempty" xml:"IsSucceed,omitempty"`
}

func (s TestEventSourceConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TestEventSourceConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *TestEventSourceConfigResponseBodyData) GetCheckItem() *string {
	return s.CheckItem
}

func (s *TestEventSourceConfigResponseBodyData) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *TestEventSourceConfigResponseBodyData) GetIsSucceed() *string {
	return s.IsSucceed
}

func (s *TestEventSourceConfigResponseBodyData) SetCheckItem(v string) *TestEventSourceConfigResponseBodyData {
	s.CheckItem = &v
	return s
}

func (s *TestEventSourceConfigResponseBodyData) SetErrorMsg(v string) *TestEventSourceConfigResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *TestEventSourceConfigResponseBodyData) SetIsSucceed(v string) *TestEventSourceConfigResponseBodyData {
	s.IsSucceed = &v
	return s
}

func (s *TestEventSourceConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
