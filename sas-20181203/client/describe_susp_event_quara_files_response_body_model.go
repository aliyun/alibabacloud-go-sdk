// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspEventQuaraFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeSuspEventQuaraFilesResponseBody
	GetCount() *int32
	SetCurrentPage(v int32) *DescribeSuspEventQuaraFilesResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeSuspEventQuaraFilesResponseBody
	GetPageSize() *int32
	SetQuaraFiles(v []*DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) *DescribeSuspEventQuaraFilesResponseBody
	GetQuaraFiles() []*DescribeSuspEventQuaraFilesResponseBodyQuaraFiles
	SetRequestId(v string) *DescribeSuspEventQuaraFilesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeSuspEventQuaraFilesResponseBody
	GetTotalCount() *int32
}

type DescribeSuspEventQuaraFilesResponseBody struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 7
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// An array that consists of the quarantined files.
	QuaraFiles []*DescribeSuspEventQuaraFilesResponseBodyQuaraFiles `json:"QuaraFiles,omitempty" xml:"QuaraFiles,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 32A73759-4C0F-4801-BE98-901223ACEE9A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 38
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSuspEventQuaraFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventQuaraFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventQuaraFilesResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeSuspEventQuaraFilesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSuspEventQuaraFilesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSuspEventQuaraFilesResponseBody) GetQuaraFiles() []*DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	return s.QuaraFiles
}

func (s *DescribeSuspEventQuaraFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSuspEventQuaraFilesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetCount(v int32) *DescribeSuspEventQuaraFilesResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetCurrentPage(v int32) *DescribeSuspEventQuaraFilesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetPageSize(v int32) *DescribeSuspEventQuaraFilesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetQuaraFiles(v []*DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) *DescribeSuspEventQuaraFilesResponseBody {
	s.QuaraFiles = v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetRequestId(v string) *DescribeSuspEventQuaraFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBody) SetTotalCount(v int32) *DescribeSuspEventQuaraFilesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBody) Validate() error {
	if s.QuaraFiles != nil {
		for _, item := range s.QuaraFiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSuspEventQuaraFilesResponseBodyQuaraFiles struct {
	// The name of the event.
	//
	// example:
	//
	// WEBSHELL
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The type of the event.
	//
	// example:
	//
	// WebshellQuaraEventType
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The ID of the quarantined file.
	//
	// example:
	//
	// 26918
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance ID of the asset.
	//
	// example:
	//
	// i-2ze9t1qp36n1436m****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server on which the quarantined file is located.
	//
	// example:
	//
	// iZwz98dkiw3vbrtqrt5v****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server on which the quarantined file is located.
	//
	// example:
	//
	// 47.XX.XX.131
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server on which the quarantined file is located.
	//
	// example:
	//
	// 172.16.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The public IP address of the server on which the quarantined file is located.
	//
	// example:
	//
	// 47.XX.XX.131
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The download link of the quarantined file.
	//
	// example:
	//
	// https://xxx.xxx/xxx
	Link *string `json:"Link,omitempty" xml:"Link,omitempty"`
	// The MD5 hash value of the quarantined file.
	//
	// example:
	//
	// 5ddebe926acc7ed39a664409bfd0ec10
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The time when the quarantined file was last modified.
	//
	// example:
	//
	// 2020-06-11 20:37:08
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The path to the quarantined file on the server.
	//
	// example:
	//
	// /var/www/html/webshell-sample-master/others/defc3e21bab59e2a2ab49f7eda99f65f83d4d349.jpg
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The status of the quarantined file. Valid values:
	//
	// 	- **quaraFailed**: The file fails to be quarantined.
	//
	// 	- **quaraDone**: The file is quarantined.
	//
	// 	- **quaraing**: The file is being quarantined.
	//
	// 	- **rollbackFailed**: The system fails to cancel quarantining the file.
	//
	// 	- **rollbackDone**: The system cancelled quarantining the file.
	//
	// 	- **rollbacking**: The system is cancelling quarantining the file.
	//
	// example:
	//
	// rollbackDone
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The unique ID of the event.
	//
	// example:
	//
	// 228f890e56eae9eec6a42c7ea801b538
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 04a0e735-ad32-4835-b635-0458d77b****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) GetEventName() *string {
	return s.EventName
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) GetEventType() *string {
	return s.EventType
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) GetId() *int32 {
	return s.Id
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) GetIp() *string {
	return s.Ip
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) GetLink() *string {
	return s.Link
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) GetMd5() *string {
	return s.Md5
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) GetPath() *string {
	return s.Path
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) GetStatus() *string {
	return s.Status
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) GetTag() *string {
	return s.Tag
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetEventName(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.EventName = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetEventType(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.EventType = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetId(v int32) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Id = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetInstanceId(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.InstanceId = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetInstanceName(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.InstanceName = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetInternetIp(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.InternetIp = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetIntranetIp(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.IntranetIp = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetIp(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Ip = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetLink(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Link = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetMd5(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Md5 = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetModifyTime(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.ModifyTime = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetPath(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Path = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetStatus(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Status = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetTag(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Tag = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) SetUuid(v string) *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles {
	s.Uuid = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponseBodyQuaraFiles) Validate() error {
	return dara.Validate(s)
}
