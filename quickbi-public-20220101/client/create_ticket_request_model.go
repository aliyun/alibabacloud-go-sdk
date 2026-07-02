// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CreateTicketRequest
	GetAccountName() *string
	SetAccountType(v int32) *CreateTicketRequest
	GetAccountType() *int32
	SetCmptId(v string) *CreateTicketRequest
	GetCmptId() *string
	SetExpireTime(v int32) *CreateTicketRequest
	GetExpireTime() *int32
	SetGlobalParam(v string) *CreateTicketRequest
	GetGlobalParam() *string
	SetTicketNum(v int32) *CreateTicketRequest
	GetTicketNum() *int32
	SetUserId(v string) *CreateTicketRequest
	GetUserId() *string
	SetWatermarkParam(v string) *CreateTicketRequest
	GetWatermarkParam() *string
	SetWorksId(v string) *CreateTicketRequest
	GetWorksId() *string
}

type CreateTicketRequest struct {
	// Deprecated
	//
	// The account name of the user.
	//
	// - If the user is an Alibaba Cloud account **wangwu**, the format is **[primary account]**, for example, **wangwu**.
	//
	// - If the user is a Resource Access Management (RAM) users account **zhangsan**@aliyun.cn**, the format is **[primary account:RAM user]**, for example, **wangwu:zhangsan**.
	//
	// > Specify either UserId or AccountName. If neither is specified, the report owner is attached by default, and the report is accessed under that user\\"s identity. To configure row-level permissions for data, see [Row-level permissions](https://help.aliyun.com/document_detail/322783.html).
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Deprecated
	//
	// The account type of the user. Valid values:
	//
	// - 1: Alibaba Cloud account
	//
	// - 3: Quick BI custom account
	//
	// - 4: DingTalk
	//
	// - 5: RAM user
	//
	// - 9: WeCom
	//
	// - 10: Lark
	//
	// > If AccountName is specified, AccountType must also be specified.
	//
	// example:
	//
	// 1
	AccountType *int32 `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The component ID. This is the ID of a specific component in the dashboard. Other report types are not supported.
	//
	// To obtain the component ID, see [QueryWorksBloodRelationship](https://next.api.aliyun.com/api/quickbi-public/2022-01-01/QueryWorksBloodRelationship?spm=a2c4g.11186623.0.0.15615d7aWVvWAl&params={}&lang=JAVA&tab=DOC&sdkStyle=old).
	//
	// example:
	//
	// 0fc6a275c7f64f17b1****a306ce0f31
	CmptId *string `json:"CmptId,omitempty" xml:"CmptId,omitempty"`
	// The expiration time.
	//
	// - Unit: minutes.
	//
	// - Default value: 240.
	//
	// example:
	//
	// 200
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The global parameter.
	//
	// example:
	//
	// [{"paramKey":"price","joinType":"and","conditionList":[{"operate":">","value":"0"}]}]
	GlobalParam *string `json:"GlobalParam,omitempty" xml:"GlobalParam,omitempty"`
	// The number of times the ticket can be used. Each time the ticket is used for access, the count decreases by 1.
	//
	// - Default value: 1.
	//
	// - Recommended value: 1.
	//
	// - Maximum value: 99999.
	//
	// example:
	//
	// 1
	TicketNum *int32 `json:"TicketNum,omitempty" xml:"TicketNum,omitempty"`
	// The Quick BI user ID, not your Alibaba Cloud account ID.
	//
	// You can call the [QueryUserInfoByAccount](https://next.api.aliyun.com/api/quickbi-public/2022-01-01/QueryUserInfoByAccount?spm=a2c4g.11186623.0.0.15615d7aWVvWAl&params={}&tab=DOC&sdkStyle=old) operation to obtain the user ID. Example: fe67f61a35a94b7da1a34ba174a7****.
	//
	// > Specify either UserId or AccountName. If neither is specified, the report owner is used by default, and the report is accessed under that user\\"s identity. To configure row-level permissions for data, see [Row-level permissions](https://help.aliyun.com/document_detail/322783.html).
	//
	// example:
	//
	// 46e537466****92704c8
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The watermark parameter for the report.
	//
	// - The value cannot exceed 50 characters.
	//
	// - Watermark parameters are not supported when the report type is data screen.
	//
	// example:
	//
	// ticket embed
	WatermarkParam *string `json:"WatermarkParam,omitempty" xml:"WatermarkParam,omitempty"`
	// The ID of the report for which embedding is enabled. Dashboards, workbooks, data screens, ad hoc queries, ad hoc analyses, and data entry forms are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// a206f5f3-****-e9b17c835b03
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
}

func (s CreateTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateTicketRequest) GetAccountType() *int32 {
	return s.AccountType
}

func (s *CreateTicketRequest) GetCmptId() *string {
	return s.CmptId
}

func (s *CreateTicketRequest) GetExpireTime() *int32 {
	return s.ExpireTime
}

func (s *CreateTicketRequest) GetGlobalParam() *string {
	return s.GlobalParam
}

func (s *CreateTicketRequest) GetTicketNum() *int32 {
	return s.TicketNum
}

func (s *CreateTicketRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateTicketRequest) GetWatermarkParam() *string {
	return s.WatermarkParam
}

func (s *CreateTicketRequest) GetWorksId() *string {
	return s.WorksId
}

func (s *CreateTicketRequest) SetAccountName(v string) *CreateTicketRequest {
	s.AccountName = &v
	return s
}

func (s *CreateTicketRequest) SetAccountType(v int32) *CreateTicketRequest {
	s.AccountType = &v
	return s
}

func (s *CreateTicketRequest) SetCmptId(v string) *CreateTicketRequest {
	s.CmptId = &v
	return s
}

func (s *CreateTicketRequest) SetExpireTime(v int32) *CreateTicketRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateTicketRequest) SetGlobalParam(v string) *CreateTicketRequest {
	s.GlobalParam = &v
	return s
}

func (s *CreateTicketRequest) SetTicketNum(v int32) *CreateTicketRequest {
	s.TicketNum = &v
	return s
}

func (s *CreateTicketRequest) SetUserId(v string) *CreateTicketRequest {
	s.UserId = &v
	return s
}

func (s *CreateTicketRequest) SetWatermarkParam(v string) *CreateTicketRequest {
	s.WatermarkParam = &v
	return s
}

func (s *CreateTicketRequest) SetWorksId(v string) *CreateTicketRequest {
	s.WorksId = &v
	return s
}

func (s *CreateTicketRequest) Validate() error {
	return dara.Validate(s)
}
