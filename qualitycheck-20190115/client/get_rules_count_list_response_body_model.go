// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRulesCountListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v int32) *GetRulesCountListResponseBody
	GetBusinessType() *int32
	SetCode(v string) *GetRulesCountListResponseBody
	GetCode() *string
	SetCount(v int32) *GetRulesCountListResponseBody
	GetCount() *int32
	SetCurrentPage(v int32) *GetRulesCountListResponseBody
	GetCurrentPage() *int32
	SetData(v *GetRulesCountListResponseBodyData) *GetRulesCountListResponseBody
	GetData() *GetRulesCountListResponseBodyData
	SetHttpStatusCode(v int32) *GetRulesCountListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetRulesCountListResponseBody
	GetMessage() *string
	SetMessages(v *GetRulesCountListResponseBodyMessages) *GetRulesCountListResponseBody
	GetMessages() *GetRulesCountListResponseBodyMessages
	SetPageNumber(v int32) *GetRulesCountListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *GetRulesCountListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetRulesCountListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRulesCountListResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *GetRulesCountListResponseBody
	GetTotalCount() *int32
}

type GetRulesCountListResponseBody struct {
	BusinessType *int32 `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        *GetRulesCountListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *GetRulesCountListResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
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
	// 9987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetRulesCountListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRulesCountListResponseBody) GoString() string {
	return s.String()
}

func (s *GetRulesCountListResponseBody) GetBusinessType() *int32 {
	return s.BusinessType
}

func (s *GetRulesCountListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRulesCountListResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *GetRulesCountListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetRulesCountListResponseBody) GetData() *GetRulesCountListResponseBodyData {
	return s.Data
}

func (s *GetRulesCountListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetRulesCountListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRulesCountListResponseBody) GetMessages() *GetRulesCountListResponseBodyMessages {
	return s.Messages
}

func (s *GetRulesCountListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetRulesCountListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetRulesCountListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRulesCountListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRulesCountListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetRulesCountListResponseBody) SetBusinessType(v int32) *GetRulesCountListResponseBody {
	s.BusinessType = &v
	return s
}

func (s *GetRulesCountListResponseBody) SetCode(v string) *GetRulesCountListResponseBody {
	s.Code = &v
	return s
}

func (s *GetRulesCountListResponseBody) SetCount(v int32) *GetRulesCountListResponseBody {
	s.Count = &v
	return s
}

func (s *GetRulesCountListResponseBody) SetCurrentPage(v int32) *GetRulesCountListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *GetRulesCountListResponseBody) SetData(v *GetRulesCountListResponseBodyData) *GetRulesCountListResponseBody {
	s.Data = v
	return s
}

func (s *GetRulesCountListResponseBody) SetHttpStatusCode(v int32) *GetRulesCountListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRulesCountListResponseBody) SetMessage(v string) *GetRulesCountListResponseBody {
	s.Message = &v
	return s
}

func (s *GetRulesCountListResponseBody) SetMessages(v *GetRulesCountListResponseBodyMessages) *GetRulesCountListResponseBody {
	s.Messages = v
	return s
}

func (s *GetRulesCountListResponseBody) SetPageNumber(v int32) *GetRulesCountListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetRulesCountListResponseBody) SetPageSize(v int32) *GetRulesCountListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetRulesCountListResponseBody) SetRequestId(v string) *GetRulesCountListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRulesCountListResponseBody) SetSuccess(v bool) *GetRulesCountListResponseBody {
	s.Success = &v
	return s
}

func (s *GetRulesCountListResponseBody) SetTotalCount(v int32) *GetRulesCountListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetRulesCountListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	if s.Messages != nil {
		if err := s.Messages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRulesCountListResponseBodyData struct {
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetRulesCountListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRulesCountListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRulesCountListResponseBodyData) GetData() []*string {
	return s.Data
}

func (s *GetRulesCountListResponseBodyData) SetData(v []*string) *GetRulesCountListResponseBodyData {
	s.Data = v
	return s
}

func (s *GetRulesCountListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetRulesCountListResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s GetRulesCountListResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s GetRulesCountListResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *GetRulesCountListResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *GetRulesCountListResponseBodyMessages) SetMessage(v []*string) *GetRulesCountListResponseBodyMessages {
	s.Message = v
	return s
}

func (s *GetRulesCountListResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
