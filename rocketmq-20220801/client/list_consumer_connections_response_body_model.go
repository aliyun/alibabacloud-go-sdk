// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerConnectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListConsumerConnectionsResponseBody
	GetCode() *string
	SetData(v *ListConsumerConnectionsResponseBodyData) *ListConsumerConnectionsResponseBody
	GetData() *ListConsumerConnectionsResponseBodyData
	SetDynamicCode(v string) *ListConsumerConnectionsResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListConsumerConnectionsResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListConsumerConnectionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListConsumerConnectionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListConsumerConnectionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListConsumerConnectionsResponseBody
	GetSuccess() *bool
}

type ListConsumerConnectionsResponseBody struct {
	// The returned error code.
	//
	// example:
	//
	// MissingPageNumber
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *ListConsumerConnectionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A3620115-6F1F-5CFB-AA3F-BBD4853B2EC4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListConsumerConnectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConsumerConnectionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListConsumerConnectionsResponseBody) GetData() *ListConsumerConnectionsResponseBodyData {
	return s.Data
}

func (s *ListConsumerConnectionsResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListConsumerConnectionsResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListConsumerConnectionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListConsumerConnectionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListConsumerConnectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConsumerConnectionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListConsumerConnectionsResponseBody) SetCode(v string) *ListConsumerConnectionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListConsumerConnectionsResponseBody) SetData(v *ListConsumerConnectionsResponseBodyData) *ListConsumerConnectionsResponseBody {
	s.Data = v
	return s
}

func (s *ListConsumerConnectionsResponseBody) SetDynamicCode(v string) *ListConsumerConnectionsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListConsumerConnectionsResponseBody) SetDynamicMessage(v string) *ListConsumerConnectionsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListConsumerConnectionsResponseBody) SetHttpStatusCode(v int32) *ListConsumerConnectionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListConsumerConnectionsResponseBody) SetMessage(v string) *ListConsumerConnectionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListConsumerConnectionsResponseBody) SetRequestId(v string) *ListConsumerConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConsumerConnectionsResponseBody) SetSuccess(v bool) *ListConsumerConnectionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListConsumerConnectionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListConsumerConnectionsResponseBodyData struct {
	// The client connection list
	Connections []*ListConsumerConnectionsResponseBodyDataConnections `json:"connections,omitempty" xml:"connections,omitempty" type:"Repeated"`
	// The consumer group ID.
	//
	// example:
	//
	// CID-TEST
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListConsumerConnectionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerConnectionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConsumerConnectionsResponseBodyData) GetConnections() []*ListConsumerConnectionsResponseBodyDataConnections {
	return s.Connections
}

func (s *ListConsumerConnectionsResponseBodyData) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *ListConsumerConnectionsResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListConsumerConnectionsResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListConsumerConnectionsResponseBodyData) SetConnections(v []*ListConsumerConnectionsResponseBodyDataConnections) *ListConsumerConnectionsResponseBodyData {
	s.Connections = v
	return s
}

func (s *ListConsumerConnectionsResponseBodyData) SetConsumerGroupId(v string) *ListConsumerConnectionsResponseBodyData {
	s.ConsumerGroupId = &v
	return s
}

func (s *ListConsumerConnectionsResponseBodyData) SetInstanceId(v string) *ListConsumerConnectionsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListConsumerConnectionsResponseBodyData) SetRegionId(v string) *ListConsumerConnectionsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListConsumerConnectionsResponseBodyData) Validate() error {
	if s.Connections != nil {
		for _, item := range s.Connections {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListConsumerConnectionsResponseBodyDataConnections struct {
	// The ID of the client.
	//
	// example:
	//
	// 172.17.135.197@17392#1936705963#551717232#9873695589062458
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// Host IP/Public IP
	//
	// example:
	//
	// xx.xx.xx.xx
	EgressIp *string `json:"egressIp,omitempty" xml:"egressIp,omitempty"`
	// The `hostname` of the cloud-native box.
	//
	// example:
	//
	// vos
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
	// The language of the client.
	//
	// example:
	//
	// java
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// Consumption Mode
	//
	// - BROADCASTING
	//
	// - CLUSTERING
	//
	// example:
	//
	// BROADCASTING
	MessageModel *string `json:"messageModel,omitempty" xml:"messageModel,omitempty"`
	// The version of the client.
	//
	// example:
	//
	// 1.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListConsumerConnectionsResponseBodyDataConnections) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerConnectionsResponseBodyDataConnections) GoString() string {
	return s.String()
}

func (s *ListConsumerConnectionsResponseBodyDataConnections) GetClientId() *string {
	return s.ClientId
}

func (s *ListConsumerConnectionsResponseBodyDataConnections) GetEgressIp() *string {
	return s.EgressIp
}

func (s *ListConsumerConnectionsResponseBodyDataConnections) GetHostname() *string {
	return s.Hostname
}

func (s *ListConsumerConnectionsResponseBodyDataConnections) GetLanguage() *string {
	return s.Language
}

func (s *ListConsumerConnectionsResponseBodyDataConnections) GetMessageModel() *string {
	return s.MessageModel
}

func (s *ListConsumerConnectionsResponseBodyDataConnections) GetVersion() *string {
	return s.Version
}

func (s *ListConsumerConnectionsResponseBodyDataConnections) SetClientId(v string) *ListConsumerConnectionsResponseBodyDataConnections {
	s.ClientId = &v
	return s
}

func (s *ListConsumerConnectionsResponseBodyDataConnections) SetEgressIp(v string) *ListConsumerConnectionsResponseBodyDataConnections {
	s.EgressIp = &v
	return s
}

func (s *ListConsumerConnectionsResponseBodyDataConnections) SetHostname(v string) *ListConsumerConnectionsResponseBodyDataConnections {
	s.Hostname = &v
	return s
}

func (s *ListConsumerConnectionsResponseBodyDataConnections) SetLanguage(v string) *ListConsumerConnectionsResponseBodyDataConnections {
	s.Language = &v
	return s
}

func (s *ListConsumerConnectionsResponseBodyDataConnections) SetMessageModel(v string) *ListConsumerConnectionsResponseBodyDataConnections {
	s.MessageModel = &v
	return s
}

func (s *ListConsumerConnectionsResponseBodyDataConnections) SetVersion(v string) *ListConsumerConnectionsResponseBodyDataConnections {
	s.Version = &v
	return s
}

func (s *ListConsumerConnectionsResponseBodyDataConnections) Validate() error {
	return dara.Validate(s)
}
