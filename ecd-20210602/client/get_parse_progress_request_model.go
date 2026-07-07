// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParseProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskKey(v string) *GetParseProgressRequest
	GetTaskKey() *string
}

type GetParseProgressRequest struct {
	// The task key for parsing the skill package.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2E7D8B71-2677-1B4C-9E25-A88B9C5******
	TaskKey *string `json:"TaskKey,omitempty" xml:"TaskKey,omitempty"`
}

func (s GetParseProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s GetParseProgressRequest) GoString() string {
	return s.String()
}

func (s *GetParseProgressRequest) GetTaskKey() *string {
	return s.TaskKey
}

func (s *GetParseProgressRequest) SetTaskKey(v string) *GetParseProgressRequest {
	s.TaskKey = &v
	return s
}

func (s *GetParseProgressRequest) Validate() error {
	return dara.Validate(s)
}
