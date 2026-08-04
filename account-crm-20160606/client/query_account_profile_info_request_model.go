// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountProfileInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHavanaId(v string) *QueryAccountProfileInfoRequest
	GetHavanaId() *string
	SetPK(v string) *QueryAccountProfileInfoRequest
	GetPK() *string
}

type QueryAccountProfileInfoRequest struct {
	HavanaId *string `json:"HavanaId,omitempty" xml:"HavanaId,omitempty"`
	PK       *string `json:"PK,omitempty" xml:"PK,omitempty"`
}

func (s QueryAccountProfileInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountProfileInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryAccountProfileInfoRequest) GetHavanaId() *string {
	return s.HavanaId
}

func (s *QueryAccountProfileInfoRequest) GetPK() *string {
	return s.PK
}

func (s *QueryAccountProfileInfoRequest) SetHavanaId(v string) *QueryAccountProfileInfoRequest {
	s.HavanaId = &v
	return s
}

func (s *QueryAccountProfileInfoRequest) SetPK(v string) *QueryAccountProfileInfoRequest {
	s.PK = &v
	return s
}

func (s *QueryAccountProfileInfoRequest) Validate() error {
	return dara.Validate(s)
}
