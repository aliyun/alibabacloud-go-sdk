// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVirusScanGlobalConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateVirusScanGlobalConfigResponseBody
	GetRequestId() *string
}

type UpdateVirusScanGlobalConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVirusScanGlobalConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVirusScanGlobalConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVirusScanGlobalConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVirusScanGlobalConfigResponseBody) SetRequestId(v string) *UpdateVirusScanGlobalConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVirusScanGlobalConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
