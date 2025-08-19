// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLayerVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *ListLayerVersionsRequest
	GetLimit() *int32
	SetStartVersion(v string) *ListLayerVersionsRequest
	GetStartVersion() *string
}

type ListLayerVersionsRequest struct {
	// The number of versions to be returned.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The initial version of the layer.
	//
	// example:
	//
	// 1
	StartVersion *string `json:"startVersion,omitempty" xml:"startVersion,omitempty"`
}

func (s ListLayerVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLayerVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListLayerVersionsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListLayerVersionsRequest) GetStartVersion() *string {
	return s.StartVersion
}

func (s *ListLayerVersionsRequest) SetLimit(v int32) *ListLayerVersionsRequest {
	s.Limit = &v
	return s
}

func (s *ListLayerVersionsRequest) SetStartVersion(v string) *ListLayerVersionsRequest {
	s.StartVersion = &v
	return s
}

func (s *ListLayerVersionsRequest) Validate() error {
	return dara.Validate(s)
}
