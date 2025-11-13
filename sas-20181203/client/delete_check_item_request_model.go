// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCheckItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckIds(v []*int64) *DeleteCheckItemRequest
	GetCheckIds() []*int64
}

type DeleteCheckItemRequest struct {
	// List of check item IDs.
	//
	// This parameter is required.
	CheckIds []*int64 `json:"CheckIds,omitempty" xml:"CheckIds,omitempty" type:"Repeated"`
}

func (s DeleteCheckItemRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCheckItemRequest) GoString() string {
	return s.String()
}

func (s *DeleteCheckItemRequest) GetCheckIds() []*int64 {
	return s.CheckIds
}

func (s *DeleteCheckItemRequest) SetCheckIds(v []*int64) *DeleteCheckItemRequest {
	s.CheckIds = v
	return s
}

func (s *DeleteCheckItemRequest) Validate() error {
	return dara.Validate(s)
}
