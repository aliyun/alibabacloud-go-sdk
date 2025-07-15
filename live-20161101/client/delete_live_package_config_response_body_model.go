// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivePackageConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLivePackageConfigResponseBody
	GetRequestId() *string
}

type DeleteLivePackageConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 72ABAD7B-B14C-52DE-B6C6-C639FECAF5AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLivePackageConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivePackageConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLivePackageConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLivePackageConfigResponseBody) SetRequestId(v string) *DeleteLivePackageConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLivePackageConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
