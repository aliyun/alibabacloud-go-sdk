// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockUpdateConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebLockUpdateConfigResponseBody
	GetRequestId() *string
}

type ModifyWebLockUpdateConfigResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D9354C1A-D709-4873-9AAE-41513327B247
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebLockUpdateConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockUpdateConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebLockUpdateConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebLockUpdateConfigResponseBody) SetRequestId(v string) *ModifyWebLockUpdateConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebLockUpdateConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
