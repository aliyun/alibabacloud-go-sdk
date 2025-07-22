// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRedisAllSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleContext(v string) *GetRedisAllSessionRequest
	GetConsoleContext() *string
	SetInstanceId(v string) *GetRedisAllSessionRequest
	GetInstanceId() *string
}

type GetRedisAllSessionRequest struct {
	// The reserved parameter.
	//
	// example:
	//
	// None
	ConsoleContext *string `json:"ConsoleContext,omitempty" xml:"ConsoleContext,omitempty"`
	// The database instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-2zemyfd1sh1u2i****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetRedisAllSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRedisAllSessionRequest) GoString() string {
	return s.String()
}

func (s *GetRedisAllSessionRequest) GetConsoleContext() *string {
	return s.ConsoleContext
}

func (s *GetRedisAllSessionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRedisAllSessionRequest) SetConsoleContext(v string) *GetRedisAllSessionRequest {
	s.ConsoleContext = &v
	return s
}

func (s *GetRedisAllSessionRequest) SetInstanceId(v string) *GetRedisAllSessionRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRedisAllSessionRequest) Validate() error {
	return dara.Validate(s)
}
