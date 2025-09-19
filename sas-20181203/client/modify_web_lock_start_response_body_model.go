// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockStartResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebLockStartResponseBody
	GetRequestId() *string
}

type ModifyWebLockStartResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D9354C1A-D709-4873-9AAE-41513327B247
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebLockStartResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockStartResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebLockStartResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebLockStartResponseBody) SetRequestId(v string) *ModifyWebLockStartResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebLockStartResponseBody) Validate() error {
	return dara.Validate(s)
}
