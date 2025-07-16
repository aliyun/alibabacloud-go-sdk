// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v map[string]*string) *UpdateServiceLabelRequest
	GetLabels() map[string]*string
}

type UpdateServiceLabelRequest struct {
	// The custom service tags.
	//
	// This parameter is required.
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s UpdateServiceLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceLabelRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceLabelRequest) GetLabels() map[string]*string {
	return s.Labels
}

func (s *UpdateServiceLabelRequest) SetLabels(v map[string]*string) *UpdateServiceLabelRequest {
	s.Labels = v
	return s
}

func (s *UpdateServiceLabelRequest) Validate() error {
	return dara.Validate(s)
}
