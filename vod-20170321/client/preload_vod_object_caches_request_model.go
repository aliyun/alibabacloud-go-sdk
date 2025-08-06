// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreloadVodObjectCachesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArea(v string) *PreloadVodObjectCachesRequest
	GetArea() *string
	SetL2Preload(v bool) *PreloadVodObjectCachesRequest
	GetL2Preload() *bool
	SetObjectPath(v string) *PreloadVodObjectCachesRequest
	GetObjectPath() *string
	SetOwnerId(v int64) *PreloadVodObjectCachesRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *PreloadVodObjectCachesRequest
	GetSecurityToken() *string
	SetWithHeader(v string) *PreloadVodObjectCachesRequest
	GetWithHeader() *string
}

type PreloadVodObjectCachesRequest struct {
	// The acceleration region in which you want to prefetch content. If you do not specify a region, the value overseas is used.
	//
	// 	- **domestic**: Chinese mainland
	//
	// 	- **overseas**: outside the Chinese mainland
	//
	// example:
	//
	// domestic
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// Specifies whether to prefetch content to POPs. Valid values:
	//
	// 	- **true**: prefetches content to nodes that include L2 DCDN nodes.
	//
	// 	- **false**: prefetches content to L2 POPs or L3 POPs.
	//
	// example:
	//
	// true
	L2Preload *bool `json:"L2Preload,omitempty" xml:"L2Preload,omitempty"`
	// The URL of the file to be prefetched. Separate multiple URLs with line breaks (\\n or \\r\\n).
	//
	// This parameter is required.
	//
	// example:
	//
	// vod.test.com/test.txt
	ObjectPath    *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The custom header for prefetch in the JSON format.
	//
	// example:
	//
	// {
	//
	//       "Accept-Encoding": [
	//
	//             "gzip, deflate, br"
	//
	//       ]
	//
	// }
	WithHeader *string `json:"WithHeader,omitempty" xml:"WithHeader,omitempty"`
}

func (s PreloadVodObjectCachesRequest) String() string {
	return dara.Prettify(s)
}

func (s PreloadVodObjectCachesRequest) GoString() string {
	return s.String()
}

func (s *PreloadVodObjectCachesRequest) GetArea() *string {
	return s.Area
}

func (s *PreloadVodObjectCachesRequest) GetL2Preload() *bool {
	return s.L2Preload
}

func (s *PreloadVodObjectCachesRequest) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *PreloadVodObjectCachesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PreloadVodObjectCachesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *PreloadVodObjectCachesRequest) GetWithHeader() *string {
	return s.WithHeader
}

func (s *PreloadVodObjectCachesRequest) SetArea(v string) *PreloadVodObjectCachesRequest {
	s.Area = &v
	return s
}

func (s *PreloadVodObjectCachesRequest) SetL2Preload(v bool) *PreloadVodObjectCachesRequest {
	s.L2Preload = &v
	return s
}

func (s *PreloadVodObjectCachesRequest) SetObjectPath(v string) *PreloadVodObjectCachesRequest {
	s.ObjectPath = &v
	return s
}

func (s *PreloadVodObjectCachesRequest) SetOwnerId(v int64) *PreloadVodObjectCachesRequest {
	s.OwnerId = &v
	return s
}

func (s *PreloadVodObjectCachesRequest) SetSecurityToken(v string) *PreloadVodObjectCachesRequest {
	s.SecurityToken = &v
	return s
}

func (s *PreloadVodObjectCachesRequest) SetWithHeader(v string) *PreloadVodObjectCachesRequest {
	s.WithHeader = &v
	return s
}

func (s *PreloadVodObjectCachesRequest) Validate() error {
	return dara.Validate(s)
}
