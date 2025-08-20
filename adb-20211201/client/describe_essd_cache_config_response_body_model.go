// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEssdCacheConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeEssdCacheConfigResponseBodyData) *DescribeEssdCacheConfigResponseBody
	GetData() *DescribeEssdCacheConfigResponseBodyData
	SetRequestId(v string) *DescribeEssdCacheConfigResponseBody
	GetRequestId() *string
}

type DescribeEssdCacheConfigResponseBody struct {
	// The returned data.
	Data *DescribeEssdCacheConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C7EDB8E4-9769-4233-88C7-DCA4C9******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEssdCacheConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEssdCacheConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEssdCacheConfigResponseBody) GetData() *DescribeEssdCacheConfigResponseBodyData {
	return s.Data
}

func (s *DescribeEssdCacheConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEssdCacheConfigResponseBody) SetData(v *DescribeEssdCacheConfigResponseBodyData) *DescribeEssdCacheConfigResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEssdCacheConfigResponseBody) SetRequestId(v string) *DescribeEssdCacheConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEssdCacheConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEssdCacheConfigResponseBodyData struct {
	// Specifies whether to enable the disk cache feature.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableEssdCache *bool `json:"EnableEssdCache,omitempty" xml:"EnableEssdCache,omitempty"`
	// The disk cache size. Unit: GB.
	//
	// example:
	//
	// 500
	EssdCacheSize *int32 `json:"EssdCacheSize,omitempty" xml:"EssdCacheSize,omitempty"`
}

func (s DescribeEssdCacheConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEssdCacheConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEssdCacheConfigResponseBodyData) GetEnableEssdCache() *bool {
	return s.EnableEssdCache
}

func (s *DescribeEssdCacheConfigResponseBodyData) GetEssdCacheSize() *int32 {
	return s.EssdCacheSize
}

func (s *DescribeEssdCacheConfigResponseBodyData) SetEnableEssdCache(v bool) *DescribeEssdCacheConfigResponseBodyData {
	s.EnableEssdCache = &v
	return s
}

func (s *DescribeEssdCacheConfigResponseBodyData) SetEssdCacheSize(v int32) *DescribeEssdCacheConfigResponseBodyData {
	s.EssdCacheSize = &v
	return s
}

func (s *DescribeEssdCacheConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
