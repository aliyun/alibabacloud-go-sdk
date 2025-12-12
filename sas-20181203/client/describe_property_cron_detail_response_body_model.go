// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyCronDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribePropertyCronDetailResponseBodyPageInfo) *DescribePropertyCronDetailResponseBody
	GetPageInfo() *DescribePropertyCronDetailResponseBodyPageInfo
	SetPropertys(v []*DescribePropertyCronDetailResponseBodyPropertys) *DescribePropertyCronDetailResponseBody
	GetPropertys() []*DescribePropertyCronDetailResponseBodyPropertys
	SetRequestId(v string) *DescribePropertyCronDetailResponseBody
	GetRequestId() *string
}

type DescribePropertyCronDetailResponseBody struct {
	// The pagination information.
	PageInfo *DescribePropertyCronDetailResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The details of the scheduled tasks.
	Propertys []*DescribePropertyCronDetailResponseBodyPropertys `json:"Propertys,omitempty" xml:"Propertys,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// B0C4E12E-CCE1-109D-9E62-7B95CBBAEF8C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePropertyCronDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyCronDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePropertyCronDetailResponseBody) GetPageInfo() *DescribePropertyCronDetailResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribePropertyCronDetailResponseBody) GetPropertys() []*DescribePropertyCronDetailResponseBodyPropertys {
	return s.Propertys
}

func (s *DescribePropertyCronDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePropertyCronDetailResponseBody) SetPageInfo(v *DescribePropertyCronDetailResponseBodyPageInfo) *DescribePropertyCronDetailResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribePropertyCronDetailResponseBody) SetPropertys(v []*DescribePropertyCronDetailResponseBodyPropertys) *DescribePropertyCronDetailResponseBody {
	s.Propertys = v
	return s
}

func (s *DescribePropertyCronDetailResponseBody) SetRequestId(v string) *DescribePropertyCronDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Propertys != nil {
		for _, item := range s.Propertys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePropertyCronDetailResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	NextToken   *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries returned per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePropertyCronDetailResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyCronDetailResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribePropertyCronDetailResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribePropertyCronDetailResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyCronDetailResponseBodyPageInfo) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePropertyCronDetailResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyCronDetailResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePropertyCronDetailResponseBodyPageInfo) SetCount(v int32) *DescribePropertyCronDetailResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPageInfo) SetCurrentPage(v int32) *DescribePropertyCronDetailResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPageInfo) SetNextToken(v string) *DescribePropertyCronDetailResponseBodyPageInfo {
	s.NextToken = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPageInfo) SetPageSize(v int32) *DescribePropertyCronDetailResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPageInfo) SetTotalCount(v int32) *DescribePropertyCronDetailResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribePropertyCronDetailResponseBodyPropertys struct {
	// The command that is used to run the scheduled task.
	//
	// example:
	//
	// /usr/lib64/sa/sa1 1 1
	Cmd *string `json:"Cmd,omitempty" xml:"Cmd,omitempty"`
	// The timestamp of the last fingerprint collection. Unit: milliseconds.
	//
	// example:
	//
	// 1649149566000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The ID of the server.
	//
	// example:
	//
	// i-hp35tftuh52wbp1g****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// hc-host-****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 100.104.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The IP addresses of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The MD5 hash value of the path to the scheduled task.
	//
	// example:
	//
	// 4cc8f97c2bf9cbabb2c2be2erqw****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The interval at which the scheduled task is performed.
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The path to the scheduled task.
	//
	// example:
	//
	// /etc/cron.d/root
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The username of the account that runs the scheduled task.
	//
	// example:
	//
	// root
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 162eb349-c2d9-4f8b-805c-75b43d4c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribePropertyCronDetailResponseBodyPropertys) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyCronDetailResponseBodyPropertys) GoString() string {
	return s.String()
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) GetCmd() *string {
	return s.Cmd
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) GetIp() *string {
	return s.Ip
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) GetMd5() *string {
	return s.Md5
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) GetPeriod() *string {
	return s.Period
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) GetSource() *string {
	return s.Source
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) GetUser() *string {
	return s.User
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) GetUuid() *string {
	return s.Uuid
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetCmd(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.Cmd = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetCreateTimestamp(v int64) *DescribePropertyCronDetailResponseBodyPropertys {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetInstanceId(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.InstanceId = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetInstanceName(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.InstanceName = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetInternetIp(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.InternetIp = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetIntranetIp(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.IntranetIp = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetIp(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.Ip = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetMd5(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.Md5 = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetPeriod(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.Period = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetSource(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.Source = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetUser(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.User = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) SetUuid(v string) *DescribePropertyCronDetailResponseBodyPropertys {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyCronDetailResponseBodyPropertys) Validate() error {
	return dara.Validate(s)
}
