// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeys(v []*string) *DeleteServiceLabelRequest
	GetKeys() []*string
}

type DeleteServiceLabelRequest struct {
	// The service tags that you want to delete.
	//
	// This parameter is required.
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
}

func (s DeleteServiceLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceLabelRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceLabelRequest) GetKeys() []*string {
	return s.Keys
}

func (s *DeleteServiceLabelRequest) SetKeys(v []*string) *DeleteServiceLabelRequest {
	s.Keys = v
	return s
}

func (s *DeleteServiceLabelRequest) Validate() error {
	return dara.Validate(s)
}
