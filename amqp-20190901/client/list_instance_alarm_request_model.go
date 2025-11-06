// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *ListInstanceAlarmRequest
	GetConsoleSessionId() *string
}

type ListInstanceAlarmRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
}

func (s ListInstanceAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAlarmRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceAlarmRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *ListInstanceAlarmRequest) SetConsoleSessionId(v string) *ListInstanceAlarmRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *ListInstanceAlarmRequest) Validate() error {
	return dara.Validate(s)
}
