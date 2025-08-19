// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAsyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQualifier(v string) *GetAsyncTaskRequest
	GetQualifier() *string
}

type GetAsyncTaskRequest struct {
	// The version or alias of the function.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s GetAsyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncTaskRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *GetAsyncTaskRequest) SetQualifier(v string) *GetAsyncTaskRequest {
	s.Qualifier = &v
	return s
}

func (s *GetAsyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
