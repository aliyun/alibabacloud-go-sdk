// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneParameterGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CloneParameterGroupResponseBody
	GetRequestId() *string
}

type CloneParameterGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1D3D5995-6BDD-43B5-93B8-2C41A2ACD6AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneParameterGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloneParameterGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CloneParameterGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloneParameterGroupResponseBody) SetRequestId(v string) *CloneParameterGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloneParameterGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
