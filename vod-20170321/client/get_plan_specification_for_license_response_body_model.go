// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlanSpecificationForLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetPlanSpecificationForLicenseResponseBody
	GetData() *string
	SetRequestId(v string) *GetPlanSpecificationForLicenseResponseBody
	GetRequestId() *string
}

type GetPlanSpecificationForLicenseResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPlanSpecificationForLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPlanSpecificationForLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *GetPlanSpecificationForLicenseResponseBody) GetData() *string {
	return s.Data
}

func (s *GetPlanSpecificationForLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPlanSpecificationForLicenseResponseBody) SetData(v string) *GetPlanSpecificationForLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *GetPlanSpecificationForLicenseResponseBody) SetRequestId(v string) *GetPlanSpecificationForLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPlanSpecificationForLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}
