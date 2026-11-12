// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetDiskResponseBody
	GetRequestId() *string
}

type ResetDiskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetDiskResponseBody) GoString() string {
	return s.String()
}

func (s *ResetDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetDiskResponseBody) SetRequestId(v string) *ResetDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetDiskResponseBody) Validate() error {
	return dara.Validate(s)
}
