// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushObjectCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArea(v string) *PushObjectCacheRequest
	GetArea() *string
	SetConsistencyHash(v bool) *PushObjectCacheRequest
	GetConsistencyHash() *bool
	SetL2Preload(v bool) *PushObjectCacheRequest
	GetL2Preload() *bool
	SetObjectPath(v string) *PushObjectCacheRequest
	GetObjectPath() *string
	SetOwnerId(v int64) *PushObjectCacheRequest
	GetOwnerId() *int64
	SetQueryHashkey(v bool) *PushObjectCacheRequest
	GetQueryHashkey() *bool
	SetSecurityToken(v string) *PushObjectCacheRequest
	GetSecurityToken() *string
	SetWithHeader(v string) *PushObjectCacheRequest
	GetWithHeader() *string
}

type PushObjectCacheRequest struct {
	// The acceleration region where content is to be prefetched. Valid values:
	//
	// 	- **domestic****: Chinese mainland**
	//
	// 	- **overseas****: regions outside the Chinese mainland**
	//
	// If you do not set this parameter, content in the service location (specified by parameter Coverage) that you configured for the domain is prefetched. Content is prefetched based on the following rules:
	//
	// 	- If the acceleration region is set to **Chinese Mainland Only**, content in regions in the Chinese mainland is prefetched.
	//
	// 	- If the acceleration region is set to **Global**, content in all regions is prefetched.
	//
	// 	- If the acceleration region is set to **Global (Excluding the Chinese Mainland)**, content in regions outside the Chinese mainland is prefetched.
	//
	// example:
	//
	// domestic
	Area            *string `json:"Area,omitempty" xml:"Area,omitempty"`
	ConsistencyHash *bool   `json:"ConsistencyHash,omitempty" xml:"ConsistencyHash,omitempty"`
	// Specifies whether to prefetch content to L2 points of presence (POPs). Valid values:
	//
	// 	- **true**: prefetches content to L2 POPs.
	//
	// 	- **false**: prefetches content to regular POPs. Regular POPs can be L2 POPs or L3 POPs. Default value: **false**.
	//
	// example:
	//
	// true
	L2Preload *bool `json:"L2Preload,omitempty" xml:"L2Preload,omitempty"`
	// The URLs based on which content is prefetched. Format: **accelerated domain name/files to be prefetched**.
	//
	// > Separate URLs with line feeds (\\n or \\r\\n). Each object path can be up to 1,024 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com/image/1.png\\nexample.org/image/2.png
	ObjectPath *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter specifies whether to enable the hash key query mode when you run a prefetch task. Valid values:
	//
	// 	- false: The default mode, in which the submitted URL is used as the hash key for the prefetch.
	//
	// 	- true: In this mode, the actual hash key used for the prefetch is queried based on the configuration of the domain name.
	//
	// example:
	//
	// true
	QueryHashkey  *bool   `json:"QueryHashkey,omitempty" xml:"QueryHashkey,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The custom header for prefetch in the JSON format.
	//
	// example:
	//
	// {
	//
	//       "Accept-Encoding": [
	//
	//             "gzip"
	//
	//       ]
	//
	// }
	WithHeader *string `json:"WithHeader,omitempty" xml:"WithHeader,omitempty"`
}

func (s PushObjectCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s PushObjectCacheRequest) GoString() string {
	return s.String()
}

func (s *PushObjectCacheRequest) GetArea() *string {
	return s.Area
}

func (s *PushObjectCacheRequest) GetConsistencyHash() *bool {
	return s.ConsistencyHash
}

func (s *PushObjectCacheRequest) GetL2Preload() *bool {
	return s.L2Preload
}

func (s *PushObjectCacheRequest) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *PushObjectCacheRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PushObjectCacheRequest) GetQueryHashkey() *bool {
	return s.QueryHashkey
}

func (s *PushObjectCacheRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *PushObjectCacheRequest) GetWithHeader() *string {
	return s.WithHeader
}

func (s *PushObjectCacheRequest) SetArea(v string) *PushObjectCacheRequest {
	s.Area = &v
	return s
}

func (s *PushObjectCacheRequest) SetConsistencyHash(v bool) *PushObjectCacheRequest {
	s.ConsistencyHash = &v
	return s
}

func (s *PushObjectCacheRequest) SetL2Preload(v bool) *PushObjectCacheRequest {
	s.L2Preload = &v
	return s
}

func (s *PushObjectCacheRequest) SetObjectPath(v string) *PushObjectCacheRequest {
	s.ObjectPath = &v
	return s
}

func (s *PushObjectCacheRequest) SetOwnerId(v int64) *PushObjectCacheRequest {
	s.OwnerId = &v
	return s
}

func (s *PushObjectCacheRequest) SetQueryHashkey(v bool) *PushObjectCacheRequest {
	s.QueryHashkey = &v
	return s
}

func (s *PushObjectCacheRequest) SetSecurityToken(v string) *PushObjectCacheRequest {
	s.SecurityToken = &v
	return s
}

func (s *PushObjectCacheRequest) SetWithHeader(v string) *PushObjectCacheRequest {
	s.WithHeader = &v
	return s
}

func (s *PushObjectCacheRequest) Validate() error {
	return dara.Validate(s)
}
