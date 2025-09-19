// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileProtectRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *UpdateFileProtectRemarkRequest
	GetId() *int64
	SetRemark(v []*string) *UpdateFileProtectRemarkRequest
	GetRemark() []*string
}

type UpdateFileProtectRemarkRequest struct {
	// The ID of the event.
	//
	// example:
	//
	// 1764
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The remarks.
	Remark []*string `json:"Remark,omitempty" xml:"Remark,omitempty" type:"Repeated"`
}

func (s UpdateFileProtectRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileProtectRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileProtectRemarkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateFileProtectRemarkRequest) GetRemark() []*string {
	return s.Remark
}

func (s *UpdateFileProtectRemarkRequest) SetId(v int64) *UpdateFileProtectRemarkRequest {
	s.Id = &v
	return s
}

func (s *UpdateFileProtectRemarkRequest) SetRemark(v []*string) *UpdateFileProtectRemarkRequest {
	s.Remark = v
	return s
}

func (s *UpdateFileProtectRemarkRequest) Validate() error {
	return dara.Validate(s)
}
