// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCsrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCsrId(v int64) *DeleteCsrRequest
	GetCsrId() *int64
}

type DeleteCsrRequest struct {
	// The ID of the CSR.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3013
	CsrId *int64 `json:"CsrId,omitempty" xml:"CsrId,omitempty"`
}

func (s DeleteCsrRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCsrRequest) GoString() string {
	return s.String()
}

func (s *DeleteCsrRequest) GetCsrId() *int64 {
	return s.CsrId
}

func (s *DeleteCsrRequest) SetCsrId(v int64) *DeleteCsrRequest {
	s.CsrId = &v
	return s
}

func (s *DeleteCsrRequest) Validate() error {
	return dara.Validate(s)
}
