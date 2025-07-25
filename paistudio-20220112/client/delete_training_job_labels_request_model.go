// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrainingJobLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeys(v string) *DeleteTrainingJobLabelsRequest
	GetKeys() *string
}

type DeleteTrainingJobLabelsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// RootModelID
	Keys *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
}

func (s DeleteTrainingJobLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrainingJobLabelsRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrainingJobLabelsRequest) GetKeys() *string {
	return s.Keys
}

func (s *DeleteTrainingJobLabelsRequest) SetKeys(v string) *DeleteTrainingJobLabelsRequest {
	s.Keys = &v
	return s
}

func (s *DeleteTrainingJobLabelsRequest) Validate() error {
	return dara.Validate(s)
}
