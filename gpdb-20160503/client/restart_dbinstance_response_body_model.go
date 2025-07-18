// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestartDBInstanceResponseBody
	GetRequestId() *string
}

type RestartDBInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A7356493-7141-4393-8951-CDA8AB5D67EC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartDBInstanceResponseBody) SetRequestId(v string) *RestartDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
