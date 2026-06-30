// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCacheClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCacheClusters(v *CacheCluster) *GetCacheClusterResponseBody
	GetCacheClusters() *CacheCluster
	SetRequestId(v string) *GetCacheClusterResponseBody
	GetRequestId() *string
}

type GetCacheClusterResponseBody struct {
	// The list of Cache clusters.
	CacheClusters *CacheCluster `json:"cacheClusters,omitempty" xml:"cacheClusters,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetCacheClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCacheClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetCacheClusterResponseBody) GetCacheClusters() *CacheCluster {
	return s.CacheClusters
}

func (s *GetCacheClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCacheClusterResponseBody) SetCacheClusters(v *CacheCluster) *GetCacheClusterResponseBody {
	s.CacheClusters = v
	return s
}

func (s *GetCacheClusterResponseBody) SetRequestId(v string) *GetCacheClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCacheClusterResponseBody) Validate() error {
	if s.CacheClusters != nil {
		if err := s.CacheClusters.Validate(); err != nil {
			return err
		}
	}
	return nil
}
