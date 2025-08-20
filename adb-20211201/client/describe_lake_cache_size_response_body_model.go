// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLakeCacheSizeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeLakeCacheSizeResponseBody
	GetCode() *int32
	SetData(v *DescribeLakeCacheSizeResponseBodyData) *DescribeLakeCacheSizeResponseBody
	GetData() *DescribeLakeCacheSizeResponseBodyData
	SetRequestId(v string) *DescribeLakeCacheSizeResponseBody
	GetRequestId() *string
}

type DescribeLakeCacheSizeResponseBody struct {
	// The status code. The value 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DescribeLakeCacheSizeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLakeCacheSizeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakeCacheSizeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLakeCacheSizeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeLakeCacheSizeResponseBody) GetData() *DescribeLakeCacheSizeResponseBodyData {
	return s.Data
}

func (s *DescribeLakeCacheSizeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLakeCacheSizeResponseBody) SetCode(v int32) *DescribeLakeCacheSizeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeLakeCacheSizeResponseBody) SetData(v *DescribeLakeCacheSizeResponseBodyData) *DescribeLakeCacheSizeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLakeCacheSizeResponseBody) SetRequestId(v string) *DescribeLakeCacheSizeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLakeCacheSizeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLakeCacheSizeResponseBodyData struct {
	// The size of the lake cache space. Unit: GB.
	//
	// example:
	//
	// 100
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// amv-bp10b6646l07akdt
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The size of the data that occupies the lake cache space. Unit: GB.
	//
	// example:
	//
	// 100
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// Indicates whether the lake cache feature is enabled.
	//
	// example:
	//
	// true
	EnableLakeCache *bool `json:"EnableLakeCache,omitempty" xml:"EnableLakeCache,omitempty"`
	// The clusters that share the lake cache space.
	Instances []*string `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
}

func (s DescribeLakeCacheSizeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakeCacheSizeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLakeCacheSizeResponseBodyData) GetCapacity() *int64 {
	return s.Capacity
}

func (s *DescribeLakeCacheSizeResponseBodyData) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeLakeCacheSizeResponseBodyData) GetDataSize() *int64 {
	return s.DataSize
}

func (s *DescribeLakeCacheSizeResponseBodyData) GetEnableLakeCache() *bool {
	return s.EnableLakeCache
}

func (s *DescribeLakeCacheSizeResponseBodyData) GetInstances() []*string {
	return s.Instances
}

func (s *DescribeLakeCacheSizeResponseBodyData) SetCapacity(v int64) *DescribeLakeCacheSizeResponseBodyData {
	s.Capacity = &v
	return s
}

func (s *DescribeLakeCacheSizeResponseBodyData) SetDBClusterId(v string) *DescribeLakeCacheSizeResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLakeCacheSizeResponseBodyData) SetDataSize(v int64) *DescribeLakeCacheSizeResponseBodyData {
	s.DataSize = &v
	return s
}

func (s *DescribeLakeCacheSizeResponseBodyData) SetEnableLakeCache(v bool) *DescribeLakeCacheSizeResponseBodyData {
	s.EnableLakeCache = &v
	return s
}

func (s *DescribeLakeCacheSizeResponseBodyData) SetInstances(v []*string) *DescribeLakeCacheSizeResponseBodyData {
	s.Instances = v
	return s
}

func (s *DescribeLakeCacheSizeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
