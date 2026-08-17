// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeAgentWithTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpgradeAgentWithTypeResponseBody
	GetCode() *string
	SetData(v *UpgradeAgentWithTypeResponseBodyData) *UpgradeAgentWithTypeResponseBody
	GetData() *UpgradeAgentWithTypeResponseBodyData
	SetMessage(v string) *UpgradeAgentWithTypeResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpgradeAgentWithTypeResponseBody
	GetRequestId() *string
}

type UpgradeAgentWithTypeResponseBody struct {
	// The status code.
	//
	// - `code == Success` indicates that the authorization is successful.
	//
	// - Other status codes indicate that the authorization failed. Check the `message` field for the detailed fault information.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *UpgradeAgentWithTypeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message.
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the request error information.
	//
	// example:
	//
	// “”
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, which can be used for end-to-end diagnostics.
	//
	// example:
	//
	// B149FD9C-ED5C-5765-B3AD-05AA4A4D64D7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpgradeAgentWithTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAgentWithTypeResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeAgentWithTypeResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpgradeAgentWithTypeResponseBody) GetData() *UpgradeAgentWithTypeResponseBodyData {
	return s.Data
}

func (s *UpgradeAgentWithTypeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradeAgentWithTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeAgentWithTypeResponseBody) SetCode(v string) *UpgradeAgentWithTypeResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeAgentWithTypeResponseBody) SetData(v *UpgradeAgentWithTypeResponseBodyData) *UpgradeAgentWithTypeResponseBody {
	s.Data = v
	return s
}

func (s *UpgradeAgentWithTypeResponseBody) SetMessage(v string) *UpgradeAgentWithTypeResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeAgentWithTypeResponseBody) SetRequestId(v string) *UpgradeAgentWithTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeAgentWithTypeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpgradeAgentWithTypeResponseBodyData struct {
	// The task ID.
	//
	// example:
	//
	// 391f5aeba2054f66b3aaef0136142fe2
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UpgradeAgentWithTypeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAgentWithTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpgradeAgentWithTypeResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UpgradeAgentWithTypeResponseBodyData) SetTaskId(v string) *UpgradeAgentWithTypeResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpgradeAgentWithTypeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
