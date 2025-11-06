// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *ListProjectRequest
	GetConsoleSessionId() *string
}

type ListProjectRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
}

func (s ListProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectRequest) GoString() string {
	return s.String()
}

func (s *ListProjectRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *ListProjectRequest) SetConsoleSessionId(v string) *ListProjectRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *ListProjectRequest) Validate() error {
	return dara.Validate(s)
}
