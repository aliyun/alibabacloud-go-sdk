// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsConsumerStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OnsConsumerStatusResponseBodyData) *OnsConsumerStatusResponseBody
	GetData() *OnsConsumerStatusResponseBodyData
	SetRequestId(v string) *OnsConsumerStatusResponseBody
	GetRequestId() *string
}

type OnsConsumerStatusResponseBody struct {
	// The data returned.
	Data *OnsConsumerStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 10EDC518-10E7-4B34-92FB-171235FA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnsConsumerStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerStatusResponseBody) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBody) GetData() *OnsConsumerStatusResponseBodyData {
	return s.Data
}

func (s *OnsConsumerStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnsConsumerStatusResponseBody) SetData(v *OnsConsumerStatusResponseBodyData) *OnsConsumerStatusResponseBody {
	s.Data = v
	return s
}

func (s *OnsConsumerStatusResponseBody) SetRequestId(v string) *OnsConsumerStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnsConsumerStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsConsumerStatusResponseBodyData struct {
	ConnectionSet *OnsConsumerStatusResponseBodyDataConnectionSet `json:"ConnectionSet,omitempty" xml:"ConnectionSet,omitempty" type:"Struct"`
	// The consumption mode. Valid values:
	//
	// 	- **CLUSTERING**: the clustering consumption mode
	//
	// 	- **BROADCASTING**: the broadcasting consumption mode
	//
	// For more information about consumption modes, see [Clustering consumption and broadcasting consumption](https://help.aliyun.com/document_detail/43163.html).
	//
	// example:
	//
	// CLUSTERING
	ConsumeModel *string `json:"ConsumeModel,omitempty" xml:"ConsumeModel,omitempty"`
	// The TPS for message consumption.
	//
	// example:
	//
	// 0
	ConsumeTps                 *float32                                                     `json:"ConsumeTps,omitempty" xml:"ConsumeTps,omitempty"`
	ConsumerConnectionInfoList *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList `json:"ConsumerConnectionInfoList,omitempty" xml:"ConsumerConnectionInfoList,omitempty" type:"Struct"`
	// The maximum latency of message consumption in all topics to which the consumer group subscribes. Unit: milliseconds.
	//
	// example:
	//
	// 100857
	DelayTime         *int64                                              `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	DetailInTopicList *OnsConsumerStatusResponseBodyDataDetailInTopicList `json:"DetailInTopicList,omitempty" xml:"DetailInTopicList,omitempty" type:"Struct"`
	// The ID of the instance
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The most recent point in time when a message was consumed.
	//
	// The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1566883844954
	LastTimestamp *int64 `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// Indicates whether the consumer group is online.
	//
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// Indicates whether load balancing is performed as expected. Valid values:
	//
	// 	- **true**: Load balancing is performed as expected.
	//
	// 	- **false**: Load balancing is not performed as expected.
	//
	// example:
	//
	// true
	RebalanceOK *bool `json:"RebalanceOK,omitempty" xml:"RebalanceOK,omitempty"`
	// Indicates whether all consumers in the consumer group subscribe to the same topics and tags.
	//
	// example:
	//
	// true
	SubscriptionSame *bool `json:"SubscriptionSame,omitempty" xml:"SubscriptionSame,omitempty"`
	// The total number of accumulated messages.
	//
	// example:
	//
	// 197
	TotalDiff *int64 `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
}

func (s OnsConsumerStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyData) GetConnectionSet() *OnsConsumerStatusResponseBodyDataConnectionSet {
	return s.ConnectionSet
}

func (s *OnsConsumerStatusResponseBodyData) GetConsumeModel() *string {
	return s.ConsumeModel
}

func (s *OnsConsumerStatusResponseBodyData) GetConsumeTps() *float32 {
	return s.ConsumeTps
}

func (s *OnsConsumerStatusResponseBodyData) GetConsumerConnectionInfoList() *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	return s.ConsumerConnectionInfoList
}

func (s *OnsConsumerStatusResponseBodyData) GetDelayTime() *int64 {
	return s.DelayTime
}

func (s *OnsConsumerStatusResponseBodyData) GetDetailInTopicList() *OnsConsumerStatusResponseBodyDataDetailInTopicList {
	return s.DetailInTopicList
}

func (s *OnsConsumerStatusResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsConsumerStatusResponseBodyData) GetLastTimestamp() *int64 {
	return s.LastTimestamp
}

func (s *OnsConsumerStatusResponseBodyData) GetOnline() *bool {
	return s.Online
}

func (s *OnsConsumerStatusResponseBodyData) GetRebalanceOK() *bool {
	return s.RebalanceOK
}

func (s *OnsConsumerStatusResponseBodyData) GetSubscriptionSame() *bool {
	return s.SubscriptionSame
}

func (s *OnsConsumerStatusResponseBodyData) GetTotalDiff() *int64 {
	return s.TotalDiff
}

func (s *OnsConsumerStatusResponseBodyData) SetConnectionSet(v *OnsConsumerStatusResponseBodyDataConnectionSet) *OnsConsumerStatusResponseBodyData {
	s.ConnectionSet = v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetConsumeModel(v string) *OnsConsumerStatusResponseBodyData {
	s.ConsumeModel = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetConsumeTps(v float32) *OnsConsumerStatusResponseBodyData {
	s.ConsumeTps = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetConsumerConnectionInfoList(v *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) *OnsConsumerStatusResponseBodyData {
	s.ConsumerConnectionInfoList = v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetDelayTime(v int64) *OnsConsumerStatusResponseBodyData {
	s.DelayTime = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetDetailInTopicList(v *OnsConsumerStatusResponseBodyDataDetailInTopicList) *OnsConsumerStatusResponseBodyData {
	s.DetailInTopicList = v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetInstanceId(v string) *OnsConsumerStatusResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetLastTimestamp(v int64) *OnsConsumerStatusResponseBodyData {
	s.LastTimestamp = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetOnline(v bool) *OnsConsumerStatusResponseBodyData {
	s.Online = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetRebalanceOK(v bool) *OnsConsumerStatusResponseBodyData {
	s.RebalanceOK = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetSubscriptionSame(v bool) *OnsConsumerStatusResponseBodyData {
	s.SubscriptionSame = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) SetTotalDiff(v int64) *OnsConsumerStatusResponseBodyData {
	s.TotalDiff = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyData) Validate() error {
	if s.ConnectionSet != nil {
		if err := s.ConnectionSet.Validate(); err != nil {
			return err
		}
	}
	if s.ConsumerConnectionInfoList != nil {
		if err := s.ConsumerConnectionInfoList.Validate(); err != nil {
			return err
		}
	}
	if s.DetailInTopicList != nil {
		if err := s.DetailInTopicList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsConsumerStatusResponseBodyDataConnectionSet struct {
	ConnectionDo []*OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo `json:"ConnectionDo,omitempty" xml:"ConnectionDo,omitempty" type:"Repeated"`
}

func (s OnsConsumerStatusResponseBodyDataConnectionSet) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConnectionSet) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSet) GetConnectionDo() []*OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo {
	return s.ConnectionDo
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSet) SetConnectionDo(v []*OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) *OnsConsumerStatusResponseBodyDataConnectionSet {
	s.ConnectionDo = v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSet) Validate() error {
	if s.ConnectionDo != nil {
		for _, item := range s.ConnectionDo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo struct {
	ClientAddr *string `json:"ClientAddr,omitempty" xml:"ClientAddr,omitempty"`
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// 100
	Diff     *int64  `json:"Diff,omitempty" xml:"Diff,omitempty"`
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	RemoteIP *string `json:"RemoteIP,omitempty" xml:"RemoteIP,omitempty"`
	Version  *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) GetClientAddr() *string {
	return s.ClientAddr
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) GetClientId() *string {
	return s.ClientId
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) GetDiff() *int64 {
	return s.Diff
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) GetLanguage() *string {
	return s.Language
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) GetRemoteIP() *string {
	return s.RemoteIP
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) GetVersion() *string {
	return s.Version
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) SetClientAddr(v string) *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo {
	s.ClientAddr = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) SetClientId(v string) *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo {
	s.ClientId = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) SetDiff(v int64) *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo {
	s.Diff = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) SetLanguage(v string) *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo {
	s.Language = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) SetRemoteIP(v string) *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo {
	s.RemoteIP = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) SetVersion(v string) *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo {
	s.Version = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConnectionSetConnectionDo) Validate() error {
	return dara.Validate(s)
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList struct {
	ConsumerConnectionInfoDo []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo `json:"ConsumerConnectionInfoDo,omitempty" xml:"ConsumerConnectionInfoDo,omitempty" type:"Repeated"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) GetConsumerConnectionInfoDo() []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	return s.ConsumerConnectionInfoDo
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) SetConsumerConnectionInfoDo(v []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList {
	s.ConsumerConnectionInfoDo = v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoList) Validate() error {
	if s.ConsumerConnectionInfoDo != nil {
		for _, item := range s.ConsumerConnectionInfoDo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo struct {
	ClientId        *string                                                                                             `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	Connection      *string                                                                                             `json:"Connection,omitempty" xml:"Connection,omitempty"`
	ConsumeModel    *string                                                                                             `json:"ConsumeModel,omitempty" xml:"ConsumeModel,omitempty"`
	ConsumeType     *string                                                                                             `json:"ConsumeType,omitempty" xml:"ConsumeType,omitempty"`
	Jstack          *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstack          `json:"Jstack,omitempty" xml:"Jstack,omitempty" type:"Struct"`
	Language        *string                                                                                             `json:"Language,omitempty" xml:"Language,omitempty"`
	LastTimeStamp   *int64                                                                                              `json:"LastTimeStamp,omitempty" xml:"LastTimeStamp,omitempty"`
	RunningDataList *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataList `json:"RunningDataList,omitempty" xml:"RunningDataList,omitempty" type:"Struct"`
	StartTimeStamp  *int64                                                                                              `json:"StartTimeStamp,omitempty" xml:"StartTimeStamp,omitempty"`
	SubscriptionSet *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSet `json:"SubscriptionSet,omitempty" xml:"SubscriptionSet,omitempty" type:"Struct"`
	ThreadCount     *int32                                                                                              `json:"ThreadCount,omitempty" xml:"ThreadCount,omitempty"`
	Version         *string                                                                                             `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) GetClientId() *string {
	return s.ClientId
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) GetConnection() *string {
	return s.Connection
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) GetConsumeModel() *string {
	return s.ConsumeModel
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) GetConsumeType() *string {
	return s.ConsumeType
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) GetJstack() *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstack {
	return s.Jstack
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) GetLanguage() *string {
	return s.Language
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) GetLastTimeStamp() *int64 {
	return s.LastTimeStamp
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) GetRunningDataList() *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataList {
	return s.RunningDataList
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) GetStartTimeStamp() *int64 {
	return s.StartTimeStamp
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) GetSubscriptionSet() *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSet {
	return s.SubscriptionSet
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) GetThreadCount() *int32 {
	return s.ThreadCount
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) GetVersion() *string {
	return s.Version
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetClientId(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.ClientId = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetConnection(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.Connection = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetConsumeModel(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.ConsumeModel = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetConsumeType(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.ConsumeType = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetJstack(v *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstack) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.Jstack = v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetLanguage(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.Language = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetLastTimeStamp(v int64) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.LastTimeStamp = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetRunningDataList(v *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataList) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.RunningDataList = v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetStartTimeStamp(v int64) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.StartTimeStamp = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetSubscriptionSet(v *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSet) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.SubscriptionSet = v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetThreadCount(v int32) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.ThreadCount = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) SetVersion(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo {
	s.Version = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDo) Validate() error {
	if s.Jstack != nil {
		if err := s.Jstack.Validate(); err != nil {
			return err
		}
	}
	if s.RunningDataList != nil {
		if err := s.RunningDataList.Validate(); err != nil {
			return err
		}
	}
	if s.SubscriptionSet != nil {
		if err := s.SubscriptionSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstack struct {
	ThreadTrackDo []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo `json:"ThreadTrackDo,omitempty" xml:"ThreadTrackDo,omitempty" type:"Repeated"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstack) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstack) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstack) GetThreadTrackDo() []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo {
	return s.ThreadTrackDo
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstack) SetThreadTrackDo(v []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstack {
	s.ThreadTrackDo = v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstack) Validate() error {
	if s.ThreadTrackDo != nil {
		for _, item := range s.ThreadTrackDo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo struct {
	Thread    *string                                                                                                          `json:"Thread,omitempty" xml:"Thread,omitempty"`
	TrackList *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDoTrackList `json:"TrackList,omitempty" xml:"TrackList,omitempty" type:"Struct"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo) GetThread() *string {
	return s.Thread
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo) GetTrackList() *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDoTrackList {
	return s.TrackList
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo) SetThread(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo {
	s.Thread = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo) SetTrackList(v *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDoTrackList) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo {
	s.TrackList = v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDo) Validate() error {
	if s.TrackList != nil {
		if err := s.TrackList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDoTrackList struct {
	Track []*string `json:"Track,omitempty" xml:"Track,omitempty" type:"Repeated"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDoTrackList) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDoTrackList) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDoTrackList) GetTrack() []*string {
	return s.Track
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDoTrackList) SetTrack(v []*string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDoTrackList {
	s.Track = v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoJstackThreadTrackDoTrackList) Validate() error {
	return dara.Validate(s)
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataList struct {
	ConsumerRunningDataDo []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo `json:"ConsumerRunningDataDo,omitempty" xml:"ConsumerRunningDataDo,omitempty" type:"Repeated"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataList) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataList) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataList) GetConsumerRunningDataDo() []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo {
	return s.ConsumerRunningDataDo
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataList) SetConsumerRunningDataDo(v []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataList {
	s.ConsumerRunningDataDo = v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataList) Validate() error {
	if s.ConsumerRunningDataDo != nil {
		for _, item := range s.ConsumerRunningDataDo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo struct {
	FailedCountPerHour *int64   `json:"FailedCountPerHour,omitempty" xml:"FailedCountPerHour,omitempty"`
	FailedTps          *float32 `json:"FailedTps,omitempty" xml:"FailedTps,omitempty"`
	OkTps              *float32 `json:"OkTps,omitempty" xml:"OkTps,omitempty"`
	Rt                 *float32 `json:"Rt,omitempty" xml:"Rt,omitempty"`
	Topic              *string  `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) GetFailedCountPerHour() *int64 {
	return s.FailedCountPerHour
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) GetFailedTps() *float32 {
	return s.FailedTps
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) GetOkTps() *float32 {
	return s.OkTps
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) GetRt() *float32 {
	return s.Rt
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) GetTopic() *string {
	return s.Topic
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) SetFailedCountPerHour(v int64) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo {
	s.FailedCountPerHour = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) SetFailedTps(v float32) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo {
	s.FailedTps = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) SetOkTps(v float32) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo {
	s.OkTps = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) SetRt(v float32) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo {
	s.Rt = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) SetTopic(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo {
	s.Topic = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoRunningDataListConsumerRunningDataDo) Validate() error {
	return dara.Validate(s)
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSet struct {
	SubscriptionData []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData `json:"SubscriptionData,omitempty" xml:"SubscriptionData,omitempty" type:"Repeated"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSet) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSet) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSet) GetSubscriptionData() []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData {
	return s.SubscriptionData
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSet) SetSubscriptionData(v []*OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSet {
	s.SubscriptionData = v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSet) Validate() error {
	if s.SubscriptionData != nil {
		for _, item := range s.SubscriptionData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData struct {
	SubString  *string                                                                                                                    `json:"SubString,omitempty" xml:"SubString,omitempty"`
	SubVersion *int64                                                                                                                     `json:"SubVersion,omitempty" xml:"SubVersion,omitempty"`
	TagsSet    *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionDataTagsSet `json:"TagsSet,omitempty" xml:"TagsSet,omitempty" type:"Struct"`
	Topic      *string                                                                                                                    `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData) GetSubString() *string {
	return s.SubString
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData) GetSubVersion() *int64 {
	return s.SubVersion
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData) GetTagsSet() *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionDataTagsSet {
	return s.TagsSet
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData) GetTopic() *string {
	return s.Topic
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData) SetSubString(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData {
	s.SubString = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData) SetSubVersion(v int64) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData {
	s.SubVersion = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData) SetTagsSet(v *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionDataTagsSet) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData {
	s.TagsSet = v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData) SetTopic(v string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData {
	s.Topic = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionData) Validate() error {
	if s.TagsSet != nil {
		if err := s.TagsSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionDataTagsSet struct {
	Tag []*string `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionDataTagsSet) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionDataTagsSet) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionDataTagsSet) GetTag() []*string {
	return s.Tag
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionDataTagsSet) SetTag(v []*string) *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionDataTagsSet {
	s.Tag = v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataConsumerConnectionInfoListConsumerConnectionInfoDoSubscriptionSetSubscriptionDataTagsSet) Validate() error {
	return dara.Validate(s)
}

type OnsConsumerStatusResponseBodyDataDetailInTopicList struct {
	DetailInTopicDo []*OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo `json:"DetailInTopicDo,omitempty" xml:"DetailInTopicDo,omitempty" type:"Repeated"`
}

func (s OnsConsumerStatusResponseBodyDataDetailInTopicList) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataDetailInTopicList) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataDetailInTopicList) GetDetailInTopicDo() []*OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo {
	return s.DetailInTopicDo
}

func (s *OnsConsumerStatusResponseBodyDataDetailInTopicList) SetDetailInTopicDo(v []*OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo) *OnsConsumerStatusResponseBodyDataDetailInTopicList {
	s.DetailInTopicDo = v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataDetailInTopicList) Validate() error {
	if s.DetailInTopicDo != nil {
		for _, item := range s.DetailInTopicDo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo struct {
	DelayTime     *int64  `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	LastTimestamp *int64  `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Topic         *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	TotalDiff     *int64  `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
}

func (s OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo) String() string {
	return dara.Prettify(s)
}

func (s OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo) GoString() string {
	return s.String()
}

func (s *OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo) GetDelayTime() *int64 {
	return s.DelayTime
}

func (s *OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo) GetLastTimestamp() *int64 {
	return s.LastTimestamp
}

func (s *OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo) GetTopic() *string {
	return s.Topic
}

func (s *OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo) GetTotalDiff() *int64 {
	return s.TotalDiff
}

func (s *OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo) SetDelayTime(v int64) *OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo {
	s.DelayTime = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo) SetLastTimestamp(v int64) *OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo {
	s.LastTimestamp = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo) SetTopic(v string) *OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo {
	s.Topic = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo) SetTotalDiff(v int64) *OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo {
	s.TotalDiff = &v
	return s
}

func (s *OnsConsumerStatusResponseBodyDataDetailInTopicListDetailInTopicDo) Validate() error {
	return dara.Validate(s)
}
