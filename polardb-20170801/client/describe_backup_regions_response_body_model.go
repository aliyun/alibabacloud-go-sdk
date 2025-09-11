// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v []*string) *DescribeBackupRegionsResponseBody
	GetRegions() []*string
	SetRequestId(v string) *DescribeBackupRegionsResponseBody
	GetRequestId() *string
}

type DescribeBackupRegionsResponseBody struct {
	Regions []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// example:
	//
	// EB07CFF0-D8A4-5C76-AED7-D00E26FC2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupRegionsResponseBody) GetRegions() []*string {
	return s.Regions
}

func (s *DescribeBackupRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupRegionsResponseBody) SetRegions(v []*string) *DescribeBackupRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeBackupRegionsResponseBody) SetRequestId(v string) *DescribeBackupRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}
