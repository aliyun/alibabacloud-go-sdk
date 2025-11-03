// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDesktopMaintenanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDesktopMaintenanceResponseBody
	GetRequestId() *string
}

type SetDesktopMaintenanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BA6A1853-3EA9-5EEB-86C8-3D14A3E01905
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDesktopMaintenanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDesktopMaintenanceResponseBody) GoString() string {
	return s.String()
}

func (s *SetDesktopMaintenanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDesktopMaintenanceResponseBody) SetRequestId(v string) *SetDesktopMaintenanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDesktopMaintenanceResponseBody) Validate() error {
	return dara.Validate(s)
}
