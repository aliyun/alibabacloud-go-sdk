// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseRenderingDataPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseRenderingDataPackageResponseBody
	GetRequestId() *string
}

type ReleaseRenderingDataPackageResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 5BEF36E7-3838-5B92-BA32-87DBF1425ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseRenderingDataPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseRenderingDataPackageResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseRenderingDataPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseRenderingDataPackageResponseBody) SetRequestId(v string) *ReleaseRenderingDataPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseRenderingDataPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
