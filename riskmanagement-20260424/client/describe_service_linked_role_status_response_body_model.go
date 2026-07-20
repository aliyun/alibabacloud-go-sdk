// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceLinkedRoleStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeServiceLinkedRoleStatusResponseBody
	GetCode() *string
	SetData(v *DescribeServiceLinkedRoleStatusResponseBodyData) *DescribeServiceLinkedRoleStatusResponseBody
	GetData() *DescribeServiceLinkedRoleStatusResponseBodyData
	SetMessage(v string) *DescribeServiceLinkedRoleStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeServiceLinkedRoleStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeServiceLinkedRoleStatusResponseBody
	GetSuccess() *bool
}

type DescribeServiceLinkedRoleStatusResponseBody struct {
	// example:
	//
	// 200
	Code *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeServiceLinkedRoleStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EF972A16-95FB-5EF2-9CED-208A74DEF040
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeServiceLinkedRoleStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceLinkedRoleStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceLinkedRoleStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeServiceLinkedRoleStatusResponseBody) GetData() *DescribeServiceLinkedRoleStatusResponseBodyData {
	return s.Data
}

func (s *DescribeServiceLinkedRoleStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeServiceLinkedRoleStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceLinkedRoleStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeServiceLinkedRoleStatusResponseBody) SetCode(v string) *DescribeServiceLinkedRoleStatusResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeServiceLinkedRoleStatusResponseBody) SetData(v *DescribeServiceLinkedRoleStatusResponseBodyData) *DescribeServiceLinkedRoleStatusResponseBody {
	s.Data = v
	return s
}

func (s *DescribeServiceLinkedRoleStatusResponseBody) SetMessage(v string) *DescribeServiceLinkedRoleStatusResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeServiceLinkedRoleStatusResponseBody) SetRequestId(v string) *DescribeServiceLinkedRoleStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceLinkedRoleStatusResponseBody) SetSuccess(v bool) *DescribeServiceLinkedRoleStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeServiceLinkedRoleStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceLinkedRoleStatusResponseBodyData struct {
	Body *DescribeServiceLinkedRoleStatusResponseBodyDataBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
}

func (s DescribeServiceLinkedRoleStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceLinkedRoleStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeServiceLinkedRoleStatusResponseBodyData) GetBody() *DescribeServiceLinkedRoleStatusResponseBodyDataBody {
	return s.Body
}

func (s *DescribeServiceLinkedRoleStatusResponseBodyData) SetBody(v *DescribeServiceLinkedRoleStatusResponseBodyDataBody) *DescribeServiceLinkedRoleStatusResponseBodyData {
	s.Body = v
	return s
}

func (s *DescribeServiceLinkedRoleStatusResponseBodyData) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceLinkedRoleStatusResponseBodyDataBody struct {
	// example:
	//
	// 7F14E3C8-A6AA-5D3C-B7E0-ABA2AC171EFC
	RequestId  *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RoleStatus *DescribeServiceLinkedRoleStatusResponseBodyDataBodyRoleStatus `json:"RoleStatus,omitempty" xml:"RoleStatus,omitempty" type:"Struct"`
}

func (s DescribeServiceLinkedRoleStatusResponseBodyDataBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceLinkedRoleStatusResponseBodyDataBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceLinkedRoleStatusResponseBodyDataBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceLinkedRoleStatusResponseBodyDataBody) GetRoleStatus() *DescribeServiceLinkedRoleStatusResponseBodyDataBodyRoleStatus {
	return s.RoleStatus
}

func (s *DescribeServiceLinkedRoleStatusResponseBodyDataBody) SetRequestId(v string) *DescribeServiceLinkedRoleStatusResponseBodyDataBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceLinkedRoleStatusResponseBodyDataBody) SetRoleStatus(v *DescribeServiceLinkedRoleStatusResponseBodyDataBodyRoleStatus) *DescribeServiceLinkedRoleStatusResponseBodyDataBody {
	s.RoleStatus = v
	return s
}

func (s *DescribeServiceLinkedRoleStatusResponseBodyDataBody) Validate() error {
	if s.RoleStatus != nil {
		if err := s.RoleStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceLinkedRoleStatusResponseBodyDataBodyRoleStatus struct {
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeServiceLinkedRoleStatusResponseBodyDataBodyRoleStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceLinkedRoleStatusResponseBodyDataBodyRoleStatus) GoString() string {
	return s.String()
}

func (s *DescribeServiceLinkedRoleStatusResponseBodyDataBodyRoleStatus) GetStatus() *bool {
	return s.Status
}

func (s *DescribeServiceLinkedRoleStatusResponseBodyDataBodyRoleStatus) SetStatus(v bool) *DescribeServiceLinkedRoleStatusResponseBodyDataBodyRoleStatus {
	s.Status = &v
	return s
}

func (s *DescribeServiceLinkedRoleStatusResponseBodyDataBodyRoleStatus) Validate() error {
	return dara.Validate(s)
}
