// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAgentDataSemanticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveAgentDataSemanticsResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *SaveAgentDataSemanticsResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *SaveAgentDataSemanticsResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveAgentDataSemanticsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveAgentDataSemanticsResponseBody
	GetSuccess() *bool
}

type SaveAgentDataSemanticsResponseBody struct {
	// The response code of the operation.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The update result. If the save is successful, an empty object is returned with no additional business fields. If none of the four knowledge categories are specified, the target state is all four categories empty: if a non-empty current version exists, an all-empty version is published. If the current version is already all empty or no current version exists, the operation succeeds idempotently and the current round of pending generation results is finalized.
	//
	// example:
	//
	// {}
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The response message. If the call fails, an error message is returned.
	//
	// example:
	//
	// Invalid data semantics request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier that Alibaba Cloud generates for the request.
	//
	// example:
	//
	// 5DAF96FB-A4DF-548C-B8A1-F2A8D2F4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveAgentDataSemanticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveAgentDataSemanticsResponseBody) GoString() string {
	return s.String()
}

func (s *SaveAgentDataSemanticsResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveAgentDataSemanticsResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *SaveAgentDataSemanticsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveAgentDataSemanticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveAgentDataSemanticsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveAgentDataSemanticsResponseBody) SetCode(v string) *SaveAgentDataSemanticsResponseBody {
	s.Code = &v
	return s
}

func (s *SaveAgentDataSemanticsResponseBody) SetData(v map[string]interface{}) *SaveAgentDataSemanticsResponseBody {
	s.Data = v
	return s
}

func (s *SaveAgentDataSemanticsResponseBody) SetMessage(v string) *SaveAgentDataSemanticsResponseBody {
	s.Message = &v
	return s
}

func (s *SaveAgentDataSemanticsResponseBody) SetRequestId(v string) *SaveAgentDataSemanticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveAgentDataSemanticsResponseBody) SetSuccess(v bool) *SaveAgentDataSemanticsResponseBody {
	s.Success = &v
	return s
}

func (s *SaveAgentDataSemanticsResponseBody) Validate() error {
	return dara.Validate(s)
}
