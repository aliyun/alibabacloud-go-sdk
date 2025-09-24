// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetVersionLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeys(v string) *DeleteDatasetVersionLabelsRequest
	GetKeys() *string
}

type DeleteDatasetVersionLabelsRequest struct {
	// The tag keys. Multiple tags are separated by commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// key1,key2
	Keys *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
}

func (s DeleteDatasetVersionLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetVersionLabelsRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasetVersionLabelsRequest) GetKeys() *string {
	return s.Keys
}

func (s *DeleteDatasetVersionLabelsRequest) SetKeys(v string) *DeleteDatasetVersionLabelsRequest {
	s.Keys = &v
	return s
}

func (s *DeleteDatasetVersionLabelsRequest) Validate() error {
	return dara.Validate(s)
}
