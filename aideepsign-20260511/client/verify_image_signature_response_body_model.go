// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyImageSignatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VerifyImageSignatureResponseBody
	GetCode() *string
	SetHttpStatusCode(v int64) *VerifyImageSignatureResponseBody
	GetHttpStatusCode() *int64
	SetIssuer(v *VerifyImageSignatureResponseBodyIssuer) *VerifyImageSignatureResponseBody
	GetIssuer() *VerifyImageSignatureResponseBodyIssuer
	SetIssuerTrusted(v bool) *VerifyImageSignatureResponseBody
	GetIssuerTrusted() *bool
	SetManifest(v *VerifyImageSignatureResponseBodyManifest) *VerifyImageSignatureResponseBody
	GetManifest() *VerifyImageSignatureResponseBodyManifest
	SetMessage(v string) *VerifyImageSignatureResponseBody
	GetMessage() *string
	SetRequestId(v string) *VerifyImageSignatureResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VerifyImageSignatureResponseBody
	GetSuccess() *bool
	SetVerificationState(v string) *VerifyImageSignatureResponseBody
	GetVerificationState() *string
}

type VerifyImageSignatureResponseBody struct {
	// The business error code. The value "OK" is returned if the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. The value `200` is returned if the request is successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The issuer information.
	Issuer *VerifyImageSignatureResponseBodyIssuer `json:"Issuer,omitempty" xml:"Issuer,omitempty" type:"Struct"`
	// Indicates whether the issuer is trusted. A value of true indicates that the issuer certificate is issued by a trusted CA.
	IssuerTrusted *bool `json:"IssuerTrusted,omitempty" xml:"IssuerTrusted,omitempty"`
	// The content credentials manifest information. This parameter is returned only when the image contains a C2PA signature.
	Manifest *VerifyImageSignatureResponseBodyManifest `json:"Manifest,omitempty" xml:"Manifest,omitempty" type:"Struct"`
	// The additional information. The value `success` is returned if the request is successful.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The signature verification status. Valid values:
	//
	// - Valid: The signature is valid.
	//
	// - Invalid: The signature is invalid or has been tampered with.
	//
	// - Trusted: The signature is valid and trusted.
	//
	// example:
	//
	// Valid
	VerificationState *string `json:"VerificationState,omitempty" xml:"VerificationState,omitempty"`
}

func (s VerifyImageSignatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyImageSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyImageSignatureResponseBody) GetCode() *string {
	return s.Code
}

func (s *VerifyImageSignatureResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *VerifyImageSignatureResponseBody) GetIssuer() *VerifyImageSignatureResponseBodyIssuer {
	return s.Issuer
}

func (s *VerifyImageSignatureResponseBody) GetIssuerTrusted() *bool {
	return s.IssuerTrusted
}

func (s *VerifyImageSignatureResponseBody) GetManifest() *VerifyImageSignatureResponseBodyManifest {
	return s.Manifest
}

func (s *VerifyImageSignatureResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VerifyImageSignatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyImageSignatureResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VerifyImageSignatureResponseBody) GetVerificationState() *string {
	return s.VerificationState
}

func (s *VerifyImageSignatureResponseBody) SetCode(v string) *VerifyImageSignatureResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyImageSignatureResponseBody) SetHttpStatusCode(v int64) *VerifyImageSignatureResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *VerifyImageSignatureResponseBody) SetIssuer(v *VerifyImageSignatureResponseBodyIssuer) *VerifyImageSignatureResponseBody {
	s.Issuer = v
	return s
}

func (s *VerifyImageSignatureResponseBody) SetIssuerTrusted(v bool) *VerifyImageSignatureResponseBody {
	s.IssuerTrusted = &v
	return s
}

func (s *VerifyImageSignatureResponseBody) SetManifest(v *VerifyImageSignatureResponseBodyManifest) *VerifyImageSignatureResponseBody {
	s.Manifest = v
	return s
}

func (s *VerifyImageSignatureResponseBody) SetMessage(v string) *VerifyImageSignatureResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyImageSignatureResponseBody) SetRequestId(v string) *VerifyImageSignatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyImageSignatureResponseBody) SetSuccess(v bool) *VerifyImageSignatureResponseBody {
	s.Success = &v
	return s
}

func (s *VerifyImageSignatureResponseBody) SetVerificationState(v string) *VerifyImageSignatureResponseBody {
	s.VerificationState = &v
	return s
}

func (s *VerifyImageSignatureResponseBody) Validate() error {
	if s.Issuer != nil {
		if err := s.Issuer.Validate(); err != nil {
			return err
		}
	}
	if s.Manifest != nil {
		if err := s.Manifest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifyImageSignatureResponseBodyIssuer struct {
	// The common name (CN) of the issuer.
	//
	// example:
	//
	// AIDeepSign User Certificate
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The organization name (O) of the issuer.
	//
	// example:
	//
	// Alibaba Cloud
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
}

func (s VerifyImageSignatureResponseBodyIssuer) String() string {
	return dara.Prettify(s)
}

func (s VerifyImageSignatureResponseBodyIssuer) GoString() string {
	return s.String()
}

func (s *VerifyImageSignatureResponseBodyIssuer) GetCommonName() *string {
	return s.CommonName
}

func (s *VerifyImageSignatureResponseBodyIssuer) GetOrganization() *string {
	return s.Organization
}

func (s *VerifyImageSignatureResponseBodyIssuer) SetCommonName(v string) *VerifyImageSignatureResponseBodyIssuer {
	s.CommonName = &v
	return s
}

func (s *VerifyImageSignatureResponseBodyIssuer) SetOrganization(v string) *VerifyImageSignatureResponseBodyIssuer {
	s.Organization = &v
	return s
}

func (s *VerifyImageSignatureResponseBodyIssuer) Validate() error {
	return dara.Validate(s)
}

type VerifyImageSignatureResponseBodyManifest struct {
	// The list of assertions, which contains metadata such as the origin and editing history of the image.
	Assertions []*VerifyImageSignatureResponseBodyManifestAssertions `json:"Assertions,omitempty" xml:"Assertions,omitempty" type:"Repeated"`
	// The signature details.
	SignatureInfo *VerifyImageSignatureResponseBodyManifestSignatureInfo `json:"SignatureInfo,omitempty" xml:"SignatureInfo,omitempty" type:"Struct"`
}

func (s VerifyImageSignatureResponseBodyManifest) String() string {
	return dara.Prettify(s)
}

func (s VerifyImageSignatureResponseBodyManifest) GoString() string {
	return s.String()
}

func (s *VerifyImageSignatureResponseBodyManifest) GetAssertions() []*VerifyImageSignatureResponseBodyManifestAssertions {
	return s.Assertions
}

func (s *VerifyImageSignatureResponseBodyManifest) GetSignatureInfo() *VerifyImageSignatureResponseBodyManifestSignatureInfo {
	return s.SignatureInfo
}

func (s *VerifyImageSignatureResponseBodyManifest) SetAssertions(v []*VerifyImageSignatureResponseBodyManifestAssertions) *VerifyImageSignatureResponseBodyManifest {
	s.Assertions = v
	return s
}

func (s *VerifyImageSignatureResponseBodyManifest) SetSignatureInfo(v *VerifyImageSignatureResponseBodyManifestSignatureInfo) *VerifyImageSignatureResponseBodyManifest {
	s.SignatureInfo = v
	return s
}

func (s *VerifyImageSignatureResponseBodyManifest) Validate() error {
	if s.Assertions != nil {
		for _, item := range s.Assertions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SignatureInfo != nil {
		if err := s.SignatureInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifyImageSignatureResponseBodyManifestAssertions struct {
	// The assertion data, which is detailed metadata in JSON format.
	//
	// example:
	//
	// {"actions":[{"action":"c2pa.created"}]}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The assertion label, such as c2pa.actions or stds.exif.
	//
	// example:
	//
	// c2pa.actions
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s VerifyImageSignatureResponseBodyManifestAssertions) String() string {
	return dara.Prettify(s)
}

func (s VerifyImageSignatureResponseBodyManifestAssertions) GoString() string {
	return s.String()
}

func (s *VerifyImageSignatureResponseBodyManifestAssertions) GetData() *string {
	return s.Data
}

func (s *VerifyImageSignatureResponseBodyManifestAssertions) GetLabel() *string {
	return s.Label
}

func (s *VerifyImageSignatureResponseBodyManifestAssertions) SetData(v string) *VerifyImageSignatureResponseBodyManifestAssertions {
	s.Data = &v
	return s
}

func (s *VerifyImageSignatureResponseBodyManifestAssertions) SetLabel(v string) *VerifyImageSignatureResponseBodyManifestAssertions {
	s.Label = &v
	return s
}

func (s *VerifyImageSignatureResponseBodyManifestAssertions) Validate() error {
	return dara.Validate(s)
}

type VerifyImageSignatureResponseBodyManifestSignatureInfo struct {
	// The signature algorithm, such as `ps256` or `es256`.
	//
	// example:
	//
	// ps256
	Alg *string `json:"Alg,omitempty" xml:"Alg,omitempty"`
	// The distinguished name (DN) of the signing certificate issuer.
	//
	// example:
	//
	// CN=AIDeepSign CA,O=Alibaba Cloud
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The signing time in ISO 8601 format.
	//
	// example:
	//
	// 2026-06-18T10:30:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s VerifyImageSignatureResponseBodyManifestSignatureInfo) String() string {
	return dara.Prettify(s)
}

func (s VerifyImageSignatureResponseBodyManifestSignatureInfo) GoString() string {
	return s.String()
}

func (s *VerifyImageSignatureResponseBodyManifestSignatureInfo) GetAlg() *string {
	return s.Alg
}

func (s *VerifyImageSignatureResponseBodyManifestSignatureInfo) GetIssuer() *string {
	return s.Issuer
}

func (s *VerifyImageSignatureResponseBodyManifestSignatureInfo) GetTime() *string {
	return s.Time
}

func (s *VerifyImageSignatureResponseBodyManifestSignatureInfo) SetAlg(v string) *VerifyImageSignatureResponseBodyManifestSignatureInfo {
	s.Alg = &v
	return s
}

func (s *VerifyImageSignatureResponseBodyManifestSignatureInfo) SetIssuer(v string) *VerifyImageSignatureResponseBodyManifestSignatureInfo {
	s.Issuer = &v
	return s
}

func (s *VerifyImageSignatureResponseBodyManifestSignatureInfo) SetTime(v string) *VerifyImageSignatureResponseBodyManifestSignatureInfo {
	s.Time = &v
	return s
}

func (s *VerifyImageSignatureResponseBodyManifestSignatureInfo) Validate() error {
	return dara.Validate(s)
}
