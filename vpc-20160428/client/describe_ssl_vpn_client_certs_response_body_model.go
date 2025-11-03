// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSslVpnClientCertsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSslVpnClientCertsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSslVpnClientCertsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSslVpnClientCertsResponseBody
	GetRequestId() *string
	SetSslVpnClientCertKeys(v *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeys) *DescribeSslVpnClientCertsResponseBody
	GetSslVpnClientCertKeys() *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeys
	SetTotalCount(v int32) *DescribeSslVpnClientCertsResponseBody
	GetTotalCount() *int32
}

type DescribeSslVpnClientCertsResponseBody struct {
	// The number of the page to return.
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
	// 5BE01CD7-5A50-472D-AC14-CA181C5C03BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the SSL client certificates.
	SslVpnClientCertKeys *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeys `json:"SslVpnClientCertKeys,omitempty" xml:"SslVpnClientCertKeys,omitempty" type:"Struct"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSslVpnClientCertsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSslVpnClientCertsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSslVpnClientCertsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSslVpnClientCertsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSslVpnClientCertsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSslVpnClientCertsResponseBody) GetSslVpnClientCertKeys() *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeys {
	return s.SslVpnClientCertKeys
}

func (s *DescribeSslVpnClientCertsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSslVpnClientCertsResponseBody) SetPageNumber(v int32) *DescribeSslVpnClientCertsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSslVpnClientCertsResponseBody) SetPageSize(v int32) *DescribeSslVpnClientCertsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSslVpnClientCertsResponseBody) SetRequestId(v string) *DescribeSslVpnClientCertsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSslVpnClientCertsResponseBody) SetSslVpnClientCertKeys(v *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeys) *DescribeSslVpnClientCertsResponseBody {
	s.SslVpnClientCertKeys = v
	return s
}

func (s *DescribeSslVpnClientCertsResponseBody) SetTotalCount(v int32) *DescribeSslVpnClientCertsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSslVpnClientCertsResponseBody) Validate() error {
	if s.SslVpnClientCertKeys != nil {
		if err := s.SslVpnClientCertKeys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeys struct {
	SslVpnClientCertKey []*DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey `json:"SslVpnClientCertKey,omitempty" xml:"SslVpnClientCertKey,omitempty" type:"Repeated"`
}

func (s DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeys) GoString() string {
	return s.String()
}

func (s *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeys) GetSslVpnClientCertKey() []*DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey {
	return s.SslVpnClientCertKey
}

func (s *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeys) SetSslVpnClientCertKey(v []*DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey) *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeys {
	s.SslVpnClientCertKey = v
	return s
}

func (s *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeys) Validate() error {
	if s.SslVpnClientCertKey != nil {
		for _, item := range s.SslVpnClientCertKey {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey struct {
	// The timestamp generated when the SSL client certificate was created. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1492747187000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The timestamp generated when the SSL client certificate expires. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1494966335000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the SSL client certificate.
	//
	// example:
	//
	// cert1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the SSL client certificate.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the SSL client certificate belongs.
	//
	// You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query resource groups.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the SSL client certificate.
	//
	// example:
	//
	// vsc-bp1n8wcf134yl0osr****
	SslVpnClientCertId *string `json:"SslVpnClientCertId,omitempty" xml:"SslVpnClientCertId,omitempty"`
	// The ID of the SSL server.
	//
	// example:
	//
	// vss-bp18q7hzj6largv4v****
	SslVpnServerId *string `json:"SslVpnServerId,omitempty" xml:"SslVpnServerId,omitempty"`
	// The status of the SSL client certificate. Valid values:
	//
	// 	- **expiring-soon**: The certificate expires in one week.
	//
	// 	- **normal**
	//
	// 	- **expired**
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey) String() string {
	return dara.Prettify(s)
}

func (s DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey) GoString() string {
	return s.String()
}

func (s *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey) GetName() *string {
	return s.Name
}

func (s *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey) GetSslVpnClientCertId() *string {
	return s.SslVpnClientCertId
}

func (s *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey) GetSslVpnServerId() *string {
	return s.SslVpnServerId
}

func (s *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey) GetStatus() *string {
	return s.Status
}

func (s *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey) SetCreateTime(v int64) *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey {
	s.CreateTime = &v
	return s
}

func (s *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey) SetEndTime(v int64) *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey {
	s.EndTime = &v
	return s
}

func (s *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey) SetName(v string) *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey {
	s.Name = &v
	return s
}

func (s *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey) SetRegionId(v string) *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey {
	s.RegionId = &v
	return s
}

func (s *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey) SetResourceGroupId(v string) *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey) SetSslVpnClientCertId(v string) *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey {
	s.SslVpnClientCertId = &v
	return s
}

func (s *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey) SetSslVpnServerId(v string) *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey {
	s.SslVpnServerId = &v
	return s
}

func (s *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey) SetStatus(v string) *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey {
	s.Status = &v
	return s
}

func (s *DescribeSslVpnClientCertsResponseBodySslVpnClientCertKeysSslVpnClientCertKey) Validate() error {
	return dara.Validate(s)
}
