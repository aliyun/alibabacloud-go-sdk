// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterpriseAccelerateLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogs(v []*ListEnterpriseAccelerateLogsResponseBodyLogs) *ListEnterpriseAccelerateLogsResponseBody
	GetLogs() []*ListEnterpriseAccelerateLogsResponseBodyLogs
	SetRequestId(v string) *ListEnterpriseAccelerateLogsResponseBody
	GetRequestId() *string
	SetTotalNumber(v int32) *ListEnterpriseAccelerateLogsResponseBody
	GetTotalNumber() *int32
}

type ListEnterpriseAccelerateLogsResponseBody struct {
	Logs []*ListEnterpriseAccelerateLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// example:
	//
	// 43F07A6A-294D-56FB-85EB-6AD00C5B60FF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 120
	TotalNumber *int32 `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s ListEnterpriseAccelerateLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseAccelerateLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnterpriseAccelerateLogsResponseBody) GetLogs() []*ListEnterpriseAccelerateLogsResponseBodyLogs {
	return s.Logs
}

func (s *ListEnterpriseAccelerateLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnterpriseAccelerateLogsResponseBody) GetTotalNumber() *int32 {
	return s.TotalNumber
}

func (s *ListEnterpriseAccelerateLogsResponseBody) SetLogs(v []*ListEnterpriseAccelerateLogsResponseBodyLogs) *ListEnterpriseAccelerateLogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListEnterpriseAccelerateLogsResponseBody) SetRequestId(v string) *ListEnterpriseAccelerateLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnterpriseAccelerateLogsResponseBody) SetTotalNumber(v int32) *ListEnterpriseAccelerateLogsResponseBody {
	s.TotalNumber = &v
	return s
}

func (s *ListEnterpriseAccelerateLogsResponseBody) Validate() error {
	if s.Logs != nil {
		for _, item := range s.Logs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEnterpriseAccelerateLogsResponseBodyLogs struct {
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// example:
	//
	// windows
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// www.bing.com:443
	DstAddr *string `json:"DstAddr,omitempty" xml:"DstAddr,omitempty"`
	// example:
	//
	// 12299
	InBytes *string `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// example:
	//
	// 2603
	OutBytes   *string `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// example:
	//
	// 8.222.179.xxx:10015
	ProxyAddr *string `json:"ProxyAddr,omitempty" xml:"ProxyAddr,omitempty"`
	// example:
	//
	// 1748422797
	UnixTime *string `json:"UnixTime,omitempty" xml:"UnixTime,omitempty"`
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListEnterpriseAccelerateLogsResponseBodyLogs) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseAccelerateLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *ListEnterpriseAccelerateLogsResponseBodyLogs) GetDepartment() *string {
	return s.Department
}

func (s *ListEnterpriseAccelerateLogsResponseBodyLogs) GetDeviceType() *string {
	return s.DeviceType
}

func (s *ListEnterpriseAccelerateLogsResponseBodyLogs) GetDstAddr() *string {
	return s.DstAddr
}

func (s *ListEnterpriseAccelerateLogsResponseBodyLogs) GetInBytes() *string {
	return s.InBytes
}

func (s *ListEnterpriseAccelerateLogsResponseBodyLogs) GetOutBytes() *string {
	return s.OutBytes
}

func (s *ListEnterpriseAccelerateLogsResponseBodyLogs) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListEnterpriseAccelerateLogsResponseBodyLogs) GetProxyAddr() *string {
	return s.ProxyAddr
}

func (s *ListEnterpriseAccelerateLogsResponseBodyLogs) GetUnixTime() *string {
	return s.UnixTime
}

func (s *ListEnterpriseAccelerateLogsResponseBodyLogs) GetUsername() *string {
	return s.Username
}

func (s *ListEnterpriseAccelerateLogsResponseBodyLogs) SetDepartment(v string) *ListEnterpriseAccelerateLogsResponseBodyLogs {
	s.Department = &v
	return s
}

func (s *ListEnterpriseAccelerateLogsResponseBodyLogs) SetDeviceType(v string) *ListEnterpriseAccelerateLogsResponseBodyLogs {
	s.DeviceType = &v
	return s
}

func (s *ListEnterpriseAccelerateLogsResponseBodyLogs) SetDstAddr(v string) *ListEnterpriseAccelerateLogsResponseBodyLogs {
	s.DstAddr = &v
	return s
}

func (s *ListEnterpriseAccelerateLogsResponseBodyLogs) SetInBytes(v string) *ListEnterpriseAccelerateLogsResponseBodyLogs {
	s.InBytes = &v
	return s
}

func (s *ListEnterpriseAccelerateLogsResponseBodyLogs) SetOutBytes(v string) *ListEnterpriseAccelerateLogsResponseBodyLogs {
	s.OutBytes = &v
	return s
}

func (s *ListEnterpriseAccelerateLogsResponseBodyLogs) SetPolicyName(v string) *ListEnterpriseAccelerateLogsResponseBodyLogs {
	s.PolicyName = &v
	return s
}

func (s *ListEnterpriseAccelerateLogsResponseBodyLogs) SetProxyAddr(v string) *ListEnterpriseAccelerateLogsResponseBodyLogs {
	s.ProxyAddr = &v
	return s
}

func (s *ListEnterpriseAccelerateLogsResponseBodyLogs) SetUnixTime(v string) *ListEnterpriseAccelerateLogsResponseBodyLogs {
	s.UnixTime = &v
	return s
}

func (s *ListEnterpriseAccelerateLogsResponseBodyLogs) SetUsername(v string) *ListEnterpriseAccelerateLogsResponseBodyLogs {
	s.Username = &v
	return s
}

func (s *ListEnterpriseAccelerateLogsResponseBodyLogs) Validate() error {
	return dara.Validate(s)
}
