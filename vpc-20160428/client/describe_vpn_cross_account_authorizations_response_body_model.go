// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnCrossAccountAuthorizationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccountAuthorizations(v []*DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations) *DescribeVpnCrossAccountAuthorizationsResponseBody
	GetCrossAccountAuthorizations() []*DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations
	SetPageNumber(v int32) *DescribeVpnCrossAccountAuthorizationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpnCrossAccountAuthorizationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVpnCrossAccountAuthorizationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpnCrossAccountAuthorizationsResponseBody
	GetTotalCount() *int32
}

type DescribeVpnCrossAccountAuthorizationsResponseBody struct {
	// The cross-account authorization information about the IPsec-VPN connection.
	CrossAccountAuthorizations []*DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations `json:"CrossAccountAuthorizations,omitempty" xml:"CrossAccountAuthorizations,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DB04E39A-6F0C-36AC-BCA0-B6D371B90062
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVpnCrossAccountAuthorizationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnCrossAccountAuthorizationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBody) GetCrossAccountAuthorizations() []*DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations {
	return s.CrossAccountAuthorizations
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBody) SetCrossAccountAuthorizations(v []*DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations) *DescribeVpnCrossAccountAuthorizationsResponseBody {
	s.CrossAccountAuthorizations = v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBody) SetPageNumber(v int32) *DescribeVpnCrossAccountAuthorizationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBody) SetPageSize(v int32) *DescribeVpnCrossAccountAuthorizationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBody) SetRequestId(v string) *DescribeVpnCrossAccountAuthorizationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBody) SetTotalCount(v int32) *DescribeVpnCrossAccountAuthorizationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBody) Validate() error {
	if s.CrossAccountAuthorizations != nil {
		for _, item := range s.CrossAccountAuthorizations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations struct {
	// The ID of the Alibaba Cloud account to which the IPsec-VPN connection belongs.
	//
	// example:
	//
	// 1250123456123456
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-vv8h0t3klfpaae****
	BindInstance *string `json:"BindInstance,omitempty" xml:"BindInstance,omitempty"`
	// The type of resource that can be associated with the IPsec-VPN connection.
	//
	// Only **CEN*	- can be returned, which indicates that the IPsec-VPN connection is authorized to be associated with the transit router of a Cloud Enterprise Network (CEN) instance that belongs to another Alibaba Cloud account.
	//
	// example:
	//
	// CEN
	BindProduct *string `json:"BindProduct,omitempty" xml:"BindProduct,omitempty"`
	// The ID of the Alibaba Cloud account whose resources the IPsec-VPN connection is authorized to be associated with.
	//
	// example:
	//
	// 1210123456123456
	BindUid *int64 `json:"BindUid,omitempty" xml:"BindUid,omitempty"`
	// The time when the authorization for the IPsec-VPN connection was created.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1658201810000
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the IPsec-VPN connection.
	//
	// example:
	//
	// vco-p0w2jpkhi2eeop6q6****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
}

func (s DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations) GoString() string {
	return s.String()
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations) GetBindInstance() *string {
	return s.BindInstance
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations) GetBindProduct() *string {
	return s.BindProduct
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations) GetBindUid() *int64 {
	return s.BindUid
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations) SetAliUid(v int64) *DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations {
	s.AliUid = &v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations) SetBindInstance(v string) *DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations {
	s.BindInstance = &v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations) SetBindProduct(v string) *DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations {
	s.BindProduct = &v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations) SetBindUid(v int64) *DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations {
	s.BindUid = &v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations) SetCreationTime(v int64) *DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations {
	s.CreationTime = &v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations) SetVpnConnectionId(v string) *DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations {
	s.VpnConnectionId = &v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsResponseBodyCrossAccountAuthorizations) Validate() error {
	return dara.Validate(s)
}
