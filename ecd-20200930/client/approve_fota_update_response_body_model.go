// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveFotaUpdateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ApproveFotaUpdateResponseBody
	GetRequestId() *string
}

type ApproveFotaUpdateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApproveFotaUpdateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApproveFotaUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveFotaUpdateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApproveFotaUpdateResponseBody) SetRequestId(v string) *ApproveFotaUpdateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApproveFotaUpdateResponseBody) Validate() error {
	return dara.Validate(s)
}
