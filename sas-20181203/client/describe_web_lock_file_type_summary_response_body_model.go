// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockFileTypeSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*DescribeWebLockFileTypeSummaryResponseBodyList) *DescribeWebLockFileTypeSummaryResponseBody
	GetList() []*DescribeWebLockFileTypeSummaryResponseBodyList
	SetRequestId(v string) *DescribeWebLockFileTypeSummaryResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeWebLockFileTypeSummaryResponseBody
	GetTotalCount() *int32
}

type DescribeWebLockFileTypeSummaryResponseBody struct {
	// An array that consists of events on web tamper proofing returned.
	List []*DescribeWebLockFileTypeSummaryResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4BB99533-4FDC-5B9C-A5E4-5AE3E9BE5C78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of events on web tamper proofing.
	//
	// example:
	//
	// 639
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWebLockFileTypeSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockFileTypeSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebLockFileTypeSummaryResponseBody) GetList() []*DescribeWebLockFileTypeSummaryResponseBodyList {
	return s.List
}

func (s *DescribeWebLockFileTypeSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebLockFileTypeSummaryResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeWebLockFileTypeSummaryResponseBody) SetList(v []*DescribeWebLockFileTypeSummaryResponseBodyList) *DescribeWebLockFileTypeSummaryResponseBody {
	s.List = v
	return s
}

func (s *DescribeWebLockFileTypeSummaryResponseBody) SetRequestId(v string) *DescribeWebLockFileTypeSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebLockFileTypeSummaryResponseBody) SetTotalCount(v int32) *DescribeWebLockFileTypeSummaryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebLockFileTypeSummaryResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWebLockFileTypeSummaryResponseBodyList struct {
	// The number of attempts.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The type of the protected file. Valid values:
	//
	// 	- **php**: PHP file
	//
	// 	- **jsp**: JSP file
	//
	// 	- **asp**: ASP file
	//
	// 	- **aspx**: ASPX file
	//
	// 	- **js**: JS file
	//
	// 	- **cgi**: CGI file
	//
	// 	- **html**: HTML file
	//
	// 	- **htm**: HTM file
	//
	// 	- **xml**: XML file
	//
	// 	- **shtml**: SHTML file
	//
	// 	- **shtm**: SHTM file
	//
	// 	- **jpg**: JPG file
	//
	// 	- **gif**: GIF file
	//
	// 	- **png**: PNG file
	//
	// example:
	//
	// jsp
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeWebLockFileTypeSummaryResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockFileTypeSummaryResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeWebLockFileTypeSummaryResponseBodyList) GetCount() *int32 {
	return s.Count
}

func (s *DescribeWebLockFileTypeSummaryResponseBodyList) GetType() *string {
	return s.Type
}

func (s *DescribeWebLockFileTypeSummaryResponseBodyList) SetCount(v int32) *DescribeWebLockFileTypeSummaryResponseBodyList {
	s.Count = &v
	return s
}

func (s *DescribeWebLockFileTypeSummaryResponseBodyList) SetType(v string) *DescribeWebLockFileTypeSummaryResponseBodyList {
	s.Type = &v
	return s
}

func (s *DescribeWebLockFileTypeSummaryResponseBodyList) Validate() error {
	return dara.Validate(s)
}
