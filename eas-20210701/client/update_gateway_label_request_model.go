// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v map[string]*string) *UpdateGatewayLabelRequest
	GetLabels() map[string]*string
}

type UpdateGatewayLabelRequest struct {
	// This parameter is required.
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s UpdateGatewayLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayLabelRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayLabelRequest) GetLabels() map[string]*string {
	return s.Labels
}

func (s *UpdateGatewayLabelRequest) SetLabels(v map[string]*string) *UpdateGatewayLabelRequest {
	s.Labels = v
	return s
}

func (s *UpdateGatewayLabelRequest) Validate() error {
	return dara.Validate(s)
}
