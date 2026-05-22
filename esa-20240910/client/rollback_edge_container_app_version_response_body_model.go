// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackEdgeContainerAppVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RollbackEdgeContainerAppVersionResponseBody
	GetRequestId() *string
}

type RollbackEdgeContainerAppVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackEdgeContainerAppVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackEdgeContainerAppVersionResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackEdgeContainerAppVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackEdgeContainerAppVersionResponseBody) SetRequestId(v string) *RollbackEdgeContainerAppVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackEdgeContainerAppVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
