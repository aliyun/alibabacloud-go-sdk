// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ActualDeductResourceCmd struct {
	Cost         *int64  `json:"cost,omitempty" xml:"cost,omitempty"`
	ExtraInfo    *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	IdempotentId *string `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
	TaskId       *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ActualDeductResourceCmd) String() string {
	return tea.Prettify(s)
}

func (s ActualDeductResourceCmd) GoString() string {
	return s.String()
}

func (s *ActualDeductResourceCmd) SetCost(v int64) *ActualDeductResourceCmd {
	s.Cost = &v
	return s
}

func (s *ActualDeductResourceCmd) SetExtraInfo(v string) *ActualDeductResourceCmd {
	s.ExtraInfo = &v
	return s
}

func (s *ActualDeductResourceCmd) SetIdempotentId(v string) *ActualDeductResourceCmd {
	s.IdempotentId = &v
	return s
}

func (s *ActualDeductResourceCmd) SetTaskId(v string) *ActualDeductResourceCmd {
	s.TaskId = &v
	return s
}

type ActualDeductResourceResult struct {
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Errorcode    *string `json:"errorcode,omitempty" xml:"errorcode,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ActualDeductResourceResult) String() string {
	return tea.Prettify(s)
}

func (s ActualDeductResourceResult) GoString() string {
	return s.String()
}

func (s *ActualDeductResourceResult) SetErrorMessage(v string) *ActualDeductResourceResult {
	s.ErrorMessage = &v
	return s
}

func (s *ActualDeductResourceResult) SetErrorcode(v string) *ActualDeductResourceResult {
	s.Errorcode = &v
	return s
}

func (s *ActualDeductResourceResult) SetRequestId(v string) *ActualDeductResourceResult {
	s.RequestId = &v
	return s
}

func (s *ActualDeductResourceResult) SetSuccess(v bool) *ActualDeductResourceResult {
	s.Success = &v
	return s
}

type DirectDeductResourceCmd struct {
	AccountId    *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Cost         *int64  `json:"cost,omitempty" xml:"cost,omitempty"`
	ExtraInfo    *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	IdempotentId *string `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
	ResourceType *int64  `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	SubAccountId *string `json:"subAccountId,omitempty" xml:"subAccountId,omitempty"`
	Token        *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s DirectDeductResourceCmd) String() string {
	return tea.Prettify(s)
}

func (s DirectDeductResourceCmd) GoString() string {
	return s.String()
}

func (s *DirectDeductResourceCmd) SetAccountId(v string) *DirectDeductResourceCmd {
	s.AccountId = &v
	return s
}

func (s *DirectDeductResourceCmd) SetCost(v int64) *DirectDeductResourceCmd {
	s.Cost = &v
	return s
}

func (s *DirectDeductResourceCmd) SetExtraInfo(v string) *DirectDeductResourceCmd {
	s.ExtraInfo = &v
	return s
}

func (s *DirectDeductResourceCmd) SetIdempotentId(v string) *DirectDeductResourceCmd {
	s.IdempotentId = &v
	return s
}

func (s *DirectDeductResourceCmd) SetResourceType(v int64) *DirectDeductResourceCmd {
	s.ResourceType = &v
	return s
}

func (s *DirectDeductResourceCmd) SetSubAccountId(v string) *DirectDeductResourceCmd {
	s.SubAccountId = &v
	return s
}

func (s *DirectDeductResourceCmd) SetToken(v string) *DirectDeductResourceCmd {
	s.Token = &v
	return s
}

type DirectDeductResourceResult struct {
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Errorcode    *string `json:"errorcode,omitempty" xml:"errorcode,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DirectDeductResourceResult) String() string {
	return tea.Prettify(s)
}

func (s DirectDeductResourceResult) GoString() string {
	return s.String()
}

func (s *DirectDeductResourceResult) SetErrorMessage(v string) *DirectDeductResourceResult {
	s.ErrorMessage = &v
	return s
}

func (s *DirectDeductResourceResult) SetErrorcode(v string) *DirectDeductResourceResult {
	s.Errorcode = &v
	return s
}

func (s *DirectDeductResourceResult) SetRequestId(v string) *DirectDeductResourceResult {
	s.RequestId = &v
	return s
}

func (s *DirectDeductResourceResult) SetSuccess(v bool) *DirectDeductResourceResult {
	s.Success = &v
	return s
}

type ExpectDeductResourceCmd struct {
	AccountId    *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Cost         *int64  `json:"cost,omitempty" xml:"cost,omitempty"`
	ExtraInfo    *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	IdempotentId *string `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
	ResourceType *int64  `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	SubAccountId *string `json:"subAccountId,omitempty" xml:"subAccountId,omitempty"`
	Token        *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s ExpectDeductResourceCmd) String() string {
	return tea.Prettify(s)
}

func (s ExpectDeductResourceCmd) GoString() string {
	return s.String()
}

func (s *ExpectDeductResourceCmd) SetAccountId(v string) *ExpectDeductResourceCmd {
	s.AccountId = &v
	return s
}

func (s *ExpectDeductResourceCmd) SetCost(v int64) *ExpectDeductResourceCmd {
	s.Cost = &v
	return s
}

func (s *ExpectDeductResourceCmd) SetExtraInfo(v string) *ExpectDeductResourceCmd {
	s.ExtraInfo = &v
	return s
}

func (s *ExpectDeductResourceCmd) SetIdempotentId(v string) *ExpectDeductResourceCmd {
	s.IdempotentId = &v
	return s
}

func (s *ExpectDeductResourceCmd) SetResourceType(v int64) *ExpectDeductResourceCmd {
	s.ResourceType = &v
	return s
}

func (s *ExpectDeductResourceCmd) SetSubAccountId(v string) *ExpectDeductResourceCmd {
	s.SubAccountId = &v
	return s
}

func (s *ExpectDeductResourceCmd) SetToken(v string) *ExpectDeductResourceCmd {
	s.Token = &v
	return s
}

type ExpectDeductResourceResult struct {
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Errorcode    *string `json:"errorcode,omitempty" xml:"errorcode,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success      *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TaskId       *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ExpectDeductResourceResult) String() string {
	return tea.Prettify(s)
}

func (s ExpectDeductResourceResult) GoString() string {
	return s.String()
}

func (s *ExpectDeductResourceResult) SetErrorMessage(v string) *ExpectDeductResourceResult {
	s.ErrorMessage = &v
	return s
}

func (s *ExpectDeductResourceResult) SetErrorcode(v string) *ExpectDeductResourceResult {
	s.Errorcode = &v
	return s
}

func (s *ExpectDeductResourceResult) SetRequestId(v string) *ExpectDeductResourceResult {
	s.RequestId = &v
	return s
}

func (s *ExpectDeductResourceResult) SetSuccess(v bool) *ExpectDeductResourceResult {
	s.Success = &v
	return s
}

func (s *ExpectDeductResourceResult) SetTaskId(v string) *ExpectDeductResourceResult {
	s.TaskId = &v
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("intelligentcreation"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
