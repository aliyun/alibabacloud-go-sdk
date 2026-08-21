// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeviceGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDeviceGroupsResponseBody
	GetRequestId() *string
}

type DeleteDeviceGroupsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// C51D9340-4604-5331-AE62-407F3B408F86
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDeviceGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeviceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeviceGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDeviceGroupsResponseBody) SetRequestId(v string) *DeleteDeviceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeviceGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
