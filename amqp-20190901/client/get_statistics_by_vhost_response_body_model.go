// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStatisticsByVhostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetStatisticsByVhostResponseBody
	GetCode() *int32
	SetData(v *GetStatisticsByVhostResponseBodyData) *GetStatisticsByVhostResponseBody
	GetData() *GetStatisticsByVhostResponseBodyData
	SetMessage(v string) *GetStatisticsByVhostResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetStatisticsByVhostResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetStatisticsByVhostResponseBody
	GetSuccess() *bool
}

type GetStatisticsByVhostResponseBody struct {
	Code      *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetStatisticsByVhostResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetStatisticsByVhostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStatisticsByVhostResponseBody) GoString() string {
	return s.String()
}

func (s *GetStatisticsByVhostResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetStatisticsByVhostResponseBody) GetData() *GetStatisticsByVhostResponseBodyData {
	return s.Data
}

func (s *GetStatisticsByVhostResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetStatisticsByVhostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStatisticsByVhostResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStatisticsByVhostResponseBody) SetCode(v int32) *GetStatisticsByVhostResponseBody {
	s.Code = &v
	return s
}

func (s *GetStatisticsByVhostResponseBody) SetData(v *GetStatisticsByVhostResponseBodyData) *GetStatisticsByVhostResponseBody {
	s.Data = v
	return s
}

func (s *GetStatisticsByVhostResponseBody) SetMessage(v string) *GetStatisticsByVhostResponseBody {
	s.Message = &v
	return s
}

func (s *GetStatisticsByVhostResponseBody) SetRequestId(v string) *GetStatisticsByVhostResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStatisticsByVhostResponseBody) SetSuccess(v bool) *GetStatisticsByVhostResponseBody {
	s.Success = &v
	return s
}

func (s *GetStatisticsByVhostResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStatisticsByVhostResponseBodyData struct {
	ConnectionStatistics []*GetStatisticsByVhostResponseBodyDataConnectionStatistics `json:"ConnectionStatistics,omitempty" xml:"ConnectionStatistics,omitempty" type:"Repeated"`
}

func (s GetStatisticsByVhostResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetStatisticsByVhostResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStatisticsByVhostResponseBodyData) GetConnectionStatistics() []*GetStatisticsByVhostResponseBodyDataConnectionStatistics {
	return s.ConnectionStatistics
}

func (s *GetStatisticsByVhostResponseBodyData) SetConnectionStatistics(v []*GetStatisticsByVhostResponseBodyDataConnectionStatistics) *GetStatisticsByVhostResponseBodyData {
	s.ConnectionStatistics = v
	return s
}

func (s *GetStatisticsByVhostResponseBodyData) Validate() error {
	if s.ConnectionStatistics != nil {
		for _, item := range s.ConnectionStatistics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetStatisticsByVhostResponseBodyDataConnectionStatistics struct {
	AccessKey             *string                                                                        `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	ChannelNum            *int32                                                                         `json:"ChannelNum,omitempty" xml:"ChannelNum,omitempty"`
	ChannelStatisticsList *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsList `json:"ChannelStatisticsList,omitempty" xml:"ChannelStatisticsList,omitempty" type:"Struct"`
	ConnectionName        *string                                                                        `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	DeliverQps            *float32                                                                       `json:"DeliverQps,omitempty" xml:"DeliverQps,omitempty"`
	Protocol              *string                                                                        `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	PublishQps            *float32                                                                       `json:"PublishQps,omitempty" xml:"PublishQps,omitempty"`
	RemoteAddress         *string                                                                        `json:"RemoteAddress,omitempty" xml:"RemoteAddress,omitempty"`
	SecurityTransport     *bool                                                                          `json:"SecurityTransport,omitempty" xml:"SecurityTransport,omitempty"`
	State                 *int32                                                                         `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetStatisticsByVhostResponseBodyDataConnectionStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetStatisticsByVhostResponseBodyDataConnectionStatistics) GoString() string {
	return s.String()
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatistics) GetAccessKey() *string {
	return s.AccessKey
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatistics) GetChannelNum() *int32 {
	return s.ChannelNum
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatistics) GetChannelStatisticsList() *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsList {
	return s.ChannelStatisticsList
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatistics) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatistics) GetDeliverQps() *float32 {
	return s.DeliverQps
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatistics) GetProtocol() *string {
	return s.Protocol
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatistics) GetPublishQps() *float32 {
	return s.PublishQps
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatistics) GetRemoteAddress() *string {
	return s.RemoteAddress
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatistics) GetSecurityTransport() *bool {
	return s.SecurityTransport
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatistics) GetState() *int32 {
	return s.State
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatistics) SetAccessKey(v string) *GetStatisticsByVhostResponseBodyDataConnectionStatistics {
	s.AccessKey = &v
	return s
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatistics) SetChannelNum(v int32) *GetStatisticsByVhostResponseBodyDataConnectionStatistics {
	s.ChannelNum = &v
	return s
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatistics) SetChannelStatisticsList(v *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsList) *GetStatisticsByVhostResponseBodyDataConnectionStatistics {
	s.ChannelStatisticsList = v
	return s
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatistics) SetConnectionName(v string) *GetStatisticsByVhostResponseBodyDataConnectionStatistics {
	s.ConnectionName = &v
	return s
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatistics) SetDeliverQps(v float32) *GetStatisticsByVhostResponseBodyDataConnectionStatistics {
	s.DeliverQps = &v
	return s
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatistics) SetProtocol(v string) *GetStatisticsByVhostResponseBodyDataConnectionStatistics {
	s.Protocol = &v
	return s
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatistics) SetPublishQps(v float32) *GetStatisticsByVhostResponseBodyDataConnectionStatistics {
	s.PublishQps = &v
	return s
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatistics) SetRemoteAddress(v string) *GetStatisticsByVhostResponseBodyDataConnectionStatistics {
	s.RemoteAddress = &v
	return s
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatistics) SetSecurityTransport(v bool) *GetStatisticsByVhostResponseBodyDataConnectionStatistics {
	s.SecurityTransport = &v
	return s
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatistics) SetState(v int32) *GetStatisticsByVhostResponseBodyDataConnectionStatistics {
	s.State = &v
	return s
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatistics) Validate() error {
	if s.ChannelStatisticsList != nil {
		if err := s.ChannelStatisticsList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsList struct {
	ChannelStatistics []*GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics `json:"ChannelStatistics,omitempty" xml:"ChannelStatistics,omitempty" type:"Repeated"`
}

func (s GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsList) String() string {
	return dara.Prettify(s)
}

func (s GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsList) GoString() string {
	return s.String()
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsList) GetChannelStatistics() []*GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics {
	return s.ChannelStatistics
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsList) SetChannelStatistics(v []*GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsList {
	s.ChannelStatistics = v
	return s
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsList) Validate() error {
	if s.ChannelStatistics != nil {
		for _, item := range s.ChannelStatistics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics struct {
	AckQps      *float32 `json:"AckQps,omitempty" xml:"AckQps,omitempty"`
	ConfirmQps  *float32 `json:"ConfirmQps,omitempty" xml:"ConfirmQps,omitempty"`
	DeliverQps  *float32 `json:"DeliverQps,omitempty" xml:"DeliverQps,omitempty"`
	GetQps      *float32 `json:"GetQps,omitempty" xml:"GetQps,omitempty"`
	Prefetch    *int32   `json:"Prefetch,omitempty" xml:"Prefetch,omitempty"`
	PublishQps  *float32 `json:"PublishQps,omitempty" xml:"PublishQps,omitempty"`
	State       *int32   `json:"State,omitempty" xml:"State,omitempty"`
	Unacked     *int32   `json:"Unacked,omitempty" xml:"Unacked,omitempty"`
	Unconfirmed *int32   `json:"Unconfirmed,omitempty" xml:"Unconfirmed,omitempty"`
}

func (s GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) GoString() string {
	return s.String()
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) GetAckQps() *float32 {
	return s.AckQps
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) GetConfirmQps() *float32 {
	return s.ConfirmQps
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) GetDeliverQps() *float32 {
	return s.DeliverQps
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) GetGetQps() *float32 {
	return s.GetQps
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) GetPrefetch() *int32 {
	return s.Prefetch
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) GetPublishQps() *float32 {
	return s.PublishQps
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) GetState() *int32 {
	return s.State
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) GetUnacked() *int32 {
	return s.Unacked
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) GetUnconfirmed() *int32 {
	return s.Unconfirmed
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) SetAckQps(v float32) *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics {
	s.AckQps = &v
	return s
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) SetConfirmQps(v float32) *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics {
	s.ConfirmQps = &v
	return s
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) SetDeliverQps(v float32) *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics {
	s.DeliverQps = &v
	return s
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) SetGetQps(v float32) *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics {
	s.GetQps = &v
	return s
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) SetPrefetch(v int32) *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics {
	s.Prefetch = &v
	return s
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) SetPublishQps(v float32) *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics {
	s.PublishQps = &v
	return s
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) SetState(v int32) *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics {
	s.State = &v
	return s
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) SetUnacked(v int32) *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics {
	s.Unacked = &v
	return s
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) SetUnconfirmed(v int32) *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics {
	s.Unconfirmed = &v
	return s
}

func (s *GetStatisticsByVhostResponseBodyDataConnectionStatisticsChannelStatisticsListChannelStatistics) Validate() error {
	return dara.Validate(s)
}
