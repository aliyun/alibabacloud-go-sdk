// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *DeleteServiceRequest
	GetDryRun() *bool
}

type DeleteServiceRequest struct {
	// Specifies whether to perform a dry run of the request. Valid values:
	//
	// - true: Sends a check request without making actual changes.
	//
	// - false: Sends a normal request.
	//
	// Default value: false.
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s DeleteServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteServiceRequest) SetDryRun(v bool) *DeleteServiceRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteServiceRequest) Validate() error {
	return dara.Validate(s)
}
