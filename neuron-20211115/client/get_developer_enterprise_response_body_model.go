// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeveloperEnterpriseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnterprise(v *Enterprise) *GetDeveloperEnterpriseResponseBody
	GetEnterprise() *Enterprise
	SetRequestId(v string) *GetDeveloperEnterpriseResponseBody
	GetRequestId() *string
}

type GetDeveloperEnterpriseResponseBody struct {
	Enterprise *Enterprise `json:"enterprise,omitempty" xml:"enterprise,omitempty"`
	RequestId  *string     `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDeveloperEnterpriseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeveloperEnterpriseResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeveloperEnterpriseResponseBody) GetEnterprise() *Enterprise {
	return s.Enterprise
}

func (s *GetDeveloperEnterpriseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeveloperEnterpriseResponseBody) SetEnterprise(v *Enterprise) *GetDeveloperEnterpriseResponseBody {
	s.Enterprise = v
	return s
}

func (s *GetDeveloperEnterpriseResponseBody) SetRequestId(v string) *GetDeveloperEnterpriseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeveloperEnterpriseResponseBody) Validate() error {
	if s.Enterprise != nil {
		if err := s.Enterprise.Validate(); err != nil {
			return err
		}
	}
	return nil
}
