// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIdcProbeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddIdcProbeFailedList(v []*AddIdcProbeResponseBodyAddIdcProbeFailedList) *AddIdcProbeResponseBody
	GetAddIdcProbeFailedList() []*AddIdcProbeResponseBodyAddIdcProbeFailedList
	SetCount(v string) *AddIdcProbeResponseBody
	GetCount() *string
	SetRequestId(v string) *AddIdcProbeResponseBody
	GetRequestId() *string
}

type AddIdcProbeResponseBody struct {
	// The records of failure.
	AddIdcProbeFailedList []*AddIdcProbeResponseBodyAddIdcProbeFailedList `json:"AddIdcProbeFailedList,omitempty" xml:"AddIdcProbeFailedList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D706F2DD-FF07-576B-9DD1-0B484A9B3065
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddIdcProbeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddIdcProbeResponseBody) GoString() string {
	return s.String()
}

func (s *AddIdcProbeResponseBody) GetAddIdcProbeFailedList() []*AddIdcProbeResponseBodyAddIdcProbeFailedList {
	return s.AddIdcProbeFailedList
}

func (s *AddIdcProbeResponseBody) GetCount() *string {
	return s.Count
}

func (s *AddIdcProbeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddIdcProbeResponseBody) SetAddIdcProbeFailedList(v []*AddIdcProbeResponseBodyAddIdcProbeFailedList) *AddIdcProbeResponseBody {
	s.AddIdcProbeFailedList = v
	return s
}

func (s *AddIdcProbeResponseBody) SetCount(v string) *AddIdcProbeResponseBody {
	s.Count = &v
	return s
}

func (s *AddIdcProbeResponseBody) SetRequestId(v string) *AddIdcProbeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddIdcProbeResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddIdcProbeResponseBodyAddIdcProbeFailedList struct {
	// The error message that is returned.
	//
	// example:
	//
	// The ResourceDirectoryId is invalid.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The name of the data center.
	//
	// example:
	//
	// test
	IdcName *string `json:"IdcName,omitempty" xml:"IdcName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// Hangzhou
	IdcRegion *string `json:"IdcRegion,omitempty" xml:"IdcRegion,omitempty"`
	// The ID of the server.
	//
	// example:
	//
	// sas-yqcl2ck3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// gl-sms-01
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 95.214.XXX.XXX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address.
	//
	// example:
	//
	// 172.29.XXX.XXX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The settings of the CIDR block.
	//
	// example:
	//
	// 192.168.XX.XX/24
	IpSegments *string `json:"IpSegments,omitempty" xml:"IpSegments,omitempty"`
	// The UUID of the server. Multiple UUIDs are separated by commas (,).
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUID.
	//
	// example:
	//
	// 076a446d-df7d-424c-bdc5-bb5dc7f1****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s AddIdcProbeResponseBodyAddIdcProbeFailedList) String() string {
	return dara.Prettify(s)
}

func (s AddIdcProbeResponseBodyAddIdcProbeFailedList) GoString() string {
	return s.String()
}

func (s *AddIdcProbeResponseBodyAddIdcProbeFailedList) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *AddIdcProbeResponseBodyAddIdcProbeFailedList) GetIdcName() *string {
	return s.IdcName
}

func (s *AddIdcProbeResponseBodyAddIdcProbeFailedList) GetIdcRegion() *string {
	return s.IdcRegion
}

func (s *AddIdcProbeResponseBodyAddIdcProbeFailedList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddIdcProbeResponseBodyAddIdcProbeFailedList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *AddIdcProbeResponseBodyAddIdcProbeFailedList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *AddIdcProbeResponseBodyAddIdcProbeFailedList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *AddIdcProbeResponseBodyAddIdcProbeFailedList) GetIpSegments() *string {
	return s.IpSegments
}

func (s *AddIdcProbeResponseBodyAddIdcProbeFailedList) GetUuid() *string {
	return s.Uuid
}

func (s *AddIdcProbeResponseBodyAddIdcProbeFailedList) SetErrorMsg(v string) *AddIdcProbeResponseBodyAddIdcProbeFailedList {
	s.ErrorMsg = &v
	return s
}

func (s *AddIdcProbeResponseBodyAddIdcProbeFailedList) SetIdcName(v string) *AddIdcProbeResponseBodyAddIdcProbeFailedList {
	s.IdcName = &v
	return s
}

func (s *AddIdcProbeResponseBodyAddIdcProbeFailedList) SetIdcRegion(v string) *AddIdcProbeResponseBodyAddIdcProbeFailedList {
	s.IdcRegion = &v
	return s
}

func (s *AddIdcProbeResponseBodyAddIdcProbeFailedList) SetInstanceId(v string) *AddIdcProbeResponseBodyAddIdcProbeFailedList {
	s.InstanceId = &v
	return s
}

func (s *AddIdcProbeResponseBodyAddIdcProbeFailedList) SetInstanceName(v string) *AddIdcProbeResponseBodyAddIdcProbeFailedList {
	s.InstanceName = &v
	return s
}

func (s *AddIdcProbeResponseBodyAddIdcProbeFailedList) SetInternetIp(v string) *AddIdcProbeResponseBodyAddIdcProbeFailedList {
	s.InternetIp = &v
	return s
}

func (s *AddIdcProbeResponseBodyAddIdcProbeFailedList) SetIntranetIp(v string) *AddIdcProbeResponseBodyAddIdcProbeFailedList {
	s.IntranetIp = &v
	return s
}

func (s *AddIdcProbeResponseBodyAddIdcProbeFailedList) SetIpSegments(v string) *AddIdcProbeResponseBodyAddIdcProbeFailedList {
	s.IpSegments = &v
	return s
}

func (s *AddIdcProbeResponseBodyAddIdcProbeFailedList) SetUuid(v string) *AddIdcProbeResponseBodyAddIdcProbeFailedList {
	s.Uuid = &v
	return s
}

func (s *AddIdcProbeResponseBodyAddIdcProbeFailedList) Validate() error {
	return dara.Validate(s)
}
