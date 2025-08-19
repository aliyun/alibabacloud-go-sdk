// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLayersOutput interface {
	dara.Model
	String() string
	GoString() string
	SetLayers(v []*Layer) *ListLayersOutput
	GetLayers() []*Layer
	SetNextToken(v string) *ListLayersOutput
	GetNextToken() *string
}

type ListLayersOutput struct {
	Layers []*Layer `json:"layers" xml:"layers" type:"Repeated"`
	// example:
	//
	// next-layer-name
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListLayersOutput) String() string {
	return dara.Prettify(s)
}

func (s ListLayersOutput) GoString() string {
	return s.String()
}

func (s *ListLayersOutput) GetLayers() []*Layer {
	return s.Layers
}

func (s *ListLayersOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLayersOutput) SetLayers(v []*Layer) *ListLayersOutput {
	s.Layers = v
	return s
}

func (s *ListLayersOutput) SetNextToken(v string) *ListLayersOutput {
	s.NextToken = &v
	return s
}

func (s *ListLayersOutput) Validate() error {
	return dara.Validate(s)
}
