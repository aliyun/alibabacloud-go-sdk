// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindFinanceTaxDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKpId(v int64) *FindFinanceTaxDetailRequest
	GetKpId() *int64
}

type FindFinanceTaxDetailRequest struct {
	// This parameter is required.
	KpId *int64 `json:"KpId,omitempty" xml:"KpId,omitempty"`
}

func (s FindFinanceTaxDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s FindFinanceTaxDetailRequest) GoString() string {
	return s.String()
}

func (s *FindFinanceTaxDetailRequest) GetKpId() *int64 {
	return s.KpId
}

func (s *FindFinanceTaxDetailRequest) SetKpId(v int64) *FindFinanceTaxDetailRequest {
	s.KpId = &v
	return s
}

func (s *FindFinanceTaxDetailRequest) Validate() error {
	return dara.Validate(s)
}
