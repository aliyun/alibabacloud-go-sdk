// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgScenedDeleteSceneShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdsShrink(v string) *DsgScenedDeleteSceneShrinkRequest
	GetIdsShrink() *string
}

type DsgScenedDeleteSceneShrinkRequest struct {
	// The IDs of level-2 data masking scenarios.
	//
	// This parameter is required.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s DsgScenedDeleteSceneShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgScenedDeleteSceneShrinkRequest) GoString() string {
	return s.String()
}

func (s *DsgScenedDeleteSceneShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *DsgScenedDeleteSceneShrinkRequest) SetIdsShrink(v string) *DsgScenedDeleteSceneShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *DsgScenedDeleteSceneShrinkRequest) Validate() error {
	return dara.Validate(s)
}
