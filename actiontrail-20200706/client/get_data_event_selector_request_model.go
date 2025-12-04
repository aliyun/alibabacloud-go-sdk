// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataEventSelectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTrailName(v string) *GetDataEventSelectorRequest
	GetTrailName() *string
}

type GetDataEventSelectorRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// trail-name
	TrailName *string `json:"TrailName,omitempty" xml:"TrailName,omitempty"`
}

func (s GetDataEventSelectorRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataEventSelectorRequest) GoString() string {
	return s.String()
}

func (s *GetDataEventSelectorRequest) GetTrailName() *string {
	return s.TrailName
}

func (s *GetDataEventSelectorRequest) SetTrailName(v string) *GetDataEventSelectorRequest {
	s.TrailName = &v
	return s
}

func (s *GetDataEventSelectorRequest) Validate() error {
	return dara.Validate(s)
}
