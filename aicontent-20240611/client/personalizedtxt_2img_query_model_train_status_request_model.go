// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPersonalizedtxt2imgQueryModelTrainStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelId(v string) *Personalizedtxt2imgQueryModelTrainStatusRequest
	GetModelId() *string
}

type Personalizedtxt2imgQueryModelTrainStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// modelId-xxxx-xxxx-xxxx
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
}

func (s Personalizedtxt2imgQueryModelTrainStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s Personalizedtxt2imgQueryModelTrainStatusRequest) GoString() string {
	return s.String()
}

func (s *Personalizedtxt2imgQueryModelTrainStatusRequest) GetModelId() *string {
	return s.ModelId
}

func (s *Personalizedtxt2imgQueryModelTrainStatusRequest) SetModelId(v string) *Personalizedtxt2imgQueryModelTrainStatusRequest {
	s.ModelId = &v
	return s
}

func (s *Personalizedtxt2imgQueryModelTrainStatusRequest) Validate() error {
	return dara.Validate(s)
}
