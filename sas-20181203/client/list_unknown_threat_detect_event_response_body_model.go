// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnknownThreatDetectEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListUnknownThreatDetectEventResponseBodyData) *ListUnknownThreatDetectEventResponseBody
	GetData() []*ListUnknownThreatDetectEventResponseBodyData
	SetPageInfo(v *ListUnknownThreatDetectEventResponseBodyPageInfo) *ListUnknownThreatDetectEventResponseBody
	GetPageInfo() *ListUnknownThreatDetectEventResponseBodyPageInfo
	SetRequestId(v string) *ListUnknownThreatDetectEventResponseBody
	GetRequestId() *string
}

type ListUnknownThreatDetectEventResponseBody struct {
	Data     []*ListUnknownThreatDetectEventResponseBodyData   `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageInfo *ListUnknownThreatDetectEventResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 7532B7EE-7CE7-5F4D-BF04-B12447******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUnknownThreatDetectEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUnknownThreatDetectEventResponseBody) GoString() string {
	return s.String()
}

func (s *ListUnknownThreatDetectEventResponseBody) GetData() []*ListUnknownThreatDetectEventResponseBodyData {
	return s.Data
}

func (s *ListUnknownThreatDetectEventResponseBody) GetPageInfo() *ListUnknownThreatDetectEventResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListUnknownThreatDetectEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUnknownThreatDetectEventResponseBody) SetData(v []*ListUnknownThreatDetectEventResponseBodyData) *ListUnknownThreatDetectEventResponseBody {
	s.Data = v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBody) SetPageInfo(v *ListUnknownThreatDetectEventResponseBodyPageInfo) *ListUnknownThreatDetectEventResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBody) SetRequestId(v string) *ListUnknownThreatDetectEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUnknownThreatDetectEventResponseBodyData struct {
	// example:
	//
	// [{"5133":"pickup -l -t unix -u"},{"1077":"/usr/libexec/postfix/master -w"},{"1":"/usr/lib/systemd/systemd --switched-root --system --deserialize 22"}]
	CmdChain *string `json:"CmdChain,omitempty" xml:"CmdChain,omitempty"`
	// example:
	//
	// /usr/sbin/sshd -D
	Cmdline *string `json:"Cmdline,omitempty" xml:"Cmdline,omitempty"`
	// example:
	//
	// 5
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1694576692000
	FirstTime *int64 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// example:
	//
	// 30368144069e7567bbb10eabc2******
	HashKey *string `json:"HashKey,omitempty" xml:"HashKey,omitempty"`
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// centos****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 172.16.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// example:
	//
	// 10.42.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// example:
	//
	// 1694576692000
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// example:
	//
	// 5b394b54ca632fe51c4ab4a6dbaf****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// example:
	//
	// /usr/sbin/sshd -D
	ParentCmdline *string `json:"ParentCmdline,omitempty" xml:"ParentCmdline,omitempty"`
	// example:
	//
	// 12
	ParentPid *string `json:"ParentPid,omitempty" xml:"ParentPid,omitempty"`
	// example:
	//
	// /usr/bin/tar
	ParentProcessPath *string `json:"ParentProcessPath,omitempty" xml:"ParentProcessPath,omitempty"`
	// example:
	//
	// 11
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// example:
	//
	// /usr/bin/tar
	ProcessPath *string `json:"ProcessPath,omitempty" xml:"ProcessPath,omitempty"`
	// example:
	//
	// 3a6fed5fc11392b3ee9f81caf017b48640d7458766a8eb0382899a605b41****
	Sha256 *string `json:"Sha256,omitempty" xml:"Sha256,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 6690a46c-0edb-4663-a641-3629d1a9****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListUnknownThreatDetectEventResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUnknownThreatDetectEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUnknownThreatDetectEventResponseBodyData) GetCmdChain() *string {
	return s.CmdChain
}

func (s *ListUnknownThreatDetectEventResponseBodyData) GetCmdline() *string {
	return s.Cmdline
}

func (s *ListUnknownThreatDetectEventResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *ListUnknownThreatDetectEventResponseBodyData) GetFirstTime() *int64 {
	return s.FirstTime
}

func (s *ListUnknownThreatDetectEventResponseBodyData) GetHashKey() *string {
	return s.HashKey
}

func (s *ListUnknownThreatDetectEventResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListUnknownThreatDetectEventResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListUnknownThreatDetectEventResponseBodyData) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListUnknownThreatDetectEventResponseBodyData) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListUnknownThreatDetectEventResponseBodyData) GetLastTime() *int64 {
	return s.LastTime
}

func (s *ListUnknownThreatDetectEventResponseBodyData) GetMd5() *string {
	return s.Md5
}

func (s *ListUnknownThreatDetectEventResponseBodyData) GetParentCmdline() *string {
	return s.ParentCmdline
}

func (s *ListUnknownThreatDetectEventResponseBodyData) GetParentPid() *string {
	return s.ParentPid
}

func (s *ListUnknownThreatDetectEventResponseBodyData) GetParentProcessPath() *string {
	return s.ParentProcessPath
}

func (s *ListUnknownThreatDetectEventResponseBodyData) GetPid() *string {
	return s.Pid
}

func (s *ListUnknownThreatDetectEventResponseBodyData) GetProcessPath() *string {
	return s.ProcessPath
}

func (s *ListUnknownThreatDetectEventResponseBodyData) GetSha256() *string {
	return s.Sha256
}

func (s *ListUnknownThreatDetectEventResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ListUnknownThreatDetectEventResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *ListUnknownThreatDetectEventResponseBodyData) SetCmdChain(v string) *ListUnknownThreatDetectEventResponseBodyData {
	s.CmdChain = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyData) SetCmdline(v string) *ListUnknownThreatDetectEventResponseBodyData {
	s.Cmdline = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyData) SetCount(v int32) *ListUnknownThreatDetectEventResponseBodyData {
	s.Count = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyData) SetFirstTime(v int64) *ListUnknownThreatDetectEventResponseBodyData {
	s.FirstTime = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyData) SetHashKey(v string) *ListUnknownThreatDetectEventResponseBodyData {
	s.HashKey = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyData) SetId(v string) *ListUnknownThreatDetectEventResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyData) SetInstanceName(v string) *ListUnknownThreatDetectEventResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyData) SetInternetIp(v string) *ListUnknownThreatDetectEventResponseBodyData {
	s.InternetIp = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyData) SetIntranetIp(v string) *ListUnknownThreatDetectEventResponseBodyData {
	s.IntranetIp = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyData) SetLastTime(v int64) *ListUnknownThreatDetectEventResponseBodyData {
	s.LastTime = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyData) SetMd5(v string) *ListUnknownThreatDetectEventResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyData) SetParentCmdline(v string) *ListUnknownThreatDetectEventResponseBodyData {
	s.ParentCmdline = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyData) SetParentPid(v string) *ListUnknownThreatDetectEventResponseBodyData {
	s.ParentPid = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyData) SetParentProcessPath(v string) *ListUnknownThreatDetectEventResponseBodyData {
	s.ParentProcessPath = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyData) SetPid(v string) *ListUnknownThreatDetectEventResponseBodyData {
	s.Pid = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyData) SetProcessPath(v string) *ListUnknownThreatDetectEventResponseBodyData {
	s.ProcessPath = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyData) SetSha256(v string) *ListUnknownThreatDetectEventResponseBodyData {
	s.Sha256 = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyData) SetStatus(v int32) *ListUnknownThreatDetectEventResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyData) SetUuid(v string) *ListUnknownThreatDetectEventResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListUnknownThreatDetectEventResponseBodyPageInfo struct {
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 149
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUnknownThreatDetectEventResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListUnknownThreatDetectEventResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListUnknownThreatDetectEventResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListUnknownThreatDetectEventResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListUnknownThreatDetectEventResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUnknownThreatDetectEventResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUnknownThreatDetectEventResponseBodyPageInfo) SetCount(v int32) *ListUnknownThreatDetectEventResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyPageInfo) SetCurrentPage(v int32) *ListUnknownThreatDetectEventResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyPageInfo) SetPageSize(v int32) *ListUnknownThreatDetectEventResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyPageInfo) SetTotalCount(v int32) *ListUnknownThreatDetectEventResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListUnknownThreatDetectEventResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
