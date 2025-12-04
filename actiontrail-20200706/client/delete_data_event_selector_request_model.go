// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataEventSelectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTrailName(v string) *DeleteDataEventSelectorRequest
	GetTrailName() *string
}

type DeleteDataEventSelectorRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// trail-name
	TrailName *string `json:"TrailName,omitempty" xml:"TrailName,omitempty"`
}

func (s DeleteDataEventSelectorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataEventSelectorRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataEventSelectorRequest) GetTrailName() *string {
	return s.TrailName
}

func (s *DeleteDataEventSelectorRequest) SetTrailName(v string) *DeleteDataEventSelectorRequest {
	s.TrailName = &v
	return s
}

func (s *DeleteDataEventSelectorRequest) Validate() error {
	return dara.Validate(s)
}
