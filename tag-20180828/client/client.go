// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AttachPolicyRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PolicyId             *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	TargetId             *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetType           *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	// Id of the request
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreatePolicyRequest struct {
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PolicyContent        *string `json:"PolicyContent,omitempty" xml:"PolicyContent,omitempty"`
	PolicyDesc           *string `json:"PolicyDesc,omitempty" xml:"PolicyDesc,omitempty"`
	PolicyName           *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	UserType             *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
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
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Id of the request
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

func (s *CreatePolicyResponseBody) SetRequestId(v string) *CreatePolicyResponseBody {
	s.RequestId = &v
	return s
}

type CreatePolicyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OwnerAccount         *string                                  `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                                   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string                                  `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
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
	Description       *string                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	Key               *string                                                   `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	// Id of the request
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PolicyId             *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
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
	// Id of the request
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	Value                *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	// Id of the request
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AcceptLanguage       *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PolicyId             *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	TargetId             *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetType           *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	// Id of the request
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GenerateConfigRuleReportRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	TargetId             *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetType           *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	UserType             *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
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
	ReportId  *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateConfigRuleReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	TargetId             *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetType           *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	UserType             *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
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
	Data           *GetConfigRuleReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
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
	CreatedTime *int64  `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	ReportId    *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	TargetId    *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetType  *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetConfigRuleReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	TargetId             *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetType           *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	EffectivePolicy *string `json:"EffectivePolicy,omitempty" xml:"EffectivePolicy,omitempty"`
	// Id of the request
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEffectivePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PolicyId             *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
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
	Policy *GetPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// Id of the request
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
	PolicyContent *string `json:"PolicyContent,omitempty" xml:"PolicyContent,omitempty"`
	PolicyDesc    *string `json:"PolicyDesc,omitempty" xml:"PolicyDesc,omitempty"`
	PolicyName    *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	UserType      *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserType             *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s GetPolicyEnableStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPolicyEnableStatusRequest) GoString() string {
	return s.String()
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
	RequestId    *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPolicyEnableStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MaxResult            *int32  `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PolicyType           *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	TagKey               *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TargetId             *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetType           *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	UserType             *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
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
	Data      []*ListConfigRulesForTargetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	NextToken *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	PolicyType   *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	Remediation  *bool   `json:"Remediation,omitempty" xml:"Remediation,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TargetId     *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetType   *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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

func (s *ListConfigRulesForTargetResponseBodyData) SetTargetId(v string) *ListConfigRulesForTargetResponseBodyData {
	s.TargetId = &v
	return s
}

func (s *ListConfigRulesForTargetResponseBodyData) SetTargetType(v string) *ListConfigRulesForTargetResponseBodyData {
	s.TargetType = &v
	return s
}

type ListConfigRulesForTargetResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListConfigRulesForTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MaxResult            *int32    `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	NextToken            *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PolicyIds            []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	PolicyNames          []*string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty" type:"Repeated"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	UserType             *string   `json:"UserType,omitempty" xml:"UserType,omitempty"`
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
	NextToken  *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PolicyList []*ListPoliciesResponseBodyPolicyList `json:"PolicyList,omitempty" xml:"PolicyList,omitempty" type:"Repeated"`
	// Id of the request
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
	PolicyContent *string `json:"PolicyContent,omitempty" xml:"PolicyContent,omitempty"`
	PolicyDesc    *string `json:"PolicyDesc,omitempty" xml:"PolicyDesc,omitempty"`
	PolicyId      *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PolicyName    *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	UserType      *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MaxResult            *int32  `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	TargetId             *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetType           *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	Data      []*ListPoliciesForTargetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	NextToken *string                                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	PolicyContent *string `json:"PolicyContent,omitempty" xml:"PolicyContent,omitempty"`
	PolicyDesc    *string `json:"PolicyDesc,omitempty" xml:"PolicyDesc,omitempty"`
	PolicyId      *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PolicyName    *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	UserType      *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPoliciesForTargetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	TagFilter            *ListResourcesByTagRequestTagFilter `json:"TagFilter,omitempty" xml:"TagFilter,omitempty" type:"Struct"`
	FuzzyType            *string                             `json:"FuzzyType,omitempty" xml:"FuzzyType,omitempty"`
	IncludeAllTags       *bool                               `json:"IncludeAllTags,omitempty" xml:"IncludeAllTags,omitempty"`
	MaxResult            *int32                              `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	NextToken            *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string                             `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                              `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string                             `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceType         *string                             `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	// 表示当前调用返回读取到的位置，空或者空字符串代表数据已经读取完毕
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ResourceId *string                                        `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Tags       []*ListResourcesByTagResponseBodyResourcesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourcesByTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MaxResult            *int32  `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProductCode          *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceTye          *string `json:"ResourceTye,omitempty" xml:"ResourceTye,omitempty"`
	ShowItems            *bool   `json:"ShowItems,omitempty" xml:"ShowItems,omitempty"`
	SupportCode          *string `json:"SupportCode,omitempty" xml:"SupportCode,omitempty"`
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
	NextToken            *string                                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ProductCode  *string                                                                 `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ResourceType *string                                                                 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
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
	Support     *bool   `json:"Support,omitempty" xml:"Support,omitempty"`
	SupportCode *string `json:"SupportCode,omitempty" xml:"SupportCode,omitempty"`
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

type ListSupportResourceTypesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSupportResourceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	TagFilter            *ListTagKeysRequestTagFilter `json:"TagFilter,omitempty" xml:"TagFilter,omitempty" type:"Struct"`
	Category             *string                      `json:"Category,omitempty" xml:"Category,omitempty"`
	FuzzyType            *string                      `json:"FuzzyType,omitempty" xml:"FuzzyType,omitempty"`
	NextToken            *string                      `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string                      `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageSize             *int32                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryType            *string                      `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	RegionId             *string                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string                      `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceType         *string                      `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
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
	Keys      *ListTagKeysResponseBodyKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
	NextToken *string                      `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Category    *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Key         *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Category             *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	NextToken            *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageSize             *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceARN          []*string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	Tags                 *string   `json:"Tags,omitempty" xml:"Tags,omitempty"`
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
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ResourceARN *string                                         `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty"`
	Tags        []*ListTagResourcesResponseBodyTagResourcesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	TagFilter            *ListTagValuesRequestTagFilter `json:"TagFilter,omitempty" xml:"TagFilter,omitempty" type:"Struct"`
	FuzzyType            *string                        `json:"FuzzyType,omitempty" xml:"FuzzyType,omitempty"`
	Key                  *string                        `json:"Key,omitempty" xml:"Key,omitempty"`
	NextToken            *string                        `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string                        `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageSize             *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryType            *string                        `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	RegionId             *string                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string                        `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceType         *string                        `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
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
	NextToken *string                          `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Values    *ListTagValuesResponseBodyValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MaxResult            *int32  `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PolicyId             *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
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
	IsRd      *bool                                      `json:"IsRd,omitempty" xml:"IsRd,omitempty"`
	NextToken *string                                    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RdId      *string                                    `json:"RdId,omitempty" xml:"RdId,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Targets   []*ListTargetsForPolicyResponseBodyTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
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
	TargetId   *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	TargetType *int32  `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTargetsForPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PolicyContent        *string `json:"PolicyContent,omitempty" xml:"PolicyContent,omitempty"`
	PolicyDesc           *string `json:"PolicyDesc,omitempty" xml:"PolicyDesc,omitempty"`
	PolicyId             *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PolicyName           *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
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
	// Id of the request
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type TagResourcesRequest struct {
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceARN          []*string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	Tags                 *string   `json:"Tags,omitempty" xml:"Tags,omitempty"`
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
	FailedResources *TagResourcesResponseBodyFailedResources `json:"FailedResources,omitempty" xml:"FailedResources,omitempty" type:"Struct"`
	RequestId       *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ResourceARN *string                                                      `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty"`
	Result      *TagResourcesResponseBodyFailedResourcesFailedResourceResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceARN          []*string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	TagKey               []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
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
	FailedResources *UntagResourcesResponseBodyFailedResources `json:"FailedResources,omitempty" xml:"FailedResources,omitempty" type:"Struct"`
	RequestId       *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ResourceARN *string                                                        `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty"`
	Result      *UntagResourcesResponseBodyFailedResourcesFailedResourceResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (client *Client) GetPolicyEnableStatusWithOptions(request *GetPolicyEnableStatusRequest, runtime *util.RuntimeOptions) (_result *GetPolicyEnableStatusResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TagFilter))) {
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TagFilter))) {
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.TagFilter))) {
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
