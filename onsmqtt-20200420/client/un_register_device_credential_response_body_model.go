// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnRegisterDeviceCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnRegisterDeviceCredentialResponseBody
	GetRequestId() *string
}

type UnRegisterDeviceCredentialResponseBody struct {
	// The unique ID that the system generates for the request. This parameter is a common parameter.
	//
	// example:
	//
	// E4581CCD-62AF-44D9-B5B4-D1DBDC0E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnRegisterDeviceCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnRegisterDeviceCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *UnRegisterDeviceCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnRegisterDeviceCredentialResponseBody) SetRequestId(v string) *UnRegisterDeviceCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnRegisterDeviceCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
