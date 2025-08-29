// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrowdsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCrowds(v []*ListCrowdsResponseBodyCrowds) *ListCrowdsResponseBody
	GetCrowds() []*ListCrowdsResponseBodyCrowds
	SetRequestId(v string) *ListCrowdsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListCrowdsResponseBody
	GetTotalCount() *int64
}

type ListCrowdsResponseBody struct {
	Crowds []*ListCrowdsResponseBodyCrowds `json:"Crowds,omitempty" xml:"Crowds,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 9763624B-5FBB-5E3A-9193-B1ADB554CEAE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCrowdsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCrowdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCrowdsResponseBody) GetCrowds() []*ListCrowdsResponseBodyCrowds {
	return s.Crowds
}

func (s *ListCrowdsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCrowdsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCrowdsResponseBody) SetCrowds(v []*ListCrowdsResponseBodyCrowds) *ListCrowdsResponseBody {
	s.Crowds = v
	return s
}

func (s *ListCrowdsResponseBody) SetRequestId(v string) *ListCrowdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCrowdsResponseBody) SetTotalCount(v int64) *ListCrowdsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCrowdsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCrowdsResponseBodyCrowds struct {
	// example:
	//
	// 3
	CrowdId *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// os=android
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// crowd1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 10
	Quantity *string `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// example:
	//
	// ManualInput
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// user1,user2
	Users *string `json:"Users,omitempty" xml:"Users,omitempty"`
}

func (s ListCrowdsResponseBodyCrowds) String() string {
	return dara.Prettify(s)
}

func (s ListCrowdsResponseBodyCrowds) GoString() string {
	return s.String()
}

func (s *ListCrowdsResponseBodyCrowds) GetCrowdId() *string {
	return s.CrowdId
}

func (s *ListCrowdsResponseBodyCrowds) GetDescription() *string {
	return s.Description
}

func (s *ListCrowdsResponseBodyCrowds) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListCrowdsResponseBodyCrowds) GetLabel() *string {
	return s.Label
}

func (s *ListCrowdsResponseBodyCrowds) GetName() *string {
	return s.Name
}

func (s *ListCrowdsResponseBodyCrowds) GetQuantity() *string {
	return s.Quantity
}

func (s *ListCrowdsResponseBodyCrowds) GetSource() *string {
	return s.Source
}

func (s *ListCrowdsResponseBodyCrowds) GetUsers() *string {
	return s.Users
}

func (s *ListCrowdsResponseBodyCrowds) SetCrowdId(v string) *ListCrowdsResponseBodyCrowds {
	s.CrowdId = &v
	return s
}

func (s *ListCrowdsResponseBodyCrowds) SetDescription(v string) *ListCrowdsResponseBodyCrowds {
	s.Description = &v
	return s
}

func (s *ListCrowdsResponseBodyCrowds) SetGmtCreateTime(v string) *ListCrowdsResponseBodyCrowds {
	s.GmtCreateTime = &v
	return s
}

func (s *ListCrowdsResponseBodyCrowds) SetLabel(v string) *ListCrowdsResponseBodyCrowds {
	s.Label = &v
	return s
}

func (s *ListCrowdsResponseBodyCrowds) SetName(v string) *ListCrowdsResponseBodyCrowds {
	s.Name = &v
	return s
}

func (s *ListCrowdsResponseBodyCrowds) SetQuantity(v string) *ListCrowdsResponseBodyCrowds {
	s.Quantity = &v
	return s
}

func (s *ListCrowdsResponseBodyCrowds) SetSource(v string) *ListCrowdsResponseBodyCrowds {
	s.Source = &v
	return s
}

func (s *ListCrowdsResponseBodyCrowds) SetUsers(v string) *ListCrowdsResponseBodyCrowds {
	s.Users = &v
	return s
}

func (s *ListCrowdsResponseBodyCrowds) Validate() error {
	return dara.Validate(s)
}
