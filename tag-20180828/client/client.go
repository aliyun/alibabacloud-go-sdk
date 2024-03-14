// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AttachPolicyRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the tag policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the object.
	//
	// >  If you use the Tag Policy feature in single-account mode, this parameter is optional. If you use the Tag Policy feature in multi-account mode, this feature is required.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the object. Valid values:
	//
	// *   USER: the current logon account. This value is available if you use the Tag Policy feature in single-account mode.
	// *   ROOT: the Root folder in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	// *   FOLDER: a folder other than the Root folder in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	// *   ACCOUNT: a member in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// >  If you use the Tag Policy feature in single-account mode, this parameter is optional. If you use the Tag Policy feature in multi-account mode, this feature is required. The value of this parameter is not case-sensitive.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s AttachPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachPolicyRequest) GoString() string {
	return s.String()
}

func (s *AttachPolicyRequest) SetOwnerAccount(v string) *AttachPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AttachPolicyRequest) SetOwnerId(v int64) *AttachPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachPolicyRequest) SetPolicyId(v string) *AttachPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *AttachPolicyRequest) SetRegionId(v string) *AttachPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *AttachPolicyRequest) SetResourceOwnerAccount(v string) *AttachPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachPolicyRequest) SetTargetId(v string) *AttachPolicyRequest {
	s.TargetId = &v
	return s
}

func (s *AttachPolicyRequest) SetTargetType(v string) *AttachPolicyRequest {
	s.TargetType = &v
	return s
}

type AttachPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *AttachPolicyResponseBody) SetRequestId(v string) *AttachPolicyResponseBody {
	s.RequestId = &v
	return s
}

type AttachPolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachPolicyResponse) GoString() string {
	return s.String()
}

func (s *AttachPolicyResponse) SetHeaders(v map[string]*string) *AttachPolicyResponse {
	s.Headers = v
	return s
}

func (s *AttachPolicyResponse) SetStatusCode(v int32) *AttachPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachPolicyResponse) SetBody(v *AttachPolicyResponseBody) *AttachPolicyResponse {
	s.Body = v
	return s
}

type CheckCreatedByEnabledRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckCreatedByEnabledRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckCreatedByEnabledRequest) GoString() string {
	return s.String()
}

func (s *CheckCreatedByEnabledRequest) SetOwnerAccount(v string) *CheckCreatedByEnabledRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckCreatedByEnabledRequest) SetOwnerId(v int64) *CheckCreatedByEnabledRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckCreatedByEnabledRequest) SetRegionId(v string) *CheckCreatedByEnabledRequest {
	s.RegionId = &v
	return s
}

func (s *CheckCreatedByEnabledRequest) SetResourceOwnerAccount(v string) *CheckCreatedByEnabledRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckCreatedByEnabledRequest) SetResourceOwnerId(v string) *CheckCreatedByEnabledRequest {
	s.ResourceOwnerId = &v
	return s
}

type CheckCreatedByEnabledResponseBody struct {
	OpenStatus *bool   `json:"OpenStatus,omitempty" xml:"OpenStatus,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckCreatedByEnabledResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckCreatedByEnabledResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCreatedByEnabledResponseBody) SetOpenStatus(v bool) *CheckCreatedByEnabledResponseBody {
	s.OpenStatus = &v
	return s
}

func (s *CheckCreatedByEnabledResponseBody) SetRequestId(v string) *CheckCreatedByEnabledResponseBody {
	s.RequestId = &v
	return s
}

type CheckCreatedByEnabledResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckCreatedByEnabledResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckCreatedByEnabledResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckCreatedByEnabledResponse) GoString() string {
	return s.String()
}

func (s *CheckCreatedByEnabledResponse) SetHeaders(v map[string]*string) *CheckCreatedByEnabledResponse {
	s.Headers = v
	return s
}

func (s *CheckCreatedByEnabledResponse) SetStatusCode(v int32) *CheckCreatedByEnabledResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckCreatedByEnabledResponse) SetBody(v *CheckCreatedByEnabledResponseBody) *CheckCreatedByEnabledResponse {
	s.Body = v
	return s
}

type CloseCreatedByRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CloseCreatedByRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseCreatedByRequest) GoString() string {
	return s.String()
}

func (s *CloseCreatedByRequest) SetOwnerAccount(v string) *CloseCreatedByRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CloseCreatedByRequest) SetOwnerId(v int64) *CloseCreatedByRequest {
	s.OwnerId = &v
	return s
}

func (s *CloseCreatedByRequest) SetRegionId(v string) *CloseCreatedByRequest {
	s.RegionId = &v
	return s
}

func (s *CloseCreatedByRequest) SetResourceOwnerAccount(v string) *CloseCreatedByRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloseCreatedByRequest) SetResourceOwnerId(v string) *CloseCreatedByRequest {
	s.ResourceOwnerId = &v
	return s
}

type CloseCreatedByResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloseCreatedByResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseCreatedByResponseBody) GoString() string {
	return s.String()
}

func (s *CloseCreatedByResponseBody) SetRequestId(v string) *CloseCreatedByResponseBody {
	s.RequestId = &v
	return s
}

type CloseCreatedByResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseCreatedByResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseCreatedByResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseCreatedByResponse) GoString() string {
	return s.String()
}

func (s *CloseCreatedByResponse) SetHeaders(v map[string]*string) *CloseCreatedByResponse {
	s.Headers = v
	return s
}

func (s *CloseCreatedByResponse) SetStatusCode(v int32) *CloseCreatedByResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseCreatedByResponse) SetBody(v *CloseCreatedByResponseBody) *CloseCreatedByResponse {
	s.Body = v
	return s
}

type CreatePolicyRequest struct {
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// *   false (default): performs a dry run and performs the actual request.
	// *   true: performs only a dry run.
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The document of the tag policy.
	//
	// For more information about the syntax of a tag policy, see [Syntax of a tag policy](~~417436~~).
	PolicyContent *string `json:"PolicyContent,omitempty" xml:"PolicyContent,omitempty"`
	// The description of the tag policy.
	//
	// The description must be 0 to 512 characters in length.
	PolicyDesc *string `json:"PolicyDesc,omitempty" xml:"PolicyDesc,omitempty"`
	// The name of the tag policy.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, and underscores (\_).
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The mode of the Tag Policy feature. Valid values:
	//
	// *   USER: single-account mode. Set the value to USER if you use an Alibaba Cloud account or a member of a resource directory to call this API operation to create a tag policy for the Alibaba Cloud account or member.
	// *   RD: multi-account mode. Set the value to RD if you use the management account of a resource directory to call this API operation to create a tag policy for the resource directory.
	//
	// For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s CreatePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequest) SetDryRun(v bool) *CreatePolicyRequest {
	s.DryRun = &v
	return s
}

func (s *CreatePolicyRequest) SetOwnerAccount(v string) *CreatePolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreatePolicyRequest) SetOwnerId(v int64) *CreatePolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *CreatePolicyRequest) SetPolicyContent(v string) *CreatePolicyRequest {
	s.PolicyContent = &v
	return s
}

func (s *CreatePolicyRequest) SetPolicyDesc(v string) *CreatePolicyRequest {
	s.PolicyDesc = &v
	return s
}

func (s *CreatePolicyRequest) SetPolicyName(v string) *CreatePolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyRequest) SetRegionId(v string) *CreatePolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePolicyRequest) SetResourceOwnerAccount(v string) *CreatePolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreatePolicyRequest) SetUserType(v string) *CreatePolicyRequest {
	s.UserType = &v
	return s
}

type CreatePolicyResponseBody struct {
	// The ID of the tag policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the tag policy.
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponseBody) SetPolicyId(v string) *CreatePolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *CreatePolicyResponseBody) SetPolicyName(v string) *CreatePolicyResponseBody {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyResponseBody) SetRequestId(v string) *CreatePolicyResponseBody {
	s.RequestId = &v
	return s
}

type CreatePolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePolicyResponse) GoString() string {
	return s.String()
}

func (s *CreatePolicyResponse) SetHeaders(v map[string]*string) *CreatePolicyResponse {
	s.Headers = v
	return s
}

func (s *CreatePolicyResponse) SetStatusCode(v int32) *CreatePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolicyResponse) SetBody(v *CreatePolicyResponseBody) *CreatePolicyResponse {
	s.Body = v
	return s
}

type CreateTagsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// > Only `cn-hangzhou` is supported.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The information about the tags.
	TagKeyValueParamList []*CreateTagsRequestTagKeyValueParamList `json:"TagKeyValueParamList,omitempty" xml:"TagKeyValueParamList,omitempty" type:"Repeated"`
}

func (s CreateTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTagsRequest) GoString() string {
	return s.String()
}

func (s *CreateTagsRequest) SetOwnerAccount(v string) *CreateTagsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTagsRequest) SetOwnerId(v int64) *CreateTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTagsRequest) SetRegionId(v string) *CreateTagsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTagsRequest) SetResourceOwnerAccount(v string) *CreateTagsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTagsRequest) SetTagKeyValueParamList(v []*CreateTagsRequestTagKeyValueParamList) *CreateTagsRequest {
	s.TagKeyValueParamList = v
	return s
}

type CreateTagsRequestTagKeyValueParamList struct {
	// The description of the key for tag N.
	//
	// Valid values of N: 1 to 10.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The key of tag N.
	//
	// The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// Valid values of N: 1 to 10.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The information about the tag value.
	TagValueParamList []*CreateTagsRequestTagKeyValueParamListTagValueParamList `json:"TagValueParamList,omitempty" xml:"TagValueParamList,omitempty" type:"Repeated"`
}

func (s CreateTagsRequestTagKeyValueParamList) String() string {
	return tea.Prettify(s)
}

func (s CreateTagsRequestTagKeyValueParamList) GoString() string {
	return s.String()
}

func (s *CreateTagsRequestTagKeyValueParamList) SetDescription(v string) *CreateTagsRequestTagKeyValueParamList {
	s.Description = &v
	return s
}

func (s *CreateTagsRequestTagKeyValueParamList) SetKey(v string) *CreateTagsRequestTagKeyValueParamList {
	s.Key = &v
	return s
}

func (s *CreateTagsRequestTagKeyValueParamList) SetTagValueParamList(v []*CreateTagsRequestTagKeyValueParamListTagValueParamList) *CreateTagsRequestTagKeyValueParamList {
	s.TagValueParamList = v
	return s
}

type CreateTagsRequestTagKeyValueParamListTagValueParamList struct {
	// The description of the value for tag N.
	//
	// Valid values of N: 1 to 10.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The value of tag N.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// Valid values of N: 1 to 10.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTagsRequestTagKeyValueParamListTagValueParamList) String() string {
	return tea.Prettify(s)
}

func (s CreateTagsRequestTagKeyValueParamListTagValueParamList) GoString() string {
	return s.String()
}

func (s *CreateTagsRequestTagKeyValueParamListTagValueParamList) SetDescription(v string) *CreateTagsRequestTagKeyValueParamListTagValueParamList {
	s.Description = &v
	return s
}

func (s *CreateTagsRequestTagKeyValueParamListTagValueParamList) SetValue(v string) *CreateTagsRequestTagKeyValueParamListTagValueParamList {
	s.Value = &v
	return s
}

type CreateTagsResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTagsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTagsResponseBody) SetRequestId(v string) *CreateTagsResponseBody {
	s.RequestId = &v
	return s
}

type CreateTagsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTagsResponse) GoString() string {
	return s.String()
}

func (s *CreateTagsResponse) SetHeaders(v map[string]*string) *CreateTagsResponse {
	s.Headers = v
	return s
}

func (s *CreateTagsResponse) SetStatusCode(v int32) *CreateTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTagsResponse) SetBody(v *CreateTagsResponseBody) *CreateTagsResponse {
	s.Body = v
	return s
}

type DeletePolicyRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the tag policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DeletePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyRequest) SetOwnerAccount(v string) *DeletePolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeletePolicyRequest) SetOwnerId(v int64) *DeletePolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DeletePolicyRequest) SetPolicyId(v string) *DeletePolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DeletePolicyRequest) SetRegionId(v string) *DeletePolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePolicyRequest) SetResourceOwnerAccount(v string) *DeletePolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type DeletePolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicyResponseBody) SetRequestId(v string) *DeletePolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeletePolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePolicyResponse) GoString() string {
	return s.String()
}

func (s *DeletePolicyResponse) SetHeaders(v map[string]*string) *DeletePolicyResponse {
	s.Headers = v
	return s
}

func (s *DeletePolicyResponse) SetStatusCode(v int32) *DeletePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolicyResponse) SetBody(v *DeletePolicyResponseBody) *DeletePolicyResponse {
	s.Body = v
	return s
}

type DeleteTagRequest struct {
	// The tag key.
	//
	// If no tag value is associated with a tag key, you can specify the `Key` parameter without specifying the Value parameter to delete the tag key. Otherwise, you must specify both the `Key` and `Value` parameters to delete a preset tag.
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// >  Only `cn-hangzhou` is supported.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DeleteTagRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagRequest) GoString() string {
	return s.String()
}

func (s *DeleteTagRequest) SetKey(v string) *DeleteTagRequest {
	s.Key = &v
	return s
}

func (s *DeleteTagRequest) SetOwnerAccount(v string) *DeleteTagRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteTagRequest) SetOwnerId(v int64) *DeleteTagRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteTagRequest) SetRegionId(v string) *DeleteTagRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTagRequest) SetResourceOwnerAccount(v string) *DeleteTagRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteTagRequest) SetValue(v string) *DeleteTagRequest {
	s.Value = &v
	return s
}

type DeleteTagResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTagResponseBody) SetRequestId(v string) *DeleteTagResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTagResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTagResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagResponse) GoString() string {
	return s.String()
}

func (s *DeleteTagResponse) SetHeaders(v map[string]*string) *DeleteTagResponse {
	s.Headers = v
	return s
}

func (s *DeleteTagResponse) SetStatusCode(v int32) *DeleteTagResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTagResponse) SetBody(v *DeleteTagResponseBody) *DeleteTagResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// The supported natural language. Valid values:
	//
	// *   zh-CN: Chinese (default value)
	// *   en-US: English
	// *   ja: Japanese
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerAccount(v string) *DescribeRegionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerId(v int64) *DescribeRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerAccount(v string) *DescribeRegionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerId(v int64) *DescribeRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// The information of the regions.
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	// The name of the region.
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint of the Tag service in the region.
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DetachPolicyRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the tag policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the object.
	//
	// >  If you use the Tag Policy feature in single-account mode, this parameter is optional. If you use the Tag Policy feature in multi-account mode, this feature is required.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the object. Valid values:
	//
	// *   USER: the current logon account. This value is available if you use the Tag Policy feature in single-account mode.
	// *   ROOT: the Root folder in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	// *   FOLDER: a folder other than the Root folder in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	// *   ACCOUNT: a member in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// >  If you use the Tag Policy feature in single-account mode, this parameter is optional. If you use the Tag Policy feature in multi-account mode, this feature is required. The value of this parameter is not case-sensitive.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s DetachPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachPolicyRequest) GoString() string {
	return s.String()
}

func (s *DetachPolicyRequest) SetOwnerAccount(v string) *DetachPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DetachPolicyRequest) SetOwnerId(v int64) *DetachPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachPolicyRequest) SetPolicyId(v string) *DetachPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DetachPolicyRequest) SetRegionId(v string) *DetachPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DetachPolicyRequest) SetResourceOwnerAccount(v string) *DetachPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachPolicyRequest) SetTargetId(v string) *DetachPolicyRequest {
	s.TargetId = &v
	return s
}

func (s *DetachPolicyRequest) SetTargetType(v string) *DetachPolicyRequest {
	s.TargetType = &v
	return s
}

type DetachPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DetachPolicyResponseBody) SetRequestId(v string) *DetachPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DetachPolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachPolicyResponse) GoString() string {
	return s.String()
}

func (s *DetachPolicyResponse) SetHeaders(v map[string]*string) *DetachPolicyResponse {
	s.Headers = v
	return s
}

func (s *DetachPolicyResponse) SetStatusCode(v int32) *DetachPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachPolicyResponse) SetBody(v *DetachPolicyResponseBody) *DetachPolicyResponse {
	s.Body = v
	return s
}

type DisablePolicyTypeRequest struct {
	OpenType             *string `json:"OpenType,omitempty" xml:"OpenType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserType             *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DisablePolicyTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DisablePolicyTypeRequest) GoString() string {
	return s.String()
}

func (s *DisablePolicyTypeRequest) SetOpenType(v string) *DisablePolicyTypeRequest {
	s.OpenType = &v
	return s
}

func (s *DisablePolicyTypeRequest) SetOwnerAccount(v string) *DisablePolicyTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DisablePolicyTypeRequest) SetOwnerId(v int64) *DisablePolicyTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *DisablePolicyTypeRequest) SetRegionId(v string) *DisablePolicyTypeRequest {
	s.RegionId = &v
	return s
}

func (s *DisablePolicyTypeRequest) SetResourceOwnerAccount(v string) *DisablePolicyTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisablePolicyTypeRequest) SetResourceOwnerId(v string) *DisablePolicyTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisablePolicyTypeRequest) SetUserType(v string) *DisablePolicyTypeRequest {
	s.UserType = &v
	return s
}

type DisablePolicyTypeResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisablePolicyTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisablePolicyTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DisablePolicyTypeResponseBody) SetRequestId(v string) *DisablePolicyTypeResponseBody {
	s.RequestId = &v
	return s
}

type DisablePolicyTypeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisablePolicyTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisablePolicyTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DisablePolicyTypeResponse) GoString() string {
	return s.String()
}

func (s *DisablePolicyTypeResponse) SetHeaders(v map[string]*string) *DisablePolicyTypeResponse {
	s.Headers = v
	return s
}

func (s *DisablePolicyTypeResponse) SetStatusCode(v int32) *DisablePolicyTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DisablePolicyTypeResponse) SetBody(v *DisablePolicyTypeResponseBody) *DisablePolicyTypeResponse {
	s.Body = v
	return s
}

type EnablePolicyTypeRequest struct {
	OpenType             *string `json:"OpenType,omitempty" xml:"OpenType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserType             *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s EnablePolicyTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s EnablePolicyTypeRequest) GoString() string {
	return s.String()
}

func (s *EnablePolicyTypeRequest) SetOpenType(v string) *EnablePolicyTypeRequest {
	s.OpenType = &v
	return s
}

func (s *EnablePolicyTypeRequest) SetOwnerAccount(v string) *EnablePolicyTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *EnablePolicyTypeRequest) SetOwnerId(v int64) *EnablePolicyTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *EnablePolicyTypeRequest) SetRegionId(v string) *EnablePolicyTypeRequest {
	s.RegionId = &v
	return s
}

func (s *EnablePolicyTypeRequest) SetResourceOwnerAccount(v string) *EnablePolicyTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *EnablePolicyTypeRequest) SetResourceOwnerId(v string) *EnablePolicyTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *EnablePolicyTypeRequest) SetUserType(v string) *EnablePolicyTypeRequest {
	s.UserType = &v
	return s
}

type EnablePolicyTypeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnablePolicyTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnablePolicyTypeResponseBody) GoString() string {
	return s.String()
}

func (s *EnablePolicyTypeResponseBody) SetRequestId(v string) *EnablePolicyTypeResponseBody {
	s.RequestId = &v
	return s
}

type EnablePolicyTypeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnablePolicyTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnablePolicyTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s EnablePolicyTypeResponse) GoString() string {
	return s.String()
}

func (s *EnablePolicyTypeResponse) SetHeaders(v map[string]*string) *EnablePolicyTypeResponse {
	s.Headers = v
	return s
}

func (s *EnablePolicyTypeResponse) SetStatusCode(v int32) *EnablePolicyTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *EnablePolicyTypeResponse) SetBody(v *EnablePolicyTypeResponseBody) *EnablePolicyTypeResponse {
	s.Body = v
	return s
}

type GenerateConfigRuleReportRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the object.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the object. Valid values:
	//
	// *   USER: the current logon account. This value is available if you use the Tag Policy feature in single-account mode.
	// *   ROOT: the Root folder in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	// *   FOLDER: a folder other than the Root folder in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	// *   ACCOUNT: a member in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// >  The value of this parameter is not case-sensitive.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The mode of the Tag Policy feature. Valid values:
	//
	// *   USER: single-account mode
	// *   RD: multi-account mode
	//
	// For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
	//
	// >  This parameter is required if the management account of your resource directory is used to enable the Tag Policy feature in both single-account mode and multi-account mode. The value of this parameter is not case-sensitive.
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s GenerateConfigRuleReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateConfigRuleReportRequest) GoString() string {
	return s.String()
}

func (s *GenerateConfigRuleReportRequest) SetOwnerAccount(v string) *GenerateConfigRuleReportRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GenerateConfigRuleReportRequest) SetOwnerId(v int64) *GenerateConfigRuleReportRequest {
	s.OwnerId = &v
	return s
}

func (s *GenerateConfigRuleReportRequest) SetRegionId(v string) *GenerateConfigRuleReportRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateConfigRuleReportRequest) SetResourceOwnerAccount(v string) *GenerateConfigRuleReportRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GenerateConfigRuleReportRequest) SetTargetId(v string) *GenerateConfigRuleReportRequest {
	s.TargetId = &v
	return s
}

func (s *GenerateConfigRuleReportRequest) SetTargetType(v string) *GenerateConfigRuleReportRequest {
	s.TargetType = &v
	return s
}

func (s *GenerateConfigRuleReportRequest) SetUserType(v string) *GenerateConfigRuleReportRequest {
	s.UserType = &v
	return s
}

type GenerateConfigRuleReportResponseBody struct {
	// The ID of the resource non-compliance report.
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateConfigRuleReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateConfigRuleReportResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateConfigRuleReportResponseBody) SetReportId(v string) *GenerateConfigRuleReportResponseBody {
	s.ReportId = &v
	return s
}

func (s *GenerateConfigRuleReportResponseBody) SetRequestId(v string) *GenerateConfigRuleReportResponseBody {
	s.RequestId = &v
	return s
}

type GenerateConfigRuleReportResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateConfigRuleReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateConfigRuleReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateConfigRuleReportResponse) GoString() string {
	return s.String()
}

func (s *GenerateConfigRuleReportResponse) SetHeaders(v map[string]*string) *GenerateConfigRuleReportResponse {
	s.Headers = v
	return s
}

func (s *GenerateConfigRuleReportResponse) SetStatusCode(v int32) *GenerateConfigRuleReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateConfigRuleReportResponse) SetBody(v *GenerateConfigRuleReportResponseBody) *GenerateConfigRuleReportResponse {
	s.Body = v
	return s
}

type GetConfigRuleReportRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the object.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the object. Valid values:
	//
	// *   USER: the current logon account. This value is available if you use the Tag Policy feature in single-account mode.
	// *   ROOT: the Root folder in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	// *   FOLDER: a folder other than the Root folder in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	// *   ACCOUNT: a member in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// >  The value of this parameter is not case-sensitive.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The mode of the Tag Policy feature. Valid values:
	//
	// *   USER: single-account mode
	// *   RD: multi-account mode
	//
	// For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
	//
	// >  The value of this parameter is not case-sensitive.
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s GetConfigRuleReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleReportRequest) GoString() string {
	return s.String()
}

func (s *GetConfigRuleReportRequest) SetOwnerAccount(v string) *GetConfigRuleReportRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetConfigRuleReportRequest) SetOwnerId(v int64) *GetConfigRuleReportRequest {
	s.OwnerId = &v
	return s
}

func (s *GetConfigRuleReportRequest) SetRegionId(v string) *GetConfigRuleReportRequest {
	s.RegionId = &v
	return s
}

func (s *GetConfigRuleReportRequest) SetResourceOwnerAccount(v string) *GetConfigRuleReportRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetConfigRuleReportRequest) SetTargetId(v string) *GetConfigRuleReportRequest {
	s.TargetId = &v
	return s
}

func (s *GetConfigRuleReportRequest) SetTargetType(v string) *GetConfigRuleReportRequest {
	s.TargetType = &v
	return s
}

func (s *GetConfigRuleReportRequest) SetUserType(v string) *GetConfigRuleReportRequest {
	s.UserType = &v
	return s
}

type GetConfigRuleReportResponseBody struct {
	// The basic information of the resource non-compliance report that is last generated.
	Data *GetConfigRuleReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   true: The request is successful.
	// *   false: The request fails.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetConfigRuleReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfigRuleReportResponseBody) SetData(v *GetConfigRuleReportResponseBodyData) *GetConfigRuleReportResponseBody {
	s.Data = v
	return s
}

func (s *GetConfigRuleReportResponseBody) SetHttpStatusCode(v int32) *GetConfigRuleReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetConfigRuleReportResponseBody) SetRequestId(v string) *GetConfigRuleReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConfigRuleReportResponseBody) SetSuccess(v bool) *GetConfigRuleReportResponseBody {
	s.Success = &v
	return s
}

type GetConfigRuleReportResponseBodyData struct {
	// The time when the report was generated. This value is a UNIX timestamp.
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The ID of the report.
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The ID of the object.
	//
	// >  This parameter is returned if you set the `TargetType` and `TargetId` parameters in the current request to the same values as the parameters that are configured when you call the [GenerateConfigRuleReport](~~433313~~) operation to generate the report.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the object. Valid values:
	//
	// *   USER: the current logon account. This value is available if you use the Tag Policy feature in single-account mode.
	// *   ROOT: the Root folder in the resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	// *   FOLDER: a folder other than the Root folder in the resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	// *   ACCOUNT: a member in the resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// >  This parameter is returned if you set the `TargetType` and `TargetId` parameters in the current request to the same values as the parameters that are configured when you call the [GenerateConfigRuleReport](~~433313~~) operation to generate the report.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GetConfigRuleReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetConfigRuleReportResponseBodyData) SetCreatedTime(v int64) *GetConfigRuleReportResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetConfigRuleReportResponseBodyData) SetReportId(v string) *GetConfigRuleReportResponseBodyData {
	s.ReportId = &v
	return s
}

func (s *GetConfigRuleReportResponseBodyData) SetTargetId(v string) *GetConfigRuleReportResponseBodyData {
	s.TargetId = &v
	return s
}

func (s *GetConfigRuleReportResponseBodyData) SetTargetType(v string) *GetConfigRuleReportResponseBodyData {
	s.TargetType = &v
	return s
}

type GetConfigRuleReportResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConfigRuleReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConfigRuleReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleReportResponse) GoString() string {
	return s.String()
}

func (s *GetConfigRuleReportResponse) SetHeaders(v map[string]*string) *GetConfigRuleReportResponse {
	s.Headers = v
	return s
}

func (s *GetConfigRuleReportResponse) SetStatusCode(v int32) *GetConfigRuleReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConfigRuleReportResponse) SetBody(v *GetConfigRuleReportResponseBody) *GetConfigRuleReportResponse {
	s.Body = v
	return s
}

type GetEffectivePolicyRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the object.
	//
	// >  If you use the Tag Policy feature in single-account mode, this parameter is optional. If you use the Tag Policy feature in multi-account mode, this feature is required.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the object. Valid values:
	//
	// *   USER: the current logon account. This value is available if you use the Tag Policy feature in single-account mode.
	// *   ROOT: the Root folder in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	// *   FOLDER: a folder other than the Root folder in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	// *   ACCOUNT: a member in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// >  If you use the Tag Policy feature in single-account mode, this parameter is optional. If you use the Tag Policy feature in multi-account mode, this feature is required. The value of this parameter is not case-sensitive.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GetEffectivePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEffectivePolicyRequest) GoString() string {
	return s.String()
}

func (s *GetEffectivePolicyRequest) SetOwnerAccount(v string) *GetEffectivePolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetEffectivePolicyRequest) SetOwnerId(v int64) *GetEffectivePolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *GetEffectivePolicyRequest) SetRegionId(v string) *GetEffectivePolicyRequest {
	s.RegionId = &v
	return s
}

func (s *GetEffectivePolicyRequest) SetResourceOwnerAccount(v string) *GetEffectivePolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetEffectivePolicyRequest) SetTargetId(v string) *GetEffectivePolicyRequest {
	s.TargetId = &v
	return s
}

func (s *GetEffectivePolicyRequest) SetTargetType(v string) *GetEffectivePolicyRequest {
	s.TargetType = &v
	return s
}

type GetEffectivePolicyResponseBody struct {
	// The effective tag policy.
	EffectivePolicy *string `json:"EffectivePolicy,omitempty" xml:"EffectivePolicy,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEffectivePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEffectivePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetEffectivePolicyResponseBody) SetEffectivePolicy(v string) *GetEffectivePolicyResponseBody {
	s.EffectivePolicy = &v
	return s
}

func (s *GetEffectivePolicyResponseBody) SetRequestId(v string) *GetEffectivePolicyResponseBody {
	s.RequestId = &v
	return s
}

type GetEffectivePolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEffectivePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEffectivePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEffectivePolicyResponse) GoString() string {
	return s.String()
}

func (s *GetEffectivePolicyResponse) SetHeaders(v map[string]*string) *GetEffectivePolicyResponse {
	s.Headers = v
	return s
}

func (s *GetEffectivePolicyResponse) SetStatusCode(v int32) *GetEffectivePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEffectivePolicyResponse) SetBody(v *GetEffectivePolicyResponseBody) *GetEffectivePolicyResponse {
	s.Body = v
	return s
}

type GetPolicyRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the tag policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s GetPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetPolicyRequest) SetOwnerAccount(v string) *GetPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetPolicyRequest) SetOwnerId(v int64) *GetPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *GetPolicyRequest) SetPolicyId(v string) *GetPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *GetPolicyRequest) SetRegionId(v string) *GetPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *GetPolicyRequest) SetResourceOwnerAccount(v string) *GetPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type GetPolicyResponseBody struct {
	// The details of the tag policy.
	Policy *GetPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBody) SetPolicy(v *GetPolicyResponseBodyPolicy) *GetPolicyResponseBody {
	s.Policy = v
	return s
}

func (s *GetPolicyResponseBody) SetRequestId(v string) *GetPolicyResponseBody {
	s.RequestId = &v
	return s
}

type GetPolicyResponseBodyPolicy struct {
	// The document of the tag policy.
	PolicyContent *string `json:"PolicyContent,omitempty" xml:"PolicyContent,omitempty"`
	// The description of the tag policy.
	PolicyDesc *string `json:"PolicyDesc,omitempty" xml:"PolicyDesc,omitempty"`
	// The name of the tag policy.
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The mode of the Tag Policy feature. Valid values:
	//
	// *   USER: single-account mode
	// *   RD: multi-account mode
	//
	// For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s GetPolicyResponseBodyPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *GetPolicyResponseBodyPolicy) SetPolicyContent(v string) *GetPolicyResponseBodyPolicy {
	s.PolicyContent = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetPolicyDesc(v string) *GetPolicyResponseBodyPolicy {
	s.PolicyDesc = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetPolicyName(v string) *GetPolicyResponseBodyPolicy {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyResponseBodyPolicy) SetUserType(v string) *GetPolicyResponseBodyPolicy {
	s.UserType = &v
	return s
}

type GetPolicyResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetPolicyResponse) SetHeaders(v map[string]*string) *GetPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetPolicyResponse) SetStatusCode(v int32) *GetPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPolicyResponse) SetBody(v *GetPolicyResponseBody) *GetPolicyResponse {
	s.Body = v
	return s
}

type GetPolicyEnableStatusRequest struct {
	OpenType     *string `json:"OpenType,omitempty" xml:"OpenType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The mode of the Tag Policy feature. This parameter specifies a filter condition for the query. Valid values:
	//
	// *   USER: single-account mode
	// *   RD: multi-account mode
	//
	// For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
	//
	// >  The value of this parameter is not case-sensitive.
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s GetPolicyEnableStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyEnableStatusRequest) GoString() string {
	return s.String()
}

func (s *GetPolicyEnableStatusRequest) SetOpenType(v string) *GetPolicyEnableStatusRequest {
	s.OpenType = &v
	return s
}

func (s *GetPolicyEnableStatusRequest) SetOwnerAccount(v string) *GetPolicyEnableStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetPolicyEnableStatusRequest) SetOwnerId(v int64) *GetPolicyEnableStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *GetPolicyEnableStatusRequest) SetRegionId(v string) *GetPolicyEnableStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetPolicyEnableStatusRequest) SetResourceOwnerAccount(v string) *GetPolicyEnableStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetPolicyEnableStatusRequest) SetResourceOwnerId(v string) *GetPolicyEnableStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetPolicyEnableStatusRequest) SetUserType(v string) *GetPolicyEnableStatusRequest {
	s.UserType = &v
	return s
}

type GetPolicyEnableStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the Tag Policy feature.
	StatusModels []*GetPolicyEnableStatusResponseBodyStatusModels `json:"StatusModels,omitempty" xml:"StatusModels,omitempty" type:"Repeated"`
}

func (s GetPolicyEnableStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyEnableStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetPolicyEnableStatusResponseBody) SetRequestId(v string) *GetPolicyEnableStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPolicyEnableStatusResponseBody) SetStatusModels(v []*GetPolicyEnableStatusResponseBodyStatusModels) *GetPolicyEnableStatusResponseBody {
	s.StatusModels = v
	return s
}

type GetPolicyEnableStatusResponseBodyStatusModels struct {
	// The status of the Tag Policy feature. Valid values:
	//
	// *   PendingEnable: The feature is being enabled.
	// *   Enabled: The feature is enabled.
	// *   Closing: The feature is being disabled.
	// *   Disabled: The feature is disabled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The mode of the Tag Policy feature. Valid values:
	//
	// *   USER: single-account mode
	// *   RD: multi-account mode
	//
	// For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s GetPolicyEnableStatusResponseBodyStatusModels) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyEnableStatusResponseBodyStatusModels) GoString() string {
	return s.String()
}

func (s *GetPolicyEnableStatusResponseBodyStatusModels) SetStatus(v string) *GetPolicyEnableStatusResponseBodyStatusModels {
	s.Status = &v
	return s
}

func (s *GetPolicyEnableStatusResponseBodyStatusModels) SetUserType(v string) *GetPolicyEnableStatusResponseBodyStatusModels {
	s.UserType = &v
	return s
}

type GetPolicyEnableStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPolicyEnableStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPolicyEnableStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyEnableStatusResponse) GoString() string {
	return s.String()
}

func (s *GetPolicyEnableStatusResponse) SetHeaders(v map[string]*string) *GetPolicyEnableStatusResponse {
	s.Headers = v
	return s
}

func (s *GetPolicyEnableStatusResponse) SetStatusCode(v int32) *GetPolicyEnableStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPolicyEnableStatusResponse) SetBody(v *GetPolicyEnableStatusResponseBody) *GetPolicyEnableStatusResponse {
	s.Body = v
	return s
}

type ListConfigRulesForTargetRequest struct {
	// The number of entries to return on each page.
	//
	// Default value: 50. Maximum value: 1000.
	MaxResult *int32 `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	// The token that is used to start the next query.
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The use scenario of the tag policy. This parameter specifies a filter condition for the query. Valid values:
	//
	// *   tags: enables tags with specified tag values to be added to resources.
	// *   rg_inherit: enables resources in a resource group to automatically inherit tags from the resource group.
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The tag key. This parameter specifies a filter condition for the query.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The ID of the object. This parameter specifies a filter condition for the query.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the object. This parameter specifies a filter condition for the query. Valid values:
	//
	// *   USER: the current logon account. This value is available if you use the Tag Policy feature in single-account mode.
	// *   ROOT: the Root folder in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	// *   FOLDER: a folder other than the Root folder in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	// *   ACCOUNT: a member in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// >  The value of this parameter is not case-sensitive.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The mode of the Tag Policy feature. This parameter specifies a filter condition for the query. Valid values:
	//
	// *   USER: single-account mode
	// *   RD: multi-account mode
	//
	// For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
	//
	// >  The value of this parameter is not case-sensitive.
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s ListConfigRulesForTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRulesForTargetRequest) GoString() string {
	return s.String()
}

func (s *ListConfigRulesForTargetRequest) SetMaxResult(v int32) *ListConfigRulesForTargetRequest {
	s.MaxResult = &v
	return s
}

func (s *ListConfigRulesForTargetRequest) SetNextToken(v string) *ListConfigRulesForTargetRequest {
	s.NextToken = &v
	return s
}

func (s *ListConfigRulesForTargetRequest) SetOwnerAccount(v string) *ListConfigRulesForTargetRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListConfigRulesForTargetRequest) SetOwnerId(v int64) *ListConfigRulesForTargetRequest {
	s.OwnerId = &v
	return s
}

func (s *ListConfigRulesForTargetRequest) SetPolicyType(v string) *ListConfigRulesForTargetRequest {
	s.PolicyType = &v
	return s
}

func (s *ListConfigRulesForTargetRequest) SetRegionId(v string) *ListConfigRulesForTargetRequest {
	s.RegionId = &v
	return s
}

func (s *ListConfigRulesForTargetRequest) SetResourceOwnerAccount(v string) *ListConfigRulesForTargetRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListConfigRulesForTargetRequest) SetTagKey(v string) *ListConfigRulesForTargetRequest {
	s.TagKey = &v
	return s
}

func (s *ListConfigRulesForTargetRequest) SetTargetId(v string) *ListConfigRulesForTargetRequest {
	s.TargetId = &v
	return s
}

func (s *ListConfigRulesForTargetRequest) SetTargetType(v string) *ListConfigRulesForTargetRequest {
	s.TargetType = &v
	return s
}

func (s *ListConfigRulesForTargetRequest) SetUserType(v string) *ListConfigRulesForTargetRequest {
	s.UserType = &v
	return s
}

type ListConfigRulesForTargetResponseBody struct {
	// The tag detection tasks.
	Data []*ListConfigRulesForTargetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Indicates whether the next query is required.
	//
	// *   If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the next query is not required.
	// *   If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConfigRulesForTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRulesForTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigRulesForTargetResponseBody) SetData(v []*ListConfigRulesForTargetResponseBodyData) *ListConfigRulesForTargetResponseBody {
	s.Data = v
	return s
}

func (s *ListConfigRulesForTargetResponseBody) SetNextToken(v string) *ListConfigRulesForTargetResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListConfigRulesForTargetResponseBody) SetRequestId(v string) *ListConfigRulesForTargetResponseBody {
	s.RequestId = &v
	return s
}

type ListConfigRulesForTargetResponseBodyData struct {
	// The ID of the account group.
	//
	// You can use the ID to query the content of the related resource non-compliance report in Cloud Config.
	//
	// >  This parameter is returned only if you use the Tag Policy feature in multi-account mode.
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the rule.
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The use scenario of the tag policy. Valid values:
	//
	// *   tags: enables tags with specified tag values to be added to resources.
	// *   rg_inherit: enables resources in a resource group to automatically inherit tags from the resource group.
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// Indicates whether automatic remediation is enabled. Valid values:
	//
	// *   true
	// *   false
	Remediation *bool `json:"Remediation,omitempty" xml:"Remediation,omitempty"`
	// The tag key.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value for automatic remediation.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	// The ID of the object.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the object. Valid values:
	//
	// *   USER: the current logon account. This value is available if you use the Tag Policy feature in single-account mode.
	// *   ROOT: the Root folder in the resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	// *   FOLDER: a folder other than the Root folder in the resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	// *   ACCOUNT: a member in the resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListConfigRulesForTargetResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRulesForTargetResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConfigRulesForTargetResponseBodyData) SetAggregatorId(v string) *ListConfigRulesForTargetResponseBodyData {
	s.AggregatorId = &v
	return s
}

func (s *ListConfigRulesForTargetResponseBodyData) SetConfigRuleId(v string) *ListConfigRulesForTargetResponseBodyData {
	s.ConfigRuleId = &v
	return s
}

func (s *ListConfigRulesForTargetResponseBodyData) SetPolicyType(v string) *ListConfigRulesForTargetResponseBodyData {
	s.PolicyType = &v
	return s
}

func (s *ListConfigRulesForTargetResponseBodyData) SetRemediation(v bool) *ListConfigRulesForTargetResponseBodyData {
	s.Remediation = &v
	return s
}

func (s *ListConfigRulesForTargetResponseBodyData) SetTagKey(v string) *ListConfigRulesForTargetResponseBodyData {
	s.TagKey = &v
	return s
}

func (s *ListConfigRulesForTargetResponseBodyData) SetTagValue(v string) *ListConfigRulesForTargetResponseBodyData {
	s.TagValue = &v
	return s
}

func (s *ListConfigRulesForTargetResponseBodyData) SetTargetId(v string) *ListConfigRulesForTargetResponseBodyData {
	s.TargetId = &v
	return s
}

func (s *ListConfigRulesForTargetResponseBodyData) SetTargetType(v string) *ListConfigRulesForTargetResponseBodyData {
	s.TargetType = &v
	return s
}

type ListConfigRulesForTargetResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConfigRulesForTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConfigRulesForTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRulesForTargetResponse) GoString() string {
	return s.String()
}

func (s *ListConfigRulesForTargetResponse) SetHeaders(v map[string]*string) *ListConfigRulesForTargetResponse {
	s.Headers = v
	return s
}

func (s *ListConfigRulesForTargetResponse) SetStatusCode(v int32) *ListConfigRulesForTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConfigRulesForTargetResponse) SetBody(v *ListConfigRulesForTargetResponseBody) *ListConfigRulesForTargetResponse {
	s.Body = v
	return s
}

type ListPoliciesRequest struct {
	// The number of entries to return on each page.
	//
	// Default value: 50. Maximum value: 1000.
	MaxResult *int32 `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	// The token that is used to start the next query.
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of a tag policy. This parameter specifies a filter condition for the query.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The name of a tag policy. This parameter specifies a filter condition for the query.
	PolicyNames []*string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty" type:"Repeated"`
	// The region ID. Set the value to cn-shanghai.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The mode of the Tag Policy feature. This parameter specifies a filter condition for the query. Valid values:
	//
	// *   USER: single-account mode
	// *   RD: multi-account mode
	//
	// For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
	//
	// >  The value of this parameter is not case-sensitive.
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s ListPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesRequest) SetMaxResult(v int32) *ListPoliciesRequest {
	s.MaxResult = &v
	return s
}

func (s *ListPoliciesRequest) SetNextToken(v string) *ListPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPoliciesRequest) SetOwnerAccount(v string) *ListPoliciesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListPoliciesRequest) SetOwnerId(v int64) *ListPoliciesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListPoliciesRequest) SetPolicyIds(v []*string) *ListPoliciesRequest {
	s.PolicyIds = v
	return s
}

func (s *ListPoliciesRequest) SetPolicyNames(v []*string) *ListPoliciesRequest {
	s.PolicyNames = v
	return s
}

func (s *ListPoliciesRequest) SetRegionId(v string) *ListPoliciesRequest {
	s.RegionId = &v
	return s
}

func (s *ListPoliciesRequest) SetResourceOwnerAccount(v string) *ListPoliciesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListPoliciesRequest) SetUserType(v string) *ListPoliciesRequest {
	s.UserType = &v
	return s
}

type ListPoliciesResponseBody struct {
	// Indicates whether the next query is required.
	//
	// *   If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the next query is not required.
	// *   If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The tag policies.
	PolicyList []*ListPoliciesResponseBodyPolicyList `json:"PolicyList,omitempty" xml:"PolicyList,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBody) SetNextToken(v string) *ListPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPoliciesResponseBody) SetPolicyList(v []*ListPoliciesResponseBodyPolicyList) *ListPoliciesResponseBody {
	s.PolicyList = v
	return s
}

func (s *ListPoliciesResponseBody) SetRequestId(v string) *ListPoliciesResponseBody {
	s.RequestId = &v
	return s
}

type ListPoliciesResponseBodyPolicyList struct {
	// The document of the tag policy.
	PolicyContent *string `json:"PolicyContent,omitempty" xml:"PolicyContent,omitempty"`
	// The description of the tag policy.
	PolicyDesc *string `json:"PolicyDesc,omitempty" xml:"PolicyDesc,omitempty"`
	// The ID of the tag policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the tag policy.
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The mode of the Tag Policy feature. Valid values:
	//
	// *   USER: single-account mode
	// *   RD: multi-account mode
	//
	// For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s ListPoliciesResponseBodyPolicyList) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponseBodyPolicyList) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBodyPolicyList) SetPolicyContent(v string) *ListPoliciesResponseBodyPolicyList {
	s.PolicyContent = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicyList) SetPolicyDesc(v string) *ListPoliciesResponseBodyPolicyList {
	s.PolicyDesc = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicyList) SetPolicyId(v string) *ListPoliciesResponseBodyPolicyList {
	s.PolicyId = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicyList) SetPolicyName(v string) *ListPoliciesResponseBodyPolicyList {
	s.PolicyName = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicyList) SetUserType(v string) *ListPoliciesResponseBodyPolicyList {
	s.UserType = &v
	return s
}

type ListPoliciesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponse) SetHeaders(v map[string]*string) *ListPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListPoliciesResponse) SetStatusCode(v int32) *ListPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPoliciesResponse) SetBody(v *ListPoliciesResponseBody) *ListPoliciesResponse {
	s.Body = v
	return s
}

type ListPoliciesForTargetRequest struct {
	// The number of entries to return on each page.
	//
	// Default value: 50. Maximum value: 1000.
	MaxResult *int32 `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	// The token that is used to start the next query.
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the object. This parameter specifies a filter condition for the query.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the object. This parameter specifies a filter condition for the query. Valid values:
	//
	// *   USER: the current logon account. This value is available if you use the Tag Policy feature in single-account mode.
	// *   ROOT: the Root folder in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	// *   FOLDER: a folder other than the Root folder in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	// *   ACCOUNT: a member in a resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	//
	// >  The value of this parameter is not case-sensitive.
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListPoliciesForTargetRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesForTargetRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesForTargetRequest) SetMaxResult(v int32) *ListPoliciesForTargetRequest {
	s.MaxResult = &v
	return s
}

func (s *ListPoliciesForTargetRequest) SetNextToken(v string) *ListPoliciesForTargetRequest {
	s.NextToken = &v
	return s
}

func (s *ListPoliciesForTargetRequest) SetOwnerAccount(v string) *ListPoliciesForTargetRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListPoliciesForTargetRequest) SetOwnerId(v int64) *ListPoliciesForTargetRequest {
	s.OwnerId = &v
	return s
}

func (s *ListPoliciesForTargetRequest) SetRegionId(v string) *ListPoliciesForTargetRequest {
	s.RegionId = &v
	return s
}

func (s *ListPoliciesForTargetRequest) SetResourceOwnerAccount(v string) *ListPoliciesForTargetRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListPoliciesForTargetRequest) SetTargetId(v string) *ListPoliciesForTargetRequest {
	s.TargetId = &v
	return s
}

func (s *ListPoliciesForTargetRequest) SetTargetType(v string) *ListPoliciesForTargetRequest {
	s.TargetType = &v
	return s
}

type ListPoliciesForTargetResponseBody struct {
	// The tag policies that are attached to the object.
	Data []*ListPoliciesForTargetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Indicates whether the next query is required.
	//
	// *   If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the next query is not required.
	// *   If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPoliciesForTargetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesForTargetResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoliciesForTargetResponseBody) SetData(v []*ListPoliciesForTargetResponseBodyData) *ListPoliciesForTargetResponseBody {
	s.Data = v
	return s
}

func (s *ListPoliciesForTargetResponseBody) SetNextToken(v string) *ListPoliciesForTargetResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPoliciesForTargetResponseBody) SetRequestId(v string) *ListPoliciesForTargetResponseBody {
	s.RequestId = &v
	return s
}

type ListPoliciesForTargetResponseBodyData struct {
	// The document of the tag policy.
	PolicyContent *string `json:"PolicyContent,omitempty" xml:"PolicyContent,omitempty"`
	// The description of the tag policy.
	PolicyDesc *string `json:"PolicyDesc,omitempty" xml:"PolicyDesc,omitempty"`
	// The ID of the tag policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the tag policy.
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The mode of the Tag Policy feature. Valid values:
	//
	// *   USER: single-account mode
	// *   RD: multi-account mode
	//
	// For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s ListPoliciesForTargetResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesForTargetResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPoliciesForTargetResponseBodyData) SetPolicyContent(v string) *ListPoliciesForTargetResponseBodyData {
	s.PolicyContent = &v
	return s
}

func (s *ListPoliciesForTargetResponseBodyData) SetPolicyDesc(v string) *ListPoliciesForTargetResponseBodyData {
	s.PolicyDesc = &v
	return s
}

func (s *ListPoliciesForTargetResponseBodyData) SetPolicyId(v string) *ListPoliciesForTargetResponseBodyData {
	s.PolicyId = &v
	return s
}

func (s *ListPoliciesForTargetResponseBodyData) SetPolicyName(v string) *ListPoliciesForTargetResponseBodyData {
	s.PolicyName = &v
	return s
}

func (s *ListPoliciesForTargetResponseBodyData) SetUserType(v string) *ListPoliciesForTargetResponseBodyData {
	s.UserType = &v
	return s
}

type ListPoliciesForTargetResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPoliciesForTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPoliciesForTargetResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesForTargetResponse) GoString() string {
	return s.String()
}

func (s *ListPoliciesForTargetResponse) SetHeaders(v map[string]*string) *ListPoliciesForTargetResponse {
	s.Headers = v
	return s
}

func (s *ListPoliciesForTargetResponse) SetStatusCode(v int32) *ListPoliciesForTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPoliciesForTargetResponse) SetBody(v *ListPoliciesForTargetResponseBody) *ListPoliciesForTargetResponse {
	s.Body = v
	return s
}

type ListResourcesByTagRequest struct {
	TagFilter *ListResourcesByTagRequestTagFilter `json:"TagFilter,omitempty" xml:"TagFilter,omitempty" type:"Struct"`
	// The type of the query. Valid values:
	//
	// *   EQUAL: exact match for resources to which the specified tag is added. This is the default value.
	// *   NOT: exact match for resources to which the specified tag is not added.
	FuzzyType *string `json:"FuzzyType,omitempty" xml:"FuzzyType,omitempty"`
	// Specifies whether to return the information of tags added to the resources. Valid values:
	//
	// *   False: does not return the information of tags added to the resources. This is the default value.
	// *   True: returns the information of all tags added to the resources.
	IncludeAllTags *bool `json:"IncludeAllTags,omitempty" xml:"IncludeAllTags,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: 50. Maximum value: 1000.
	MaxResult *int32 `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	// The token that is used to start the next query.
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// For more information about region IDs, see [Endpoints](~~2330902~~).
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The resource type. This parameter specifies a filter condition for the query.
	//
	// *   If you set the FuzzyType parameter to EQUAL, you can set this parameter to a value obtained from the response of the [ListSupportResourceTypes](~~2330915~~) operation.
	// *   If you set the FuzzyType parameter to NOT, you can set this parameter to a resource type provided in **Types of resources that support queries based on the NOT operator**.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourcesByTagRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesByTagRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesByTagRequest) SetTagFilter(v *ListResourcesByTagRequestTagFilter) *ListResourcesByTagRequest {
	s.TagFilter = v
	return s
}

func (s *ListResourcesByTagRequest) SetFuzzyType(v string) *ListResourcesByTagRequest {
	s.FuzzyType = &v
	return s
}

func (s *ListResourcesByTagRequest) SetIncludeAllTags(v bool) *ListResourcesByTagRequest {
	s.IncludeAllTags = &v
	return s
}

func (s *ListResourcesByTagRequest) SetMaxResult(v int32) *ListResourcesByTagRequest {
	s.MaxResult = &v
	return s
}

func (s *ListResourcesByTagRequest) SetNextToken(v string) *ListResourcesByTagRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourcesByTagRequest) SetOwnerAccount(v string) *ListResourcesByTagRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListResourcesByTagRequest) SetOwnerId(v int64) *ListResourcesByTagRequest {
	s.OwnerId = &v
	return s
}

func (s *ListResourcesByTagRequest) SetRegionId(v string) *ListResourcesByTagRequest {
	s.RegionId = &v
	return s
}

func (s *ListResourcesByTagRequest) SetResourceOwnerAccount(v string) *ListResourcesByTagRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListResourcesByTagRequest) SetResourceType(v string) *ListResourcesByTagRequest {
	s.ResourceType = &v
	return s
}

type ListResourcesByTagRequestTagFilter struct {
	// The tag key. This parameter specifies a filter condition for the query.
	//
	// The tag key can be a maximum of 128 characters in length. It cannot contain `http://` or `https://` and cannot start with `acs:` or `aliyun`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. This parameter specifies a filter condition for the query.
	//
	// The tag value can be a maximum of 128 characters in length. It cannot contain `http://` or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListResourcesByTagRequestTagFilter) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesByTagRequestTagFilter) GoString() string {
	return s.String()
}

func (s *ListResourcesByTagRequestTagFilter) SetKey(v string) *ListResourcesByTagRequestTagFilter {
	s.Key = &v
	return s
}

func (s *ListResourcesByTagRequestTagFilter) SetValue(v string) *ListResourcesByTagRequestTagFilter {
	s.Value = &v
	return s
}

type ListResourcesByTagResponseBody struct {
	// Indicates whether the `next query` is required.
	//
	// *   If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the `next query` is not required.
	// *   If the value of this parameter is not empty, the next query is required, and the value is the `token` used to start the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resources.
	Resources []*ListResourcesByTagResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s ListResourcesByTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesByTagResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesByTagResponseBody) SetNextToken(v string) *ListResourcesByTagResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourcesByTagResponseBody) SetRequestId(v string) *ListResourcesByTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourcesByTagResponseBody) SetResources(v []*ListResourcesByTagResponseBodyResources) *ListResourcesByTagResponseBody {
	s.Resources = v
	return s
}

type ListResourcesByTagResponseBodyResources struct {
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The information of the tags.
	//
	// This parameter is returned only if the `IncludeAllTags` parameter is set to `True`.
	Tags []*ListResourcesByTagResponseBodyResourcesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListResourcesByTagResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesByTagResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListResourcesByTagResponseBodyResources) SetResourceId(v string) *ListResourcesByTagResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *ListResourcesByTagResponseBodyResources) SetTags(v []*ListResourcesByTagResponseBodyResourcesTags) *ListResourcesByTagResponseBodyResources {
	s.Tags = v
	return s
}

type ListResourcesByTagResponseBodyResourcesTags struct {
	// The type of the tag. Valid values:
	//
	// *   custom
	// *   system
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListResourcesByTagResponseBodyResourcesTags) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesByTagResponseBodyResourcesTags) GoString() string {
	return s.String()
}

func (s *ListResourcesByTagResponseBodyResourcesTags) SetCategory(v string) *ListResourcesByTagResponseBodyResourcesTags {
	s.Category = &v
	return s
}

func (s *ListResourcesByTagResponseBodyResourcesTags) SetKey(v string) *ListResourcesByTagResponseBodyResourcesTags {
	s.Key = &v
	return s
}

func (s *ListResourcesByTagResponseBodyResourcesTags) SetValue(v string) *ListResourcesByTagResponseBodyResourcesTags {
	s.Value = &v
	return s
}

type ListResourcesByTagResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcesByTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcesByTagResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesByTagResponse) GoString() string {
	return s.String()
}

func (s *ListResourcesByTagResponse) SetHeaders(v map[string]*string) *ListResourcesByTagResponse {
	s.Headers = v
	return s
}

func (s *ListResourcesByTagResponse) SetStatusCode(v int32) *ListResourcesByTagResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcesByTagResponse) SetBody(v *ListResourcesByTagResponseBody) *ListResourcesByTagResponse {
	s.Body = v
	return s
}

type ListSupportResourceTypesRequest struct {
	// The number of entries to return on each page.
	//
	// Maximum value: 1000. Default value: 50.
	MaxResult *int32 `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	// The token that is used to start the next query.
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The service code. This parameter specifies a filter condition for the query.
	//
	// This parameter is obtained from the response.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The region ID.
	//
	// For more information about region IDs, see [Endpoints](~~2330902~~).
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The resource type. This parameter specifies a filter condition for the query.
	//
	// This parameter is obtained from the response.
	ResourceTye *string `json:"ResourceTye,omitempty" xml:"ResourceTye,omitempty"`
	// Specifies whether to return tag-related capability items. Valid values:
	//
	// *   true: The system returns tag-related capability items.
	// *   false (default value): The system does not return tag-related capability items.
	ShowItems *bool `json:"ShowItems,omitempty" xml:"ShowItems,omitempty"`
	// The code of the tag-related capability item. This parameter specifies a filter condition for the query.
	//
	// For more information, see **Tag-related capability items**.
	SupportCode *string `json:"SupportCode,omitempty" xml:"SupportCode,omitempty"`
}

func (s ListSupportResourceTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSupportResourceTypesRequest) GoString() string {
	return s.String()
}

func (s *ListSupportResourceTypesRequest) SetMaxResult(v int32) *ListSupportResourceTypesRequest {
	s.MaxResult = &v
	return s
}

func (s *ListSupportResourceTypesRequest) SetNextToken(v string) *ListSupportResourceTypesRequest {
	s.NextToken = &v
	return s
}

func (s *ListSupportResourceTypesRequest) SetOwnerAccount(v string) *ListSupportResourceTypesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListSupportResourceTypesRequest) SetOwnerId(v int64) *ListSupportResourceTypesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListSupportResourceTypesRequest) SetProductCode(v string) *ListSupportResourceTypesRequest {
	s.ProductCode = &v
	return s
}

func (s *ListSupportResourceTypesRequest) SetRegionId(v string) *ListSupportResourceTypesRequest {
	s.RegionId = &v
	return s
}

func (s *ListSupportResourceTypesRequest) SetResourceOwnerAccount(v string) *ListSupportResourceTypesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListSupportResourceTypesRequest) SetResourceTye(v string) *ListSupportResourceTypesRequest {
	s.ResourceTye = &v
	return s
}

func (s *ListSupportResourceTypesRequest) SetShowItems(v bool) *ListSupportResourceTypesRequest {
	s.ShowItems = &v
	return s
}

func (s *ListSupportResourceTypesRequest) SetSupportCode(v string) *ListSupportResourceTypesRequest {
	s.SupportCode = &v
	return s
}

type ListSupportResourceTypesResponseBody struct {
	// Indicates whether the next query is required.
	//
	// *   If the value of this parameter is empty, all results are returned, and the next query is not required.
	// *   If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The supported resource types.
	SupportResourceTypes []*ListSupportResourceTypesResponseBodySupportResourceTypes `json:"SupportResourceTypes,omitempty" xml:"SupportResourceTypes,omitempty" type:"Repeated"`
}

func (s ListSupportResourceTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSupportResourceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSupportResourceTypesResponseBody) SetNextToken(v string) *ListSupportResourceTypesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSupportResourceTypesResponseBody) SetRequestId(v string) *ListSupportResourceTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSupportResourceTypesResponseBody) SetSupportResourceTypes(v []*ListSupportResourceTypesResponseBodySupportResourceTypes) *ListSupportResourceTypesResponseBody {
	s.SupportResourceTypes = v
	return s
}

type ListSupportResourceTypesResponseBodySupportResourceTypes struct {
	// The service code.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The supported tag-related capability items.
	//
	// >  This parameter is returned only if the `ShowItems` parameter is set to `true`.
	SupportItems []*ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems `json:"SupportItems,omitempty" xml:"SupportItems,omitempty" type:"Repeated"`
}

func (s ListSupportResourceTypesResponseBodySupportResourceTypes) String() string {
	return tea.Prettify(s)
}

func (s ListSupportResourceTypesResponseBodySupportResourceTypes) GoString() string {
	return s.String()
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypes) SetProductCode(v string) *ListSupportResourceTypesResponseBodySupportResourceTypes {
	s.ProductCode = &v
	return s
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypes) SetResourceType(v string) *ListSupportResourceTypesResponseBodySupportResourceTypes {
	s.ResourceType = &v
	return s
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypes) SetSupportItems(v []*ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems) *ListSupportResourceTypesResponseBodySupportResourceTypes {
	s.SupportItems = v
	return s
}

type ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems struct {
	// Indicates whether the tag-related capability item is supported. Valid values:
	//
	// *   true
	// *   false
	Support *bool `json:"Support,omitempty" xml:"Support,omitempty"`
	// The code of the tag-related capability item.
	SupportCode *string `json:"SupportCode,omitempty" xml:"SupportCode,omitempty"`
	// The details of the support for the tag-related capability item.
	SupportDetails []map[string]*string `json:"SupportDetails,omitempty" xml:"SupportDetails,omitempty" type:"Repeated"`
}

func (s ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems) String() string {
	return tea.Prettify(s)
}

func (s ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems) GoString() string {
	return s.String()
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems) SetSupport(v bool) *ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems {
	s.Support = &v
	return s
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems) SetSupportCode(v string) *ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems {
	s.SupportCode = &v
	return s
}

func (s *ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems) SetSupportDetails(v []map[string]*string) *ListSupportResourceTypesResponseBodySupportResourceTypesSupportItems {
	s.SupportDetails = v
	return s
}

type ListSupportResourceTypesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSupportResourceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSupportResourceTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSupportResourceTypesResponse) GoString() string {
	return s.String()
}

func (s *ListSupportResourceTypesResponse) SetHeaders(v map[string]*string) *ListSupportResourceTypesResponse {
	s.Headers = v
	return s
}

func (s *ListSupportResourceTypesResponse) SetStatusCode(v int32) *ListSupportResourceTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSupportResourceTypesResponse) SetBody(v *ListSupportResourceTypesResponseBody) *ListSupportResourceTypesResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	TagFilter *ListTagKeysRequestTagFilter `json:"TagFilter,omitempty" xml:"TagFilter,omitempty" type:"Struct"`
	// The type of the resource tags. This parameter specifies a filter condition for the query. Valid values:
	//
	// *   all (default value)
	// *   custom
	// *   system
	//
	// >  The value of this parameter is not case-sensitive.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The type of the query. Valid values:
	//
	// *   EQUAL: exact match. This is the default value.
	// *   PREFIX: prefix-based fuzzy match.
	//
	// >  This parameter is available only in the China (Shenzhen) and China (Hong Kong) regions.
	FuzzyType *string `json:"FuzzyType,omitempty" xml:"FuzzyType,omitempty"`
	// The token that is used to start the next query.
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of tag keys to return on each page.
	//
	// Maximum value: 1000. Default value: 50.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The category of the tags. This parameter specifies a filter condition for the query. Valid values:
	//
	// *   ResourceTag: resource tags, including custom and system tags. This is the default value.
	// *   MetaTag: preset tags.
	//
	// >  The value of this parameter is not case-sensitive.
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The region ID.
	//
	// For more information about region IDs, see [Endpoints](~~2330902~~).
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The resource type. This parameter specifies a filter condition for the query.
	//
	// Format: `ALIYUN::${ProductCode}::${ResourceType}`. All letters in the value of this parameter must be in uppercase.
	//
	// *   `ProductCode`: the service code. You can set this field to a value obtained from the response of the [ListSupportResourceTypes](~~2330915~~) operation.
	// *   `ResourceType`: the resource type. You can set this field to a value obtained from the response of the [ListSupportResourceTypes](~~2330915~~) operation.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetTagFilter(v *ListTagKeysRequestTagFilter) *ListTagKeysRequest {
	s.TagFilter = v
	return s
}

func (s *ListTagKeysRequest) SetCategory(v string) *ListTagKeysRequest {
	s.Category = &v
	return s
}

func (s *ListTagKeysRequest) SetFuzzyType(v string) *ListTagKeysRequest {
	s.FuzzyType = &v
	return s
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetOwnerAccount(v string) *ListTagKeysRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagKeysRequest) SetOwnerId(v int64) *ListTagKeysRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagKeysRequest) SetPageSize(v int32) *ListTagKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysRequest) SetQueryType(v string) *ListTagKeysRequest {
	s.QueryType = &v
	return s
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceOwnerAccount(v string) *ListTagKeysRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

type ListTagKeysRequestTagFilter struct {
	// The tag key.
	//
	// This parameter is used together with the `FuzzyType` parameter.
	//
	// >  This parameter is available only in the China (Shenzhen) and China (Hong Kong) regions.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListTagKeysRequestTagFilter) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequestTagFilter) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequestTagFilter) SetKey(v string) *ListTagKeysRequestTagFilter {
	s.Key = &v
	return s
}

type ListTagKeysResponseBody struct {
	// The information of the tag keys.
	Keys *ListTagKeysResponseBodyKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
	// Indicates whether the next query is required. The value of this parameter may be empty.
	//
	// *   If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the next query is not required.
	// *   If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetKeys(v *ListTagKeysResponseBodyKeys) *ListTagKeysResponseBody {
	s.Keys = v
	return s
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

type ListTagKeysResponseBodyKeys struct {
	Key []*ListTagKeysResponseBodyKeysKey `json:"Key,omitempty" xml:"Key,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBodyKeys) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBodyKeys) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyKeys) SetKey(v []*ListTagKeysResponseBodyKeysKey) *ListTagKeysResponseBodyKeys {
	s.Key = v
	return s
}

type ListTagKeysResponseBodyKeysKey struct {
	// The type of the resource tag. Valid values:
	//
	// *   custom
	// *   system
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The description of the tag key.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListTagKeysResponseBodyKeysKey) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBodyKeysKey) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyKeysKey) SetCategory(v string) *ListTagKeysResponseBodyKeysKey {
	s.Category = &v
	return s
}

func (s *ListTagKeysResponseBodyKeysKey) SetDescription(v string) *ListTagKeysResponseBodyKeysKey {
	s.Description = &v
	return s
}

func (s *ListTagKeysResponseBodyKeysKey) SetKey(v string) *ListTagKeysResponseBodyKeysKey {
	s.Key = &v
	return s
}

type ListTagKeysResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The type of the tag. Valid values:
	//
	// *   Custom
	// *   System
	// *   All
	//
	// Default value: All.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The token that is used to start the next query.
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of entries to return on each page.
	//
	// Maximum value: 1000. Default value: 50.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// *   If the resources belong to a service that is centrally deployed, set the value to the region ID of the resources by referring to [Regions supported by tag-related operations on resources of centrally deployed Alibaba Cloud services](~~2579691~~).
	// *   If the resources belong to a service that is not centrally deployed, set the value to the region ID of the resources.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of a resource.
	ResourceARN          []*string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The key-value pairs of tags. You can specify 1 to 10 key-value pairs.
	//
	// If you specify multiple tags, the system queries the resources to which all these tags are added.
	//
	// Limits:
	//
	// *   A tag key must be 1 to 128 characters in length.
	// *   A tag value must be 1 to 128 characters in length.
	// *   Tag keys and tag values are case-sensitive.
	// *   Each tag key on a resource can have only one tag value. If you create a tag that has the same key as an existing tag, the value of the existing tag is overwritten.
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetCategory(v string) *ListTagResourcesRequest {
	s.Category = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerAccount(v string) *ListTagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetPageSize(v int32) *ListTagResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceARN(v []*string) *ListTagResourcesRequest {
	s.ResourceARN = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetTags(v string) *ListTagResourcesRequest {
	s.Tags = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// Indicates whether the `next query` is required.
	//
	// *   If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the `next query` is not required.
	// *   If the value of this parameter is not empty, the next query is required, and the value is the `token` used to start the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the tags that are added to the resources.
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	// The ARN of the resource.
	ResourceARN *string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty"`
	// The information of the tags.
	Tags []*ListTagResourcesResponseBodyTagResourcesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceARN(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceARN = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTags(v []*ListTagResourcesResponseBodyTagResourcesTags) *ListTagResourcesResponseBodyTagResources {
	s.Tags = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTags struct {
	// The type of the tag. Valid values:
	//
	// *   Custom
	// *   System
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTags) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTags) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTags) SetCategory(v string) *ListTagResourcesResponseBodyTagResourcesTags {
	s.Category = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTags) SetKey(v string) *ListTagResourcesResponseBodyTagResourcesTags {
	s.Key = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTags) SetValue(v string) *ListTagResourcesResponseBodyTagResourcesTags {
	s.Value = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListTagValuesRequest struct {
	TagFilter *ListTagValuesRequestTagFilter `json:"TagFilter,omitempty" xml:"TagFilter,omitempty" type:"Struct"`
	// The type of the query. Valid values:
	//
	// *   EQUAL: exact match. This is the default value.
	// *   PREFIX: prefix-based fuzzy match.
	//
	// >  This parameter is available only in the China (Shenzhen) and China (Hong Kong) regions.
	FuzzyType *string `json:"FuzzyType,omitempty" xml:"FuzzyType,omitempty"`
	// The tag key. This parameter specifies a filter condition for the query.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The token that is used to start the next query.
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of tag values to return on each page.
	//
	// Maximum value: 1000. Default value: 50.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The category of the tags. This parameter specifies a filter condition for the query. Valid values:
	//
	// *   ResourceTag: resource tags, including custom and system tags. This is the default value.
	// *   MetaTag: preset tags.
	//
	// >  The value of this parameter is not case-sensitive.
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The region ID.
	//
	// For more information about region IDs, see [Endpoints](~~2330902~~).
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The resource type. This parameter specifies a filter condition for the query.
	//
	// Format: `ALIYUN::${ProductCode}::${ResourceType}`. All letters in the value of this parameter must be in uppercase.
	//
	// *   `ProductCode`: the service code. You can set this field to a value obtained from the response of the [ListSupportResourceTypes](~~2330915~~) operation.
	// *   `ResourceType`: the resource type. You can set this field to a value obtained from the response of the [ListSupportResourceTypes](~~2330915~~) operation.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetTagFilter(v *ListTagValuesRequestTagFilter) *ListTagValuesRequest {
	s.TagFilter = v
	return s
}

func (s *ListTagValuesRequest) SetFuzzyType(v string) *ListTagValuesRequest {
	s.FuzzyType = &v
	return s
}

func (s *ListTagValuesRequest) SetKey(v string) *ListTagValuesRequest {
	s.Key = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetOwnerAccount(v string) *ListTagValuesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagValuesRequest) SetOwnerId(v int64) *ListTagValuesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagValuesRequest) SetPageSize(v int32) *ListTagValuesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagValuesRequest) SetQueryType(v string) *ListTagValuesRequest {
	s.QueryType = &v
	return s
}

func (s *ListTagValuesRequest) SetRegionId(v string) *ListTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceOwnerAccount(v string) *ListTagValuesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

type ListTagValuesRequestTagFilter struct {
	// The tag value.
	//
	// This parameter is used together with the `FuzzyType` parameter.
	//
	// >  This parameter is available only in the China (Shenzhen) and China (Hong Kong) regions.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagValuesRequestTagFilter) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequestTagFilter) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequestTagFilter) SetValue(v string) *ListTagValuesRequestTagFilter {
	s.Value = &v
	return s
}

type ListTagValuesResponseBody struct {
	// Indicates whether the next query is required. The value of this parameter may be empty.
	//
	// *   If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the next query is not required.
	// *   If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the tag values.
	Values *ListTagValuesResponseBodyValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
}

func (s ListTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) SetNextToken(v string) *ListTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetValues(v *ListTagValuesResponseBodyValues) *ListTagValuesResponseBody {
	s.Values = v
	return s
}

type ListTagValuesResponseBodyValues struct {
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBodyValues) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBodyValues) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBodyValues) SetValue(v []*string) *ListTagValuesResponseBodyValues {
	s.Value = v
	return s
}

type ListTagValuesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponse) SetHeaders(v map[string]*string) *ListTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListTagValuesResponse) SetStatusCode(v int32) *ListTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type ListTargetsForPolicyRequest struct {
	// The number of entries to return on each page.
	//
	// Default value: 50. Maximum value: 1000.
	MaxResult *int32 `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	// The token that is used to start the next query.
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the tag policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s ListTargetsForPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTargetsForPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListTargetsForPolicyRequest) SetMaxResult(v int32) *ListTargetsForPolicyRequest {
	s.MaxResult = &v
	return s
}

func (s *ListTargetsForPolicyRequest) SetNextToken(v string) *ListTargetsForPolicyRequest {
	s.NextToken = &v
	return s
}

func (s *ListTargetsForPolicyRequest) SetOwnerAccount(v string) *ListTargetsForPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTargetsForPolicyRequest) SetOwnerId(v int64) *ListTargetsForPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTargetsForPolicyRequest) SetPolicyId(v string) *ListTargetsForPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *ListTargetsForPolicyRequest) SetRegionId(v string) *ListTargetsForPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ListTargetsForPolicyRequest) SetResourceOwnerAccount(v string) *ListTargetsForPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type ListTargetsForPolicyResponseBody struct {
	// Indicates whether the object belongs to the resource directory. Valid values:
	//
	// *   true: The object belongs to the resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	// *   false: The object does not belong to the resource directory. This value is available if you use the Tag Policy feature in single-account mode.
	IsRd *bool `json:"IsRd,omitempty" xml:"IsRd,omitempty"`
	// Indicates whether the next query is required.
	//
	// *   If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the next query is not required.
	// *   If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource directory.
	//
	// >  This parameter is returned only if you use the Tag Policy feature in multi-account mode.
	RdId *string `json:"RdId,omitempty" xml:"RdId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The objects to which the tag policy is attached.
	Targets []*ListTargetsForPolicyResponseBodyTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
}

func (s ListTargetsForPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTargetsForPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListTargetsForPolicyResponseBody) SetIsRd(v bool) *ListTargetsForPolicyResponseBody {
	s.IsRd = &v
	return s
}

func (s *ListTargetsForPolicyResponseBody) SetNextToken(v string) *ListTargetsForPolicyResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTargetsForPolicyResponseBody) SetRdId(v string) *ListTargetsForPolicyResponseBody {
	s.RdId = &v
	return s
}

func (s *ListTargetsForPolicyResponseBody) SetRequestId(v string) *ListTargetsForPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTargetsForPolicyResponseBody) SetTargets(v []*ListTargetsForPolicyResponseBodyTargets) *ListTargetsForPolicyResponseBody {
	s.Targets = v
	return s
}

type ListTargetsForPolicyResponseBodyTargets struct {
	// The ID of the object.
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the object. Valid values:
	//
	// *   USER: the current logon account. This value is available if you use the Tag Policy feature in single-account mode.
	// *   ROOT: the Root folder in the resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	// *   FOLDER: a folder other than the Root folder in the resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	// *   ACCOUNT: a member in the resource directory. This value is available if you use the Tag Policy feature in multi-account mode.
	TargetType *int32 `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListTargetsForPolicyResponseBodyTargets) String() string {
	return tea.Prettify(s)
}

func (s ListTargetsForPolicyResponseBodyTargets) GoString() string {
	return s.String()
}

func (s *ListTargetsForPolicyResponseBodyTargets) SetTargetId(v string) *ListTargetsForPolicyResponseBodyTargets {
	s.TargetId = &v
	return s
}

func (s *ListTargetsForPolicyResponseBodyTargets) SetTargetType(v int32) *ListTargetsForPolicyResponseBodyTargets {
	s.TargetType = &v
	return s
}

type ListTargetsForPolicyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTargetsForPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTargetsForPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTargetsForPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListTargetsForPolicyResponse) SetHeaders(v map[string]*string) *ListTargetsForPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListTargetsForPolicyResponse) SetStatusCode(v int32) *ListTargetsForPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTargetsForPolicyResponse) SetBody(v *ListTargetsForPolicyResponseBody) *ListTargetsForPolicyResponse {
	s.Body = v
	return s
}

type ModifyPolicyRequest struct {
	// Specifies whether to perform a dry run for the request. Valid values:
	//
	// *   false: The system performs the related operation based on the parameter settings in the request. This is the default value.
	// *   true: The system does not perform the related operation based on the parameter settings in the request but only verifies the parameter settings.
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The document of the tag policy.
	//
	// For more information about the syntax of a tag policy, see [Syntax of a tag policy](~~417436~~).
	PolicyContent *string `json:"PolicyContent,omitempty" xml:"PolicyContent,omitempty"`
	// The description of the tag policy.
	//
	// The description must be 0 to 512 characters in length.
	PolicyDesc *string `json:"PolicyDesc,omitempty" xml:"PolicyDesc,omitempty"`
	// The ID of the tag policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the tag policy.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, and underscores (\_).
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s ModifyPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyRequest) SetDryRun(v bool) *ModifyPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyPolicyRequest) SetOwnerAccount(v string) *ModifyPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyPolicyRequest) SetOwnerId(v int64) *ModifyPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyPolicyRequest) SetPolicyContent(v string) *ModifyPolicyRequest {
	s.PolicyContent = &v
	return s
}

func (s *ModifyPolicyRequest) SetPolicyDesc(v string) *ModifyPolicyRequest {
	s.PolicyDesc = &v
	return s
}

func (s *ModifyPolicyRequest) SetPolicyId(v string) *ModifyPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *ModifyPolicyRequest) SetPolicyName(v string) *ModifyPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *ModifyPolicyRequest) SetRegionId(v string) *ModifyPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPolicyRequest) SetResourceOwnerAccount(v string) *ModifyPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type ModifyPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPolicyResponseBody) SetRequestId(v string) *ModifyPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyPolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolicyResponse) SetHeaders(v map[string]*string) *ModifyPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyPolicyResponse) SetStatusCode(v int32) *ModifyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPolicyResponse) SetBody(v *ModifyPolicyResponseBody) *ModifyPolicyResponse {
	s.Body = v
	return s
}

type OpenCreatedByRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. Set the value to cn-shanghai.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s OpenCreatedByRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenCreatedByRequest) GoString() string {
	return s.String()
}

func (s *OpenCreatedByRequest) SetOwnerAccount(v string) *OpenCreatedByRequest {
	s.OwnerAccount = &v
	return s
}

func (s *OpenCreatedByRequest) SetOwnerId(v int64) *OpenCreatedByRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenCreatedByRequest) SetRegionId(v string) *OpenCreatedByRequest {
	s.RegionId = &v
	return s
}

func (s *OpenCreatedByRequest) SetResourceOwnerAccount(v string) *OpenCreatedByRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OpenCreatedByRequest) SetResourceOwnerId(v string) *OpenCreatedByRequest {
	s.ResourceOwnerId = &v
	return s
}

type OpenCreatedByResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenCreatedByResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenCreatedByResponseBody) GoString() string {
	return s.String()
}

func (s *OpenCreatedByResponseBody) SetRequestId(v string) *OpenCreatedByResponseBody {
	s.RequestId = &v
	return s
}

type OpenCreatedByResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenCreatedByResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenCreatedByResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenCreatedByResponse) GoString() string {
	return s.String()
}

func (s *OpenCreatedByResponse) SetHeaders(v map[string]*string) *OpenCreatedByResponse {
	s.Headers = v
	return s
}

func (s *OpenCreatedByResponse) SetStatusCode(v int32) *OpenCreatedByResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenCreatedByResponse) SetBody(v *OpenCreatedByResponseBody) *OpenCreatedByResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// *   If the resources belong to a service that is centrally deployed, set the value to `cn-hangzhou` or to the region ID of the resources by referring to [Regions supported by tag-related operations on resources of centrally deployed Alibaba Cloud services](~~2579691~~).
	// *   If the resources belong to a service that is not centrally deployed, set the value to the region ID of the resources.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of a resource.
	ResourceARN          []*string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The key-value pairs of tags. You can specify 1 to 10 key-value pairs.
	//
	// If you specify multiple tags, the system adds all the tags to the specified resources.
	//
	// Limits:
	//
	// *   A tag key must be 1 to 128 characters in length.
	// *   A tag value must be 1 to 128 characters in length.
	// *   Tag keys and tag values are case-sensitive.
	// *   Each tag key on a resource can have only one tag value. If you create a tag that has the same key as an existing tag, the value of the existing tag is overwritten.
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetOwnerAccount(v string) *TagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceARN(v []*string) *TagResourcesRequest {
	s.ResourceARN = v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetTags(v string) *TagResourcesRequest {
	s.Tags = &v
	return s
}

type TagResourcesResponseBody struct {
	// The information about the resources to which tags fail to be added.
	//
	// >
	//
	// *   If tags are added to all resources, the value of `FailedResources` is empty.
	//
	// *   If tags fail to be added to some or all resources, the value of `FailedResources` contains the detailed information about the resources.
	FailedResources *TagResourcesResponseBodyFailedResources `json:"FailedResources,omitempty" xml:"FailedResources,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetFailedResources(v *TagResourcesResponseBodyFailedResources) *TagResourcesResponseBody {
	s.FailedResources = v
	return s
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponseBodyFailedResources struct {
	FailedResource []*TagResourcesResponseBodyFailedResourcesFailedResource `json:"FailedResource,omitempty" xml:"FailedResource,omitempty" type:"Repeated"`
}

func (s TagResourcesResponseBodyFailedResources) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBodyFailedResources) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBodyFailedResources) SetFailedResource(v []*TagResourcesResponseBodyFailedResourcesFailedResource) *TagResourcesResponseBodyFailedResources {
	s.FailedResource = v
	return s
}

type TagResourcesResponseBodyFailedResourcesFailedResource struct {
	// The ARN of the resource.
	ResourceARN *string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty"`
	// The information about the error.
	Result *TagResourcesResponseBodyFailedResourcesFailedResourceResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s TagResourcesResponseBodyFailedResourcesFailedResource) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBodyFailedResourcesFailedResource) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBodyFailedResourcesFailedResource) SetResourceARN(v string) *TagResourcesResponseBodyFailedResourcesFailedResource {
	s.ResourceARN = &v
	return s
}

func (s *TagResourcesResponseBodyFailedResourcesFailedResource) SetResult(v *TagResourcesResponseBodyFailedResourcesFailedResourceResult) *TagResourcesResponseBodyFailedResourcesFailedResource {
	s.Result = v
	return s
}

type TagResourcesResponseBodyFailedResourcesFailedResourceResult struct {
	// The error code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s TagResourcesResponseBodyFailedResourcesFailedResourceResult) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBodyFailedResourcesFailedResourceResult) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBodyFailedResourcesFailedResourceResult) SetCode(v string) *TagResourcesResponseBodyFailedResourcesFailedResourceResult {
	s.Code = &v
	return s
}

func (s *TagResourcesResponseBodyFailedResourcesFailedResourceResult) SetMessage(v string) *TagResourcesResponseBodyFailedResourcesFailedResourceResult {
	s.Message = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// *   If the resources belong to a service that is centrally deployed, set the value to `cn-hangzhou` or to the region ID of the resources by referring to [Regions supported by tag-related operations on resources of centrally deployed Alibaba Cloud services](~~2579691~~).
	// *   If the resources belong to a service that is not centrally deployed, set the value to the region ID of the resources.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of a resource.
	ResourceARN          []*string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// A tag key.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetOwnerAccount(v string) *UntagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceARN(v []*string) *UntagResourcesRequest {
	s.ResourceARN = v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerAccount(v string) *UntagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// The information about the resources from which tags fail to be removed.
	//
	// >
	//
	// *   If tags are removed from all resources, the value of FailedResources is empty.
	//
	// *   If tags fail to be removed from some or all resources, the value of FailedResources contains the detailed information about the resources.
	FailedResources *UntagResourcesResponseBodyFailedResources `json:"FailedResources,omitempty" xml:"FailedResources,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetFailedResources(v *UntagResourcesResponseBodyFailedResources) *UntagResourcesResponseBody {
	s.FailedResources = v
	return s
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponseBodyFailedResources struct {
	FailedResource []*UntagResourcesResponseBodyFailedResourcesFailedResource `json:"FailedResource,omitempty" xml:"FailedResource,omitempty" type:"Repeated"`
}

func (s UntagResourcesResponseBodyFailedResources) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBodyFailedResources) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBodyFailedResources) SetFailedResource(v []*UntagResourcesResponseBodyFailedResourcesFailedResource) *UntagResourcesResponseBodyFailedResources {
	s.FailedResource = v
	return s
}

type UntagResourcesResponseBodyFailedResourcesFailedResource struct {
	// The ARN of the resource.
	ResourceARN *string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty"`
	// The information about the error.
	Result *UntagResourcesResponseBodyFailedResourcesFailedResourceResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UntagResourcesResponseBodyFailedResourcesFailedResource) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBodyFailedResourcesFailedResource) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBodyFailedResourcesFailedResource) SetResourceARN(v string) *UntagResourcesResponseBodyFailedResourcesFailedResource {
	s.ResourceARN = &v
	return s
}

func (s *UntagResourcesResponseBodyFailedResourcesFailedResource) SetResult(v *UntagResourcesResponseBodyFailedResourcesFailedResourceResult) *UntagResourcesResponseBodyFailedResourcesFailedResource {
	s.Result = v
	return s
}

type UntagResourcesResponseBodyFailedResourcesFailedResourceResult struct {
	// The error code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UntagResourcesResponseBodyFailedResourcesFailedResourceResult) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBodyFailedResourcesFailedResourceResult) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBodyFailedResourcesFailedResourceResult) SetCode(v string) *UntagResourcesResponseBodyFailedResourcesFailedResourceResult {
	s.Code = &v
	return s
}

func (s *UntagResourcesResponseBodyFailedResourcesFailedResourceResult) SetMessage(v string) *UntagResourcesResponseBodyFailedResourcesFailedResourceResult {
	s.Message = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
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
	client.EndpointMap = map[string]*string{
		"cn-qingdao":                  tea.String("tag.aliyuncs.com"),
		"cn-beijing":                  tea.String("tag.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("tag.aliyuncs.com"),
		"cn-shanghai":                 tea.String("tag.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("tag.aliyuncs.com"),
		"cn-hongkong":                 tea.String("tag.aliyuncs.com"),
		"ap-southeast-1":              tea.String("tag.aliyuncs.com"),
		"us-west-1":                   tea.String("tag.aliyuncs.com"),
		"us-east-1":                   tea.String("tag.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("tag.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("tag.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("tag.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("tag.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("tag.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("tag.aliyuncs.com"),
		"cn-edge-1":                   tea.String("tag.aliyuncs.com"),
		"cn-fujian":                   tea.String("tag.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("tag.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("tag.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("tag.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("tag.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("tag.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("tag.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("tag.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("tag.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("tag.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("tag.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("tag.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("tag.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("tag.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("tag.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("tag.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("tag.aliyuncs.com"),
		"cn-wuhan":                    tea.String("tag.aliyuncs.com"),
		"cn-yushanfang":               tea.String("tag.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("tag.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("tag.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("tag.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("tag.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("tag.cn-shenzhen-cloudstone.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("tag.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("tag"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

/**
 * If you use the Tag Policy feature in single-account mode, you can call this API operation to attach a tag policy to the current logon account. If you use the Tag Policy feature in multi-account mode, you can call this API operation to attach a tag policy to the Root folder, a folder other than the Root folder, or a member in a resource directory. For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
 * This topic provides an example on how to call the API operation to attach the tag policy with an ID of `p-de62a0bf400e4b69****` to the current logon account. In this example, the Tag Policy feature in single-account mode is used.
 *
 * @param request AttachPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AttachPolicyResponse
 */
func (client *Client) AttachPolicyWithOptions(request *AttachPolicyRequest, runtime *util.RuntimeOptions) (_result *AttachPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachPolicy"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you use the Tag Policy feature in single-account mode, you can call this API operation to attach a tag policy to the current logon account. If you use the Tag Policy feature in multi-account mode, you can call this API operation to attach a tag policy to the Root folder, a folder other than the Root folder, or a member in a resource directory. For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
 * This topic provides an example on how to call the API operation to attach the tag policy with an ID of `p-de62a0bf400e4b69****` to the current logon account. In this example, the Tag Policy feature in single-account mode is used.
 *
 * @param request AttachPolicyRequest
 * @return AttachPolicyResponse
 */
func (client *Client) AttachPolicy(request *AttachPolicyRequest) (_result *AttachPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachPolicyResponse{}
	_body, _err := client.AttachPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckCreatedByEnabledWithOptions(request *CheckCreatedByEnabledRequest, runtime *util.RuntimeOptions) (_result *CheckCreatedByEnabledResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckCreatedByEnabled"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckCreatedByEnabledResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckCreatedByEnabled(request *CheckCreatedByEnabledRequest) (_result *CheckCreatedByEnabledResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckCreatedByEnabledResponse{}
	_body, _err := client.CheckCreatedByEnabledWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseCreatedByWithOptions(request *CloseCreatedByRequest, runtime *util.RuntimeOptions) (_result *CloseCreatedByResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseCreatedBy"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseCreatedByResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseCreatedBy(request *CloseCreatedByRequest) (_result *CloseCreatedByResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseCreatedByResponse{}
	_body, _err := client.CloseCreatedByWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * This topic provides an example on how to call the API operation to create a tag policy named `test`. In this example, the Tag Policy feature in multi-account mode is used. The tag policy defines that resources to which the `CostCenter:Beijing` or `CostCenter:Shanghai` tag is added are compliant and other resources are not compliant.
 *
 * @param request CreatePolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreatePolicyResponse
 */
func (client *Client) CreatePolicyWithOptions(request *CreatePolicyRequest, runtime *util.RuntimeOptions) (_result *CreatePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyContent)) {
		query["PolicyContent"] = request.PolicyContent
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyDesc)) {
		query["PolicyDesc"] = request.PolicyDesc
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePolicy"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * This topic provides an example on how to call the API operation to create a tag policy named `test`. In this example, the Tag Policy feature in multi-account mode is used. The tag policy defines that resources to which the `CostCenter:Beijing` or `CostCenter:Shanghai` tag is added are compliant and other resources are not compliant.
 *
 * @param request CreatePolicyRequest
 * @return CreatePolicyResponse
 */
func (client *Client) CreatePolicy(request *CreatePolicyRequest) (_result *CreatePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePolicyResponse{}
	_body, _err := client.CreatePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * A preset tag is a tag that you create in advance and is available for the resources in all regions. You can create preset tags in the stage of tag planning and add them to specific resources in the stage of tag implementation. When you create a preset tag, you can specify only the tag key. You can specify a tag value in the future.
 * This topic provides an example on how to call the API operation to create a preset tag whose tag key is `Environment` to indicate the business environment.
 *
 * @param request CreateTagsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateTagsResponse
 */
func (client *Client) CreateTagsWithOptions(request *CreateTagsRequest, runtime *util.RuntimeOptions) (_result *CreateTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeyValueParamList)) {
		query["TagKeyValueParamList"] = request.TagKeyValueParamList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTags"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * A preset tag is a tag that you create in advance and is available for the resources in all regions. You can create preset tags in the stage of tag planning and add them to specific resources in the stage of tag implementation. When you create a preset tag, you can specify only the tag key. You can specify a tag value in the future.
 * This topic provides an example on how to call the API operation to create a preset tag whose tag key is `Environment` to indicate the business environment.
 *
 * @param request CreateTagsRequest
 * @return CreateTagsResponse
 */
func (client *Client) CreateTags(request *CreateTagsRequest) (_result *CreateTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTagsResponse{}
	_body, _err := client.CreateTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Before you delete a tag policy, make sure that the tag policy is detached from all objects to which the tag policy is attached. For more information about how to detach a tag policy, see [DetachPolicy](~~429724~~).
 * This topic provides an example on how to call the API operation to delete the tag policy with an ID of `p-557cb141331f41c7****`.
 *
 * @param request DeletePolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeletePolicyResponse
 */
func (client *Client) DeletePolicyWithOptions(request *DeletePolicyRequest, runtime *util.RuntimeOptions) (_result *DeletePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePolicy"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Before you delete a tag policy, make sure that the tag policy is detached from all objects to which the tag policy is attached. For more information about how to detach a tag policy, see [DetachPolicy](~~429724~~).
 * This topic provides an example on how to call the API operation to delete the tag policy with an ID of `p-557cb141331f41c7****`.
 *
 * @param request DeletePolicyRequest
 * @return DeletePolicyResponse
 */
func (client *Client) DeletePolicy(request *DeletePolicyRequest) (_result *DeletePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePolicyResponse{}
	_body, _err := client.DeletePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to delete the preset tag whose tag key is `Environment` and tag value is `test`.
 *
 * @param request DeleteTagRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteTagResponse
 */
func (client *Client) DeleteTagWithOptions(request *DeleteTagRequest, runtime *util.RuntimeOptions) (_result *DeleteTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTag"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to delete the preset tag whose tag key is `Environment` and tag value is `test`.
 *
 * @param request DeleteTagRequest
 * @return DeleteTagResponse
 */
func (client *Client) DeleteTag(request *DeleteTagRequest) (_result *DeleteTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTagResponse{}
	_body, _err := client.DeleteTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If you use the Tag Policy feature in single-account mode, you can call this API operation to detach a tag policy from the current logon account. If you use the Tag Policy feature in multi-account mode, you can call this API operation to detach a tag policy from the Root folder, a folder other than the Root folder, or a member in a resource directory. For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
 * This topic provides an example on how to call the API operation to detach the tag policy with an ID of `p-a3381efe2fe34a75****` from the current logon account. In this example, the Tag Policy feature in single-account mode is used.
 *
 * @param request DetachPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DetachPolicyResponse
 */
func (client *Client) DetachPolicyWithOptions(request *DetachPolicyRequest, runtime *util.RuntimeOptions) (_result *DetachPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachPolicy"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you use the Tag Policy feature in single-account mode, you can call this API operation to detach a tag policy from the current logon account. If you use the Tag Policy feature in multi-account mode, you can call this API operation to detach a tag policy from the Root folder, a folder other than the Root folder, or a member in a resource directory. For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
 * This topic provides an example on how to call the API operation to detach the tag policy with an ID of `p-a3381efe2fe34a75****` from the current logon account. In this example, the Tag Policy feature in single-account mode is used.
 *
 * @param request DetachPolicyRequest
 * @return DetachPolicyResponse
 */
func (client *Client) DetachPolicy(request *DetachPolicyRequest) (_result *DetachPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachPolicyResponse{}
	_body, _err := client.DetachPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisablePolicyTypeWithOptions(request *DisablePolicyTypeRequest, runtime *util.RuntimeOptions) (_result *DisablePolicyTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenType)) {
		query["OpenType"] = request.OpenType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisablePolicyType"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisablePolicyTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisablePolicyType(request *DisablePolicyTypeRequest) (_result *DisablePolicyTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisablePolicyTypeResponse{}
	_body, _err := client.DisablePolicyTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnablePolicyTypeWithOptions(request *EnablePolicyTypeRequest, runtime *util.RuntimeOptions) (_result *EnablePolicyTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenType)) {
		query["OpenType"] = request.OpenType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnablePolicyType"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnablePolicyTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnablePolicyType(request *EnablePolicyTypeRequest) (_result *EnablePolicyTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnablePolicyTypeResponse{}
	_body, _err := client.EnablePolicyTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If you use the Tag Policy feature in single-account mode, you can call this API operation to generate a resource non-compliance report for the current logon account. If you use the Tag Policy feature in multi-account mode, you can call this API operation to generate a resource non-compliance report for the Root folder, a folder other than the Root folder, or a member in a resource directory. For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
 * This topic provides an example on how to call this API operation to generate a resource non-compliance report for the current logon account. In this example, the Tag Policy feature in single-account mode is used.
 *
 * @param request GenerateConfigRuleReportRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GenerateConfigRuleReportResponse
 */
func (client *Client) GenerateConfigRuleReportWithOptions(request *GenerateConfigRuleReportRequest, runtime *util.RuntimeOptions) (_result *GenerateConfigRuleReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateConfigRuleReport"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateConfigRuleReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you use the Tag Policy feature in single-account mode, you can call this API operation to generate a resource non-compliance report for the current logon account. If you use the Tag Policy feature in multi-account mode, you can call this API operation to generate a resource non-compliance report for the Root folder, a folder other than the Root folder, or a member in a resource directory. For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
 * This topic provides an example on how to call this API operation to generate a resource non-compliance report for the current logon account. In this example, the Tag Policy feature in single-account mode is used.
 *
 * @param request GenerateConfigRuleReportRequest
 * @return GenerateConfigRuleReportResponse
 */
func (client *Client) GenerateConfigRuleReport(request *GenerateConfigRuleReportRequest) (_result *GenerateConfigRuleReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateConfigRuleReportResponse{}
	_body, _err := client.GenerateConfigRuleReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If you use the Tag Policy feature in single-account mode, you can use the current logon account to call this API operation to query the basic information of the resource non-compliance report that is last generated for the account. If you use the Tag Policy feature in multi-account mode, you can use the management account of a resource directory to call this API operation to query the basic information of the resource non-compliance report that is last generated for an object in the resource directory. The object can be the Root folder, a folder other than the Root folder, or a member. For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
 * This topic provides an example on how to call this API operation to query the basic information of the resource non-compliance report that is last generated for the current logon account. In this example, the Tag Policy feature in single-account mode is used. The response shows that the ID of the report is `crp-ao0786618088006c****`.
 *
 * @param request GetConfigRuleReportRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetConfigRuleReportResponse
 */
func (client *Client) GetConfigRuleReportWithOptions(request *GetConfigRuleReportRequest, runtime *util.RuntimeOptions) (_result *GetConfigRuleReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConfigRuleReport"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConfigRuleReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you use the Tag Policy feature in single-account mode, you can use the current logon account to call this API operation to query the basic information of the resource non-compliance report that is last generated for the account. If you use the Tag Policy feature in multi-account mode, you can use the management account of a resource directory to call this API operation to query the basic information of the resource non-compliance report that is last generated for an object in the resource directory. The object can be the Root folder, a folder other than the Root folder, or a member. For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
 * This topic provides an example on how to call this API operation to query the basic information of the resource non-compliance report that is last generated for the current logon account. In this example, the Tag Policy feature in single-account mode is used. The response shows that the ID of the report is `crp-ao0786618088006c****`.
 *
 * @param request GetConfigRuleReportRequest
 * @return GetConfigRuleReportResponse
 */
func (client *Client) GetConfigRuleReport(request *GetConfigRuleReportRequest) (_result *GetConfigRuleReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConfigRuleReportResponse{}
	_body, _err := client.GetConfigRuleReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If you use the Tag Policy feature in single-account mode, you can use the current logon account to call this API operation to query the effective tag policy for the account. If you use the Tag Policy feature in multi-account mode, you can use the management account of a resource directory to call this API operation to query the effective tag policy for the Root folder, a folder other than the Root folder, or a member in the resource directory. You can also use a member of a resource directory to call this API operation to query the effective tag policy for the member. For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
 * An effective tag policy is obtained based on tag policy inheritance. For more information, see [Inheritance of a tag policy and calculation of an effective tag policy](~~417435~~).
 * This topic provides an example on how to call the API operation to query the effective tag policy for the current logon account. In this example, the Tag Policy feature in single-account mode is used.
 *
 * @param request GetEffectivePolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetEffectivePolicyResponse
 */
func (client *Client) GetEffectivePolicyWithOptions(request *GetEffectivePolicyRequest, runtime *util.RuntimeOptions) (_result *GetEffectivePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEffectivePolicy"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEffectivePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you use the Tag Policy feature in single-account mode, you can use the current logon account to call this API operation to query the effective tag policy for the account. If you use the Tag Policy feature in multi-account mode, you can use the management account of a resource directory to call this API operation to query the effective tag policy for the Root folder, a folder other than the Root folder, or a member in the resource directory. You can also use a member of a resource directory to call this API operation to query the effective tag policy for the member. For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
 * An effective tag policy is obtained based on tag policy inheritance. For more information, see [Inheritance of a tag policy and calculation of an effective tag policy](~~417435~~).
 * This topic provides an example on how to call the API operation to query the effective tag policy for the current logon account. In this example, the Tag Policy feature in single-account mode is used.
 *
 * @param request GetEffectivePolicyRequest
 * @return GetEffectivePolicyResponse
 */
func (client *Client) GetEffectivePolicy(request *GetEffectivePolicyRequest) (_result *GetEffectivePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEffectivePolicyResponse{}
	_body, _err := client.GetEffectivePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the details of the tag policy with an ID of `p-557cb141331f41c7****`.
 *
 * @param request GetPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetPolicyResponse
 */
func (client *Client) GetPolicyWithOptions(request *GetPolicyRequest, runtime *util.RuntimeOptions) (_result *GetPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPolicy"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the details of the tag policy with an ID of `p-557cb141331f41c7****`.
 *
 * @param request GetPolicyRequest
 * @return GetPolicyResponse
 */
func (client *Client) GetPolicy(request *GetPolicyRequest) (_result *GetPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPolicyResponse{}
	_body, _err := client.GetPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the status of the Tag Policy feature for the current logon account. The response shows that the Tag Policy feature in multi-account mode is enabled for the current logon account.
 *
 * @param request GetPolicyEnableStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetPolicyEnableStatusResponse
 */
func (client *Client) GetPolicyEnableStatusWithOptions(request *GetPolicyEnableStatusRequest, runtime *util.RuntimeOptions) (_result *GetPolicyEnableStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpenType)) {
		query["OpenType"] = request.OpenType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPolicyEnableStatus"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPolicyEnableStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the status of the Tag Policy feature for the current logon account. The response shows that the Tag Policy feature in multi-account mode is enabled for the current logon account.
 *
 * @param request GetPolicyEnableStatusRequest
 * @return GetPolicyEnableStatusResponse
 */
func (client *Client) GetPolicyEnableStatus(request *GetPolicyEnableStatusRequest) (_result *GetPolicyEnableStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPolicyEnableStatusResponse{}
	_body, _err := client.GetPolicyEnableStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If you use the Tag Policy feature in single-account mode, you can use the current logon account to call this API operation to query the tag detection tasks for the account. If you use the Tag Policy feature in multi-account mode, you can use the management account of a resource directory to call this API operation to query the tag detection tasks for the Root folder, a folder other than the Root folder, or a member in the resource directory. For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
 * This topic provides an example on how to call the API operation to query the tag detection tasks for the current logon account. In this example, the Tag Policy feature in single-account mode is used. The response shows that only one tag detection task exists.
 *
 * @param request ListConfigRulesForTargetRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListConfigRulesForTargetResponse
 */
func (client *Client) ListConfigRulesForTargetWithOptions(request *ListConfigRulesForTargetRequest, runtime *util.RuntimeOptions) (_result *ListConfigRulesForTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResult)) {
		query["MaxResult"] = request.MaxResult
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["PolicyType"] = request.PolicyType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConfigRulesForTarget"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConfigRulesForTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you use the Tag Policy feature in single-account mode, you can use the current logon account to call this API operation to query the tag detection tasks for the account. If you use the Tag Policy feature in multi-account mode, you can use the management account of a resource directory to call this API operation to query the tag detection tasks for the Root folder, a folder other than the Root folder, or a member in the resource directory. For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
 * This topic provides an example on how to call the API operation to query the tag detection tasks for the current logon account. In this example, the Tag Policy feature in single-account mode is used. The response shows that only one tag detection task exists.
 *
 * @param request ListConfigRulesForTargetRequest
 * @return ListConfigRulesForTargetResponse
 */
func (client *Client) ListConfigRulesForTarget(request *ListConfigRulesForTargetRequest) (_result *ListConfigRulesForTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConfigRulesForTargetResponse{}
	_body, _err := client.ListConfigRulesForTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If you use the Tag Policy feature in single-account mode, you can use the current logon account to call this API operation to query all tag policies that are created for the account. If you use the Tag Policy feature in multi-account mode, you can use the management account of a resource directory to call this API operation to query all tag policies that are created for the resource directory. For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
 * This topic provides an example on how to call the API operation to query all tag policies that are created for the current logon account. In this example, the Tag Policy feature in single-account mode is used. The response shows that two tag policies are created.
 *
 * @param request ListPoliciesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListPoliciesResponse
 */
func (client *Client) ListPoliciesWithOptions(request *ListPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResult)) {
		query["MaxResult"] = request.MaxResult
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyIds)) {
		query["PolicyIds"] = request.PolicyIds
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyNames)) {
		query["PolicyNames"] = request.PolicyNames
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPolicies"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you use the Tag Policy feature in single-account mode, you can use the current logon account to call this API operation to query all tag policies that are created for the account. If you use the Tag Policy feature in multi-account mode, you can use the management account of a resource directory to call this API operation to query all tag policies that are created for the resource directory. For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
 * This topic provides an example on how to call the API operation to query all tag policies that are created for the current logon account. In this example, the Tag Policy feature in single-account mode is used. The response shows that two tag policies are created.
 *
 * @param request ListPoliciesRequest
 * @return ListPoliciesResponse
 */
func (client *Client) ListPolicies(request *ListPoliciesRequest) (_result *ListPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPoliciesResponse{}
	_body, _err := client.ListPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If you use the Tag Policy feature in single-account mode, you can use the current logon account to call this API operation to query the tag policies that are attached to the account. If you use the Tag Policy feature in multi-account mode, you can use the management account of a resource directory to call this API operation to query the tag policies that are attached to the Root folder, a folder other than the Root folder, or a member in the resource directory. For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
 * This topic provides an example on how to call the API operation to query the tag policies that are attached to the current logon account. In this example, the Tag Policy feature in single-account mode is used. The response shows that only one tag policy is attached to the current logon account.
 *
 * @param request ListPoliciesForTargetRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListPoliciesForTargetResponse
 */
func (client *Client) ListPoliciesForTargetWithOptions(request *ListPoliciesForTargetRequest, runtime *util.RuntimeOptions) (_result *ListPoliciesForTargetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResult)) {
		query["MaxResult"] = request.MaxResult
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPoliciesForTarget"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPoliciesForTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you use the Tag Policy feature in single-account mode, you can use the current logon account to call this API operation to query the tag policies that are attached to the account. If you use the Tag Policy feature in multi-account mode, you can use the management account of a resource directory to call this API operation to query the tag policies that are attached to the Root folder, a folder other than the Root folder, or a member in the resource directory. For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
 * This topic provides an example on how to call the API operation to query the tag policies that are attached to the current logon account. In this example, the Tag Policy feature in single-account mode is used. The response shows that only one tag policy is attached to the current logon account.
 *
 * @param request ListPoliciesForTargetRequest
 * @return ListPoliciesForTargetResponse
 */
func (client *Client) ListPoliciesForTarget(request *ListPoliciesForTargetRequest) (_result *ListPoliciesForTargetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPoliciesForTargetResponse{}
	_body, _err := client.ListPoliciesForTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation in the China (Shenzhen) region to query virtual private clouds (VPCs) to which the tag key `k1` is added. The response shows that the tag key is added to two VPCs.
 *
 * @param request ListResourcesByTagRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListResourcesByTagResponse
 */
func (client *Client) ListResourcesByTagWithOptions(request *ListResourcesByTagRequest, runtime *util.RuntimeOptions) (_result *ListResourcesByTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FuzzyType)) {
		query["FuzzyType"] = request.FuzzyType
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeAllTags)) {
		query["IncludeAllTags"] = request.IncludeAllTags
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResult)) {
		query["MaxResult"] = request.MaxResult
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagFilter)) {
		query["TagFilter"] = request.TagFilter
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourcesByTag"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourcesByTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation in the China (Shenzhen) region to query virtual private clouds (VPCs) to which the tag key `k1` is added. The response shows that the tag key is added to two VPCs.
 *
 * @param request ListResourcesByTagRequest
 * @return ListResourcesByTagResponse
 */
func (client *Client) ListResourcesByTag(request *ListResourcesByTagRequest) (_result *ListResourcesByTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourcesByTagResponse{}
	_body, _err := client.ListResourcesByTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### [](#)Call examples
 * *   Query a list of resource types supported by TagResources or UntagResources. For more information, see [Example](https://api.alibabacloud.com/api/Tag/2018-08-28/ListSupportResourceTypes?tab=DEBUG\\&params=%7B%22RegionId%22:%22cn-hangzhou%22,%22SupportCode%22:%22TAG_CONSOLE_SUPPORT%22%7D).
 * *   Query a list of resource types supported by ListTagResources or ListResourcesByTag. For more information, see [Example](https://api.alibabacloud.com/api/Tag/2018-08-28/ListSupportResourceTypes?tab=DEBUG\\&params=%7B%22RegionId%22:%22cn-hangzhou%22%7D).
 * *   Query a list of resource types that support createdby tags. For more information, see [Example](https://api.alibabacloud.com/api/Tag/2018-08-28/ListSupportResourceTypes?tab=DEBUG\\&params=%7B%22RegionId%22:%22cn-hangzhou%22,%22SupportCode%22:%22CREATED_BY_TAG_CONSOLE_SUPPORT%22%7D).
 *
 * @param request ListSupportResourceTypesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListSupportResourceTypesResponse
 */
func (client *Client) ListSupportResourceTypesWithOptions(request *ListSupportResourceTypesRequest, runtime *util.RuntimeOptions) (_result *ListSupportResourceTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResult)) {
		query["MaxResult"] = request.MaxResult
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTye)) {
		query["ResourceTye"] = request.ResourceTye
	}

	if !tea.BoolValue(util.IsUnset(request.ShowItems)) {
		query["ShowItems"] = request.ShowItems
	}

	if !tea.BoolValue(util.IsUnset(request.SupportCode)) {
		query["SupportCode"] = request.SupportCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSupportResourceTypes"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSupportResourceTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### [](#)Call examples
 * *   Query a list of resource types supported by TagResources or UntagResources. For more information, see [Example](https://api.alibabacloud.com/api/Tag/2018-08-28/ListSupportResourceTypes?tab=DEBUG\\&params=%7B%22RegionId%22:%22cn-hangzhou%22,%22SupportCode%22:%22TAG_CONSOLE_SUPPORT%22%7D).
 * *   Query a list of resource types supported by ListTagResources or ListResourcesByTag. For more information, see [Example](https://api.alibabacloud.com/api/Tag/2018-08-28/ListSupportResourceTypes?tab=DEBUG\\&params=%7B%22RegionId%22:%22cn-hangzhou%22%7D).
 * *   Query a list of resource types that support createdby tags. For more information, see [Example](https://api.alibabacloud.com/api/Tag/2018-08-28/ListSupportResourceTypes?tab=DEBUG\\&params=%7B%22RegionId%22:%22cn-hangzhou%22,%22SupportCode%22:%22CREATED_BY_TAG_CONSOLE_SUPPORT%22%7D).
 *
 * @param request ListSupportResourceTypesRequest
 * @return ListSupportResourceTypesResponse
 */
func (client *Client) ListSupportResourceTypes(request *ListSupportResourceTypesRequest) (_result *ListSupportResourceTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSupportResourceTypesResponse{}
	_body, _err := client.ListSupportResourceTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the tag keys in the `cn-hangzhou` region. The response shows that the following tag keys exist: `team`, `k1`, and `k2`.
 *
 * @param request ListTagKeysRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListTagKeysResponse
 */
func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyType)) {
		query["FuzzyType"] = request.FuzzyType
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		query["QueryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagFilter)) {
		query["TagFilter"] = request.TagFilter
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagKeys"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the tag keys in the `cn-hangzhou` region. The response shows that the following tag keys exist: `team`, `k1`, and `k2`.
 *
 * @param request ListTagKeysRequest
 * @return ListTagKeysResponse
 */
func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * For information about the Alibaba Cloud services that support tags, see [Services that work with Tag](~~171455~~).
 *
 * @param request ListTagResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListTagResourcesResponse
 */
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceARN)) {
		query["ResourceARN"] = request.ResourceARN
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * For information about the Alibaba Cloud services that support tags, see [Services that work with Tag](~~171455~~).
 *
 * @param request ListTagResourcesRequest
 * @return ListTagResourcesResponse
 */
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the values of the tag key `k1` in the `cn-hangzhou` region. The response shows that the value of the tag key `k1` is `v1`.
 *
 * @param request ListTagValuesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListTagValuesResponse
 */
func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FuzzyType)) {
		query["FuzzyType"] = request.FuzzyType
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		query["QueryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagFilter)) {
		query["TagFilter"] = request.TagFilter
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagValues"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to query the values of the tag key `k1` in the `cn-hangzhou` region. The response shows that the value of the tag key `k1` is `v1`.
 *
 * @param request ListTagValuesRequest
 * @return ListTagValuesResponse
 */
func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If you use the Tag Policy feature in single-account mode, you can use the current logon account to call this API operation to query the object to which a tag policy is attached. The object is the current logon account. If you use the Tag Policy feature in multi-account mode, you can use the management account of a resource directory to call this API operation to query the objects to which a tag policy is attached. The objects include the Root folder, folders other than the Root folder, and members in the resource directory. For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
 * This topic provides an example on how to call the API operation to query the objects to which the tag policy with an ID of `p-de62a0bf400e4b69****` is attached. In this example, the Tag Policy feature in multi-account mode is used. The response shows that the tag policy is attached to two members in the related resource directory.
 *
 * @param request ListTargetsForPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListTargetsForPolicyResponse
 */
func (client *Client) ListTargetsForPolicyWithOptions(request *ListTargetsForPolicyRequest, runtime *util.RuntimeOptions) (_result *ListTargetsForPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResult)) {
		query["MaxResult"] = request.MaxResult
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTargetsForPolicy"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTargetsForPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If you use the Tag Policy feature in single-account mode, you can use the current logon account to call this API operation to query the object to which a tag policy is attached. The object is the current logon account. If you use the Tag Policy feature in multi-account mode, you can use the management account of a resource directory to call this API operation to query the objects to which a tag policy is attached. The objects include the Root folder, folders other than the Root folder, and members in the resource directory. For more information about the modes of the Tag Policy feature, see [Modes of the Tag Policy feature](~~417434~~).
 * This topic provides an example on how to call the API operation to query the objects to which the tag policy with an ID of `p-de62a0bf400e4b69****` is attached. In this example, the Tag Policy feature in multi-account mode is used. The response shows that the tag policy is attached to two members in the related resource directory.
 *
 * @param request ListTargetsForPolicyRequest
 * @return ListTargetsForPolicyResponse
 */
func (client *Client) ListTargetsForPolicy(request *ListTargetsForPolicyRequest) (_result *ListTargetsForPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTargetsForPolicyResponse{}
	_body, _err := client.ListTargetsForPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to change the name of a tag policy to `test`.
 *
 * @param request ModifyPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyPolicyResponse
 */
func (client *Client) ModifyPolicyWithOptions(request *ModifyPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyContent)) {
		query["PolicyContent"] = request.PolicyContent
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyDesc)) {
		query["PolicyDesc"] = request.PolicyDesc
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPolicy"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to call the API operation to change the name of a tag policy to `test`.
 *
 * @param request ModifyPolicyRequest
 * @return ModifyPolicyResponse
 */
func (client *Client) ModifyPolicy(request *ModifyPolicyRequest) (_result *ModifyPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPolicyResponse{}
	_body, _err := client.ModifyPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * createdby tags can help you analyze costs and bills and manage the costs of cloud resources in an efficient manner. You can identify the creators of resources based on the createdby tags added to the resources. createdby tags are system tags that are provided by Alibaba Cloud and automatically added to resources. The key of createdby tags is `acs:tag:createdby`.
 *
 * @param request OpenCreatedByRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return OpenCreatedByResponse
 */
func (client *Client) OpenCreatedByWithOptions(request *OpenCreatedByRequest, runtime *util.RuntimeOptions) (_result *OpenCreatedByResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenCreatedBy"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenCreatedByResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * createdby tags can help you analyze costs and bills and manage the costs of cloud resources in an efficient manner. You can identify the creators of resources based on the createdby tags added to the resources. createdby tags are system tags that are provided by Alibaba Cloud and automatically added to resources. The key of createdby tags is `acs:tag:createdby`.
 *
 * @param request OpenCreatedByRequest
 * @return OpenCreatedByResponse
 */
func (client *Client) OpenCreatedBy(request *OpenCreatedByRequest) (_result *OpenCreatedByResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenCreatedByResponse{}
	_body, _err := client.OpenCreatedByWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Tags are used to identify resources. Tags allow you to categorize, search for, and aggregate resources that have the same characteristics from different dimensions. This facilitates resource management. For more information, see [Tag overview](~~156983~~).
 * For information about the Alibaba Cloud services that support tags, see [Services that work with Tag](~~171455~~).
 *
 * @param request TagResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return TagResourcesResponse
 */
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceARN)) {
		query["ResourceARN"] = request.ResourceARN
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Tags are used to identify resources. Tags allow you to categorize, search for, and aggregate resources that have the same characteristics from different dimensions. This facilitates resource management. For more information, see [Tag overview](~~156983~~).
 * For information about the Alibaba Cloud services that support tags, see [Services that work with Tag](~~171455~~).
 *
 * @param request TagResourcesRequest
 * @return TagResourcesResponse
 */
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * After you remove a tag, the tag is automatically deleted within 24 hours if it is not added to other resources.
 * For information about the Alibaba Cloud services that support tags, see [Services that work with Tag](~~171455~~).
 *
 * @param request UntagResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UntagResourcesResponse
 */
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceARN)) {
		query["ResourceARN"] = request.ResourceARN
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2018-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After you remove a tag, the tag is automatically deleted within 24 hours if it is not added to other resources.
 * For information about the Alibaba Cloud services that support tags, see [Services that work with Tag](~~171455~~).
 *
 * @param request UntagResourcesRequest
 * @return UntagResourcesResponse
 */
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
