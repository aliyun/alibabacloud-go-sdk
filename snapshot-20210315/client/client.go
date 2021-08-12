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
	"io"
)

type GetSnapshotInfoRequest struct {
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 待读取数据的快照名称
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// 是否返回详细信息，详细信息需要更多查询时间，默认为 False
	ShowDetail *bool `json:"ShowDetail,omitempty" xml:"ShowDetail,omitempty"`
}

func (s GetSnapshotInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSnapshotInfoRequest) SetClientToken(v string) *GetSnapshotInfoRequest {
	s.ClientToken = &v
	return s
}

func (s *GetSnapshotInfoRequest) SetSnapshotId(v string) *GetSnapshotInfoRequest {
	s.SnapshotId = &v
	return s
}

func (s *GetSnapshotInfoRequest) SetShowDetail(v bool) *GetSnapshotInfoRequest {
	s.ShowDetail = &v
	return s
}

type GetSnapshotInfoResponseBody struct {
	// 快照大小，单位 GB，最小 1GB
	VolumeSize *int64 `json:"VolumeSize,omitempty" xml:"VolumeSize,omitempty"`
	// 快照数据快大小，单位 Bytes
	BlockSize *int64 `json:"BlockSize,omitempty" xml:"BlockSize,omitempty"`
	// 快照数据块总数量，包含空数据块
	BlockCount *int64 `json:"BlockCount,omitempty" xml:"BlockCount,omitempty"`
	// 快照中非空数据块总数量，仅在 ShowDetail 为 True 时返回
	ValidBlockCount *int64 `json:"ValidBlockCount,omitempty" xml:"ValidBlockCount,omitempty"`
	// 快照状态。    "SNAPSHOT_INVALID",     "SNAPSHOT_PENDING",     "SNAPSHOT_COMMITTED",     "SNAPSHOT_DELETED"
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 快照创建时间戳
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 快照是否为加密快照
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// 快照格式。"FORMAT_VHD", "FORMAT_RAW", "FORMAT_INVALID"
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
}

func (s GetSnapshotInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSnapshotInfoResponseBody) SetVolumeSize(v int64) *GetSnapshotInfoResponseBody {
	s.VolumeSize = &v
	return s
}

func (s *GetSnapshotInfoResponseBody) SetBlockSize(v int64) *GetSnapshotInfoResponseBody {
	s.BlockSize = &v
	return s
}

func (s *GetSnapshotInfoResponseBody) SetBlockCount(v int64) *GetSnapshotInfoResponseBody {
	s.BlockCount = &v
	return s
}

func (s *GetSnapshotInfoResponseBody) SetValidBlockCount(v int64) *GetSnapshotInfoResponseBody {
	s.ValidBlockCount = &v
	return s
}

func (s *GetSnapshotInfoResponseBody) SetStatus(v string) *GetSnapshotInfoResponseBody {
	s.Status = &v
	return s
}

func (s *GetSnapshotInfoResponseBody) SetCreateTime(v int64) *GetSnapshotInfoResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetSnapshotInfoResponseBody) SetEncrypted(v bool) *GetSnapshotInfoResponseBody {
	s.Encrypted = &v
	return s
}

func (s *GetSnapshotInfoResponseBody) SetFormat(v string) *GetSnapshotInfoResponseBody {
	s.Format = &v
	return s
}

type GetSnapshotInfoResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSnapshotInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSnapshotInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotInfoResponse) GoString() string {
	return s.String()
}

func (s *GetSnapshotInfoResponse) SetHeaders(v map[string]*string) *GetSnapshotInfoResponse {
	s.Headers = v
	return s
}

func (s *GetSnapshotInfoResponse) SetBody(v *GetSnapshotInfoResponseBody) *GetSnapshotInfoResponse {
	s.Body = v
	return s
}

type GetSnapshotBlockRequest struct {
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 待读取的数据块索引，从零开始。从 ListChangedBlocks 或者 ListSnapshotBlocks 返回。最大为快照数据块数量减一
	BlockIndex *int64 `json:"BlockIndex,omitempty" xml:"BlockIndex,omitempty"`
	// 待读取的数据块Token，从 ListChangedBlocks 或者 ListSnapshotBlocks 返回，最大长度 256 字节
	BlockToken *string `json:"BlockToken,omitempty" xml:"BlockToken,omitempty"`
	// 待读取数据的快照名称，最大 256 字节
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s GetSnapshotBlockRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotBlockRequest) GoString() string {
	return s.String()
}

func (s *GetSnapshotBlockRequest) SetClientToken(v string) *GetSnapshotBlockRequest {
	s.ClientToken = &v
	return s
}

func (s *GetSnapshotBlockRequest) SetBlockIndex(v int64) *GetSnapshotBlockRequest {
	s.BlockIndex = &v
	return s
}

func (s *GetSnapshotBlockRequest) SetBlockToken(v string) *GetSnapshotBlockRequest {
	s.BlockToken = &v
	return s
}

func (s *GetSnapshotBlockRequest) SetSnapshotId(v string) *GetSnapshotBlockRequest {
	s.SnapshotId = &v
	return s
}

type GetSnapshotBlockResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    io.Reader          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSnapshotBlockResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSnapshotBlockResponse) GoString() string {
	return s.String()
}

func (s *GetSnapshotBlockResponse) SetHeaders(v map[string]*string) *GetSnapshotBlockResponse {
	s.Headers = v
	return s
}

func (s *GetSnapshotBlockResponse) SetBody(v io.Reader) *GetSnapshotBlockResponse {
	s.Body = v
	return s
}

type ListSnapshotBlocksRequest struct {
	// 标记当前开始读取的位置，置空表示从 StartingBlockIndex 指定 Block 开始，最大长度 256 字节
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据记录数量，最小 100， 最大 10000， 默认为 100
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 待列出数据块的快照名称，最大 256 字节
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// 列出快照中数据块起始索引值，从零开始，最大值为快照数据块总数减一。NextToken 不为空时忽略此参数
	StartingBlockIndex *int64 `json:"StartingBlockIndex,omitempty" xml:"StartingBlockIndex,omitempty"`
}

func (s ListSnapshotBlocksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotBlocksRequest) GoString() string {
	return s.String()
}

func (s *ListSnapshotBlocksRequest) SetNextToken(v string) *ListSnapshotBlocksRequest {
	s.NextToken = &v
	return s
}

func (s *ListSnapshotBlocksRequest) SetMaxResults(v int64) *ListSnapshotBlocksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSnapshotBlocksRequest) SetClientToken(v string) *ListSnapshotBlocksRequest {
	s.ClientToken = &v
	return s
}

func (s *ListSnapshotBlocksRequest) SetSnapshotId(v string) *ListSnapshotBlocksRequest {
	s.SnapshotId = &v
	return s
}

func (s *ListSnapshotBlocksRequest) SetStartingBlockIndex(v int64) *ListSnapshotBlocksRequest {
	s.StartingBlockIndex = &v
	return s
}

type ListSnapshotBlocksResponseBody struct {
	// 返回下一个有效数据块 token，为空时代表无其他有效数据块，最大  256 字节
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 快照有效数据块信息列表，不包含空数据块
	Blocks []*ListSnapshotBlocksResponseBodyBlocks `json:"Blocks,omitempty" xml:"Blocks,omitempty" type:"Repeated"`
	// BlockToken 过期 UTC 时间戳, 单位微秒
	ExpiryTime *int64 `json:"ExpiryTime,omitempty" xml:"ExpiryTime,omitempty"`
	// 快照大小，单位 GB，最小 1GB
	VolumeSize *int64 `json:"VolumeSize,omitempty" xml:"VolumeSize,omitempty"`
	// 数据块大小，单位 Byte
	BlockSize *int64 `json:"BlockSize,omitempty" xml:"BlockSize,omitempty"`
	// 本次查询中快照有效数据块数量
	BlockCount *int64 `json:"BlockCount,omitempty" xml:"BlockCount,omitempty"`
	// 快照有效数据块总数量
	TotalBlockCount *int64 `json:"TotalBlockCount,omitempty" xml:"TotalBlockCount,omitempty"`
}

func (s ListSnapshotBlocksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotBlocksResponseBody) GoString() string {
	return s.String()
}

func (s *ListSnapshotBlocksResponseBody) SetNextToken(v string) *ListSnapshotBlocksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSnapshotBlocksResponseBody) SetBlocks(v []*ListSnapshotBlocksResponseBodyBlocks) *ListSnapshotBlocksResponseBody {
	s.Blocks = v
	return s
}

func (s *ListSnapshotBlocksResponseBody) SetExpiryTime(v int64) *ListSnapshotBlocksResponseBody {
	s.ExpiryTime = &v
	return s
}

func (s *ListSnapshotBlocksResponseBody) SetVolumeSize(v int64) *ListSnapshotBlocksResponseBody {
	s.VolumeSize = &v
	return s
}

func (s *ListSnapshotBlocksResponseBody) SetBlockSize(v int64) *ListSnapshotBlocksResponseBody {
	s.BlockSize = &v
	return s
}

func (s *ListSnapshotBlocksResponseBody) SetBlockCount(v int64) *ListSnapshotBlocksResponseBody {
	s.BlockCount = &v
	return s
}

func (s *ListSnapshotBlocksResponseBody) SetTotalBlockCount(v int64) *ListSnapshotBlocksResponseBody {
	s.TotalBlockCount = &v
	return s
}

type ListSnapshotBlocksResponseBodyBlocks struct {
	// 数据块的索引值，从零开始
	BlockIndex *int64 `json:"BlockIndex,omitempty" xml:"BlockIndex,omitempty"`
	// 数据块的 Token，用于后续的数据读取。最大长度 256 字节
	BlockToken *string `json:"BlockToken,omitempty" xml:"BlockToken,omitempty"`
}

func (s ListSnapshotBlocksResponseBodyBlocks) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotBlocksResponseBodyBlocks) GoString() string {
	return s.String()
}

func (s *ListSnapshotBlocksResponseBodyBlocks) SetBlockIndex(v int64) *ListSnapshotBlocksResponseBodyBlocks {
	s.BlockIndex = &v
	return s
}

func (s *ListSnapshotBlocksResponseBodyBlocks) SetBlockToken(v string) *ListSnapshotBlocksResponseBodyBlocks {
	s.BlockToken = &v
	return s
}

type ListSnapshotBlocksResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSnapshotBlocksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSnapshotBlocksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotBlocksResponse) GoString() string {
	return s.String()
}

func (s *ListSnapshotBlocksResponse) SetHeaders(v map[string]*string) *ListSnapshotBlocksResponse {
	s.Headers = v
	return s
}

func (s *ListSnapshotBlocksResponse) SetBody(v *ListSnapshotBlocksResponseBody) *ListSnapshotBlocksResponse {
	s.Body = v
	return s
}

type ListChangedBlocksRequest struct {
	// 标记开始比较差异的位置，为空时代表从 StartingBlockIndex 指定数据块开始比较， 最大长度 256 字节
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据记录数量，最小 100，最大 10000
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 待比较的第一个快照名称，最大 256 字节
	FirstSnapshotId *string `json:"FirstSnapshotId,omitempty" xml:"FirstSnapshotId,omitempty"`
	// 待比较的第二个快照名称，最大 256 字节
	SecondSnapshotId *string `json:"SecondSnapshotId,omitempty" xml:"SecondSnapshotId,omitempty"`
	// 两个快照比较的起始数据块 index，从零开始，最大值为第二快照的数据块总数量减一。当 NextToken 不为空时忽略该参数
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

func (s *ListChangedBlocksRequest) SetMaxResults(v int64) *ListChangedBlocksRequest {
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
	// 下一个待比较的数据块 BlockToken，为空时代表无其他待比较数据块，最大长度 256 字节
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 两个快照差异的数据块列表
	ChangedBlocks []*ListChangedBlocksResponseBodyChangedBlocks `json:"ChangedBlocks,omitempty" xml:"ChangedBlocks,omitempty" type:"Repeated"`
	// 差异数据块 token 过期 UTC 时间戳， 单位微秒
	ExpiryTime *int64 `json:"ExpiryTime,omitempty" xml:"ExpiryTime,omitempty"`
	// 第二个快照大小，单位 GB，最小 1GB
	VolumeSize *int64 `json:"VolumeSize,omitempty" xml:"VolumeSize,omitempty"`
	// 数据块大小，单位 Byte
	BlockSize *int64 `json:"BlockSize,omitempty" xml:"BlockSize,omitempty"`
	// 本次查询中变化数据块数量
	BlockCount *int64 `json:"BlockCount,omitempty" xml:"BlockCount,omitempty"`
	// 第二个快照相对于第一个快照差异数据快总数
	TotalBlockCount *int64 `json:"TotalBlockCount,omitempty" xml:"TotalBlockCount,omitempty"`
}

func (s ListChangedBlocksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListChangedBlocksResponseBody) GoString() string {
	return s.String()
}

func (s *ListChangedBlocksResponseBody) SetNextToken(v string) *ListChangedBlocksResponseBody {
	s.NextToken = &v
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

func (s *ListChangedBlocksResponseBody) SetVolumeSize(v int64) *ListChangedBlocksResponseBody {
	s.VolumeSize = &v
	return s
}

func (s *ListChangedBlocksResponseBody) SetBlockSize(v int64) *ListChangedBlocksResponseBody {
	s.BlockSize = &v
	return s
}

func (s *ListChangedBlocksResponseBody) SetBlockCount(v int64) *ListChangedBlocksResponseBody {
	s.BlockCount = &v
	return s
}

func (s *ListChangedBlocksResponseBody) SetTotalBlockCount(v int64) *ListChangedBlocksResponseBody {
	s.TotalBlockCount = &v
	return s
}

type ListChangedBlocksResponseBodyChangedBlocks struct {
	// 数据块的索引值，从零开始
	BlockIndex *int64 `json:"BlockIndex,omitempty" xml:"BlockIndex,omitempty"`
	// FirstSnapshotId 中数据块的 Token，用于后续的数据读取，第一个快照未变化时可省略。最大长度 256 字节
	FirstBlockToken *string `json:"FirstBlockToken,omitempty" xml:"FirstBlockToken,omitempty"`
	// SecondBlockToken指定的快照中相对于 FirstBlockIndex 变化的数据块 Token，用于后续读取数据。最大长度 256 字节
	SecondBlockToken *string `json:"SecondBlockToken,omitempty" xml:"SecondBlockToken,omitempty"`
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

func (s *ListChangedBlocksResponseBodyChangedBlocks) SetFirstBlockToken(v string) *ListChangedBlocksResponseBodyChangedBlocks {
	s.FirstBlockToken = &v
	return s
}

func (s *ListChangedBlocksResponseBodyChangedBlocks) SetSecondBlockToken(v string) *ListChangedBlocksResponseBodyChangedBlocks {
	s.SecondBlockToken = &v
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
	client.EndpointRule = tea.String("regional")
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

func (client *Client) GetSnapshotInfo(request *GetSnapshotInfoRequest) (_result *GetSnapshotInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSnapshotInfoResponse{}
	_body, _err := client.GetSnapshotInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSnapshotInfoWithOptions(request *GetSnapshotInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSnapshotInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowDetail)) {
		query["ShowDetail"] = request.ShowDetail
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetSnapshotInfoResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSnapshotInfo"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/snapshots/info"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSnapshotBlock(request *GetSnapshotBlockRequest) (_result *GetSnapshotBlockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSnapshotBlockResponse{}
	_body, _err := client.GetSnapshotBlockWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSnapshotBlockWithOptions(request *GetSnapshotBlockRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSnapshotBlockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.BlockIndex)) {
		query["BlockIndex"] = request.BlockIndex
	}

	if !tea.BoolValue(util.IsUnset(request.BlockToken)) {
		query["BlockToken"] = request.BlockToken
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	res := &GetSnapshotBlockResponse{}
	doROARequestTmp, err := client.DoROARequest(tea.String("GetSnapshotBlock"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/snapshots/block"), tea.String("binary"), req, runtime)
	if err != nil {
		_err = err
		return
	}
	tmp := util.AssertAsMap(doROARequestTmp)
	if !tea.BoolValue(util.IsUnset(tmp["body"])) {
		respBody := util.AssertAsReadable(tmp["body"])
		res.Body = respBody
	}

	if !tea.BoolValue(util.IsUnset(tmp["headers"])) {
		respHeaders := util.AssertAsMap(tmp["headers"])
		res.Headers = util.StringifyMapValue(respHeaders)
	}

	_result = res
	return _result, _err
}

func (client *Client) ListSnapshotBlocks(request *ListSnapshotBlocksRequest) (_result *ListSnapshotBlocksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSnapshotBlocksResponse{}
	_body, _err := client.ListSnapshotBlocksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSnapshotBlocksWithOptions(request *ListSnapshotBlocksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSnapshotBlocksResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.StartingBlockIndex)) {
		query["StartingBlockIndex"] = request.StartingBlockIndex
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListSnapshotBlocksResponse{}
	_body, _err := client.DoROARequest(tea.String("ListSnapshotBlocks"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/snapshots/listblocks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	_body, _err := client.DoROARequest(tea.String("ListChangedBlocks"), tea.String("2021-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/snapshots/changedblocks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
