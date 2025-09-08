// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DescribeStorageResponseBody
	GetData() *bool
	SetRequestId(v string) *DescribeStorageResponseBody
	GetRequestId() *string
}

type DescribeStorageResponseBody struct {
	// Indicates whether the projects and Logstores that are created for the threat analysis feature exist in Simple Log Service. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CCEEE128-6607-503E-AAA6-C5E57D94****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStorageResponseBody) GetData() *bool {
	return s.Data
}

func (s *DescribeStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStorageResponseBody) SetData(v bool) *DescribeStorageResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeStorageResponseBody) SetRequestId(v string) *DescribeStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStorageResponseBody) Validate() error {
	return dara.Validate(s)
}
