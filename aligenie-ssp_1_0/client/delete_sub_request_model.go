// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSubId(v int64) *DeleteSubRequest
	GetSubId() *int64
}

type DeleteSubRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 81
	SubId *int64 `json:"SubId,omitempty" xml:"SubId,omitempty"`
}

func (s DeleteSubRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubRequest) GetSubId() *int64 {
	return s.SubId
}

func (s *DeleteSubRequest) SetSubId(v int64) *DeleteSubRequest {
	s.SubId = &v
	return s
}

func (s *DeleteSubRequest) Validate() error {
	return dara.Validate(s)
}
