// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushObjectCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPushTaskId(v string) *PushObjectCacheResponseBody
	GetPushTaskId() *string
	SetRequestId(v string) *PushObjectCacheResponseBody
	GetRequestId() *string
}

type PushObjectCacheResponseBody struct {
	// The ID of the prefetch task. If multiple tasks are returned, the IDs are separated by commas (,). The task IDs are merged based on the following rules:
	//
	// 	- If the tasks are set for the same accelerated domain name, submitted within the same second, and prefetch content from URLs instead of directories, the tasks IDs are merged into the same task ID (RushTaskId).
	//
	// 	- If the number of tasks that are set for the same accelerated domain name, submitted within the same second, and prefetch content from URLs instead of directories exceeds 500, every 500 task IDs are merged into the same task ID (RushTaskId).
	//
	// example:
	//
	// 9524xxxx
	PushTaskId *string `json:"PushTaskId,omitempty" xml:"PushTaskId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushObjectCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushObjectCacheResponseBody) GoString() string {
	return s.String()
}

func (s *PushObjectCacheResponseBody) GetPushTaskId() *string {
	return s.PushTaskId
}

func (s *PushObjectCacheResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushObjectCacheResponseBody) SetPushTaskId(v string) *PushObjectCacheResponseBody {
	s.PushTaskId = &v
	return s
}

func (s *PushObjectCacheResponseBody) SetRequestId(v string) *PushObjectCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushObjectCacheResponseBody) Validate() error {
	return dara.Validate(s)
}
