// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInsightSelectorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTrailName(v string) *GetInsightSelectorsRequest
	GetTrailName() *string
}

type GetInsightSelectorsRequest struct {
	// The name of the trail.
	//
	// This parameter is required.
	//
	// example:
	//
	// trail-name
	TrailName *string `json:"TrailName,omitempty" xml:"TrailName,omitempty"`
}

func (s GetInsightSelectorsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInsightSelectorsRequest) GoString() string {
	return s.String()
}

func (s *GetInsightSelectorsRequest) GetTrailName() *string {
	return s.TrailName
}

func (s *GetInsightSelectorsRequest) SetTrailName(v string) *GetInsightSelectorsRequest {
	s.TrailName = &v
	return s
}

func (s *GetInsightSelectorsRequest) Validate() error {
	return dara.Validate(s)
}
