// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIAgentPhoneNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListAIAgentPhoneNumberResponseBodyData) *ListAIAgentPhoneNumberResponseBody
	GetData() []*ListAIAgentPhoneNumberResponseBodyData
	SetPageNumber(v int32) *ListAIAgentPhoneNumberResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAIAgentPhoneNumberResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAIAgentPhoneNumberResponseBody
	GetRequestId() *string
	SetTotalNumber(v int32) *ListAIAgentPhoneNumberResponseBody
	GetTotalNumber() *int32
}

type ListAIAgentPhoneNumberResponseBody struct {
	Data []*ListAIAgentPhoneNumberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ************16-412C-B127-******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 186
	TotalNumber *int32 `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s ListAIAgentPhoneNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIAgentPhoneNumberResponseBody) GetData() []*ListAIAgentPhoneNumberResponseBodyData {
	return s.Data
}

func (s *ListAIAgentPhoneNumberResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAIAgentPhoneNumberResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAIAgentPhoneNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAIAgentPhoneNumberResponseBody) GetTotalNumber() *int32 {
	return s.TotalNumber
}

func (s *ListAIAgentPhoneNumberResponseBody) SetData(v []*ListAIAgentPhoneNumberResponseBodyData) *ListAIAgentPhoneNumberResponseBody {
	s.Data = v
	return s
}

func (s *ListAIAgentPhoneNumberResponseBody) SetPageNumber(v int32) *ListAIAgentPhoneNumberResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAIAgentPhoneNumberResponseBody) SetPageSize(v int32) *ListAIAgentPhoneNumberResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAIAgentPhoneNumberResponseBody) SetRequestId(v string) *ListAIAgentPhoneNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAIAgentPhoneNumberResponseBody) SetTotalNumber(v int32) *ListAIAgentPhoneNumberResponseBody {
	s.TotalNumber = &v
	return s
}

func (s *ListAIAgentPhoneNumberResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAIAgentPhoneNumberResponseBodyData struct {
	// example:
	//
	// 132*****683
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAIAgentPhoneNumberResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentPhoneNumberResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAIAgentPhoneNumberResponseBodyData) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ListAIAgentPhoneNumberResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ListAIAgentPhoneNumberResponseBodyData) SetPhoneNumber(v string) *ListAIAgentPhoneNumberResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *ListAIAgentPhoneNumberResponseBodyData) SetStatus(v int32) *ListAIAgentPhoneNumberResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListAIAgentPhoneNumberResponseBodyData) Validate() error {
	return dara.Validate(s)
}
