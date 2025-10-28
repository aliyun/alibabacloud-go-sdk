// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEccInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEccId(v string) *QueryEccInfoRequest
	GetEccId() *string
}

type QueryEccInfoRequest struct {
	// The ID of the ECC.
	//
	// This parameter is required.
	//
	// example:
	//
	// b197-40ab-9155-****
	EccId *string `json:"EccId,omitempty" xml:"EccId,omitempty"`
}

func (s QueryEccInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryEccInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryEccInfoRequest) GetEccId() *string {
	return s.EccId
}

func (s *QueryEccInfoRequest) SetEccId(v string) *QueryEccInfoRequest {
	s.EccId = &v
	return s
}

func (s *QueryEccInfoRequest) Validate() error {
	return dara.Validate(s)
}
