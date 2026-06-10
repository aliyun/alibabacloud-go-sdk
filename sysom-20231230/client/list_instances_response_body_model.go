// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstancesResponseBody
	GetCode() *string
	SetData(v []*ListInstancesResponseBodyData) *ListInstancesResponseBody
	GetData() []*ListInstancesResponseBodyData
	SetMessage(v string) *ListInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListInstancesResponseBody
	GetTotal() *int64
}

type ListInstancesResponseBody struct {
	// Status code
	//
	// - `code == Success` indicates that authorization succeeded.
	//
	// - Any other status code indicates that authorization failed. When authorization fails, check the `message` field for detailed error information.
	//
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Return Result.
	Data []*ListInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Error message
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the request error message.
	//
	// example:
	//
	// Requests for llm service failed
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request RequestId
	//
	// example:
	//
	// 9515E5A0-8905-59B0-9BBF-5F0BE568C3A0
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Total number of records
	//
	// example:
	//
	// 623
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstancesResponseBody) GetData() []*ListInstancesResponseBodyData {
	return s.Data
}

func (s *ListInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListInstancesResponseBody) SetCode(v string) *ListInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesResponseBody) SetData(v []*ListInstancesResponseBodyData) *ListInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListInstancesResponseBody) SetMessage(v string) *ListInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotal(v int64) *ListInstancesResponseBody {
	s.Total = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
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

type ListInstancesResponseBodyData struct {
	// Cluster ID
	//
	// example:
	//
	// c2218ca2b76ec45e7b7ee1693f6fcd374
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// ECS instance ID
	//
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// Milvus version of the instance
	//
	// example:
	//
	// 5.10.134-16.1.an8.x86_64
	KernelVersion *string `json:"kernel_version,omitempty" xml:"kernel_version,omitempty"`
	// Metadata of the instance
	//
	// example:
	//
	// {
	//
	//     "uname": "Linux",
	//
	//     "oncpu": "off",
	//
	//     "release": "5.10.134-16.3.al8.aarch64",
	//
	//     "monitor": "on",
	//
	//     "version_id": "3",
	//
	//     "version": "3 (Soaring Falcon)",
	//
	//     "podNs": [
	//
	//     ],
	//
	//     "machine": "aarch64",
	//
	//     "name": "Alibaba Cloud Linux",
	//
	//     "sysak": "3.4.0-1",
	//
	//     "id": "alinux",
	//
	//     "region": "cn-hangzhou",
	//
	//     "centos-release": "Alibaba Cloud Linux release 3 (Soaring Falcon)"
	//
	// }
	Meta interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	// Architecture of the ECS instance
	//
	// example:
	//
	// x86
	OsArch *string `json:"os_arch,omitempty" xml:"os_arch,omitempty"`
	// Health score of the instance
	//
	// example:
	//
	// 100
	OsHealthScore *string `json:"os_health_score,omitempty" xml:"os_health_score,omitempty"`
	// Operating system name of the instance (retrieved from /etc/os-release)
	//
	// example:
	//
	// Anolis OS
	OsName *string `json:"os_name,omitempty" xml:"os_name,omitempty"`
	// Operating system name ID of the instance (retrieved from /etc/os-release)
	//
	// example:
	//
	// anolis
	OsNameId *string `json:"os_name_id,omitempty" xml:"os_name_id,omitempty"`
	// Operating system version of the instance (obtained from /etc/os-release)
	//
	// example:
	//
	// 8.9
	OsVersion *string `json:"os_version,omitempty" xml:"os_version,omitempty"`
	// Operating system version ID of the instance (retrieved from /etc/os-release)
	//
	// example:
	//
	// rhel fedora centos
	OsVersionId *string `json:"os_version_id,omitempty" xml:"os_version_id,omitempty"`
	// Region where the instance is located
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// Status of the instance
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListInstancesResponseBodyData) GetInstance() *string {
	return s.Instance
}

func (s *ListInstancesResponseBodyData) GetKernelVersion() *string {
	return s.KernelVersion
}

func (s *ListInstancesResponseBodyData) GetMeta() interface{} {
	return s.Meta
}

func (s *ListInstancesResponseBodyData) GetOsArch() *string {
	return s.OsArch
}

func (s *ListInstancesResponseBodyData) GetOsHealthScore() *string {
	return s.OsHealthScore
}

func (s *ListInstancesResponseBodyData) GetOsName() *string {
	return s.OsName
}

func (s *ListInstancesResponseBodyData) GetOsNameId() *string {
	return s.OsNameId
}

func (s *ListInstancesResponseBodyData) GetOsVersion() *string {
	return s.OsVersion
}

func (s *ListInstancesResponseBodyData) GetOsVersionId() *string {
	return s.OsVersionId
}

func (s *ListInstancesResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *ListInstancesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyData) SetClusterId(v string) *ListInstancesResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetInstance(v string) *ListInstancesResponseBodyData {
	s.Instance = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetKernelVersion(v string) *ListInstancesResponseBodyData {
	s.KernelVersion = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetMeta(v interface{}) *ListInstancesResponseBodyData {
	s.Meta = v
	return s
}

func (s *ListInstancesResponseBodyData) SetOsArch(v string) *ListInstancesResponseBodyData {
	s.OsArch = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetOsHealthScore(v string) *ListInstancesResponseBodyData {
	s.OsHealthScore = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetOsName(v string) *ListInstancesResponseBodyData {
	s.OsName = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetOsNameId(v string) *ListInstancesResponseBodyData {
	s.OsNameId = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetOsVersion(v string) *ListInstancesResponseBodyData {
	s.OsVersion = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetOsVersionId(v string) *ListInstancesResponseBodyData {
	s.OsVersionId = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetRegion(v string) *ListInstancesResponseBodyData {
	s.Region = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetStatus(v string) *ListInstancesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
