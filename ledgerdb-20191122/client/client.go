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

type AcceptMemberRequest struct {
	LedgerId  *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	KeyType   *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
}

func (s AcceptMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s AcceptMemberRequest) GoString() string {
	return s.String()
}

func (s *AcceptMemberRequest) SetLedgerId(v string) *AcceptMemberRequest {
	s.LedgerId = &v
	return s
}

func (s *AcceptMemberRequest) SetKeyType(v string) *AcceptMemberRequest {
	s.KeyType = &v
	return s
}

func (s *AcceptMemberRequest) SetPublicKey(v string) *AcceptMemberRequest {
	s.PublicKey = &v
	return s
}

type AcceptMemberResponseBody struct {
	MemberId  *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AcceptMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AcceptMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptMemberResponseBody) SetMemberId(v string) *AcceptMemberResponseBody {
	s.MemberId = &v
	return s
}

func (s *AcceptMemberResponseBody) SetRequestId(v string) *AcceptMemberResponseBody {
	s.RequestId = &v
	return s
}

type AcceptMemberResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AcceptMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AcceptMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s AcceptMemberResponse) GoString() string {
	return s.String()
}

func (s *AcceptMemberResponse) SetHeaders(v map[string]*string) *AcceptMemberResponse {
	s.Headers = v
	return s
}

func (s *AcceptMemberResponse) SetBody(v *AcceptMemberResponseBody) *AcceptMemberResponse {
	s.Body = v
	return s
}

type CreateVpcEndpointRequest struct {
	LedgerId    *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateVpcEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointRequest) SetLedgerId(v string) *CreateVpcEndpointRequest {
	s.LedgerId = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetVpcId(v string) *CreateVpcEndpointRequest {
	s.VpcId = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetVSwitchId(v string) *CreateVpcEndpointRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetClientToken(v string) *CreateVpcEndpointRequest {
	s.ClientToken = &v
	return s
}

type CreateVpcEndpointResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VpcEndpointId *string `json:"VpcEndpointId,omitempty" xml:"VpcEndpointId,omitempty"`
}

func (s CreateVpcEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointResponseBody) SetRequestId(v string) *CreateVpcEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcEndpointResponseBody) SetVpcEndpointId(v string) *CreateVpcEndpointResponseBody {
	s.VpcEndpointId = &v
	return s
}

type CreateVpcEndpointResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateVpcEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointResponse) SetHeaders(v map[string]*string) *CreateVpcEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcEndpointResponse) SetBody(v *CreateVpcEndpointResponseBody) *CreateVpcEndpointResponse {
	s.Body = v
	return s
}

type DeleteLedgerRequest struct {
	LedgerId *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
}

func (s DeleteLedgerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLedgerRequest) GoString() string {
	return s.String()
}

func (s *DeleteLedgerRequest) SetLedgerId(v string) *DeleteLedgerRequest {
	s.LedgerId = &v
	return s
}

type DeleteLedgerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLedgerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLedgerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLedgerResponseBody) SetRequestId(v string) *DeleteLedgerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLedgerResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteLedgerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLedgerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLedgerResponse) GoString() string {
	return s.String()
}

func (s *DeleteLedgerResponse) SetHeaders(v map[string]*string) *DeleteLedgerResponse {
	s.Headers = v
	return s
}

func (s *DeleteLedgerResponse) SetBody(v *DeleteLedgerResponseBody) *DeleteLedgerResponse {
	s.Body = v
	return s
}

type DeleteMemberRequest struct {
	LedgerId *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s DeleteMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMemberRequest) GoString() string {
	return s.String()
}

func (s *DeleteMemberRequest) SetLedgerId(v string) *DeleteMemberRequest {
	s.LedgerId = &v
	return s
}

func (s *DeleteMemberRequest) SetMemberId(v string) *DeleteMemberRequest {
	s.MemberId = &v
	return s
}

type DeleteMemberResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMemberResponseBody) SetRequestId(v string) *DeleteMemberResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMemberResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteMemberResponse) SetHeaders(v map[string]*string) *DeleteMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteMemberResponse) SetBody(v *DeleteMemberResponseBody) *DeleteMemberResponse {
	s.Body = v
	return s
}

type DeleteVpcEndpointRequest struct {
	LedgerId      *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId     *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcEndpointId *string `json:"VpcEndpointId,omitempty" xml:"VpcEndpointId,omitempty"`
}

func (s DeleteVpcEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcEndpointRequest) SetLedgerId(v string) *DeleteVpcEndpointRequest {
	s.LedgerId = &v
	return s
}

func (s *DeleteVpcEndpointRequest) SetVpcId(v string) *DeleteVpcEndpointRequest {
	s.VpcId = &v
	return s
}

func (s *DeleteVpcEndpointRequest) SetVSwitchId(v string) *DeleteVpcEndpointRequest {
	s.VSwitchId = &v
	return s
}

func (s *DeleteVpcEndpointRequest) SetVpcEndpointId(v string) *DeleteVpcEndpointRequest {
	s.VpcEndpointId = &v
	return s
}

type DeleteVpcEndpointResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpcEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpcEndpointResponseBody) SetRequestId(v string) *DeleteVpcEndpointResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVpcEndpointResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteVpcEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcEndpointResponse) SetHeaders(v map[string]*string) *DeleteVpcEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcEndpointResponse) SetBody(v *DeleteVpcEndpointResponseBody) *DeleteVpcEndpointResponse {
	s.Body = v
	return s
}

type DescribeLedgerRequest struct {
	LedgerId *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
}

func (s DescribeLedgerRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLedgerRequest) GoString() string {
	return s.String()
}

func (s *DescribeLedgerRequest) SetLedgerId(v string) *DescribeLedgerRequest {
	s.LedgerId = &v
	return s
}

type DescribeLedgerResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Ledger    *DescribeLedgerResponseBodyLedger `json:"Ledger,omitempty" xml:"Ledger,omitempty" type:"Struct"`
}

func (s DescribeLedgerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLedgerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLedgerResponseBody) SetRequestId(v string) *DescribeLedgerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLedgerResponseBody) SetLedger(v *DescribeLedgerResponseBodyLedger) *DescribeLedgerResponseBody {
	s.Ledger = v
	return s
}

type DescribeLedgerResponseBodyLedger struct {
	StorageClass      *string                                         `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	UpdateTime        *string                                         `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	JournalCount      *int64                                          `json:"JournalCount,omitempty" xml:"JournalCount,omitempty"`
	LedgerDescription *string                                         `json:"LedgerDescription,omitempty" xml:"LedgerDescription,omitempty"`
	CreateTime        *string                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LedgerType        *string                                         `json:"LedgerType,omitempty" xml:"LedgerType,omitempty"`
	LastTimeAnchor    *DescribeLedgerResponseBodyLedgerLastTimeAnchor `json:"LastTimeAnchor,omitempty" xml:"LastTimeAnchor,omitempty" type:"Struct"`
	LedgerId          *string                                         `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	RegionId          *string                                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MemberCount       *int64                                          `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
	LedgerStatus      *string                                         `json:"LedgerStatus,omitempty" xml:"LedgerStatus,omitempty"`
	TimeAnchorCount   *int64                                          `json:"TimeAnchorCount,omitempty" xml:"TimeAnchorCount,omitempty"`
	ZoneId            *string                                         `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	LedgerName        *string                                         `json:"LedgerName,omitempty" xml:"LedgerName,omitempty"`
	OwnerAliUid       *string                                         `json:"OwnerAliUid,omitempty" xml:"OwnerAliUid,omitempty"`
}

func (s DescribeLedgerResponseBodyLedger) String() string {
	return tea.Prettify(s)
}

func (s DescribeLedgerResponseBodyLedger) GoString() string {
	return s.String()
}

func (s *DescribeLedgerResponseBodyLedger) SetStorageClass(v string) *DescribeLedgerResponseBodyLedger {
	s.StorageClass = &v
	return s
}

func (s *DescribeLedgerResponseBodyLedger) SetUpdateTime(v string) *DescribeLedgerResponseBodyLedger {
	s.UpdateTime = &v
	return s
}

func (s *DescribeLedgerResponseBodyLedger) SetJournalCount(v int64) *DescribeLedgerResponseBodyLedger {
	s.JournalCount = &v
	return s
}

func (s *DescribeLedgerResponseBodyLedger) SetLedgerDescription(v string) *DescribeLedgerResponseBodyLedger {
	s.LedgerDescription = &v
	return s
}

func (s *DescribeLedgerResponseBodyLedger) SetCreateTime(v string) *DescribeLedgerResponseBodyLedger {
	s.CreateTime = &v
	return s
}

func (s *DescribeLedgerResponseBodyLedger) SetLedgerType(v string) *DescribeLedgerResponseBodyLedger {
	s.LedgerType = &v
	return s
}

func (s *DescribeLedgerResponseBodyLedger) SetLastTimeAnchor(v *DescribeLedgerResponseBodyLedgerLastTimeAnchor) *DescribeLedgerResponseBodyLedger {
	s.LastTimeAnchor = v
	return s
}

func (s *DescribeLedgerResponseBodyLedger) SetLedgerId(v string) *DescribeLedgerResponseBodyLedger {
	s.LedgerId = &v
	return s
}

func (s *DescribeLedgerResponseBodyLedger) SetRegionId(v string) *DescribeLedgerResponseBodyLedger {
	s.RegionId = &v
	return s
}

func (s *DescribeLedgerResponseBodyLedger) SetMemberCount(v int64) *DescribeLedgerResponseBodyLedger {
	s.MemberCount = &v
	return s
}

func (s *DescribeLedgerResponseBodyLedger) SetLedgerStatus(v string) *DescribeLedgerResponseBodyLedger {
	s.LedgerStatus = &v
	return s
}

func (s *DescribeLedgerResponseBodyLedger) SetTimeAnchorCount(v int64) *DescribeLedgerResponseBodyLedger {
	s.TimeAnchorCount = &v
	return s
}

func (s *DescribeLedgerResponseBodyLedger) SetZoneId(v string) *DescribeLedgerResponseBodyLedger {
	s.ZoneId = &v
	return s
}

func (s *DescribeLedgerResponseBodyLedger) SetLedgerName(v string) *DescribeLedgerResponseBodyLedger {
	s.LedgerName = &v
	return s
}

func (s *DescribeLedgerResponseBodyLedger) SetOwnerAliUid(v string) *DescribeLedgerResponseBodyLedger {
	s.OwnerAliUid = &v
	return s
}

type DescribeLedgerResponseBodyLedgerLastTimeAnchor struct {
	JournalId        *string `json:"JournalId,omitempty" xml:"JournalId,omitempty"`
	LedgerVersion    *string `json:"LedgerVersion,omitempty" xml:"LedgerVersion,omitempty"`
	TimeStamp        *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	LedgerDigest     *string `json:"LedgerDigest,omitempty" xml:"LedgerDigest,omitempty"`
	LedgerDigestType *string `json:"LedgerDigestType,omitempty" xml:"LedgerDigestType,omitempty"`
	Proof            *string `json:"Proof,omitempty" xml:"Proof,omitempty"`
}

func (s DescribeLedgerResponseBodyLedgerLastTimeAnchor) String() string {
	return tea.Prettify(s)
}

func (s DescribeLedgerResponseBodyLedgerLastTimeAnchor) GoString() string {
	return s.String()
}

func (s *DescribeLedgerResponseBodyLedgerLastTimeAnchor) SetJournalId(v string) *DescribeLedgerResponseBodyLedgerLastTimeAnchor {
	s.JournalId = &v
	return s
}

func (s *DescribeLedgerResponseBodyLedgerLastTimeAnchor) SetLedgerVersion(v string) *DescribeLedgerResponseBodyLedgerLastTimeAnchor {
	s.LedgerVersion = &v
	return s
}

func (s *DescribeLedgerResponseBodyLedgerLastTimeAnchor) SetTimeStamp(v string) *DescribeLedgerResponseBodyLedgerLastTimeAnchor {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLedgerResponseBodyLedgerLastTimeAnchor) SetLedgerDigest(v string) *DescribeLedgerResponseBodyLedgerLastTimeAnchor {
	s.LedgerDigest = &v
	return s
}

func (s *DescribeLedgerResponseBodyLedgerLastTimeAnchor) SetLedgerDigestType(v string) *DescribeLedgerResponseBodyLedgerLastTimeAnchor {
	s.LedgerDigestType = &v
	return s
}

func (s *DescribeLedgerResponseBodyLedgerLastTimeAnchor) SetProof(v string) *DescribeLedgerResponseBodyLedgerLastTimeAnchor {
	s.Proof = &v
	return s
}

type DescribeLedgerResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLedgerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLedgerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLedgerResponse) GoString() string {
	return s.String()
}

func (s *DescribeLedgerResponse) SetHeaders(v map[string]*string) *DescribeLedgerResponse {
	s.Headers = v
	return s
}

func (s *DescribeLedgerResponse) SetBody(v *DescribeLedgerResponseBody) *DescribeLedgerResponse {
	s.Body = v
	return s
}

type DescribeLedgersRequest struct {
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s DescribeLedgersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLedgersRequest) GoString() string {
	return s.String()
}

func (s *DescribeLedgersRequest) SetNextToken(v string) *DescribeLedgersRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeLedgersRequest) SetMaxResults(v int32) *DescribeLedgersRequest {
	s.MaxResults = &v
	return s
}

type DescribeLedgersResponseBody struct {
	NextToken  *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults *int32                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Ledgers    []*DescribeLedgersResponseBodyLedgers `json:"Ledgers,omitempty" xml:"Ledgers,omitempty" type:"Repeated"`
}

func (s DescribeLedgersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLedgersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLedgersResponseBody) SetNextToken(v string) *DescribeLedgersResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeLedgersResponseBody) SetRequestId(v string) *DescribeLedgersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLedgersResponseBody) SetMaxResults(v int32) *DescribeLedgersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeLedgersResponseBody) SetLedgers(v []*DescribeLedgersResponseBodyLedgers) *DescribeLedgersResponseBody {
	s.Ledgers = v
	return s
}

type DescribeLedgersResponseBodyLedgers struct {
	StorageClass      *string                                           `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	UpdateTime        *string                                           `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	JournalCount      *int64                                            `json:"JournalCount,omitempty" xml:"JournalCount,omitempty"`
	LedgerDescription *string                                           `json:"LedgerDescription,omitempty" xml:"LedgerDescription,omitempty"`
	CreateTime        *string                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LedgerType        *string                                           `json:"LedgerType,omitempty" xml:"LedgerType,omitempty"`
	LastTimeAnchor    *DescribeLedgersResponseBodyLedgersLastTimeAnchor `json:"LastTimeAnchor,omitempty" xml:"LastTimeAnchor,omitempty" type:"Struct"`
	LedgerId          *string                                           `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	RegionId          *string                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MemberCount       *int64                                            `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
	LedgerStatus      *string                                           `json:"LedgerStatus,omitempty" xml:"LedgerStatus,omitempty"`
	TimeAnchorCount   *int64                                            `json:"TimeAnchorCount,omitempty" xml:"TimeAnchorCount,omitempty"`
	ZoneId            *string                                           `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	LedgerName        *string                                           `json:"LedgerName,omitempty" xml:"LedgerName,omitempty"`
	OwnerAliUid       *string                                           `json:"OwnerAliUid,omitempty" xml:"OwnerAliUid,omitempty"`
}

func (s DescribeLedgersResponseBodyLedgers) String() string {
	return tea.Prettify(s)
}

func (s DescribeLedgersResponseBodyLedgers) GoString() string {
	return s.String()
}

func (s *DescribeLedgersResponseBodyLedgers) SetStorageClass(v string) *DescribeLedgersResponseBodyLedgers {
	s.StorageClass = &v
	return s
}

func (s *DescribeLedgersResponseBodyLedgers) SetUpdateTime(v string) *DescribeLedgersResponseBodyLedgers {
	s.UpdateTime = &v
	return s
}

func (s *DescribeLedgersResponseBodyLedgers) SetJournalCount(v int64) *DescribeLedgersResponseBodyLedgers {
	s.JournalCount = &v
	return s
}

func (s *DescribeLedgersResponseBodyLedgers) SetLedgerDescription(v string) *DescribeLedgersResponseBodyLedgers {
	s.LedgerDescription = &v
	return s
}

func (s *DescribeLedgersResponseBodyLedgers) SetCreateTime(v string) *DescribeLedgersResponseBodyLedgers {
	s.CreateTime = &v
	return s
}

func (s *DescribeLedgersResponseBodyLedgers) SetLedgerType(v string) *DescribeLedgersResponseBodyLedgers {
	s.LedgerType = &v
	return s
}

func (s *DescribeLedgersResponseBodyLedgers) SetLastTimeAnchor(v *DescribeLedgersResponseBodyLedgersLastTimeAnchor) *DescribeLedgersResponseBodyLedgers {
	s.LastTimeAnchor = v
	return s
}

func (s *DescribeLedgersResponseBodyLedgers) SetLedgerId(v string) *DescribeLedgersResponseBodyLedgers {
	s.LedgerId = &v
	return s
}

func (s *DescribeLedgersResponseBodyLedgers) SetRegionId(v string) *DescribeLedgersResponseBodyLedgers {
	s.RegionId = &v
	return s
}

func (s *DescribeLedgersResponseBodyLedgers) SetMemberCount(v int64) *DescribeLedgersResponseBodyLedgers {
	s.MemberCount = &v
	return s
}

func (s *DescribeLedgersResponseBodyLedgers) SetLedgerStatus(v string) *DescribeLedgersResponseBodyLedgers {
	s.LedgerStatus = &v
	return s
}

func (s *DescribeLedgersResponseBodyLedgers) SetTimeAnchorCount(v int64) *DescribeLedgersResponseBodyLedgers {
	s.TimeAnchorCount = &v
	return s
}

func (s *DescribeLedgersResponseBodyLedgers) SetZoneId(v string) *DescribeLedgersResponseBodyLedgers {
	s.ZoneId = &v
	return s
}

func (s *DescribeLedgersResponseBodyLedgers) SetLedgerName(v string) *DescribeLedgersResponseBodyLedgers {
	s.LedgerName = &v
	return s
}

func (s *DescribeLedgersResponseBodyLedgers) SetOwnerAliUid(v string) *DescribeLedgersResponseBodyLedgers {
	s.OwnerAliUid = &v
	return s
}

type DescribeLedgersResponseBodyLedgersLastTimeAnchor struct {
	JournalId        *string `json:"JournalId,omitempty" xml:"JournalId,omitempty"`
	LedgerVersion    *string `json:"LedgerVersion,omitempty" xml:"LedgerVersion,omitempty"`
	TimeStamp        *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	LedgerDigest     *string `json:"LedgerDigest,omitempty" xml:"LedgerDigest,omitempty"`
	LedgerDigestType *string `json:"LedgerDigestType,omitempty" xml:"LedgerDigestType,omitempty"`
	Proof            *string `json:"Proof,omitempty" xml:"Proof,omitempty"`
}

func (s DescribeLedgersResponseBodyLedgersLastTimeAnchor) String() string {
	return tea.Prettify(s)
}

func (s DescribeLedgersResponseBodyLedgersLastTimeAnchor) GoString() string {
	return s.String()
}

func (s *DescribeLedgersResponseBodyLedgersLastTimeAnchor) SetJournalId(v string) *DescribeLedgersResponseBodyLedgersLastTimeAnchor {
	s.JournalId = &v
	return s
}

func (s *DescribeLedgersResponseBodyLedgersLastTimeAnchor) SetLedgerVersion(v string) *DescribeLedgersResponseBodyLedgersLastTimeAnchor {
	s.LedgerVersion = &v
	return s
}

func (s *DescribeLedgersResponseBodyLedgersLastTimeAnchor) SetTimeStamp(v string) *DescribeLedgersResponseBodyLedgersLastTimeAnchor {
	s.TimeStamp = &v
	return s
}

func (s *DescribeLedgersResponseBodyLedgersLastTimeAnchor) SetLedgerDigest(v string) *DescribeLedgersResponseBodyLedgersLastTimeAnchor {
	s.LedgerDigest = &v
	return s
}

func (s *DescribeLedgersResponseBodyLedgersLastTimeAnchor) SetLedgerDigestType(v string) *DescribeLedgersResponseBodyLedgersLastTimeAnchor {
	s.LedgerDigestType = &v
	return s
}

func (s *DescribeLedgersResponseBodyLedgersLastTimeAnchor) SetProof(v string) *DescribeLedgersResponseBodyLedgersLastTimeAnchor {
	s.Proof = &v
	return s
}

type DescribeLedgersResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLedgersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLedgersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLedgersResponse) GoString() string {
	return s.String()
}

func (s *DescribeLedgersResponse) SetHeaders(v map[string]*string) *DescribeLedgersResponse {
	s.Headers = v
	return s
}

func (s *DescribeLedgersResponse) SetBody(v *DescribeLedgersResponseBody) *DescribeLedgersResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	ErrorCode *int32                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetErrorCode(v int32) *DescribeRegionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetSuccess(v bool) *DescribeRegionsResponseBody {
	s.Success = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

type DescribeTSARequest struct {
	SequenceWithinLedger *int64 `json:"SequenceWithinLedger,omitempty" xml:"SequenceWithinLedger,omitempty"`
}

func (s DescribeTSARequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTSARequest) GoString() string {
	return s.String()
}

func (s *DescribeTSARequest) SetSequenceWithinLedger(v int64) *DescribeTSARequest {
	s.SequenceWithinLedger = &v
	return s
}

type DescribeTSAResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TSADetail *DescribeTSAResponseBodyTSADetail `json:"TSADetail,omitempty" xml:"TSADetail,omitempty" type:"Struct"`
}

func (s DescribeTSAResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTSAResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTSAResponseBody) SetRequestId(v string) *DescribeTSAResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTSAResponseBody) SetTSADetail(v *DescribeTSAResponseBodyTSADetail) *DescribeTSAResponseBody {
	s.TSADetail = v
	return s
}

type DescribeTSAResponseBodyTSADetail struct {
	TS          *int64  `json:"TS,omitempty" xml:"TS,omitempty"`
	SN          *string `json:"SN,omitempty" xml:"SN,omitempty"`
	RootHash    *string `json:"RootHash,omitempty" xml:"RootHash,omitempty"`
	BlockNumber *int64  `json:"BlockNumber,omitempty" xml:"BlockNumber,omitempty"`
	CTSR        *string `json:"CTSR,omitempty" xml:"CTSR,omitempty"`
}

func (s DescribeTSAResponseBodyTSADetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeTSAResponseBodyTSADetail) GoString() string {
	return s.String()
}

func (s *DescribeTSAResponseBodyTSADetail) SetTS(v int64) *DescribeTSAResponseBodyTSADetail {
	s.TS = &v
	return s
}

func (s *DescribeTSAResponseBodyTSADetail) SetSN(v string) *DescribeTSAResponseBodyTSADetail {
	s.SN = &v
	return s
}

func (s *DescribeTSAResponseBodyTSADetail) SetRootHash(v string) *DescribeTSAResponseBodyTSADetail {
	s.RootHash = &v
	return s
}

func (s *DescribeTSAResponseBodyTSADetail) SetBlockNumber(v int64) *DescribeTSAResponseBodyTSADetail {
	s.BlockNumber = &v
	return s
}

func (s *DescribeTSAResponseBodyTSADetail) SetCTSR(v string) *DescribeTSAResponseBodyTSADetail {
	s.CTSR = &v
	return s
}

type DescribeTSAResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTSAResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTSAResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTSAResponse) GoString() string {
	return s.String()
}

func (s *DescribeTSAResponse) SetHeaders(v map[string]*string) *DescribeTSAResponse {
	s.Headers = v
	return s
}

func (s *DescribeTSAResponse) SetBody(v *DescribeTSAResponseBody) *DescribeTSAResponse {
	s.Body = v
	return s
}

type DisableMemberRequest struct {
	LedgerId *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s DisableMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableMemberRequest) GoString() string {
	return s.String()
}

func (s *DisableMemberRequest) SetLedgerId(v string) *DisableMemberRequest {
	s.LedgerId = &v
	return s
}

func (s *DisableMemberRequest) SetMemberId(v string) *DisableMemberRequest {
	s.MemberId = &v
	return s
}

type DisableMemberResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DisableMemberResponseBody) SetRequestId(v string) *DisableMemberResponseBody {
	s.RequestId = &v
	return s
}

type DisableMemberResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableMemberResponse) GoString() string {
	return s.String()
}

func (s *DisableMemberResponse) SetHeaders(v map[string]*string) *DisableMemberResponse {
	s.Headers = v
	return s
}

func (s *DisableMemberResponse) SetBody(v *DisableMemberResponseBody) *DisableMemberResponse {
	s.Body = v
	return s
}

type EnableMemberRequest struct {
	LedgerId *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s EnableMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableMemberRequest) GoString() string {
	return s.String()
}

func (s *EnableMemberRequest) SetLedgerId(v string) *EnableMemberRequest {
	s.LedgerId = &v
	return s
}

func (s *EnableMemberRequest) SetMemberId(v string) *EnableMemberRequest {
	s.MemberId = &v
	return s
}

type EnableMemberResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableMemberResponseBody) GoString() string {
	return s.String()
}

func (s *EnableMemberResponseBody) SetRequestId(v string) *EnableMemberResponseBody {
	s.RequestId = &v
	return s
}

type EnableMemberResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableMemberResponse) GoString() string {
	return s.String()
}

func (s *EnableMemberResponse) SetHeaders(v map[string]*string) *EnableMemberResponse {
	s.Headers = v
	return s
}

func (s *EnableMemberResponse) SetBody(v *EnableMemberResponseBody) *EnableMemberResponse {
	s.Body = v
	return s
}

type GetAccessAttributeRequest struct {
	LedgerId *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
}

func (s GetAccessAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccessAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetAccessAttributeRequest) SetLedgerId(v string) *GetAccessAttributeRequest {
	s.LedgerId = &v
	return s
}

type GetAccessAttributeResponseBody struct {
	EnableOpenAccess *string `json:"EnableOpenAccess,omitempty" xml:"EnableOpenAccess,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LedgerId         *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	OpenUri          *string `json:"OpenUri,omitempty" xml:"OpenUri,omitempty"`
}

func (s GetAccessAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccessAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessAttributeResponseBody) SetEnableOpenAccess(v string) *GetAccessAttributeResponseBody {
	s.EnableOpenAccess = &v
	return s
}

func (s *GetAccessAttributeResponseBody) SetRequestId(v string) *GetAccessAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccessAttributeResponseBody) SetLedgerId(v string) *GetAccessAttributeResponseBody {
	s.LedgerId = &v
	return s
}

func (s *GetAccessAttributeResponseBody) SetOpenUri(v string) *GetAccessAttributeResponseBody {
	s.OpenUri = &v
	return s
}

type GetAccessAttributeResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAccessAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAccessAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccessAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetAccessAttributeResponse) SetHeaders(v map[string]*string) *GetAccessAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetAccessAttributeResponse) SetBody(v *GetAccessAttributeResponseBody) *GetAccessAttributeResponse {
	s.Body = v
	return s
}

type GetIpWhiteListRequest struct {
	LedgerId *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
}

func (s GetIpWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIpWhiteListRequest) GoString() string {
	return s.String()
}

func (s *GetIpWhiteListRequest) SetLedgerId(v string) *GetIpWhiteListRequest {
	s.LedgerId = &v
	return s
}

type GetIpWhiteListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IpList    *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	LedgerId  *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
}

func (s GetIpWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIpWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *GetIpWhiteListResponseBody) SetRequestId(v string) *GetIpWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIpWhiteListResponseBody) SetIpList(v string) *GetIpWhiteListResponseBody {
	s.IpList = &v
	return s
}

func (s *GetIpWhiteListResponseBody) SetLedgerId(v string) *GetIpWhiteListResponseBody {
	s.LedgerId = &v
	return s
}

type GetIpWhiteListResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetIpWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIpWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIpWhiteListResponse) GoString() string {
	return s.String()
}

func (s *GetIpWhiteListResponse) SetHeaders(v map[string]*string) *GetIpWhiteListResponse {
	s.Headers = v
	return s
}

func (s *GetIpWhiteListResponse) SetBody(v *GetIpWhiteListResponseBody) *GetIpWhiteListResponse {
	s.Body = v
	return s
}

type GetJournalRequest struct {
	LedgerId  *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	JournalId *int64  `json:"JournalId,omitempty" xml:"JournalId,omitempty"`
}

func (s GetJournalRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJournalRequest) GoString() string {
	return s.String()
}

func (s *GetJournalRequest) SetLedgerId(v string) *GetJournalRequest {
	s.LedgerId = &v
	return s
}

func (s *GetJournalRequest) SetJournalId(v int64) *GetJournalRequest {
	s.JournalId = &v
	return s
}

type GetJournalResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Journal   *GetJournalResponseBodyJournal `json:"Journal,omitempty" xml:"Journal,omitempty" type:"Struct"`
}

func (s GetJournalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJournalResponseBody) GoString() string {
	return s.String()
}

func (s *GetJournalResponseBody) SetRequestId(v string) *GetJournalResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJournalResponseBody) SetJournal(v *GetJournalResponseBodyJournal) *GetJournalResponseBody {
	s.Journal = v
	return s
}

type GetJournalResponseBodyJournal struct {
	JournalId         *string   `json:"JournalId,omitempty" xml:"JournalId,omitempty"`
	Clues             []*string `json:"Clues,omitempty" xml:"Clues,omitempty" type:"Repeated"`
	PayloadJsonString *string   `json:"PayloadJsonString,omitempty" xml:"PayloadJsonString,omitempty"`
	JournalHash       *string   `json:"JournalHash,omitempty" xml:"JournalHash,omitempty"`
	Timestamp         *int64    `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	LedgerId          *string   `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	MemberId          *string   `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	PayloadType       *string   `json:"PayloadType,omitempty" xml:"PayloadType,omitempty"`
	ClientId          *string   `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
}

func (s GetJournalResponseBodyJournal) String() string {
	return tea.Prettify(s)
}

func (s GetJournalResponseBodyJournal) GoString() string {
	return s.String()
}

func (s *GetJournalResponseBodyJournal) SetJournalId(v string) *GetJournalResponseBodyJournal {
	s.JournalId = &v
	return s
}

func (s *GetJournalResponseBodyJournal) SetClues(v []*string) *GetJournalResponseBodyJournal {
	s.Clues = v
	return s
}

func (s *GetJournalResponseBodyJournal) SetPayloadJsonString(v string) *GetJournalResponseBodyJournal {
	s.PayloadJsonString = &v
	return s
}

func (s *GetJournalResponseBodyJournal) SetJournalHash(v string) *GetJournalResponseBodyJournal {
	s.JournalHash = &v
	return s
}

func (s *GetJournalResponseBodyJournal) SetTimestamp(v int64) *GetJournalResponseBodyJournal {
	s.Timestamp = &v
	return s
}

func (s *GetJournalResponseBodyJournal) SetLedgerId(v string) *GetJournalResponseBodyJournal {
	s.LedgerId = &v
	return s
}

func (s *GetJournalResponseBodyJournal) SetMemberId(v string) *GetJournalResponseBodyJournal {
	s.MemberId = &v
	return s
}

func (s *GetJournalResponseBodyJournal) SetPayloadType(v string) *GetJournalResponseBodyJournal {
	s.PayloadType = &v
	return s
}

func (s *GetJournalResponseBodyJournal) SetClientId(v string) *GetJournalResponseBodyJournal {
	s.ClientId = &v
	return s
}

type GetJournalResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJournalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJournalResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJournalResponse) GoString() string {
	return s.String()
}

func (s *GetJournalResponse) SetHeaders(v map[string]*string) *GetJournalResponse {
	s.Headers = v
	return s
}

func (s *GetJournalResponse) SetBody(v *GetJournalResponseBody) *GetJournalResponse {
	s.Body = v
	return s
}

type GetMemberRequest struct {
	LedgerId *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s GetMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMemberRequest) GoString() string {
	return s.String()
}

func (s *GetMemberRequest) SetLedgerId(v string) *GetMemberRequest {
	s.LedgerId = &v
	return s
}

func (s *GetMemberRequest) SetMemberId(v string) *GetMemberRequest {
	s.MemberId = &v
	return s
}

type GetMemberResponseBody struct {
	KMSKeyId      *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	KeyType       *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PublicKey     *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	CreateTime    *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LedgerId      *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	KeySource     *string `json:"KeySource,omitempty" xml:"KeySource,omitempty"`
	Role          *string `json:"Role,omitempty" xml:"Role,omitempty"`
	MemberId      *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	State         *string `json:"State,omitempty" xml:"State,omitempty"`
	UpdateTime    *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	KMSKeyVersion *string `json:"KMSKeyVersion,omitempty" xml:"KMSKeyVersion,omitempty"`
	KeyMeta       *string `json:"KeyMeta,omitempty" xml:"KeyMeta,omitempty"`
	AliUid        *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
}

func (s GetMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMemberResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemberResponseBody) SetKMSKeyId(v string) *GetMemberResponseBody {
	s.KMSKeyId = &v
	return s
}

func (s *GetMemberResponseBody) SetKeyType(v string) *GetMemberResponseBody {
	s.KeyType = &v
	return s
}

func (s *GetMemberResponseBody) SetRequestId(v string) *GetMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemberResponseBody) SetPublicKey(v string) *GetMemberResponseBody {
	s.PublicKey = &v
	return s
}

func (s *GetMemberResponseBody) SetCreateTime(v int64) *GetMemberResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetMemberResponseBody) SetLedgerId(v string) *GetMemberResponseBody {
	s.LedgerId = &v
	return s
}

func (s *GetMemberResponseBody) SetKeySource(v string) *GetMemberResponseBody {
	s.KeySource = &v
	return s
}

func (s *GetMemberResponseBody) SetRole(v string) *GetMemberResponseBody {
	s.Role = &v
	return s
}

func (s *GetMemberResponseBody) SetMemberId(v string) *GetMemberResponseBody {
	s.MemberId = &v
	return s
}

func (s *GetMemberResponseBody) SetState(v string) *GetMemberResponseBody {
	s.State = &v
	return s
}

func (s *GetMemberResponseBody) SetUpdateTime(v int64) *GetMemberResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetMemberResponseBody) SetKMSKeyVersion(v string) *GetMemberResponseBody {
	s.KMSKeyVersion = &v
	return s
}

func (s *GetMemberResponseBody) SetKeyMeta(v string) *GetMemberResponseBody {
	s.KeyMeta = &v
	return s
}

func (s *GetMemberResponseBody) SetAliUid(v string) *GetMemberResponseBody {
	s.AliUid = &v
	return s
}

type GetMemberResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMemberResponse) GoString() string {
	return s.String()
}

func (s *GetMemberResponse) SetHeaders(v map[string]*string) *GetMemberResponse {
	s.Headers = v
	return s
}

func (s *GetMemberResponse) SetBody(v *GetMemberResponseBody) *GetMemberResponse {
	s.Body = v
	return s
}

type GrantServiceLinkedRoleRequest struct {
	LedgerId *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
}

func (s GrantServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *GrantServiceLinkedRoleRequest) SetLedgerId(v string) *GrantServiceLinkedRoleRequest {
	s.LedgerId = &v
	return s
}

type GrantServiceLinkedRoleResponseBody struct {
	SLRStatus *string `json:"SLRStatus,omitempty" xml:"SLRStatus,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantServiceLinkedRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *GrantServiceLinkedRoleResponseBody) SetSLRStatus(v string) *GrantServiceLinkedRoleResponseBody {
	s.SLRStatus = &v
	return s
}

func (s *GrantServiceLinkedRoleResponseBody) SetRequestId(v string) *GrantServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

type GrantServiceLinkedRoleResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GrantServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrantServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *GrantServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *GrantServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *GrantServiceLinkedRoleResponse) SetBody(v *GrantServiceLinkedRoleResponseBody) *GrantServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type InviteMembersRequest struct {
	LedgerId *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	AliUids  *string `json:"AliUids,omitempty" xml:"AliUids,omitempty"`
}

func (s InviteMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s InviteMembersRequest) GoString() string {
	return s.String()
}

func (s *InviteMembersRequest) SetLedgerId(v string) *InviteMembersRequest {
	s.LedgerId = &v
	return s
}

func (s *InviteMembersRequest) SetAliUids(v string) *InviteMembersRequest {
	s.AliUids = &v
	return s
}

type InviteMembersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InviteMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InviteMembersResponseBody) GoString() string {
	return s.String()
}

func (s *InviteMembersResponseBody) SetRequestId(v string) *InviteMembersResponseBody {
	s.RequestId = &v
	return s
}

type InviteMembersResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InviteMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InviteMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s InviteMembersResponse) GoString() string {
	return s.String()
}

func (s *InviteMembersResponse) SetHeaders(v map[string]*string) *InviteMembersResponse {
	s.Headers = v
	return s
}

func (s *InviteMembersResponse) SetBody(v *InviteMembersResponseBody) *InviteMembersResponse {
	s.Body = v
	return s
}

type ListJournalsRequest struct {
	LedgerId   *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	Clue       *string `json:"Clue,omitempty" xml:"Clue,omitempty"`
	MemberId   *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Reverse    *bool   `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
}

func (s ListJournalsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJournalsRequest) GoString() string {
	return s.String()
}

func (s *ListJournalsRequest) SetLedgerId(v string) *ListJournalsRequest {
	s.LedgerId = &v
	return s
}

func (s *ListJournalsRequest) SetClue(v string) *ListJournalsRequest {
	s.Clue = &v
	return s
}

func (s *ListJournalsRequest) SetMemberId(v string) *ListJournalsRequest {
	s.MemberId = &v
	return s
}

func (s *ListJournalsRequest) SetNextToken(v string) *ListJournalsRequest {
	s.NextToken = &v
	return s
}

func (s *ListJournalsRequest) SetMaxResults(v int32) *ListJournalsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListJournalsRequest) SetReverse(v bool) *ListJournalsRequest {
	s.Reverse = &v
	return s
}

type ListJournalsResponseBody struct {
	NextToken  *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults *int32                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Journals   []*ListJournalsResponseBodyJournals `json:"Journals,omitempty" xml:"Journals,omitempty" type:"Repeated"`
}

func (s ListJournalsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJournalsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJournalsResponseBody) SetNextToken(v string) *ListJournalsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListJournalsResponseBody) SetRequestId(v string) *ListJournalsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJournalsResponseBody) SetMaxResults(v int32) *ListJournalsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListJournalsResponseBody) SetJournals(v []*ListJournalsResponseBodyJournals) *ListJournalsResponseBody {
	s.Journals = v
	return s
}

type ListJournalsResponseBodyJournals struct {
	JournalId         *string   `json:"JournalId,omitempty" xml:"JournalId,omitempty"`
	Clues             []*string `json:"Clues,omitempty" xml:"Clues,omitempty" type:"Repeated"`
	PayloadJsonString *string   `json:"PayloadJsonString,omitempty" xml:"PayloadJsonString,omitempty"`
	JournalHash       *string   `json:"JournalHash,omitempty" xml:"JournalHash,omitempty"`
	Timestamp         *int64    `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	LedgerId          *string   `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	MemberId          *string   `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	PayloadType       *string   `json:"PayloadType,omitempty" xml:"PayloadType,omitempty"`
	ClientId          *string   `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
}

func (s ListJournalsResponseBodyJournals) String() string {
	return tea.Prettify(s)
}

func (s ListJournalsResponseBodyJournals) GoString() string {
	return s.String()
}

func (s *ListJournalsResponseBodyJournals) SetJournalId(v string) *ListJournalsResponseBodyJournals {
	s.JournalId = &v
	return s
}

func (s *ListJournalsResponseBodyJournals) SetClues(v []*string) *ListJournalsResponseBodyJournals {
	s.Clues = v
	return s
}

func (s *ListJournalsResponseBodyJournals) SetPayloadJsonString(v string) *ListJournalsResponseBodyJournals {
	s.PayloadJsonString = &v
	return s
}

func (s *ListJournalsResponseBodyJournals) SetJournalHash(v string) *ListJournalsResponseBodyJournals {
	s.JournalHash = &v
	return s
}

func (s *ListJournalsResponseBodyJournals) SetTimestamp(v int64) *ListJournalsResponseBodyJournals {
	s.Timestamp = &v
	return s
}

func (s *ListJournalsResponseBodyJournals) SetLedgerId(v string) *ListJournalsResponseBodyJournals {
	s.LedgerId = &v
	return s
}

func (s *ListJournalsResponseBodyJournals) SetMemberId(v string) *ListJournalsResponseBodyJournals {
	s.MemberId = &v
	return s
}

func (s *ListJournalsResponseBodyJournals) SetPayloadType(v string) *ListJournalsResponseBodyJournals {
	s.PayloadType = &v
	return s
}

func (s *ListJournalsResponseBodyJournals) SetClientId(v string) *ListJournalsResponseBodyJournals {
	s.ClientId = &v
	return s
}

type ListJournalsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListJournalsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListJournalsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJournalsResponse) GoString() string {
	return s.String()
}

func (s *ListJournalsResponse) SetHeaders(v map[string]*string) *ListJournalsResponse {
	s.Headers = v
	return s
}

func (s *ListJournalsResponse) SetBody(v *ListJournalsResponseBody) *ListJournalsResponse {
	s.Body = v
	return s
}

type ListMembersRequest struct {
	LedgerId   *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMembersRequest) GoString() string {
	return s.String()
}

func (s *ListMembersRequest) SetLedgerId(v string) *ListMembersRequest {
	s.LedgerId = &v
	return s
}

func (s *ListMembersRequest) SetNextToken(v string) *ListMembersRequest {
	s.NextToken = &v
	return s
}

func (s *ListMembersRequest) SetMaxResults(v int32) *ListMembersRequest {
	s.MaxResults = &v
	return s
}

type ListMembersResponseBody struct {
	NextToken  *string                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults *int32                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Members    []*ListMembersResponseBodyMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
}

func (s ListMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListMembersResponseBody) SetNextToken(v string) *ListMembersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMembersResponseBody) SetRequestId(v string) *ListMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMembersResponseBody) SetMaxResults(v int32) *ListMembersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMembersResponseBody) SetMembers(v []*ListMembersResponseBodyMembers) *ListMembersResponseBody {
	s.Members = v
	return s
}

type ListMembersResponseBodyMembers struct {
	UpdateTime *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	KeyType    *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	State      *string `json:"State,omitempty" xml:"State,omitempty"`
	CreateTime *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PublicKey  *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	AliUid     *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	MemberId   *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	LedgerId   *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	Role       *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ListMembersResponseBodyMembers) String() string {
	return tea.Prettify(s)
}

func (s ListMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *ListMembersResponseBodyMembers) SetUpdateTime(v int64) *ListMembersResponseBodyMembers {
	s.UpdateTime = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetKeyType(v string) *ListMembersResponseBodyMembers {
	s.KeyType = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetState(v string) *ListMembersResponseBodyMembers {
	s.State = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetCreateTime(v int64) *ListMembersResponseBodyMembers {
	s.CreateTime = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetPublicKey(v string) *ListMembersResponseBodyMembers {
	s.PublicKey = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetAliUid(v string) *ListMembersResponseBodyMembers {
	s.AliUid = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetMemberId(v string) *ListMembersResponseBodyMembers {
	s.MemberId = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetLedgerId(v string) *ListMembersResponseBodyMembers {
	s.LedgerId = &v
	return s
}

func (s *ListMembersResponseBodyMembers) SetRole(v string) *ListMembersResponseBodyMembers {
	s.Role = &v
	return s
}

type ListMembersResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMembersResponse) GoString() string {
	return s.String()
}

func (s *ListMembersResponse) SetHeaders(v map[string]*string) *ListMembersResponse {
	s.Headers = v
	return s
}

func (s *ListMembersResponse) SetBody(v *ListMembersResponseBody) *ListMembersResponse {
	s.Body = v
	return s
}

type ListTimeAnchorsRequest struct {
	LedgerId   *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	Reverse    *bool   `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListTimeAnchorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTimeAnchorsRequest) GoString() string {
	return s.String()
}

func (s *ListTimeAnchorsRequest) SetLedgerId(v string) *ListTimeAnchorsRequest {
	s.LedgerId = &v
	return s
}

func (s *ListTimeAnchorsRequest) SetReverse(v bool) *ListTimeAnchorsRequest {
	s.Reverse = &v
	return s
}

func (s *ListTimeAnchorsRequest) SetNextToken(v string) *ListTimeAnchorsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTimeAnchorsRequest) SetMaxResults(v int32) *ListTimeAnchorsRequest {
	s.MaxResults = &v
	return s
}

type ListTimeAnchorsResponseBody struct {
	TimeAnchors []*ListTimeAnchorsResponseBodyTimeAnchors `json:"TimeAnchors,omitempty" xml:"TimeAnchors,omitempty" type:"Repeated"`
	NextToken   *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MaxResults  *int32                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListTimeAnchorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTimeAnchorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTimeAnchorsResponseBody) SetTimeAnchors(v []*ListTimeAnchorsResponseBodyTimeAnchors) *ListTimeAnchorsResponseBody {
	s.TimeAnchors = v
	return s
}

func (s *ListTimeAnchorsResponseBody) SetNextToken(v string) *ListTimeAnchorsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTimeAnchorsResponseBody) SetRequestId(v string) *ListTimeAnchorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTimeAnchorsResponseBody) SetMaxResults(v int32) *ListTimeAnchorsResponseBody {
	s.MaxResults = &v
	return s
}

type ListTimeAnchorsResponseBodyTimeAnchors struct {
	JournalId        *int64  `json:"JournalId,omitempty" xml:"JournalId,omitempty"`
	LedgerVersion    *int64  `json:"LedgerVersion,omitempty" xml:"LedgerVersion,omitempty"`
	TimeStamp        *int64  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	LedgerDigest     *string `json:"LedgerDigest,omitempty" xml:"LedgerDigest,omitempty"`
	LedgerDigestType *string `json:"LedgerDigestType,omitempty" xml:"LedgerDigestType,omitempty"`
	Proof            *string `json:"Proof,omitempty" xml:"Proof,omitempty"`
}

func (s ListTimeAnchorsResponseBodyTimeAnchors) String() string {
	return tea.Prettify(s)
}

func (s ListTimeAnchorsResponseBodyTimeAnchors) GoString() string {
	return s.String()
}

func (s *ListTimeAnchorsResponseBodyTimeAnchors) SetJournalId(v int64) *ListTimeAnchorsResponseBodyTimeAnchors {
	s.JournalId = &v
	return s
}

func (s *ListTimeAnchorsResponseBodyTimeAnchors) SetLedgerVersion(v int64) *ListTimeAnchorsResponseBodyTimeAnchors {
	s.LedgerVersion = &v
	return s
}

func (s *ListTimeAnchorsResponseBodyTimeAnchors) SetTimeStamp(v int64) *ListTimeAnchorsResponseBodyTimeAnchors {
	s.TimeStamp = &v
	return s
}

func (s *ListTimeAnchorsResponseBodyTimeAnchors) SetLedgerDigest(v string) *ListTimeAnchorsResponseBodyTimeAnchors {
	s.LedgerDigest = &v
	return s
}

func (s *ListTimeAnchorsResponseBodyTimeAnchors) SetLedgerDigestType(v string) *ListTimeAnchorsResponseBodyTimeAnchors {
	s.LedgerDigestType = &v
	return s
}

func (s *ListTimeAnchorsResponseBodyTimeAnchors) SetProof(v string) *ListTimeAnchorsResponseBodyTimeAnchors {
	s.Proof = &v
	return s
}

type ListTimeAnchorsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTimeAnchorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTimeAnchorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTimeAnchorsResponse) GoString() string {
	return s.String()
}

func (s *ListTimeAnchorsResponse) SetHeaders(v map[string]*string) *ListTimeAnchorsResponse {
	s.Headers = v
	return s
}

func (s *ListTimeAnchorsResponse) SetBody(v *ListTimeAnchorsResponseBody) *ListTimeAnchorsResponse {
	s.Body = v
	return s
}

type ListVpcEndpointsRequest struct {
	LedgerId   *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListVpcEndpointsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsRequest) SetLedgerId(v string) *ListVpcEndpointsRequest {
	s.LedgerId = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetNextToken(v string) *ListVpcEndpointsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointsRequest) SetMaxResults(v int32) *ListVpcEndpointsRequest {
	s.MaxResults = &v
	return s
}

type ListVpcEndpointsResponseBody struct {
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VpcEndpoints []*ListVpcEndpointsResponseBodyVpcEndpoints `json:"VpcEndpoints,omitempty" xml:"VpcEndpoints,omitempty" type:"Repeated"`
	MaxResults   *int32                                      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListVpcEndpointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsResponseBody) SetNextToken(v string) *ListVpcEndpointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointsResponseBody) SetRequestId(v string) *ListVpcEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVpcEndpointsResponseBody) SetVpcEndpoints(v []*ListVpcEndpointsResponseBodyVpcEndpoints) *ListVpcEndpointsResponseBody {
	s.VpcEndpoints = v
	return s
}

func (s *ListVpcEndpointsResponseBody) SetMaxResults(v int32) *ListVpcEndpointsResponseBody {
	s.MaxResults = &v
	return s
}

type ListVpcEndpointsResponseBodyVpcEndpoints struct {
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId     *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	CreateTime    *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Address       *string `json:"Address,omitempty" xml:"Address,omitempty"`
	VpcEndpointId *string `json:"VpcEndpointId,omitempty" xml:"VpcEndpointId,omitempty"`
	MemberId      *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	LedgerId      *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListVpcEndpointsResponseBodyVpcEndpoints) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointsResponseBodyVpcEndpoints) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsResponseBodyVpcEndpoints) SetStatus(v string) *ListVpcEndpointsResponseBodyVpcEndpoints {
	s.Status = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyVpcEndpoints) SetVpcId(v string) *ListVpcEndpointsResponseBodyVpcEndpoints {
	s.VpcId = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyVpcEndpoints) SetVSwitchId(v string) *ListVpcEndpointsResponseBodyVpcEndpoints {
	s.VSwitchId = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyVpcEndpoints) SetCreateTime(v int64) *ListVpcEndpointsResponseBodyVpcEndpoints {
	s.CreateTime = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyVpcEndpoints) SetAddress(v string) *ListVpcEndpointsResponseBodyVpcEndpoints {
	s.Address = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyVpcEndpoints) SetVpcEndpointId(v string) *ListVpcEndpointsResponseBodyVpcEndpoints {
	s.VpcEndpointId = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyVpcEndpoints) SetMemberId(v string) *ListVpcEndpointsResponseBodyVpcEndpoints {
	s.MemberId = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyVpcEndpoints) SetLedgerId(v string) *ListVpcEndpointsResponseBodyVpcEndpoints {
	s.LedgerId = &v
	return s
}

func (s *ListVpcEndpointsResponseBodyVpcEndpoints) SetRegionId(v string) *ListVpcEndpointsResponseBodyVpcEndpoints {
	s.RegionId = &v
	return s
}

type ListVpcEndpointsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVpcEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVpcEndpointsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVpcEndpointsResponse) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointsResponse) SetHeaders(v map[string]*string) *ListVpcEndpointsResponse {
	s.Headers = v
	return s
}

func (s *ListVpcEndpointsResponse) SetBody(v *ListVpcEndpointsResponseBody) *ListVpcEndpointsResponse {
	s.Body = v
	return s
}

type ModifyAccessAttributeRequest struct {
	LedgerId         *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	EnableOpenAccess *bool   `json:"EnableOpenAccess,omitempty" xml:"EnableOpenAccess,omitempty"`
}

func (s ModifyAccessAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccessAttributeRequest) SetLedgerId(v string) *ModifyAccessAttributeRequest {
	s.LedgerId = &v
	return s
}

func (s *ModifyAccessAttributeRequest) SetEnableOpenAccess(v bool) *ModifyAccessAttributeRequest {
	s.EnableOpenAccess = &v
	return s
}

type ModifyAccessAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccessAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccessAttributeResponseBody) SetRequestId(v string) *ModifyAccessAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccessAttributeResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAccessAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAccessAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccessAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccessAttributeResponse) SetHeaders(v map[string]*string) *ModifyAccessAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccessAttributeResponse) SetBody(v *ModifyAccessAttributeResponseBody) *ModifyAccessAttributeResponse {
	s.Body = v
	return s
}

type ModifyIpWhiteListRequest struct {
	LedgerId   *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	IpList     *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
}

func (s ModifyIpWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ModifyIpWhiteListRequest) SetLedgerId(v string) *ModifyIpWhiteListRequest {
	s.LedgerId = &v
	return s
}

func (s *ModifyIpWhiteListRequest) SetModifyType(v string) *ModifyIpWhiteListRequest {
	s.ModifyType = &v
	return s
}

func (s *ModifyIpWhiteListRequest) SetIpList(v string) *ModifyIpWhiteListRequest {
	s.IpList = &v
	return s
}

type ModifyIpWhiteListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIpWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIpWhiteListResponseBody) SetRequestId(v string) *ModifyIpWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type ModifyIpWhiteListResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyIpWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyIpWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ModifyIpWhiteListResponse) SetHeaders(v map[string]*string) *ModifyIpWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ModifyIpWhiteListResponse) SetBody(v *ModifyIpWhiteListResponseBody) *ModifyIpWhiteListResponse {
	s.Body = v
	return s
}

type ModifyLedgerAttributeRequest struct {
	LedgerId          *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	LedgerName        *string `json:"LedgerName,omitempty" xml:"LedgerName,omitempty"`
	LedgerDescription *string `json:"LedgerDescription,omitempty" xml:"LedgerDescription,omitempty"`
}

func (s ModifyLedgerAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLedgerAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyLedgerAttributeRequest) SetLedgerId(v string) *ModifyLedgerAttributeRequest {
	s.LedgerId = &v
	return s
}

func (s *ModifyLedgerAttributeRequest) SetLedgerName(v string) *ModifyLedgerAttributeRequest {
	s.LedgerName = &v
	return s
}

func (s *ModifyLedgerAttributeRequest) SetLedgerDescription(v string) *ModifyLedgerAttributeRequest {
	s.LedgerDescription = &v
	return s
}

type ModifyLedgerAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLedgerAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyLedgerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLedgerAttributeResponseBody) SetRequestId(v string) *ModifyLedgerAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyLedgerAttributeResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyLedgerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyLedgerAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLedgerAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyLedgerAttributeResponse) SetHeaders(v map[string]*string) *ModifyLedgerAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyLedgerAttributeResponse) SetBody(v *ModifyLedgerAttributeResponseBody) *ModifyLedgerAttributeResponse {
	s.Body = v
	return s
}

type ModifyMemberACLsRequest struct {
	LedgerId *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	Role     *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ModifyMemberACLsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMemberACLsRequest) GoString() string {
	return s.String()
}

func (s *ModifyMemberACLsRequest) SetLedgerId(v string) *ModifyMemberACLsRequest {
	s.LedgerId = &v
	return s
}

func (s *ModifyMemberACLsRequest) SetMemberId(v string) *ModifyMemberACLsRequest {
	s.MemberId = &v
	return s
}

func (s *ModifyMemberACLsRequest) SetRole(v string) *ModifyMemberACLsRequest {
	s.Role = &v
	return s
}

type ModifyMemberACLsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMemberACLsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMemberACLsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMemberACLsResponseBody) SetRequestId(v string) *ModifyMemberACLsResponseBody {
	s.RequestId = &v
	return s
}

type ModifyMemberACLsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyMemberACLsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyMemberACLsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMemberACLsResponse) GoString() string {
	return s.String()
}

func (s *ModifyMemberACLsResponse) SetHeaders(v map[string]*string) *ModifyMemberACLsResponse {
	s.Headers = v
	return s
}

func (s *ModifyMemberACLsResponse) SetBody(v *ModifyMemberACLsResponseBody) *ModifyMemberACLsResponse {
	s.Body = v
	return s
}

type ModifyMemberKeyRequest struct {
	LedgerId  *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	MemberId  *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	KeyType   *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
}

func (s ModifyMemberKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMemberKeyRequest) GoString() string {
	return s.String()
}

func (s *ModifyMemberKeyRequest) SetLedgerId(v string) *ModifyMemberKeyRequest {
	s.LedgerId = &v
	return s
}

func (s *ModifyMemberKeyRequest) SetMemberId(v string) *ModifyMemberKeyRequest {
	s.MemberId = &v
	return s
}

func (s *ModifyMemberKeyRequest) SetKeyType(v string) *ModifyMemberKeyRequest {
	s.KeyType = &v
	return s
}

func (s *ModifyMemberKeyRequest) SetPublicKey(v string) *ModifyMemberKeyRequest {
	s.PublicKey = &v
	return s
}

type ModifyMemberKeyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMemberKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMemberKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMemberKeyResponseBody) SetRequestId(v string) *ModifyMemberKeyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyMemberKeyResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyMemberKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyMemberKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMemberKeyResponse) GoString() string {
	return s.String()
}

func (s *ModifyMemberKeyResponse) SetHeaders(v map[string]*string) *ModifyMemberKeyResponse {
	s.Headers = v
	return s
}

func (s *ModifyMemberKeyResponse) SetBody(v *ModifyMemberKeyResponseBody) *ModifyMemberKeyResponse {
	s.Body = v
	return s
}

type UpdateMemberKeyByKMSRequest struct {
	LedgerId      *string `json:"LedgerId,omitempty" xml:"LedgerId,omitempty"`
	MemberId      *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	KeyType       *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	KMSKeyId      *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	KMSKeyVersion *string `json:"KMSKeyVersion,omitempty" xml:"KMSKeyVersion,omitempty"`
}

func (s UpdateMemberKeyByKMSRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemberKeyByKMSRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemberKeyByKMSRequest) SetLedgerId(v string) *UpdateMemberKeyByKMSRequest {
	s.LedgerId = &v
	return s
}

func (s *UpdateMemberKeyByKMSRequest) SetMemberId(v string) *UpdateMemberKeyByKMSRequest {
	s.MemberId = &v
	return s
}

func (s *UpdateMemberKeyByKMSRequest) SetKeyType(v string) *UpdateMemberKeyByKMSRequest {
	s.KeyType = &v
	return s
}

func (s *UpdateMemberKeyByKMSRequest) SetKMSKeyId(v string) *UpdateMemberKeyByKMSRequest {
	s.KMSKeyId = &v
	return s
}

func (s *UpdateMemberKeyByKMSRequest) SetKMSKeyVersion(v string) *UpdateMemberKeyByKMSRequest {
	s.KMSKeyVersion = &v
	return s
}

type UpdateMemberKeyByKMSResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMemberKeyByKMSResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemberKeyByKMSResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMemberKeyByKMSResponseBody) SetRequestId(v string) *UpdateMemberKeyByKMSResponseBody {
	s.RequestId = &v
	return s
}

type UpdateMemberKeyByKMSResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateMemberKeyByKMSResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMemberKeyByKMSResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemberKeyByKMSResponse) GoString() string {
	return s.String()
}

func (s *UpdateMemberKeyByKMSResponse) SetHeaders(v map[string]*string) *UpdateMemberKeyByKMSResponse {
	s.Headers = v
	return s
}

func (s *UpdateMemberKeyByKMSResponse) SetBody(v *UpdateMemberKeyByKMSResponseBody) *UpdateMemberKeyByKMSResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("ledgerdb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AcceptMemberWithOptions(request *AcceptMemberRequest, runtime *util.RuntimeOptions) (_result *AcceptMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AcceptMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AcceptMember"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AcceptMember(request *AcceptMemberRequest) (_result *AcceptMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AcceptMemberResponse{}
	_body, _err := client.AcceptMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVpcEndpointWithOptions(request *CreateVpcEndpointRequest, runtime *util.RuntimeOptions) (_result *CreateVpcEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateVpcEndpointResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateVpcEndpoint"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVpcEndpoint(request *CreateVpcEndpointRequest) (_result *CreateVpcEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVpcEndpointResponse{}
	_body, _err := client.CreateVpcEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLedgerWithOptions(request *DeleteLedgerRequest, runtime *util.RuntimeOptions) (_result *DeleteLedgerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteLedgerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteLedger"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLedger(request *DeleteLedgerRequest) (_result *DeleteLedgerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLedgerResponse{}
	_body, _err := client.DeleteLedgerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMemberWithOptions(request *DeleteMemberRequest, runtime *util.RuntimeOptions) (_result *DeleteMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMember"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMember(request *DeleteMemberRequest) (_result *DeleteMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMemberResponse{}
	_body, _err := client.DeleteMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVpcEndpointWithOptions(request *DeleteVpcEndpointRequest, runtime *util.RuntimeOptions) (_result *DeleteVpcEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVpcEndpointResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVpcEndpoint"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVpcEndpoint(request *DeleteVpcEndpointRequest) (_result *DeleteVpcEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVpcEndpointResponse{}
	_body, _err := client.DeleteVpcEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLedgerWithOptions(request *DescribeLedgerRequest, runtime *util.RuntimeOptions) (_result *DescribeLedgerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeLedgerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLedger"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLedger(request *DescribeLedgerRequest) (_result *DescribeLedgerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLedgerResponse{}
	_body, _err := client.DescribeLedgerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLedgersWithOptions(request *DescribeLedgersRequest, runtime *util.RuntimeOptions) (_result *DescribeLedgersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeLedgersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLedgers"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLedgers(request *DescribeLedgersRequest) (_result *DescribeLedgersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLedgersResponse{}
	_body, _err := client.DescribeLedgersWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeTSAWithOptions(request *DescribeTSARequest, runtime *util.RuntimeOptions) (_result *DescribeTSAResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTSAResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTSA"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTSA(request *DescribeTSARequest) (_result *DescribeTSAResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTSAResponse{}
	_body, _err := client.DescribeTSAWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableMemberWithOptions(request *DisableMemberRequest, runtime *util.RuntimeOptions) (_result *DisableMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableMember"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableMember(request *DisableMemberRequest) (_result *DisableMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableMemberResponse{}
	_body, _err := client.DisableMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableMemberWithOptions(request *EnableMemberRequest, runtime *util.RuntimeOptions) (_result *EnableMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableMember"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableMember(request *EnableMemberRequest) (_result *EnableMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableMemberResponse{}
	_body, _err := client.EnableMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccessAttributeWithOptions(request *GetAccessAttributeRequest, runtime *util.RuntimeOptions) (_result *GetAccessAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetAccessAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAccessAttribute"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccessAttribute(request *GetAccessAttributeRequest) (_result *GetAccessAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccessAttributeResponse{}
	_body, _err := client.GetAccessAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIpWhiteListWithOptions(request *GetIpWhiteListRequest, runtime *util.RuntimeOptions) (_result *GetIpWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetIpWhiteListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetIpWhiteList"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIpWhiteList(request *GetIpWhiteListRequest) (_result *GetIpWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIpWhiteListResponse{}
	_body, _err := client.GetIpWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJournalWithOptions(request *GetJournalRequest, runtime *util.RuntimeOptions) (_result *GetJournalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetJournalResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetJournal"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJournal(request *GetJournalRequest) (_result *GetJournalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJournalResponse{}
	_body, _err := client.GetJournalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMemberWithOptions(request *GetMemberRequest, runtime *util.RuntimeOptions) (_result *GetMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMember"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMember(request *GetMemberRequest) (_result *GetMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMemberResponse{}
	_body, _err := client.GetMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantServiceLinkedRoleWithOptions(request *GrantServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *GrantServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GrantServiceLinkedRoleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GrantServiceLinkedRole"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantServiceLinkedRole(request *GrantServiceLinkedRoleRequest) (_result *GrantServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantServiceLinkedRoleResponse{}
	_body, _err := client.GrantServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InviteMembersWithOptions(request *InviteMembersRequest, runtime *util.RuntimeOptions) (_result *InviteMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InviteMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InviteMembers"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InviteMembers(request *InviteMembersRequest) (_result *InviteMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InviteMembersResponse{}
	_body, _err := client.InviteMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJournalsWithOptions(request *ListJournalsRequest, runtime *util.RuntimeOptions) (_result *ListJournalsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListJournalsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListJournals"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJournals(request *ListJournalsRequest) (_result *ListJournalsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListJournalsResponse{}
	_body, _err := client.ListJournalsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMembersWithOptions(request *ListMembersRequest, runtime *util.RuntimeOptions) (_result *ListMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListMembers"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMembers(request *ListMembersRequest) (_result *ListMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMembersResponse{}
	_body, _err := client.ListMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTimeAnchorsWithOptions(request *ListTimeAnchorsRequest, runtime *util.RuntimeOptions) (_result *ListTimeAnchorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListTimeAnchorsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTimeAnchors"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTimeAnchors(request *ListTimeAnchorsRequest) (_result *ListTimeAnchorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTimeAnchorsResponse{}
	_body, _err := client.ListTimeAnchorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVpcEndpointsWithOptions(request *ListVpcEndpointsRequest, runtime *util.RuntimeOptions) (_result *ListVpcEndpointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListVpcEndpointsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListVpcEndpoints"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVpcEndpoints(request *ListVpcEndpointsRequest) (_result *ListVpcEndpointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVpcEndpointsResponse{}
	_body, _err := client.ListVpcEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAccessAttributeWithOptions(request *ModifyAccessAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyAccessAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAccessAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAccessAttribute"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccessAttribute(request *ModifyAccessAttributeRequest) (_result *ModifyAccessAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccessAttributeResponse{}
	_body, _err := client.ModifyAccessAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyIpWhiteListWithOptions(request *ModifyIpWhiteListRequest, runtime *util.RuntimeOptions) (_result *ModifyIpWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyIpWhiteListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyIpWhiteList"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyIpWhiteList(request *ModifyIpWhiteListRequest) (_result *ModifyIpWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyIpWhiteListResponse{}
	_body, _err := client.ModifyIpWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyLedgerAttributeWithOptions(request *ModifyLedgerAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyLedgerAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyLedgerAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyLedgerAttribute"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyLedgerAttribute(request *ModifyLedgerAttributeRequest) (_result *ModifyLedgerAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLedgerAttributeResponse{}
	_body, _err := client.ModifyLedgerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMemberACLsWithOptions(request *ModifyMemberACLsRequest, runtime *util.RuntimeOptions) (_result *ModifyMemberACLsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyMemberACLsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyMemberACLs"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMemberACLs(request *ModifyMemberACLsRequest) (_result *ModifyMemberACLsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMemberACLsResponse{}
	_body, _err := client.ModifyMemberACLsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMemberKeyWithOptions(request *ModifyMemberKeyRequest, runtime *util.RuntimeOptions) (_result *ModifyMemberKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyMemberKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyMemberKey"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMemberKey(request *ModifyMemberKeyRequest) (_result *ModifyMemberKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMemberKeyResponse{}
	_body, _err := client.ModifyMemberKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMemberKeyByKMSWithOptions(request *UpdateMemberKeyByKMSRequest, runtime *util.RuntimeOptions) (_result *UpdateMemberKeyByKMSResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateMemberKeyByKMSResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateMemberKeyByKMS"), tea.String("2019-11-22"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMemberKeyByKMS(request *UpdateMemberKeyByKMSRequest) (_result *UpdateMemberKeyByKMSResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMemberKeyByKMSResponse{}
	_body, _err := client.UpdateMemberKeyByKMSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
