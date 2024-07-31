// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BatchEnrollAccountsRequest struct {
	Accounts []*BatchEnrollAccountsRequestAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// example:
	//
	// afb-bp1durvn3lgqe28v****
	BaselineId    *string                                    `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	BaselineItems []*BatchEnrollAccountsRequestBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s BatchEnrollAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchEnrollAccountsRequest) GoString() string {
	return s.String()
}

func (s *BatchEnrollAccountsRequest) SetAccounts(v []*BatchEnrollAccountsRequestAccounts) *BatchEnrollAccountsRequest {
	s.Accounts = v
	return s
}

func (s *BatchEnrollAccountsRequest) SetBaselineId(v string) *BatchEnrollAccountsRequest {
	s.BaselineId = &v
	return s
}

func (s *BatchEnrollAccountsRequest) SetBaselineItems(v []*BatchEnrollAccountsRequestBaselineItems) *BatchEnrollAccountsRequest {
	s.BaselineItems = v
	return s
}

func (s *BatchEnrollAccountsRequest) SetRegionId(v string) *BatchEnrollAccountsRequest {
	s.RegionId = &v
	return s
}

type BatchEnrollAccountsRequestAccounts struct {
	// example:
	//
	// 12868156179****
	AccountUid *int64 `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
}

func (s BatchEnrollAccountsRequestAccounts) String() string {
	return tea.Prettify(s)
}

func (s BatchEnrollAccountsRequestAccounts) GoString() string {
	return s.String()
}

func (s *BatchEnrollAccountsRequestAccounts) SetAccountUid(v int64) *BatchEnrollAccountsRequestAccounts {
	s.AccountUid = &v
	return s
}

type BatchEnrollAccountsRequestBaselineItems struct {
	// example:
	//
	// {\\"Notifications\\":[{\\"GroupKey\\":\\"account_msg\\",\\"Contacts\\":[{\\"Name\\":\\"aa\\"}],\\"PmsgStatus\\":1,\\"EmailStatus\\":1,\\"SmsStatus\\":1}]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// ACS-BP_ACCOUNT_FACTORY_VPC
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// false
	Skip *bool `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s BatchEnrollAccountsRequestBaselineItems) String() string {
	return tea.Prettify(s)
}

func (s BatchEnrollAccountsRequestBaselineItems) GoString() string {
	return s.String()
}

func (s *BatchEnrollAccountsRequestBaselineItems) SetConfig(v string) *BatchEnrollAccountsRequestBaselineItems {
	s.Config = &v
	return s
}

func (s *BatchEnrollAccountsRequestBaselineItems) SetName(v string) *BatchEnrollAccountsRequestBaselineItems {
	s.Name = &v
	return s
}

func (s *BatchEnrollAccountsRequestBaselineItems) SetSkip(v bool) *BatchEnrollAccountsRequestBaselineItems {
	s.Skip = &v
	return s
}

func (s *BatchEnrollAccountsRequestBaselineItems) SetVersion(v string) *BatchEnrollAccountsRequestBaselineItems {
	s.Version = &v
	return s
}

type BatchEnrollAccountsResponseBody struct {
	// example:
	//
	// 16B208DD-86BD-5E7D-AC93-FFD44B6FBDF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchEnrollAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchEnrollAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchEnrollAccountsResponseBody) SetRequestId(v string) *BatchEnrollAccountsResponseBody {
	s.RequestId = &v
	return s
}

type BatchEnrollAccountsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchEnrollAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchEnrollAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchEnrollAccountsResponse) GoString() string {
	return s.String()
}

func (s *BatchEnrollAccountsResponse) SetHeaders(v map[string]*string) *BatchEnrollAccountsResponse {
	s.Headers = v
	return s
}

func (s *BatchEnrollAccountsResponse) SetStatusCode(v int32) *BatchEnrollAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchEnrollAccountsResponse) SetBody(v *BatchEnrollAccountsResponseBody) *BatchEnrollAccountsResponse {
	s.Body = v
	return s
}

type CreateAccountFactoryBaselineRequest struct {
	BaselineItems []*CreateAccountFactoryBaselineRequestBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
	BaselineName  *string                                             `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	Description   *string                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	// RegionId
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateAccountFactoryBaselineRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountFactoryBaselineRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountFactoryBaselineRequest) SetBaselineItems(v []*CreateAccountFactoryBaselineRequestBaselineItems) *CreateAccountFactoryBaselineRequest {
	s.BaselineItems = v
	return s
}

func (s *CreateAccountFactoryBaselineRequest) SetBaselineName(v string) *CreateAccountFactoryBaselineRequest {
	s.BaselineName = &v
	return s
}

func (s *CreateAccountFactoryBaselineRequest) SetDescription(v string) *CreateAccountFactoryBaselineRequest {
	s.Description = &v
	return s
}

func (s *CreateAccountFactoryBaselineRequest) SetRegionId(v string) *CreateAccountFactoryBaselineRequest {
	s.RegionId = &v
	return s
}

type CreateAccountFactoryBaselineRequestBaselineItems struct {
	// example:
	//
	// {\\"EnabledServices\\":[\\"CEN_TR\\",\\"CDT\\",\\"CMS\\",\\"KMS\\"]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// ACS-BP_ACCOUNT_FACTORY_VPC
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateAccountFactoryBaselineRequestBaselineItems) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountFactoryBaselineRequestBaselineItems) GoString() string {
	return s.String()
}

func (s *CreateAccountFactoryBaselineRequestBaselineItems) SetConfig(v string) *CreateAccountFactoryBaselineRequestBaselineItems {
	s.Config = &v
	return s
}

func (s *CreateAccountFactoryBaselineRequestBaselineItems) SetName(v string) *CreateAccountFactoryBaselineRequestBaselineItems {
	s.Name = &v
	return s
}

func (s *CreateAccountFactoryBaselineRequestBaselineItems) SetVersion(v string) *CreateAccountFactoryBaselineRequestBaselineItems {
	s.Version = &v
	return s
}

type CreateAccountFactoryBaselineResponseBody struct {
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// example:
	//
	// A5592E2E-0FC4-557C-B989-DF229B5EBE13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccountFactoryBaselineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountFactoryBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountFactoryBaselineResponseBody) SetBaselineId(v string) *CreateAccountFactoryBaselineResponseBody {
	s.BaselineId = &v
	return s
}

func (s *CreateAccountFactoryBaselineResponseBody) SetRequestId(v string) *CreateAccountFactoryBaselineResponseBody {
	s.RequestId = &v
	return s
}

type CreateAccountFactoryBaselineResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccountFactoryBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccountFactoryBaselineResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountFactoryBaselineResponse) GoString() string {
	return s.String()
}

func (s *CreateAccountFactoryBaselineResponse) SetHeaders(v map[string]*string) *CreateAccountFactoryBaselineResponse {
	s.Headers = v
	return s
}

func (s *CreateAccountFactoryBaselineResponse) SetStatusCode(v int32) *CreateAccountFactoryBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccountFactoryBaselineResponse) SetBody(v *CreateAccountFactoryBaselineResponseBody) *CreateAccountFactoryBaselineResponse {
	s.Body = v
	return s
}

type DeleteAccountFactoryBaselineRequest struct {
	// example:
	//
	// afb-bp1durvn3lgqe28v****
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAccountFactoryBaselineRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountFactoryBaselineRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountFactoryBaselineRequest) SetBaselineId(v string) *DeleteAccountFactoryBaselineRequest {
	s.BaselineId = &v
	return s
}

func (s *DeleteAccountFactoryBaselineRequest) SetRegionId(v string) *DeleteAccountFactoryBaselineRequest {
	s.RegionId = &v
	return s
}

type DeleteAccountFactoryBaselineResponseBody struct {
	// example:
	//
	// 0F45D888-8C4D-55E5-ACA2-D1515159181D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccountFactoryBaselineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountFactoryBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccountFactoryBaselineResponseBody) SetRequestId(v string) *DeleteAccountFactoryBaselineResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAccountFactoryBaselineResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccountFactoryBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccountFactoryBaselineResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountFactoryBaselineResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccountFactoryBaselineResponse) SetHeaders(v map[string]*string) *DeleteAccountFactoryBaselineResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccountFactoryBaselineResponse) SetStatusCode(v int32) *DeleteAccountFactoryBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccountFactoryBaselineResponse) SetBody(v *DeleteAccountFactoryBaselineResponseBody) *DeleteAccountFactoryBaselineResponse {
	s.Body = v
	return s
}

type EnrollAccountRequest struct {
	// The prefix for the account name of the member.
	//
	// 	- If the account baseline is applied to an account that is newly created, you must configure this parameter.
	//
	// 	- If the account baseline is applied to an existing account, you do not need to configure this parameter.
	//
	// example:
	//
	// test-account
	AccountNamePrefix *string `json:"AccountNamePrefix,omitempty" xml:"AccountNamePrefix,omitempty"`
	// The account ID.
	//
	// 	- If the account baseline is applied to an account that is newly created, you do not need to configure this parameter.
	//
	// 	- If the account baseline is applied to an existing account, you must configure this parameter.
	//
	// example:
	//
	// 12868156179****
	AccountUid *int64 `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
	// The baseline ID.
	//
	// If this parameter is left empty, the default baseline is used.
	//
	// example:
	//
	// afb-bp1durvn3lgqe28v****
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// An array that contains baseline items.
	//
	// If this parameter is specified, the configurations of the baseline items are merged with the baseline of the specified account. The configurations of the same baseline items are subject to the configuration of this parameter. We recommend that you leave this parameter empty and configure the `BaselineId` parameter to specify an account baseline and apply the configuration of the account baseline to the account.
	BaselineItems []*EnrollAccountRequestBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
	// The display name of the account.
	//
	// 	- If the account baseline is applied to an account that is newly created, you must configure this parameter.
	//
	// 	- If the account baseline is applied to an existing account, you do not need to configure this parameter.
	//
	// example:
	//
	// test-account
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the parent folder.
	//
	// 	- If the account baseline is applied to an account that is newly created, you need to specify a parent folder. If you do not configure this parameter, the account is created in the Root folder.
	//
	// 	- If the account baseline is applied to an existing account, you do not need to configure this parameter.
	//
	// example:
	//
	// fd-5ESoku****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The ID of the billing account.
	//
	// 	- If the account baseline is applied to an account that is newly created, you need to specify a billing account. If you do not configure this parameter, the self-pay settlement method is used for the account.
	//
	// 	- If the account baseline is applied to an existing account, you do not need to configure this parameter.
	//
	// example:
	//
	// 19534534552****
	PayerAccountUid *int64 `json:"PayerAccountUid,omitempty" xml:"PayerAccountUid,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The identity type of the member. Valid values:
	//
	// 	- resell (default): The member is an account for a reseller. A relationship is automatically established between the member and the reseller. The management account of the resource directory must be used as the billing account of the member.
	//
	// 	- non_resell: The member is not an account for a reseller. The member is an account that is not associated with a reseller. You can directly use the account to purchase Alibaba Cloud resources. The member is used as its own billing account.
	//
	// > This parameter is available only for resellers at the international site (alibabacloud.com).
	//
	// example:
	//
	// resell
	ResellAccountType *string `json:"ResellAccountType,omitempty" xml:"ResellAccountType,omitempty"`
}

func (s EnrollAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s EnrollAccountRequest) GoString() string {
	return s.String()
}

func (s *EnrollAccountRequest) SetAccountNamePrefix(v string) *EnrollAccountRequest {
	s.AccountNamePrefix = &v
	return s
}

func (s *EnrollAccountRequest) SetAccountUid(v int64) *EnrollAccountRequest {
	s.AccountUid = &v
	return s
}

func (s *EnrollAccountRequest) SetBaselineId(v string) *EnrollAccountRequest {
	s.BaselineId = &v
	return s
}

func (s *EnrollAccountRequest) SetBaselineItems(v []*EnrollAccountRequestBaselineItems) *EnrollAccountRequest {
	s.BaselineItems = v
	return s
}

func (s *EnrollAccountRequest) SetDisplayName(v string) *EnrollAccountRequest {
	s.DisplayName = &v
	return s
}

func (s *EnrollAccountRequest) SetFolderId(v string) *EnrollAccountRequest {
	s.FolderId = &v
	return s
}

func (s *EnrollAccountRequest) SetPayerAccountUid(v int64) *EnrollAccountRequest {
	s.PayerAccountUid = &v
	return s
}

func (s *EnrollAccountRequest) SetRegionId(v string) *EnrollAccountRequest {
	s.RegionId = &v
	return s
}

func (s *EnrollAccountRequest) SetResellAccountType(v string) *EnrollAccountRequest {
	s.ResellAccountType = &v
	return s
}

type EnrollAccountRequestBaselineItems struct {
	// The configurations of the baseline item.
	//
	// example:
	//
	// {\\"Notifications\\":[{\\"GroupKey\\":\\"account_msg\\",\\"Contacts\\":[{\\"Name\\":\\"aa\\"}],\\"PmsgStatus\\":1,\\"EmailStatus\\":1,\\"SmsStatus\\":1}]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The name of the baseline item.
	//
	// example:
	//
	// ACS-BP_ACCOUNT_FACTORY_VPC
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to skip the baseline item. Valid values:
	//
	// 	- false: The baseline item is not skipped.
	//
	// 	- true: The baseline item is skipped.
	//
	// example:
	//
	// false
	Skip *bool `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The version of the baseline item.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s EnrollAccountRequestBaselineItems) String() string {
	return tea.Prettify(s)
}

func (s EnrollAccountRequestBaselineItems) GoString() string {
	return s.String()
}

func (s *EnrollAccountRequestBaselineItems) SetConfig(v string) *EnrollAccountRequestBaselineItems {
	s.Config = &v
	return s
}

func (s *EnrollAccountRequestBaselineItems) SetName(v string) *EnrollAccountRequestBaselineItems {
	s.Name = &v
	return s
}

func (s *EnrollAccountRequestBaselineItems) SetSkip(v bool) *EnrollAccountRequestBaselineItems {
	s.Skip = &v
	return s
}

func (s *EnrollAccountRequestBaselineItems) SetVersion(v string) *EnrollAccountRequestBaselineItems {
	s.Version = &v
	return s
}

type EnrollAccountResponseBody struct {
	// The account ID.
	//
	// example:
	//
	// 143165363236****
	AccountUid *int64 `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7071E5FA-515E-5F53-B335-B87D619C6A66
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnrollAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnrollAccountResponseBody) GoString() string {
	return s.String()
}

func (s *EnrollAccountResponseBody) SetAccountUid(v int64) *EnrollAccountResponseBody {
	s.AccountUid = &v
	return s
}

func (s *EnrollAccountResponseBody) SetRequestId(v string) *EnrollAccountResponseBody {
	s.RequestId = &v
	return s
}

type EnrollAccountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnrollAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnrollAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s EnrollAccountResponse) GoString() string {
	return s.String()
}

func (s *EnrollAccountResponse) SetHeaders(v map[string]*string) *EnrollAccountResponse {
	s.Headers = v
	return s
}

func (s *EnrollAccountResponse) SetStatusCode(v int32) *EnrollAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *EnrollAccountResponse) SetBody(v *EnrollAccountResponseBody) *EnrollAccountResponse {
	s.Body = v
	return s
}

type GetAccountFactoryBaselineRequest struct {
	// The baseline ID.
	//
	// example:
	//
	// afb-bp1nf0enuzb89az*****
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAccountFactoryBaselineRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountFactoryBaselineRequest) GoString() string {
	return s.String()
}

func (s *GetAccountFactoryBaselineRequest) SetBaselineId(v string) *GetAccountFactoryBaselineRequest {
	s.BaselineId = &v
	return s
}

func (s *GetAccountFactoryBaselineRequest) SetRegionId(v string) *GetAccountFactoryBaselineRequest {
	s.RegionId = &v
	return s
}

type GetAccountFactoryBaselineResponseBody struct {
	// The baseline ID.
	//
	// example:
	//
	// afb-bp16ae2k8a3yo3d*****
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The baseline items.
	BaselineItems []*GetAccountFactoryBaselineResponseBodyBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
	// The name of the baseline.
	//
	// example:
	//
	// Default
	BaselineName *string `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	// The time when the baseline was created.
	//
	// example:
	//
	// 2022-11-28T00:46:34Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the baseline.
	//
	// example:
	//
	// Default baseline
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 60D54503-F1F6-51B6-B6FA-A70CBDA2A68C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the baseline. Valid values:
	//
	// 	- System: default baseline.
	//
	// 	- Custom: custom baseline.
	//
	// example:
	//
	// Custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the baseline was updated.
	//
	// example:
	//
	// 2022-11-02T01:00:07Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetAccountFactoryBaselineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountFactoryBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountFactoryBaselineResponseBody) SetBaselineId(v string) *GetAccountFactoryBaselineResponseBody {
	s.BaselineId = &v
	return s
}

func (s *GetAccountFactoryBaselineResponseBody) SetBaselineItems(v []*GetAccountFactoryBaselineResponseBodyBaselineItems) *GetAccountFactoryBaselineResponseBody {
	s.BaselineItems = v
	return s
}

func (s *GetAccountFactoryBaselineResponseBody) SetBaselineName(v string) *GetAccountFactoryBaselineResponseBody {
	s.BaselineName = &v
	return s
}

func (s *GetAccountFactoryBaselineResponseBody) SetCreateTime(v string) *GetAccountFactoryBaselineResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetAccountFactoryBaselineResponseBody) SetDescription(v string) *GetAccountFactoryBaselineResponseBody {
	s.Description = &v
	return s
}

func (s *GetAccountFactoryBaselineResponseBody) SetRequestId(v string) *GetAccountFactoryBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountFactoryBaselineResponseBody) SetType(v string) *GetAccountFactoryBaselineResponseBody {
	s.Type = &v
	return s
}

func (s *GetAccountFactoryBaselineResponseBody) SetUpdateTime(v string) *GetAccountFactoryBaselineResponseBody {
	s.UpdateTime = &v
	return s
}

type GetAccountFactoryBaselineResponseBodyBaselineItems struct {
	// The configuration of the baseline item.
	//
	// The value is a JSON string.
	//
	// example:
	//
	// {\\"Notifications\\":[{\\"GroupKey\\":\\"account_msg\\",\\"Contacts\\":[{\\"Name\\":\\"aa\\"}],\\"PmsgStatus\\":1,\\"EmailStatus\\":1,\\"SmsStatus\\":1}]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The name of the baseline item.
	//
	// example:
	//
	// 1097526274671790
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the baseline item.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetAccountFactoryBaselineResponseBodyBaselineItems) String() string {
	return tea.Prettify(s)
}

func (s GetAccountFactoryBaselineResponseBodyBaselineItems) GoString() string {
	return s.String()
}

func (s *GetAccountFactoryBaselineResponseBodyBaselineItems) SetConfig(v string) *GetAccountFactoryBaselineResponseBodyBaselineItems {
	s.Config = &v
	return s
}

func (s *GetAccountFactoryBaselineResponseBodyBaselineItems) SetName(v string) *GetAccountFactoryBaselineResponseBodyBaselineItems {
	s.Name = &v
	return s
}

func (s *GetAccountFactoryBaselineResponseBodyBaselineItems) SetVersion(v string) *GetAccountFactoryBaselineResponseBodyBaselineItems {
	s.Version = &v
	return s
}

type GetAccountFactoryBaselineResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountFactoryBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountFactoryBaselineResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountFactoryBaselineResponse) GoString() string {
	return s.String()
}

func (s *GetAccountFactoryBaselineResponse) SetHeaders(v map[string]*string) *GetAccountFactoryBaselineResponse {
	s.Headers = v
	return s
}

func (s *GetAccountFactoryBaselineResponse) SetStatusCode(v int32) *GetAccountFactoryBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountFactoryBaselineResponse) SetBody(v *GetAccountFactoryBaselineResponseBody) *GetAccountFactoryBaselineResponse {
	s.Body = v
	return s
}

type GetEnrolledAccountRequest struct {
	// The account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 19534534552****
	AccountUid *int64 `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetEnrolledAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEnrolledAccountRequest) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountRequest) SetAccountUid(v int64) *GetEnrolledAccountRequest {
	s.AccountUid = &v
	return s
}

func (s *GetEnrolledAccountRequest) SetRegionId(v string) *GetEnrolledAccountRequest {
	s.RegionId = &v
	return s
}

type GetEnrolledAccountResponseBody struct {
	// The account ID.
	//
	// example:
	//
	// 12868156179*****
	AccountUid *int64 `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
	// The ID of the baseline that is implemented.
	//
	// example:
	//
	// afb-bp1adadfadsf***
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// An array that contains baseline items.
	BaselineItems []*GetEnrolledAccountResponseBodyBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
	// The time when the account was created.
	//
	// example:
	//
	// 2021-11-01T02:38:27Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The display name of the account.
	//
	// example:
	//
	// test-account
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The error message.
	//
	// >  This parameter is returned if the value of `Status` is `Failed` or `ScheduleFailed`.
	ErrorInfo *GetEnrolledAccountResponseBodyErrorInfo `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty" type:"Struct"`
	// The ID of the parent folder.
	//
	// example:
	//
	// fd-5ESoku****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// Indicates whether the initialization is complete. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// true
	Initialized *bool `json:"Initialized,omitempty" xml:"Initialized,omitempty"`
	// The input parameters that are used when the account was registered.
	Inputs *GetEnrolledAccountResponseBodyInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Struct"`
	// The ID of the management account of the resource directory to which the account belongs.
	//
	// example:
	//
	// 19534534552*****
	MasterAccountUid *int64 `json:"MasterAccountUid,omitempty" xml:"MasterAccountUid,omitempty"`
	// The ID of the settlement account.
	//
	// example:
	//
	// 19534534552*****
	PayerAccountUid *int64 `json:"PayerAccountUid,omitempty" xml:"PayerAccountUid,omitempty"`
	// The progress of the applying the baseline to the account.
	Progress []*GetEnrolledAccountResponseBodyProgress `json:"Progress,omitempty" xml:"Progress,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 768F908D-A66A-5A5D-816C-20C93CBBFEE3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the account. Valid values:
	//
	// 	- Pending: The account is pending to be created.
	//
	// 	- Running: The account is being created.
	//
	// 	- Finished: The account is created.
	//
	// 	- Failed: The account fails to be created.
	//
	// 	- Scheduling: The account is being scheduled.
	//
	// 	- ScheduleFailed: The account fails to be scheduled.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the information about the account was updated.
	//
	// example:
	//
	// 2021-11-01T02:38:27Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetEnrolledAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEnrolledAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponseBody) SetAccountUid(v int64) *GetEnrolledAccountResponseBody {
	s.AccountUid = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetBaselineId(v string) *GetEnrolledAccountResponseBody {
	s.BaselineId = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetBaselineItems(v []*GetEnrolledAccountResponseBodyBaselineItems) *GetEnrolledAccountResponseBody {
	s.BaselineItems = v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetCreateTime(v string) *GetEnrolledAccountResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetDisplayName(v string) *GetEnrolledAccountResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetErrorInfo(v *GetEnrolledAccountResponseBodyErrorInfo) *GetEnrolledAccountResponseBody {
	s.ErrorInfo = v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetFolderId(v string) *GetEnrolledAccountResponseBody {
	s.FolderId = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetInitialized(v bool) *GetEnrolledAccountResponseBody {
	s.Initialized = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetInputs(v *GetEnrolledAccountResponseBodyInputs) *GetEnrolledAccountResponseBody {
	s.Inputs = v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetMasterAccountUid(v int64) *GetEnrolledAccountResponseBody {
	s.MasterAccountUid = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetPayerAccountUid(v int64) *GetEnrolledAccountResponseBody {
	s.PayerAccountUid = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetProgress(v []*GetEnrolledAccountResponseBodyProgress) *GetEnrolledAccountResponseBody {
	s.Progress = v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetRequestId(v string) *GetEnrolledAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetStatus(v string) *GetEnrolledAccountResponseBody {
	s.Status = &v
	return s
}

func (s *GetEnrolledAccountResponseBody) SetUpdateTime(v string) *GetEnrolledAccountResponseBody {
	s.UpdateTime = &v
	return s
}

type GetEnrolledAccountResponseBodyBaselineItems struct {
	// The configurations of the baseline item.
	//
	// example:
	//
	// {\\"Notifications\\":[{\\"GroupKey\\":\\"account_msg\\",\\"Contacts\\":[{\\"Name\\":\\"aa\\"}],\\"PmsgStatus\\":1,\\"EmailStatus\\":1,\\"SmsStatus\\":1}]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The name of the baseline item.
	//
	// example:
	//
	// ACS-BP_ACCOUNT_FACTORY_VPC
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether baseline item is skipped. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	Skip *bool `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The version of the baseline item.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetEnrolledAccountResponseBodyBaselineItems) String() string {
	return tea.Prettify(s)
}

func (s GetEnrolledAccountResponseBodyBaselineItems) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponseBodyBaselineItems) SetConfig(v string) *GetEnrolledAccountResponseBodyBaselineItems {
	s.Config = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyBaselineItems) SetName(v string) *GetEnrolledAccountResponseBodyBaselineItems {
	s.Name = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyBaselineItems) SetSkip(v bool) *GetEnrolledAccountResponseBodyBaselineItems {
	s.Skip = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyBaselineItems) SetVersion(v string) *GetEnrolledAccountResponseBodyBaselineItems {
	s.Version = &v
	return s
}

type GetEnrolledAccountResponseBodyErrorInfo struct {
	// The error code.
	//
	// example:
	//
	// CompliancePackExists
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The compliance pack already exists.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The recommended solution.
	//
	// example:
	//
	// https://next.api.aliyun.com/troubleshoot?q=CompliancePackExists\\\\u0026product=Config
	Recommend *string `json:"Recommend,omitempty" xml:"Recommend,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6D5EAA86-2D41-5CB7-8DA7-B60093ACAA4E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEnrolledAccountResponseBodyErrorInfo) String() string {
	return tea.Prettify(s)
}

func (s GetEnrolledAccountResponseBodyErrorInfo) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponseBodyErrorInfo) SetCode(v string) *GetEnrolledAccountResponseBodyErrorInfo {
	s.Code = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyErrorInfo) SetMessage(v string) *GetEnrolledAccountResponseBodyErrorInfo {
	s.Message = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyErrorInfo) SetRecommend(v string) *GetEnrolledAccountResponseBodyErrorInfo {
	s.Recommend = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyErrorInfo) SetRequestId(v string) *GetEnrolledAccountResponseBodyErrorInfo {
	s.RequestId = &v
	return s
}

type GetEnrolledAccountResponseBodyInputs struct {
	// The prefix of the account name.
	//
	// example:
	//
	// test-account
	AccountNamePrefix *string `json:"AccountNamePrefix,omitempty" xml:"AccountNamePrefix,omitempty"`
	// The account ID.
	//
	// example:
	//
	// 12868156179*****
	AccountUid *int64 `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
	// The baseline items.
	BaselineItems []*GetEnrolledAccountResponseBodyInputsBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
	// The display name of the account.
	//
	// example:
	//
	// test-account
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the parent folder.
	//
	// example:
	//
	// fd-5ESoku****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The ID of the settlement account.
	//
	// example:
	//
	// 19534534552*****
	PayerAccountUid *int64 `json:"PayerAccountUid,omitempty" xml:"PayerAccountUid,omitempty"`
}

func (s GetEnrolledAccountResponseBodyInputs) String() string {
	return tea.Prettify(s)
}

func (s GetEnrolledAccountResponseBodyInputs) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponseBodyInputs) SetAccountNamePrefix(v string) *GetEnrolledAccountResponseBodyInputs {
	s.AccountNamePrefix = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputs) SetAccountUid(v int64) *GetEnrolledAccountResponseBodyInputs {
	s.AccountUid = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputs) SetBaselineItems(v []*GetEnrolledAccountResponseBodyInputsBaselineItems) *GetEnrolledAccountResponseBodyInputs {
	s.BaselineItems = v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputs) SetDisplayName(v string) *GetEnrolledAccountResponseBodyInputs {
	s.DisplayName = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputs) SetFolderId(v string) *GetEnrolledAccountResponseBodyInputs {
	s.FolderId = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputs) SetPayerAccountUid(v int64) *GetEnrolledAccountResponseBodyInputs {
	s.PayerAccountUid = &v
	return s
}

type GetEnrolledAccountResponseBodyInputsBaselineItems struct {
	// The configurations of the baseline item.
	//
	// example:
	//
	// {\\"Contacts\\":[{\\"Name\\":\\"governance\\",\\"Email\\":\\"wibud****@gmail.com\\",\\"Mobile\\":\\"1234\\",\\"Position\\":\\"Other\\"}]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The name of the baseline item.
	//
	// example:
	//
	// ACS-BP_ACCOUNT_FACTORY_VPC
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether baseline item is skipped. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	Skip *bool `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The version of the baseline item.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetEnrolledAccountResponseBodyInputsBaselineItems) String() string {
	return tea.Prettify(s)
}

func (s GetEnrolledAccountResponseBodyInputsBaselineItems) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponseBodyInputsBaselineItems) SetConfig(v string) *GetEnrolledAccountResponseBodyInputsBaselineItems {
	s.Config = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputsBaselineItems) SetName(v string) *GetEnrolledAccountResponseBodyInputsBaselineItems {
	s.Name = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputsBaselineItems) SetSkip(v bool) *GetEnrolledAccountResponseBodyInputsBaselineItems {
	s.Skip = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputsBaselineItems) SetVersion(v string) *GetEnrolledAccountResponseBodyInputsBaselineItems {
	s.Version = &v
	return s
}

type GetEnrolledAccountResponseBodyProgress struct {
	// The name of the baseline item.
	//
	// example:
	//
	// ACS-BP_ACCOUNT_FACTORY_VPC
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of applying the baseline to the account. Valid values:
	//
	// 	- Pending: The baseline is pending to be applied to the account.
	//
	// 	- Running: The baseline is being applied to the account.
	//
	// 	- Finished: : The baseline is applied to the account.
	//
	// 	- Failed: : The baseline fails to be applied to the account.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetEnrolledAccountResponseBodyProgress) String() string {
	return tea.Prettify(s)
}

func (s GetEnrolledAccountResponseBodyProgress) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponseBodyProgress) SetName(v string) *GetEnrolledAccountResponseBodyProgress {
	s.Name = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyProgress) SetStatus(v string) *GetEnrolledAccountResponseBodyProgress {
	s.Status = &v
	return s
}

type GetEnrolledAccountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEnrolledAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEnrolledAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEnrolledAccountResponse) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponse) SetHeaders(v map[string]*string) *GetEnrolledAccountResponse {
	s.Headers = v
	return s
}

func (s *GetEnrolledAccountResponse) SetStatusCode(v int32) *GetEnrolledAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEnrolledAccountResponse) SetBody(v *GetEnrolledAccountResponseBody) *GetEnrolledAccountResponse {
	s.Body = v
	return s
}

type ListAccountFactoryBaselineItemsRequest struct {
	// example:
	//
	// 10
	MaxResults *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Names      []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAACDGQdAEX3m42z3sQ+f3VTK2Xr2DzYbz/SAfc/zJRqod
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// RegionId
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// AccountFactory
	Type     *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Versions []*string `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s ListAccountFactoryBaselineItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccountFactoryBaselineItemsRequest) GoString() string {
	return s.String()
}

func (s *ListAccountFactoryBaselineItemsRequest) SetMaxResults(v int32) *ListAccountFactoryBaselineItemsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsRequest) SetNames(v []*string) *ListAccountFactoryBaselineItemsRequest {
	s.Names = v
	return s
}

func (s *ListAccountFactoryBaselineItemsRequest) SetNextToken(v string) *ListAccountFactoryBaselineItemsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsRequest) SetRegionId(v string) *ListAccountFactoryBaselineItemsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsRequest) SetType(v string) *ListAccountFactoryBaselineItemsRequest {
	s.Type = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsRequest) SetVersions(v []*string) *ListAccountFactoryBaselineItemsRequest {
	s.Versions = v
	return s
}

type ListAccountFactoryBaselineItemsResponseBody struct {
	BaselineItems []*ListAccountFactoryBaselineItemsResponseBodyBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
	// example:
	//
	// AAAAACDGQdAEX3m42z3sQ+f3VTK2Xr2DzYbz/SAfc/zJRqod
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// B40D73D8-76AC-5D3C-AC63-4FC8AFCE6671
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAccountFactoryBaselineItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccountFactoryBaselineItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountFactoryBaselineItemsResponseBody) SetBaselineItems(v []*ListAccountFactoryBaselineItemsResponseBodyBaselineItems) *ListAccountFactoryBaselineItemsResponseBody {
	s.BaselineItems = v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponseBody) SetNextToken(v string) *ListAccountFactoryBaselineItemsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponseBody) SetRequestId(v string) *ListAccountFactoryBaselineItemsResponseBody {
	s.RequestId = &v
	return s
}

type ListAccountFactoryBaselineItemsResponseBodyBaselineItems struct {
	DependsOn []*ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn `json:"DependsOn,omitempty" xml:"DependsOn,omitempty" type:"Repeated"`
	// example:
	//
	// Notification.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// ACS-BP_ACCOUNT_FACTORY_ACCOUNT_NOTIFICATION
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// AccountFactory
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAccountFactoryBaselineItemsResponseBodyBaselineItems) String() string {
	return tea.Prettify(s)
}

func (s ListAccountFactoryBaselineItemsResponseBodyBaselineItems) GoString() string {
	return s.String()
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItems) SetDependsOn(v []*ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn) *ListAccountFactoryBaselineItemsResponseBodyBaselineItems {
	s.DependsOn = v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItems) SetDescription(v string) *ListAccountFactoryBaselineItemsResponseBodyBaselineItems {
	s.Description = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItems) SetName(v string) *ListAccountFactoryBaselineItemsResponseBodyBaselineItems {
	s.Name = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItems) SetType(v string) *ListAccountFactoryBaselineItemsResponseBodyBaselineItems {
	s.Type = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItems) SetVersion(v string) *ListAccountFactoryBaselineItemsResponseBodyBaselineItems {
	s.Version = &v
	return s
}

type ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn struct {
	// example:
	//
	// ACS-BP_ACCOUNT_FACTORY_VPC
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// AccountFactory
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn) String() string {
	return tea.Prettify(s)
}

func (s ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn) GoString() string {
	return s.String()
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn) SetName(v string) *ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn {
	s.Name = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn) SetType(v string) *ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn {
	s.Type = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn) SetVersion(v string) *ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn {
	s.Version = &v
	return s
}

type ListAccountFactoryBaselineItemsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccountFactoryBaselineItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccountFactoryBaselineItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccountFactoryBaselineItemsResponse) GoString() string {
	return s.String()
}

func (s *ListAccountFactoryBaselineItemsResponse) SetHeaders(v map[string]*string) *ListAccountFactoryBaselineItemsResponse {
	s.Headers = v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponse) SetStatusCode(v int32) *ListAccountFactoryBaselineItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccountFactoryBaselineItemsResponse) SetBody(v *ListAccountFactoryBaselineItemsResponseBody) *ListAccountFactoryBaselineItemsResponse {
	s.Body = v
	return s
}

type ListAccountFactoryBaselinesRequest struct {
	// The maximum number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// AAAAALHWGpGoYCcYMxiFfmlhvh62Xr2DzYbz/SAfc*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// RegionId
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAccountFactoryBaselinesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccountFactoryBaselinesRequest) GoString() string {
	return s.String()
}

func (s *ListAccountFactoryBaselinesRequest) SetMaxResults(v int32) *ListAccountFactoryBaselinesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAccountFactoryBaselinesRequest) SetNextToken(v string) *ListAccountFactoryBaselinesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAccountFactoryBaselinesRequest) SetRegionId(v string) *ListAccountFactoryBaselinesRequest {
	s.RegionId = &v
	return s
}

type ListAccountFactoryBaselinesResponseBody struct {
	// An array that consists of baselines.
	Baselines []*ListAccountFactoryBaselinesResponseBodyBaselines `json:"Baselines,omitempty" xml:"Baselines,omitempty" type:"Repeated"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAALHWGpGoYCcYMxiFfmlhvh62Xr2DzYbz/SAfc*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3245E413-7CDD-5287-8988-6A94DE8A8369
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAccountFactoryBaselinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccountFactoryBaselinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountFactoryBaselinesResponseBody) SetBaselines(v []*ListAccountFactoryBaselinesResponseBodyBaselines) *ListAccountFactoryBaselinesResponseBody {
	s.Baselines = v
	return s
}

func (s *ListAccountFactoryBaselinesResponseBody) SetNextToken(v string) *ListAccountFactoryBaselinesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAccountFactoryBaselinesResponseBody) SetRequestId(v string) *ListAccountFactoryBaselinesResponseBody {
	s.RequestId = &v
	return s
}

type ListAccountFactoryBaselinesResponseBodyBaselines struct {
	// The baseline ID.
	//
	// example:
	//
	// afb-bp1durvn3lgqe28v****
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The name of the baseline.
	//
	// example:
	//
	// Default
	BaselineName *string `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	// The time at which the baseline was created.
	//
	// example:
	//
	// 2021-11-30T09:09:28Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the baseline.
	//
	// example:
	//
	// Default baseline
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the baseline. Valid values:
	//
	// 	- System: default baseline
	//
	// 	- Custom: custom baseline
	//
	// example:
	//
	// Custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the baseline was updated.
	//
	// example:
	//
	// 2022-12-29T07:08:40Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListAccountFactoryBaselinesResponseBodyBaselines) String() string {
	return tea.Prettify(s)
}

func (s ListAccountFactoryBaselinesResponseBodyBaselines) GoString() string {
	return s.String()
}

func (s *ListAccountFactoryBaselinesResponseBodyBaselines) SetBaselineId(v string) *ListAccountFactoryBaselinesResponseBodyBaselines {
	s.BaselineId = &v
	return s
}

func (s *ListAccountFactoryBaselinesResponseBodyBaselines) SetBaselineName(v string) *ListAccountFactoryBaselinesResponseBodyBaselines {
	s.BaselineName = &v
	return s
}

func (s *ListAccountFactoryBaselinesResponseBodyBaselines) SetCreateTime(v string) *ListAccountFactoryBaselinesResponseBodyBaselines {
	s.CreateTime = &v
	return s
}

func (s *ListAccountFactoryBaselinesResponseBodyBaselines) SetDescription(v string) *ListAccountFactoryBaselinesResponseBodyBaselines {
	s.Description = &v
	return s
}

func (s *ListAccountFactoryBaselinesResponseBodyBaselines) SetType(v string) *ListAccountFactoryBaselinesResponseBodyBaselines {
	s.Type = &v
	return s
}

func (s *ListAccountFactoryBaselinesResponseBodyBaselines) SetUpdateTime(v string) *ListAccountFactoryBaselinesResponseBodyBaselines {
	s.UpdateTime = &v
	return s
}

type ListAccountFactoryBaselinesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccountFactoryBaselinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccountFactoryBaselinesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccountFactoryBaselinesResponse) GoString() string {
	return s.String()
}

func (s *ListAccountFactoryBaselinesResponse) SetHeaders(v map[string]*string) *ListAccountFactoryBaselinesResponse {
	s.Headers = v
	return s
}

func (s *ListAccountFactoryBaselinesResponse) SetStatusCode(v int32) *ListAccountFactoryBaselinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccountFactoryBaselinesResponse) SetBody(v *ListAccountFactoryBaselinesResponseBody) *ListAccountFactoryBaselinesResponse {
	s.Body = v
	return s
}

type ListEnrolledAccountsRequest struct {
	// The maximum number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// AAAAALHWGpGoYCcYMxiFfmlhvh62Xr2DzYbz/SAfc*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListEnrolledAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEnrolledAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListEnrolledAccountsRequest) SetMaxResults(v int32) *ListEnrolledAccountsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEnrolledAccountsRequest) SetNextToken(v string) *ListEnrolledAccountsRequest {
	s.NextToken = &v
	return s
}

func (s *ListEnrolledAccountsRequest) SetRegionId(v string) *ListEnrolledAccountsRequest {
	s.RegionId = &v
	return s
}

type ListEnrolledAccountsResponseBody struct {
	// The accounts.
	EnrolledAccounts []*ListEnrolledAccountsResponseBodyEnrolledAccounts `json:"EnrolledAccounts,omitempty" xml:"EnrolledAccounts,omitempty" type:"Repeated"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAALHWGpGoYCcYMxiFfmlhvh62Xr2DzYbz/SAfc*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 768F908D-A66A-5A5D-816C-20C93CBBFEE3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEnrolledAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnrolledAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnrolledAccountsResponseBody) SetEnrolledAccounts(v []*ListEnrolledAccountsResponseBodyEnrolledAccounts) *ListEnrolledAccountsResponseBody {
	s.EnrolledAccounts = v
	return s
}

func (s *ListEnrolledAccountsResponseBody) SetNextToken(v string) *ListEnrolledAccountsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEnrolledAccountsResponseBody) SetRequestId(v string) *ListEnrolledAccountsResponseBody {
	s.RequestId = &v
	return s
}

type ListEnrolledAccountsResponseBodyEnrolledAccounts struct {
	// The account ID.
	//
	// example:
	//
	// 19534534552*****
	AccountUid *int64 `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
	// The baseline ID.
	//
	// example:
	//
	// afb-bp1durvn3lgqe28v****
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The time at which the account was created.
	//
	// example:
	//
	// 2021-11-01T02:38:27Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The display name of the account.
	//
	// example:
	//
	// TestAccount
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the parent folder.
	//
	// example:
	//
	// fd-5ESoku****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The ID of the billing account.
	//
	// example:
	//
	// 13161210500*****
	PayerAccountUid *int64 `json:"PayerAccountUid,omitempty" xml:"PayerAccountUid,omitempty"`
	// The creation status of the account. Valid values:
	//
	// 	- Pending: The account is waiting to be created.
	//
	// 	- Running: The account is being created.
	//
	// 	- Finished: The account is created.
	//
	// 	- Failed: The account failed to be created.
	//
	// 	- Scheduling: The account is being scheduled.
	//
	// 	- ScheduleFailed: The account failed to be scheduled.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the information about the account was updated.
	//
	// example:
	//
	// 2021-11-01T02:38:27Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListEnrolledAccountsResponseBodyEnrolledAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListEnrolledAccountsResponseBodyEnrolledAccounts) GoString() string {
	return s.String()
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetAccountUid(v int64) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.AccountUid = &v
	return s
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetBaselineId(v string) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.BaselineId = &v
	return s
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetCreateTime(v string) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.CreateTime = &v
	return s
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetDisplayName(v string) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.DisplayName = &v
	return s
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetFolderId(v string) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.FolderId = &v
	return s
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetPayerAccountUid(v int64) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.PayerAccountUid = &v
	return s
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetStatus(v string) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.Status = &v
	return s
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetUpdateTime(v string) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.UpdateTime = &v
	return s
}

type ListEnrolledAccountsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnrolledAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnrolledAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnrolledAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListEnrolledAccountsResponse) SetHeaders(v map[string]*string) *ListEnrolledAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListEnrolledAccountsResponse) SetStatusCode(v int32) *ListEnrolledAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnrolledAccountsResponse) SetBody(v *ListEnrolledAccountsResponseBody) *ListEnrolledAccountsResponse {
	s.Body = v
	return s
}

type UpdateAccountFactoryBaselineRequest struct {
	BaselineId    *string                                             `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	BaselineItems []*UpdateAccountFactoryBaselineRequestBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
	BaselineName  *string                                             `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	Description   *string                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	// RegionId
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateAccountFactoryBaselineRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountFactoryBaselineRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccountFactoryBaselineRequest) SetBaselineId(v string) *UpdateAccountFactoryBaselineRequest {
	s.BaselineId = &v
	return s
}

func (s *UpdateAccountFactoryBaselineRequest) SetBaselineItems(v []*UpdateAccountFactoryBaselineRequestBaselineItems) *UpdateAccountFactoryBaselineRequest {
	s.BaselineItems = v
	return s
}

func (s *UpdateAccountFactoryBaselineRequest) SetBaselineName(v string) *UpdateAccountFactoryBaselineRequest {
	s.BaselineName = &v
	return s
}

func (s *UpdateAccountFactoryBaselineRequest) SetDescription(v string) *UpdateAccountFactoryBaselineRequest {
	s.Description = &v
	return s
}

func (s *UpdateAccountFactoryBaselineRequest) SetRegionId(v string) *UpdateAccountFactoryBaselineRequest {
	s.RegionId = &v
	return s
}

type UpdateAccountFactoryBaselineRequestBaselineItems struct {
	// example:
	//
	// {\\"EnabledServices\\":[\\"CEN_TR\\",\\"CDT\\",\\"CMS\\",\\"KMS\\"]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// ACS-BP_ACCOUNT_FACTORY_VPC
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateAccountFactoryBaselineRequestBaselineItems) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountFactoryBaselineRequestBaselineItems) GoString() string {
	return s.String()
}

func (s *UpdateAccountFactoryBaselineRequestBaselineItems) SetConfig(v string) *UpdateAccountFactoryBaselineRequestBaselineItems {
	s.Config = &v
	return s
}

func (s *UpdateAccountFactoryBaselineRequestBaselineItems) SetName(v string) *UpdateAccountFactoryBaselineRequestBaselineItems {
	s.Name = &v
	return s
}

func (s *UpdateAccountFactoryBaselineRequestBaselineItems) SetVersion(v string) *UpdateAccountFactoryBaselineRequestBaselineItems {
	s.Version = &v
	return s
}

type UpdateAccountFactoryBaselineResponseBody struct {
	// example:
	//
	// C18A891D-7B04-51A1-AAC6-201727A361CE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAccountFactoryBaselineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountFactoryBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAccountFactoryBaselineResponseBody) SetRequestId(v string) *UpdateAccountFactoryBaselineResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAccountFactoryBaselineResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAccountFactoryBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAccountFactoryBaselineResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccountFactoryBaselineResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccountFactoryBaselineResponse) SetHeaders(v map[string]*string) *UpdateAccountFactoryBaselineResponse {
	s.Headers = v
	return s
}

func (s *UpdateAccountFactoryBaselineResponse) SetStatusCode(v int32) *UpdateAccountFactoryBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAccountFactoryBaselineResponse) SetBody(v *UpdateAccountFactoryBaselineResponseBody) *UpdateAccountFactoryBaselineResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("governance"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 账号工厂批量注册账号
//
// @param request - BatchEnrollAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchEnrollAccountsResponse
func (client *Client) BatchEnrollAccountsWithOptions(request *BatchEnrollAccountsRequest, runtime *util.RuntimeOptions) (_result *BatchEnrollAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accounts)) {
		query["Accounts"] = request.Accounts
	}

	if !tea.BoolValue(util.IsUnset(request.BaselineId)) {
		query["BaselineId"] = request.BaselineId
	}

	if !tea.BoolValue(util.IsUnset(request.BaselineItems)) {
		query["BaselineItems"] = request.BaselineItems
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchEnrollAccounts"),
		Version:     tea.String("2021-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchEnrollAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 账号工厂批量注册账号
//
// @param request - BatchEnrollAccountsRequest
//
// @return BatchEnrollAccountsResponse
func (client *Client) BatchEnrollAccounts(request *BatchEnrollAccountsRequest) (_result *BatchEnrollAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchEnrollAccountsResponse{}
	_body, _err := client.BatchEnrollAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建账号工厂基线
//
// @param request - CreateAccountFactoryBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccountFactoryBaselineResponse
func (client *Client) CreateAccountFactoryBaselineWithOptions(request *CreateAccountFactoryBaselineRequest, runtime *util.RuntimeOptions) (_result *CreateAccountFactoryBaselineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaselineItems)) {
		query["BaselineItems"] = request.BaselineItems
	}

	if !tea.BoolValue(util.IsUnset(request.BaselineName)) {
		query["BaselineName"] = request.BaselineName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccountFactoryBaseline"),
		Version:     tea.String("2021-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccountFactoryBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建账号工厂基线
//
// @param request - CreateAccountFactoryBaselineRequest
//
// @return CreateAccountFactoryBaselineResponse
func (client *Client) CreateAccountFactoryBaseline(request *CreateAccountFactoryBaselineRequest) (_result *CreateAccountFactoryBaselineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccountFactoryBaselineResponse{}
	_body, _err := client.CreateAccountFactoryBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除账号工厂基线
//
// @param request - DeleteAccountFactoryBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccountFactoryBaselineResponse
func (client *Client) DeleteAccountFactoryBaselineWithOptions(request *DeleteAccountFactoryBaselineRequest, runtime *util.RuntimeOptions) (_result *DeleteAccountFactoryBaselineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaselineId)) {
		query["BaselineId"] = request.BaselineId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccountFactoryBaseline"),
		Version:     tea.String("2021-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccountFactoryBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除账号工厂基线
//
// @param request - DeleteAccountFactoryBaselineRequest
//
// @return DeleteAccountFactoryBaselineResponse
func (client *Client) DeleteAccountFactoryBaseline(request *DeleteAccountFactoryBaselineRequest) (_result *DeleteAccountFactoryBaselineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccountFactoryBaselineResponse{}
	_body, _err := client.DeleteAccountFactoryBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enrolls an account. You can create a new account or manage an existing account in the account factory.
//
// Description:
//
// You can call this API operation to create a new account or manage an existing account and apply the account baseline to the account.
//
// Accounts are created in asynchronous mode. After you create an account, you can apply the account baseline to the account. You can call the [GetEnrolledAccount API](~~GetEnrolledAccount~~) operation to view the details about the account to obtain the result of applying the account baseline to the account.
//
// @param request - EnrollAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnrollAccountResponse
func (client *Client) EnrollAccountWithOptions(request *EnrollAccountRequest, runtime *util.RuntimeOptions) (_result *EnrollAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountNamePrefix)) {
		query["AccountNamePrefix"] = request.AccountNamePrefix
	}

	if !tea.BoolValue(util.IsUnset(request.AccountUid)) {
		query["AccountUid"] = request.AccountUid
	}

	if !tea.BoolValue(util.IsUnset(request.BaselineId)) {
		query["BaselineId"] = request.BaselineId
	}

	if !tea.BoolValue(util.IsUnset(request.BaselineItems)) {
		query["BaselineItems"] = request.BaselineItems
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.FolderId)) {
		query["FolderId"] = request.FolderId
	}

	if !tea.BoolValue(util.IsUnset(request.PayerAccountUid)) {
		query["PayerAccountUid"] = request.PayerAccountUid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResellAccountType)) {
		query["ResellAccountType"] = request.ResellAccountType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnrollAccount"),
		Version:     tea.String("2021-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnrollAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enrolls an account. You can create a new account or manage an existing account in the account factory.
//
// Description:
//
// You can call this API operation to create a new account or manage an existing account and apply the account baseline to the account.
//
// Accounts are created in asynchronous mode. After you create an account, you can apply the account baseline to the account. You can call the [GetEnrolledAccount API](~~GetEnrolledAccount~~) operation to view the details about the account to obtain the result of applying the account baseline to the account.
//
// @param request - EnrollAccountRequest
//
// @return EnrollAccountResponse
func (client *Client) EnrollAccount(request *EnrollAccountRequest) (_result *EnrollAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnrollAccountResponse{}
	_body, _err := client.EnrollAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of an account factory baseline.
//
// @param request - GetAccountFactoryBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountFactoryBaselineResponse
func (client *Client) GetAccountFactoryBaselineWithOptions(request *GetAccountFactoryBaselineRequest, runtime *util.RuntimeOptions) (_result *GetAccountFactoryBaselineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaselineId)) {
		query["BaselineId"] = request.BaselineId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccountFactoryBaseline"),
		Version:     tea.String("2021-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountFactoryBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of an account factory baseline.
//
// @param request - GetAccountFactoryBaselineRequest
//
// @return GetAccountFactoryBaselineResponse
func (client *Client) GetAccountFactoryBaseline(request *GetAccountFactoryBaselineRequest) (_result *GetAccountFactoryBaselineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountFactoryBaselineResponse{}
	_body, _err := client.GetAccountFactoryBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about an account that is enrolled in the account factory.
//
// @param request - GetEnrolledAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEnrolledAccountResponse
func (client *Client) GetEnrolledAccountWithOptions(request *GetEnrolledAccountRequest, runtime *util.RuntimeOptions) (_result *GetEnrolledAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountUid)) {
		query["AccountUid"] = request.AccountUid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEnrolledAccount"),
		Version:     tea.String("2021-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEnrolledAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about an account that is enrolled in the account factory.
//
// @param request - GetEnrolledAccountRequest
//
// @return GetEnrolledAccountResponse
func (client *Client) GetEnrolledAccount(request *GetEnrolledAccountRequest) (_result *GetEnrolledAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEnrolledAccountResponse{}
	_body, _err := client.GetEnrolledAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取账号工厂基线元素列表
//
// @param request - ListAccountFactoryBaselineItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccountFactoryBaselineItemsResponse
func (client *Client) ListAccountFactoryBaselineItemsWithOptions(request *ListAccountFactoryBaselineItemsRequest, runtime *util.RuntimeOptions) (_result *ListAccountFactoryBaselineItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Names)) {
		query["Names"] = request.Names
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Versions)) {
		query["Versions"] = request.Versions
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccountFactoryBaselineItems"),
		Version:     tea.String("2021-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccountFactoryBaselineItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取账号工厂基线元素列表
//
// @param request - ListAccountFactoryBaselineItemsRequest
//
// @return ListAccountFactoryBaselineItemsResponse
func (client *Client) ListAccountFactoryBaselineItems(request *ListAccountFactoryBaselineItemsRequest) (_result *ListAccountFactoryBaselineItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccountFactoryBaselineItemsResponse{}
	_body, _err := client.ListAccountFactoryBaselineItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a list of baselines in the account factory.
//
// @param request - ListAccountFactoryBaselinesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccountFactoryBaselinesResponse
func (client *Client) ListAccountFactoryBaselinesWithOptions(request *ListAccountFactoryBaselinesRequest, runtime *util.RuntimeOptions) (_result *ListAccountFactoryBaselinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccountFactoryBaselines"),
		Version:     tea.String("2021-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccountFactoryBaselinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of baselines in the account factory.
//
// @param request - ListAccountFactoryBaselinesRequest
//
// @return ListAccountFactoryBaselinesResponse
func (client *Client) ListAccountFactoryBaselines(request *ListAccountFactoryBaselinesRequest) (_result *ListAccountFactoryBaselinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccountFactoryBaselinesResponse{}
	_body, _err := client.ListAccountFactoryBaselinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of accounts that are enrolled in the account factory.
//
// @param request - ListEnrolledAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnrolledAccountsResponse
func (client *Client) ListEnrolledAccountsWithOptions(request *ListEnrolledAccountsRequest, runtime *util.RuntimeOptions) (_result *ListEnrolledAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEnrolledAccounts"),
		Version:     tea.String("2021-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEnrolledAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of accounts that are enrolled in the account factory.
//
// @param request - ListEnrolledAccountsRequest
//
// @return ListEnrolledAccountsResponse
func (client *Client) ListEnrolledAccounts(request *ListEnrolledAccountsRequest) (_result *ListEnrolledAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEnrolledAccountsResponse{}
	_body, _err := client.ListEnrolledAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新账号工厂基线
//
// @param request - UpdateAccountFactoryBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAccountFactoryBaselineResponse
func (client *Client) UpdateAccountFactoryBaselineWithOptions(request *UpdateAccountFactoryBaselineRequest, runtime *util.RuntimeOptions) (_result *UpdateAccountFactoryBaselineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BaselineId)) {
		query["BaselineId"] = request.BaselineId
	}

	if !tea.BoolValue(util.IsUnset(request.BaselineItems)) {
		query["BaselineItems"] = request.BaselineItems
	}

	if !tea.BoolValue(util.IsUnset(request.BaselineName)) {
		query["BaselineName"] = request.BaselineName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAccountFactoryBaseline"),
		Version:     tea.String("2021-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAccountFactoryBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新账号工厂基线
//
// @param request - UpdateAccountFactoryBaselineRequest
//
// @return UpdateAccountFactoryBaselineResponse
func (client *Client) UpdateAccountFactoryBaseline(request *UpdateAccountFactoryBaselineRequest) (_result *UpdateAccountFactoryBaselineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAccountFactoryBaselineResponse{}
	_body, _err := client.UpdateAccountFactoryBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
