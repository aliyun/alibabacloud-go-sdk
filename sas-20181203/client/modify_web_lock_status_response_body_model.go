// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebLockStatusResponseBody
	GetRequestId() *string
}

type ModifyWebLockStatusResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D9354C1A-D709-4873-9AAE-41513327B247
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebLockStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebLockStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebLockStatusResponseBody) SetRequestId(v string) *ModifyWebLockStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebLockStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
