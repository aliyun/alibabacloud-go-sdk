// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelSelectionDelta interface {
	dara.Model
	String() string
	GoString() string
	SetAddModelCardIds(v []*string) *ModelSelectionDelta
	GetAddModelCardIds() []*string
	SetRemoveModelCardIds(v []*string) *ModelSelectionDelta
	GetRemoveModelCardIds() []*string
}

type ModelSelectionDelta struct {
	// The actual source model IDs to add to the unique ModelAPI. Already selected IDs must not be added again.
	AddModelCardIds []*string `json:"addModelCardIds,omitempty" xml:"addModelCardIds,omitempty" type:"Repeated"`
	// The actual source model IDs to remove from the unique ModelAPI. This operation does not delete models in the source. You cannot clear the API or remove backup models that are still referenced by other entry points.
	RemoveModelCardIds []*string `json:"removeModelCardIds,omitempty" xml:"removeModelCardIds,omitempty" type:"Repeated"`
}

func (s ModelSelectionDelta) String() string {
	return dara.Prettify(s)
}

func (s ModelSelectionDelta) GoString() string {
	return s.String()
}

func (s *ModelSelectionDelta) GetAddModelCardIds() []*string {
	return s.AddModelCardIds
}

func (s *ModelSelectionDelta) GetRemoveModelCardIds() []*string {
	return s.RemoveModelCardIds
}

func (s *ModelSelectionDelta) SetAddModelCardIds(v []*string) *ModelSelectionDelta {
	s.AddModelCardIds = v
	return s
}

func (s *ModelSelectionDelta) SetRemoveModelCardIds(v []*string) *ModelSelectionDelta {
	s.RemoveModelCardIds = v
	return s
}

func (s *ModelSelectionDelta) Validate() error {
	return dara.Validate(s)
}
