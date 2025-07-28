// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProductInstances(v []*DescribeProductInstancesResponseBodyProductInstances) *DescribeProductInstancesResponseBody
	GetProductInstances() []*DescribeProductInstancesResponseBodyProductInstances
	SetRequestId(v string) *DescribeProductInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeProductInstancesResponseBody
	GetTotalCount() *int64
}

type DescribeProductInstancesResponseBody struct {
	// The information about the instances.
	ProductInstances []*DescribeProductInstancesResponseBodyProductInstances `json:"ProductInstances,omitempty" xml:"ProductInstances,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// FDCBAE1E-2B3F-5C13-AD20-844B9473****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeProductInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProductInstancesResponseBody) GetProductInstances() []*DescribeProductInstancesResponseBodyProductInstances {
	return s.ProductInstances
}

func (s *DescribeProductInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProductInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeProductInstancesResponseBody) SetProductInstances(v []*DescribeProductInstancesResponseBodyProductInstances) *DescribeProductInstancesResponseBody {
	s.ProductInstances = v
	return s
}

func (s *DescribeProductInstancesResponseBody) SetRequestId(v string) *DescribeProductInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProductInstancesResponseBody) SetTotalCount(v int64) *DescribeProductInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeProductInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeProductInstancesResponseBodyProductInstances struct {
	AccessInstanceId       *string                                                                       `json:"AccessInstanceId,omitempty" xml:"AccessInstanceId,omitempty"`
	AccessPortAndProtocols []*DescribeProductInstancesResponseBodyProductInstancesAccessPortAndProtocols `json:"AccessPortAndProtocols,omitempty" xml:"AccessPortAndProtocols,omitempty" type:"Repeated"`
	AccessPorts            []*int32                                                                      `json:"AccessPorts,omitempty" xml:"AccessPorts,omitempty" type:"Repeated"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 1704********9107
	OwnerUserId                  *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	ResourceInstanceAccessStatus *string `json:"ResourceInstanceAccessStatus,omitempty" xml:"ResourceInstanceAccessStatus,omitempty"`
	ResourceInstanceEdition      *string `json:"ResourceInstanceEdition,omitempty" xml:"ResourceInstanceEdition,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-2ze1tm4pvghp****cluv
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The IP address of the instance that is added to WAF.
	//
	// example:
	//
	// 1.X.X.1
	ResourceInstanceIp *string `json:"ResourceInstanceIp,omitempty" xml:"ResourceInstanceIp,omitempty"`
	// The name of the instance that is added to WAF.
	//
	// example:
	//
	// demoInstanceName
	ResourceInstanceName *string `json:"ResourceInstanceName,omitempty" xml:"ResourceInstanceName,omitempty"`
	// Deprecated
	//
	// The public IP address of the instance.
	//
	// example:
	//
	// 1.X.X.1
	ResourceIp *string `json:"ResourceIp,omitempty" xml:"ResourceIp,omitempty"`
	// Deprecated
	//
	// The name of the instance.
	//
	// example:
	//
	// ecs-test
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The information about the ports.
	ResourcePorts []*DescribeProductInstancesResponseBodyProductInstancesResourcePorts `json:"ResourcePorts,omitempty" xml:"ResourcePorts,omitempty" type:"Repeated"`
	// The cloud service to which the instance belongs. Valid values:
	//
	// 	- **clb4**: Layer 4 CLB.
	//
	// 	- **clb7**: Layer 7 CLB.
	//
	// 	- **ecs**: ECS.
	//
	// example:
	//
	// clb4
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
	// The region ID of the instance. Valid values:
	//
	// 	- **cn-chengdu**: China (Chengdu).
	//
	// 	- **cn-beijing**: China (Beijing).
	//
	// 	- **cn-zhangjiakou**: China (Zhangjiakou).
	//
	// 	- **cn-hangzhou**: China (Hangzhou).
	//
	// 	- **cn-shanghai**: China (Shanghai).
	//
	// 	- **cn-shenzhen**: China (Shenzhen).
	//
	// 	- **cn-qingdao**: China (Qingdao).
	//
	// 	- **cn-hongkong**: China (Hong Kong).
	//
	// 	- **ap-southeast-3**: Malaysia (Kuala Lumpur).
	//
	// 	- **ap-southeast-5**: Indonesia (Jakarta).
	//
	// example:
	//
	// cn-beijing
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
}

func (s DescribeProductInstancesResponseBodyProductInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductInstancesResponseBodyProductInstances) GoString() string {
	return s.String()
}

func (s *DescribeProductInstancesResponseBodyProductInstances) GetAccessInstanceId() *string {
	return s.AccessInstanceId
}

func (s *DescribeProductInstancesResponseBodyProductInstances) GetAccessPortAndProtocols() []*DescribeProductInstancesResponseBodyProductInstancesAccessPortAndProtocols {
	return s.AccessPortAndProtocols
}

func (s *DescribeProductInstancesResponseBodyProductInstances) GetAccessPorts() []*int32 {
	return s.AccessPorts
}

func (s *DescribeProductInstancesResponseBodyProductInstances) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *DescribeProductInstancesResponseBodyProductInstances) GetResourceInstanceAccessStatus() *string {
	return s.ResourceInstanceAccessStatus
}

func (s *DescribeProductInstancesResponseBodyProductInstances) GetResourceInstanceEdition() *string {
	return s.ResourceInstanceEdition
}

func (s *DescribeProductInstancesResponseBodyProductInstances) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *DescribeProductInstancesResponseBodyProductInstances) GetResourceInstanceIp() *string {
	return s.ResourceInstanceIp
}

func (s *DescribeProductInstancesResponseBodyProductInstances) GetResourceInstanceName() *string {
	return s.ResourceInstanceName
}

func (s *DescribeProductInstancesResponseBodyProductInstances) GetResourceIp() *string {
	return s.ResourceIp
}

func (s *DescribeProductInstancesResponseBodyProductInstances) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeProductInstancesResponseBodyProductInstances) GetResourcePorts() []*DescribeProductInstancesResponseBodyProductInstancesResourcePorts {
	return s.ResourcePorts
}

func (s *DescribeProductInstancesResponseBodyProductInstances) GetResourceProduct() *string {
	return s.ResourceProduct
}

func (s *DescribeProductInstancesResponseBodyProductInstances) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *DescribeProductInstancesResponseBodyProductInstances) SetAccessInstanceId(v string) *DescribeProductInstancesResponseBodyProductInstances {
	s.AccessInstanceId = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstances) SetAccessPortAndProtocols(v []*DescribeProductInstancesResponseBodyProductInstancesAccessPortAndProtocols) *DescribeProductInstancesResponseBodyProductInstances {
	s.AccessPortAndProtocols = v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstances) SetAccessPorts(v []*int32) *DescribeProductInstancesResponseBodyProductInstances {
	s.AccessPorts = v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstances) SetOwnerUserId(v string) *DescribeProductInstancesResponseBodyProductInstances {
	s.OwnerUserId = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstances) SetResourceInstanceAccessStatus(v string) *DescribeProductInstancesResponseBodyProductInstances {
	s.ResourceInstanceAccessStatus = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstances) SetResourceInstanceEdition(v string) *DescribeProductInstancesResponseBodyProductInstances {
	s.ResourceInstanceEdition = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstances) SetResourceInstanceId(v string) *DescribeProductInstancesResponseBodyProductInstances {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstances) SetResourceInstanceIp(v string) *DescribeProductInstancesResponseBodyProductInstances {
	s.ResourceInstanceIp = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstances) SetResourceInstanceName(v string) *DescribeProductInstancesResponseBodyProductInstances {
	s.ResourceInstanceName = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstances) SetResourceIp(v string) *DescribeProductInstancesResponseBodyProductInstances {
	s.ResourceIp = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstances) SetResourceName(v string) *DescribeProductInstancesResponseBodyProductInstances {
	s.ResourceName = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstances) SetResourcePorts(v []*DescribeProductInstancesResponseBodyProductInstancesResourcePorts) *DescribeProductInstancesResponseBodyProductInstances {
	s.ResourcePorts = v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstances) SetResourceProduct(v string) *DescribeProductInstancesResponseBodyProductInstances {
	s.ResourceProduct = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstances) SetResourceRegionId(v string) *DescribeProductInstancesResponseBodyProductInstances {
	s.ResourceRegionId = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeProductInstancesResponseBodyProductInstancesAccessPortAndProtocols struct {
	CertificateIds []*string `json:"CertificateIds,omitempty" xml:"CertificateIds,omitempty" type:"Repeated"`
	Port           *int32    `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol       *string   `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeProductInstancesResponseBodyProductInstancesAccessPortAndProtocols) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductInstancesResponseBodyProductInstancesAccessPortAndProtocols) GoString() string {
	return s.String()
}

func (s *DescribeProductInstancesResponseBodyProductInstancesAccessPortAndProtocols) GetCertificateIds() []*string {
	return s.CertificateIds
}

func (s *DescribeProductInstancesResponseBodyProductInstancesAccessPortAndProtocols) GetPort() *int32 {
	return s.Port
}

func (s *DescribeProductInstancesResponseBodyProductInstancesAccessPortAndProtocols) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeProductInstancesResponseBodyProductInstancesAccessPortAndProtocols) SetCertificateIds(v []*string) *DescribeProductInstancesResponseBodyProductInstancesAccessPortAndProtocols {
	s.CertificateIds = v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstancesAccessPortAndProtocols) SetPort(v int32) *DescribeProductInstancesResponseBodyProductInstancesAccessPortAndProtocols {
	s.Port = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstancesAccessPortAndProtocols) SetProtocol(v string) *DescribeProductInstancesResponseBodyProductInstancesAccessPortAndProtocols {
	s.Protocol = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstancesAccessPortAndProtocols) Validate() error {
	return dara.Validate(s)
}

type DescribeProductInstancesResponseBodyProductInstancesResourcePorts struct {
	// The information about the certificates.
	Certificates []*DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The port number.
	//
	// example:
	//
	// 443
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type. Valid values:
	//
	// 	- **http**
	//
	// 	- **https**
	//
	// example:
	//
	// https
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeProductInstancesResponseBodyProductInstancesResourcePorts) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductInstancesResponseBodyProductInstancesResourcePorts) GoString() string {
	return s.String()
}

func (s *DescribeProductInstancesResponseBodyProductInstancesResourcePorts) GetCertificates() []*DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates {
	return s.Certificates
}

func (s *DescribeProductInstancesResponseBodyProductInstancesResourcePorts) GetPort() *int32 {
	return s.Port
}

func (s *DescribeProductInstancesResponseBodyProductInstancesResourcePorts) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeProductInstancesResponseBodyProductInstancesResourcePorts) SetCertificates(v []*DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates) *DescribeProductInstancesResponseBodyProductInstancesResourcePorts {
	s.Certificates = v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstancesResourcePorts) SetPort(v int32) *DescribeProductInstancesResponseBodyProductInstancesResourcePorts {
	s.Port = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstancesResourcePorts) SetProtocol(v string) *DescribeProductInstancesResponseBodyProductInstancesResourcePorts {
	s.Protocol = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstancesResourcePorts) Validate() error {
	return dara.Validate(s)
}

type DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates struct {
	AppliedType *string `json:"AppliedType,omitempty" xml:"AppliedType,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// 10106183
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// trafficxxxx.cn
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates) GoString() string {
	return s.String()
}

func (s *DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates) GetAppliedType() *string {
	return s.AppliedType
}

func (s *DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates) GetCertificateId() *string {
	return s.CertificateId
}

func (s *DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates) GetCertificateName() *string {
	return s.CertificateName
}

func (s *DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates) GetDomain() *string {
	return s.Domain
}

func (s *DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates) SetAppliedType(v string) *DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates {
	s.AppliedType = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates) SetCertificateId(v string) *DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates {
	s.CertificateId = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates) SetCertificateName(v string) *DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates {
	s.CertificateName = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates) SetDomain(v string) *DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates {
	s.Domain = &v
	return s
}

func (s *DescribeProductInstancesResponseBodyProductInstancesResourcePortsCertificates) Validate() error {
	return dara.Validate(s)
}
