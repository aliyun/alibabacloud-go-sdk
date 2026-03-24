// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstance(v int64) *ScaleServiceRequest
	GetInstance() *int64
	SetInstancesToDelete(v []*string) *ScaleServiceRequest
	GetInstancesToDelete() []*string
}

type ScaleServiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2
	Instance *int64 `json:"Instance,omitempty" xml:"Instance,omitempty"`
	// if can be null:
	// true
	InstancesToDelete []*string `json:"InstancesToDelete,omitempty" xml:"InstancesToDelete,omitempty" type:"Repeated"`
}

func (s ScaleServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ScaleServiceRequest) GoString() string {
	return s.String()
}

func (s *ScaleServiceRequest) GetInstance() *int64 {
	return s.Instance
}

func (s *ScaleServiceRequest) GetInstancesToDelete() []*string {
	return s.InstancesToDelete
}

func (s *ScaleServiceRequest) SetInstance(v int64) *ScaleServiceRequest {
	s.Instance = &v
	return s
}

func (s *ScaleServiceRequest) SetInstancesToDelete(v []*string) *ScaleServiceRequest {
	s.InstancesToDelete = v
	return s
}

func (s *ScaleServiceRequest) Validate() error {
	return dara.Validate(s)
}
