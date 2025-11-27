// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDbProxyInstanceSslResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDbProxyCertListItems(v *GetDbProxyInstanceSslResponseBodyDbProxyCertListItems) *GetDbProxyInstanceSslResponseBody
	GetDbProxyCertListItems() *GetDbProxyInstanceSslResponseBodyDbProxyCertListItems
	SetRequestId(v string) *GetDbProxyInstanceSslResponseBody
	GetRequestId() *string
}

type GetDbProxyInstanceSslResponseBody struct {
	// An array that consists of SSL encryption settings.
	DbProxyCertListItems *GetDbProxyInstanceSslResponseBodyDbProxyCertListItems `json:"DbProxyCertListItems,omitempty" xml:"DbProxyCertListItems,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D330E60C-8AAA-4D63-8F64-5B78F4692F98
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDbProxyInstanceSslResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDbProxyInstanceSslResponseBody) GoString() string {
	return s.String()
}

func (s *GetDbProxyInstanceSslResponseBody) GetDbProxyCertListItems() *GetDbProxyInstanceSslResponseBodyDbProxyCertListItems {
	return s.DbProxyCertListItems
}

func (s *GetDbProxyInstanceSslResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDbProxyInstanceSslResponseBody) SetDbProxyCertListItems(v *GetDbProxyInstanceSslResponseBodyDbProxyCertListItems) *GetDbProxyInstanceSslResponseBody {
	s.DbProxyCertListItems = v
	return s
}

func (s *GetDbProxyInstanceSslResponseBody) SetRequestId(v string) *GetDbProxyInstanceSslResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDbProxyInstanceSslResponseBody) Validate() error {
	if s.DbProxyCertListItems != nil {
		if err := s.DbProxyCertListItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDbProxyInstanceSslResponseBodyDbProxyCertListItems struct {
	DbProxyCertListItems []*GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems `json:"DbProxyCertListItems,omitempty" xml:"DbProxyCertListItems,omitempty" type:"Repeated"`
}

func (s GetDbProxyInstanceSslResponseBodyDbProxyCertListItems) String() string {
	return dara.Prettify(s)
}

func (s GetDbProxyInstanceSslResponseBodyDbProxyCertListItems) GoString() string {
	return s.String()
}

func (s *GetDbProxyInstanceSslResponseBodyDbProxyCertListItems) GetDbProxyCertListItems() []*GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems {
	return s.DbProxyCertListItems
}

func (s *GetDbProxyInstanceSslResponseBodyDbProxyCertListItems) SetDbProxyCertListItems(v []*GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems) *GetDbProxyInstanceSslResponseBodyDbProxyCertListItems {
	s.DbProxyCertListItems = v
	return s
}

func (s *GetDbProxyInstanceSslResponseBodyDbProxyCertListItems) Validate() error {
	if s.DbProxyCertListItems != nil {
		for _, item := range s.DbProxyCertListItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems struct {
	// The dedicated proxy endpoint for which SSL encryption is enabled.
	//
	// example:
	//
	// test1234.rwlb.rds.aliyuncs.com
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-t4n3axxxxx
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	// The ID of the dedicated proxy endpoint.
	//
	// example:
	//
	// buxxxxxxx
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The default identifier of the dedicated proxy endpoint. The value is fixed as **RWSplit**.
	//
	// example:
	//
	// RWSplit
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The time at which the certificate expires.
	//
	// example:
	//
	// 2021-12-16T08:43:20Z
	SslExpiredTime *string `json:"SslExpiredTime,omitempty" xml:"SslExpiredTime,omitempty"`
}

func (s GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems) String() string {
	return dara.Prettify(s)
}

func (s GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems) GoString() string {
	return s.String()
}

func (s *GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems) GetCertCommonName() *string {
	return s.CertCommonName
}

func (s *GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems) GetDbInstanceName() *string {
	return s.DbInstanceName
}

func (s *GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems) GetEndpointName() *string {
	return s.EndpointName
}

func (s *GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems) GetEndpointType() *string {
	return s.EndpointType
}

func (s *GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems) GetSslExpiredTime() *string {
	return s.SslExpiredTime
}

func (s *GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems) SetCertCommonName(v string) *GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems {
	s.CertCommonName = &v
	return s
}

func (s *GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems) SetDbInstanceName(v string) *GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems {
	s.DbInstanceName = &v
	return s
}

func (s *GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems) SetEndpointName(v string) *GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems {
	s.EndpointName = &v
	return s
}

func (s *GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems) SetEndpointType(v string) *GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems {
	s.EndpointType = &v
	return s
}

func (s *GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems) SetSslExpiredTime(v string) *GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems {
	s.SslExpiredTime = &v
	return s
}

func (s *GetDbProxyInstanceSslResponseBodyDbProxyCertListItemsDbProxyCertListItems) Validate() error {
	return dara.Validate(s)
}
