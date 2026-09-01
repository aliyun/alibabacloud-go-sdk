// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDistinctReleasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecords(v []*DescribeDistinctReleasesResponseBodyRecords) *DescribeDistinctReleasesResponseBody
	GetRecords() []*DescribeDistinctReleasesResponseBodyRecords
	SetRequestId(v string) *DescribeDistinctReleasesResponseBody
	GetRequestId() *string
}

type DescribeDistinctReleasesResponseBody struct {
	// The version records.
	Records []*DescribeDistinctReleasesResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The ID of the request. This ID is unique to each request and is used to troubleshoot issues.
	//
	// example:
	//
	// 145CACF6-D276-5197-8549-CB1AD76E2AC8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDistinctReleasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDistinctReleasesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDistinctReleasesResponseBody) GetRecords() []*DescribeDistinctReleasesResponseBodyRecords {
	return s.Records
}

func (s *DescribeDistinctReleasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDistinctReleasesResponseBody) SetRecords(v []*DescribeDistinctReleasesResponseBodyRecords) *DescribeDistinctReleasesResponseBody {
	s.Records = v
	return s
}

func (s *DescribeDistinctReleasesResponseBody) SetRequestId(v string) *DescribeDistinctReleasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDistinctReleasesResponseBody) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDistinctReleasesResponseBodyRecords struct {
	// The description of the version.
	//
	// example:
	//
	// demo version
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The MD5 value of the version\\"s XML configuration.
	//
	// example:
	//
	// 17cf53049bc8efa941207xxxxx
	TaskflowMd5 *string `json:"TaskflowMd5,omitempty" xml:"TaskflowMd5,omitempty"`
	// The configuration format of the playbook.
	//
	// - **xml**: XML format.
	//
	// - **x6**: JSON format.
	//
	// example:
	//
	// x6
	TaskflowType *string `json:"TaskflowType,omitempty" xml:"TaskflowType,omitempty"`
}

func (s DescribeDistinctReleasesResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeDistinctReleasesResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeDistinctReleasesResponseBodyRecords) GetDescription() *string {
	return s.Description
}

func (s *DescribeDistinctReleasesResponseBodyRecords) GetTaskflowMd5() *string {
	return s.TaskflowMd5
}

func (s *DescribeDistinctReleasesResponseBodyRecords) GetTaskflowType() *string {
	return s.TaskflowType
}

func (s *DescribeDistinctReleasesResponseBodyRecords) SetDescription(v string) *DescribeDistinctReleasesResponseBodyRecords {
	s.Description = &v
	return s
}

func (s *DescribeDistinctReleasesResponseBodyRecords) SetTaskflowMd5(v string) *DescribeDistinctReleasesResponseBodyRecords {
	s.TaskflowMd5 = &v
	return s
}

func (s *DescribeDistinctReleasesResponseBodyRecords) SetTaskflowType(v string) *DescribeDistinctReleasesResponseBodyRecords {
	s.TaskflowType = &v
	return s
}

func (s *DescribeDistinctReleasesResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}
