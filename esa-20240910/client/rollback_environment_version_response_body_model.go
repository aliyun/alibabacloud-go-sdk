// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackEnvironmentVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RollbackEnvironmentVersionResponseBody
	GetRequestId() *string
}

type RollbackEnvironmentVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B955107D-E658-4E77-B913-E0AC3D31693E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackEnvironmentVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackEnvironmentVersionResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackEnvironmentVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackEnvironmentVersionResponseBody) SetRequestId(v string) *RollbackEnvironmentVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackEnvironmentVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
