// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetAreaCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudGetAreaCodeRequest
	GetEnterpriseId() *int64
	SetTel(v string) *CloudGetAreaCodeRequest
	GetTel() *string
}

type CloudGetAreaCodeRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 11位手机号；只能查询手机号
	//
	// This parameter is required.
	//
	// example:
	//
	// 17750247753
	Tel *string `json:"Tel,omitempty" xml:"Tel,omitempty"`
}

func (s CloudGetAreaCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAreaCodeRequest) GoString() string {
	return s.String()
}

func (s *CloudGetAreaCodeRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudGetAreaCodeRequest) GetTel() *string {
	return s.Tel
}

func (s *CloudGetAreaCodeRequest) SetEnterpriseId(v int64) *CloudGetAreaCodeRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudGetAreaCodeRequest) SetTel(v string) *CloudGetAreaCodeRequest {
	s.Tel = &v
	return s
}

func (s *CloudGetAreaCodeRequest) Validate() error {
	return dara.Validate(s)
}
