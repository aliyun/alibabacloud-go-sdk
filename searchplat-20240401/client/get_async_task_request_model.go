// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *GetAsyncTaskRequest
	GetDryRun() *bool
}

type GetAsyncTaskRequest struct {
	// Specifies whether to perform a dry run request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s GetAsyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncTaskRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *GetAsyncTaskRequest) SetDryRun(v bool) *GetAsyncTaskRequest {
	s.DryRun = &v
	return s
}

func (s *GetAsyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
