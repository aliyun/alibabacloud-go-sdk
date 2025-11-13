// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckIpExistsInSecurityIpListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CheckIpExistsInSecurityIpListResponseBodyData) *CheckIpExistsInSecurityIpListResponseBody
	GetData() *CheckIpExistsInSecurityIpListResponseBodyData
	SetRequestId(v string) *CheckIpExistsInSecurityIpListResponseBody
	GetRequestId() *string
}

type CheckIpExistsInSecurityIpListResponseBody struct {
	Data *CheckIpExistsInSecurityIpListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 4773E4EC-025D-509F-AEA9-D53123FDFB0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckIpExistsInSecurityIpListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckIpExistsInSecurityIpListResponseBody) GoString() string {
	return s.String()
}

func (s *CheckIpExistsInSecurityIpListResponseBody) GetData() *CheckIpExistsInSecurityIpListResponseBodyData {
	return s.Data
}

func (s *CheckIpExistsInSecurityIpListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckIpExistsInSecurityIpListResponseBody) SetData(v *CheckIpExistsInSecurityIpListResponseBodyData) *CheckIpExistsInSecurityIpListResponseBody {
	s.Data = v
	return s
}

func (s *CheckIpExistsInSecurityIpListResponseBody) SetRequestId(v string) *CheckIpExistsInSecurityIpListResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckIpExistsInSecurityIpListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckIpExistsInSecurityIpListResponseBodyData struct {
	// example:
	//
	// true
	IpExists *bool `json:"IpExists,omitempty" xml:"IpExists,omitempty"`
}

func (s CheckIpExistsInSecurityIpListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CheckIpExistsInSecurityIpListResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckIpExistsInSecurityIpListResponseBodyData) GetIpExists() *bool {
	return s.IpExists
}

func (s *CheckIpExistsInSecurityIpListResponseBodyData) SetIpExists(v bool) *CheckIpExistsInSecurityIpListResponseBodyData {
	s.IpExists = &v
	return s
}

func (s *CheckIpExistsInSecurityIpListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
