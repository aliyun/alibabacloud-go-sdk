// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommonBuyUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v string) *GetCommonBuyUrlRequest
	GetActionType() *string
	SetConsoleSessionId(v string) *GetCommonBuyUrlRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *GetCommonBuyUrlRequest
	GetInstanceId() *string
}

type GetCommonBuyUrlRequest struct {
	// This parameter is required.
	ActionType       *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetCommonBuyUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCommonBuyUrlRequest) GoString() string {
	return s.String()
}

func (s *GetCommonBuyUrlRequest) GetActionType() *string {
	return s.ActionType
}

func (s *GetCommonBuyUrlRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetCommonBuyUrlRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCommonBuyUrlRequest) SetActionType(v string) *GetCommonBuyUrlRequest {
	s.ActionType = &v
	return s
}

func (s *GetCommonBuyUrlRequest) SetConsoleSessionId(v string) *GetCommonBuyUrlRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetCommonBuyUrlRequest) SetInstanceId(v string) *GetCommonBuyUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCommonBuyUrlRequest) Validate() error {
	return dara.Validate(s)
}
