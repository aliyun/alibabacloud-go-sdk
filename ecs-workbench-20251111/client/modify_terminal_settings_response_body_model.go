// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTerminalSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyTerminalSettingsResponseBody
	GetRequestId() *string
}

type ModifyTerminalSettingsResponseBody struct {
	// 请求ID
	//
	// example:
	//
	// 47348885-C929-489A-93D7-B2E99D50D77B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTerminalSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTerminalSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTerminalSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTerminalSettingsResponseBody) SetRequestId(v string) *ModifyTerminalSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTerminalSettingsResponseBody) Validate() error {
	return dara.Validate(s)
}
