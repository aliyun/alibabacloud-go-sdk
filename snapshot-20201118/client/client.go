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

type GetSnapshotBlockRequest struct {
	// 幂等参数
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 待读取的数据块索引，从零开始。从 ListChangedBlocks 或者 ListSnapshotBlocks 返回
	BlockIndex *string `json:"BlockIndex,omitempty" xml:"BlockIndex,omitempty"`
	// 待读取的数据块Token，从零开始。从 ListChangedBlocks 或者 ListSnapshotBlocks 返回
	BlockToken *string `json:"BlockToken,omitempty" xml:"BlockToken,omitempty"`
	// 待读取数据的快照名称
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

func (s *GetSnapshotBlockRequest) SetBlockIndex(v string) *GetSnapshotBlockRequest {
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

type ListChangedBlocksRequest struct {
	NextToken          *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults         *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	FirstSnapshotId    *string `json:"FirstSnapshotId,omitempty" xml:"FirstSnapshotId,omitempty"`
	SecondSnapshotId   *string `json:"SecondSnapshotId,omitempty" xml:"SecondSnapshotId,omitempty"`
	StartingBlockIndex *int64  `json:"StartingBlockIndex,omitempty" xml:"StartingBlockIndex,omitempty"`
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
	NextToken     *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ExpiryTime    *int64                                        `json:"ExpiryTime,omitempty" xml:"ExpiryTime,omitempty"`
	BlockSize     *int32                                        `json:"BlockSize,omitempty" xml:"BlockSize,omitempty"`
	VolumeSize    *int32                                        `json:"VolumeSize,omitempty" xml:"VolumeSize,omitempty"`
	ChangedBlocks []*ListChangedBlocksResponseBodyChangedBlocks `json:"ChangedBlocks,omitempty" xml:"ChangedBlocks,omitempty" type:"Repeated"`
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

func (s *ListChangedBlocksResponseBody) SetExpiryTime(v int64) *ListChangedBlocksResponseBody {
	s.ExpiryTime = &v
	return s
}

func (s *ListChangedBlocksResponseBody) SetBlockSize(v int32) *ListChangedBlocksResponseBody {
	s.BlockSize = &v
	return s
}

func (s *ListChangedBlocksResponseBody) SetVolumeSize(v int32) *ListChangedBlocksResponseBody {
	s.VolumeSize = &v
	return s
}

func (s *ListChangedBlocksResponseBody) SetChangedBlocks(v []*ListChangedBlocksResponseBodyChangedBlocks) *ListChangedBlocksResponseBody {
	s.ChangedBlocks = v
	return s
}

type ListChangedBlocksResponseBodyChangedBlocks struct {
	FirstBlockToken  *string `json:"FirstBlockToken,omitempty" xml:"FirstBlockToken,omitempty"`
	BlockIndex       *int64  `json:"BlockIndex,omitempty" xml:"BlockIndex,omitempty"`
	SecondBlockToken *string `json:"SecondBlockToken,omitempty" xml:"SecondBlockToken,omitempty"`
}

func (s ListChangedBlocksResponseBodyChangedBlocks) String() string {
	return tea.Prettify(s)
}

func (s ListChangedBlocksResponseBodyChangedBlocks) GoString() string {
	return s.String()
}

func (s *ListChangedBlocksResponseBodyChangedBlocks) SetFirstBlockToken(v string) *ListChangedBlocksResponseBodyChangedBlocks {
	s.FirstBlockToken = &v
	return s
}

func (s *ListChangedBlocksResponseBodyChangedBlocks) SetBlockIndex(v int64) *ListChangedBlocksResponseBodyChangedBlocks {
	s.BlockIndex = &v
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

type ListSnapshotBlocksRequest struct {
	NextToken          *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults         *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	SnapshotId         *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	StartingBlockIndex *int64  `json:"StartingBlockIndex,omitempty" xml:"StartingBlockIndex,omitempty"`
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

func (s *ListSnapshotBlocksRequest) SetMaxResults(v int32) *ListSnapshotBlocksRequest {
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
	NextToken  *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ExpiryTime *int64                                  `json:"ExpiryTime,omitempty" xml:"ExpiryTime,omitempty"`
	BlockSize  *int32                                  `json:"BlockSize,omitempty" xml:"BlockSize,omitempty"`
	VolumeSize *int32                                  `json:"VolumeSize,omitempty" xml:"VolumeSize,omitempty"`
	Blocks     []*ListSnapshotBlocksResponseBodyBlocks `json:"Blocks,omitempty" xml:"Blocks,omitempty" type:"Repeated"`
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

func (s *ListSnapshotBlocksResponseBody) SetExpiryTime(v int64) *ListSnapshotBlocksResponseBody {
	s.ExpiryTime = &v
	return s
}

func (s *ListSnapshotBlocksResponseBody) SetBlockSize(v int32) *ListSnapshotBlocksResponseBody {
	s.BlockSize = &v
	return s
}

func (s *ListSnapshotBlocksResponseBody) SetVolumeSize(v int32) *ListSnapshotBlocksResponseBody {
	s.VolumeSize = &v
	return s
}

func (s *ListSnapshotBlocksResponseBody) SetBlocks(v []*ListSnapshotBlocksResponseBodyBlocks) *ListSnapshotBlocksResponseBody {
	s.Blocks = v
	return s
}

type ListSnapshotBlocksResponseBodyBlocks struct {
	BlockIndex *int64  `json:"BlockIndex,omitempty" xml:"BlockIndex,omitempty"`
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
	tmp, _err := client.DoROARequest(tea.String("GetSnapshotBlock"), tea.String("2020-11-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/snapshots/block"), tea.String("binary"), req, runtime)
	if _err != nil {
		return _result, _err
	}

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
	_body, _err := client.DoROARequest(tea.String("ListSnapshotBlocks"), tea.String("2020-11-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/snapshots/listblocks"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
