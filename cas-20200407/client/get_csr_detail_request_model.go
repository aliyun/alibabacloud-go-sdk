// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCsrDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCsrId(v int64) *GetCsrDetailRequest
	GetCsrId() *int64
}

type GetCsrDetailRequest struct {
	// The ID of the CSR.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3924
	CsrId *int64 `json:"CsrId,omitempty" xml:"CsrId,omitempty"`
}

func (s GetCsrDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCsrDetailRequest) GoString() string {
	return s.String()
}

func (s *GetCsrDetailRequest) GetCsrId() *int64 {
	return s.CsrId
}

func (s *GetCsrDetailRequest) SetCsrId(v int64) *GetCsrDetailRequest {
	s.CsrId = &v
	return s
}

func (s *GetCsrDetailRequest) Validate() error {
	return dara.Validate(s)
}
