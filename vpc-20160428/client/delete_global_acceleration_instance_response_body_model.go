// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGlobalAccelerationInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteGlobalAccelerationInstanceResponseBody
	GetRequestId() *string
}

type DeleteGlobalAccelerationInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E6E63B2A-9820-44A8-A359-9BB2DAEE6424
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGlobalAccelerationInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGlobalAccelerationInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGlobalAccelerationInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGlobalAccelerationInstanceResponseBody) SetRequestId(v string) *DeleteGlobalAccelerationInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGlobalAccelerationInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
