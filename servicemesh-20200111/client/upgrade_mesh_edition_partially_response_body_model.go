// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeMeshEditionPartiallyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeMeshEditionPartiallyResponseBody
	GetRequestId() *string
}

type UpgradeMeshEditionPartiallyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 31d3a0f0-07ed-4f6e-9004-1804498c****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeMeshEditionPartiallyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeMeshEditionPartiallyResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeMeshEditionPartiallyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeMeshEditionPartiallyResponseBody) SetRequestId(v string) *UpgradeMeshEditionPartiallyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeMeshEditionPartiallyResponseBody) Validate() error {
	return dara.Validate(s)
}
