// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllInstanceIdListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetAllInstanceIdListResponseBody
	GetCode() *int32
	SetInstanceIds(v map[string]interface{}) *GetAllInstanceIdListResponseBody
	GetInstanceIds() map[string]interface{}
	SetMessage(v string) *GetAllInstanceIdListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAllInstanceIdListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAllInstanceIdListResponseBody
	GetSuccess() *bool
}

type GetAllInstanceIdListResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The IDs of instances that are managed by the Alibaba Cloud account in all the regions.
	//
	// example:
	//
	// [{"cn-shenzhen": ["alikafka_post-cn-7pp2btvo****"],"us-west-1": ["alikafka_pre-cn-i7m2lxid****"],"cn-hangzhou": ["alikafka_pre-cn-i7m2hflj****","alikafka_pre-cn-zvp2hsje****","alikafka_pre-cn-zvp2kvc9****"]}]
	InstanceIds map[string]interface{} `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ABA4A7FD-E10F-45C7-9774-A5236015****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAllInstanceIdListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAllInstanceIdListResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllInstanceIdListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetAllInstanceIdListResponseBody) GetInstanceIds() map[string]interface{} {
	return s.InstanceIds
}

func (s *GetAllInstanceIdListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAllInstanceIdListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAllInstanceIdListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAllInstanceIdListResponseBody) SetCode(v int32) *GetAllInstanceIdListResponseBody {
	s.Code = &v
	return s
}

func (s *GetAllInstanceIdListResponseBody) SetInstanceIds(v map[string]interface{}) *GetAllInstanceIdListResponseBody {
	s.InstanceIds = v
	return s
}

func (s *GetAllInstanceIdListResponseBody) SetMessage(v string) *GetAllInstanceIdListResponseBody {
	s.Message = &v
	return s
}

func (s *GetAllInstanceIdListResponseBody) SetRequestId(v string) *GetAllInstanceIdListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAllInstanceIdListResponseBody) SetSuccess(v bool) *GetAllInstanceIdListResponseBody {
	s.Success = &v
	return s
}

func (s *GetAllInstanceIdListResponseBody) Validate() error {
	return dara.Validate(s)
}
