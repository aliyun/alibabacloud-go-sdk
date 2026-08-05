// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsyncTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAsyncTasksResponseBody
	GetRequestId() *string
	SetResult(v []*ListAsyncTasksResponseBodyResult) *ListAsyncTasksResponseBody
	GetResult() []*ListAsyncTasksResponseBodyResult
}

type ListAsyncTasksResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2BA0504F-B179-586D-8210-A7C7C09A9907
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result []*ListAsyncTasksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListAsyncTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAsyncTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAsyncTasksResponseBody) GetResult() []*ListAsyncTasksResponseBodyResult {
	return s.Result
}

func (s *ListAsyncTasksResponseBody) SetRequestId(v string) *ListAsyncTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetResult(v []*ListAsyncTasksResponseBodyResult) *ListAsyncTasksResponseBody {
	s.Result = v
	return s
}

func (s *ListAsyncTasksResponseBody) Validate() error {
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

type ListAsyncTasksResponseBodyResult struct {
	// The creation timestamp.
	//
	// example:
	//
	// 1745893195510
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// The trial data ID.
	//
	// example:
	//
	// 1232
	DataId *int32 `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// The asynchronous task ID.
	//
	// example:
	//
	// ae1cea6dc680b98b908a757050c406c9
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The file name.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The task execution result.
	//
	// example:
	//
	// xxxxx
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// The service ID.
	//
	// example:
	//
	// ops-document-analyze-001
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The service type.
	//
	// - document-analyze.
	//
	// example:
	//
	// document-analyze
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// The task status. Valid values:
	//
	// - PENDING: in progress.
	//
	// - SUCCESS: parsing succeeded.
	//
	// - FAILED: parsing failed.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1729665694
	Updated *int64 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListAsyncTasksResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListAsyncTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksResponseBodyResult) GetCreated() *int64 {
	return s.Created
}

func (s *ListAsyncTasksResponseBodyResult) GetDataId() *int32 {
	return s.DataId
}

func (s *ListAsyncTasksResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListAsyncTasksResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListAsyncTasksResponseBodyResult) GetResult() *string {
	return s.Result
}

func (s *ListAsyncTasksResponseBodyResult) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListAsyncTasksResponseBodyResult) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListAsyncTasksResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListAsyncTasksResponseBodyResult) GetUpdated() *int64 {
	return s.Updated
}

func (s *ListAsyncTasksResponseBodyResult) SetCreated(v int64) *ListAsyncTasksResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListAsyncTasksResponseBodyResult) SetDataId(v int32) *ListAsyncTasksResponseBodyResult {
	s.DataId = &v
	return s
}

func (s *ListAsyncTasksResponseBodyResult) SetId(v string) *ListAsyncTasksResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListAsyncTasksResponseBodyResult) SetName(v string) *ListAsyncTasksResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListAsyncTasksResponseBodyResult) SetResult(v string) *ListAsyncTasksResponseBodyResult {
	s.Result = &v
	return s
}

func (s *ListAsyncTasksResponseBodyResult) SetServiceId(v string) *ListAsyncTasksResponseBodyResult {
	s.ServiceId = &v
	return s
}

func (s *ListAsyncTasksResponseBodyResult) SetServiceType(v string) *ListAsyncTasksResponseBodyResult {
	s.ServiceType = &v
	return s
}

func (s *ListAsyncTasksResponseBodyResult) SetStatus(v string) *ListAsyncTasksResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListAsyncTasksResponseBodyResult) SetUpdated(v int64) *ListAsyncTasksResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListAsyncTasksResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
