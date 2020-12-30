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

type ListChangedBlocksRequest struct {
	// 标记当前开始读取的位置，置空表示从头开始
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据记录数量
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 待比较的第一个快照名称
	FirstSnapshotId *string `json:"FirstSnapshotId,omitempty" xml:"FirstSnapshotId,omitempty"`
	// 待比较的第二个快照名称
	SecondSnapshotId *string `json:"SecondSnapshotId,omitempty" xml:"SecondSnapshotId,omitempty"`
	// 两个快照比较的起始数据块 index，从零开始
	StartingBlockIndex *int64 `json:"StartingBlockIndex,omitempty" xml:"StartingBlockIndex,omitempty"`
}

func (s ListChangedBlocksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListChangedBlocksRequest) GoString() string {
	return s.String()
}

func (s *ListChangedBlocksRequest) SetNextToken(v string) *ListChangedBlocksRequest {
	s.NextToken = &v
	return s
}

func (s *ListChangedBlocksRequest) SetMaxResults(v int32) *ListChangedBlocksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListChangedBlocksRequest) SetClientToken(v string) *ListChangedBlocksRequest {
	s.ClientToken = &v
	return s
}

func (s *ListChangedBlocksRequest) SetFirstSnapshotId(v string) *ListChangedBlocksRequest {
	s.FirstSnapshotId = &v
	return s
}

func (s *ListChangedBlocksRequest) SetSecondSnapshotId(v string) *ListChangedBlocksRequest {
	s.SecondSnapshotId = &v
	return s
}

func (s *ListChangedBlocksRequest) SetStartingBlockIndex(v int64) *ListChangedBlocksRequest {
	s.StartingBlockIndex = &v
	return s
}

type ListChangedBlocksResponseBody struct {
	// TotalCount本次请求条件下的数据总量，此参数为可选参数，默认可不返回
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 下一页结果的 token，为空时代表无新增页
	NextToken []byte `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// MaxResults本次请求所返回的最大记录条数
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 两个快照差异的数据块列表
	ChangedBlocks []*ListChangedBlocksResponseBodyChangedBlocks `json:"ChangedBlocks,omitempty" xml:"ChangedBlocks,omitempty" type:"Repeated"`
	// 差异数据块 token 过期时间戳
	ExpiryTime *int64 `json:"ExpiryTime,omitempty" xml:"ExpiryTime,omitempty"`
	// 快照大小，单位 GB，最小 1GB
	VolumeSize *int32 `json:"VolumeSize,omitempty" xml:"VolumeSize,omitempty"`
	// 数据块大小，单位 Byte
	BlockSize *int32 `json:"BlockSize,omitempty" xml:"BlockSize,omitempty"`
}

func (s ListChangedBlocksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListChangedBlocksResponseBody) GoString() string {
	return s.String()
}

func (s *ListChangedBlocksResponseBody) SetTotalCount(v int32) *ListChangedBlocksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListChangedBlocksResponseBody) SetRequestId(v string) *ListChangedBlocksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChangedBlocksResponseBody) SetNextToken(v []byte) *ListChangedBlocksResponseBody {
	s.NextToken = v
	return s
}

func (s *ListChangedBlocksResponseBody) SetMaxResults(v int32) *ListChangedBlocksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListChangedBlocksResponseBody) SetChangedBlocks(v []*ListChangedBlocksResponseBodyChangedBlocks) *ListChangedBlocksResponseBody {
	s.ChangedBlocks = v
	return s
}

func (s *ListChangedBlocksResponseBody) SetExpiryTime(v int64) *ListChangedBlocksResponseBody {
	s.ExpiryTime = &v
	return s
}

func (s *ListChangedBlocksResponseBody) SetVolumeSize(v int32) *ListChangedBlocksResponseBody {
	s.VolumeSize = &v
	return s
}

func (s *ListChangedBlocksResponseBody) SetBlockSize(v int32) *ListChangedBlocksResponseBody {
	s.BlockSize = &v
	return s
}

type ListChangedBlocksResponseBodyChangedBlocks struct {
	// 数据块的索引值，从零开始
	BlockIndex *int64 `json:"BlockIndex,omitempty" xml:"BlockIndex,omitempty"`
	// 数据块的 Token，用于后续的数据读取
	BlockToken []byte `json:"BlockToken,omitempty" xml:"BlockToken,omitempty"`
}

func (s ListChangedBlocksResponseBodyChangedBlocks) String() string {
	return tea.Prettify(s)
}

func (s ListChangedBlocksResponseBodyChangedBlocks) GoString() string {
	return s.String()
}

func (s *ListChangedBlocksResponseBodyChangedBlocks) SetBlockIndex(v int64) *ListChangedBlocksResponseBodyChangedBlocks {
	s.BlockIndex = &v
	return s
}

func (s *ListChangedBlocksResponseBodyChangedBlocks) SetBlockToken(v []byte) *ListChangedBlocksResponseBodyChangedBlocks {
	s.BlockToken = v
	return s
}

type ListChangedBlocksResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListChangedBlocksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListChangedBlocksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListChangedBlocksResponse) GoString() string {
	return s.String()
}

func (s *ListChangedBlocksResponse) SetHeaders(v map[string]*string) *ListChangedBlocksResponse {
	s.Headers = v
	return s
}

func (s *ListChangedBlocksResponse) SetBody(v *ListChangedBlocksResponseBody) *ListChangedBlocksResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("snapshot"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ListChangedBlocks(request *ListChangedBlocksRequest) (_result *ListChangedBlocksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListChangedBlocksResponse{}
	_body, _err := client.ListChangedBlocksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListChangedBlocksWithOptions(request *ListChangedBlocksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListChangedBlocksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.FirstSnapshotId)) {
		query["FirstSnapshotId"] = request.FirstSnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.SecondSnapshotId)) {
		query["SecondSnapshotId"] = request.SecondSnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.StartingBlockIndex)) {
		query["StartingBlockIndex"] = request.StartingBlockIndex
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListChangedBlocksResponse{}
	_body, _err := client.DoROARequest(tea.String("ListChangedBlocks"), tea.String("2020-11-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/snapshots/changedblocks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
