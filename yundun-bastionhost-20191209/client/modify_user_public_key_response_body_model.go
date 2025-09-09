// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserPublicKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyUserPublicKeyResponseBody
	GetRequestId() *string
}

type ModifyUserPublicKeyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// AAB631FB-ABD0-5783-99F3-F29573B129E4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserPublicKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserPublicKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserPublicKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyUserPublicKeyResponseBody) SetRequestId(v string) *ModifyUserPublicKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUserPublicKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
