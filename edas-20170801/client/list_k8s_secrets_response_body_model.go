// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8sSecretsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListK8sSecretsResponseBody
	GetCode() *int32
	SetMessage(v string) *ListK8sSecretsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListK8sSecretsResponseBody
	GetRequestId() *string
	SetResult(v *ListK8sSecretsResponseBodyResult) *ListK8sSecretsResponseBody
	GetResult() *ListK8sSecretsResponseBodyResult
}

type ListK8sSecretsResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// b197-40ab-9155-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned query results of Kubernetes Secrets.
	Result *ListK8sSecretsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListK8sSecretsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListK8sSecretsResponseBody) GoString() string {
	return s.String()
}

func (s *ListK8sSecretsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListK8sSecretsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListK8sSecretsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListK8sSecretsResponseBody) GetResult() *ListK8sSecretsResponseBodyResult {
	return s.Result
}

func (s *ListK8sSecretsResponseBody) SetCode(v int32) *ListK8sSecretsResponseBody {
	s.Code = &v
	return s
}

func (s *ListK8sSecretsResponseBody) SetMessage(v string) *ListK8sSecretsResponseBody {
	s.Message = &v
	return s
}

func (s *ListK8sSecretsResponseBody) SetRequestId(v string) *ListK8sSecretsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListK8sSecretsResponseBody) SetResult(v *ListK8sSecretsResponseBodyResult) *ListK8sSecretsResponseBody {
	s.Result = v
	return s
}

func (s *ListK8sSecretsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListK8sSecretsResponseBodyResult struct {
	// The information about Kubernetes Secrets.
	Secrets []*ListK8sSecretsResponseBodyResultSecrets `json:"Secrets,omitempty" xml:"Secrets,omitempty" type:"Repeated"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 6
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListK8sSecretsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListK8sSecretsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListK8sSecretsResponseBodyResult) GetSecrets() []*ListK8sSecretsResponseBodyResultSecrets {
	return s.Secrets
}

func (s *ListK8sSecretsResponseBodyResult) GetTotal() *int32 {
	return s.Total
}

func (s *ListK8sSecretsResponseBodyResult) SetSecrets(v []*ListK8sSecretsResponseBodyResultSecrets) *ListK8sSecretsResponseBodyResult {
	s.Secrets = v
	return s
}

func (s *ListK8sSecretsResponseBodyResult) SetTotal(v int32) *ListK8sSecretsResponseBodyResult {
	s.Total = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResult) Validate() error {
	if s.Secrets != nil {
		for _, item := range s.Secrets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListK8sSecretsResponseBodyResultSecrets struct {
	// Indicates whether the data is Base64-encoded. Valid values:
	//
	// 	- true: The data is Base64-encoded.
	//
	// 	- false: The data is not Base64-encoded.
	//
	// example:
	//
	// false
	Base64Encoded *bool `json:"Base64Encoded,omitempty" xml:"Base64Encoded,omitempty"`
	// The details of the Secure Sockets Layer (SSL) certificate.
	CertDetail *ListK8sSecretsResponseBodyResultSecretsCertDetail `json:"CertDetail,omitempty" xml:"CertDetail,omitempty" type:"Struct"`
	// The ID of the certificate provided by Alibaba Cloud Certificate Management Service.
	//
	// example:
	//
	// 123456
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The region in which the certificate is stored.
	//
	// example:
	//
	// cn-hangzhou
	CertRegionId *string `json:"CertRegionId,omitempty" xml:"CertRegionId,omitempty"`
	// The ID of the cluster in Enterprise Distributed Application Service (EDAS).
	//
	// example:
	//
	// 4472a6d3-f01d-4087-85a7-3dc52********
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// test-cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The time when the Secret was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*hh:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-05-26T02:57:02Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The data of the Kubernetes Secret.
	Data []*ListK8sSecretsResponseBodyResultSecretsData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The name of the Secret. The name must start with a letter, and can contain digits, letters, and hyphens (-). It can be up to 63 characters in length.
	//
	// example:
	//
	// my-secret
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace of the Kubernetes cluster.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Applications that use the Secret.
	RelatedApps []*ListK8sSecretsResponseBodyResultSecretsRelatedApps `json:"RelatedApps,omitempty" xml:"RelatedApps,omitempty" type:"Repeated"`
	// Rules in the Ingress that is associated with the Secret.
	RelatedIngressRules []*ListK8sSecretsResponseBodyResultSecretsRelatedIngressRules `json:"RelatedIngressRules,omitempty" xml:"RelatedIngressRules,omitempty" type:"Repeated"`
	// The type of the Secret. Valid values:
	//
	// 	- Opaque: user-defined data
	//
	// 	- kubernetes.io/tls: Transport Layer Security (TLS) certificate
	//
	// example:
	//
	// Opaque
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListK8sSecretsResponseBodyResultSecrets) String() string {
	return dara.Prettify(s)
}

func (s ListK8sSecretsResponseBodyResultSecrets) GoString() string {
	return s.String()
}

func (s *ListK8sSecretsResponseBodyResultSecrets) GetBase64Encoded() *bool {
	return s.Base64Encoded
}

func (s *ListK8sSecretsResponseBodyResultSecrets) GetCertDetail() *ListK8sSecretsResponseBodyResultSecretsCertDetail {
	return s.CertDetail
}

func (s *ListK8sSecretsResponseBodyResultSecrets) GetCertId() *string {
	return s.CertId
}

func (s *ListK8sSecretsResponseBodyResultSecrets) GetCertRegionId() *string {
	return s.CertRegionId
}

func (s *ListK8sSecretsResponseBodyResultSecrets) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListK8sSecretsResponseBodyResultSecrets) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListK8sSecretsResponseBodyResultSecrets) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListK8sSecretsResponseBodyResultSecrets) GetData() []*ListK8sSecretsResponseBodyResultSecretsData {
	return s.Data
}

func (s *ListK8sSecretsResponseBodyResultSecrets) GetName() *string {
	return s.Name
}

func (s *ListK8sSecretsResponseBodyResultSecrets) GetNamespace() *string {
	return s.Namespace
}

func (s *ListK8sSecretsResponseBodyResultSecrets) GetRelatedApps() []*ListK8sSecretsResponseBodyResultSecretsRelatedApps {
	return s.RelatedApps
}

func (s *ListK8sSecretsResponseBodyResultSecrets) GetRelatedIngressRules() []*ListK8sSecretsResponseBodyResultSecretsRelatedIngressRules {
	return s.RelatedIngressRules
}

func (s *ListK8sSecretsResponseBodyResultSecrets) GetType() *string {
	return s.Type
}

func (s *ListK8sSecretsResponseBodyResultSecrets) SetBase64Encoded(v bool) *ListK8sSecretsResponseBodyResultSecrets {
	s.Base64Encoded = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecrets) SetCertDetail(v *ListK8sSecretsResponseBodyResultSecretsCertDetail) *ListK8sSecretsResponseBodyResultSecrets {
	s.CertDetail = v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecrets) SetCertId(v string) *ListK8sSecretsResponseBodyResultSecrets {
	s.CertId = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecrets) SetCertRegionId(v string) *ListK8sSecretsResponseBodyResultSecrets {
	s.CertRegionId = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecrets) SetClusterId(v string) *ListK8sSecretsResponseBodyResultSecrets {
	s.ClusterId = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecrets) SetClusterName(v string) *ListK8sSecretsResponseBodyResultSecrets {
	s.ClusterName = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecrets) SetCreationTime(v string) *ListK8sSecretsResponseBodyResultSecrets {
	s.CreationTime = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecrets) SetData(v []*ListK8sSecretsResponseBodyResultSecretsData) *ListK8sSecretsResponseBodyResultSecrets {
	s.Data = v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecrets) SetName(v string) *ListK8sSecretsResponseBodyResultSecrets {
	s.Name = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecrets) SetNamespace(v string) *ListK8sSecretsResponseBodyResultSecrets {
	s.Namespace = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecrets) SetRelatedApps(v []*ListK8sSecretsResponseBodyResultSecretsRelatedApps) *ListK8sSecretsResponseBodyResultSecrets {
	s.RelatedApps = v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecrets) SetRelatedIngressRules(v []*ListK8sSecretsResponseBodyResultSecretsRelatedIngressRules) *ListK8sSecretsResponseBodyResultSecrets {
	s.RelatedIngressRules = v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecrets) SetType(v string) *ListK8sSecretsResponseBodyResultSecrets {
	s.Type = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecrets) Validate() error {
	if s.CertDetail != nil {
		if err := s.CertDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RelatedApps != nil {
		for _, item := range s.RelatedApps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RelatedIngressRules != nil {
		for _, item := range s.RelatedIngressRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListK8sSecretsResponseBodyResultSecretsCertDetail struct {
	// Domain names that are associated with the SSL certificate.
	DomainNames []*string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty" type:"Repeated"`
	// The time when the SSL certificate expired.
	//
	// example:
	//
	// 2022-02-22T02:32:41Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The certificate authority (CA) that issued the SSL certificate.
	//
	// example:
	//
	// CN=GlobalSign Root CA, OU=Root CA, O=GlobalSign nv-sa, C=BE
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The time when the SSL certificate started to take effect.
	//
	// example:
	//
	// 2022-01-02T22:40:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the SSL certificate. Valid values:
	//
	// 	- normal: The SSL certificate is valid.
	//
	// 	- invalid: The SSL certificate is invalid.
	//
	// 	- expired: The SSL certificate has expired.
	//
	// 	- not_yet_valid: The SSL certificate is currently invalid.
	//
	// 	- about_to_expire: The SSL certificate is about to expire.
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListK8sSecretsResponseBodyResultSecretsCertDetail) String() string {
	return dara.Prettify(s)
}

func (s ListK8sSecretsResponseBodyResultSecretsCertDetail) GoString() string {
	return s.String()
}

func (s *ListK8sSecretsResponseBodyResultSecretsCertDetail) GetDomainNames() []*string {
	return s.DomainNames
}

func (s *ListK8sSecretsResponseBodyResultSecretsCertDetail) GetEndTime() *string {
	return s.EndTime
}

func (s *ListK8sSecretsResponseBodyResultSecretsCertDetail) GetIssuer() *string {
	return s.Issuer
}

func (s *ListK8sSecretsResponseBodyResultSecretsCertDetail) GetStartTime() *string {
	return s.StartTime
}

func (s *ListK8sSecretsResponseBodyResultSecretsCertDetail) GetStatus() *string {
	return s.Status
}

func (s *ListK8sSecretsResponseBodyResultSecretsCertDetail) SetDomainNames(v []*string) *ListK8sSecretsResponseBodyResultSecretsCertDetail {
	s.DomainNames = v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecretsCertDetail) SetEndTime(v string) *ListK8sSecretsResponseBodyResultSecretsCertDetail {
	s.EndTime = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecretsCertDetail) SetIssuer(v string) *ListK8sSecretsResponseBodyResultSecretsCertDetail {
	s.Issuer = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecretsCertDetail) SetStartTime(v string) *ListK8sSecretsResponseBodyResultSecretsCertDetail {
	s.StartTime = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecretsCertDetail) SetStatus(v string) *ListK8sSecretsResponseBodyResultSecretsCertDetail {
	s.Status = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecretsCertDetail) Validate() error {
	return dara.Validate(s)
}

type ListK8sSecretsResponseBodyResultSecretsData struct {
	// The user-defined key of the Kubernetes Secret.
	//
	// example:
	//
	// name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The user-defined value of the Kubernetes Secret.
	//
	// example:
	//
	// william
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListK8sSecretsResponseBodyResultSecretsData) String() string {
	return dara.Prettify(s)
}

func (s ListK8sSecretsResponseBodyResultSecretsData) GoString() string {
	return s.String()
}

func (s *ListK8sSecretsResponseBodyResultSecretsData) GetKey() *string {
	return s.Key
}

func (s *ListK8sSecretsResponseBodyResultSecretsData) GetValue() *string {
	return s.Value
}

func (s *ListK8sSecretsResponseBodyResultSecretsData) SetKey(v string) *ListK8sSecretsResponseBodyResultSecretsData {
	s.Key = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecretsData) SetValue(v string) *ListK8sSecretsResponseBodyResultSecretsData {
	s.Value = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecretsData) Validate() error {
	return dara.Validate(s)
}

type ListK8sSecretsResponseBodyResultSecretsRelatedApps struct {
	// The ID of the application.
	//
	// example:
	//
	// b08eeb18-8946-410c-a1ea-dbbc********
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// my-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s ListK8sSecretsResponseBodyResultSecretsRelatedApps) String() string {
	return dara.Prettify(s)
}

func (s ListK8sSecretsResponseBodyResultSecretsRelatedApps) GoString() string {
	return s.String()
}

func (s *ListK8sSecretsResponseBodyResultSecretsRelatedApps) GetAppId() *string {
	return s.AppId
}

func (s *ListK8sSecretsResponseBodyResultSecretsRelatedApps) GetAppName() *string {
	return s.AppName
}

func (s *ListK8sSecretsResponseBodyResultSecretsRelatedApps) SetAppId(v string) *ListK8sSecretsResponseBodyResultSecretsRelatedApps {
	s.AppId = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecretsRelatedApps) SetAppName(v string) *ListK8sSecretsResponseBodyResultSecretsRelatedApps {
	s.AppName = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecretsRelatedApps) Validate() error {
	return dara.Validate(s)
}

type ListK8sSecretsResponseBodyResultSecretsRelatedIngressRules struct {
	// The name of the rule in the Ingress.
	//
	// example:
	//
	// testrulename
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespaces of the Kubernetes cluster.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Aplications that are associated with the Ingress.
	RelatedApps []*ListK8sSecretsResponseBodyResultSecretsRelatedIngressRulesRelatedApps `json:"RelatedApps,omitempty" xml:"RelatedApps,omitempty" type:"Repeated"`
}

func (s ListK8sSecretsResponseBodyResultSecretsRelatedIngressRules) String() string {
	return dara.Prettify(s)
}

func (s ListK8sSecretsResponseBodyResultSecretsRelatedIngressRules) GoString() string {
	return s.String()
}

func (s *ListK8sSecretsResponseBodyResultSecretsRelatedIngressRules) GetName() *string {
	return s.Name
}

func (s *ListK8sSecretsResponseBodyResultSecretsRelatedIngressRules) GetNamespace() *string {
	return s.Namespace
}

func (s *ListK8sSecretsResponseBodyResultSecretsRelatedIngressRules) GetRelatedApps() []*ListK8sSecretsResponseBodyResultSecretsRelatedIngressRulesRelatedApps {
	return s.RelatedApps
}

func (s *ListK8sSecretsResponseBodyResultSecretsRelatedIngressRules) SetName(v string) *ListK8sSecretsResponseBodyResultSecretsRelatedIngressRules {
	s.Name = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecretsRelatedIngressRules) SetNamespace(v string) *ListK8sSecretsResponseBodyResultSecretsRelatedIngressRules {
	s.Namespace = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecretsRelatedIngressRules) SetRelatedApps(v []*ListK8sSecretsResponseBodyResultSecretsRelatedIngressRulesRelatedApps) *ListK8sSecretsResponseBodyResultSecretsRelatedIngressRules {
	s.RelatedApps = v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecretsRelatedIngressRules) Validate() error {
	if s.RelatedApps != nil {
		for _, item := range s.RelatedApps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListK8sSecretsResponseBodyResultSecretsRelatedIngressRulesRelatedApps struct {
	// The ID of the application.
	//
	// example:
	//
	// 6dc74432-5a35-4e68-8aaa-3700********
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the EDAS application.
	//
	// example:
	//
	// app-test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s ListK8sSecretsResponseBodyResultSecretsRelatedIngressRulesRelatedApps) String() string {
	return dara.Prettify(s)
}

func (s ListK8sSecretsResponseBodyResultSecretsRelatedIngressRulesRelatedApps) GoString() string {
	return s.String()
}

func (s *ListK8sSecretsResponseBodyResultSecretsRelatedIngressRulesRelatedApps) GetAppId() *string {
	return s.AppId
}

func (s *ListK8sSecretsResponseBodyResultSecretsRelatedIngressRulesRelatedApps) GetAppName() *string {
	return s.AppName
}

func (s *ListK8sSecretsResponseBodyResultSecretsRelatedIngressRulesRelatedApps) SetAppId(v string) *ListK8sSecretsResponseBodyResultSecretsRelatedIngressRulesRelatedApps {
	s.AppId = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecretsRelatedIngressRulesRelatedApps) SetAppName(v string) *ListK8sSecretsResponseBodyResultSecretsRelatedIngressRulesRelatedApps {
	s.AppName = &v
	return s
}

func (s *ListK8sSecretsResponseBodyResultSecretsRelatedIngressRulesRelatedApps) Validate() error {
	return dara.Validate(s)
}
