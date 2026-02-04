// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshErObjectCachesByCacheTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRefreshTaskId(v string) *RefreshErObjectCachesByCacheTagResponseBody
	GetRefreshTaskId() *string
	SetRequestId(v string) *RefreshErObjectCachesByCacheTagResponseBody
	GetRequestId() *string
}

type RefreshErObjectCachesByCacheTagResponseBody struct {
	// example:
	//
	// 17410889914
	RefreshTaskId *string `json:"RefreshTaskId,omitempty" xml:"RefreshTaskId,omitempty"`
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshErObjectCachesByCacheTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshErObjectCachesByCacheTagResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshErObjectCachesByCacheTagResponseBody) GetRefreshTaskId() *string {
	return s.RefreshTaskId
}

func (s *RefreshErObjectCachesByCacheTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshErObjectCachesByCacheTagResponseBody) SetRefreshTaskId(v string) *RefreshErObjectCachesByCacheTagResponseBody {
	s.RefreshTaskId = &v
	return s
}

func (s *RefreshErObjectCachesByCacheTagResponseBody) SetRequestId(v string) *RefreshErObjectCachesByCacheTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshErObjectCachesByCacheTagResponseBody) Validate() error {
	return dara.Validate(s)
}
