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

type AttachExpressConnectRouterChildInstanceRequest struct {
	ChildInstanceId       *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceOwnerId  *int64  `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceType     *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId                 *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
}

func (s AttachExpressConnectRouterChildInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachExpressConnectRouterChildInstanceRequest) GoString() string {
	return s.String()
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetChildInstanceId(v string) *AttachExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetChildInstanceOwnerId(v int64) *AttachExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetChildInstanceRegionId(v string) *AttachExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetChildInstanceType(v string) *AttachExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetClientToken(v string) *AttachExpressConnectRouterChildInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetDryRun(v bool) *AttachExpressConnectRouterChildInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceRequest) SetEcrId(v string) *AttachExpressConnectRouterChildInstanceRequest {
	s.EcrId = &v
	return s
}

type AttachExpressConnectRouterChildInstanceResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode        *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachExpressConnectRouterChildInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachExpressConnectRouterChildInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetAccessDeniedDetail(v string) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetCode(v string) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetDynamicCode(v string) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetDynamicMessage(v string) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetHttpStatusCode(v int32) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetMessage(v string) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetRequestId(v string) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponseBody) SetSuccess(v bool) *AttachExpressConnectRouterChildInstanceResponseBody {
	s.Success = &v
	return s
}

type AttachExpressConnectRouterChildInstanceResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachExpressConnectRouterChildInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachExpressConnectRouterChildInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachExpressConnectRouterChildInstanceResponse) GoString() string {
	return s.String()
}

func (s *AttachExpressConnectRouterChildInstanceResponse) SetHeaders(v map[string]*string) *AttachExpressConnectRouterChildInstanceResponse {
	s.Headers = v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponse) SetStatusCode(v int32) *AttachExpressConnectRouterChildInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachExpressConnectRouterChildInstanceResponse) SetBody(v *AttachExpressConnectRouterChildInstanceResponseBody) *AttachExpressConnectRouterChildInstanceResponse {
	s.Body = v
	return s
}

type CheckAddRegionToExpressConnectRouterRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun        *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId         *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	FreshRegionId *string `json:"FreshRegionId,omitempty" xml:"FreshRegionId,omitempty"`
}

func (s CheckAddRegionToExpressConnectRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckAddRegionToExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *CheckAddRegionToExpressConnectRouterRequest) SetClientToken(v string) *CheckAddRegionToExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterRequest) SetDryRun(v bool) *CheckAddRegionToExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterRequest) SetEcrId(v string) *CheckAddRegionToExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterRequest) SetFreshRegionId(v string) *CheckAddRegionToExpressConnectRouterRequest {
	s.FreshRegionId = &v
	return s
}

type CheckAddRegionToExpressConnectRouterResponseBody struct {
	AccessDeniedDetail                   *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	AnyCrossBorderLink                   *bool   `json:"AnyCrossBorderLink,omitempty" xml:"AnyCrossBorderLink,omitempty"`
	AnyInterRegionLink                   *bool   `json:"AnyInterRegionLink,omitempty" xml:"AnyInterRegionLink,omitempty"`
	Code                                 *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode                          *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage                       *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode                       *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	IsCdtCrossBorderEnabled              *bool   `json:"IsCdtCrossBorderEnabled,omitempty" xml:"IsCdtCrossBorderEnabled,omitempty"`
	IsCdtInterRegionEnabled              *bool   `json:"IsCdtInterRegionEnabled,omitempty" xml:"IsCdtInterRegionEnabled,omitempty"`
	IsUserAllowedToCreateCrossBorderLink *bool   `json:"IsUserAllowedToCreateCrossBorderLink,omitempty" xml:"IsUserAllowedToCreateCrossBorderLink,omitempty"`
	Message                              *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId                            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                              *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckAddRegionToExpressConnectRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckAddRegionToExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetAnyCrossBorderLink(v bool) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.AnyCrossBorderLink = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetAnyInterRegionLink(v bool) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.AnyInterRegionLink = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetCode(v string) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetDynamicCode(v string) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetDynamicMessage(v string) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetIsCdtCrossBorderEnabled(v bool) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.IsCdtCrossBorderEnabled = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetIsCdtInterRegionEnabled(v bool) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.IsCdtInterRegionEnabled = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetIsUserAllowedToCreateCrossBorderLink(v bool) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.IsUserAllowedToCreateCrossBorderLink = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetMessage(v string) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetRequestId(v string) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponseBody) SetSuccess(v bool) *CheckAddRegionToExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

type CheckAddRegionToExpressConnectRouterResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAddRegionToExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAddRegionToExpressConnectRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckAddRegionToExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *CheckAddRegionToExpressConnectRouterResponse) SetHeaders(v map[string]*string) *CheckAddRegionToExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponse) SetStatusCode(v int32) *CheckAddRegionToExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAddRegionToExpressConnectRouterResponse) SetBody(v *CheckAddRegionToExpressConnectRouterResponseBody) *CheckAddRegionToExpressConnectRouterResponse {
	s.Body = v
	return s
}

type CreateExpressConnectRouterRequest struct {
	AlibabaSideAsn  *int64  `json:"AlibabaSideAsn,omitempty" xml:"AlibabaSideAsn,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun          *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateExpressConnectRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectRouterRequest) SetAlibabaSideAsn(v int64) *CreateExpressConnectRouterRequest {
	s.AlibabaSideAsn = &v
	return s
}

func (s *CreateExpressConnectRouterRequest) SetClientToken(v string) *CreateExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateExpressConnectRouterRequest) SetDescription(v string) *CreateExpressConnectRouterRequest {
	s.Description = &v
	return s
}

func (s *CreateExpressConnectRouterRequest) SetDryRun(v bool) *CreateExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *CreateExpressConnectRouterRequest) SetName(v string) *CreateExpressConnectRouterRequest {
	s.Name = &v
	return s
}

func (s *CreateExpressConnectRouterRequest) SetResourceGroupId(v string) *CreateExpressConnectRouterRequest {
	s.ResourceGroupId = &v
	return s
}

type CreateExpressConnectRouterResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode        *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	EcrId              *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	HttpStatusCode     *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateExpressConnectRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *CreateExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetCode(v string) *CreateExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetDynamicCode(v string) *CreateExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetDynamicMessage(v string) *CreateExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetEcrId(v string) *CreateExpressConnectRouterResponseBody {
	s.EcrId = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *CreateExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetMessage(v string) *CreateExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetRequestId(v string) *CreateExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExpressConnectRouterResponseBody) SetSuccess(v bool) *CreateExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

type CreateExpressConnectRouterResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExpressConnectRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectRouterResponse) SetHeaders(v map[string]*string) *CreateExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *CreateExpressConnectRouterResponse) SetStatusCode(v int32) *CreateExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExpressConnectRouterResponse) SetBody(v *CreateExpressConnectRouterResponseBody) *CreateExpressConnectRouterResponse {
	s.Body = v
	return s
}

type CreateExpressConnectRouterAssociationRequest struct {
	AllowedPrefixes      []*string `json:"AllowedPrefixes,omitempty" xml:"AllowedPrefixes,omitempty" type:"Repeated"`
	AssociationRegionId  *string   `json:"AssociationRegionId,omitempty" xml:"AssociationRegionId,omitempty"`
	CenId                *string   `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CreateAttachment     *bool     `json:"CreateAttachment,omitempty" xml:"CreateAttachment,omitempty"`
	DryRun               *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId                *string   `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	TransitRouterId      *string   `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	TransitRouterOwnerId *int64    `json:"TransitRouterOwnerId,omitempty" xml:"TransitRouterOwnerId,omitempty"`
	VpcId                *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcOwnerId           *int64    `json:"VpcOwnerId,omitempty" xml:"VpcOwnerId,omitempty"`
}

func (s CreateExpressConnectRouterAssociationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExpressConnectRouterAssociationRequest) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectRouterAssociationRequest) SetAllowedPrefixes(v []*string) *CreateExpressConnectRouterAssociationRequest {
	s.AllowedPrefixes = v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetAssociationRegionId(v string) *CreateExpressConnectRouterAssociationRequest {
	s.AssociationRegionId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetCenId(v string) *CreateExpressConnectRouterAssociationRequest {
	s.CenId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetClientToken(v string) *CreateExpressConnectRouterAssociationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetCreateAttachment(v bool) *CreateExpressConnectRouterAssociationRequest {
	s.CreateAttachment = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetDryRun(v bool) *CreateExpressConnectRouterAssociationRequest {
	s.DryRun = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetEcrId(v string) *CreateExpressConnectRouterAssociationRequest {
	s.EcrId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetTransitRouterId(v string) *CreateExpressConnectRouterAssociationRequest {
	s.TransitRouterId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetTransitRouterOwnerId(v int64) *CreateExpressConnectRouterAssociationRequest {
	s.TransitRouterOwnerId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetVpcId(v string) *CreateExpressConnectRouterAssociationRequest {
	s.VpcId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetVpcOwnerId(v int64) *CreateExpressConnectRouterAssociationRequest {
	s.VpcOwnerId = &v
	return s
}

type CreateExpressConnectRouterAssociationResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	AssociationId      *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode        *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateExpressConnectRouterAssociationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateExpressConnectRouterAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetAccessDeniedDetail(v string) *CreateExpressConnectRouterAssociationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetAssociationId(v string) *CreateExpressConnectRouterAssociationResponseBody {
	s.AssociationId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetCode(v string) *CreateExpressConnectRouterAssociationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetDynamicCode(v string) *CreateExpressConnectRouterAssociationResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetDynamicMessage(v string) *CreateExpressConnectRouterAssociationResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetHttpStatusCode(v int32) *CreateExpressConnectRouterAssociationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetMessage(v string) *CreateExpressConnectRouterAssociationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetRequestId(v string) *CreateExpressConnectRouterAssociationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponseBody) SetSuccess(v bool) *CreateExpressConnectRouterAssociationResponseBody {
	s.Success = &v
	return s
}

type CreateExpressConnectRouterAssociationResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExpressConnectRouterAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExpressConnectRouterAssociationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateExpressConnectRouterAssociationResponse) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectRouterAssociationResponse) SetHeaders(v map[string]*string) *CreateExpressConnectRouterAssociationResponse {
	s.Headers = v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponse) SetStatusCode(v int32) *CreateExpressConnectRouterAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationResponse) SetBody(v *CreateExpressConnectRouterAssociationResponseBody) *CreateExpressConnectRouterAssociationResponse {
	s.Body = v
	return s
}

type DeleteExpressConnectRouterRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId       *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
}

func (s DeleteExpressConnectRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectRouterRequest) SetClientToken(v string) *DeleteExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteExpressConnectRouterRequest) SetDryRun(v bool) *DeleteExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteExpressConnectRouterRequest) SetEcrId(v string) *DeleteExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

type DeleteExpressConnectRouterResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode        *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteExpressConnectRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *DeleteExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) SetCode(v string) *DeleteExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) SetDynamicCode(v string) *DeleteExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) SetDynamicMessage(v string) *DeleteExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *DeleteExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) SetMessage(v string) *DeleteExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) SetRequestId(v string) *DeleteExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExpressConnectRouterResponseBody) SetSuccess(v bool) *DeleteExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

type DeleteExpressConnectRouterResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExpressConnectRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectRouterResponse) SetHeaders(v map[string]*string) *DeleteExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *DeleteExpressConnectRouterResponse) SetStatusCode(v int32) *DeleteExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExpressConnectRouterResponse) SetBody(v *DeleteExpressConnectRouterResponseBody) *DeleteExpressConnectRouterResponse {
	s.Body = v
	return s
}

type DeleteExpressConnectRouterAssociationRequest struct {
	AssociationId    *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DeleteAttachment *bool   `json:"DeleteAttachment,omitempty" xml:"DeleteAttachment,omitempty"`
	DryRun           *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId            *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
}

func (s DeleteExpressConnectRouterAssociationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteExpressConnectRouterAssociationRequest) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectRouterAssociationRequest) SetAssociationId(v string) *DeleteExpressConnectRouterAssociationRequest {
	s.AssociationId = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationRequest) SetClientToken(v string) *DeleteExpressConnectRouterAssociationRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationRequest) SetDeleteAttachment(v bool) *DeleteExpressConnectRouterAssociationRequest {
	s.DeleteAttachment = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationRequest) SetDryRun(v bool) *DeleteExpressConnectRouterAssociationRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationRequest) SetEcrId(v string) *DeleteExpressConnectRouterAssociationRequest {
	s.EcrId = &v
	return s
}

type DeleteExpressConnectRouterAssociationResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode        *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteExpressConnectRouterAssociationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExpressConnectRouterAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetAccessDeniedDetail(v string) *DeleteExpressConnectRouterAssociationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetCode(v string) *DeleteExpressConnectRouterAssociationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetDynamicCode(v string) *DeleteExpressConnectRouterAssociationResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetDynamicMessage(v string) *DeleteExpressConnectRouterAssociationResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetHttpStatusCode(v int32) *DeleteExpressConnectRouterAssociationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetMessage(v string) *DeleteExpressConnectRouterAssociationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetRequestId(v string) *DeleteExpressConnectRouterAssociationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponseBody) SetSuccess(v bool) *DeleteExpressConnectRouterAssociationResponseBody {
	s.Success = &v
	return s
}

type DeleteExpressConnectRouterAssociationResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExpressConnectRouterAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExpressConnectRouterAssociationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExpressConnectRouterAssociationResponse) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectRouterAssociationResponse) SetHeaders(v map[string]*string) *DeleteExpressConnectRouterAssociationResponse {
	s.Headers = v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponse) SetStatusCode(v int32) *DeleteExpressConnectRouterAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationResponse) SetBody(v *DeleteExpressConnectRouterAssociationResponseBody) *DeleteExpressConnectRouterAssociationResponse {
	s.Body = v
	return s
}

type DescribeDisabledExpressConnectRouterRouteEntriesRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId       *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	MaxResults  *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) SetClientToken(v string) *DescribeDisabledExpressConnectRouterRouteEntriesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) SetDryRun(v bool) *DescribeDisabledExpressConnectRouterRouteEntriesRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) SetEcrId(v string) *DescribeDisabledExpressConnectRouterRouteEntriesRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) SetMaxResults(v int32) *DescribeDisabledExpressConnectRouterRouteEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesRequest) SetNextToken(v string) *DescribeDisabledExpressConnectRouterRouteEntriesRequest {
	s.NextToken = &v
	return s
}

type DescribeDisabledExpressConnectRouterRouteEntriesResponseBody struct {
	AccessDeniedDetail     *string                                                                               `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code                   *string                                                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	DisabledRouteEntryList []*DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList `json:"DisabledRouteEntryList,omitempty" xml:"DisabledRouteEntryList,omitempty" type:"Repeated"`
	DynamicCode            *string                                                                               `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage         *string                                                                               `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode         *int32                                                                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	MaxResults             *int32                                                                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message                *string                                                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken              *string                                                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId              *string                                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                *bool                                                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount             *int32                                                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetAccessDeniedDetail(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetCode(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetDisabledRouteEntryList(v []*DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.DisabledRouteEntryList = v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetDynamicCode(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetDynamicMessage(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetHttpStatusCode(v int32) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetMaxResults(v int32) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetMessage(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetNextToken(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetRequestId(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetSuccess(v bool) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) SetTotalCount(v int32) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList struct {
	DestinationCidrBlock    *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	EcrId                   *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	GmtCreate               *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	NexthopInstanceId       *string `json:"NexthopInstanceId,omitempty" xml:"NexthopInstanceId,omitempty"`
	NexthopInstanceRegionId *string `json:"NexthopInstanceRegionId,omitempty" xml:"NexthopInstanceRegionId,omitempty"`
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) SetDestinationCidrBlock(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) SetEcrId(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList {
	s.EcrId = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) SetGmtCreate(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) SetNexthopInstanceId(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList {
	s.NexthopInstanceId = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList) SetNexthopInstanceRegionId(v string) *DescribeDisabledExpressConnectRouterRouteEntriesResponseBodyDisabledRouteEntryList {
	s.NexthopInstanceRegionId = &v
	return s
}

type DescribeDisabledExpressConnectRouterRouteEntriesResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDisabledExpressConnectRouterRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponse) SetHeaders(v map[string]*string) *DescribeDisabledExpressConnectRouterRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponse) SetStatusCode(v int32) *DescribeDisabledExpressConnectRouterRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDisabledExpressConnectRouterRouteEntriesResponse) SetBody(v *DescribeDisabledExpressConnectRouterRouteEntriesResponseBody) *DescribeDisabledExpressConnectRouterRouteEntriesResponse {
	s.Body = v
	return s
}

type DescribeExpressConnectRouterRequest struct {
	ClientToken     *string                                         `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun          *bool                                           `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId           *string                                         `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	MaxResults      *int32                                          `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Name            *string                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	NextToken       *string                                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceGroupId *string                                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TagModels       []*DescribeExpressConnectRouterRequestTagModels `json:"TagModels,omitempty" xml:"TagModels,omitempty" type:"Repeated"`
}

func (s DescribeExpressConnectRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRequest) SetClientToken(v string) *DescribeExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetDryRun(v bool) *DescribeExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetEcrId(v string) *DescribeExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetMaxResults(v int32) *DescribeExpressConnectRouterRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetName(v string) *DescribeExpressConnectRouterRequest {
	s.Name = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetNextToken(v string) *DescribeExpressConnectRouterRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetResourceGroupId(v string) *DescribeExpressConnectRouterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeExpressConnectRouterRequest) SetTagModels(v []*DescribeExpressConnectRouterRequestTagModels) *DescribeExpressConnectRouterRequest {
	s.TagModels = v
	return s
}

type DescribeExpressConnectRouterRequestTagModels struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeExpressConnectRouterRequestTagModels) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterRequestTagModels) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRequestTagModels) SetTagKey(v string) *DescribeExpressConnectRouterRequestTagModels {
	s.TagKey = &v
	return s
}

func (s *DescribeExpressConnectRouterRequestTagModels) SetTagValue(v string) *DescribeExpressConnectRouterRequestTagModels {
	s.TagValue = &v
	return s
}

type DescribeExpressConnectRouterResponseBody struct {
	AccessDeniedDetail *string                                            `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode        *string                                            `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string                                            `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	EcrList            []*DescribeExpressConnectRouterResponseBodyEcrList `json:"EcrList,omitempty" xml:"EcrList,omitempty" type:"Repeated"`
	HttpStatusCode     *int32                                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	MaxResults         *int32                                             `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message            *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken          *string                                            `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId          *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount         *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeExpressConnectRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetCode(v string) *DescribeExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetDynamicCode(v string) *DescribeExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetDynamicMessage(v string) *DescribeExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetEcrList(v []*DescribeExpressConnectRouterResponseBodyEcrList) *DescribeExpressConnectRouterResponseBody {
	s.EcrList = v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *DescribeExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetMaxResults(v int32) *DescribeExpressConnectRouterResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetMessage(v string) *DescribeExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetNextToken(v string) *DescribeExpressConnectRouterResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetRequestId(v string) *DescribeExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetSuccess(v bool) *DescribeExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBody) SetTotalCount(v int32) *DescribeExpressConnectRouterResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeExpressConnectRouterResponseBodyEcrList struct {
	AlibabaSideAsn  *int64                                                 `json:"AlibabaSideAsn,omitempty" xml:"AlibabaSideAsn,omitempty"`
	BizStatus       *string                                                `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	Description     *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	EcrId           *string                                                `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	GmtCreate       *string                                                `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified     *string                                                `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Name            *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId         *int64                                                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceGroupId *string                                                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Status          *string                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags            []*DescribeExpressConnectRouterResponseBodyEcrListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeExpressConnectRouterResponseBodyEcrList) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterResponseBodyEcrList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetAlibabaSideAsn(v int64) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.AlibabaSideAsn = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetBizStatus(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.BizStatus = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetDescription(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.Description = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetEcrId(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetGmtCreate(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetGmtModified(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.GmtModified = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetName(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.Name = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetOwnerId(v int64) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.OwnerId = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetResourceGroupId(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetStatus(v string) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.Status = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrList) SetTags(v []*DescribeExpressConnectRouterResponseBodyEcrListTags) *DescribeExpressConnectRouterResponseBodyEcrList {
	s.Tags = v
	return s
}

type DescribeExpressConnectRouterResponseBodyEcrListTags struct {
	AliUid       *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Category     *int32  `json:"Category,omitempty" xml:"Category,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionNo     *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResuorceType *string `json:"ResuorceType,omitempty" xml:"ResuorceType,omitempty"`
	Scope        *int32  `json:"Scope,omitempty" xml:"Scope,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeExpressConnectRouterResponseBodyEcrListTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterResponseBodyEcrListTags) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterResponseBodyEcrListTags) SetAliUid(v int64) *DescribeExpressConnectRouterResponseBodyEcrListTags {
	s.AliUid = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrListTags) SetCategory(v int32) *DescribeExpressConnectRouterResponseBodyEcrListTags {
	s.Category = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrListTags) SetId(v int64) *DescribeExpressConnectRouterResponseBodyEcrListTags {
	s.Id = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrListTags) SetRegionNo(v string) *DescribeExpressConnectRouterResponseBodyEcrListTags {
	s.RegionNo = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrListTags) SetResourceId(v string) *DescribeExpressConnectRouterResponseBodyEcrListTags {
	s.ResourceId = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrListTags) SetResuorceType(v string) *DescribeExpressConnectRouterResponseBodyEcrListTags {
	s.ResuorceType = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrListTags) SetScope(v int32) *DescribeExpressConnectRouterResponseBodyEcrListTags {
	s.Scope = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrListTags) SetTagKey(v string) *DescribeExpressConnectRouterResponseBodyEcrListTags {
	s.TagKey = &v
	return s
}

func (s *DescribeExpressConnectRouterResponseBodyEcrListTags) SetTagValue(v string) *DescribeExpressConnectRouterResponseBodyEcrListTags {
	s.TagValue = &v
	return s
}

type DescribeExpressConnectRouterResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectRouterResponse) SetStatusCode(v int32) *DescribeExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterResponse) SetBody(v *DescribeExpressConnectRouterResponseBody) *DescribeExpressConnectRouterResponse {
	s.Body = v
	return s
}

type DescribeExpressConnectRouterAllowedPrefixHistoryRequest struct {
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun        *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId         *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType  *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) SetAssociationId(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest {
	s.AssociationId = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) SetClientToken(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) SetDryRun(v bool) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) SetEcrId(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) SetInstanceId(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) SetInstanceType(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryRequest {
	s.InstanceType = &v
	return s
}

type DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody struct {
	AccessDeniedDetail       *string                                                                                 `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	AllowedPrefixHistoryList []*DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList `json:"AllowedPrefixHistoryList,omitempty" xml:"AllowedPrefixHistoryList,omitempty" type:"Repeated"`
	Code                     *string                                                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode              *string                                                                                 `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage           *string                                                                                 `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode           *int32                                                                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message                  *string                                                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId                *string                                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                  *bool                                                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetAllowedPrefixHistoryList(v []*DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.AllowedPrefixHistoryList = v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetCode(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetDynamicCode(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetDynamicMessage(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetHttpStatusCode(v int32) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetMessage(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetRequestId(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) SetSuccess(v bool) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody {
	s.Success = &v
	return s
}

type DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList struct {
	AllowedPrefix []*string `json:"AllowedPrefix,omitempty" xml:"AllowedPrefix,omitempty" type:"Repeated"`
	GmtCreate     *string   `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList) SetAllowedPrefix(v []*string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList {
	s.AllowedPrefix = v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList) SetGmtCreate(v string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBodyAllowedPrefixHistoryList {
	s.GmtCreate = &v
	return s
}

type DescribeExpressConnectRouterAllowedPrefixHistoryResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterAllowedPrefixHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectRouterAllowedPrefixHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponse) SetStatusCode(v int32) *DescribeExpressConnectRouterAllowedPrefixHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterAllowedPrefixHistoryResponse) SetBody(v *DescribeExpressConnectRouterAllowedPrefixHistoryResponseBody) *DescribeExpressConnectRouterAllowedPrefixHistoryResponse {
	s.Body = v
	return s
}

type DescribeExpressConnectRouterAssociationRequest struct {
	AssociationId       *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	AssociationNodeType *string `json:"AssociationNodeType,omitempty" xml:"AssociationNodeType,omitempty"`
	AssociationRegionId *string `json:"AssociationRegionId,omitempty" xml:"AssociationRegionId,omitempty"`
	CenId               *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun              *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId               *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	MaxResults          *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken           *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TransitRouterId     *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	VpcId               *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeExpressConnectRouterAssociationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterAssociationRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetAssociationId(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.AssociationId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetAssociationNodeType(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.AssociationNodeType = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetAssociationRegionId(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.AssociationRegionId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetCenId(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.CenId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetClientToken(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetDryRun(v bool) *DescribeExpressConnectRouterAssociationRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetEcrId(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetMaxResults(v int32) *DescribeExpressConnectRouterAssociationRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetNextToken(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetTransitRouterId(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationRequest) SetVpcId(v string) *DescribeExpressConnectRouterAssociationRequest {
	s.VpcId = &v
	return s
}

type DescribeExpressConnectRouterAssociationResponseBody struct {
	AccessDeniedDetail *string                                                               `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	AssociationList    []*DescribeExpressConnectRouterAssociationResponseBodyAssociationList `json:"AssociationList,omitempty" xml:"AssociationList,omitempty" type:"Repeated"`
	Code               *string                                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode        *string                                                               `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string                                                               `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32                                                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	MaxResults         *int32                                                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message            *string                                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken          *string                                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId          *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount         *int32                                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeExpressConnectRouterAssociationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterAssociationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetAssociationList(v []*DescribeExpressConnectRouterAssociationResponseBodyAssociationList) *DescribeExpressConnectRouterAssociationResponseBody {
	s.AssociationList = v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetCode(v string) *DescribeExpressConnectRouterAssociationResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetDynamicCode(v string) *DescribeExpressConnectRouterAssociationResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetDynamicMessage(v string) *DescribeExpressConnectRouterAssociationResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetHttpStatusCode(v int32) *DescribeExpressConnectRouterAssociationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetMaxResults(v int32) *DescribeExpressConnectRouterAssociationResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetMessage(v string) *DescribeExpressConnectRouterAssociationResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetNextToken(v string) *DescribeExpressConnectRouterAssociationResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetRequestId(v string) *DescribeExpressConnectRouterAssociationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetSuccess(v bool) *DescribeExpressConnectRouterAssociationResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBody) SetTotalCount(v int32) *DescribeExpressConnectRouterAssociationResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeExpressConnectRouterAssociationResponseBodyAssociationList struct {
	AllowedPrefixes      []*string `json:"AllowedPrefixes,omitempty" xml:"AllowedPrefixes,omitempty" type:"Repeated"`
	AssociationId        *string   `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	AssociationNodeType  *string   `json:"AssociationNodeType,omitempty" xml:"AssociationNodeType,omitempty"`
	CenId                *string   `json:"CenId,omitempty" xml:"CenId,omitempty"`
	EcrId                *string   `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	GmtCreate            *string   `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified          *string   `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status               *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TransitRouterId      *string   `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	TransitRouterOwnerId *int64    `json:"TransitRouterOwnerId,omitempty" xml:"TransitRouterOwnerId,omitempty"`
	VpcId                *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcOwnerId           *int64    `json:"VpcOwnerId,omitempty" xml:"VpcOwnerId,omitempty"`
}

func (s DescribeExpressConnectRouterAssociationResponseBodyAssociationList) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterAssociationResponseBodyAssociationList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetAllowedPrefixes(v []*string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.AllowedPrefixes = v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetAssociationId(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.AssociationId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetAssociationNodeType(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.AssociationNodeType = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetCenId(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.CenId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetEcrId(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetGmtCreate(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetGmtModified(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.GmtModified = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetOwnerId(v int64) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.OwnerId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetRegionId(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.RegionId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetStatus(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.Status = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetTransitRouterId(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetTransitRouterOwnerId(v int64) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.TransitRouterOwnerId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetVpcId(v string) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.VpcId = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponseBodyAssociationList) SetVpcOwnerId(v int64) *DescribeExpressConnectRouterAssociationResponseBodyAssociationList {
	s.VpcOwnerId = &v
	return s
}

type DescribeExpressConnectRouterAssociationResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectRouterAssociationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectRouterAssociationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterAssociationResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterAssociationResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectRouterAssociationResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponse) SetStatusCode(v int32) *DescribeExpressConnectRouterAssociationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterAssociationResponse) SetBody(v *DescribeExpressConnectRouterAssociationResponseBody) *DescribeExpressConnectRouterAssociationResponse {
	s.Body = v
	return s
}

type DescribeExpressConnectRouterChildInstanceRequest struct {
	AssociationId         *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	ChildInstanceId       *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceType     *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun                *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId                 *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	MaxResults            *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken             *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s DescribeExpressConnectRouterChildInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterChildInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetAssociationId(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.AssociationId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetChildInstanceId(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetChildInstanceRegionId(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetChildInstanceType(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetClientToken(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetDryRun(v bool) *DescribeExpressConnectRouterChildInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetEcrId(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetMaxResults(v int32) *DescribeExpressConnectRouterChildInstanceRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceRequest) SetNextToken(v string) *DescribeExpressConnectRouterChildInstanceRequest {
	s.NextToken = &v
	return s
}

type DescribeExpressConnectRouterChildInstanceResponseBody struct {
	AccessDeniedDetail *string                                                                   `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	ChildInstanceList  []*DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList `json:"ChildInstanceList,omitempty" xml:"ChildInstanceList,omitempty" type:"Repeated"`
	Code               *string                                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode        *string                                                                   `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string                                                                   `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32                                                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	MaxResults         *int32                                                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message            *string                                                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken          *string                                                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId          *string                                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount         *int32                                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeExpressConnectRouterChildInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterChildInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetChildInstanceList(v []*DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.ChildInstanceList = v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetCode(v string) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetDynamicCode(v string) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetDynamicMessage(v string) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetHttpStatusCode(v int32) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetMaxResults(v int32) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetMessage(v string) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetNextToken(v string) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetRequestId(v string) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetSuccess(v bool) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBody) SetTotalCount(v int32) *DescribeExpressConnectRouterChildInstanceResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList struct {
	AssociationId         *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	ChildInstanceId       *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceOwnerId  *int64  `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	ChildInstanceType     *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	EcrId                 *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	GmtCreate             *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified           *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetAssociationId(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.AssociationId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetChildInstanceId(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.ChildInstanceId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetChildInstanceOwnerId(v int64) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetChildInstanceRegionId(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetChildInstanceType(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.ChildInstanceType = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetEcrId(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetGmtCreate(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetGmtModified(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.GmtModified = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetOwnerId(v int64) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.OwnerId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetRegionId(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.RegionId = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList) SetStatus(v string) *DescribeExpressConnectRouterChildInstanceResponseBodyChildInstanceList {
	s.Status = &v
	return s
}

type DescribeExpressConnectRouterChildInstanceResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectRouterChildInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectRouterChildInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterChildInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterChildInstanceResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectRouterChildInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponse) SetStatusCode(v int32) *DescribeExpressConnectRouterChildInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterChildInstanceResponse) SetBody(v *DescribeExpressConnectRouterChildInstanceResponseBody) *DescribeExpressConnectRouterChildInstanceResponse {
	s.Body = v
	return s
}

type DescribeExpressConnectRouterInterRegionTransitModeRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId       *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
}

func (s DescribeExpressConnectRouterInterRegionTransitModeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterInterRegionTransitModeRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeRequest) SetClientToken(v string) *DescribeExpressConnectRouterInterRegionTransitModeRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeRequest) SetDryRun(v bool) *DescribeExpressConnectRouterInterRegionTransitModeRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeRequest) SetEcrId(v string) *DescribeExpressConnectRouterInterRegionTransitModeRequest {
	s.EcrId = &v
	return s
}

type DescribeExpressConnectRouterInterRegionTransitModeResponseBody struct {
	AccessDeniedDetail         *string                                                                                     `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code                       *string                                                                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode                *string                                                                                     `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage             *string                                                                                     `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode             *int32                                                                                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InterRegionTransitModeList []*DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList `json:"InterRegionTransitModeList,omitempty" xml:"InterRegionTransitModeList,omitempty" type:"Repeated"`
	Message                    *string                                                                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId                  *string                                                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                    *bool                                                                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeExpressConnectRouterInterRegionTransitModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterInterRegionTransitModeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetCode(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetDynamicCode(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetDynamicMessage(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetHttpStatusCode(v int32) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetInterRegionTransitModeList(v []*DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.InterRegionTransitModeList = v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetMessage(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetRequestId(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) SetSuccess(v bool) *DescribeExpressConnectRouterInterRegionTransitModeResponseBody {
	s.Success = &v
	return s
}

type DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList struct {
	Mode     *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList) SetMode(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList {
	s.Mode = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList) SetRegionId(v string) *DescribeExpressConnectRouterInterRegionTransitModeResponseBodyInterRegionTransitModeList {
	s.RegionId = &v
	return s
}

type DescribeExpressConnectRouterInterRegionTransitModeResponse struct {
	Headers    map[string]*string                                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectRouterInterRegionTransitModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectRouterInterRegionTransitModeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterInterRegionTransitModeResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectRouterInterRegionTransitModeResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponse) SetStatusCode(v int32) *DescribeExpressConnectRouterInterRegionTransitModeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterInterRegionTransitModeResponse) SetBody(v *DescribeExpressConnectRouterInterRegionTransitModeResponseBody) *DescribeExpressConnectRouterInterRegionTransitModeResponse {
	s.Body = v
	return s
}

type DescribeExpressConnectRouterRegionRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId       *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
}

func (s DescribeExpressConnectRouterRegionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterRegionRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRegionRequest) SetClientToken(v string) *DescribeExpressConnectRouterRegionRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionRequest) SetDryRun(v bool) *DescribeExpressConnectRouterRegionRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionRequest) SetEcrId(v string) *DescribeExpressConnectRouterRegionRequest {
	s.EcrId = &v
	return s
}

type DescribeExpressConnectRouterRegionResponseBody struct {
	AccessDeniedDetail *string   `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode        *string   `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string   `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RegionIdList       []*string `json:"RegionIdList,omitempty" xml:"RegionIdList,omitempty" type:"Repeated"`
	RequestId          *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeExpressConnectRouterRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterRegionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterRegionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetCode(v string) *DescribeExpressConnectRouterRegionResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetDynamicCode(v string) *DescribeExpressConnectRouterRegionResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetDynamicMessage(v string) *DescribeExpressConnectRouterRegionResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetHttpStatusCode(v int32) *DescribeExpressConnectRouterRegionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetMessage(v string) *DescribeExpressConnectRouterRegionResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetRegionIdList(v []*string) *DescribeExpressConnectRouterRegionResponseBody {
	s.RegionIdList = v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetRequestId(v string) *DescribeExpressConnectRouterRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponseBody) SetSuccess(v bool) *DescribeExpressConnectRouterRegionResponseBody {
	s.Success = &v
	return s
}

type DescribeExpressConnectRouterRegionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectRouterRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectRouterRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterRegionResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRegionResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectRouterRegionResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponse) SetStatusCode(v int32) *DescribeExpressConnectRouterRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterRegionResponse) SetBody(v *DescribeExpressConnectRouterRegionResponseBody) *DescribeExpressConnectRouterRegionResponse {
	s.Body = v
	return s
}

type DescribeExpressConnectRouterRouteEntriesRequest struct {
	AsPath               *string `json:"AsPath,omitempty" xml:"AsPath,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Community            *string `json:"Community,omitempty" xml:"Community,omitempty"`
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId                *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	MaxResults           *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	NexthopInstanceId    *string `json:"NexthopInstanceId,omitempty" xml:"NexthopInstanceId,omitempty"`
	QueryRegionId        *string `json:"QueryRegionId,omitempty" xml:"QueryRegionId,omitempty"`
}

func (s DescribeExpressConnectRouterRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetAsPath(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.AsPath = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetClientToken(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetCommunity(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.Community = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetDestinationCidrBlock(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetDryRun(v bool) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetEcrId(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetMaxResults(v int32) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetNextToken(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetNexthopInstanceId(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.NexthopInstanceId = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesRequest) SetQueryRegionId(v string) *DescribeExpressConnectRouterRouteEntriesRequest {
	s.QueryRegionId = &v
	return s
}

type DescribeExpressConnectRouterRouteEntriesResponseBody struct {
	AccessDeniedDetail *string                                                                 `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode        *string                                                                 `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string                                                                 `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32                                                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	MaxResults         *int32                                                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message            *string                                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken          *string                                                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId          *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RouteEntriesList   []*DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList `json:"RouteEntriesList,omitempty" xml:"RouteEntriesList,omitempty" type:"Repeated"`
	Success            *bool                                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount         *int32                                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeExpressConnectRouterRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetAccessDeniedDetail(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetCode(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetDynamicCode(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetDynamicMessage(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetHttpStatusCode(v int32) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetMaxResults(v int32) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetMessage(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetNextToken(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetRequestId(v string) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetRouteEntriesList(v []*DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.RouteEntriesList = v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetSuccess(v bool) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBody) SetTotalCount(v int32) *DescribeExpressConnectRouterRouteEntriesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList struct {
	AsPath                  *string `json:"AsPath,omitempty" xml:"AsPath,omitempty"`
	Community               *string `json:"Community,omitempty" xml:"Community,omitempty"`
	DestinationCidrBlock    *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	NexthopInstanceId       *string `json:"NexthopInstanceId,omitempty" xml:"NexthopInstanceId,omitempty"`
	NexthopInstanceRegionId *string `json:"NexthopInstanceRegionId,omitempty" xml:"NexthopInstanceRegionId,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) SetAsPath(v string) *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList {
	s.AsPath = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) SetCommunity(v string) *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList {
	s.Community = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) SetDestinationCidrBlock(v string) *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) SetNexthopInstanceId(v string) *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList {
	s.NexthopInstanceId = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) SetNexthopInstanceRegionId(v string) *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList {
	s.NexthopInstanceRegionId = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList) SetStatus(v string) *DescribeExpressConnectRouterRouteEntriesResponseBodyRouteEntriesList {
	s.Status = &v
	return s
}

type DescribeExpressConnectRouterRouteEntriesResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpressConnectRouterRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpressConnectRouterRouteEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExpressConnectRouterRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectRouterRouteEntriesResponse) SetHeaders(v map[string]*string) *DescribeExpressConnectRouterRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponse) SetStatusCode(v int32) *DescribeExpressConnectRouterRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpressConnectRouterRouteEntriesResponse) SetBody(v *DescribeExpressConnectRouterRouteEntriesResponseBody) *DescribeExpressConnectRouterRouteEntriesResponse {
	s.Body = v
	return s
}

type DescribeInstanceGrantedToExpressConnectRouterRequest struct {
	ClientToken      *string                                                          `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun           *bool                                                            `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId            *string                                                          `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	InstanceId       *string                                                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceOwnerId  *int64                                                           `json:"InstanceOwnerId,omitempty" xml:"InstanceOwnerId,omitempty"`
	InstanceRegionId *string                                                          `json:"InstanceRegionId,omitempty" xml:"InstanceRegionId,omitempty"`
	InstanceType     *string                                                          `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	MaxResults       *int32                                                           `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string                                                          `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceGroupId  *string                                                          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	TagModels        []*DescribeInstanceGrantedToExpressConnectRouterRequestTagModels `json:"TagModels,omitempty" xml:"TagModels,omitempty" type:"Repeated"`
}

func (s DescribeInstanceGrantedToExpressConnectRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceGrantedToExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetClientToken(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetDryRun(v bool) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetEcrId(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetInstanceId(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetInstanceOwnerId(v int64) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.InstanceOwnerId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetInstanceRegionId(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.InstanceRegionId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetInstanceType(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetMaxResults(v int32) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetNextToken(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetResourceGroupId(v string) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequest) SetTagModels(v []*DescribeInstanceGrantedToExpressConnectRouterRequestTagModels) *DescribeInstanceGrantedToExpressConnectRouterRequest {
	s.TagModels = v
	return s
}

type DescribeInstanceGrantedToExpressConnectRouterRequestTagModels struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeInstanceGrantedToExpressConnectRouterRequestTagModels) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceGrantedToExpressConnectRouterRequestTagModels) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequestTagModels) SetTagKey(v string) *DescribeInstanceGrantedToExpressConnectRouterRequestTagModels {
	s.TagKey = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterRequestTagModels) SetTagValue(v string) *DescribeInstanceGrantedToExpressConnectRouterRequestTagModels {
	s.TagValue = &v
	return s
}

type DescribeInstanceGrantedToExpressConnectRouterResponseBody struct {
	AccessDeniedDetail     *string                                                                            `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code                   *string                                                                            `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode            *string                                                                            `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage         *string                                                                            `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	EcrGrantedInstanceList []*DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList `json:"EcrGrantedInstanceList,omitempty" xml:"EcrGrantedInstanceList,omitempty" type:"Repeated"`
	HttpStatusCode         *int32                                                                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	MaxResults             *int32                                                                             `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message                *string                                                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken              *string                                                                            `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId              *string                                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                *bool                                                                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount             *int32                                                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceGrantedToExpressConnectRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceGrantedToExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetCode(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetDynamicCode(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetDynamicMessage(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetEcrGrantedInstanceList(v []*DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.EcrGrantedInstanceList = v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetMaxResults(v int32) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetMessage(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetNextToken(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetRequestId(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetSuccess(v bool) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBody) SetTotalCount(v int32) *DescribeInstanceGrantedToExpressConnectRouterResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList struct {
	EcrId        *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	GmtCreate    *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GrantId      *string `json:"GrantId,omitempty" xml:"GrantId,omitempty"`
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeOwnerBid *string `json:"NodeOwnerBid,omitempty" xml:"NodeOwnerBid,omitempty"`
	NodeOwnerUid *int64  `json:"NodeOwnerUid,omitempty" xml:"NodeOwnerUid,omitempty"`
	NodeRegionId *string `json:"NodeRegionId,omitempty" xml:"NodeRegionId,omitempty"`
	NodeType     *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetEcrId(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.EcrId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetGmtCreate(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetGmtModified(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.GmtModified = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetGrantId(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.GrantId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetNodeId(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.NodeId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetNodeOwnerBid(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.NodeOwnerBid = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetNodeOwnerUid(v int64) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.NodeOwnerUid = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetNodeRegionId(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.NodeRegionId = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetNodeType(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.NodeType = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList) SetStatus(v string) *DescribeInstanceGrantedToExpressConnectRouterResponseBodyEcrGrantedInstanceList {
	s.Status = &v
	return s
}

type DescribeInstanceGrantedToExpressConnectRouterResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceGrantedToExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceGrantedToExpressConnectRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceGrantedToExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponse) SetHeaders(v map[string]*string) *DescribeInstanceGrantedToExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponse) SetStatusCode(v int32) *DescribeInstanceGrantedToExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceGrantedToExpressConnectRouterResponse) SetBody(v *DescribeInstanceGrantedToExpressConnectRouterResponseBody) *DescribeInstanceGrantedToExpressConnectRouterResponse {
	s.Body = v
	return s
}

type DetachExpressConnectRouterChildInstanceRequest struct {
	ChildInstanceId   *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	ChildInstanceType *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun            *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId             *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
}

func (s DetachExpressConnectRouterChildInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachExpressConnectRouterChildInstanceRequest) GoString() string {
	return s.String()
}

func (s *DetachExpressConnectRouterChildInstanceRequest) SetChildInstanceId(v string) *DetachExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceRequest) SetChildInstanceType(v string) *DetachExpressConnectRouterChildInstanceRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceRequest) SetClientToken(v string) *DetachExpressConnectRouterChildInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceRequest) SetDryRun(v bool) *DetachExpressConnectRouterChildInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceRequest) SetEcrId(v string) *DetachExpressConnectRouterChildInstanceRequest {
	s.EcrId = &v
	return s
}

type DetachExpressConnectRouterChildInstanceResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode        *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetachExpressConnectRouterChildInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachExpressConnectRouterChildInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetAccessDeniedDetail(v string) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetCode(v string) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetDynamicCode(v string) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetDynamicMessage(v string) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetHttpStatusCode(v int32) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetMessage(v string) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetRequestId(v string) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponseBody) SetSuccess(v bool) *DetachExpressConnectRouterChildInstanceResponseBody {
	s.Success = &v
	return s
}

type DetachExpressConnectRouterChildInstanceResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachExpressConnectRouterChildInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachExpressConnectRouterChildInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachExpressConnectRouterChildInstanceResponse) GoString() string {
	return s.String()
}

func (s *DetachExpressConnectRouterChildInstanceResponse) SetHeaders(v map[string]*string) *DetachExpressConnectRouterChildInstanceResponse {
	s.Headers = v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponse) SetStatusCode(v int32) *DetachExpressConnectRouterChildInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachExpressConnectRouterChildInstanceResponse) SetBody(v *DetachExpressConnectRouterChildInstanceResponseBody) *DetachExpressConnectRouterChildInstanceResponse {
	s.Body = v
	return s
}

type DisableExpressConnectRouterRouteEntriesRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId                *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	NexthopInstanceId    *string `json:"NexthopInstanceId,omitempty" xml:"NexthopInstanceId,omitempty"`
}

func (s DisableExpressConnectRouterRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableExpressConnectRouterRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) SetClientToken(v string) *DisableExpressConnectRouterRouteEntriesRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) SetDestinationCidrBlock(v string) *DisableExpressConnectRouterRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) SetDryRun(v bool) *DisableExpressConnectRouterRouteEntriesRequest {
	s.DryRun = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) SetEcrId(v string) *DisableExpressConnectRouterRouteEntriesRequest {
	s.EcrId = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesRequest) SetNexthopInstanceId(v string) *DisableExpressConnectRouterRouteEntriesRequest {
	s.NexthopInstanceId = &v
	return s
}

type DisableExpressConnectRouterRouteEntriesResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode        *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableExpressConnectRouterRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableExpressConnectRouterRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetAccessDeniedDetail(v string) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetCode(v string) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.Code = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetDynamicCode(v string) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetDynamicMessage(v string) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetHttpStatusCode(v int32) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetMessage(v string) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.Message = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetRequestId(v string) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponseBody) SetSuccess(v bool) *DisableExpressConnectRouterRouteEntriesResponseBody {
	s.Success = &v
	return s
}

type DisableExpressConnectRouterRouteEntriesResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableExpressConnectRouterRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableExpressConnectRouterRouteEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableExpressConnectRouterRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *DisableExpressConnectRouterRouteEntriesResponse) SetHeaders(v map[string]*string) *DisableExpressConnectRouterRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponse) SetStatusCode(v int32) *DisableExpressConnectRouterRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableExpressConnectRouterRouteEntriesResponse) SetBody(v *DisableExpressConnectRouterRouteEntriesResponseBody) *DisableExpressConnectRouterRouteEntriesResponse {
	s.Body = v
	return s
}

type EnableExpressConnectRouterRouteEntriesRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId                *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	NexthopInstanceId    *string `json:"NexthopInstanceId,omitempty" xml:"NexthopInstanceId,omitempty"`
}

func (s EnableExpressConnectRouterRouteEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableExpressConnectRouterRouteEntriesRequest) GoString() string {
	return s.String()
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) SetClientToken(v string) *EnableExpressConnectRouterRouteEntriesRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) SetDestinationCidrBlock(v string) *EnableExpressConnectRouterRouteEntriesRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) SetDryRun(v bool) *EnableExpressConnectRouterRouteEntriesRequest {
	s.DryRun = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) SetEcrId(v string) *EnableExpressConnectRouterRouteEntriesRequest {
	s.EcrId = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesRequest) SetNexthopInstanceId(v string) *EnableExpressConnectRouterRouteEntriesRequest {
	s.NexthopInstanceId = &v
	return s
}

type EnableExpressConnectRouterRouteEntriesResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode        *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableExpressConnectRouterRouteEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableExpressConnectRouterRouteEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetAccessDeniedDetail(v string) *EnableExpressConnectRouterRouteEntriesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetCode(v string) *EnableExpressConnectRouterRouteEntriesResponseBody {
	s.Code = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetDynamicCode(v string) *EnableExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetDynamicMessage(v string) *EnableExpressConnectRouterRouteEntriesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetHttpStatusCode(v int32) *EnableExpressConnectRouterRouteEntriesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetMessage(v string) *EnableExpressConnectRouterRouteEntriesResponseBody {
	s.Message = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetRequestId(v string) *EnableExpressConnectRouterRouteEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponseBody) SetSuccess(v bool) *EnableExpressConnectRouterRouteEntriesResponseBody {
	s.Success = &v
	return s
}

type EnableExpressConnectRouterRouteEntriesResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableExpressConnectRouterRouteEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableExpressConnectRouterRouteEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableExpressConnectRouterRouteEntriesResponse) GoString() string {
	return s.String()
}

func (s *EnableExpressConnectRouterRouteEntriesResponse) SetHeaders(v map[string]*string) *EnableExpressConnectRouterRouteEntriesResponse {
	s.Headers = v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponse) SetStatusCode(v int32) *EnableExpressConnectRouterRouteEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableExpressConnectRouterRouteEntriesResponse) SetBody(v *EnableExpressConnectRouterRouteEntriesResponseBody) *EnableExpressConnectRouterRouteEntriesResponse {
	s.Body = v
	return s
}

type ForceDeleteExpressConnectRouterRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId       *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
}

func (s ForceDeleteExpressConnectRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s ForceDeleteExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *ForceDeleteExpressConnectRouterRequest) SetClientToken(v string) *ForceDeleteExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterRequest) SetDryRun(v bool) *ForceDeleteExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterRequest) SetEcrId(v string) *ForceDeleteExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

type ForceDeleteExpressConnectRouterResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode        *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ForceDeleteExpressConnectRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ForceDeleteExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *ForceDeleteExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetCode(v string) *ForceDeleteExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetDynamicCode(v string) *ForceDeleteExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetDynamicMessage(v string) *ForceDeleteExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *ForceDeleteExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetMessage(v string) *ForceDeleteExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetRequestId(v string) *ForceDeleteExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponseBody) SetSuccess(v bool) *ForceDeleteExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

type ForceDeleteExpressConnectRouterResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ForceDeleteExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ForceDeleteExpressConnectRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s ForceDeleteExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *ForceDeleteExpressConnectRouterResponse) SetHeaders(v map[string]*string) *ForceDeleteExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponse) SetStatusCode(v int32) *ForceDeleteExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *ForceDeleteExpressConnectRouterResponse) SetBody(v *ForceDeleteExpressConnectRouterResponseBody) *ForceDeleteExpressConnectRouterResponse {
	s.Body = v
	return s
}

type GrantInstanceToExpressConnectRouterRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun           *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId            *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	EcrOwnerAliUid   *int64  `json:"EcrOwnerAliUid,omitempty" xml:"EcrOwnerAliUid,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceRegionId *string `json:"InstanceRegionId,omitempty" xml:"InstanceRegionId,omitempty"`
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s GrantInstanceToExpressConnectRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantInstanceToExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *GrantInstanceToExpressConnectRouterRequest) SetClientToken(v string) *GrantInstanceToExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterRequest) SetDryRun(v bool) *GrantInstanceToExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterRequest) SetEcrId(v string) *GrantInstanceToExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterRequest) SetEcrOwnerAliUid(v int64) *GrantInstanceToExpressConnectRouterRequest {
	s.EcrOwnerAliUid = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterRequest) SetInstanceId(v string) *GrantInstanceToExpressConnectRouterRequest {
	s.InstanceId = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterRequest) SetInstanceRegionId(v string) *GrantInstanceToExpressConnectRouterRequest {
	s.InstanceRegionId = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterRequest) SetInstanceType(v string) *GrantInstanceToExpressConnectRouterRequest {
	s.InstanceType = &v
	return s
}

type GrantInstanceToExpressConnectRouterResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode        *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GrantInstanceToExpressConnectRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantInstanceToExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *GrantInstanceToExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetCode(v string) *GrantInstanceToExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetDynamicCode(v string) *GrantInstanceToExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetDynamicMessage(v string) *GrantInstanceToExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *GrantInstanceToExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetMessage(v string) *GrantInstanceToExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetRequestId(v string) *GrantInstanceToExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponseBody) SetSuccess(v bool) *GrantInstanceToExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

type GrantInstanceToExpressConnectRouterResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantInstanceToExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantInstanceToExpressConnectRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantInstanceToExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *GrantInstanceToExpressConnectRouterResponse) SetHeaders(v map[string]*string) *GrantInstanceToExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponse) SetStatusCode(v int32) *GrantInstanceToExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantInstanceToExpressConnectRouterResponse) SetBody(v *GrantInstanceToExpressConnectRouterResponseBody) *GrantInstanceToExpressConnectRouterResponse {
	s.Body = v
	return s
}

type ListExpressConnectRouterSupportedRegionRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	NodeType    *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s ListExpressConnectRouterSupportedRegionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExpressConnectRouterSupportedRegionRequest) GoString() string {
	return s.String()
}

func (s *ListExpressConnectRouterSupportedRegionRequest) SetClientToken(v string) *ListExpressConnectRouterSupportedRegionRequest {
	s.ClientToken = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionRequest) SetNodeType(v string) *ListExpressConnectRouterSupportedRegionRequest {
	s.NodeType = &v
	return s
}

type ListExpressConnectRouterSupportedRegionResponseBody struct {
	Code                  *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode        *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message               *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId             *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success               *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
	SupportedRegionIdList []*string `json:"SupportedRegionIdList,omitempty" xml:"SupportedRegionIdList,omitempty" type:"Repeated"`
}

func (s ListExpressConnectRouterSupportedRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExpressConnectRouterSupportedRegionResponseBody) GoString() string {
	return s.String()
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) SetCode(v string) *ListExpressConnectRouterSupportedRegionResponseBody {
	s.Code = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) SetHttpStatusCode(v int32) *ListExpressConnectRouterSupportedRegionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) SetMessage(v string) *ListExpressConnectRouterSupportedRegionResponseBody {
	s.Message = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) SetRequestId(v string) *ListExpressConnectRouterSupportedRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) SetSuccess(v bool) *ListExpressConnectRouterSupportedRegionResponseBody {
	s.Success = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponseBody) SetSupportedRegionIdList(v []*string) *ListExpressConnectRouterSupportedRegionResponseBody {
	s.SupportedRegionIdList = v
	return s
}

type ListExpressConnectRouterSupportedRegionResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExpressConnectRouterSupportedRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExpressConnectRouterSupportedRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExpressConnectRouterSupportedRegionResponse) GoString() string {
	return s.String()
}

func (s *ListExpressConnectRouterSupportedRegionResponse) SetHeaders(v map[string]*string) *ListExpressConnectRouterSupportedRegionResponse {
	s.Headers = v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponse) SetStatusCode(v int32) *ListExpressConnectRouterSupportedRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExpressConnectRouterSupportedRegionResponse) SetBody(v *ListExpressConnectRouterSupportedRegionResponseBody) *ListExpressConnectRouterSupportedRegionResponse {
	s.Body = v
	return s
}

type ModifyExpressConnectRouterRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId       *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyExpressConnectRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterRequest) SetClientToken(v string) *ModifyExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyExpressConnectRouterRequest) SetDescription(v string) *ModifyExpressConnectRouterRequest {
	s.Description = &v
	return s
}

func (s *ModifyExpressConnectRouterRequest) SetDryRun(v bool) *ModifyExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyExpressConnectRouterRequest) SetEcrId(v string) *ModifyExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

func (s *ModifyExpressConnectRouterRequest) SetName(v string) *ModifyExpressConnectRouterRequest {
	s.Name = &v
	return s
}

type ModifyExpressConnectRouterResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode        *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyExpressConnectRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *ModifyExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) SetCode(v string) *ModifyExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) SetDynamicCode(v string) *ModifyExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) SetDynamicMessage(v string) *ModifyExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *ModifyExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) SetMessage(v string) *ModifyExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) SetRequestId(v string) *ModifyExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyExpressConnectRouterResponseBody) SetSuccess(v bool) *ModifyExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

type ModifyExpressConnectRouterResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyExpressConnectRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterResponse) SetHeaders(v map[string]*string) *ModifyExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *ModifyExpressConnectRouterResponse) SetStatusCode(v int32) *ModifyExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterResponse) SetBody(v *ModifyExpressConnectRouterResponseBody) *ModifyExpressConnectRouterResponse {
	s.Body = v
	return s
}

type ModifyExpressConnectRouterAssociationAllowedPrefixRequest struct {
	AllowedPrefixes []*string `json:"AllowedPrefixes,omitempty" xml:"AllowedPrefixes,omitempty" type:"Repeated"`
	AssociationId   *string   `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	ClientToken     *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun          *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId           *string   `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	OwnerAccount    *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s ModifyExpressConnectRouterAssociationAllowedPrefixRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyExpressConnectRouterAssociationAllowedPrefixRequest) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) SetAllowedPrefixes(v []*string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest {
	s.AllowedPrefixes = v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) SetAssociationId(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest {
	s.AssociationId = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) SetClientToken(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) SetDryRun(v bool) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) SetEcrId(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest {
	s.EcrId = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) SetOwnerAccount(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest {
	s.OwnerAccount = &v
	return s
}

type ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode        *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetAccessDeniedDetail(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetCode(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetDynamicCode(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetDynamicMessage(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetHttpStatusCode(v int32) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetMessage(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetRequestId(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) SetSuccess(v bool) *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody {
	s.Success = &v
	return s
}

type ModifyExpressConnectRouterAssociationAllowedPrefixResponse struct {
	Headers    map[string]*string                                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyExpressConnectRouterAssociationAllowedPrefixResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyExpressConnectRouterAssociationAllowedPrefixResponse) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponse) SetHeaders(v map[string]*string) *ModifyExpressConnectRouterAssociationAllowedPrefixResponse {
	s.Headers = v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponse) SetStatusCode(v int32) *ModifyExpressConnectRouterAssociationAllowedPrefixResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixResponse) SetBody(v *ModifyExpressConnectRouterAssociationAllowedPrefixResponseBody) *ModifyExpressConnectRouterAssociationAllowedPrefixResponse {
	s.Body = v
	return s
}

type ModifyExpressConnectRouterInterRegionTransitModeRequest struct {
	ClientToken     *string                                                                   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun          *bool                                                                     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId           *string                                                                   `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	TransitModeList []*ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList `json:"TransitModeList,omitempty" xml:"TransitModeList,omitempty" type:"Repeated"`
}

func (s ModifyExpressConnectRouterInterRegionTransitModeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyExpressConnectRouterInterRegionTransitModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequest) SetClientToken(v string) *ModifyExpressConnectRouterInterRegionTransitModeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequest) SetDryRun(v bool) *ModifyExpressConnectRouterInterRegionTransitModeRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequest) SetEcrId(v string) *ModifyExpressConnectRouterInterRegionTransitModeRequest {
	s.EcrId = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequest) SetTransitModeList(v []*ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList) *ModifyExpressConnectRouterInterRegionTransitModeRequest {
	s.TransitModeList = v
	return s
}

type ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList struct {
	Mode     *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList) String() string {
	return tea.Prettify(s)
}

func (s ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList) SetMode(v string) *ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList {
	s.Mode = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList) SetRegionId(v string) *ModifyExpressConnectRouterInterRegionTransitModeRequestTransitModeList {
	s.RegionId = &v
	return s
}

type ModifyExpressConnectRouterInterRegionTransitModeResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode        *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyExpressConnectRouterInterRegionTransitModeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyExpressConnectRouterInterRegionTransitModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetAccessDeniedDetail(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetCode(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetDynamicCode(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetDynamicMessage(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetHttpStatusCode(v int32) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetMessage(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetRequestId(v string) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) SetSuccess(v bool) *ModifyExpressConnectRouterInterRegionTransitModeResponseBody {
	s.Success = &v
	return s
}

type ModifyExpressConnectRouterInterRegionTransitModeResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyExpressConnectRouterInterRegionTransitModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyExpressConnectRouterInterRegionTransitModeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyExpressConnectRouterInterRegionTransitModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponse) SetHeaders(v map[string]*string) *ModifyExpressConnectRouterInterRegionTransitModeResponse {
	s.Headers = v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponse) SetStatusCode(v int32) *ModifyExpressConnectRouterInterRegionTransitModeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyExpressConnectRouterInterRegionTransitModeResponse) SetBody(v *ModifyExpressConnectRouterInterRegionTransitModeResponseBody) *ModifyExpressConnectRouterInterRegionTransitModeResponse {
	s.Body = v
	return s
}

type RevokeInstanceFromExpressConnectRouterRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun           *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId            *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	EcrOwnerAliUid   *int64  `json:"EcrOwnerAliUid,omitempty" xml:"EcrOwnerAliUid,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceRegionId *string `json:"InstanceRegionId,omitempty" xml:"InstanceRegionId,omitempty"`
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s RevokeInstanceFromExpressConnectRouterRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeInstanceFromExpressConnectRouterRequest) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetClientToken(v string) *RevokeInstanceFromExpressConnectRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetDryRun(v bool) *RevokeInstanceFromExpressConnectRouterRequest {
	s.DryRun = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetEcrId(v string) *RevokeInstanceFromExpressConnectRouterRequest {
	s.EcrId = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetEcrOwnerAliUid(v int64) *RevokeInstanceFromExpressConnectRouterRequest {
	s.EcrOwnerAliUid = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetInstanceId(v string) *RevokeInstanceFromExpressConnectRouterRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetInstanceRegionId(v string) *RevokeInstanceFromExpressConnectRouterRequest {
	s.InstanceRegionId = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterRequest) SetInstanceType(v string) *RevokeInstanceFromExpressConnectRouterRequest {
	s.InstanceType = &v
	return s
}

type RevokeInstanceFromExpressConnectRouterResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode        *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RevokeInstanceFromExpressConnectRouterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeInstanceFromExpressConnectRouterResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetAccessDeniedDetail(v string) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetCode(v string) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.Code = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetDynamicCode(v string) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetDynamicMessage(v string) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetHttpStatusCode(v int32) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetMessage(v string) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.Message = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetRequestId(v string) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponseBody) SetSuccess(v bool) *RevokeInstanceFromExpressConnectRouterResponseBody {
	s.Success = &v
	return s
}

type RevokeInstanceFromExpressConnectRouterResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeInstanceFromExpressConnectRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeInstanceFromExpressConnectRouterResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeInstanceFromExpressConnectRouterResponse) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromExpressConnectRouterResponse) SetHeaders(v map[string]*string) *RevokeInstanceFromExpressConnectRouterResponse {
	s.Headers = v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponse) SetStatusCode(v int32) *RevokeInstanceFromExpressConnectRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeInstanceFromExpressConnectRouterResponse) SetBody(v *RevokeInstanceFromExpressConnectRouterResponseBody) *RevokeInstanceFromExpressConnectRouterResponse {
	s.Body = v
	return s
}

type SynchronizeExpressConnectRouterInterRegionBandwidthRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DryRun      *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EcrId       *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
}

func (s SynchronizeExpressConnectRouterInterRegionBandwidthRequest) String() string {
	return tea.Prettify(s)
}

func (s SynchronizeExpressConnectRouterInterRegionBandwidthRequest) GoString() string {
	return s.String()
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthRequest) SetClientToken(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthRequest {
	s.ClientToken = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthRequest) SetDryRun(v bool) *SynchronizeExpressConnectRouterInterRegionBandwidthRequest {
	s.DryRun = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthRequest) SetEcrId(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthRequest {
	s.EcrId = &v
	return s
}

type SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicCode        *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetAccessDeniedDetail(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetCode(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.Code = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetDynamicCode(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetDynamicMessage(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetHttpStatusCode(v int32) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetMessage(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.Message = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetRequestId(v string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) SetSuccess(v bool) *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody {
	s.Success = &v
	return s
}

type SynchronizeExpressConnectRouterInterRegionBandwidthResponse struct {
	Headers    map[string]*string                                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SynchronizeExpressConnectRouterInterRegionBandwidthResponse) String() string {
	return tea.Prettify(s)
}

func (s SynchronizeExpressConnectRouterInterRegionBandwidthResponse) GoString() string {
	return s.String()
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponse) SetHeaders(v map[string]*string) *SynchronizeExpressConnectRouterInterRegionBandwidthResponse {
	s.Headers = v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponse) SetStatusCode(v int32) *SynchronizeExpressConnectRouterInterRegionBandwidthResponse {
	s.StatusCode = &v
	return s
}

func (s *SynchronizeExpressConnectRouterInterRegionBandwidthResponse) SetBody(v *SynchronizeExpressConnectRouterInterRegionBandwidthResponseBody) *SynchronizeExpressConnectRouterInterRegionBandwidthResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("expressconnectrouter"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AttachExpressConnectRouterChildInstanceWithOptions(request *AttachExpressConnectRouterChildInstanceRequest, runtime *util.RuntimeOptions) (_result *AttachExpressConnectRouterChildInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChildInstanceId)) {
		body["ChildInstanceId"] = request.ChildInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ChildInstanceOwnerId)) {
		body["ChildInstanceOwnerId"] = request.ChildInstanceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ChildInstanceRegionId)) {
		body["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ChildInstanceType)) {
		body["ChildInstanceType"] = request.ChildInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachExpressConnectRouterChildInstance"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachExpressConnectRouterChildInstance(request *AttachExpressConnectRouterChildInstanceRequest) (_result *AttachExpressConnectRouterChildInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.AttachExpressConnectRouterChildInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckAddRegionToExpressConnectRouterWithOptions(request *CheckAddRegionToExpressConnectRouterRequest, runtime *util.RuntimeOptions) (_result *CheckAddRegionToExpressConnectRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.FreshRegionId)) {
		body["FreshRegionId"] = request.FreshRegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckAddRegionToExpressConnectRouter"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckAddRegionToExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckAddRegionToExpressConnectRouter(request *CheckAddRegionToExpressConnectRouterRequest) (_result *CheckAddRegionToExpressConnectRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckAddRegionToExpressConnectRouterResponse{}
	_body, _err := client.CheckAddRegionToExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateExpressConnectRouterWithOptions(request *CreateExpressConnectRouterRequest, runtime *util.RuntimeOptions) (_result *CreateExpressConnectRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlibabaSideAsn)) {
		body["AlibabaSideAsn"] = request.AlibabaSideAsn
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateExpressConnectRouter"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateExpressConnectRouter(request *CreateExpressConnectRouterRequest) (_result *CreateExpressConnectRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateExpressConnectRouterResponse{}
	_body, _err := client.CreateExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateExpressConnectRouterAssociationWithOptions(request *CreateExpressConnectRouterAssociationRequest, runtime *util.RuntimeOptions) (_result *CreateExpressConnectRouterAssociationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowedPrefixes)) {
		body["AllowedPrefixes"] = request.AllowedPrefixes
	}

	if !tea.BoolValue(util.IsUnset(request.AssociationRegionId)) {
		body["AssociationRegionId"] = request.AssociationRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.CenId)) {
		body["CenId"] = request.CenId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CreateAttachment)) {
		body["CreateAttachment"] = request.CreateAttachment
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.TransitRouterId)) {
		body["TransitRouterId"] = request.TransitRouterId
	}

	if !tea.BoolValue(util.IsUnset(request.TransitRouterOwnerId)) {
		body["TransitRouterOwnerId"] = request.TransitRouterOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcOwnerId)) {
		body["VpcOwnerId"] = request.VpcOwnerId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateExpressConnectRouterAssociation"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateExpressConnectRouterAssociationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateExpressConnectRouterAssociation(request *CreateExpressConnectRouterAssociationRequest) (_result *CreateExpressConnectRouterAssociationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateExpressConnectRouterAssociationResponse{}
	_body, _err := client.CreateExpressConnectRouterAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteExpressConnectRouterWithOptions(request *DeleteExpressConnectRouterRequest, runtime *util.RuntimeOptions) (_result *DeleteExpressConnectRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteExpressConnectRouter"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteExpressConnectRouter(request *DeleteExpressConnectRouterRequest) (_result *DeleteExpressConnectRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteExpressConnectRouterResponse{}
	_body, _err := client.DeleteExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteExpressConnectRouterAssociationWithOptions(request *DeleteExpressConnectRouterAssociationRequest, runtime *util.RuntimeOptions) (_result *DeleteExpressConnectRouterAssociationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssociationId)) {
		body["AssociationId"] = request.AssociationId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DeleteAttachment)) {
		body["DeleteAttachment"] = request.DeleteAttachment
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteExpressConnectRouterAssociation"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteExpressConnectRouterAssociationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteExpressConnectRouterAssociation(request *DeleteExpressConnectRouterAssociationRequest) (_result *DeleteExpressConnectRouterAssociationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteExpressConnectRouterAssociationResponse{}
	_body, _err := client.DeleteExpressConnectRouterAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDisabledExpressConnectRouterRouteEntriesWithOptions(request *DescribeDisabledExpressConnectRouterRouteEntriesRequest, runtime *util.RuntimeOptions) (_result *DescribeDisabledExpressConnectRouterRouteEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDisabledExpressConnectRouterRouteEntries"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDisabledExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDisabledExpressConnectRouterRouteEntries(request *DescribeDisabledExpressConnectRouterRouteEntriesRequest) (_result *DescribeDisabledExpressConnectRouterRouteEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDisabledExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.DescribeDisabledExpressConnectRouterRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExpressConnectRouterWithOptions(request *DescribeExpressConnectRouterRequest, runtime *util.RuntimeOptions) (_result *DescribeExpressConnectRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TagModels)) {
		body["TagModels"] = request.TagModels
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExpressConnectRouter"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExpressConnectRouter(request *DescribeExpressConnectRouterRequest) (_result *DescribeExpressConnectRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterResponse{}
	_body, _err := client.DescribeExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExpressConnectRouterAllowedPrefixHistoryWithOptions(request *DescribeExpressConnectRouterAllowedPrefixHistoryRequest, runtime *util.RuntimeOptions) (_result *DescribeExpressConnectRouterAllowedPrefixHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssociationId)) {
		body["AssociationId"] = request.AssociationId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["InstanceType"] = request.InstanceType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExpressConnectRouterAllowedPrefixHistory"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeExpressConnectRouterAllowedPrefixHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExpressConnectRouterAllowedPrefixHistory(request *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) (_result *DescribeExpressConnectRouterAllowedPrefixHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterAllowedPrefixHistoryResponse{}
	_body, _err := client.DescribeExpressConnectRouterAllowedPrefixHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExpressConnectRouterAssociationWithOptions(request *DescribeExpressConnectRouterAssociationRequest, runtime *util.RuntimeOptions) (_result *DescribeExpressConnectRouterAssociationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssociationId)) {
		body["AssociationId"] = request.AssociationId
	}

	if !tea.BoolValue(util.IsUnset(request.AssociationNodeType)) {
		body["AssociationNodeType"] = request.AssociationNodeType
	}

	if !tea.BoolValue(util.IsUnset(request.AssociationRegionId)) {
		body["AssociationRegionId"] = request.AssociationRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.CenId)) {
		body["CenId"] = request.CenId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.TransitRouterId)) {
		body["TransitRouterId"] = request.TransitRouterId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExpressConnectRouterAssociation"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeExpressConnectRouterAssociationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExpressConnectRouterAssociation(request *DescribeExpressConnectRouterAssociationRequest) (_result *DescribeExpressConnectRouterAssociationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterAssociationResponse{}
	_body, _err := client.DescribeExpressConnectRouterAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExpressConnectRouterChildInstanceWithOptions(request *DescribeExpressConnectRouterChildInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeExpressConnectRouterChildInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssociationId)) {
		body["AssociationId"] = request.AssociationId
	}

	if !tea.BoolValue(util.IsUnset(request.ChildInstanceId)) {
		body["ChildInstanceId"] = request.ChildInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ChildInstanceRegionId)) {
		body["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ChildInstanceType)) {
		body["ChildInstanceType"] = request.ChildInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExpressConnectRouterChildInstance"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExpressConnectRouterChildInstance(request *DescribeExpressConnectRouterChildInstanceRequest) (_result *DescribeExpressConnectRouterChildInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.DescribeExpressConnectRouterChildInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExpressConnectRouterInterRegionTransitModeWithOptions(request *DescribeExpressConnectRouterInterRegionTransitModeRequest, runtime *util.RuntimeOptions) (_result *DescribeExpressConnectRouterInterRegionTransitModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExpressConnectRouterInterRegionTransitMode"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeExpressConnectRouterInterRegionTransitModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExpressConnectRouterInterRegionTransitMode(request *DescribeExpressConnectRouterInterRegionTransitModeRequest) (_result *DescribeExpressConnectRouterInterRegionTransitModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterInterRegionTransitModeResponse{}
	_body, _err := client.DescribeExpressConnectRouterInterRegionTransitModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExpressConnectRouterRegionWithOptions(request *DescribeExpressConnectRouterRegionRequest, runtime *util.RuntimeOptions) (_result *DescribeExpressConnectRouterRegionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExpressConnectRouterRegion"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeExpressConnectRouterRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExpressConnectRouterRegion(request *DescribeExpressConnectRouterRegionRequest) (_result *DescribeExpressConnectRouterRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterRegionResponse{}
	_body, _err := client.DescribeExpressConnectRouterRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExpressConnectRouterRouteEntriesWithOptions(request *DescribeExpressConnectRouterRouteEntriesRequest, runtime *util.RuntimeOptions) (_result *DescribeExpressConnectRouterRouteEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AsPath)) {
		body["AsPath"] = request.AsPath
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Community)) {
		body["Community"] = request.Community
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationCidrBlock)) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.NexthopInstanceId)) {
		body["NexthopInstanceId"] = request.NexthopInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryRegionId)) {
		body["QueryRegionId"] = request.QueryRegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExpressConnectRouterRouteEntries"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExpressConnectRouterRouteEntries(request *DescribeExpressConnectRouterRouteEntriesRequest) (_result *DescribeExpressConnectRouterRouteEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.DescribeExpressConnectRouterRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceGrantedToExpressConnectRouterWithOptions(request *DescribeInstanceGrantedToExpressConnectRouterRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceGrantedToExpressConnectRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceOwnerId)) {
		body["InstanceOwnerId"] = request.InstanceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceRegionId)) {
		body["InstanceRegionId"] = request.InstanceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TagModels)) {
		body["TagModels"] = request.TagModels
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceGrantedToExpressConnectRouter"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceGrantedToExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceGrantedToExpressConnectRouter(request *DescribeInstanceGrantedToExpressConnectRouterRequest) (_result *DescribeInstanceGrantedToExpressConnectRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceGrantedToExpressConnectRouterResponse{}
	_body, _err := client.DescribeInstanceGrantedToExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachExpressConnectRouterChildInstanceWithOptions(request *DetachExpressConnectRouterChildInstanceRequest, runtime *util.RuntimeOptions) (_result *DetachExpressConnectRouterChildInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChildInstanceId)) {
		body["ChildInstanceId"] = request.ChildInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ChildInstanceType)) {
		body["ChildInstanceType"] = request.ChildInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachExpressConnectRouterChildInstance"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachExpressConnectRouterChildInstance(request *DetachExpressConnectRouterChildInstanceRequest) (_result *DetachExpressConnectRouterChildInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.DetachExpressConnectRouterChildInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableExpressConnectRouterRouteEntriesWithOptions(request *DisableExpressConnectRouterRouteEntriesRequest, runtime *util.RuntimeOptions) (_result *DisableExpressConnectRouterRouteEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationCidrBlock)) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.NexthopInstanceId)) {
		body["NexthopInstanceId"] = request.NexthopInstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableExpressConnectRouterRouteEntries"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableExpressConnectRouterRouteEntries(request *DisableExpressConnectRouterRouteEntriesRequest) (_result *DisableExpressConnectRouterRouteEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.DisableExpressConnectRouterRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableExpressConnectRouterRouteEntriesWithOptions(request *EnableExpressConnectRouterRouteEntriesRequest, runtime *util.RuntimeOptions) (_result *EnableExpressConnectRouterRouteEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationCidrBlock)) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.NexthopInstanceId)) {
		body["NexthopInstanceId"] = request.NexthopInstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableExpressConnectRouterRouteEntries"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableExpressConnectRouterRouteEntries(request *EnableExpressConnectRouterRouteEntriesRequest) (_result *EnableExpressConnectRouterRouteEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.EnableExpressConnectRouterRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ForceDeleteExpressConnectRouterWithOptions(request *ForceDeleteExpressConnectRouterRequest, runtime *util.RuntimeOptions) (_result *ForceDeleteExpressConnectRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ForceDeleteExpressConnectRouter"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ForceDeleteExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ForceDeleteExpressConnectRouter(request *ForceDeleteExpressConnectRouterRequest) (_result *ForceDeleteExpressConnectRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ForceDeleteExpressConnectRouterResponse{}
	_body, _err := client.ForceDeleteExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantInstanceToExpressConnectRouterWithOptions(request *GrantInstanceToExpressConnectRouterRequest, runtime *util.RuntimeOptions) (_result *GrantInstanceToExpressConnectRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.EcrOwnerAliUid)) {
		body["EcrOwnerAliUid"] = request.EcrOwnerAliUid
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceRegionId)) {
		body["InstanceRegionId"] = request.InstanceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["InstanceType"] = request.InstanceType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GrantInstanceToExpressConnectRouter"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantInstanceToExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantInstanceToExpressConnectRouter(request *GrantInstanceToExpressConnectRouterRequest) (_result *GrantInstanceToExpressConnectRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantInstanceToExpressConnectRouterResponse{}
	_body, _err := client.GrantInstanceToExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListExpressConnectRouterSupportedRegionWithOptions(request *ListExpressConnectRouterSupportedRegionRequest, runtime *util.RuntimeOptions) (_result *ListExpressConnectRouterSupportedRegionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.NodeType)) {
		body["NodeType"] = request.NodeType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExpressConnectRouterSupportedRegion"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExpressConnectRouterSupportedRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListExpressConnectRouterSupportedRegion(request *ListExpressConnectRouterSupportedRegionRequest) (_result *ListExpressConnectRouterSupportedRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExpressConnectRouterSupportedRegionResponse{}
	_body, _err := client.ListExpressConnectRouterSupportedRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyExpressConnectRouterWithOptions(request *ModifyExpressConnectRouterRequest, runtime *util.RuntimeOptions) (_result *ModifyExpressConnectRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyExpressConnectRouter"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyExpressConnectRouter(request *ModifyExpressConnectRouterRequest) (_result *ModifyExpressConnectRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyExpressConnectRouterResponse{}
	_body, _err := client.ModifyExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyExpressConnectRouterAssociationAllowedPrefixWithOptions(request *ModifyExpressConnectRouterAssociationAllowedPrefixRequest, runtime *util.RuntimeOptions) (_result *ModifyExpressConnectRouterAssociationAllowedPrefixResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowedPrefixes)) {
		body["AllowedPrefixes"] = request.AllowedPrefixes
	}

	if !tea.BoolValue(util.IsUnset(request.AssociationId)) {
		body["AssociationId"] = request.AssociationId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		body["OwnerAccount"] = request.OwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyExpressConnectRouterAssociationAllowedPrefix"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyExpressConnectRouterAssociationAllowedPrefixResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyExpressConnectRouterAssociationAllowedPrefix(request *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) (_result *ModifyExpressConnectRouterAssociationAllowedPrefixResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyExpressConnectRouterAssociationAllowedPrefixResponse{}
	_body, _err := client.ModifyExpressConnectRouterAssociationAllowedPrefixWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyExpressConnectRouterInterRegionTransitModeWithOptions(request *ModifyExpressConnectRouterInterRegionTransitModeRequest, runtime *util.RuntimeOptions) (_result *ModifyExpressConnectRouterInterRegionTransitModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.TransitModeList)) {
		body["TransitModeList"] = request.TransitModeList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyExpressConnectRouterInterRegionTransitMode"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyExpressConnectRouterInterRegionTransitModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyExpressConnectRouterInterRegionTransitMode(request *ModifyExpressConnectRouterInterRegionTransitModeRequest) (_result *ModifyExpressConnectRouterInterRegionTransitModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyExpressConnectRouterInterRegionTransitModeResponse{}
	_body, _err := client.ModifyExpressConnectRouterInterRegionTransitModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevokeInstanceFromExpressConnectRouterWithOptions(request *RevokeInstanceFromExpressConnectRouterRequest, runtime *util.RuntimeOptions) (_result *RevokeInstanceFromExpressConnectRouterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	if !tea.BoolValue(util.IsUnset(request.EcrOwnerAliUid)) {
		body["EcrOwnerAliUid"] = request.EcrOwnerAliUid
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceRegionId)) {
		body["InstanceRegionId"] = request.InstanceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["InstanceType"] = request.InstanceType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeInstanceFromExpressConnectRouter"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeInstanceFromExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokeInstanceFromExpressConnectRouter(request *RevokeInstanceFromExpressConnectRouterRequest) (_result *RevokeInstanceFromExpressConnectRouterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeInstanceFromExpressConnectRouterResponse{}
	_body, _err := client.RevokeInstanceFromExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SynchronizeExpressConnectRouterInterRegionBandwidthWithOptions(request *SynchronizeExpressConnectRouterInterRegionBandwidthRequest, runtime *util.RuntimeOptions) (_result *SynchronizeExpressConnectRouterInterRegionBandwidthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		body["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EcrId)) {
		body["EcrId"] = request.EcrId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SynchronizeExpressConnectRouterInterRegionBandwidth"),
		Version:     tea.String("2023-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SynchronizeExpressConnectRouterInterRegionBandwidthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SynchronizeExpressConnectRouterInterRegionBandwidth(request *SynchronizeExpressConnectRouterInterRegionBandwidthRequest) (_result *SynchronizeExpressConnectRouterInterRegionBandwidthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SynchronizeExpressConnectRouterInterRegionBandwidthResponse{}
	_body, _err := client.SynchronizeExpressConnectRouterInterRegionBandwidthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
