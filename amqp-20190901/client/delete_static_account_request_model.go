// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStaticAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *DeleteStaticAccountRequest
	GetConsoleSessionId() *string
	SetCreateTimeStamp(v int64) *DeleteStaticAccountRequest
	GetCreateTimeStamp() *int64
	SetInstanceId(v string) *DeleteStaticAccountRequest
	GetInstanceId() *string
	SetUserName(v string) *DeleteStaticAccountRequest
	GetUserName() *string
}

type DeleteStaticAccountRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	CreateTimeStamp *int64  `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DeleteStaticAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStaticAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteStaticAccountRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *DeleteStaticAccountRequest) GetCreateTimeStamp() *int64 {
	return s.CreateTimeStamp
}

func (s *DeleteStaticAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteStaticAccountRequest) GetUserName() *string {
	return s.UserName
}

func (s *DeleteStaticAccountRequest) SetConsoleSessionId(v string) *DeleteStaticAccountRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *DeleteStaticAccountRequest) SetCreateTimeStamp(v int64) *DeleteStaticAccountRequest {
	s.CreateTimeStamp = &v
	return s
}

func (s *DeleteStaticAccountRequest) SetInstanceId(v string) *DeleteStaticAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteStaticAccountRequest) SetUserName(v string) *DeleteStaticAccountRequest {
	s.UserName = &v
	return s
}

func (s *DeleteStaticAccountRequest) Validate() error {
	return dara.Validate(s)
}
