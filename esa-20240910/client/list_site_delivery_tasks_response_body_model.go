// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSiteDeliveryTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListSiteDeliveryTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSiteDeliveryTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSiteDeliveryTasksResponseBody
	GetRequestId() *string
	SetTasks(v []*ListSiteDeliveryTasksResponseBodyTasks) *ListSiteDeliveryTasksResponseBody
	GetTasks() []*ListSiteDeliveryTasksResponseBodyTasks
	SetTotalCount(v int32) *ListSiteDeliveryTasksResponseBody
	GetTotalCount() *int32
}

type ListSiteDeliveryTasksResponseBody struct {
	// The page number. Default value: 0.
	//
	// example:
	//
	// 0
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1 to 500**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 34DCBC8A-****-****-****-6DAA11D7DDBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The delivery tasks.
	Tasks []*ListSiteDeliveryTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of log delivery tasks.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSiteDeliveryTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSiteDeliveryTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListSiteDeliveryTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSiteDeliveryTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSiteDeliveryTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSiteDeliveryTasksResponseBody) GetTasks() []*ListSiteDeliveryTasksResponseBodyTasks {
	return s.Tasks
}

func (s *ListSiteDeliveryTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSiteDeliveryTasksResponseBody) SetPageNumber(v int32) *ListSiteDeliveryTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSiteDeliveryTasksResponseBody) SetPageSize(v int32) *ListSiteDeliveryTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSiteDeliveryTasksResponseBody) SetRequestId(v string) *ListSiteDeliveryTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSiteDeliveryTasksResponseBody) SetTasks(v []*ListSiteDeliveryTasksResponseBodyTasks) *ListSiteDeliveryTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListSiteDeliveryTasksResponseBody) SetTotalCount(v int32) *ListSiteDeliveryTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSiteDeliveryTasksResponseBody) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSiteDeliveryTasksResponseBodyTasks struct {
	// The log category. Valid values:
	//
	// 	- dcdn_log_access_l1 (default): access logs.
	//
	// 	- dcdn_log_er: Edge Routine logs.
	//
	// 	- dcdn_log_waf: firewall logs.
	//
	// 	- dcdn_log_ipa: TCP/UDP proxy logs.
	//
	// example:
	//
	// dcdn_log_access_l1
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The data center. Valid values:
	//
	// 	- cn: the Chinese mainland.
	//
	// 	- sg: outside the Chinese mainland.
	//
	// example:
	//
	// cn
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The destination of the delivery. Valid values:
	//
	// 1.  sls: Alibaba Cloud Simple Log Service (SLS).
	//
	// 2.  http: HTTP server.
	//
	// 3.  aws3: Amazon Simple Storage Service (S3).
	//
	// 4.  oss: Alibaba Cloud Object Storage Service (OSS).
	//
	// 5.  kafka: Kafka.
	//
	// 6.  aws3cmpt: S3-compatible storage service.
	//
	// example:
	//
	// sls
	DeliveryType *string `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// The status of the delivery task.
	//
	// 	- **online**
	//
	// 	- **offline**
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the delivery task.
	//
	// example:
	//
	// cdn-test-task
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListSiteDeliveryTasksResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s ListSiteDeliveryTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListSiteDeliveryTasksResponseBodyTasks) GetBusinessType() *string {
	return s.BusinessType
}

func (s *ListSiteDeliveryTasksResponseBodyTasks) GetDataCenter() *string {
	return s.DataCenter
}

func (s *ListSiteDeliveryTasksResponseBodyTasks) GetDeliveryType() *string {
	return s.DeliveryType
}

func (s *ListSiteDeliveryTasksResponseBodyTasks) GetStatus() *string {
	return s.Status
}

func (s *ListSiteDeliveryTasksResponseBodyTasks) GetTaskName() *string {
	return s.TaskName
}

func (s *ListSiteDeliveryTasksResponseBodyTasks) SetBusinessType(v string) *ListSiteDeliveryTasksResponseBodyTasks {
	s.BusinessType = &v
	return s
}

func (s *ListSiteDeliveryTasksResponseBodyTasks) SetDataCenter(v string) *ListSiteDeliveryTasksResponseBodyTasks {
	s.DataCenter = &v
	return s
}

func (s *ListSiteDeliveryTasksResponseBodyTasks) SetDeliveryType(v string) *ListSiteDeliveryTasksResponseBodyTasks {
	s.DeliveryType = &v
	return s
}

func (s *ListSiteDeliveryTasksResponseBodyTasks) SetStatus(v string) *ListSiteDeliveryTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListSiteDeliveryTasksResponseBodyTasks) SetTaskName(v string) *ListSiteDeliveryTasksResponseBodyTasks {
	s.TaskName = &v
	return s
}

func (s *ListSiteDeliveryTasksResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}
