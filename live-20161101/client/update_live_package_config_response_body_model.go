// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePackageConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLivePackageConfigResponseBody
	GetRequestId() *string
}

type UpdateLivePackageConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLivePackageConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePackageConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLivePackageConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLivePackageConfigResponseBody) SetRequestId(v string) *UpdateLivePackageConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLivePackageConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
