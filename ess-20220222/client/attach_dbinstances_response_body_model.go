// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDBInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachDBInstancesResponseBody
	GetRequestId() *string
}

type AttachDBInstancesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachDBInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachDBInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *AttachDBInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachDBInstancesResponseBody) SetRequestId(v string) *AttachDBInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachDBInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
