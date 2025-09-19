// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockProcessStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebLockProcessStatusResponseBody
	GetRequestId() *string
}

type ModifyWebLockProcessStatusResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 6A540B52-2441-5493-902B-37376C412776
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebLockProcessStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockProcessStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebLockProcessStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebLockProcessStatusResponseBody) SetRequestId(v string) *ModifyWebLockProcessStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebLockProcessStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
