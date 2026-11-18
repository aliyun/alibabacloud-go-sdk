// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudDriveServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCloudDriveServicesResponseBody
	GetRequestId() *string
}

type DeleteCloudDriveServicesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20CA01E9-C4FF-5A64-9775-7EF48F8E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCloudDriveServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudDriveServicesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudDriveServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloudDriveServicesResponseBody) SetRequestId(v string) *DeleteCloudDriveServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloudDriveServicesResponseBody) Validate() error {
	return dara.Validate(s)
}
