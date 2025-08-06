// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpecificationsForLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetSpecificationsForLicenseResponseBody
	GetData() *string
	SetRequestId(v string) *GetSpecificationsForLicenseResponseBody
	GetRequestId() *string
}

type GetSpecificationsForLicenseResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSpecificationsForLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSpecificationsForLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpecificationsForLicenseResponseBody) GetData() *string {
	return s.Data
}

func (s *GetSpecificationsForLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSpecificationsForLicenseResponseBody) SetData(v string) *GetSpecificationsForLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *GetSpecificationsForLicenseResponseBody) SetRequestId(v string) *GetSpecificationsForLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSpecificationsForLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}
