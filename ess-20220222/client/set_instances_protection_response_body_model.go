// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetInstancesProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetInstancesProtectionResponseBody
	GetRequestId() *string
}

type SetInstancesProtectionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetInstancesProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetInstancesProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *SetInstancesProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetInstancesProtectionResponseBody) SetRequestId(v string) *SetInstancesProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetInstancesProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
