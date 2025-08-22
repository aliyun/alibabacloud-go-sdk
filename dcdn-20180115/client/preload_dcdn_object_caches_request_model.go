// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreloadDcdnObjectCachesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArea(v string) *PreloadDcdnObjectCachesRequest
	GetArea() *string
	SetL2Preload(v bool) *PreloadDcdnObjectCachesRequest
	GetL2Preload() *bool
	SetObjectPath(v string) *PreloadDcdnObjectCachesRequest
	GetObjectPath() *string
	SetOwnerId(v int64) *PreloadDcdnObjectCachesRequest
	GetOwnerId() *int64
	SetQueryHashkey(v bool) *PreloadDcdnObjectCachesRequest
	GetQueryHashkey() *bool
	SetSecurityToken(v string) *PreloadDcdnObjectCachesRequest
	GetSecurityToken() *string
	SetWithHeader(v string) *PreloadDcdnObjectCachesRequest
	GetWithHeader() *string
}

type PreloadDcdnObjectCachesRequest struct {
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
	// The path of the content that you want to prefetch. Separate multiple URLs with line feed characters (\\n) or a pair of carriage return and line feed characters (\\r\\n).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com/examplefile.txt
	ObjectPath    *string `json:"ObjectPath,omitempty" xml:"ObjectPath,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s PreloadDcdnObjectCachesRequest) String() string {
	return dara.Prettify(s)
}

func (s PreloadDcdnObjectCachesRequest) GoString() string {
	return s.String()
}

func (s *PreloadDcdnObjectCachesRequest) GetArea() *string {
	return s.Area
}

func (s *PreloadDcdnObjectCachesRequest) GetL2Preload() *bool {
	return s.L2Preload
}

func (s *PreloadDcdnObjectCachesRequest) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *PreloadDcdnObjectCachesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PreloadDcdnObjectCachesRequest) GetQueryHashkey() *bool {
	return s.QueryHashkey
}

func (s *PreloadDcdnObjectCachesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *PreloadDcdnObjectCachesRequest) GetWithHeader() *string {
	return s.WithHeader
}

func (s *PreloadDcdnObjectCachesRequest) SetArea(v string) *PreloadDcdnObjectCachesRequest {
	s.Area = &v
	return s
}

func (s *PreloadDcdnObjectCachesRequest) SetL2Preload(v bool) *PreloadDcdnObjectCachesRequest {
	s.L2Preload = &v
	return s
}

func (s *PreloadDcdnObjectCachesRequest) SetObjectPath(v string) *PreloadDcdnObjectCachesRequest {
	s.ObjectPath = &v
	return s
}

func (s *PreloadDcdnObjectCachesRequest) SetOwnerId(v int64) *PreloadDcdnObjectCachesRequest {
	s.OwnerId = &v
	return s
}

func (s *PreloadDcdnObjectCachesRequest) SetQueryHashkey(v bool) *PreloadDcdnObjectCachesRequest {
	s.QueryHashkey = &v
	return s
}

func (s *PreloadDcdnObjectCachesRequest) SetSecurityToken(v string) *PreloadDcdnObjectCachesRequest {
	s.SecurityToken = &v
	return s
}

func (s *PreloadDcdnObjectCachesRequest) SetWithHeader(v string) *PreloadDcdnObjectCachesRequest {
	s.WithHeader = &v
	return s
}

func (s *PreloadDcdnObjectCachesRequest) Validate() error {
	return dara.Validate(s)
}
