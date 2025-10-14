// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v map[string]interface{}) *DeleteInstancesResponseBody
	GetAccessDeniedDetail() map[string]interface{}
	SetCode(v string) *DeleteInstancesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteInstancesResponseBody
	GetHttpStatusCode() *int32
	SetInstanceIds(v []*string) *DeleteInstancesResponseBody
	GetInstanceIds() []*string
	SetMessage(v string) *DeleteInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteInstancesResponseBody
	GetSuccess() *bool
}

type DeleteInstancesResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail map[string]interface{} `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
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
	// The instance can not be deleted.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// AA527C1A-F259-5E53-A4DD-D0941193F02D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstancesResponseBody) GetAccessDeniedDetail() map[string]interface{} {
	return s.AccessDeniedDetail
}

func (s *DeleteInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteInstancesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteInstancesResponseBody) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DeleteInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteInstancesResponseBody) SetAccessDeniedDetail(v map[string]interface{}) *DeleteInstancesResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DeleteInstancesResponseBody) SetCode(v string) *DeleteInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstancesResponseBody) SetHttpStatusCode(v int32) *DeleteInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteInstancesResponseBody) SetInstanceIds(v []*string) *DeleteInstancesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *DeleteInstancesResponseBody) SetMessage(v string) *DeleteInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstancesResponseBody) SetRequestId(v string) *DeleteInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstancesResponseBody) SetSuccess(v bool) *DeleteInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
