// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserProvisioningEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUserProvisioningEventResponseBody
	GetRequestId() *string
}

type DeleteUserProvisioningEventResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A9287DA5-FD59-32A0-A810-1962E8B58ABB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserProvisioningEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserProvisioningEventResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserProvisioningEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserProvisioningEventResponseBody) SetRequestId(v string) *DeleteUserProvisioningEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserProvisioningEventResponseBody) Validate() error {
	return dara.Validate(s)
}
