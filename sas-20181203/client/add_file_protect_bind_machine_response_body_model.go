// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFileProtectBindMachineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddFileProtectBindMachineResponseBody
	GetRequestId() *string
}

type AddFileProtectBindMachineResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6673D49C-A9AB-40DD-B4A2-B92306701AE7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddFileProtectBindMachineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddFileProtectBindMachineResponseBody) GoString() string {
	return s.String()
}

func (s *AddFileProtectBindMachineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddFileProtectBindMachineResponseBody) SetRequestId(v string) *AddFileProtectBindMachineResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFileProtectBindMachineResponseBody) Validate() error {
	return dara.Validate(s)
}
