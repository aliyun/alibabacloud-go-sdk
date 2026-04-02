// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateDatasetRequest
	GetDescription() *string
}

type UpdateDatasetRequest struct {
	// example:
	//
	// test dataset
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDatasetRequest) SetDescription(v string) *UpdateDatasetRequest {
	s.Description = &v
	return s
}

func (s *UpdateDatasetRequest) Validate() error {
	return dara.Validate(s)
}
