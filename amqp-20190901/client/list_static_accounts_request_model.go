// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStaticAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *ListStaticAccountsRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *ListStaticAccountsRequest
	GetInstanceId() *string
}

type ListStaticAccountsRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListStaticAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStaticAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListStaticAccountsRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *ListStaticAccountsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListStaticAccountsRequest) SetConsoleSessionId(v string) *ListStaticAccountsRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *ListStaticAccountsRequest) SetInstanceId(v string) *ListStaticAccountsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListStaticAccountsRequest) Validate() error {
	return dara.Validate(s)
}
