// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfiguredDomainNamesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v []*DescribeConfiguredDomainNamesResponseBodyDomainNames) *DescribeConfiguredDomainNamesResponseBody
	GetDomainNames() []*DescribeConfiguredDomainNamesResponseBodyDomainNames
	SetModule(v string) *DescribeConfiguredDomainNamesResponseBody
	GetModule() *string
	SetRequestId(v string) *DescribeConfiguredDomainNamesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeConfiguredDomainNamesResponseBody
	GetTotalCount() *int32
}

type DescribeConfiguredDomainNamesResponseBody struct {
	DomainNames []*DescribeConfiguredDomainNamesResponseBodyDomainNames `json:"DomainNames,omitempty" xml:"DomainNames,omitempty" type:"Repeated"`
	// example:
	//
	// sg_server
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// example:
	//
	// 09A2D6F1-EA1B-56D9-977D-74878405****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 16
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeConfiguredDomainNamesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfiguredDomainNamesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfiguredDomainNamesResponseBody) GetDomainNames() []*DescribeConfiguredDomainNamesResponseBodyDomainNames {
	return s.DomainNames
}

func (s *DescribeConfiguredDomainNamesResponseBody) GetModule() *string {
	return s.Module
}

func (s *DescribeConfiguredDomainNamesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeConfiguredDomainNamesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeConfiguredDomainNamesResponseBody) SetDomainNames(v []*DescribeConfiguredDomainNamesResponseBodyDomainNames) *DescribeConfiguredDomainNamesResponseBody {
	s.DomainNames = v
	return s
}

func (s *DescribeConfiguredDomainNamesResponseBody) SetModule(v string) *DescribeConfiguredDomainNamesResponseBody {
	s.Module = &v
	return s
}

func (s *DescribeConfiguredDomainNamesResponseBody) SetRequestId(v string) *DescribeConfiguredDomainNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConfiguredDomainNamesResponseBody) SetTotalCount(v int32) *DescribeConfiguredDomainNamesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeConfiguredDomainNamesResponseBody) Validate() error {
	if s.DomainNames != nil {
		for _, item := range s.DomainNames {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeConfiguredDomainNamesResponseBodyDomainNames struct {
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 0
	IsMalicious *bool `json:"IsMalicious,omitempty" xml:"IsMalicious,omitempty"`
	// example:
	//
	// 1534408189
	OperationTime *int32 `json:"OperationTime,omitempty" xml:"OperationTime,omitempty"`
}

func (s DescribeConfiguredDomainNamesResponseBodyDomainNames) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfiguredDomainNamesResponseBodyDomainNames) GoString() string {
	return s.String()
}

func (s *DescribeConfiguredDomainNamesResponseBodyDomainNames) GetComment() *string {
	return s.Comment
}

func (s *DescribeConfiguredDomainNamesResponseBodyDomainNames) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeConfiguredDomainNamesResponseBodyDomainNames) GetIsMalicious() *bool {
	return s.IsMalicious
}

func (s *DescribeConfiguredDomainNamesResponseBodyDomainNames) GetOperationTime() *int32 {
	return s.OperationTime
}

func (s *DescribeConfiguredDomainNamesResponseBodyDomainNames) SetComment(v string) *DescribeConfiguredDomainNamesResponseBodyDomainNames {
	s.Comment = &v
	return s
}

func (s *DescribeConfiguredDomainNamesResponseBodyDomainNames) SetDomainName(v string) *DescribeConfiguredDomainNamesResponseBodyDomainNames {
	s.DomainName = &v
	return s
}

func (s *DescribeConfiguredDomainNamesResponseBodyDomainNames) SetIsMalicious(v bool) *DescribeConfiguredDomainNamesResponseBodyDomainNames {
	s.IsMalicious = &v
	return s
}

func (s *DescribeConfiguredDomainNamesResponseBodyDomainNames) SetOperationTime(v int32) *DescribeConfiguredDomainNamesResponseBodyDomainNames {
	s.OperationTime = &v
	return s
}

func (s *DescribeConfiguredDomainNamesResponseBodyDomainNames) Validate() error {
	return dara.Validate(s)
}
