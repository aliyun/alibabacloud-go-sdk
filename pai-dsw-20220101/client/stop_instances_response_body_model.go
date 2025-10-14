// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v map[string]*string) *StopInstancesResponseBody
	GetAccessDeniedDetail() map[string]*string
	SetCode(v string) *StopInstancesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *StopInstancesResponseBody
	GetHttpStatusCode() *int32
	SetInstanceIds(v []*string) *StopInstancesResponseBody
	GetInstanceIds() []*string
	SetMessage(v string) *StopInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopInstancesResponseBody
	GetSuccess() *bool
}

type StopInstancesResponseBody struct {
	AccessDeniedDetail map[string]*string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// PermissionError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceIds    []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstancesResponseBody) GetAccessDeniedDetail() map[string]*string {
	return s.AccessDeniedDetail
}

func (s *StopInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopInstancesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StopInstancesResponseBody) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *StopInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopInstancesResponseBody) SetAccessDeniedDetail(v map[string]*string) *StopInstancesResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *StopInstancesResponseBody) SetCode(v string) *StopInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *StopInstancesResponseBody) SetHttpStatusCode(v int32) *StopInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopInstancesResponseBody) SetInstanceIds(v []*string) *StopInstancesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *StopInstancesResponseBody) SetMessage(v string) *StopInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *StopInstancesResponseBody) SetRequestId(v string) *StopInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopInstancesResponseBody) SetSuccess(v bool) *StopInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *StopInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
