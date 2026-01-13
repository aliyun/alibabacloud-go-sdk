// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperationResult interface {
	dara.Model
	String() string
	GoString() string
	SetSuccess(v bool) *OperationResult
	GetSuccess() *bool
}

type OperationResult struct {
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OperationResult) String() string {
	return dara.Prettify(s)
}

func (s OperationResult) GoString() string {
	return s.String()
}

func (s *OperationResult) GetSuccess() *bool {
	return s.Success
}

func (s *OperationResult) SetSuccess(v bool) *OperationResult {
	s.Success = &v
	return s
}

func (s *OperationResult) Validate() error {
	return dara.Validate(s)
}
