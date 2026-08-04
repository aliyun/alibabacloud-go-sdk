// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoadRealNameInfoByPkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPK(v string) *LoadRealNameInfoByPkRequest
	GetPK() *string
}

type LoadRealNameInfoByPkRequest struct {
	// This parameter is required.
	PK *string `json:"PK,omitempty" xml:"PK,omitempty"`
}

func (s LoadRealNameInfoByPkRequest) String() string {
	return dara.Prettify(s)
}

func (s LoadRealNameInfoByPkRequest) GoString() string {
	return s.String()
}

func (s *LoadRealNameInfoByPkRequest) GetPK() *string {
	return s.PK
}

func (s *LoadRealNameInfoByPkRequest) SetPK(v string) *LoadRealNameInfoByPkRequest {
	s.PK = &v
	return s
}

func (s *LoadRealNameInfoByPkRequest) Validate() error {
	return dara.Validate(s)
}
