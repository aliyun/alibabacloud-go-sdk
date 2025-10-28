// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLayerVersionOutput interface {
	dara.Model
	String() string
	GoString() string
	SetLayers(v []*Layer) *ListLayerVersionOutput
	GetLayers() []*Layer
	SetNextVersion(v int32) *ListLayerVersionOutput
	GetNextVersion() *int32
}

type ListLayerVersionOutput struct {
	Layers []*Layer `json:"layers" xml:"layers" type:"Repeated"`
	// example:
	//
	// 10
	NextVersion *int32 `json:"nextVersion,omitempty" xml:"nextVersion,omitempty"`
}

func (s ListLayerVersionOutput) String() string {
	return dara.Prettify(s)
}

func (s ListLayerVersionOutput) GoString() string {
	return s.String()
}

func (s *ListLayerVersionOutput) GetLayers() []*Layer {
	return s.Layers
}

func (s *ListLayerVersionOutput) GetNextVersion() *int32 {
	return s.NextVersion
}

func (s *ListLayerVersionOutput) SetLayers(v []*Layer) *ListLayerVersionOutput {
	s.Layers = v
	return s
}

func (s *ListLayerVersionOutput) SetNextVersion(v int32) *ListLayerVersionOutput {
	s.NextVersion = &v
	return s
}

func (s *ListLayerVersionOutput) Validate() error {
	if s.Layers != nil {
		for _, item := range s.Layers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
