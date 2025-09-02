// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgWhiteListDeleteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v []*int32) *DsgWhiteListDeleteListRequest
	GetIds() []*int32
}

type DsgWhiteListDeleteListRequest struct {
	// The IDs of the whitelists.
	//
	// This parameter is required.
	Ids []*int32 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s DsgWhiteListDeleteListRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgWhiteListDeleteListRequest) GoString() string {
	return s.String()
}

func (s *DsgWhiteListDeleteListRequest) GetIds() []*int32 {
	return s.Ids
}

func (s *DsgWhiteListDeleteListRequest) SetIds(v []*int32) *DsgWhiteListDeleteListRequest {
	s.Ids = v
	return s
}

func (s *DsgWhiteListDeleteListRequest) Validate() error {
	return dara.Validate(s)
}
