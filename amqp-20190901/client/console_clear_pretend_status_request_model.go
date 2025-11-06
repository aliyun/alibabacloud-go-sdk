// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConsoleClearPretendStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *ConsoleClearPretendStatusRequest
	GetConsoleSessionId() *string
}

type ConsoleClearPretendStatusRequest struct {
	// This parameter is required.
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
}

func (s ConsoleClearPretendStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ConsoleClearPretendStatusRequest) GoString() string {
	return s.String()
}

func (s *ConsoleClearPretendStatusRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *ConsoleClearPretendStatusRequest) SetConsoleSessionId(v string) *ConsoleClearPretendStatusRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *ConsoleClearPretendStatusRequest) Validate() error {
	return dara.Validate(s)
}
