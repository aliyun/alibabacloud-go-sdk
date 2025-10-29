// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutRecordStorageLifeCycleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *PutRecordStorageLifeCycleResponseBody
	GetCode() *int32
	SetMsg(v string) *PutRecordStorageLifeCycleResponseBody
	GetMsg() *string
	SetRequestId(v string) *PutRecordStorageLifeCycleResponseBody
	GetRequestId() *string
}

type PutRecordStorageLifeCycleResponseBody struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutRecordStorageLifeCycleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutRecordStorageLifeCycleResponseBody) GoString() string {
	return s.String()
}

func (s *PutRecordStorageLifeCycleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *PutRecordStorageLifeCycleResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *PutRecordStorageLifeCycleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutRecordStorageLifeCycleResponseBody) SetCode(v int32) *PutRecordStorageLifeCycleResponseBody {
	s.Code = &v
	return s
}

func (s *PutRecordStorageLifeCycleResponseBody) SetMsg(v string) *PutRecordStorageLifeCycleResponseBody {
	s.Msg = &v
	return s
}

func (s *PutRecordStorageLifeCycleResponseBody) SetRequestId(v string) *PutRecordStorageLifeCycleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutRecordStorageLifeCycleResponseBody) Validate() error {
	return dara.Validate(s)
}
