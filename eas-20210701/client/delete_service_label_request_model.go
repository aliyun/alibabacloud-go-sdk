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
	SetLabelKeys(v []*string) *DeleteServiceLabelRequest
	GetLabelKeys() []*string
}

type DeleteServiceLabelRequest struct {
	// Deprecated
	//
	// The keys of the service labels that you want to delete. This parameter is deprecated. Use the LabelKeys parameter instead.
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// The keys of the service labels that you want to delete.
	LabelKeys []*string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty" type:"Repeated"`
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

func (s *DeleteServiceLabelRequest) GetLabelKeys() []*string {
	return s.LabelKeys
}

func (s *DeleteServiceLabelRequest) SetKeys(v []*string) *DeleteServiceLabelRequest {
	s.Keys = v
	return s
}

func (s *DeleteServiceLabelRequest) SetLabelKeys(v []*string) *DeleteServiceLabelRequest {
	s.LabelKeys = v
	return s
}

func (s *DeleteServiceLabelRequest) Validate() error {
	return dara.Validate(s)
}
