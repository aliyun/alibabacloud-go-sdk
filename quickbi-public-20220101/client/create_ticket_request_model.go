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
	// The user\\"s account name.
	//
	// - If the user is an Alibaba Cloud primary account **wangwu**, the format is **[Primary Account]**, for example, **wangwu**.
	//
	// - If the user is a RAM account **zhangsan**@aliyun.cn**, the format is **[Primary Account: Sub-Account]**, for example, **wangwu:zhangsan**.
	//
	// > Only one of UserId and AccountName needs to be filled in. If neither is filled in, it will default to binding the report\\"s Owner, and the report will be accessed with that user\\"s identity. If you need to configure row-level permissions, please refer to [Row-Level Permissions](https://help.aliyun.com/document_detail/322783.html).
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Deprecated
	//
	// The type of the user\\"s account.
	//
	// - 1: Alibaba Cloud account
	//
	// - 3: Quick BI self-built account
	//
	// - 4: DingTalk
	//
	// - 5: RAM sub-account
	//
	// - 9: WeCom
	//
	// - 10: Feishu
	//
	// > If AccountName is not empty, then AccountType must also not be empty.
	//
	// example:
	//
	// 1
	AccountType *int32 `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// Component ID. This is the ID of a component within the above-mentioned dashboard; other types of works do not support this.
	//
	// Refer to [QueryWorksBloodRelationship](https://next.api.aliyun.com/api/quickbi-public/2022-01-01/QueryWorksBloodRelationship?spm=a2c4g.11186623.0.0.15615d7aWVvWAl&params={}&lang=JAVA&tab=DOC&sdkStyle=old) for the API to get the component ID.
	//
	// example:
	//
	// 0fc6a275c7f64f17b1****a306ce0f31
	CmptId *string `json:"CmptId,omitempty" xml:"CmptId,omitempty"`
	// Expiration time
	//
	// - Unit: minutes
	//
	// - Default: 240
	//
	// example:
	//
	// 200
	ExpireTime *int32 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Global parameters for the report filter conditions.
	//
	// - A string in JsonArray format.
	//
	// > If you need to use global parameter capabilities, please contact the [Quick BI Operations Manager](https://h5-alimebot.dingtalk.com/intl/index.htm?spm=a2c4g.11186623.0.0.3da14f6chrDv9e&sourceType=ding_talk&from=DEFFB9G5KBByQkwq23wneFIOmaJ).
	//
	// example:
	//
	// [{"paramKey":"price","joinType":"and","conditionList":[{"operate":">","value":"0"}]}]
	GlobalParam *string `json:"GlobalParam,omitempty" xml:"GlobalParam,omitempty"`
	// The number of tickets. Each time a ticket is used, the number of tickets decreases by 1.
	//
	// - Default value: 1
	//
	// - Recommended value: 1
	//
	// - Maximum value: 99999
	//
	// example:
	//
	// 1
	TicketNum *int32 `json:"TicketNum,omitempty" xml:"TicketNum,omitempty"`
	// Quick BI\\"s UserId, which is not your Alibaba Cloud account ID.
	//
	// You can call the [QueryUserInfoByAccount](https://next.api.aliyun.com/api/quickbi-public/2022-01-01/QueryUserInfoByAccount?spm=a2c4g.11186623.0.0.15615d7aWVvWAl&params={}&tab=DOC&sdkStyle=old) interface to obtain the UserId. An example of a UserId is fe67f61a35a94b7da1a34ba174a7****.
	//
	// > Only one of UserId and AccountName needs to be filled in. If neither is filled in, it will default to binding the report\\"s Owner, and the report will be accessed with that user\\"s identity. If you need to configure row-level permissions, please refer to [Row-Level Permissions](https://help.aliyun.com/document_detail/322783.html).
	//
	// example:
	//
	// 46e537466****92704c8
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Watermark parameters for the report.
	//
	// - Must not exceed 50 characters.
	//
	// - When the report type is a large screen, watermark parameter passing is not supported.
	//
	// example:
	//
	// test
	WatermarkParam *string `json:"WatermarkParam,omitempty" xml:"WatermarkParam,omitempty"`
	// The ID of the report to be embedded. Currently supports dashboards, spreadsheets, data screens, self-service data retrieval, ad-hoc analysis, and data entry.
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
