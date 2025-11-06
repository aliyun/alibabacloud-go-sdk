// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstancesInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*QueryInstancesInfoResponseBodyData) *QueryInstancesInfoResponseBody
	GetData() []*QueryInstancesInfoResponseBodyData
	SetErrorCode(v string) *QueryInstancesInfoResponseBody
	GetErrorCode() *string
	SetHttpCode(v string) *QueryInstancesInfoResponseBody
	GetHttpCode() *string
	SetMessage(v string) *QueryInstancesInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryInstancesInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryInstancesInfoResponseBody
	GetSuccess() *bool
}

type QueryInstancesInfoResponseBody struct {
	// The details of the data.
	Data []*QueryInstancesInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 202
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The message that is returned.
	//
	// 	- If the request is successful, a success message is returned.
	//
	// 	- If the request fails, an error message is returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 54973C90-F379-4372-9AA5-053A3F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryInstancesInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryInstancesInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstancesInfoResponseBody) GetData() []*QueryInstancesInfoResponseBodyData {
	return s.Data
}

func (s *QueryInstancesInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryInstancesInfoResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *QueryInstancesInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryInstancesInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryInstancesInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryInstancesInfoResponseBody) SetData(v []*QueryInstancesInfoResponseBodyData) *QueryInstancesInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryInstancesInfoResponseBody) SetErrorCode(v string) *QueryInstancesInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryInstancesInfoResponseBody) SetHttpCode(v string) *QueryInstancesInfoResponseBody {
	s.HttpCode = &v
	return s
}

func (s *QueryInstancesInfoResponseBody) SetMessage(v string) *QueryInstancesInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryInstancesInfoResponseBody) SetRequestId(v string) *QueryInstancesInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInstancesInfoResponseBody) SetSuccess(v bool) *QueryInstancesInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryInstancesInfoResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryInstancesInfoResponseBodyData struct {
	// The enabled port.
	//
	// example:
	//
	// 8848
	ClientPort *string `json:"ClientPort,omitempty" xml:"ClientPort,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-12-15T02:02:15Z
	CreationTimestamp *string `json:"CreationTimestamp,omitempty" xml:"CreationTimestamp,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// null
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 120.55.71.x
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The IP address of the pod.
	//
	// example:
	//
	// 25.24.91.x
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The pod name.
	//
	// example:
	//
	// mse-xxxxx-xxxxx-reg-center-0-1
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// null
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The internal IP address.
	//
	// example:
	//
	// 172.16.66.x
	SingleTunnelVip *string `json:"SingleTunnelVip,omitempty" xml:"SingleTunnelVip,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-k
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
	// Indicates whether all pods in the cluster are distributed in the specified zones.
	//
	// example:
	//
	// true
	ZoneDistributed *bool `json:"ZoneDistributed,omitempty" xml:"ZoneDistributed,omitempty"`
}

func (s QueryInstancesInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryInstancesInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryInstancesInfoResponseBodyData) GetClientPort() *string {
	return s.ClientPort
}

func (s *QueryInstancesInfoResponseBodyData) GetCreationTimestamp() *string {
	return s.CreationTimestamp
}

func (s *QueryInstancesInfoResponseBodyData) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *QueryInstancesInfoResponseBodyData) GetInternetIp() *string {
	return s.InternetIp
}

func (s *QueryInstancesInfoResponseBodyData) GetIp() *string {
	return s.Ip
}

func (s *QueryInstancesInfoResponseBodyData) GetPodName() *string {
	return s.PodName
}

func (s *QueryInstancesInfoResponseBodyData) GetRole() *string {
	return s.Role
}

func (s *QueryInstancesInfoResponseBodyData) GetSingleTunnelVip() *string {
	return s.SingleTunnelVip
}

func (s *QueryInstancesInfoResponseBodyData) GetZone() *string {
	return s.Zone
}

func (s *QueryInstancesInfoResponseBodyData) GetZoneDistributed() *bool {
	return s.ZoneDistributed
}

func (s *QueryInstancesInfoResponseBodyData) SetClientPort(v string) *QueryInstancesInfoResponseBodyData {
	s.ClientPort = &v
	return s
}

func (s *QueryInstancesInfoResponseBodyData) SetCreationTimestamp(v string) *QueryInstancesInfoResponseBodyData {
	s.CreationTimestamp = &v
	return s
}

func (s *QueryInstancesInfoResponseBodyData) SetHealthStatus(v string) *QueryInstancesInfoResponseBodyData {
	s.HealthStatus = &v
	return s
}

func (s *QueryInstancesInfoResponseBodyData) SetInternetIp(v string) *QueryInstancesInfoResponseBodyData {
	s.InternetIp = &v
	return s
}

func (s *QueryInstancesInfoResponseBodyData) SetIp(v string) *QueryInstancesInfoResponseBodyData {
	s.Ip = &v
	return s
}

func (s *QueryInstancesInfoResponseBodyData) SetPodName(v string) *QueryInstancesInfoResponseBodyData {
	s.PodName = &v
	return s
}

func (s *QueryInstancesInfoResponseBodyData) SetRole(v string) *QueryInstancesInfoResponseBodyData {
	s.Role = &v
	return s
}

func (s *QueryInstancesInfoResponseBodyData) SetSingleTunnelVip(v string) *QueryInstancesInfoResponseBodyData {
	s.SingleTunnelVip = &v
	return s
}

func (s *QueryInstancesInfoResponseBodyData) SetZone(v string) *QueryInstancesInfoResponseBodyData {
	s.Zone = &v
	return s
}

func (s *QueryInstancesInfoResponseBodyData) SetZoneDistributed(v bool) *QueryInstancesInfoResponseBodyData {
	s.ZoneDistributed = &v
	return s
}

func (s *QueryInstancesInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
