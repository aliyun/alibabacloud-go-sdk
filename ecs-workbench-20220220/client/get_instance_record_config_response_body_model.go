// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceRecordConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceRecordConfigResponseBody
	GetCode() *string
	SetMessage(v string) *GetInstanceRecordConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstanceRecordConfigResponseBody
	GetRequestId() *string
	SetRoot(v *GetInstanceRecordConfigResponseBodyRoot) *GetInstanceRecordConfigResponseBody
	GetRoot() *GetInstanceRecordConfigResponseBodyRoot
	SetSuccess(v bool) *GetInstanceRecordConfigResponseBody
	GetSuccess() *bool
}

type GetInstanceRecordConfigResponseBody struct {
	// example:
	//
	// InvalidParamter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter is null or invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Root      *GetInstanceRecordConfigResponseBodyRoot `json:"Root,omitempty" xml:"Root,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceRecordConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceRecordConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceRecordConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceRecordConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceRecordConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceRecordConfigResponseBody) GetRoot() *GetInstanceRecordConfigResponseBodyRoot {
	return s.Root
}

func (s *GetInstanceRecordConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceRecordConfigResponseBody) SetCode(v string) *GetInstanceRecordConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceRecordConfigResponseBody) SetMessage(v string) *GetInstanceRecordConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceRecordConfigResponseBody) SetRequestId(v string) *GetInstanceRecordConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceRecordConfigResponseBody) SetRoot(v *GetInstanceRecordConfigResponseBodyRoot) *GetInstanceRecordConfigResponseBody {
	s.Root = v
	return s
}

func (s *GetInstanceRecordConfigResponseBody) SetSuccess(v bool) *GetInstanceRecordConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceRecordConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetInstanceRecordConfigResponseBodyRoot struct {
	// example:
	//
	// 7
	ExpirationDays *int32 `json:"ExpirationDays,omitempty" xml:"ExpirationDays,omitempty"`
	// example:
	//
	// i-xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 123
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// acs:oss:cn-shanghai:123:workbench-record-123-1/record
	RecordStorageTarget *string `json:"RecordStorageTarget,omitempty" xml:"RecordStorageTarget,omitempty"`
}

func (s GetInstanceRecordConfigResponseBodyRoot) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceRecordConfigResponseBodyRoot) GoString() string {
	return s.String()
}

func (s *GetInstanceRecordConfigResponseBodyRoot) GetExpirationDays() *int32 {
	return s.ExpirationDays
}

func (s *GetInstanceRecordConfigResponseBodyRoot) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceRecordConfigResponseBodyRoot) GetParentId() *string {
	return s.ParentId
}

func (s *GetInstanceRecordConfigResponseBodyRoot) GetRecordStorageTarget() *string {
	return s.RecordStorageTarget
}

func (s *GetInstanceRecordConfigResponseBodyRoot) SetExpirationDays(v int32) *GetInstanceRecordConfigResponseBodyRoot {
	s.ExpirationDays = &v
	return s
}

func (s *GetInstanceRecordConfigResponseBodyRoot) SetInstanceId(v string) *GetInstanceRecordConfigResponseBodyRoot {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceRecordConfigResponseBodyRoot) SetParentId(v string) *GetInstanceRecordConfigResponseBodyRoot {
	s.ParentId = &v
	return s
}

func (s *GetInstanceRecordConfigResponseBodyRoot) SetRecordStorageTarget(v string) *GetInstanceRecordConfigResponseBodyRoot {
	s.RecordStorageTarget = &v
	return s
}

func (s *GetInstanceRecordConfigResponseBodyRoot) Validate() error {
	return dara.Validate(s)
}
