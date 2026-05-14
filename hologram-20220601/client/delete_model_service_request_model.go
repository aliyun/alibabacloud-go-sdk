// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelServiceName(v string) *DeleteModelServiceRequest
	GetModelServiceName() *string
}

type DeleteModelServiceRequest struct {
	// example:
	//
	// model-qwen
	ModelServiceName *string `json:"modelServiceName,omitempty" xml:"modelServiceName,omitempty"`
}

func (s DeleteModelServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelServiceRequest) GetModelServiceName() *string {
	return s.ModelServiceName
}

func (s *DeleteModelServiceRequest) SetModelServiceName(v string) *DeleteModelServiceRequest {
	s.ModelServiceName = &v
	return s
}

func (s *DeleteModelServiceRequest) Validate() error {
	return dara.Validate(s)
}
