// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileProtectEventStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFileProtectEventStatusResponseBody
	GetRequestId() *string
}

type UpdateFileProtectEventStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C286491D-4A2F-589A-B63B-D2AD3DA9BD71
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFileProtectEventStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileProtectEventStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFileProtectEventStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFileProtectEventStatusResponseBody) SetRequestId(v string) *UpdateFileProtectEventStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFileProtectEventStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
