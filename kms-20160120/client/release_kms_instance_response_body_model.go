// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseKmsInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseKmsInstanceResponseBody
	GetRequestId() *string
}

type ReleaseKmsInstanceResponseBody struct {
	// example:
	//
	// 475f1620-b9d3-4d35-b5c6-3fbdd941423d
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseKmsInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseKmsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseKmsInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseKmsInstanceResponseBody) SetRequestId(v string) *ReleaseKmsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseKmsInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
