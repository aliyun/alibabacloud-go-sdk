// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceTestCaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteServiceTestCaseRequest
	GetRegionId() *string
	SetTestCaseId(v string) *DeleteServiceTestCaseRequest
	GetTestCaseId() *string
}

type DeleteServiceTestCaseRequest struct {
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service test case id.
	//
	// This parameter is required.
	//
	// example:
	//
	// stc-0b2a3ad7e1de4c299eec
	TestCaseId *string `json:"TestCaseId,omitempty" xml:"TestCaseId,omitempty"`
}

func (s DeleteServiceTestCaseRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceTestCaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceTestCaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteServiceTestCaseRequest) GetTestCaseId() *string {
	return s.TestCaseId
}

func (s *DeleteServiceTestCaseRequest) SetRegionId(v string) *DeleteServiceTestCaseRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServiceTestCaseRequest) SetTestCaseId(v string) *DeleteServiceTestCaseRequest {
	s.TestCaseId = &v
	return s
}

func (s *DeleteServiceTestCaseRequest) Validate() error {
	return dara.Validate(s)
}
