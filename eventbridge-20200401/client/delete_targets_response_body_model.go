// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTargetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteTargetsResponseBody
	GetCode() *string
	SetData(v *DeleteTargetsResponseBodyData) *DeleteTargetsResponseBody
	GetData() *DeleteTargetsResponseBodyData
	SetMessage(v string) *DeleteTargetsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTargetsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTargetsResponseBody
	GetSuccess() *bool
}

type DeleteTargetsResponseBody struct {
	// The response code. The code 200 indicates that the request was successful. Other codes indicate that the request failed. For information about error codes, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DeleteTargetsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	//
	// example:
	//
	// EventBusNotExist
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 78FA9EAC-F0C0-58B0-871E-9F9756CE1D29
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTargetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTargetsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteTargetsResponseBody) GetData() *DeleteTargetsResponseBodyData {
	return s.Data
}

func (s *DeleteTargetsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTargetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTargetsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTargetsResponseBody) SetCode(v string) *DeleteTargetsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTargetsResponseBody) SetData(v *DeleteTargetsResponseBodyData) *DeleteTargetsResponseBody {
	s.Data = v
	return s
}

func (s *DeleteTargetsResponseBody) SetMessage(v string) *DeleteTargetsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTargetsResponseBody) SetRequestId(v string) *DeleteTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTargetsResponseBody) SetSuccess(v bool) *DeleteTargetsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTargetsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteTargetsResponseBodyData struct {
	// The information about the event body that failed to be processed.
	ErrorEntries []*DeleteTargetsResponseBodyDataErrorEntries `json:"ErrorEntries,omitempty" xml:"ErrorEntries,omitempty" type:"Repeated"`
	// The number of event bodies that failed to be processed. Valid values: 0: No event bodies failed to be processed. An integer other than 0: the number of event bodies that failed to be processed.
	//
	// example:
	//
	// 0
	ErrorEntriesCount *int32 `json:"ErrorEntriesCount,omitempty" xml:"ErrorEntriesCount,omitempty"`
}

func (s DeleteTargetsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteTargetsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteTargetsResponseBodyData) GetErrorEntries() []*DeleteTargetsResponseBodyDataErrorEntries {
	return s.ErrorEntries
}

func (s *DeleteTargetsResponseBodyData) GetErrorEntriesCount() *int32 {
	return s.ErrorEntriesCount
}

func (s *DeleteTargetsResponseBodyData) SetErrorEntries(v []*DeleteTargetsResponseBodyDataErrorEntries) *DeleteTargetsResponseBodyData {
	s.ErrorEntries = v
	return s
}

func (s *DeleteTargetsResponseBodyData) SetErrorEntriesCount(v int32) *DeleteTargetsResponseBodyData {
	s.ErrorEntriesCount = &v
	return s
}

func (s *DeleteTargetsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DeleteTargetsResponseBodyDataErrorEntries struct {
	// The ID of the event body that failed to be processed.
	//
	// example:
	//
	// target5
	EntryId *string `json:"EntryId,omitempty" xml:"EntryId,omitempty"`
	// The error code.
	//
	// example:
	//
	// EventRuleTargetIdDuplicate
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The id of event target is duplicate!
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s DeleteTargetsResponseBodyDataErrorEntries) String() string {
	return dara.Prettify(s)
}

func (s DeleteTargetsResponseBodyDataErrorEntries) GoString() string {
	return s.String()
}

func (s *DeleteTargetsResponseBodyDataErrorEntries) GetEntryId() *string {
	return s.EntryId
}

func (s *DeleteTargetsResponseBodyDataErrorEntries) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteTargetsResponseBodyDataErrorEntries) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteTargetsResponseBodyDataErrorEntries) SetEntryId(v string) *DeleteTargetsResponseBodyDataErrorEntries {
	s.EntryId = &v
	return s
}

func (s *DeleteTargetsResponseBodyDataErrorEntries) SetErrorCode(v string) *DeleteTargetsResponseBodyDataErrorEntries {
	s.ErrorCode = &v
	return s
}

func (s *DeleteTargetsResponseBodyDataErrorEntries) SetErrorMessage(v string) *DeleteTargetsResponseBodyDataErrorEntries {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteTargetsResponseBodyDataErrorEntries) Validate() error {
	return dara.Validate(s)
}
