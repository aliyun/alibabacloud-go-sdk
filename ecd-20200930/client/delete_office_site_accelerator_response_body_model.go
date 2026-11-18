// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOfficeSiteAcceleratorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteOfficeSiteAcceleratorResponseBody
	GetRequestId() *string
}

type DeleteOfficeSiteAcceleratorResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// EE9472BC-0B5D-5458-85CD-C52BDD******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOfficeSiteAcceleratorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOfficeSiteAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOfficeSiteAcceleratorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOfficeSiteAcceleratorResponseBody) SetRequestId(v string) *DeleteOfficeSiteAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOfficeSiteAcceleratorResponseBody) Validate() error {
	return dara.Validate(s)
}
