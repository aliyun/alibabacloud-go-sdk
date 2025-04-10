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
	// The resource accounts.
	Accounts []*BatchEnrollAccountsRequestAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// The baseline ID.
	//
	// If this parameter is left empty, the default baseline is used.
	//
	// example:
	//
	// afb-bp1durvn3lgqe28v****
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The baseline items.
	//
	// If this parameter is specified, the configurations of the baseline items are merged with the baseline applied to the specified account. The configurations of the same baseline items are subject to the configurations of this parameter. We recommend that you leave this parameter empty and configure the `BaselineId` parameter to specify an account baseline and apply the configurations of the account baseline to the account.
	BaselineItems []*BatchEnrollAccountsRequestBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
	// The region ID.
	//
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
	// The account ID. This parameter is required.
	//
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
	// The request ID.
	//
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
	// An array that contains the baseline items.
	//
	// You can call the [ListAccountFactoryBaselineItems](~~ListAccountFactoryBaselineItems~~) operation to query a list of baseline items supported by the account factory in Cloud Governance Center.
	BaselineItems []*CreateAccountFactoryBaselineRequestBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
	// The name of the baseline.
	//
	// example:
	//
	// Default
	BaselineName *string `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	// The description of the baseline.
	//
	// example:
	//
	// Default Baseline.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The region ID.
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
	// The configurations of the baseline item. The value of this parameter is a JSON string.
	//
	// example:
	//
	// {\\"EnabledServices\\":[\\"CEN_TR\\",\\"CDT\\",\\"CMS\\",\\"KMS\\"]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The name of the baseline item.
	//
	// example:
	//
	// ACS-BP_ACCOUNT_FACTORY_VPC
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the baseline item.
	//
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
	// The baseline ID.
	//
	// example:
	//
	// afb-bp1e6ixtiwupap8m****
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The request ID.
	//
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
	// The baseline ID.
	//
	// example:
	//
	// afb-bp1durvn3lgqe28v****
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The region ID.
	//
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
	// The request ID.
	//
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
	// The array that contains baseline items.
	//
	// If this parameter is specified, the configurations of the baseline items are merged with the baseline applied to the specified account. The configurations of the same baseline items are subject to the configurations of this parameter. We recommend that you leave this parameter empty and configure the `BaselineId` parameter to specify an account baseline and apply the configurations of the account baseline to the account.
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
	// The tags. You can specify up to 20 tags.
	Tag []*EnrollAccountRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *EnrollAccountRequest) SetTag(v []*EnrollAccountRequestTag) *EnrollAccountRequest {
	s.Tag = v
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
	// Whether to skip the baseline item. Valid values:
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

type EnrollAccountRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// tagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// tagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s EnrollAccountRequestTag) String() string {
	return tea.Prettify(s)
}

func (s EnrollAccountRequestTag) GoString() string {
	return s.String()
}

func (s *EnrollAccountRequestTag) SetKey(v string) *EnrollAccountRequestTag {
	s.Key = &v
	return s
}

func (s *EnrollAccountRequestTag) SetValue(v string) *EnrollAccountRequestTag {
	s.Value = &v
	return s
}

type EnrollAccountShrinkRequest struct {
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
	// The array that contains baseline items.
	//
	// If this parameter is specified, the configurations of the baseline items are merged with the baseline applied to the specified account. The configurations of the same baseline items are subject to the configurations of this parameter. We recommend that you leave this parameter empty and configure the `BaselineId` parameter to specify an account baseline and apply the configurations of the account baseline to the account.
	BaselineItems []*EnrollAccountShrinkRequestBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
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
	// The tags. You can specify up to 20 tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s EnrollAccountShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s EnrollAccountShrinkRequest) GoString() string {
	return s.String()
}

func (s *EnrollAccountShrinkRequest) SetAccountNamePrefix(v string) *EnrollAccountShrinkRequest {
	s.AccountNamePrefix = &v
	return s
}

func (s *EnrollAccountShrinkRequest) SetAccountUid(v int64) *EnrollAccountShrinkRequest {
	s.AccountUid = &v
	return s
}

func (s *EnrollAccountShrinkRequest) SetBaselineId(v string) *EnrollAccountShrinkRequest {
	s.BaselineId = &v
	return s
}

func (s *EnrollAccountShrinkRequest) SetBaselineItems(v []*EnrollAccountShrinkRequestBaselineItems) *EnrollAccountShrinkRequest {
	s.BaselineItems = v
	return s
}

func (s *EnrollAccountShrinkRequest) SetDisplayName(v string) *EnrollAccountShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *EnrollAccountShrinkRequest) SetFolderId(v string) *EnrollAccountShrinkRequest {
	s.FolderId = &v
	return s
}

func (s *EnrollAccountShrinkRequest) SetPayerAccountUid(v int64) *EnrollAccountShrinkRequest {
	s.PayerAccountUid = &v
	return s
}

func (s *EnrollAccountShrinkRequest) SetRegionId(v string) *EnrollAccountShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *EnrollAccountShrinkRequest) SetResellAccountType(v string) *EnrollAccountShrinkRequest {
	s.ResellAccountType = &v
	return s
}

func (s *EnrollAccountShrinkRequest) SetTagShrink(v string) *EnrollAccountShrinkRequest {
	s.TagShrink = &v
	return s
}

type EnrollAccountShrinkRequestBaselineItems struct {
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
	// Whether to skip the baseline item. Valid values:
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

func (s EnrollAccountShrinkRequestBaselineItems) String() string {
	return tea.Prettify(s)
}

func (s EnrollAccountShrinkRequestBaselineItems) GoString() string {
	return s.String()
}

func (s *EnrollAccountShrinkRequestBaselineItems) SetConfig(v string) *EnrollAccountShrinkRequestBaselineItems {
	s.Config = &v
	return s
}

func (s *EnrollAccountShrinkRequestBaselineItems) SetName(v string) *EnrollAccountShrinkRequestBaselineItems {
	s.Name = &v
	return s
}

func (s *EnrollAccountShrinkRequestBaselineItems) SetSkip(v bool) *EnrollAccountShrinkRequestBaselineItems {
	s.Skip = &v
	return s
}

func (s *EnrollAccountShrinkRequestBaselineItems) SetVersion(v string) *EnrollAccountShrinkRequestBaselineItems {
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
	// The array that contains baseline items.
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
	// Input parameters used to create an account.
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
	// The update time.
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
	// The configuration of the baseline item.
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
	// The tag.
	Tag []*GetEnrolledAccountResponseBodyInputsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *GetEnrolledAccountResponseBodyInputs) SetTag(v []*GetEnrolledAccountResponseBodyInputsTag) *GetEnrolledAccountResponseBodyInputs {
	s.Tag = v
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

type GetEnrolledAccountResponseBodyInputsTag struct {
	// The tag key.
	//
	// example:
	//
	// product
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// governance
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEnrolledAccountResponseBodyInputsTag) String() string {
	return tea.Prettify(s)
}

func (s GetEnrolledAccountResponseBodyInputsTag) GoString() string {
	return s.String()
}

func (s *GetEnrolledAccountResponseBodyInputsTag) SetKey(v string) *GetEnrolledAccountResponseBodyInputsTag {
	s.Key = &v
	return s
}

func (s *GetEnrolledAccountResponseBodyInputsTag) SetValue(v string) *GetEnrolledAccountResponseBodyInputsTag {
	s.Value = &v
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
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The names of the baseline items.
	Names []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// AAAAACDGQdAEX3m42z3sQ+f3VTK2Xr2DzYbz/SAfc/zJRqod
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the baseline items.
	//
	// example:
	//
	// AccountFactory
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The versions of the baseline items.
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
	// The baseline items.
	BaselineItems []*ListAccountFactoryBaselineItemsResponseBodyBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAACDGQdAEX3m42z3sQ+f3VTK2Xr2DzYbz/SAfc/zJRqod
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
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
	// The dependency of the baseline item.
	DependsOn []*ListAccountFactoryBaselineItemsResponseBodyBaselineItemsDependsOn `json:"DependsOn,omitempty" xml:"DependsOn,omitempty" type:"Repeated"`
	// The description of the baseline item.
	//
	// example:
	//
	// Notification.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the baseline item.
	//
	// example:
	//
	// ACS-BP_ACCOUNT_FACTORY_ACCOUNT_NOTIFICATION
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the baseline item.
	//
	// example:
	//
	// AccountFactory
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the baseline item.
	//
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
	// The name of the baseline item.
	//
	// example:
	//
	// ACS-BP_ACCOUNT_FACTORY_VPC
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the baseline item.
	//
	// example:
	//
	// AccountFactory
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the baseline item.
	//
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
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// You do not need to specify this parameter for the first request.
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
	// The baselines.
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
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// You do not need to specify this parameter for the first request.
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
	// The enrolled accounts.
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
	// The ID of the baseline that is implemented.
	//
	// example:
	//
	// afb-bp1durvn3lgqe28v****
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The creation time.
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
	// The ID of the settlement account.
	//
	// example:
	//
	// 13161210500*****
	PayerAccountUid *int64 `json:"PayerAccountUid,omitempty" xml:"PayerAccountUid,omitempty"`
	// The creation status. Valid values:
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
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The update time.
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

type ListEvaluationMetadataRequest struct {
	// The language. The information is returned in the specified language. Valid values:
	//
	// 	- en: English
	//
	// 	- zh: Chinese
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListEvaluationMetadataRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationMetadataRequest) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataRequest) SetLanguage(v string) *ListEvaluationMetadataRequest {
	s.Language = &v
	return s
}

func (s *ListEvaluationMetadataRequest) SetRegionId(v string) *ListEvaluationMetadataRequest {
	s.RegionId = &v
	return s
}

type ListEvaluationMetadataResponseBody struct {
	// The metadata of a governance maturity check.
	EvaluationMetadata []*ListEvaluationMetadataResponseBodyEvaluationMetadata `json:"EvaluationMetadata,omitempty" xml:"EvaluationMetadata,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 16B208DD-86BD-5E7D-AC93-FFD44B6FBDF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEvaluationMetadataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataResponseBody) SetEvaluationMetadata(v []*ListEvaluationMetadataResponseBodyEvaluationMetadata) *ListEvaluationMetadataResponseBody {
	s.EvaluationMetadata = v
	return s
}

func (s *ListEvaluationMetadataResponseBody) SetRequestId(v string) *ListEvaluationMetadataResponseBody {
	s.RequestId = &v
	return s
}

type ListEvaluationMetadataResponseBodyEvaluationMetadata struct {
	// The metadata objects of a specific metadata type.
	Metadata []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata `json:"Metadata,omitempty" xml:"Metadata,omitempty" type:"Repeated"`
	// The type of the metadata. Valid values:
	//
	// 	- Metric: the check item
	//
	// example:
	//
	// Metric
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadata) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadata) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadata) SetMetadata(v []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) *ListEvaluationMetadataResponseBodyEvaluationMetadata {
	s.Metadata = v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadata) SetType(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadata {
	s.Type = &v
	return s
}

type ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata struct {
	// The category of the check item.
	//
	// example:
	//
	// Security
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The description of the check item.
	//
	// example:
	//
	// If you use an AccessKey pair of an Alibaba Cloud account, you have full permissions on the resources of the account. You cannot set limits on the account, such as setting limits on source IP addresses or access duration. If the AccessKey pair is leaked, resources within the account are exposed to high security risks. If your Alibaba Cloud account has an existing AccessKey pair, the check result is Non-compliant.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the check item.
	//
	// example:
	//
	// An AccessKey pair is enabled for the Alibaba Cloud account.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the metadata.
	//
	// example:
	//
	// pxgtda****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The governance level of the check item.
	//
	// example:
	//
	// High
	RecommendationLevel *string `json:"RecommendationLevel,omitempty" xml:"RecommendationLevel,omitempty"`
	// The metadata of the fixing task.
	RemediationMetadata *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadata `json:"RemediationMetadata,omitempty" xml:"RemediationMetadata,omitempty" type:"Struct"`
	// The metadata of the checked resources.
	ResourceMetadata *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadata `json:"ResourceMetadata,omitempty" xml:"ResourceMetadata,omitempty" type:"Struct"`
	// The scope of the check item. Valid values:
	//
	// 	- Account: the check item in a single-account governance maturity check
	//
	// 	- ResourceDirectory: the check item in a multi-account governance maturity check
	//
	// example:
	//
	// Account
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The status of the check item. Valid values:
	//
	// 	- Released: The check item is released.
	//
	// 	- Beta: The check item is pre-released.
	//
	// example:
	//
	// Released
	Stage *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) SetCategory(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata {
	s.Category = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) SetDescription(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata {
	s.Description = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) SetDisplayName(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata {
	s.DisplayName = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) SetId(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata {
	s.Id = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) SetRecommendationLevel(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata {
	s.RecommendationLevel = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) SetRemediationMetadata(v *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadata) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata {
	s.RemediationMetadata = v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) SetResourceMetadata(v *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadata) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata {
	s.ResourceMetadata = v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) SetScope(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata {
	s.Scope = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata) SetStage(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadata {
	s.Stage = &v
	return s
}

type ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadata struct {
	// The fixing items.
	Remediation []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation `json:"Remediation,omitempty" xml:"Remediation,omitempty" type:"Repeated"`
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadata) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadata) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadata) SetRemediation(v []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadata {
	s.Remediation = v
	return s
}

type ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation struct {
	// The fixing operations.
	Actions []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	// The type of the fixing method. Valid values:
	//
	// 	- Manual: manual fixing
	//
	// 	- QuickFix: quick fixing
	//
	// 	- Analysis: auxiliary decision-making
	//
	// example:
	//
	// Manual
	RemediationType *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation) SetActions(v []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation {
	s.Actions = v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation) SetRemediationType(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediation {
	s.RemediationType = &v
	return s
}

type ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions struct {
	// The fixing method.
	//
	// >  This parameter is returned only if the value of `RemediationType` is `Analysis`.
	//
	// example:
	//
	// UnusedAccessKeyInRamUser
	Classification *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	// The fixing cost.
	//
	// example:
	//
	// You are not charged for this operation.
	CostDescription *string `json:"CostDescription,omitempty" xml:"CostDescription,omitempty"`
	// The description of the fixing item.
	//
	// >  This parameter is returned only if the value of `RemediationType` is `Analysis`.
	//
	// example:
	//
	// Console logon is enabled for the RAM user. The RAM user owns an AccessKey pair that is never used.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The content of the fixing items.
	Guidance []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance `json:"Guidance,omitempty" xml:"Guidance,omitempty" type:"Repeated"`
	// The usage notes of the fixing item.
	//
	// example:
	//
	// The BestPracticesForIdentityAndPermissions compliance package is enabled in Cloud Config to check the settings and usage of the AccessKey pair, Alibaba Cloud account, and RAM users.
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// The fixing suggestion.
	//
	// >  This parameter is returned only if the value of `RemediationType` is `Analysis`.
	//
	// example:
	//
	// Console logon is enabled for the RAM user and the RAM user owns an AccessKey pair, while the AccessKey pair has never been used by the RAM user. We recommend that you disable the AccessKey pair for 90 days. If no related issue occurs during this period, you can delete the AccessKey pair.
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) SetClassification(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions {
	s.Classification = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) SetCostDescription(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions {
	s.CostDescription = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) SetDescription(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions {
	s.Description = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) SetGuidance(v []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions {
	s.Guidance = v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) SetNotice(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions {
	s.Notice = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions) SetSuggestion(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActions {
	s.Suggestion = &v
	return s
}

type ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance struct {
	// The display name of the fixing button.
	//
	// example:
	//
	// Manual fixing
	ButtonName *string `json:"ButtonName,omitempty" xml:"ButtonName,omitempty"`
	// The navigation URL of the fixing button.
	//
	// example:
	//
	// https://ram.console.aliyun.com/users
	ButtonRef *string `json:"ButtonRef,omitempty" xml:"ButtonRef,omitempty"`
	// The fixing procedure.
	//
	// example:
	//
	// You must replace the AccessKey pair of your Alibaba Cloud account. To do so, perform the following steps:</br>1. Log on to the RAM console. In the left-side navigation pane, choose Identities > Users. On the Users page, click Create User.</br>2. On the Create User page, enter a logon name and select OpenAPI Access for the Access Mode parameter.</br>3. After the RAM user is created, save the AccessKey pair. Then, find the user that you created on the Users page and click Add Permissions in the Actions column. In the Grant Permission panel, find the AdministratorAccess policy and attach it to the RAM user.</br>4. In a program, replace the AccessKey pair of the Alibaba Cloud account with the AccessKey pair of the RAM user created in the previous step and check whether the program runs as expected in the test environment.</br>5. If the program runs as expected, publish the program to the production environment and disable the previous AccessKey pair of your Alibaba Cloud account. Then, check whether the program runs as expected.</br>6. If the program runs as expected, delete the disabled AccessKey pair after the specified period of time, such as 90 days.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The title of the fixing procedure.
	//
	// example:
	//
	// Scenario 3: AccessKey pair that is used within the last 90 days
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance) SetButtonName(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance {
	s.ButtonName = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance) SetButtonRef(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance {
	s.ButtonRef = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance) SetContent(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance {
	s.Content = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance) SetTitle(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataRemediationMetadataRemediationActionsGuidance {
	s.Title = &v
	return s
}

type ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadata struct {
	// The metadata of the resource properties.
	ResourcePropertyMetadata []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata `json:"ResourcePropertyMetadata,omitempty" xml:"ResourcePropertyMetadata,omitempty" type:"Repeated"`
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadata) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadata) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadata) SetResourcePropertyMetadata(v []*ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadata {
	s.ResourcePropertyMetadata = v
	return s
}

type ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata struct {
	// The display name of the resource property.
	//
	// example:
	//
	// AccessKey Pair Last Used At
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The name of the resource property.
	//
	// example:
	//
	// AkLastUsedTime
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	// The type of the resource property.
	//
	// example:
	//
	// String
	PropertyType *string `json:"PropertyType,omitempty" xml:"PropertyType,omitempty"`
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata) SetDisplayName(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata {
	s.DisplayName = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata) SetPropertyName(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata {
	s.PropertyName = &v
	return s
}

func (s *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata) SetPropertyType(v string) *ListEvaluationMetadataResponseBodyEvaluationMetadataMetadataResourceMetadataResourcePropertyMetadata {
	s.PropertyType = &v
	return s
}

type ListEvaluationMetadataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEvaluationMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEvaluationMetadataResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationMetadataResponse) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetadataResponse) SetHeaders(v map[string]*string) *ListEvaluationMetadataResponse {
	s.Headers = v
	return s
}

func (s *ListEvaluationMetadataResponse) SetStatusCode(v int32) *ListEvaluationMetadataResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEvaluationMetadataResponse) SetBody(v *ListEvaluationMetadataResponseBody) *ListEvaluationMetadataResponse {
	s.Body = v
	return s
}

type ListEvaluationMetricDetailsRequest struct {
	// The Alibaba Cloud account ID of the member. This parameter takes effect only when a multi-account governance maturity check is performed.
	//
	// example:
	//
	// 103144549568****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the check item.
	//
	// You can call the [ListEvaluationMetadata](https://help.aliyun.com/document_detail/2841889.html) operation to query the ID of the check item.
	//
	// example:
	//
	// xfyve5****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The maximum number of entries to return for a single request. Default value: 5.
	//
	// example:
	//
	// 5
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAGEaXR18y1rqykZHIqRuBejOqED4S3Xne33c7zbn****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ListEvaluationMetricDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationMetricDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetricDetailsRequest) SetAccountId(v int64) *ListEvaluationMetricDetailsRequest {
	s.AccountId = &v
	return s
}

func (s *ListEvaluationMetricDetailsRequest) SetId(v string) *ListEvaluationMetricDetailsRequest {
	s.Id = &v
	return s
}

func (s *ListEvaluationMetricDetailsRequest) SetMaxResults(v int32) *ListEvaluationMetricDetailsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEvaluationMetricDetailsRequest) SetNextToken(v string) *ListEvaluationMetricDetailsRequest {
	s.NextToken = &v
	return s
}

func (s *ListEvaluationMetricDetailsRequest) SetRegionId(v string) *ListEvaluationMetricDetailsRequest {
	s.RegionId = &v
	return s
}

func (s *ListEvaluationMetricDetailsRequest) SetSnapshotId(v string) *ListEvaluationMetricDetailsRequest {
	s.SnapshotId = &v
	return s
}

type ListEvaluationMetricDetailsResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAGEaXR18y1rqykZHIqRuBejOqED4S3Xne33c7zbn****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AC9BD94C-D20C-4D27-88D4-89E8D75C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the non-compliant resources.
	Resources []*ListEvaluationMetricDetailsResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s ListEvaluationMetricDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationMetricDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetricDetailsResponseBody) SetNextToken(v string) *ListEvaluationMetricDetailsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBody) SetRequestId(v string) *ListEvaluationMetricDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBody) SetResources(v []*ListEvaluationMetricDetailsResponseBodyResources) *ListEvaluationMetricDetailsResponseBody {
	s.Resources = v
	return s
}

type ListEvaluationMetricDetailsResponseBodyResources struct {
	// The compliance status of the resource. Valid values:
	//
	// 	- NonCompliant: non-compliant.
	//
	// 	- Excluded: ignored.
	//
	// 	- PendingExclusion: to be ignored.
	//
	// 	- PendingInclusion: to be unignored.
	//
	// example:
	//
	// NonCompliant
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The check results further analyzed by auxiliary decision-making.
	//
	// >  This parameter is returned only when the check item supports the auxiliary decision-making feature.
	//
	// example:
	//
	// RecentUnloginRamUser
	ResourceClassification *string `json:"ResourceClassification,omitempty" xml:"ResourceClassification,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// 26435103783237****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// test
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The ID of the Alibaba Cloud account that owns the resource.
	//
	// example:
	//
	// 176618589410****
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The attributes of the resource.
	ResourceProperties []*ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties `json:"ResourceProperties,omitempty" xml:"ResourceProperties,omitempty" type:"Repeated"`
	// The type of the resource.
	//
	// example:
	//
	// ACS::RAM::User
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListEvaluationMetricDetailsResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationMetricDetailsResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) SetComplianceType(v string) *ListEvaluationMetricDetailsResponseBodyResources {
	s.ComplianceType = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) SetRegionId(v string) *ListEvaluationMetricDetailsResponseBodyResources {
	s.RegionId = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) SetResourceClassification(v string) *ListEvaluationMetricDetailsResponseBodyResources {
	s.ResourceClassification = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) SetResourceId(v string) *ListEvaluationMetricDetailsResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) SetResourceName(v string) *ListEvaluationMetricDetailsResponseBodyResources {
	s.ResourceName = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) SetResourceOwnerId(v int64) *ListEvaluationMetricDetailsResponseBodyResources {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) SetResourceProperties(v []*ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties) *ListEvaluationMetricDetailsResponseBodyResources {
	s.ResourceProperties = v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBodyResources) SetResourceType(v string) *ListEvaluationMetricDetailsResponseBodyResources {
	s.ResourceType = &v
	return s
}

type ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties struct {
	// The name of the resource attribute.
	//
	// example:
	//
	// DisplayName
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	// The value of the resource attribute.
	//
	// example:
	//
	// example
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
}

func (s ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties) SetPropertyName(v string) *ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties {
	s.PropertyName = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties) SetPropertyValue(v string) *ListEvaluationMetricDetailsResponseBodyResourcesResourceProperties {
	s.PropertyValue = &v
	return s
}

type ListEvaluationMetricDetailsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEvaluationMetricDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEvaluationMetricDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationMetricDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListEvaluationMetricDetailsResponse) SetHeaders(v map[string]*string) *ListEvaluationMetricDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListEvaluationMetricDetailsResponse) SetStatusCode(v int32) *ListEvaluationMetricDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEvaluationMetricDetailsResponse) SetBody(v *ListEvaluationMetricDetailsResponseBody) *ListEvaluationMetricDetailsResponse {
	s.Body = v
	return s
}

type ListEvaluationResultsRequest struct {
	// The Alibaba Cloud account ID of the member. This parameter takes effect only when a multi-account governance maturity check is performed.
	//
	// example:
	//
	// 176618589410****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The filter conditions.
	Filters []*ListEvaluationResultsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ListEvaluationResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationResultsRequest) GoString() string {
	return s.String()
}

func (s *ListEvaluationResultsRequest) SetAccountId(v int64) *ListEvaluationResultsRequest {
	s.AccountId = &v
	return s
}

func (s *ListEvaluationResultsRequest) SetFilters(v []*ListEvaluationResultsRequestFilters) *ListEvaluationResultsRequest {
	s.Filters = v
	return s
}

func (s *ListEvaluationResultsRequest) SetRegionId(v string) *ListEvaluationResultsRequest {
	s.RegionId = &v
	return s
}

func (s *ListEvaluationResultsRequest) SetSnapshotId(v string) *ListEvaluationResultsRequest {
	s.SnapshotId = &v
	return s
}

type ListEvaluationResultsRequestFilters struct {
	// The key of the filter condition. Valid values:
	//
	// 	- ResourceId: the resource ID.
	//
	// 	- ResourceName: the name of the resource.
	//
	// 	- ResourceType: the resource type.
	//
	// example:
	//
	// ResourceId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The list of filter condition values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListEvaluationResultsRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationResultsRequestFilters) GoString() string {
	return s.String()
}

func (s *ListEvaluationResultsRequestFilters) SetKey(v string) *ListEvaluationResultsRequestFilters {
	s.Key = &v
	return s
}

func (s *ListEvaluationResultsRequestFilters) SetValues(v []*string) *ListEvaluationResultsRequestFilters {
	s.Values = v
	return s
}

type ListEvaluationResultsResponseBody struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 176618589410****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BD57329E-131A-59F4-8746-E1CD8D7B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The check results, including the status of the overall check and the results of check items.
	Results *ListEvaluationResultsResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
}

func (s ListEvaluationResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEvaluationResultsResponseBody) SetAccountId(v int64) *ListEvaluationResultsResponseBody {
	s.AccountId = &v
	return s
}

func (s *ListEvaluationResultsResponseBody) SetRequestId(v string) *ListEvaluationResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEvaluationResultsResponseBody) SetResults(v *ListEvaluationResultsResponseBodyResults) *ListEvaluationResultsResponseBody {
	s.Results = v
	return s
}

type ListEvaluationResultsResponseBodyResults struct {
	// The end time of the overall check. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-12-13T03:35:00Z
	EvaluationTime *string `json:"EvaluationTime,omitempty" xml:"EvaluationTime,omitempty"`
	// The check results.
	MetricResults []*ListEvaluationResultsResponseBodyResultsMetricResults `json:"MetricResults,omitempty" xml:"MetricResults,omitempty" type:"Repeated"`
	// The status of the overall check. Valid values:
	//
	// 	- Running: The check is in progress.
	//
	// 	- Finished: The check is complete.
	//
	// 	- failed: The check fails.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The overall score.
	//
	// example:
	//
	// 0.6453
	TotalScore *float64 `json:"TotalScore,omitempty" xml:"TotalScore,omitempty"`
}

func (s ListEvaluationResultsResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationResultsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ListEvaluationResultsResponseBodyResults) SetEvaluationTime(v string) *ListEvaluationResultsResponseBodyResults {
	s.EvaluationTime = &v
	return s
}

func (s *ListEvaluationResultsResponseBodyResults) SetMetricResults(v []*ListEvaluationResultsResponseBodyResultsMetricResults) *ListEvaluationResultsResponseBodyResults {
	s.MetricResults = v
	return s
}

func (s *ListEvaluationResultsResponseBodyResults) SetStatus(v string) *ListEvaluationResultsResponseBodyResults {
	s.Status = &v
	return s
}

func (s *ListEvaluationResultsResponseBodyResults) SetTotalScore(v float64) *ListEvaluationResultsResponseBodyResults {
	s.TotalScore = &v
	return s
}

type ListEvaluationResultsResponseBodyResultsMetricResults struct {
	// The error information.
	//
	// >  This parameter is returned only if the value of `Status` is `Failed`.
	ErrorInfo *ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty" type:"Struct"`
	// The end time of the check item. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-12-13T03:34:02Z
	EvaluationTime *string `json:"EvaluationTime,omitempty" xml:"EvaluationTime,omitempty"`
	// The ID of the check item.
	//
	// example:
	//
	// r7xdcu****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The checked resources.
	ResourcesSummary *ListEvaluationResultsResponseBodyResultsMetricResultsResourcesSummary `json:"ResourcesSummary,omitempty" xml:"ResourcesSummary,omitempty" type:"Struct"`
	// The rate of the non-compliant resources.
	//
	// example:
	//
	// 0.67
	Result *float64 `json:"Result,omitempty" xml:"Result,omitempty"`
	// The risk level. Valid values:
	//
	// 	- Error: high risk
	//
	// 	- Warning: medium risk
	//
	// 	- None: no risk
	//
	// example:
	//
	// Error
	Risk *string `json:"Risk,omitempty" xml:"Risk,omitempty"`
	// The status of the check item. Valid values:
	//
	// 	- Running: The check is in progress.
	//
	// 	- Finished: The check is complete.
	//
	// 	- failed: The check fails.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListEvaluationResultsResponseBodyResultsMetricResults) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationResultsResponseBodyResultsMetricResults) GoString() string {
	return s.String()
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) SetErrorInfo(v *ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo) *ListEvaluationResultsResponseBodyResultsMetricResults {
	s.ErrorInfo = v
	return s
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) SetEvaluationTime(v string) *ListEvaluationResultsResponseBodyResultsMetricResults {
	s.EvaluationTime = &v
	return s
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) SetId(v string) *ListEvaluationResultsResponseBodyResultsMetricResults {
	s.Id = &v
	return s
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) SetResourcesSummary(v *ListEvaluationResultsResponseBodyResultsMetricResultsResourcesSummary) *ListEvaluationResultsResponseBodyResultsMetricResults {
	s.ResourcesSummary = v
	return s
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) SetResult(v float64) *ListEvaluationResultsResponseBodyResultsMetricResults {
	s.Result = &v
	return s
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) SetRisk(v string) *ListEvaluationResultsResponseBodyResultsMetricResults {
	s.Risk = &v
	return s
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResults) SetStatus(v string) *ListEvaluationResultsResponseBodyResultsMetricResults {
	s.Status = &v
	return s
}

type ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo struct {
	// The error code.
	//
	// example:
	//
	// EcsInsightEnableFailed
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unable to enable ECS Insight due to a server error.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo) GoString() string {
	return s.String()
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo) SetCode(v string) *ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo {
	s.Code = &v
	return s
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo) SetMessage(v string) *ListEvaluationResultsResponseBodyResultsMetricResultsErrorInfo {
	s.Message = &v
	return s
}

type ListEvaluationResultsResponseBodyResultsMetricResultsResourcesSummary struct {
	// The number of non-compliant resources.
	//
	// example:
	//
	// 2
	NonCompliant *int32 `json:"NonCompliant,omitempty" xml:"NonCompliant,omitempty"`
}

func (s ListEvaluationResultsResponseBodyResultsMetricResultsResourcesSummary) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationResultsResponseBodyResultsMetricResultsResourcesSummary) GoString() string {
	return s.String()
}

func (s *ListEvaluationResultsResponseBodyResultsMetricResultsResourcesSummary) SetNonCompliant(v int32) *ListEvaluationResultsResponseBodyResultsMetricResultsResourcesSummary {
	s.NonCompliant = &v
	return s
}

type ListEvaluationResultsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEvaluationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEvaluationResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationResultsResponse) GoString() string {
	return s.String()
}

func (s *ListEvaluationResultsResponse) SetHeaders(v map[string]*string) *ListEvaluationResultsResponse {
	s.Headers = v
	return s
}

func (s *ListEvaluationResultsResponse) SetStatusCode(v int32) *ListEvaluationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEvaluationResultsResponse) SetBody(v *ListEvaluationResultsResponseBody) *ListEvaluationResultsResponse {
	s.Body = v
	return s
}

type ListEvaluationScoreHistoryRequest struct {
	// The Alibaba Cloud account ID of the member. This parameter takes effect only when a multi-account governance maturity check is performed.
	//
	// example:
	//
	// 176618589410****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The end of the time range to query. Specify the time in the YYYY-MM-DD format.
	//
	// By default, the historical scores that were generated in the seven days before the current date are queried.
	//
	// example:
	//
	// 2024-07-11
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the YYYY-MM-DD format.
	//
	// You can query the historical scores within the previous 180 days.
	//
	// example:
	//
	// 2024-06-11
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ListEvaluationScoreHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationScoreHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListEvaluationScoreHistoryRequest) SetAccountId(v int64) *ListEvaluationScoreHistoryRequest {
	s.AccountId = &v
	return s
}

func (s *ListEvaluationScoreHistoryRequest) SetEndDate(v string) *ListEvaluationScoreHistoryRequest {
	s.EndDate = &v
	return s
}

func (s *ListEvaluationScoreHistoryRequest) SetRegionId(v string) *ListEvaluationScoreHistoryRequest {
	s.RegionId = &v
	return s
}

func (s *ListEvaluationScoreHistoryRequest) SetStartDate(v string) *ListEvaluationScoreHistoryRequest {
	s.StartDate = &v
	return s
}

type ListEvaluationScoreHistoryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AC9BD94C-D20C-4D27-88D4-89E8D75C051B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The historical scores.
	ScoreHistory *ListEvaluationScoreHistoryResponseBodyScoreHistory `json:"ScoreHistory,omitempty" xml:"ScoreHistory,omitempty" type:"Struct"`
}

func (s ListEvaluationScoreHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationScoreHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListEvaluationScoreHistoryResponseBody) SetRequestId(v string) *ListEvaluationScoreHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEvaluationScoreHistoryResponseBody) SetScoreHistory(v *ListEvaluationScoreHistoryResponseBodyScoreHistory) *ListEvaluationScoreHistoryResponseBody {
	s.ScoreHistory = v
	return s
}

type ListEvaluationScoreHistoryResponseBodyScoreHistory struct {
	// The historical scores.
	TotalScoreHistory []*ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory `json:"TotalScoreHistory,omitempty" xml:"TotalScoreHistory,omitempty" type:"Repeated"`
}

func (s ListEvaluationScoreHistoryResponseBodyScoreHistory) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationScoreHistoryResponseBodyScoreHistory) GoString() string {
	return s.String()
}

func (s *ListEvaluationScoreHistoryResponseBodyScoreHistory) SetTotalScoreHistory(v []*ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory) *ListEvaluationScoreHistoryResponseBodyScoreHistory {
	s.TotalScoreHistory = v
	return s
}

type ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory struct {
	// The time when the score was generated. The time is in UTC.
	//
	// example:
	//
	// 2024-06-30T03:34:02Z
	EvaluationTime *string `json:"EvaluationTime,omitempty" xml:"EvaluationTime,omitempty"`
	// The score.
	//
	// Valid values: 0 to 1.
	//
	// example:
	//
	// 0.6753
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory) GoString() string {
	return s.String()
}

func (s *ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory) SetEvaluationTime(v string) *ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory {
	s.EvaluationTime = &v
	return s
}

func (s *ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory) SetScore(v float64) *ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory {
	s.Score = &v
	return s
}

type ListEvaluationScoreHistoryResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEvaluationScoreHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEvaluationScoreHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEvaluationScoreHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListEvaluationScoreHistoryResponse) SetHeaders(v map[string]*string) *ListEvaluationScoreHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListEvaluationScoreHistoryResponse) SetStatusCode(v int32) *ListEvaluationScoreHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEvaluationScoreHistoryResponse) SetBody(v *ListEvaluationScoreHistoryResponseBody) *ListEvaluationScoreHistoryResponse {
	s.Body = v
	return s
}

type RunEvaluationRequest struct {
	// The Alibaba Cloud account ID of the member. This parameter takes effect only when a multi-account governance maturity check is performed.
	//
	// example:
	//
	// 176618589410****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The IDs of the check items to be checked.
	MetricIds []*string `json:"MetricIds,omitempty" xml:"MetricIds,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The check range of the governance maturity check. Valid values:
	//
	// 	- Account (default): A single-account governance maturity check is performed to check only the Alibaba Cloud account that you use to access Cloud Governance Center.
	//
	// 	- ResourceDirectory: A multi-account governance maturity check is performed to check all members within a resource directory. Before you perform a multi-account governance maturity check, you must enable the multi-account governance maturity check feature.
	//
	// example:
	//
	// ResourceDirectory
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s RunEvaluationRequest) String() string {
	return tea.Prettify(s)
}

func (s RunEvaluationRequest) GoString() string {
	return s.String()
}

func (s *RunEvaluationRequest) SetAccountId(v int64) *RunEvaluationRequest {
	s.AccountId = &v
	return s
}

func (s *RunEvaluationRequest) SetMetricIds(v []*string) *RunEvaluationRequest {
	s.MetricIds = v
	return s
}

func (s *RunEvaluationRequest) SetRegionId(v string) *RunEvaluationRequest {
	s.RegionId = &v
	return s
}

func (s *RunEvaluationRequest) SetScope(v string) *RunEvaluationRequest {
	s.Scope = &v
	return s
}

type RunEvaluationShrinkRequest struct {
	// The Alibaba Cloud account ID of the member. This parameter takes effect only when a multi-account governance maturity check is performed.
	//
	// example:
	//
	// 176618589410****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The IDs of the check items to be checked.
	MetricIdsShrink *string `json:"MetricIds,omitempty" xml:"MetricIds,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The check range of the governance maturity check. Valid values:
	//
	// 	- Account (default): A single-account governance maturity check is performed to check only the Alibaba Cloud account that you use to access Cloud Governance Center.
	//
	// 	- ResourceDirectory: A multi-account governance maturity check is performed to check all members within a resource directory. Before you perform a multi-account governance maturity check, you must enable the multi-account governance maturity check feature.
	//
	// example:
	//
	// ResourceDirectory
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s RunEvaluationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunEvaluationShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunEvaluationShrinkRequest) SetAccountId(v int64) *RunEvaluationShrinkRequest {
	s.AccountId = &v
	return s
}

func (s *RunEvaluationShrinkRequest) SetMetricIdsShrink(v string) *RunEvaluationShrinkRequest {
	s.MetricIdsShrink = &v
	return s
}

func (s *RunEvaluationShrinkRequest) SetRegionId(v string) *RunEvaluationShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RunEvaluationShrinkRequest) SetScope(v string) *RunEvaluationShrinkRequest {
	s.Scope = &v
	return s
}

type RunEvaluationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2D3E2A3A-F2B8-578D-9659-3195F94A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunEvaluationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunEvaluationResponseBody) GoString() string {
	return s.String()
}

func (s *RunEvaluationResponseBody) SetRequestId(v string) *RunEvaluationResponseBody {
	s.RequestId = &v
	return s
}

type RunEvaluationResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunEvaluationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunEvaluationResponse) String() string {
	return tea.Prettify(s)
}

func (s RunEvaluationResponse) GoString() string {
	return s.String()
}

func (s *RunEvaluationResponse) SetHeaders(v map[string]*string) *RunEvaluationResponse {
	s.Headers = v
	return s
}

func (s *RunEvaluationResponse) SetStatusCode(v int32) *RunEvaluationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunEvaluationResponse) SetBody(v *RunEvaluationResponseBody) *RunEvaluationResponse {
	s.Body = v
	return s
}

type UpdateAccountFactoryBaselineRequest struct {
	// The baseline ID.
	//
	// example:
	//
	// afb-bp1pq3emlkt27vsj****
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The baseline items.
	//
	// You can call the [ListAccountFactoryBaselineItems](~~ListAccountFactoryBaselineItems~~) operation to query a list of baseline items supported by the account factory in Cloud Governance Center.
	BaselineItems []*UpdateAccountFactoryBaselineRequestBaselineItems `json:"BaselineItems,omitempty" xml:"BaselineItems,omitempty" type:"Repeated"`
	// The name of the baseline.
	BaselineName *string `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	// The description of the baseline.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The region ID.
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
	// The configurations of the baseline item. The value of this parameter is a JSON string.
	//
	// example:
	//
	// {\\"EnabledServices\\":[\\"CEN_TR\\",\\"CDT\\",\\"CMS\\",\\"KMS\\"]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The name of the baseline item.
	//
	// example:
	//
	// ACS-BP_ACCOUNT_FACTORY_VPC
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the baseline item.
	//
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
	// The request ID.
	//
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
	client.SignatureAlgorithm = tea.String("v2")
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
// Applies an account baseline to multiple existing resource accounts at a time.
//
// Description:
//
// You can call this operation to apply an account baseline to existing resource accounts.
//
// Accounts are enrolled in the account factory in asynchronous mode. After a resource account is created, an account baseline is applied to the account. You can call the [GetEnrolledAccount](https://help.aliyun.com/document_detail/609062.html) operation to query the details of the account enrolled in the account factory and check whether the account baseline is applied to the account.
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &BatchEnrollAccountsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &BatchEnrollAccountsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Applies an account baseline to multiple existing resource accounts at a time.
//
// Description:
//
// You can call this operation to apply an account baseline to existing resource accounts.
//
// Accounts are enrolled in the account factory in asynchronous mode. After a resource account is created, an account baseline is applied to the account. You can call the [GetEnrolledAccount](https://help.aliyun.com/document_detail/609062.html) operation to query the details of the account enrolled in the account factory and check whether the account baseline is applied to the account.
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
// Creates a baseline of the account factory.
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateAccountFactoryBaselineResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateAccountFactoryBaselineResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a baseline of the account factory.
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
// Deletes an account factory baseline.
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteAccountFactoryBaselineResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteAccountFactoryBaselineResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes an account factory baseline.
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
// @param tmpReq - EnrollAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnrollAccountResponse
func (client *Client) EnrollAccountWithOptions(tmpReq *EnrollAccountRequest, runtime *util.RuntimeOptions) (_result *EnrollAccountResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &EnrollAccountShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("Tag"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		query["Tag"] = request.TagShrink
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &EnrollAccountResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &EnrollAccountResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetAccountFactoryBaselineResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetAccountFactoryBaselineResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetEnrolledAccountResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetEnrolledAccountResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// Queries a list of baseline items that are supported by the account factory of Cloud Governance Center (CGC).
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListAccountFactoryBaselineItemsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListAccountFactoryBaselineItemsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries a list of baseline items that are supported by the account factory of Cloud Governance Center (CGC).
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListAccountFactoryBaselinesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListAccountFactoryBaselinesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListEnrolledAccountsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListEnrolledAccountsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// Queries all available information about check items in a governance maturity check, including the name, ID, description, stage, resource metadata, and fixing guide.
//
// @param request - ListEvaluationMetadataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEvaluationMetadataResponse
func (client *Client) ListEvaluationMetadataWithOptions(request *ListEvaluationMetadataRequest, runtime *util.RuntimeOptions) (_result *ListEvaluationMetadataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEvaluationMetadata"),
		Version:     tea.String("2021-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListEvaluationMetadataResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListEvaluationMetadataResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries all available information about check items in a governance maturity check, including the name, ID, description, stage, resource metadata, and fixing guide.
//
// @param request - ListEvaluationMetadataRequest
//
// @return ListEvaluationMetadataResponse
func (client *Client) ListEvaluationMetadata(request *ListEvaluationMetadataRequest) (_result *ListEvaluationMetadataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEvaluationMetadataResponse{}
	_body, _err := client.ListEvaluationMetadataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the non-compliant resource information of a check item, including the name, ID, category, type, region, and related metadata of non-compliant resources.
//
// @param request - ListEvaluationMetricDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEvaluationMetricDetailsResponse
func (client *Client) ListEvaluationMetricDetailsWithOptions(request *ListEvaluationMetricDetailsRequest, runtime *util.RuntimeOptions) (_result *ListEvaluationMetricDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEvaluationMetricDetails"),
		Version:     tea.String("2021-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListEvaluationMetricDetailsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListEvaluationMetricDetailsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the non-compliant resource information of a check item, including the name, ID, category, type, region, and related metadata of non-compliant resources.
//
// @param request - ListEvaluationMetricDetailsRequest
//
// @return ListEvaluationMetricDetailsResponse
func (client *Client) ListEvaluationMetricDetails(request *ListEvaluationMetricDetailsRequest) (_result *ListEvaluationMetricDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEvaluationMetricDetailsResponse{}
	_body, _err := client.ListEvaluationMetricDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the result and status of a governance check.
//
// @param request - ListEvaluationResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEvaluationResultsResponse
func (client *Client) ListEvaluationResultsWithOptions(request *ListEvaluationResultsRequest, runtime *util.RuntimeOptions) (_result *ListEvaluationResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEvaluationResults"),
		Version:     tea.String("2021-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListEvaluationResultsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListEvaluationResultsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the result and status of a governance check.
//
// @param request - ListEvaluationResultsRequest
//
// @return ListEvaluationResultsResponse
func (client *Client) ListEvaluationResults(request *ListEvaluationResultsRequest) (_result *ListEvaluationResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEvaluationResultsResponse{}
	_body, _err := client.ListEvaluationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the historical scores of a governance maturity check.
//
// @param request - ListEvaluationScoreHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEvaluationScoreHistoryResponse
func (client *Client) ListEvaluationScoreHistoryWithOptions(request *ListEvaluationScoreHistoryRequest, runtime *util.RuntimeOptions) (_result *ListEvaluationScoreHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEvaluationScoreHistory"),
		Version:     tea.String("2021-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListEvaluationScoreHistoryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListEvaluationScoreHistoryResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the historical scores of a governance maturity check.
//
// @param request - ListEvaluationScoreHistoryRequest
//
// @return ListEvaluationScoreHistoryResponse
func (client *Client) ListEvaluationScoreHistory(request *ListEvaluationScoreHistoryRequest) (_result *ListEvaluationScoreHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEvaluationScoreHistoryResponse{}
	_body, _err := client.ListEvaluationScoreHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a governance maturity check.
//
// @param tmpReq - RunEvaluationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunEvaluationResponse
func (client *Client) RunEvaluationWithOptions(tmpReq *RunEvaluationRequest, runtime *util.RuntimeOptions) (_result *RunEvaluationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunEvaluationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.MetricIds)) {
		request.MetricIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MetricIds, tea.String("MetricIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.MetricIdsShrink)) {
		query["MetricIds"] = request.MetricIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RunEvaluation"),
		Version:     tea.String("2021-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunEvaluationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunEvaluationResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Performs a governance maturity check.
//
// @param request - RunEvaluationRequest
//
// @return RunEvaluationResponse
func (client *Client) RunEvaluation(request *RunEvaluationRequest) (_result *RunEvaluationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunEvaluationResponse{}
	_body, _err := client.RunEvaluationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a baseline of the account factory.
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateAccountFactoryBaselineResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateAccountFactoryBaselineResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Updates a baseline of the account factory.
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
