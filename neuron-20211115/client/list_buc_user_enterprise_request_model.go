// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBucUserEnterpriseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmpId(v string) *ListBucUserEnterpriseRequest
	GetEmpId() *string
}

type ListBucUserEnterpriseRequest struct {
	EmpId *string `json:"empId,omitempty" xml:"empId,omitempty"`
}

func (s ListBucUserEnterpriseRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBucUserEnterpriseRequest) GoString() string {
	return s.String()
}

func (s *ListBucUserEnterpriseRequest) GetEmpId() *string {
	return s.EmpId
}

func (s *ListBucUserEnterpriseRequest) SetEmpId(v string) *ListBucUserEnterpriseRequest {
	s.EmpId = &v
	return s
}

func (s *ListBucUserEnterpriseRequest) Validate() error {
	return dara.Validate(s)
}
