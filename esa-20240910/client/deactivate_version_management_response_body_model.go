// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactivateVersionManagementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeactivateVersionManagementResponseBody
	GetRequestId() *string
}

type DeactivateVersionManagementResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D61E4801-EAFF-4A63-AAE1-FBF6CE1CFD1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeactivateVersionManagementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeactivateVersionManagementResponseBody) GoString() string {
	return s.String()
}

func (s *DeactivateVersionManagementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeactivateVersionManagementResponseBody) SetRequestId(v string) *DeactivateVersionManagementResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeactivateVersionManagementResponseBody) Validate() error {
	return dara.Validate(s)
}
