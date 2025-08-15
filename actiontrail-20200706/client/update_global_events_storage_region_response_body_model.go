// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGlobalEventsStorageRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateGlobalEventsStorageRegionResponseBody
	GetRequestId() *string
}

type UpdateGlobalEventsStorageRegionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D7A0694E-C8FE-574E-92E3-63C5B5D23BD4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGlobalEventsStorageRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGlobalEventsStorageRegionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGlobalEventsStorageRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGlobalEventsStorageRegionResponseBody) SetRequestId(v string) *UpdateGlobalEventsStorageRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGlobalEventsStorageRegionResponseBody) Validate() error {
	return dara.Validate(s)
}
