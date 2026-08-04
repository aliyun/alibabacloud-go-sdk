// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountRealNameInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPK(v string) *QueryAccountRealNameInfoRequest
	GetPK() *string
}

type QueryAccountRealNameInfoRequest struct {
	PK *string `json:"PK,omitempty" xml:"PK,omitempty"`
}

func (s QueryAccountRealNameInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountRealNameInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryAccountRealNameInfoRequest) GetPK() *string {
	return s.PK
}

func (s *QueryAccountRealNameInfoRequest) SetPK(v string) *QueryAccountRealNameInfoRequest {
	s.PK = &v
	return s
}

func (s *QueryAccountRealNameInfoRequest) Validate() error {
	return dara.Validate(s)
}
