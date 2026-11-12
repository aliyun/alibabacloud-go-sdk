// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceVersionRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRemark(v string) *UpdateServiceVersionRemarkRequest
	GetRemark() *string
}

type UpdateServiceVersionRemarkRequest struct {
	// The remarks.
	//
	// example:
	//
	// version info.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateServiceVersionRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceVersionRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceVersionRemarkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateServiceVersionRemarkRequest) SetRemark(v string) *UpdateServiceVersionRemarkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateServiceVersionRemarkRequest) Validate() error {
	return dara.Validate(s)
}
