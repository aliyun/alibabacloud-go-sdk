// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebAIProtectModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebAIProtectModeResponseBody
	GetRequestId() *string
}

type ModifyWebAIProtectModeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebAIProtectModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebAIProtectModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebAIProtectModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebAIProtectModeResponseBody) SetRequestId(v string) *ModifyWebAIProtectModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebAIProtectModeResponseBody) Validate() error {
	return dara.Validate(s)
}
