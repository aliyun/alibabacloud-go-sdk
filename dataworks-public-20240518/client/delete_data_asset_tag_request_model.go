// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAssetTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *DeleteDataAssetTagRequest
	GetKey() *string
	SetValues(v []*string) *DeleteDataAssetTagRequest
	GetValues() []*string
}

type DeleteDataAssetTagRequest struct {
	// The tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DeleteDataAssetTagRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAssetTagRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataAssetTagRequest) GetKey() *string {
	return s.Key
}

func (s *DeleteDataAssetTagRequest) GetValues() []*string {
	return s.Values
}

func (s *DeleteDataAssetTagRequest) SetKey(v string) *DeleteDataAssetTagRequest {
	s.Key = &v
	return s
}

func (s *DeleteDataAssetTagRequest) SetValues(v []*string) *DeleteDataAssetTagRequest {
	s.Values = v
	return s
}

func (s *DeleteDataAssetTagRequest) Validate() error {
	return dara.Validate(s)
}
