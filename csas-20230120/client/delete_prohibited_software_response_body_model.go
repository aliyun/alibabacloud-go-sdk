// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProhibitedSoftwareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteProhibitedSoftwareResponseBody
	GetRequestId() *string
}

type DeleteProhibitedSoftwareResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8E53BDC2-5630-58A6-BA3D-5761D4A80A99
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProhibitedSoftwareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProhibitedSoftwareResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProhibitedSoftwareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProhibitedSoftwareResponseBody) SetRequestId(v string) *DeleteProhibitedSoftwareResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProhibitedSoftwareResponseBody) Validate() error {
	return dara.Validate(s)
}
