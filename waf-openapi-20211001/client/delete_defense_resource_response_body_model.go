// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDefenseResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDefenseResourceResponseBody
	GetRequestId() *string
}

type DeleteDefenseResourceResponseBody struct {
	// example:
	//
	// 745F051D-95FD-57EC-9DC1-79BB4883C6A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDefenseResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDefenseResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDefenseResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDefenseResourceResponseBody) SetRequestId(v string) *DeleteDefenseResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDefenseResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
