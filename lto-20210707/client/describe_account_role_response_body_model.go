// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeAccountRoleResponseBody
	GetCode() *string
	SetData(v *DescribeAccountRoleResponseBodyData) *DescribeAccountRoleResponseBody
	GetData() *DescribeAccountRoleResponseBodyData
	SetHttpStatusCode(v int32) *DescribeAccountRoleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeAccountRoleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAccountRoleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAccountRoleResponseBody
	GetSuccess() *bool
}

type DescribeAccountRoleResponseBody struct {
	Code           *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DescribeAccountRoleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAccountRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountRoleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAccountRoleResponseBody) GetData() *DescribeAccountRoleResponseBodyData {
	return s.Data
}

func (s *DescribeAccountRoleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeAccountRoleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAccountRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccountRoleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAccountRoleResponseBody) SetCode(v string) *DescribeAccountRoleResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAccountRoleResponseBody) SetData(v *DescribeAccountRoleResponseBodyData) *DescribeAccountRoleResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAccountRoleResponseBody) SetHttpStatusCode(v int32) *DescribeAccountRoleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAccountRoleResponseBody) SetMessage(v string) *DescribeAccountRoleResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAccountRoleResponseBody) SetRequestId(v string) *DescribeAccountRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountRoleResponseBody) SetSuccess(v bool) *DescribeAccountRoleResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAccountRoleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccountRoleResponseBodyData struct {
	AuthorizedBaaS *bool   `json:"AuthorizedBaaS,omitempty" xml:"AuthorizedBaaS,omitempty"`
	RoleType       *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeAccountRoleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountRoleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAccountRoleResponseBodyData) GetAuthorizedBaaS() *bool {
	return s.AuthorizedBaaS
}

func (s *DescribeAccountRoleResponseBodyData) GetRoleType() *string {
	return s.RoleType
}

func (s *DescribeAccountRoleResponseBodyData) SetAuthorizedBaaS(v bool) *DescribeAccountRoleResponseBodyData {
	s.AuthorizedBaaS = &v
	return s
}

func (s *DescribeAccountRoleResponseBodyData) SetRoleType(v string) *DescribeAccountRoleResponseBodyData {
	s.RoleType = &v
	return s
}

func (s *DescribeAccountRoleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
