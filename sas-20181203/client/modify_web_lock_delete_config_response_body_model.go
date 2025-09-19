// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockDeleteConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebLockDeleteConfigResponseBody
	GetRequestId() *string
}

type ModifyWebLockDeleteConfigResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 299D0992-271A-5750-ACEB-46D322862BFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebLockDeleteConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockDeleteConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebLockDeleteConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebLockDeleteConfigResponseBody) SetRequestId(v string) *ModifyWebLockDeleteConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebLockDeleteConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
