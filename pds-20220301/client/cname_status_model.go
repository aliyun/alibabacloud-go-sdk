// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCNameStatus interface {
	dara.Model
	String() string
	GoString() string
	SetBingdingState(v string) *CNameStatus
	GetBingdingState() *string
	SetLegalState(v string) *CNameStatus
	GetLegalState() *string
	SetRemark(v string) *CNameStatus
	GetRemark() *string
}

type CNameStatus struct {
	// example:
	//
	// BINDING/BOUND
	BingdingState *string `json:"bingding_state,omitempty" xml:"bingding_state,omitempty"`
	// example:
	//
	// NORMAL/ABNORMAL
	LegalState *string `json:"legal_state,omitempty" xml:"legal_state,omitempty"`
	// example:
	//
	// beian
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s CNameStatus) String() string {
	return dara.Prettify(s)
}

func (s CNameStatus) GoString() string {
	return s.String()
}

func (s *CNameStatus) GetBingdingState() *string {
	return s.BingdingState
}

func (s *CNameStatus) GetLegalState() *string {
	return s.LegalState
}

func (s *CNameStatus) GetRemark() *string {
	return s.Remark
}

func (s *CNameStatus) SetBingdingState(v string) *CNameStatus {
	s.BingdingState = &v
	return s
}

func (s *CNameStatus) SetLegalState(v string) *CNameStatus {
	s.LegalState = &v
	return s
}

func (s *CNameStatus) SetRemark(v string) *CNameStatus {
	s.Remark = &v
	return s
}

func (s *CNameStatus) Validate() error {
	return dara.Validate(s)
}
