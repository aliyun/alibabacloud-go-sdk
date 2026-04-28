// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvisorChecksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeAdvisorChecksResponseBody
	GetCode() *string
	SetData(v *DescribeAdvisorChecksResponseBodyData) *DescribeAdvisorChecksResponseBody
	GetData() *DescribeAdvisorChecksResponseBodyData
	SetRequestId(v string) *DescribeAdvisorChecksResponseBody
	GetRequestId() *string
}

type DescribeAdvisorChecksResponseBody struct {
	// example:
	//
	// 200
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeAdvisorChecksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 566331F9-5AB3-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAdvisorChecksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvisorChecksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAdvisorChecksResponseBody) GetData() *DescribeAdvisorChecksResponseBodyData {
	return s.Data
}

func (s *DescribeAdvisorChecksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAdvisorChecksResponseBody) SetCode(v string) *DescribeAdvisorChecksResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBody) SetData(v *DescribeAdvisorChecksResponseBodyData) *DescribeAdvisorChecksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAdvisorChecksResponseBody) SetRequestId(v string) *DescribeAdvisorChecksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAdvisorChecksResponseBodyData struct {
	AdvisorCheck []*DescribeAdvisorChecksResponseBodyDataAdvisorCheck `json:"AdvisorCheck,omitempty" xml:"AdvisorCheck,omitempty" type:"Repeated"`
}

func (s DescribeAdvisorChecksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvisorChecksResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksResponseBodyData) GetAdvisorCheck() []*DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	return s.AdvisorCheck
}

func (s *DescribeAdvisorChecksResponseBodyData) SetAdvisorCheck(v []*DescribeAdvisorChecksResponseBodyDataAdvisorCheck) *DescribeAdvisorChecksResponseBodyData {
	s.AdvisorCheck = v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyData) Validate() error {
	if s.AdvisorCheck != nil {
		for _, item := range s.AdvisorCheck {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAdvisorChecksResponseBodyDataAdvisorCheck struct {
	Category      *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreated    *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified   *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OperateColumn *string `json:"OperateColumn,omitempty" xml:"OperateColumn,omitempty"`
	Product       *string `json:"Product,omitempty" xml:"Product,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Tips          *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
	ViewColumn    *string `json:"ViewColumn,omitempty" xml:"ViewColumn,omitempty"`
}

func (s DescribeAdvisorChecksResponseBodyDataAdvisorCheck) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvisorChecksResponseBodyDataAdvisorCheck) GoString() string {
	return s.String()
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) GetCategory() *string {
	return s.Category
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) GetCode() *string {
	return s.Code
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) GetDescription() *string {
	return s.Description
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) GetName() *string {
	return s.Name
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) GetOperateColumn() *string {
	return s.OperateColumn
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) GetProduct() *string {
	return s.Product
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) GetStatus() *string {
	return s.Status
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) GetTips() *string {
	return s.Tips
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) GetViewColumn() *string {
	return s.ViewColumn
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetCategory(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.Category = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetCode(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.Code = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetDescription(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.Description = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetGmtCreated(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.GmtCreated = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetGmtModified(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.GmtModified = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetName(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.Name = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetOperateColumn(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.OperateColumn = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetProduct(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.Product = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetStatus(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.Status = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetTips(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.Tips = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) SetViewColumn(v string) *DescribeAdvisorChecksResponseBodyDataAdvisorCheck {
	s.ViewColumn = &v
	return s
}

func (s *DescribeAdvisorChecksResponseBodyDataAdvisorCheck) Validate() error {
	return dara.Validate(s)
}
