// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEvaluatorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersion(v string) *GetEvaluatorRequest
	GetVersion() *string
}

type GetEvaluatorRequest struct {
	// The target version number. If not specified, the latest version is used.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetEvaluatorRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEvaluatorRequest) GoString() string {
	return s.String()
}

func (s *GetEvaluatorRequest) GetVersion() *string {
	return s.Version
}

func (s *GetEvaluatorRequest) SetVersion(v string) *GetEvaluatorRequest {
	s.Version = &v
	return s
}

func (s *GetEvaluatorRequest) Validate() error {
	return dara.Validate(s)
}
