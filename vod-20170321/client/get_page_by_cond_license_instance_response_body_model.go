// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPageByCondLicenseInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetPageByCondLicenseInstanceResponseBodyData) *GetPageByCondLicenseInstanceResponseBody
	GetData() *GetPageByCondLicenseInstanceResponseBodyData
	SetRequestId(v string) *GetPageByCondLicenseInstanceResponseBody
	GetRequestId() *string
}

type GetPageByCondLicenseInstanceResponseBody struct {
	Data      *GetPageByCondLicenseInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPageByCondLicenseInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPageByCondLicenseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetPageByCondLicenseInstanceResponseBody) GetData() *GetPageByCondLicenseInstanceResponseBodyData {
	return s.Data
}

func (s *GetPageByCondLicenseInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPageByCondLicenseInstanceResponseBody) SetData(v *GetPageByCondLicenseInstanceResponseBodyData) *GetPageByCondLicenseInstanceResponseBody {
	s.Data = v
	return s
}

func (s *GetPageByCondLicenseInstanceResponseBody) SetRequestId(v string) *GetPageByCondLicenseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPageByCondLicenseInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPageByCondLicenseInstanceResponseBodyData struct {
	List  []*LicenseInstanceAppDTO `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Total *int64                   `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetPageByCondLicenseInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPageByCondLicenseInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPageByCondLicenseInstanceResponseBodyData) GetList() []*LicenseInstanceAppDTO {
	return s.List
}

func (s *GetPageByCondLicenseInstanceResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetPageByCondLicenseInstanceResponseBodyData) SetList(v []*LicenseInstanceAppDTO) *GetPageByCondLicenseInstanceResponseBodyData {
	s.List = v
	return s
}

func (s *GetPageByCondLicenseInstanceResponseBodyData) SetTotal(v int64) *GetPageByCondLicenseInstanceResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetPageByCondLicenseInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
