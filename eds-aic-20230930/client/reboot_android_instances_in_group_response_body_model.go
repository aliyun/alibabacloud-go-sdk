// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootAndroidInstancesInGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RebootAndroidInstancesInGroupResponseBody
	GetRequestId() *string
}

type RebootAndroidInstancesInGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 227CBB4C-F5DC-589D-A667-C5CA3D52****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootAndroidInstancesInGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebootAndroidInstancesInGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RebootAndroidInstancesInGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebootAndroidInstancesInGroupResponseBody) SetRequestId(v string) *RebootAndroidInstancesInGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootAndroidInstancesInGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
