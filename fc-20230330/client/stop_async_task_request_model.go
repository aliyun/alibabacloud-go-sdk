// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAsyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQualifier(v string) *StopAsyncTaskRequest
	GetQualifier() *string
}

type StopAsyncTaskRequest struct {
	// The version or alias of the function.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s StopAsyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *StopAsyncTaskRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *StopAsyncTaskRequest) SetQualifier(v string) *StopAsyncTaskRequest {
	s.Qualifier = &v
	return s
}

func (s *StopAsyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
