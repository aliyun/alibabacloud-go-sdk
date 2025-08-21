// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebCCGlobalSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebCCGlobalSwitchResponseBody
	GetRequestId() *string
}

type ModifyWebCCGlobalSwitchResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5AE2FC86-C840-41AE-9F1A-3A2747C7C1DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebCCGlobalSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebCCGlobalSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebCCGlobalSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebCCGlobalSwitchResponseBody) SetRequestId(v string) *ModifyWebCCGlobalSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebCCGlobalSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
