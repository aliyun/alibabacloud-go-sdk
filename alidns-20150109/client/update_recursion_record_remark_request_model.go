// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionRecordRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateRecursionRecordRemarkRequest
	GetClientToken() *string
	SetRecordId(v string) *UpdateRecursionRecordRemarkRequest
	GetRecordId() *string
	SetRemark(v string) *UpdateRecursionRecordRemarkRequest
	GetRemark() *string
}

type UpdateRecursionRecordRemarkRequest struct {
	// example:
	//
	// e432232342423ew423
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// record id
	//
	// example:
	//
	// 173671468000010
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateRecursionRecordRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionRecordRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecursionRecordRemarkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRecursionRecordRemarkRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *UpdateRecursionRecordRemarkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateRecursionRecordRemarkRequest) SetClientToken(v string) *UpdateRecursionRecordRemarkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRecursionRecordRemarkRequest) SetRecordId(v string) *UpdateRecursionRecordRemarkRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateRecursionRecordRemarkRequest) SetRemark(v string) *UpdateRecursionRecordRemarkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateRecursionRecordRemarkRequest) Validate() error {
	return dara.Validate(s)
}
