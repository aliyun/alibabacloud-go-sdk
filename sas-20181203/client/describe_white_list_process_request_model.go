// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteListProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeWhiteListProcessRequest
	GetCurrentPage() *int32
	SetDesc(v int32) *DescribeWhiteListProcessRequest
	GetDesc() *int32
	SetLang(v string) *DescribeWhiteListProcessRequest
	GetLang() *string
	SetOrderBy(v int32) *DescribeWhiteListProcessRequest
	GetOrderBy() *int32
	SetPageSize(v int32) *DescribeWhiteListProcessRequest
	GetPageSize() *int32
	SetProcessName(v string) *DescribeWhiteListProcessRequest
	GetProcessName() *string
	SetProcessType(v int32) *DescribeWhiteListProcessRequest
	GetProcessType() *int32
	SetSourceIp(v string) *DescribeWhiteListProcessRequest
	GetSourceIp() *string
	SetStrategyId(v int64) *DescribeWhiteListProcessRequest
	GetStrategyId() *int64
}

type DescribeWhiteListProcessRequest struct {
	// The page number. Default value: **1**. Pages start from page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The sort order. Default value: descending order. Valid values:
	//
	// 	- **1**: ascending order
	//
	// 	- **2**: descending order
	//
	// example:
	//
	// 2
	Desc *int32 `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The item based on which you want to sort the returned results. Default value: **process type**. Valid values:
	//
	// 	- **1**: process type
	//
	// 	- **2**: degree of trustability
	//
	// example:
	//
	// 1
	OrderBy *int32 `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of entries per page. Maximum value: 1000. Default value: 20. If you leave this parameter empty, 20 data entries are returned per page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the process.
	//
	// example:
	//
	// JAVA
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// The type of the process. Valid values:
	//
	// 	- **1**: trusted
	//
	// 	- **2**: suspicious
	//
	// 	- **3**: malicious
	//
	// example:
	//
	// 1
	ProcessType *int32 `json:"ProcessType,omitempty" xml:"ProcessType,omitempty"`
	// The source IP address of the request. You do not need to specify this parameter. It is automatically obtained by the system.
	//
	// example:
	//
	// 27.223.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the policy.
	//
	// >  You can call the [DescribeWhiteListStrategyList](~~DescribeWhiteListStrategyList~~) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8562
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s DescribeWhiteListProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListProcessRequest) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListProcessRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWhiteListProcessRequest) GetDesc() *int32 {
	return s.Desc
}

func (s *DescribeWhiteListProcessRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeWhiteListProcessRequest) GetOrderBy() *int32 {
	return s.OrderBy
}

func (s *DescribeWhiteListProcessRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWhiteListProcessRequest) GetProcessName() *string {
	return s.ProcessName
}

func (s *DescribeWhiteListProcessRequest) GetProcessType() *int32 {
	return s.ProcessType
}

func (s *DescribeWhiteListProcessRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeWhiteListProcessRequest) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *DescribeWhiteListProcessRequest) SetCurrentPage(v int32) *DescribeWhiteListProcessRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWhiteListProcessRequest) SetDesc(v int32) *DescribeWhiteListProcessRequest {
	s.Desc = &v
	return s
}

func (s *DescribeWhiteListProcessRequest) SetLang(v string) *DescribeWhiteListProcessRequest {
	s.Lang = &v
	return s
}

func (s *DescribeWhiteListProcessRequest) SetOrderBy(v int32) *DescribeWhiteListProcessRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeWhiteListProcessRequest) SetPageSize(v int32) *DescribeWhiteListProcessRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWhiteListProcessRequest) SetProcessName(v string) *DescribeWhiteListProcessRequest {
	s.ProcessName = &v
	return s
}

func (s *DescribeWhiteListProcessRequest) SetProcessType(v int32) *DescribeWhiteListProcessRequest {
	s.ProcessType = &v
	return s
}

func (s *DescribeWhiteListProcessRequest) SetSourceIp(v string) *DescribeWhiteListProcessRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeWhiteListProcessRequest) SetStrategyId(v int64) *DescribeWhiteListProcessRequest {
	s.StrategyId = &v
	return s
}

func (s *DescribeWhiteListProcessRequest) Validate() error {
	return dara.Validate(s)
}
