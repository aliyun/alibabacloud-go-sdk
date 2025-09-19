// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLoginSwitchConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyLoginSwitchConfigResponseBody
	GetRequestId() *string
}

type ModifyLoginSwitchConfigResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// B256A525-7E42-4BB9-A27C-9017FDDFF1A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLoginSwitchConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLoginSwitchConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLoginSwitchConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLoginSwitchConfigResponseBody) SetRequestId(v string) *ModifyLoginSwitchConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLoginSwitchConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
