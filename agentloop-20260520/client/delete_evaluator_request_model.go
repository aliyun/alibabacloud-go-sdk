// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEvaluatorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersion(v string) *DeleteEvaluatorRequest
	GetVersion() *string
}

type DeleteEvaluatorRequest struct {
	// The version to delete. If this parameter is not specified, the entire evaluator is deleted.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DeleteEvaluatorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEvaluatorRequest) GoString() string {
	return s.String()
}

func (s *DeleteEvaluatorRequest) GetVersion() *string {
	return s.Version
}

func (s *DeleteEvaluatorRequest) SetVersion(v string) *DeleteEvaluatorRequest {
	s.Version = &v
	return s
}

func (s *DeleteEvaluatorRequest) Validate() error {
	return dara.Validate(s)
}
