// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessKeyLeakDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DescribeAccessKeyLeakDetailRequest
	GetId() *int64
	SetResourceDirectoryAccountId(v int64) *DescribeAccessKeyLeakDetailRequest
	GetResourceDirectoryAccountId() *int64
}

type DescribeAccessKeyLeakDetailRequest struct {
	// The ID of the AccessKey pair leak event.
	//
	// > You can call the [DescribeAccesskeyLeakList](~~DescribeAccesskeyLeakList~~) operation to obtain the event ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 389357
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to query the ID.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
}

func (s DescribeAccessKeyLeakDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessKeyLeakDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessKeyLeakDetailRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeAccessKeyLeakDetailRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeAccessKeyLeakDetailRequest) SetId(v int64) *DescribeAccessKeyLeakDetailRequest {
	s.Id = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailRequest) SetResourceDirectoryAccountId(v int64) *DescribeAccessKeyLeakDetailRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeAccessKeyLeakDetailRequest) Validate() error {
	return dara.Validate(s)
}
