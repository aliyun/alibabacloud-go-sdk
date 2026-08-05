// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAsyncTaskResponseBody
	GetRequestId() *string
	SetResult(v *GetAsyncTaskResponseBodyResult) *GetAsyncTaskResponseBody
	GetResult() *GetAsyncTaskResponseBodyResult
}

type GetAsyncTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 58113A95-1858-5674-87E5-192AEE6FD9DD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The response result.
	Result *GetAsyncTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetAsyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAsyncTaskResponseBody) GetResult() *GetAsyncTaskResponseBodyResult {
	return s.Result
}

func (s *GetAsyncTaskResponseBody) SetRequestId(v string) *GetAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetResult(v *GetAsyncTaskResponseBodyResult) *GetAsyncTaskResponseBody {
	s.Result = v
	return s
}

func (s *GetAsyncTaskResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAsyncTaskResponseBodyResult struct {
	// The creation time.
	//
	// example:
	//
	// 1729684154
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// The playground data ID.
	//
	// example:
	//
	// 123
	DataId *int32 `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// The asynchronous task ID.
	//
	// example:
	//
	// taskf90b77d481c47b05c18266a31b6cdbdd
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
	// {}
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
	// - PENDING: in progress.
	//
	// - SUCCESS: parsing succeeded.
	//
	// - FAILED: parsing failed.
	//
	// example:
	//
	// PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1729684154
	Updated *int64 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s GetAsyncTaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetAsyncTaskResponseBodyResult) GetCreated() *int64 {
	return s.Created
}

func (s *GetAsyncTaskResponseBodyResult) GetDataId() *int32 {
	return s.DataId
}

func (s *GetAsyncTaskResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *GetAsyncTaskResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetAsyncTaskResponseBodyResult) GetResult() *string {
	return s.Result
}

func (s *GetAsyncTaskResponseBodyResult) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetAsyncTaskResponseBodyResult) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetAsyncTaskResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetAsyncTaskResponseBodyResult) GetUpdated() *int64 {
	return s.Updated
}

func (s *GetAsyncTaskResponseBodyResult) SetCreated(v int64) *GetAsyncTaskResponseBodyResult {
	s.Created = &v
	return s
}

func (s *GetAsyncTaskResponseBodyResult) SetDataId(v int32) *GetAsyncTaskResponseBodyResult {
	s.DataId = &v
	return s
}

func (s *GetAsyncTaskResponseBodyResult) SetId(v string) *GetAsyncTaskResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetAsyncTaskResponseBodyResult) SetName(v string) *GetAsyncTaskResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetAsyncTaskResponseBodyResult) SetResult(v string) *GetAsyncTaskResponseBodyResult {
	s.Result = &v
	return s
}

func (s *GetAsyncTaskResponseBodyResult) SetServiceId(v string) *GetAsyncTaskResponseBodyResult {
	s.ServiceId = &v
	return s
}

func (s *GetAsyncTaskResponseBodyResult) SetServiceType(v string) *GetAsyncTaskResponseBodyResult {
	s.ServiceType = &v
	return s
}

func (s *GetAsyncTaskResponseBodyResult) SetStatus(v string) *GetAsyncTaskResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetAsyncTaskResponseBodyResult) SetUpdated(v int64) *GetAsyncTaskResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *GetAsyncTaskResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
