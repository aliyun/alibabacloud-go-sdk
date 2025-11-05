// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCreditInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUid(v int64) *GetCreditInfoRequest
	GetUid() *int64
}

type GetCreditInfoRequest struct {
	// Sub Account UID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1792155717328010
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetCreditInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCreditInfoRequest) GoString() string {
	return s.String()
}

func (s *GetCreditInfoRequest) GetUid() *int64 {
	return s.Uid
}

func (s *GetCreditInfoRequest) SetUid(v int64) *GetCreditInfoRequest {
	s.Uid = &v
	return s
}

func (s *GetCreditInfoRequest) Validate() error {
	return dara.Validate(s)
}
