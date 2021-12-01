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

type CreateDiskReplicaPairRequest struct {
	Bandwidth           *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ChargeType          *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClientToken         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationDiskId   *string `json:"DestinationDiskId,omitempty" xml:"DestinationDiskId,omitempty"`
	DestinationRegionId *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	DestinationZoneId   *string `json:"DestinationZoneId,omitempty" xml:"DestinationZoneId,omitempty"`
	DiskId              *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	PairName            *string `json:"PairName,omitempty" xml:"PairName,omitempty"`
	Period              *int64  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit          *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	SourceRegionId      *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	SourceZoneId        *string `json:"SourceZoneId,omitempty" xml:"SourceZoneId,omitempty"`
}

func (s CreateDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaPairRequest) SetBandwidth(v int64) *CreateDiskReplicaPairRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetChargeType(v string) *CreateDiskReplicaPairRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetClientToken(v string) *CreateDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetDescription(v string) *CreateDiskReplicaPairRequest {
	s.Description = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetDestinationDiskId(v string) *CreateDiskReplicaPairRequest {
	s.DestinationDiskId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetDestinationRegionId(v string) *CreateDiskReplicaPairRequest {
	s.DestinationRegionId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetDestinationZoneId(v string) *CreateDiskReplicaPairRequest {
	s.DestinationZoneId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetDiskId(v string) *CreateDiskReplicaPairRequest {
	s.DiskId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetPairName(v string) *CreateDiskReplicaPairRequest {
	s.PairName = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetPeriod(v int64) *CreateDiskReplicaPairRequest {
	s.Period = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetPeriodUnit(v string) *CreateDiskReplicaPairRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetSourceRegionId(v string) *CreateDiskReplicaPairRequest {
	s.SourceRegionId = &v
	return s
}

func (s *CreateDiskReplicaPairRequest) SetSourceZoneId(v string) *CreateDiskReplicaPairRequest {
	s.SourceZoneId = &v
	return s
}

type CreateDiskReplicaPairResponseBody struct {
	OrderId       *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaPairResponseBody) SetOrderId(v string) *CreateDiskReplicaPairResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDiskReplicaPairResponseBody) SetReplicaPairId(v string) *CreateDiskReplicaPairResponseBody {
	s.ReplicaPairId = &v
	return s
}

func (s *CreateDiskReplicaPairResponseBody) SetRequestId(v string) *CreateDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type CreateDiskReplicaPairResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *CreateDiskReplicaPairResponse) SetHeaders(v map[string]*string) *CreateDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *CreateDiskReplicaPairResponse) SetBody(v *CreateDiskReplicaPairResponseBody) *CreateDiskReplicaPairResponse {
	s.Body = v
	return s
}

type DeleteDiskReplicaPairRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s DeleteDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *DeleteDiskReplicaPairRequest) SetClientToken(v string) *DeleteDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDiskReplicaPairRequest) SetRegionId(v string) *DeleteDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDiskReplicaPairRequest) SetReplicaPairId(v string) *DeleteDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

type DeleteDiskReplicaPairResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDiskReplicaPairResponseBody) SetRequestId(v string) *DeleteDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDiskReplicaPairResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *DeleteDiskReplicaPairResponse) SetHeaders(v map[string]*string) *DeleteDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *DeleteDiskReplicaPairResponse) SetBody(v *DeleteDiskReplicaPairResponseBody) *DeleteDiskReplicaPairResponse {
	s.Body = v
	return s
}

type DescribeDiskReplicaPairsRequest struct {
	MaxResults *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PairIds    *string `json:"PairIds,omitempty" xml:"PairIds,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDiskReplicaPairsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaPairsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairsRequest) SetMaxResults(v int64) *DescribeDiskReplicaPairsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetNextToken(v string) *DescribeDiskReplicaPairsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetPairIds(v string) *DescribeDiskReplicaPairsRequest {
	s.PairIds = &v
	return s
}

func (s *DescribeDiskReplicaPairsRequest) SetRegionId(v string) *DescribeDiskReplicaPairsRequest {
	s.RegionId = &v
	return s
}

type DescribeDiskReplicaPairsResponseBody struct {
	NextToken    *string                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ReplicaPairs []*DescribeDiskReplicaPairsResponseBodyReplicaPairs `json:"ReplicaPairs,omitempty" xml:"ReplicaPairs,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDiskReplicaPairsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaPairsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairsResponseBody) SetNextToken(v string) *DescribeDiskReplicaPairsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBody) SetReplicaPairs(v []*DescribeDiskReplicaPairsResponseBodyReplicaPairs) *DescribeDiskReplicaPairsResponseBody {
	s.ReplicaPairs = v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBody) SetRequestId(v string) *DescribeDiskReplicaPairsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDiskReplicaPairsResponseBodyReplicaPairs struct {
	AsyncCycle        *int64  `json:"AsyncCycle,omitempty" xml:"AsyncCycle,omitempty"`
	Bandwidth         *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationDiskId *string `json:"DestinationDiskId,omitempty" xml:"DestinationDiskId,omitempty"`
	DestinationRegion *string `json:"DestinationRegion,omitempty" xml:"DestinationRegion,omitempty"`
	LastRPO           *int64  `json:"LastRPO,omitempty" xml:"LastRPO,omitempty"`
	PairName          *string `json:"PairName,omitempty" xml:"PairName,omitempty"`
	ReplicaPairId     *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
	SourceDiskId      *string `json:"SourceDiskId,omitempty" xml:"SourceDiskId,omitempty"`
	SourceRegion      *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusMessage     *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s DescribeDiskReplicaPairsResponseBodyReplicaPairs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaPairsResponseBodyReplicaPairs) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetAsyncCycle(v int64) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.AsyncCycle = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetBandwidth(v int64) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.Bandwidth = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetDescription(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.Description = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetDestinationDiskId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.DestinationDiskId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetDestinationRegion(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.DestinationRegion = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetLastRPO(v int64) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.LastRPO = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetPairName(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.PairName = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetReplicaPairId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.ReplicaPairId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetSourceDiskId(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.SourceDiskId = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetSourceRegion(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.SourceRegion = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetStatus(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.Status = &v
	return s
}

func (s *DescribeDiskReplicaPairsResponseBodyReplicaPairs) SetStatusMessage(v string) *DescribeDiskReplicaPairsResponseBodyReplicaPairs {
	s.StatusMessage = &v
	return s
}

type DescribeDiskReplicaPairsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDiskReplicaPairsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiskReplicaPairsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiskReplicaPairsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairsResponse) SetHeaders(v map[string]*string) *DescribeDiskReplicaPairsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskReplicaPairsResponse) SetBody(v *DescribeDiskReplicaPairsResponseBody) *DescribeDiskReplicaPairsResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	ResourceType   *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
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

func (s *DescribeRegionsRequest) SetResourceType(v string) *DescribeRegionsRequest {
	s.ResourceType = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	LocalName      *string                                    `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string                                    `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Zones          []*DescribeRegionsResponseBodyRegionsZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetZones(v []*DescribeRegionsResponseBodyRegionsZones) *DescribeRegionsResponseBodyRegions {
	s.Zones = v
	return s
}

type DescribeRegionsResponseBodyRegionsZones struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsZones) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsZones) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsZones {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsZones) SetZoneId(v string) *DescribeRegionsResponseBodyRegionsZones {
	s.ZoneId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type FailoverDiskReplicaPairRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s FailoverDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s FailoverDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *FailoverDiskReplicaPairRequest) SetClientToken(v string) *FailoverDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *FailoverDiskReplicaPairRequest) SetRegionId(v string) *FailoverDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *FailoverDiskReplicaPairRequest) SetReplicaPairId(v string) *FailoverDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

type FailoverDiskReplicaPairResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FailoverDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FailoverDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *FailoverDiskReplicaPairResponseBody) SetRequestId(v string) *FailoverDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type FailoverDiskReplicaPairResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FailoverDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FailoverDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s FailoverDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *FailoverDiskReplicaPairResponse) SetHeaders(v map[string]*string) *FailoverDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *FailoverDiskReplicaPairResponse) SetBody(v *FailoverDiskReplicaPairResponseBody) *FailoverDiskReplicaPairResponse {
	s.Body = v
	return s
}

type ModifyDiskReplicaPairRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PairName      *string `json:"PairName,omitempty" xml:"PairName,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s ModifyDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskReplicaPairRequest) SetClientToken(v string) *ModifyDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetDescription(v string) *ModifyDiskReplicaPairRequest {
	s.Description = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetPairName(v string) *ModifyDiskReplicaPairRequest {
	s.PairName = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetRegionId(v string) *ModifyDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDiskReplicaPairRequest) SetReplicaPairId(v string) *ModifyDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

type ModifyDiskReplicaPairResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDiskReplicaPairResponseBody) SetRequestId(v string) *ModifyDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDiskReplicaPairResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *ModifyDiskReplicaPairResponse) SetHeaders(v map[string]*string) *ModifyDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *ModifyDiskReplicaPairResponse) SetBody(v *ModifyDiskReplicaPairResponseBody) *ModifyDiskReplicaPairResponse {
	s.Body = v
	return s
}

type ReprotectDiskReplicaPairRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s ReprotectDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s ReprotectDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *ReprotectDiskReplicaPairRequest) SetClientToken(v string) *ReprotectDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *ReprotectDiskReplicaPairRequest) SetRegionId(v string) *ReprotectDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *ReprotectDiskReplicaPairRequest) SetReplicaPairId(v string) *ReprotectDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

type ReprotectDiskReplicaPairResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReprotectDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReprotectDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *ReprotectDiskReplicaPairResponseBody) SetRequestId(v string) *ReprotectDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type ReprotectDiskReplicaPairResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReprotectDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReprotectDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s ReprotectDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *ReprotectDiskReplicaPairResponse) SetHeaders(v map[string]*string) *ReprotectDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *ReprotectDiskReplicaPairResponse) SetBody(v *ReprotectDiskReplicaPairResponseBody) *ReprotectDiskReplicaPairResponse {
	s.Body = v
	return s
}

type StartDiskReplicaPairRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s StartDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *StartDiskReplicaPairRequest) SetClientToken(v string) *StartDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *StartDiskReplicaPairRequest) SetRegionId(v string) *StartDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *StartDiskReplicaPairRequest) SetReplicaPairId(v string) *StartDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

type StartDiskReplicaPairResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *StartDiskReplicaPairResponseBody) SetRequestId(v string) *StartDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type StartDiskReplicaPairResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *StartDiskReplicaPairResponse) SetHeaders(v map[string]*string) *StartDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *StartDiskReplicaPairResponse) SetBody(v *StartDiskReplicaPairResponseBody) *StartDiskReplicaPairResponse {
	s.Body = v
	return s
}

type StopDiskReplicaPairRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaPairId *string `json:"ReplicaPairId,omitempty" xml:"ReplicaPairId,omitempty"`
}

func (s StopDiskReplicaPairRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDiskReplicaPairRequest) GoString() string {
	return s.String()
}

func (s *StopDiskReplicaPairRequest) SetClientToken(v string) *StopDiskReplicaPairRequest {
	s.ClientToken = &v
	return s
}

func (s *StopDiskReplicaPairRequest) SetRegionId(v string) *StopDiskReplicaPairRequest {
	s.RegionId = &v
	return s
}

func (s *StopDiskReplicaPairRequest) SetReplicaPairId(v string) *StopDiskReplicaPairRequest {
	s.ReplicaPairId = &v
	return s
}

type StopDiskReplicaPairResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDiskReplicaPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDiskReplicaPairResponseBody) GoString() string {
	return s.String()
}

func (s *StopDiskReplicaPairResponseBody) SetRequestId(v string) *StopDiskReplicaPairResponseBody {
	s.RequestId = &v
	return s
}

type StopDiskReplicaPairResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopDiskReplicaPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopDiskReplicaPairResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDiskReplicaPairResponse) GoString() string {
	return s.String()
}

func (s *StopDiskReplicaPairResponse) SetHeaders(v map[string]*string) *StopDiskReplicaPairResponse {
	s.Headers = v
	return s
}

func (s *StopDiskReplicaPairResponse) SetBody(v *StopDiskReplicaPairResponseBody) *StopDiskReplicaPairResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("ebs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateDiskReplicaPairWithOptions(request *CreateDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *CreateDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Bandwidth"] = request.Bandwidth
	query["ChargeType"] = request.ChargeType
	query["ClientToken"] = request.ClientToken
	query["Description"] = request.Description
	query["DestinationDiskId"] = request.DestinationDiskId
	query["DestinationRegionId"] = request.DestinationRegionId
	query["DestinationZoneId"] = request.DestinationZoneId
	query["DiskId"] = request.DiskId
	query["PairName"] = request.PairName
	query["Period"] = request.Period
	query["PeriodUnit"] = request.PeriodUnit
	query["SourceRegionId"] = request.SourceRegionId
	query["SourceZoneId"] = request.SourceZoneId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDiskReplicaPair(request *CreateDiskReplicaPairRequest) (_result *CreateDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDiskReplicaPairResponse{}
	_body, _err := client.CreateDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDiskReplicaPairWithOptions(request *DeleteDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *DeleteDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["RegionId"] = request.RegionId
	query["ReplicaPairId"] = request.ReplicaPairId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDiskReplicaPair(request *DeleteDiskReplicaPairRequest) (_result *DeleteDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDiskReplicaPairResponse{}
	_body, _err := client.DeleteDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiskReplicaPairsWithOptions(request *DescribeDiskReplicaPairsRequest, runtime *util.RuntimeOptions) (_result *DescribeDiskReplicaPairsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["MaxResults"] = request.MaxResults
	query["NextToken"] = request.NextToken
	query["PairIds"] = request.PairIds
	query["RegionId"] = request.RegionId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiskReplicaPairs"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiskReplicaPairsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiskReplicaPairs(request *DescribeDiskReplicaPairsRequest) (_result *DescribeDiskReplicaPairsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiskReplicaPairsResponse{}
	_body, _err := client.DescribeDiskReplicaPairsWithOptions(request, runtime)
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
	query["AcceptLanguage"] = request.AcceptLanguage
	query["ResourceType"] = request.ResourceType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) FailoverDiskReplicaPairWithOptions(request *FailoverDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *FailoverDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["RegionId"] = request.RegionId
	query["ReplicaPairId"] = request.ReplicaPairId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("FailoverDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FailoverDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FailoverDiskReplicaPair(request *FailoverDiskReplicaPairRequest) (_result *FailoverDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FailoverDiskReplicaPairResponse{}
	_body, _err := client.FailoverDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDiskReplicaPairWithOptions(request *ModifyDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *ModifyDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["Description"] = request.Description
	query["PairName"] = request.PairName
	query["RegionId"] = request.RegionId
	query["ReplicaPairId"] = request.ReplicaPairId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDiskReplicaPair(request *ModifyDiskReplicaPairRequest) (_result *ModifyDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDiskReplicaPairResponse{}
	_body, _err := client.ModifyDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReprotectDiskReplicaPairWithOptions(request *ReprotectDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *ReprotectDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["RegionId"] = request.RegionId
	query["ReplicaPairId"] = request.ReplicaPairId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ReprotectDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ReprotectDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReprotectDiskReplicaPair(request *ReprotectDiskReplicaPairRequest) (_result *ReprotectDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReprotectDiskReplicaPairResponse{}
	_body, _err := client.ReprotectDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartDiskReplicaPairWithOptions(request *StartDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *StartDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["RegionId"] = request.RegionId
	query["ReplicaPairId"] = request.ReplicaPairId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("StartDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartDiskReplicaPair(request *StartDiskReplicaPairRequest) (_result *StartDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDiskReplicaPairResponse{}
	_body, _err := client.StartDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopDiskReplicaPairWithOptions(request *StopDiskReplicaPairRequest, runtime *util.RuntimeOptions) (_result *StopDiskReplicaPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientToken"] = request.ClientToken
	query["RegionId"] = request.RegionId
	query["ReplicaPairId"] = request.ReplicaPairId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("StopDiskReplicaPair"),
		Version:     tea.String("2021-07-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopDiskReplicaPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopDiskReplicaPair(request *StopDiskReplicaPairRequest) (_result *StopDiskReplicaPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDiskReplicaPairResponse{}
	_body, _err := client.StopDiskReplicaPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
