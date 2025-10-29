// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceIndicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v *ListInstanceIndicesResponseBodyHeaders) *ListInstanceIndicesResponseBody
	GetHeaders() *ListInstanceIndicesResponseBodyHeaders
	SetRequestId(v string) *ListInstanceIndicesResponseBody
	GetRequestId() *string
	SetResult(v []*ListInstanceIndicesResponseBodyResult) *ListInstanceIndicesResponseBody
	GetResult() []*ListInstanceIndicesResponseBodyResult
}

type ListInstanceIndicesResponseBody struct {
	// The total size of the OpenStore cold stage index for this instance. Unit: bytes.
	Headers *ListInstanceIndicesResponseBodyHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
	// The total number of indexes in Cloud Hosting.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total storage space occupied by the current index. Unit: bytes.
	Result []*ListInstanceIndicesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListInstanceIndicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceIndicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceIndicesResponseBody) GetHeaders() *ListInstanceIndicesResponseBodyHeaders {
	return s.Headers
}

func (s *ListInstanceIndicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceIndicesResponseBody) GetResult() []*ListInstanceIndicesResponseBodyResult {
	return s.Result
}

func (s *ListInstanceIndicesResponseBody) SetHeaders(v *ListInstanceIndicesResponseBodyHeaders) *ListInstanceIndicesResponseBody {
	s.Headers = v
	return s
}

func (s *ListInstanceIndicesResponseBody) SetRequestId(v string) *ListInstanceIndicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceIndicesResponseBody) SetResult(v []*ListInstanceIndicesResponseBodyResult) *ListInstanceIndicesResponseBody {
	s.Result = v
	return s
}

func (s *ListInstanceIndicesResponseBody) Validate() error {
	if s.Headers != nil {
		if err := s.Headers.Validate(); err != nil {
			return err
		}
	}
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceIndicesResponseBodyHeaders struct {
	// The details of the index list.
	//
	// example:
	//
	// 15
	XManagedCount *int32 `json:"X-Managed-Count,omitempty" xml:"X-Managed-Count,omitempty"`
	// The total number of indexes in the OpenStore cold phase.
	//
	// example:
	//
	// 18093942932
	XManagedStorageSize *int64 `json:"X-Managed-StorageSize,omitempty" xml:"X-Managed-StorageSize,omitempty"`
	// The time when the index list was queried.
	//
	// example:
	//
	// 5
	XOSSCount *int32 `json:"X-OSS-Count,omitempty" xml:"X-OSS-Count,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 9093942932
	XOSSStorageSize *int64 `json:"X-OSS-StorageSize,omitempty" xml:"X-OSS-StorageSize,omitempty"`
}

func (s ListInstanceIndicesResponseBodyHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceIndicesResponseBodyHeaders) GoString() string {
	return s.String()
}

func (s *ListInstanceIndicesResponseBodyHeaders) GetXManagedCount() *int32 {
	return s.XManagedCount
}

func (s *ListInstanceIndicesResponseBodyHeaders) GetXManagedStorageSize() *int64 {
	return s.XManagedStorageSize
}

func (s *ListInstanceIndicesResponseBodyHeaders) GetXOSSCount() *int32 {
	return s.XOSSCount
}

func (s *ListInstanceIndicesResponseBodyHeaders) GetXOSSStorageSize() *int64 {
	return s.XOSSStorageSize
}

func (s *ListInstanceIndicesResponseBodyHeaders) SetXManagedCount(v int32) *ListInstanceIndicesResponseBodyHeaders {
	s.XManagedCount = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyHeaders) SetXManagedStorageSize(v int64) *ListInstanceIndicesResponseBodyHeaders {
	s.XManagedStorageSize = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyHeaders) SetXOSSCount(v int32) *ListInstanceIndicesResponseBodyHeaders {
	s.XOSSCount = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyHeaders) SetXOSSStorageSize(v int64) *ListInstanceIndicesResponseBodyHeaders {
	s.XOSSStorageSize = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyHeaders) Validate() error {
	return dara.Validate(s)
}

type ListInstanceIndicesResponseBodyResult struct {
	// The name of the Elasticsearch index.
	//
	// example:
	//
	// 2021-01-11T05:49:41.114Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// green
	Health *string `json:"health,omitempty" xml:"health,omitempty"`
	// example:
	//
	// {    "indices": {         ".ds-console-2021.08.18-000002": {             "index": ".ds-console-2021.08.18-000002",             "managed": true,             "policy": "console",             "lifecycle_date_millis": 1629277498775,             "age": "2.64h",             "phase": "hot",             "phase_time_millis": 1629277450334,             "action": "complete",             "action_time_millis": 1629278605586,             "step": "complete",             "step_time_millis": 1629278605586,             "phase_execution": {                 "policy": "console",                 "phase_definition": {                     "min_age": "0s",                     "actions": {                         "rollover": {                             "max_size": "1gb",                             "max_age": "1d",                             "max_docs": 10000                         },                         "set_priority": {                             "priority": 1000                         }                     }                 },                 "version": 1,                 "modified_date_in_millis": 1629277370953             }         }     } }
	IlmExplain *string `json:"ilmExplain,omitempty" xml:"ilmExplain,omitempty"`
	// The managed status of the index. The following three statuses are supported:
	//
	// 	- following: Hosting.
	//
	// 	- closing: The instance is being unhosted.
	//
	// 	- closed: unmanaged.
	//
	// example:
	//
	// false
	IsManaged *string `json:"isManaged,omitempty" xml:"isManaged,omitempty"`
	// The current storage lifecycle. Value meaning:
	//
	// 	- warm: warm.
	//
	// 	- cold: the cold phase.
	//
	// 	- hot: hot phase.
	//
	// 	- delete: deletes a stage.
	//
	// >  If this parameter is empty, the current index is not managed by the lifecycle.
	//
	// example:
	//
	// closing
	ManagedStatus *string `json:"managedStatus,omitempty" xml:"managedStatus,omitempty"`
	// The full lifecycle status of the current index.
	//
	// example:
	//
	// .kibana_task_manager_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// warm
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
	// The running status of the index. The following three statuses are supported:
	//
	// 	- green: healthy.
	//
	// 	- yellow: alerts.
	//
	// 	- red: an exception.
	//
	// example:
	//
	// 49298589
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListInstanceIndicesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceIndicesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInstanceIndicesResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListInstanceIndicesResponseBodyResult) GetHealth() *string {
	return s.Health
}

func (s *ListInstanceIndicesResponseBodyResult) GetIlmExplain() *string {
	return s.IlmExplain
}

func (s *ListInstanceIndicesResponseBodyResult) GetIsManaged() *string {
	return s.IsManaged
}

func (s *ListInstanceIndicesResponseBodyResult) GetManagedStatus() *string {
	return s.ManagedStatus
}

func (s *ListInstanceIndicesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListInstanceIndicesResponseBodyResult) GetPhase() *string {
	return s.Phase
}

func (s *ListInstanceIndicesResponseBodyResult) GetSize() *int64 {
	return s.Size
}

func (s *ListInstanceIndicesResponseBodyResult) SetCreateTime(v string) *ListInstanceIndicesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyResult) SetHealth(v string) *ListInstanceIndicesResponseBodyResult {
	s.Health = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyResult) SetIlmExplain(v string) *ListInstanceIndicesResponseBodyResult {
	s.IlmExplain = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyResult) SetIsManaged(v string) *ListInstanceIndicesResponseBodyResult {
	s.IsManaged = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyResult) SetManagedStatus(v string) *ListInstanceIndicesResponseBodyResult {
	s.ManagedStatus = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyResult) SetName(v string) *ListInstanceIndicesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyResult) SetPhase(v string) *ListInstanceIndicesResponseBodyResult {
	s.Phase = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyResult) SetSize(v int64) *ListInstanceIndicesResponseBodyResult {
	s.Size = &v
	return s
}

func (s *ListInstanceIndicesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
