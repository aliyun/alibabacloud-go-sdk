// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarTaskAndActionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSoarTaskAndActionsRequest
	GetLang() *string
	SetPageNumber(v int32) *DescribeSoarTaskAndActionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSoarTaskAndActionsRequest
	GetPageSize() *int32
	SetQueryType(v string) *DescribeSoarTaskAndActionsRequest
	GetQueryType() *string
	SetQueryValue(v string) *DescribeSoarTaskAndActionsRequest
	GetQueryValue() *string
	SetRequestUuid(v string) *DescribeSoarTaskAndActionsRequest
	GetRequestUuid() *string
}

type DescribeSoarTaskAndActionsRequest struct {
	// The language of the request and response. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The maximum value is 20. If you do not specify this parameter, 10 entries are returned.
	//
	// > Specify a value for this parameter.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The trigger type of the task. Valid values:
	//
	// - **stream**: The task is triggered by a data stream.
	//
	// - **debug**: The task is triggered by a debugging process.
	//
	// - **manual**: The task is triggered manually.
	//
	// - **timer**: The task is triggered by a timer.
	//
	// - **SubInvoke**: The task is triggered by a child flow.
	//
	// example:
	//
	// debug
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The input parameter of the playbook.
	//
	// example:
	//
	// input
	QueryValue *string `json:"QueryValue,omitempty" xml:"QueryValue,omitempty"`
	// The UUID of the playbook task.
	//
	// example:
	//
	// 1077f2f9-25e8-42d9-bfdf-1528e1313f6d
	RequestUuid *string `json:"RequestUuid,omitempty" xml:"RequestUuid,omitempty"`
}

func (s DescribeSoarTaskAndActionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarTaskAndActionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSoarTaskAndActionsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSoarTaskAndActionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSoarTaskAndActionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSoarTaskAndActionsRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *DescribeSoarTaskAndActionsRequest) GetQueryValue() *string {
	return s.QueryValue
}

func (s *DescribeSoarTaskAndActionsRequest) GetRequestUuid() *string {
	return s.RequestUuid
}

func (s *DescribeSoarTaskAndActionsRequest) SetLang(v string) *DescribeSoarTaskAndActionsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSoarTaskAndActionsRequest) SetPageNumber(v int32) *DescribeSoarTaskAndActionsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSoarTaskAndActionsRequest) SetPageSize(v int32) *DescribeSoarTaskAndActionsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSoarTaskAndActionsRequest) SetQueryType(v string) *DescribeSoarTaskAndActionsRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeSoarTaskAndActionsRequest) SetQueryValue(v string) *DescribeSoarTaskAndActionsRequest {
	s.QueryValue = &v
	return s
}

func (s *DescribeSoarTaskAndActionsRequest) SetRequestUuid(v string) *DescribeSoarTaskAndActionsRequest {
	s.RequestUuid = &v
	return s
}

func (s *DescribeSoarTaskAndActionsRequest) Validate() error {
	return dara.Validate(s)
}
