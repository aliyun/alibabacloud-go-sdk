// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultOlapComputeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDefaultOlapComputeGroupResponseBody
	GetRequestId() *string
}

type SetDefaultOlapComputeGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDefaultOlapComputeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultOlapComputeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultOlapComputeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDefaultOlapComputeGroupResponseBody) SetRequestId(v string) *SetDefaultOlapComputeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDefaultOlapComputeGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
