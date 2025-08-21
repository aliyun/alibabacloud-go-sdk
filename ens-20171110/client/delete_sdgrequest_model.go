// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSDGRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSDGId(v []*string) *DeleteSDGRequest
	GetSDGId() []*string
}

type DeleteSDGRequest struct {
	// The IDs of the SDGs that you want to delete.
	//
	// This parameter is required.
	SDGId []*string `json:"SDGId,omitempty" xml:"SDGId,omitempty" type:"Repeated"`
}

func (s DeleteSDGRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSDGRequest) GoString() string {
	return s.String()
}

func (s *DeleteSDGRequest) GetSDGId() []*string {
	return s.SDGId
}

func (s *DeleteSDGRequest) SetSDGId(v []*string) *DeleteSDGRequest {
	s.SDGId = v
	return s
}

func (s *DeleteSDGRequest) Validate() error {
	return dara.Validate(s)
}
