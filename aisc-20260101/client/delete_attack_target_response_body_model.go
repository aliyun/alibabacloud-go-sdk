// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAttackTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAttackTargetResponseBody
	GetRequestId() *string
}

type DeleteAttackTargetResponseBody struct {
	// The request ID. You can use this ID for troubleshooting and log tracing.
	//
	// example:
	//
	// 1EBD0C05-6C1F-4C95-9C63-B7AB7B5A9C8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAttackTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAttackTargetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAttackTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAttackTargetResponseBody) SetRequestId(v string) *DeleteAttackTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAttackTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
