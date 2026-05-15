// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigNumListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *GetConfigNumListRequest
	GetAccountName() *string
	SetDepartmentId(v int64) *GetConfigNumListRequest
	GetDepartmentId() *int64
	SetInstanceId(v string) *GetConfigNumListRequest
	GetInstanceId() *string
}

type GetConfigNumListRequest struct {
	// example:
	//
	// 123@123.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// 12345
	DepartmentId *int64 `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetConfigNumListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConfigNumListRequest) GoString() string {
	return s.String()
}

func (s *GetConfigNumListRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *GetConfigNumListRequest) GetDepartmentId() *int64 {
	return s.DepartmentId
}

func (s *GetConfigNumListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetConfigNumListRequest) SetAccountName(v string) *GetConfigNumListRequest {
	s.AccountName = &v
	return s
}

func (s *GetConfigNumListRequest) SetDepartmentId(v int64) *GetConfigNumListRequest {
	s.DepartmentId = &v
	return s
}

func (s *GetConfigNumListRequest) SetInstanceId(v string) *GetConfigNumListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetConfigNumListRequest) Validate() error {
	return dara.Validate(s)
}
