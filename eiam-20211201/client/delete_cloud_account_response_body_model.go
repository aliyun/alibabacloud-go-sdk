// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCloudAccountResponseBody
	GetRequestId() *string
}

type DeleteCloudAccountResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCloudAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloudAccountResponseBody) SetRequestId(v string) *DeleteCloudAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloudAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
