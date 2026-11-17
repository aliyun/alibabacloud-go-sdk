// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvadeEventPayloadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetsInstanceId(v string) *DescribeInvadeEventPayloadRequest
	GetAssetsInstanceId() *string
	SetEndTime(v string) *DescribeInvadeEventPayloadRequest
	GetEndTime() *string
	SetEventUuid(v string) *DescribeInvadeEventPayloadRequest
	GetEventUuid() *string
	SetLang(v string) *DescribeInvadeEventPayloadRequest
	GetLang() *string
	SetPrivateIP(v string) *DescribeInvadeEventPayloadRequest
	GetPrivateIP() *string
	SetPublicIP(v string) *DescribeInvadeEventPayloadRequest
	GetPublicIP() *string
	SetSourceIp(v string) *DescribeInvadeEventPayloadRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeInvadeEventPayloadRequest
	GetStartTime() *string
}

type DescribeInvadeEventPayloadRequest struct {
	// The instance ID of the asset where the intrusion event occurred.
	//
	// example:
	//
	// i-uf6faknmuby7ezht****
	AssetsInstanceId *string `json:"AssetsInstanceId,omitempty" xml:"AssetsInstanceId,omitempty"`
	// The end time of the time range during which the event occurred. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1731551104
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 92498578-7c42-4845-8c73-7e824782****
	EventUuid *string `json:"EventUuid,omitempty" xml:"EventUuid,omitempty"`
	// The language of the request and response. Valid values: - **zh**: Chinese. - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The private IP address of the asset where the intrusion event occurred.
	//
	// example:
	//
	// 10.200.33.XXX
	PrivateIP *string `json:"PrivateIP,omitempty" xml:"PrivateIP,omitempty"`
	// The public IP address of the asset where the intrusion event occurred.
	//
	// example:
	//
	// 192.0.XX.XX
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// The source IP address of the caller.
	//
	// example:
	//
	// 58.34.174.194
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The start time of the time range during which the event occurred. The value is a UNIX timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1656837360
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInvadeEventPayloadRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEventPayloadRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventPayloadRequest) GetAssetsInstanceId() *string {
	return s.AssetsInstanceId
}

func (s *DescribeInvadeEventPayloadRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInvadeEventPayloadRequest) GetEventUuid() *string {
	return s.EventUuid
}

func (s *DescribeInvadeEventPayloadRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInvadeEventPayloadRequest) GetPrivateIP() *string {
	return s.PrivateIP
}

func (s *DescribeInvadeEventPayloadRequest) GetPublicIP() *string {
	return s.PublicIP
}

func (s *DescribeInvadeEventPayloadRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInvadeEventPayloadRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInvadeEventPayloadRequest) SetAssetsInstanceId(v string) *DescribeInvadeEventPayloadRequest {
	s.AssetsInstanceId = &v
	return s
}

func (s *DescribeInvadeEventPayloadRequest) SetEndTime(v string) *DescribeInvadeEventPayloadRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInvadeEventPayloadRequest) SetEventUuid(v string) *DescribeInvadeEventPayloadRequest {
	s.EventUuid = &v
	return s
}

func (s *DescribeInvadeEventPayloadRequest) SetLang(v string) *DescribeInvadeEventPayloadRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInvadeEventPayloadRequest) SetPrivateIP(v string) *DescribeInvadeEventPayloadRequest {
	s.PrivateIP = &v
	return s
}

func (s *DescribeInvadeEventPayloadRequest) SetPublicIP(v string) *DescribeInvadeEventPayloadRequest {
	s.PublicIP = &v
	return s
}

func (s *DescribeInvadeEventPayloadRequest) SetSourceIp(v string) *DescribeInvadeEventPayloadRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInvadeEventPayloadRequest) SetStartTime(v string) *DescribeInvadeEventPayloadRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInvadeEventPayloadRequest) Validate() error {
	return dara.Validate(s)
}
