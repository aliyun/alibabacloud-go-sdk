// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEntityStoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *CreateEntityStoreRequest
	GetDryRun() *bool
}

type CreateEntityStoreRequest struct {
	// Specifies whether to perform a dry run. Valid values: true: performs a dry run without making actual changes. false: sends a normal request. Default value: false.
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateEntityStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEntityStoreRequest) GoString() string {
	return s.String()
}

func (s *CreateEntityStoreRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateEntityStoreRequest) SetDryRun(v bool) *CreateEntityStoreRequest {
	s.DryRun = &v
	return s
}

func (s *CreateEntityStoreRequest) Validate() error {
	return dara.Validate(s)
}
