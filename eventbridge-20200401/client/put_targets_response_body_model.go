// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutTargetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PutTargetsResponseBody
	GetCode() *string
	SetData(v *PutTargetsResponseBodyData) *PutTargetsResponseBody
	GetData() *PutTargetsResponseBodyData
	SetMessage(v string) *PutTargetsResponseBody
	GetMessage() *string
	SetRequestId(v string) *PutTargetsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PutTargetsResponseBody
	GetSuccess() *bool
}

type PutTargetsResponseBody struct {
	// The response code.
	//
	// - Success: The request was successful.
	//
	// - Other values indicate an error. For details, see Error codes.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *PutTargetsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// The event rule not existed!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID that Alibaba Cloud generates for the request.
	//
	// example:
	//
	// 6FB52207-7621-5292-BDF2-A17E2E984160
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns true if the operation is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutTargetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *PutTargetsResponseBody) GetCode() *string {
	return s.Code
}

func (s *PutTargetsResponseBody) GetData() *PutTargetsResponseBodyData {
	return s.Data
}

func (s *PutTargetsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PutTargetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutTargetsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PutTargetsResponseBody) SetCode(v string) *PutTargetsResponseBody {
	s.Code = &v
	return s
}

func (s *PutTargetsResponseBody) SetData(v *PutTargetsResponseBodyData) *PutTargetsResponseBody {
	s.Data = v
	return s
}

func (s *PutTargetsResponseBody) SetMessage(v string) *PutTargetsResponseBody {
	s.Message = &v
	return s
}

func (s *PutTargetsResponseBody) SetRequestId(v string) *PutTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutTargetsResponseBody) SetSuccess(v bool) *PutTargetsResponseBody {
	s.Success = &v
	return s
}

func (s *PutTargetsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PutTargetsResponseBodyData struct {
	// Details about the event targets that failed to be processed.
	ErrorEntries []*PutTargetsResponseBodyDataErrorEntries `json:"ErrorEntries,omitempty" xml:"ErrorEntries,omitempty" type:"Repeated"`
	// The number of event targets that failed to be processed. A value of 0 indicates that all event targets were processed successfully.
	//
	// -
	//
	// -
	//
	// example:
	//
	// 0
	ErrorEntriesCount *int32 `json:"ErrorEntriesCount,omitempty" xml:"ErrorEntriesCount,omitempty"`
}

func (s PutTargetsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PutTargetsResponseBodyData) GoString() string {
	return s.String()
}

func (s *PutTargetsResponseBodyData) GetErrorEntries() []*PutTargetsResponseBodyDataErrorEntries {
	return s.ErrorEntries
}

func (s *PutTargetsResponseBodyData) GetErrorEntriesCount() *int32 {
	return s.ErrorEntriesCount
}

func (s *PutTargetsResponseBodyData) SetErrorEntries(v []*PutTargetsResponseBodyDataErrorEntries) *PutTargetsResponseBodyData {
	s.ErrorEntries = v
	return s
}

func (s *PutTargetsResponseBodyData) SetErrorEntriesCount(v int32) *PutTargetsResponseBodyData {
	s.ErrorEntriesCount = &v
	return s
}

func (s *PutTargetsResponseBodyData) Validate() error {
	if s.ErrorEntries != nil {
		for _, item := range s.ErrorEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutTargetsResponseBodyDataErrorEntries struct {
	// The ID of the failed event target.
	//
	// example:
	//
	// Mlm123456JHd2RsRoKw
	EntryId *string `json:"EntryId,omitempty" xml:"EntryId,omitempty"`
	// The error code.
	//
	// example:
	//
	// EventRuleTargetIdDuplicate
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The description of the error.
	//
	// example:
	//
	// The id of event target is duplicate!
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s PutTargetsResponseBodyDataErrorEntries) String() string {
	return dara.Prettify(s)
}

func (s PutTargetsResponseBodyDataErrorEntries) GoString() string {
	return s.String()
}

func (s *PutTargetsResponseBodyDataErrorEntries) GetEntryId() *string {
	return s.EntryId
}

func (s *PutTargetsResponseBodyDataErrorEntries) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *PutTargetsResponseBodyDataErrorEntries) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *PutTargetsResponseBodyDataErrorEntries) SetEntryId(v string) *PutTargetsResponseBodyDataErrorEntries {
	s.EntryId = &v
	return s
}

func (s *PutTargetsResponseBodyDataErrorEntries) SetErrorCode(v string) *PutTargetsResponseBodyDataErrorEntries {
	s.ErrorCode = &v
	return s
}

func (s *PutTargetsResponseBodyDataErrorEntries) SetErrorMessage(v string) *PutTargetsResponseBodyDataErrorEntries {
	s.ErrorMessage = &v
	return s
}

func (s *PutTargetsResponseBodyDataErrorEntries) Validate() error {
	return dara.Validate(s)
}
