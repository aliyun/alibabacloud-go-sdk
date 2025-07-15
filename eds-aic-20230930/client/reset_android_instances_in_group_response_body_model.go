// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAndroidInstancesInGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetAndroidInstancesInGroupResponseBody
	GetRequestId() *string
}

type ResetAndroidInstancesInGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 69BCBBE4-FCF2-59B8-AD9D-531EB422****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAndroidInstancesInGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetAndroidInstancesInGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAndroidInstancesInGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetAndroidInstancesInGroupResponseBody) SetRequestId(v string) *ResetAndroidInstancesInGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetAndroidInstancesInGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
