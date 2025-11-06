// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *ListInstanceRequest
	GetConsoleSessionId() *string
}

type ListInstanceRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
}

func (s ListInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *ListInstanceRequest) SetConsoleSessionId(v string) *ListInstanceRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *ListInstanceRequest) Validate() error {
	return dara.Validate(s)
}
