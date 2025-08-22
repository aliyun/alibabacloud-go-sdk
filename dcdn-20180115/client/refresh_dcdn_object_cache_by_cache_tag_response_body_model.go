// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshDcdnObjectCacheByCacheTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRefreshTaskId(v string) *RefreshDcdnObjectCacheByCacheTagResponseBody
	GetRefreshTaskId() *string
	SetRequestId(v string) *RefreshDcdnObjectCacheByCacheTagResponseBody
	GetRequestId() *string
}

type RefreshDcdnObjectCacheByCacheTagResponseBody struct {
	// example:
	//
	// 17410889914
	RefreshTaskId *string `json:"RefreshTaskId,omitempty" xml:"RefreshTaskId,omitempty"`
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshDcdnObjectCacheByCacheTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshDcdnObjectCacheByCacheTagResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshDcdnObjectCacheByCacheTagResponseBody) GetRefreshTaskId() *string {
	return s.RefreshTaskId
}

func (s *RefreshDcdnObjectCacheByCacheTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshDcdnObjectCacheByCacheTagResponseBody) SetRefreshTaskId(v string) *RefreshDcdnObjectCacheByCacheTagResponseBody {
	s.RefreshTaskId = &v
	return s
}

func (s *RefreshDcdnObjectCacheByCacheTagResponseBody) SetRequestId(v string) *RefreshDcdnObjectCacheByCacheTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshDcdnObjectCacheByCacheTagResponseBody) Validate() error {
	return dara.Validate(s)
}
