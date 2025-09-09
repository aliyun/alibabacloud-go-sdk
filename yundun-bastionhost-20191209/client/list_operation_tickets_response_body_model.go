// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationTicketsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationTickets(v []*ListOperationTicketsResponseBodyOperationTickets) *ListOperationTicketsResponseBody
	GetOperationTickets() []*ListOperationTicketsResponseBodyOperationTickets
	SetRequestId(v string) *ListOperationTicketsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListOperationTicketsResponseBody
	GetTotalCount() *int64
}

type ListOperationTicketsResponseBody struct {
	// The O\\&M applications to be reviewed.
	OperationTickets []*ListOperationTicketsResponseBodyOperationTickets `json:"OperationTickets,omitempty" xml:"OperationTickets,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of O\\&M applications to be reviewed.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOperationTicketsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOperationTicketsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationTicketsResponseBody) GetOperationTickets() []*ListOperationTicketsResponseBodyOperationTickets {
	return s.OperationTickets
}

func (s *ListOperationTicketsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOperationTicketsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListOperationTicketsResponseBody) SetOperationTickets(v []*ListOperationTicketsResponseBodyOperationTickets) *ListOperationTicketsResponseBody {
	s.OperationTickets = v
	return s
}

func (s *ListOperationTicketsResponseBody) SetRequestId(v string) *ListOperationTicketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOperationTicketsResponseBody) SetTotalCount(v int64) *ListOperationTicketsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOperationTicketsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListOperationTicketsResponseBodyOperationTickets struct {
	// The ID of the O\\&M applicant.
	//
	// example:
	//
	// 1
	ApplyUserId *string `json:"ApplyUserId,omitempty" xml:"ApplyUserId,omitempty"`
	// The username of the O\\&M applicant.
	//
	// example:
	//
	// test
	ApplyUsername *string `json:"ApplyUsername,omitempty" xml:"ApplyUsername,omitempty"`
	// The remarks entered when the O\\&M personnel applies for O\\&M permissions.
	//
	// example:
	//
	// Apply for O\\&M
	ApproveComment *string `json:"ApproveComment,omitempty" xml:"ApproveComment,omitempty"`
	// The ID of the asset account.
	//
	// example:
	//
	// 1
	AssetAccountId *string `json:"AssetAccountId,omitempty" xml:"AssetAccountId,omitempty"`
	// The username of the asset account.
	//
	// example:
	//
	// root
	AssetAccountName *string `json:"AssetAccountName,omitempty" xml:"AssetAccountName,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 10.167.XX.XX
	AssetAddress *string `json:"AssetAddress,omitempty" xml:"AssetAddress,omitempty"`
	// The ID of the asset.
	//
	// example:
	//
	// 2
	AssetId *string `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// poros-test
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The network domain ID of the asset.
	//
	// example:
	//
	// 2
	AssetNetworkDomainId *string `json:"AssetNetworkDomainId,omitempty" xml:"AssetNetworkDomainId,omitempty"`
	// The operating system of the asset.
	//
	// example:
	//
	// Linux
	AssetOs *string `json:"AssetOs,omitempty" xml:"AssetOs,omitempty"`
	// The name of the asset source to which the asset belongs. Valid values:
	//
	// 	- **Local**: an on-premises host.
	//
	// 	- **Ecs**: an Elastic Compute Service (ECS) instance.
	//
	// 	- **Rds**: an ApsaraDB RDS instance.
	//
	// 	- A third-party asset source.
	//
	// example:
	//
	// Local
	AssetSource *string `json:"AssetSource,omitempty" xml:"AssetSource,omitempty"`
	// The ID of the asset source to which the asset belongs.
	//
	// example:
	//
	// 1
	AssetSourceInstanceId *string `json:"AssetSourceInstanceId,omitempty" xml:"AssetSourceInstanceId,omitempty"`
	// The time when the O\\&M application was submitted. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1669965908
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The maximum number of logons applied by the O\\&M engineer. Valid values:
	//
	// 	- **0**: The number of logons is unlimited. The O\\&M engineer can log on to the specified asset for unlimited times during the validity period.
	//
	// 	- **1**: The O\\&M engineer can log on to the specified asset only once during the validity period.
	//
	// example:
	//
	// 0
	EffectCount *int32 `json:"EffectCount,omitempty" xml:"EffectCount,omitempty"`
	// The end time of the validity period. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1679393152
	EffectEndTime *int32 `json:"EffectEndTime,omitempty" xml:"EffectEndTime,omitempty"`
	// The start time of the validity period. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1685600242
	EffectStartTime *int32 `json:"EffectStartTime,omitempty" xml:"EffectStartTime,omitempty"`
	// The ID of the O\\&M application to be reviewed.
	//
	// example:
	//
	// 1
	OperationTicketId *string `json:"OperationTicketId,omitempty" xml:"OperationTicketId,omitempty"`
	// The O\\&M protocol.
	//
	// example:
	//
	// SSH
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	// The status of the review. Valid value:
	//
	// 	- Normal: to be reviewed
	//
	// example:
	//
	// Normal
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListOperationTicketsResponseBodyOperationTickets) String() string {
	return dara.Prettify(s)
}

func (s ListOperationTicketsResponseBodyOperationTickets) GoString() string {
	return s.String()
}

func (s *ListOperationTicketsResponseBodyOperationTickets) GetApplyUserId() *string {
	return s.ApplyUserId
}

func (s *ListOperationTicketsResponseBodyOperationTickets) GetApplyUsername() *string {
	return s.ApplyUsername
}

func (s *ListOperationTicketsResponseBodyOperationTickets) GetApproveComment() *string {
	return s.ApproveComment
}

func (s *ListOperationTicketsResponseBodyOperationTickets) GetAssetAccountId() *string {
	return s.AssetAccountId
}

func (s *ListOperationTicketsResponseBodyOperationTickets) GetAssetAccountName() *string {
	return s.AssetAccountName
}

func (s *ListOperationTicketsResponseBodyOperationTickets) GetAssetAddress() *string {
	return s.AssetAddress
}

func (s *ListOperationTicketsResponseBodyOperationTickets) GetAssetId() *string {
	return s.AssetId
}

func (s *ListOperationTicketsResponseBodyOperationTickets) GetAssetName() *string {
	return s.AssetName
}

func (s *ListOperationTicketsResponseBodyOperationTickets) GetAssetNetworkDomainId() *string {
	return s.AssetNetworkDomainId
}

func (s *ListOperationTicketsResponseBodyOperationTickets) GetAssetOs() *string {
	return s.AssetOs
}

func (s *ListOperationTicketsResponseBodyOperationTickets) GetAssetSource() *string {
	return s.AssetSource
}

func (s *ListOperationTicketsResponseBodyOperationTickets) GetAssetSourceInstanceId() *string {
	return s.AssetSourceInstanceId
}

func (s *ListOperationTicketsResponseBodyOperationTickets) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListOperationTicketsResponseBodyOperationTickets) GetEffectCount() *int32 {
	return s.EffectCount
}

func (s *ListOperationTicketsResponseBodyOperationTickets) GetEffectEndTime() *int32 {
	return s.EffectEndTime
}

func (s *ListOperationTicketsResponseBodyOperationTickets) GetEffectStartTime() *int32 {
	return s.EffectStartTime
}

func (s *ListOperationTicketsResponseBodyOperationTickets) GetOperationTicketId() *string {
	return s.OperationTicketId
}

func (s *ListOperationTicketsResponseBodyOperationTickets) GetProtocolName() *string {
	return s.ProtocolName
}

func (s *ListOperationTicketsResponseBodyOperationTickets) GetState() *string {
	return s.State
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetApplyUserId(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.ApplyUserId = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetApplyUsername(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.ApplyUsername = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetApproveComment(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.ApproveComment = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetAssetAccountId(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.AssetAccountId = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetAssetAccountName(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.AssetAccountName = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetAssetAddress(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.AssetAddress = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetAssetId(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.AssetId = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetAssetName(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.AssetName = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetAssetNetworkDomainId(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.AssetNetworkDomainId = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetAssetOs(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.AssetOs = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetAssetSource(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.AssetSource = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetAssetSourceInstanceId(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.AssetSourceInstanceId = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetCreatedTime(v int64) *ListOperationTicketsResponseBodyOperationTickets {
	s.CreatedTime = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetEffectCount(v int32) *ListOperationTicketsResponseBodyOperationTickets {
	s.EffectCount = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetEffectEndTime(v int32) *ListOperationTicketsResponseBodyOperationTickets {
	s.EffectEndTime = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetEffectStartTime(v int32) *ListOperationTicketsResponseBodyOperationTickets {
	s.EffectStartTime = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetOperationTicketId(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.OperationTicketId = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetProtocolName(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.ProtocolName = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetState(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.State = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) Validate() error {
	return dara.Validate(s)
}
