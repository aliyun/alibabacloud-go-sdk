// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelKeys(v []*string) *DeleteGatewayLabelRequest
	GetLabelKeys() []*string
}

type DeleteGatewayLabelRequest struct {
	// This parameter is required.
	LabelKeys []*string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty" type:"Repeated"`
}

func (s DeleteGatewayLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayLabelRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayLabelRequest) GetLabelKeys() []*string {
	return s.LabelKeys
}

func (s *DeleteGatewayLabelRequest) SetLabelKeys(v []*string) *DeleteGatewayLabelRequest {
	s.LabelKeys = v
	return s
}

func (s *DeleteGatewayLabelRequest) Validate() error {
	return dara.Validate(s)
}
