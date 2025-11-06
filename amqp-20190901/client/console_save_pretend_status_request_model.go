// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConsoleSavePretendStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *ConsoleSavePretendStatusRequest
	GetConsoleSessionId() *string
	SetKey(v string) *ConsoleSavePretendStatusRequest
	GetKey() *string
	SetType(v int32) *ConsoleSavePretendStatusRequest
	GetType() *int32
}

type ConsoleSavePretendStatusRequest struct {
	// This parameter is required.
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ConsoleSavePretendStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ConsoleSavePretendStatusRequest) GoString() string {
	return s.String()
}

func (s *ConsoleSavePretendStatusRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *ConsoleSavePretendStatusRequest) GetKey() *string {
	return s.Key
}

func (s *ConsoleSavePretendStatusRequest) GetType() *int32 {
	return s.Type
}

func (s *ConsoleSavePretendStatusRequest) SetConsoleSessionId(v string) *ConsoleSavePretendStatusRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *ConsoleSavePretendStatusRequest) SetKey(v string) *ConsoleSavePretendStatusRequest {
	s.Key = &v
	return s
}

func (s *ConsoleSavePretendStatusRequest) SetType(v int32) *ConsoleSavePretendStatusRequest {
	s.Type = &v
	return s
}

func (s *ConsoleSavePretendStatusRequest) Validate() error {
	return dara.Validate(s)
}
