// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLivePackageConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLivePackageConfigResponseBody
	GetRequestId() *string
}

type AddLivePackageConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 05D00B48-DF50-5DC0-A07D-A250DFAE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLivePackageConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLivePackageConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddLivePackageConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLivePackageConfigResponseBody) SetRequestId(v string) *AddLivePackageConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLivePackageConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
