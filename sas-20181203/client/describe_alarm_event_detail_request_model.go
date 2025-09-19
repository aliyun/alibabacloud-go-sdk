// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlarmEventDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmUniqueInfo(v string) *DescribeAlarmEventDetailRequest
	GetAlarmUniqueInfo() *string
	SetFrom(v string) *DescribeAlarmEventDetailRequest
	GetFrom() *string
	SetLang(v string) *DescribeAlarmEventDetailRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeAlarmEventDetailRequest
	GetSourceIp() *string
}

type DescribeAlarmEventDetailRequest struct {
	// The unique identifier of the alert event.
	//
	// > To query the details of an alert event, you must provide the unique identifier of the alert event. You can call the [DescribeSuspEvents](~~DescribeSuspEvents~~) operation to obtain the identifier.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9f62555666f177aa84ee1eaf465a****
	AlarmUniqueInfo *string `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	// The ID of the request source. Set the value to **sas**.
	//
	// This parameter is required.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeAlarmEventDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlarmEventDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlarmEventDetailRequest) GetAlarmUniqueInfo() *string {
	return s.AlarmUniqueInfo
}

func (s *DescribeAlarmEventDetailRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeAlarmEventDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAlarmEventDetailRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeAlarmEventDetailRequest) SetAlarmUniqueInfo(v string) *DescribeAlarmEventDetailRequest {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *DescribeAlarmEventDetailRequest) SetFrom(v string) *DescribeAlarmEventDetailRequest {
	s.From = &v
	return s
}

func (s *DescribeAlarmEventDetailRequest) SetLang(v string) *DescribeAlarmEventDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAlarmEventDetailRequest) SetSourceIp(v string) *DescribeAlarmEventDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAlarmEventDetailRequest) Validate() error {
	return dara.Validate(s)
}
