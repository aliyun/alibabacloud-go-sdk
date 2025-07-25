// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmInstanceSystemCnameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeGtmInstanceSystemCnameRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeGtmInstanceSystemCnameRequest
	GetLang() *string
}

type DescribeGtmInstanceSystemCnameRequest struct {
	// The ID of the Global Traffic Manager (GTM) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language used by the user.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeGtmInstanceSystemCnameRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstanceSystemCnameRequest) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceSystemCnameRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGtmInstanceSystemCnameRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGtmInstanceSystemCnameRequest) SetInstanceId(v string) *DescribeGtmInstanceSystemCnameRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGtmInstanceSystemCnameRequest) SetLang(v string) *DescribeGtmInstanceSystemCnameRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGtmInstanceSystemCnameRequest) Validate() error {
	return dara.Validate(s)
}
