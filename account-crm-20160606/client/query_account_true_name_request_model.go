// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountTrueNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHavanaId(v string) *QueryAccountTrueNameRequest
	GetHavanaId() *string
	SetPK(v string) *QueryAccountTrueNameRequest
	GetPK() *string
}

type QueryAccountTrueNameRequest struct {
	HavanaId *string `json:"HavanaId,omitempty" xml:"HavanaId,omitempty"`
	PK       *string `json:"PK,omitempty" xml:"PK,omitempty"`
}

func (s QueryAccountTrueNameRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountTrueNameRequest) GoString() string {
	return s.String()
}

func (s *QueryAccountTrueNameRequest) GetHavanaId() *string {
	return s.HavanaId
}

func (s *QueryAccountTrueNameRequest) GetPK() *string {
	return s.PK
}

func (s *QueryAccountTrueNameRequest) SetHavanaId(v string) *QueryAccountTrueNameRequest {
	s.HavanaId = &v
	return s
}

func (s *QueryAccountTrueNameRequest) SetPK(v string) *QueryAccountTrueNameRequest {
	s.PK = &v
	return s
}

func (s *QueryAccountTrueNameRequest) Validate() error {
	return dara.Validate(s)
}
