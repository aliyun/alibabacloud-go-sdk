// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAbnormalCloudResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAbnormalCloudResources(v []*DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources) *DescribeAbnormalCloudResourcesResponseBody
	GetAbnormalCloudResources() []*DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources
	SetRequestId(v string) *DescribeAbnormalCloudResourcesResponseBody
	GetRequestId() *string
}

type DescribeAbnormalCloudResourcesResponseBody struct {
	AbnormalCloudResources []*DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources `json:"AbnormalCloudResources,omitempty" xml:"AbnormalCloudResources,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 66A98669-CC6E-4F3E-80A6-***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAbnormalCloudResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAbnormalCloudResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAbnormalCloudResourcesResponseBody) GetAbnormalCloudResources() []*DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources {
	return s.AbnormalCloudResources
}

func (s *DescribeAbnormalCloudResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAbnormalCloudResourcesResponseBody) SetAbnormalCloudResources(v []*DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources) *DescribeAbnormalCloudResourcesResponseBody {
	s.AbnormalCloudResources = v
	return s
}

func (s *DescribeAbnormalCloudResourcesResponseBody) SetRequestId(v string) *DescribeAbnormalCloudResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAbnormalCloudResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources struct {
	Details []*DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// example:
	//
	// CertExpired
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// lb-***
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// example:
	//
	// test-name
	ResourceInstanceName *string `json:"ResourceInstanceName,omitempty" xml:"ResourceInstanceName,omitempty"`
	// example:
	//
	// 80
	ResourceInstancePort *int32 `json:"ResourceInstancePort,omitempty" xml:"ResourceInstancePort,omitempty"`
	// example:
	//
	// clb7
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
}

func (s DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources) GoString() string {
	return s.String()
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources) GetDetails() []*DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails {
	return s.Details
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources) GetReason() *string {
	return s.Reason
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources) GetResourceInstanceName() *string {
	return s.ResourceInstanceName
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources) GetResourceInstancePort() *int32 {
	return s.ResourceInstancePort
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources) GetResourceProduct() *string {
	return s.ResourceProduct
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources) SetDetails(v []*DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails) *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources {
	s.Details = v
	return s
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources) SetReason(v string) *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources {
	s.Reason = &v
	return s
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources) SetResourceInstanceId(v string) *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources) SetResourceInstanceName(v string) *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources {
	s.ResourceInstanceName = &v
	return s
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources) SetResourceInstancePort(v int32) *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources {
	s.ResourceInstancePort = &v
	return s
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources) SetResourceProduct(v string) *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources {
	s.ResourceProduct = &v
	return s
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResources) Validate() error {
	return dara.Validate(s)
}

type DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails struct {
	// example:
	//
	// default
	AppliedType *string `json:"AppliedType,omitempty" xml:"AppliedType,omitempty"`
	// example:
	//
	// test-name
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// example:
	//
	// CertExpired
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// www.test.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// example:
	//
	// 1735009193
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 123-cn-hangzhou
	ProductCertId *string `json:"ProductCertId,omitempty" xml:"ProductCertId,omitempty"`
	// example:
	//
	// test-name
	ProductCertName *string `json:"ProductCertName,omitempty" xml:"ProductCertName,omitempty"`
	// example:
	//
	// www.test.com
	ProductDomainExtension *string `json:"ProductDomainExtension,omitempty" xml:"ProductDomainExtension,omitempty"`
}

func (s DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails) GoString() string {
	return s.String()
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails) GetAppliedType() *string {
	return s.AppliedType
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails) GetCertName() *string {
	return s.CertName
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails) GetCode() *string {
	return s.Code
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails) GetCommonName() *string {
	return s.CommonName
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails) GetProductCertId() *string {
	return s.ProductCertId
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails) GetProductCertName() *string {
	return s.ProductCertName
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails) GetProductDomainExtension() *string {
	return s.ProductDomainExtension
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails) SetAppliedType(v string) *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails {
	s.AppliedType = &v
	return s
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails) SetCertName(v string) *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails {
	s.CertName = &v
	return s
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails) SetCode(v string) *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails {
	s.Code = &v
	return s
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails) SetCommonName(v string) *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails {
	s.CommonName = &v
	return s
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails) SetExpireTime(v int64) *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails {
	s.ExpireTime = &v
	return s
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails) SetProductCertId(v string) *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails {
	s.ProductCertId = &v
	return s
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails) SetProductCertName(v string) *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails {
	s.ProductCertName = &v
	return s
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails) SetProductDomainExtension(v string) *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails {
	s.ProductDomainExtension = &v
	return s
}

func (s *DescribeAbnormalCloudResourcesResponseBodyAbnormalCloudResourcesDetails) Validate() error {
	return dara.Validate(s)
}
