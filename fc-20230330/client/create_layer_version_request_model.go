// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLayerVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateLayerVersionInput) *CreateLayerVersionRequest
	GetBody() *CreateLayerVersionInput
}

type CreateLayerVersionRequest struct {
	// The information about layer configurations.
	//
	// This parameter is required.
	Body *CreateLayerVersionInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLayerVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLayerVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateLayerVersionRequest) GetBody() *CreateLayerVersionInput {
	return s.Body
}

func (s *CreateLayerVersionRequest) SetBody(v *CreateLayerVersionInput) *CreateLayerVersionRequest {
	s.Body = v
	return s
}

func (s *CreateLayerVersionRequest) Validate() error {
	return dara.Validate(s)
}
