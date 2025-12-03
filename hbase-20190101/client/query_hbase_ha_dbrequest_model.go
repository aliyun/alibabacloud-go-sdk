// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHBaseHaDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBdsId(v string) *QueryHBaseHaDBRequest
	GetBdsId() *string
}

type QueryHBaseHaDBRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// bds-t4n3496whj23ia4k
	BdsId *string `json:"BdsId,omitempty" xml:"BdsId,omitempty"`
}

func (s QueryHBaseHaDBRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryHBaseHaDBRequest) GoString() string {
	return s.String()
}

func (s *QueryHBaseHaDBRequest) GetBdsId() *string {
	return s.BdsId
}

func (s *QueryHBaseHaDBRequest) SetBdsId(v string) *QueryHBaseHaDBRequest {
	s.BdsId = &v
	return s
}

func (s *QueryHBaseHaDBRequest) Validate() error {
	return dara.Validate(s)
}
