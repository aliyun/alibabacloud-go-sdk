// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOSSStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetColdStorage(v bool) *DescribeOSSStorageResponseBody
	GetColdStorage() *bool
	SetPolicy(v string) *DescribeOSSStorageResponseBody
	GetPolicy() *string
	SetRequestId(v string) *DescribeOSSStorageResponseBody
	GetRequestId() *string
	SetState(v string) *DescribeOSSStorageResponseBody
	GetState() *string
	SetStorageUsage(v string) *DescribeOSSStorageResponseBody
	GetStorageUsage() *string
}

type DescribeOSSStorageResponseBody struct {
	// Indicates whether tiered storage for hot and cold data can be enabled. Valid values:
	//
	// - **true**: Tiered storage can be enabled.
	//
	// - **false**: Tiered storage cannot be enabled.
	//
	// example:
	//
	// true
	ColdStorage *bool `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	// The parameters of the tiered storage policy for hot and cold data.
	//
	// example:
	//
	// [{ "currentValue":"0.1", "defaultValue":"0.1", "desc":"Ratio of free disk space. When the ratio exceeds the value of configuration parameter, ClickHouse start to move data to the cold storage", "name":"move_factor", "restart":true, "valueRange":"(0, 1]" },{ "currentValue":"true", "defaultValue":"true", "desc":"Disables merging of data parts on cold storage", "name":"prefer_not_to_merge", "restart":true, "valueRange":"true|false" }]
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The request ID.
	//
	// example:
	//
	// aadbb456-ebf7-4ed8-9671-fad9f3846ca4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of tiered storage for hot and cold data. Valid values:
	//
	// - **CREATING**: Tiered storage is being enabled.
	//
	// - **DISABLE**: Tiered storage is disabled.
	//
	// - **ENABLE**: Tiered storage is enabled.
	//
	// example:
	//
	// ENABLE
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The cold storage space used. Unit: GB.
	//
	// example:
	//
	// 0.00
	StorageUsage *string `json:"StorageUsage,omitempty" xml:"StorageUsage,omitempty"`
}

func (s DescribeOSSStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOSSStorageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOSSStorageResponseBody) GetColdStorage() *bool {
	return s.ColdStorage
}

func (s *DescribeOSSStorageResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeOSSStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOSSStorageResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeOSSStorageResponseBody) GetStorageUsage() *string {
	return s.StorageUsage
}

func (s *DescribeOSSStorageResponseBody) SetColdStorage(v bool) *DescribeOSSStorageResponseBody {
	s.ColdStorage = &v
	return s
}

func (s *DescribeOSSStorageResponseBody) SetPolicy(v string) *DescribeOSSStorageResponseBody {
	s.Policy = &v
	return s
}

func (s *DescribeOSSStorageResponseBody) SetRequestId(v string) *DescribeOSSStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOSSStorageResponseBody) SetState(v string) *DescribeOSSStorageResponseBody {
	s.State = &v
	return s
}

func (s *DescribeOSSStorageResponseBody) SetStorageUsage(v string) *DescribeOSSStorageResponseBody {
	s.StorageUsage = &v
	return s
}

func (s *DescribeOSSStorageResponseBody) Validate() error {
	return dara.Validate(s)
}
