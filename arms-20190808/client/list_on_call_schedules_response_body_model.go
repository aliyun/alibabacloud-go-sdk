// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOnCallSchedulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageBean(v *ListOnCallSchedulesResponseBodyPageBean) *ListOnCallSchedulesResponseBody
	GetPageBean() *ListOnCallSchedulesResponseBodyPageBean
	SetRequestId(v string) *ListOnCallSchedulesResponseBody
	GetRequestId() *string
}

type ListOnCallSchedulesResponseBody struct {
	// The objects that were returned.
	PageBean *ListOnCallSchedulesResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 78901766-3806-4E96-8E47-CFEF59E4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListOnCallSchedulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOnCallSchedulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOnCallSchedulesResponseBody) GetPageBean() *ListOnCallSchedulesResponseBodyPageBean {
	return s.PageBean
}

func (s *ListOnCallSchedulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOnCallSchedulesResponseBody) SetPageBean(v *ListOnCallSchedulesResponseBodyPageBean) *ListOnCallSchedulesResponseBody {
	s.PageBean = v
	return s
}

func (s *ListOnCallSchedulesResponseBody) SetRequestId(v string) *ListOnCallSchedulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOnCallSchedulesResponseBody) Validate() error {
	if s.PageBean != nil {
		if err := s.PageBean.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListOnCallSchedulesResponseBodyPageBean struct {
	// The information about the scheduling policy.
	OnCallSchedules []*ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules `json:"OnCallSchedules,omitempty" xml:"OnCallSchedules,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListOnCallSchedulesResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s ListOnCallSchedulesResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *ListOnCallSchedulesResponseBodyPageBean) GetOnCallSchedules() []*ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules {
	return s.OnCallSchedules
}

func (s *ListOnCallSchedulesResponseBodyPageBean) GetPage() *int64 {
	return s.Page
}

func (s *ListOnCallSchedulesResponseBodyPageBean) GetSize() *int64 {
	return s.Size
}

func (s *ListOnCallSchedulesResponseBodyPageBean) GetTotal() *int64 {
	return s.Total
}

func (s *ListOnCallSchedulesResponseBodyPageBean) SetOnCallSchedules(v []*ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules) *ListOnCallSchedulesResponseBodyPageBean {
	s.OnCallSchedules = v
	return s
}

func (s *ListOnCallSchedulesResponseBodyPageBean) SetPage(v int64) *ListOnCallSchedulesResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *ListOnCallSchedulesResponseBodyPageBean) SetSize(v int64) *ListOnCallSchedulesResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *ListOnCallSchedulesResponseBodyPageBean) SetTotal(v int64) *ListOnCallSchedulesResponseBodyPageBean {
	s.Total = &v
	return s
}

func (s *ListOnCallSchedulesResponseBodyPageBean) Validate() error {
	if s.OnCallSchedules != nil {
		for _, item := range s.OnCallSchedules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules struct {
	// The description of the scheduling policy.
	//
	// example:
	//
	// Test scheduling policy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the scheduling policy.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the scheduling policy.
	//
	// example:
	//
	// OnCallSchedule_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules) String() string {
	return dara.Prettify(s)
}

func (s ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules) GoString() string {
	return s.String()
}

func (s *ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules) GetDescription() *string {
	return s.Description
}

func (s *ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules) GetId() *int64 {
	return s.Id
}

func (s *ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules) GetName() *string {
	return s.Name
}

func (s *ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules) SetDescription(v string) *ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules {
	s.Description = &v
	return s
}

func (s *ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules) SetId(v int64) *ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules {
	s.Id = &v
	return s
}

func (s *ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules) SetName(v string) *ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules {
	s.Name = &v
	return s
}

func (s *ListOnCallSchedulesResponseBodyPageBeanOnCallSchedules) Validate() error {
	return dara.Validate(s)
}
