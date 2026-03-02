// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterBucUserCmd interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *RegisterBucUserCmd
	GetEnterpriseId() *int64
	SetEnterpriseName(v string) *RegisterBucUserCmd
	GetEnterpriseName() *string
}

type RegisterBucUserCmd struct {
	EnterpriseId   *int64  `json:"enterpriseId,omitempty" xml:"enterpriseId,omitempty"`
	EnterpriseName *string `json:"enterpriseName,omitempty" xml:"enterpriseName,omitempty"`
}

func (s RegisterBucUserCmd) String() string {
	return dara.Prettify(s)
}

func (s RegisterBucUserCmd) GoString() string {
	return s.String()
}

func (s *RegisterBucUserCmd) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *RegisterBucUserCmd) GetEnterpriseName() *string {
	return s.EnterpriseName
}

func (s *RegisterBucUserCmd) SetEnterpriseId(v int64) *RegisterBucUserCmd {
	s.EnterpriseId = &v
	return s
}

func (s *RegisterBucUserCmd) SetEnterpriseName(v string) *RegisterBucUserCmd {
	s.EnterpriseName = &v
	return s
}

func (s *RegisterBucUserCmd) Validate() error {
	return dara.Validate(s)
}
