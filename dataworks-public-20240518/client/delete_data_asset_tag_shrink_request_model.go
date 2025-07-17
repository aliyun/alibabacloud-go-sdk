// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAssetTagShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *DeleteDataAssetTagShrinkRequest
	GetKey() *string
	SetValuesShrink(v string) *DeleteDataAssetTagShrinkRequest
	GetValuesShrink() *string
}

type DeleteDataAssetTagShrinkRequest struct {
	// The tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag values.
	ValuesShrink *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DeleteDataAssetTagShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAssetTagShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataAssetTagShrinkRequest) GetKey() *string {
	return s.Key
}

func (s *DeleteDataAssetTagShrinkRequest) GetValuesShrink() *string {
	return s.ValuesShrink
}

func (s *DeleteDataAssetTagShrinkRequest) SetKey(v string) *DeleteDataAssetTagShrinkRequest {
	s.Key = &v
	return s
}

func (s *DeleteDataAssetTagShrinkRequest) SetValuesShrink(v string) *DeleteDataAssetTagShrinkRequest {
	s.ValuesShrink = &v
	return s
}

func (s *DeleteDataAssetTagShrinkRequest) Validate() error {
	return dara.Validate(s)
}
