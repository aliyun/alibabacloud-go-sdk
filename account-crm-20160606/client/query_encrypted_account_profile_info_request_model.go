// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEncryptedAccountProfileInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHavanaId(v string) *QueryEncryptedAccountProfileInfoRequest
	GetHavanaId() *string
	SetPK(v string) *QueryEncryptedAccountProfileInfoRequest
	GetPK() *string
}

type QueryEncryptedAccountProfileInfoRequest struct {
	HavanaId *string `json:"HavanaId,omitempty" xml:"HavanaId,omitempty"`
	PK       *string `json:"PK,omitempty" xml:"PK,omitempty"`
}

func (s QueryEncryptedAccountProfileInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryEncryptedAccountProfileInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryEncryptedAccountProfileInfoRequest) GetHavanaId() *string {
	return s.HavanaId
}

func (s *QueryEncryptedAccountProfileInfoRequest) GetPK() *string {
	return s.PK
}

func (s *QueryEncryptedAccountProfileInfoRequest) SetHavanaId(v string) *QueryEncryptedAccountProfileInfoRequest {
	s.HavanaId = &v
	return s
}

func (s *QueryEncryptedAccountProfileInfoRequest) SetPK(v string) *QueryEncryptedAccountProfileInfoRequest {
	s.PK = &v
	return s
}

func (s *QueryEncryptedAccountProfileInfoRequest) Validate() error {
	return dara.Validate(s)
}
