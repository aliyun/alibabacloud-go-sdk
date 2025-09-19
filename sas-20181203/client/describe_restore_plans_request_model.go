// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestorePlansRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeRestorePlansRequest
	GetCurrentPage() *int32
	SetInstanceName(v string) *DescribeRestorePlansRequest
	GetInstanceName() *string
	SetPageSize(v int32) *DescribeRestorePlansRequest
	GetPageSize() *int32
	SetStatus(v string) *DescribeRestorePlansRequest
	GetStatus() *string
}

type DescribeRestorePlansRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// sql-test-001
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the restoration task. Valid values:
	//
	// 	- **init**: initializing
	//
	// 	- **created**: creating
	//
	// 	- **running**: running
	//
	// 	- **completed**: complete
	//
	// 	- **error**: failed
	//
	// 	- **restoring**: restoring
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeRestorePlansRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestorePlansRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestorePlansRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeRestorePlansRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeRestorePlansRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRestorePlansRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeRestorePlansRequest) SetCurrentPage(v int32) *DescribeRestorePlansRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRestorePlansRequest) SetInstanceName(v string) *DescribeRestorePlansRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeRestorePlansRequest) SetPageSize(v int32) *DescribeRestorePlansRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRestorePlansRequest) SetStatus(v string) *DescribeRestorePlansRequest {
	s.Status = &v
	return s
}

func (s *DescribeRestorePlansRequest) Validate() error {
	return dara.Validate(s)
}
