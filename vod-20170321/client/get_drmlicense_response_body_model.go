// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDRMLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLicense(v string) *GetDRMLicenseResponseBody
	GetLicense() *string
	SetRequestId(v string) *GetDRMLicenseResponseBody
	GetRequestId() *string
}

type GetDRMLicenseResponseBody struct {
	License   *string `json:"License,omitempty" xml:"License,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDRMLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDRMLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *GetDRMLicenseResponseBody) GetLicense() *string {
	return s.License
}

func (s *GetDRMLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDRMLicenseResponseBody) SetLicense(v string) *GetDRMLicenseResponseBody {
	s.License = &v
	return s
}

func (s *GetDRMLicenseResponseBody) SetRequestId(v string) *GetDRMLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDRMLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}
