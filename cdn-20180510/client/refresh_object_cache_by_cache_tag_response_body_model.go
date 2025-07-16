// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshObjectCacheByCacheTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRefreshTaskId(v string) *RefreshObjectCacheByCacheTagResponseBody
	GetRefreshTaskId() *string
	SetRequestId(v string) *RefreshObjectCacheByCacheTagResponseBody
	GetRequestId() *string
}

type RefreshObjectCacheByCacheTagResponseBody struct {
	// The ID of the refresh task. If multiple tasks are returned, the task IDs are separated by commas (,). The task IDs are merged based on the following rules:
	//
	// 	- If the tasks are specified for the same accelerated domain name, submitted within the same second, and run to refresh content based on URLs instead of directories, the task IDs are merged into one task ID (RefreshTaskId).
	//
	// 	- If the number of tasks that are specified for the same accelerated domain name, submitted within the same second, and run to refresh content based on URLs instead of directories exceeds 2,000, every 2,000 task IDs are merged into one task ID (RefreshTaskId).
	//
	// example:
	//
	// 17772470184
	RefreshTaskId *string `json:"RefreshTaskId,omitempty" xml:"RefreshTaskId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2E5AD83F-BD7B-462E-8319-2E30E305519A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshObjectCacheByCacheTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshObjectCacheByCacheTagResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshObjectCacheByCacheTagResponseBody) GetRefreshTaskId() *string {
	return s.RefreshTaskId
}

func (s *RefreshObjectCacheByCacheTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshObjectCacheByCacheTagResponseBody) SetRefreshTaskId(v string) *RefreshObjectCacheByCacheTagResponseBody {
	s.RefreshTaskId = &v
	return s
}

func (s *RefreshObjectCacheByCacheTagResponseBody) SetRequestId(v string) *RefreshObjectCacheByCacheTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshObjectCacheByCacheTagResponseBody) Validate() error {
	return dara.Validate(s)
}
