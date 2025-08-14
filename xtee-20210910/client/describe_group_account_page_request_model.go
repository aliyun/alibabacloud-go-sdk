// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupAccountPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeGroupAccountPageRequest
	GetLang() *string
	SetCommunityNo(v string) *DescribeGroupAccountPageRequest
	GetCommunityNo() *string
	SetCurrentPage(v string) *DescribeGroupAccountPageRequest
	GetCurrentPage() *string
	SetDirection(v string) *DescribeGroupAccountPageRequest
	GetDirection() *string
	SetFieldKey(v string) *DescribeGroupAccountPageRequest
	GetFieldKey() *string
	SetFieldVal(v string) *DescribeGroupAccountPageRequest
	GetFieldVal() *string
	SetIsPage(v bool) *DescribeGroupAccountPageRequest
	GetIsPage() *bool
	SetOrder(v string) *DescribeGroupAccountPageRequest
	GetOrder() *string
	SetPageSize(v string) *DescribeGroupAccountPageRequest
	GetPageSize() *string
	SetRegId(v string) *DescribeGroupAccountPageRequest
	GetRegId() *string
	SetTaskId(v string) *DescribeGroupAccountPageRequest
	GetTaskId() *string
}

type DescribeGroupAccountPageRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Community number.
	//
	// example:
	//
	// 129838420210118141502KiJ1SZL2
	CommunityNo *string `json:"communityNo,omitempty" xml:"communityNo,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Order direction.
	//
	// example:
	//
	// 0
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// fieldKey.
	//
	// example:
	//
	// mobile
	FieldKey *string `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	// fieldVal.
	//
	// example:
	//
	// 18000000000
	FieldVal *string `json:"fieldVal,omitempty" xml:"fieldVal,omitempty"`
	// Whether to paginate.
	//
	// example:
	//
	// true
	IsPage *bool `json:"isPage,omitempty" xml:"isPage,omitempty"`
	// Sorting condition.
	//
	// example:
	//
	// asc
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
	// Page size, default value is 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Task ID.
	//
	// example:
	//
	// 6770764
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s DescribeGroupAccountPageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupAccountPageRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupAccountPageRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGroupAccountPageRequest) GetCommunityNo() *string {
	return s.CommunityNo
}

func (s *DescribeGroupAccountPageRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeGroupAccountPageRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeGroupAccountPageRequest) GetFieldKey() *string {
	return s.FieldKey
}

func (s *DescribeGroupAccountPageRequest) GetFieldVal() *string {
	return s.FieldVal
}

func (s *DescribeGroupAccountPageRequest) GetIsPage() *bool {
	return s.IsPage
}

func (s *DescribeGroupAccountPageRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeGroupAccountPageRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeGroupAccountPageRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeGroupAccountPageRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeGroupAccountPageRequest) SetLang(v string) *DescribeGroupAccountPageRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGroupAccountPageRequest) SetCommunityNo(v string) *DescribeGroupAccountPageRequest {
	s.CommunityNo = &v
	return s
}

func (s *DescribeGroupAccountPageRequest) SetCurrentPage(v string) *DescribeGroupAccountPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupAccountPageRequest) SetDirection(v string) *DescribeGroupAccountPageRequest {
	s.Direction = &v
	return s
}

func (s *DescribeGroupAccountPageRequest) SetFieldKey(v string) *DescribeGroupAccountPageRequest {
	s.FieldKey = &v
	return s
}

func (s *DescribeGroupAccountPageRequest) SetFieldVal(v string) *DescribeGroupAccountPageRequest {
	s.FieldVal = &v
	return s
}

func (s *DescribeGroupAccountPageRequest) SetIsPage(v bool) *DescribeGroupAccountPageRequest {
	s.IsPage = &v
	return s
}

func (s *DescribeGroupAccountPageRequest) SetOrder(v string) *DescribeGroupAccountPageRequest {
	s.Order = &v
	return s
}

func (s *DescribeGroupAccountPageRequest) SetPageSize(v string) *DescribeGroupAccountPageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupAccountPageRequest) SetRegId(v string) *DescribeGroupAccountPageRequest {
	s.RegId = &v
	return s
}

func (s *DescribeGroupAccountPageRequest) SetTaskId(v string) *DescribeGroupAccountPageRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeGroupAccountPageRequest) Validate() error {
	return dara.Validate(s)
}
