// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedLicenseItemIdList(v []*string) *DeleteAppLicenseResponseBody
	GetFailedLicenseItemIdList() []*string
	SetRequestId(v string) *DeleteAppLicenseResponseBody
	GetRequestId() *string
}

type DeleteAppLicenseResponseBody struct {
	FailedLicenseItemIdList []*string `json:"FailedLicenseItemIdList,omitempty" xml:"FailedLicenseItemIdList,omitempty" type:"Repeated"`
	RequestId               *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppLicenseResponseBody) GetFailedLicenseItemIdList() []*string {
	return s.FailedLicenseItemIdList
}

func (s *DeleteAppLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppLicenseResponseBody) SetFailedLicenseItemIdList(v []*string) *DeleteAppLicenseResponseBody {
	s.FailedLicenseItemIdList = v
	return s
}

func (s *DeleteAppLicenseResponseBody) SetRequestId(v string) *DeleteAppLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}
