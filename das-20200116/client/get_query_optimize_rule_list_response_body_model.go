// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryOptimizeRuleListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQueryOptimizeRuleListResponseBody
	GetCode() *string
	SetData(v *GetQueryOptimizeRuleListResponseBodyData) *GetQueryOptimizeRuleListResponseBody
	GetData() *GetQueryOptimizeRuleListResponseBodyData
	SetMessage(v string) *GetQueryOptimizeRuleListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQueryOptimizeRuleListResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetQueryOptimizeRuleListResponseBody
	GetSuccess() *string
}

type GetQueryOptimizeRuleListResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information.
	Data *GetQueryOptimizeRuleListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQueryOptimizeRuleListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeRuleListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQueryOptimizeRuleListResponseBody) GetData() *GetQueryOptimizeRuleListResponseBodyData {
	return s.Data
}

func (s *GetQueryOptimizeRuleListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQueryOptimizeRuleListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQueryOptimizeRuleListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetQueryOptimizeRuleListResponseBody) SetCode(v string) *GetQueryOptimizeRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBody) SetData(v *GetQueryOptimizeRuleListResponseBodyData) *GetQueryOptimizeRuleListResponseBody {
	s.Data = v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBody) SetMessage(v string) *GetQueryOptimizeRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBody) SetRequestId(v string) *GetQueryOptimizeRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBody) SetSuccess(v string) *GetQueryOptimizeRuleListResponseBody {
	s.Success = &v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQueryOptimizeRuleListResponseBodyData struct {
	// A reserved parameter.
	//
	// example:
	//
	// None
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The information about tags.
	List []*GetQueryOptimizeRuleListResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetQueryOptimizeRuleListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeRuleListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeRuleListResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *GetQueryOptimizeRuleListResponseBodyData) GetList() []*GetQueryOptimizeRuleListResponseBodyDataList {
	return s.List
}

func (s *GetQueryOptimizeRuleListResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetQueryOptimizeRuleListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetQueryOptimizeRuleListResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetQueryOptimizeRuleListResponseBodyData) SetExtra(v string) *GetQueryOptimizeRuleListResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBodyData) SetList(v []*GetQueryOptimizeRuleListResponseBodyDataList) *GetQueryOptimizeRuleListResponseBodyData {
	s.List = v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBodyData) SetPageNo(v int32) *GetQueryOptimizeRuleListResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBodyData) SetPageSize(v int32) *GetQueryOptimizeRuleListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBodyData) SetTotal(v int64) *GetQueryOptimizeRuleListResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetQueryOptimizeRuleListResponseBodyDataList struct {
	// The name of the tag.
	//
	// example:
	//
	// LARGE_ROWS_EXAMINED
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// None
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The type of the tag. **Predefined*	- is returned, which indicates that the tag is added by the system.
	//
	// example:
	//
	// Predefined
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetQueryOptimizeRuleListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetQueryOptimizeRuleListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetQueryOptimizeRuleListResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *GetQueryOptimizeRuleListResponseBodyDataList) GetRuleId() *string {
	return s.RuleId
}

func (s *GetQueryOptimizeRuleListResponseBodyDataList) GetType() *string {
	return s.Type
}

func (s *GetQueryOptimizeRuleListResponseBodyDataList) SetName(v string) *GetQueryOptimizeRuleListResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBodyDataList) SetRuleId(v string) *GetQueryOptimizeRuleListResponseBodyDataList {
	s.RuleId = &v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBodyDataList) SetType(v string) *GetQueryOptimizeRuleListResponseBodyDataList {
	s.Type = &v
	return s
}

func (s *GetQueryOptimizeRuleListResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
