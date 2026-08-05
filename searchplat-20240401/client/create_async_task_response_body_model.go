// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAsyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAsyncTaskResponseBody
	GetRequestId() *string
	SetResult(v []*CreateAsyncTaskResponseBodyResult) *CreateAsyncTaskResponseBody
	GetResult() []*CreateAsyncTaskResponseBodyResult
}

type CreateAsyncTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1CC93E65-6734-5060-BEF7-0EB0A4862BCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result []*CreateAsyncTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s CreateAsyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAsyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAsyncTaskResponseBody) GetResult() []*CreateAsyncTaskResponseBodyResult {
	return s.Result
}

func (s *CreateAsyncTaskResponseBody) SetRequestId(v string) *CreateAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAsyncTaskResponseBody) SetResult(v []*CreateAsyncTaskResponseBodyResult) *CreateAsyncTaskResponseBody {
	s.Result = v
	return s
}

func (s *CreateAsyncTaskResponseBody) Validate() error {
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

type CreateAsyncTaskResponseBodyResult struct {
	// The creation time.
	//
	// example:
	//
	// 1729669284
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// The data ID.
	//
	// example:
	//
	// 1232
	DataId *int32 `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// fae9bcc5-949f-4c31-b9b7-a273bf891699
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The task name.
	//
	// example:
	//
	// 文档解析任务
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The parsing result.
	//
	// example:
	//
	// xx
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// The service ID.
	//
	// example:
	//
	// ops-document-analyze-001
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The service type.
	//
	// example:
	//
	// document-analyze
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// The task status. Valid values:
	//
	// - PENDING: In progress.
	//
	// - SUCCESS: Parsing succeeded.
	//
	// - FAILED: Parsing failed.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1729669284
	Updated *int64 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s CreateAsyncTaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateAsyncTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAsyncTaskResponseBodyResult) GetCreated() *int64 {
	return s.Created
}

func (s *CreateAsyncTaskResponseBodyResult) GetDataId() *int32 {
	return s.DataId
}

func (s *CreateAsyncTaskResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *CreateAsyncTaskResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *CreateAsyncTaskResponseBodyResult) GetResult() *string {
	return s.Result
}

func (s *CreateAsyncTaskResponseBodyResult) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateAsyncTaskResponseBodyResult) GetServiceType() *string {
	return s.ServiceType
}

func (s *CreateAsyncTaskResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *CreateAsyncTaskResponseBodyResult) GetUpdated() *int64 {
	return s.Updated
}

func (s *CreateAsyncTaskResponseBodyResult) SetCreated(v int64) *CreateAsyncTaskResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateAsyncTaskResponseBodyResult) SetDataId(v int32) *CreateAsyncTaskResponseBodyResult {
	s.DataId = &v
	return s
}

func (s *CreateAsyncTaskResponseBodyResult) SetId(v string) *CreateAsyncTaskResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateAsyncTaskResponseBodyResult) SetName(v string) *CreateAsyncTaskResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateAsyncTaskResponseBodyResult) SetResult(v string) *CreateAsyncTaskResponseBodyResult {
	s.Result = &v
	return s
}

func (s *CreateAsyncTaskResponseBodyResult) SetServiceId(v string) *CreateAsyncTaskResponseBodyResult {
	s.ServiceId = &v
	return s
}

func (s *CreateAsyncTaskResponseBodyResult) SetServiceType(v string) *CreateAsyncTaskResponseBodyResult {
	s.ServiceType = &v
	return s
}

func (s *CreateAsyncTaskResponseBodyResult) SetStatus(v string) *CreateAsyncTaskResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateAsyncTaskResponseBodyResult) SetUpdated(v int64) *CreateAsyncTaskResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *CreateAsyncTaskResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
