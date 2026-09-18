// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitUploadPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFilePath(v string) *CreateTransitUploadPolicyResponseBody
	GetFilePath() *string
	SetPolicyInfo(v *CreateTransitUploadPolicyResponseBodyPolicyInfo) *CreateTransitUploadPolicyResponseBody
	GetPolicyInfo() *CreateTransitUploadPolicyResponseBodyPolicyInfo
	SetRequestId(v string) *CreateTransitUploadPolicyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTransitUploadPolicyResponseBody
	GetSuccess() *bool
	SetTransitId(v string) *CreateTransitUploadPolicyResponseBody
	GetTransitId() *string
}

type CreateTransitUploadPolicyResponseBody struct {
	// The object storage key, which is also the `key` field in the PostObject form.
	//
	// example:
	//
	// skill-bundle/tenant-demo/user-demo/20260904120000_0123456789abcdef0123456789abcdef.zip
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The upload policy object. For the complete list of subfields, see the following table.
	PolicyInfo *CreateTransitUploadPolicyResponseBodyPolicyInfo `json:"PolicyInfo,omitempty" xml:"PolicyInfo,omitempty" type:"Struct"`
	// The request ID, used for Tracing Analysis and troubleshooting.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-EF0123456789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the upload policy is generated. A successful response always returns `true`. An error response is returned upon failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the newly created Transit record, used for subsequent queries and storage operations.
	//
	// example:
	//
	// transit_0123456789abcdef0123456789abcdef
	TransitId *string `json:"TransitId,omitempty" xml:"TransitId,omitempty"`
}

func (s CreateTransitUploadPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitUploadPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransitUploadPolicyResponseBody) GetFilePath() *string {
	return s.FilePath
}

func (s *CreateTransitUploadPolicyResponseBody) GetPolicyInfo() *CreateTransitUploadPolicyResponseBodyPolicyInfo {
	return s.PolicyInfo
}

func (s *CreateTransitUploadPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTransitUploadPolicyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTransitUploadPolicyResponseBody) GetTransitId() *string {
	return s.TransitId
}

func (s *CreateTransitUploadPolicyResponseBody) SetFilePath(v string) *CreateTransitUploadPolicyResponseBody {
	s.FilePath = &v
	return s
}

func (s *CreateTransitUploadPolicyResponseBody) SetPolicyInfo(v *CreateTransitUploadPolicyResponseBodyPolicyInfo) *CreateTransitUploadPolicyResponseBody {
	s.PolicyInfo = v
	return s
}

func (s *CreateTransitUploadPolicyResponseBody) SetRequestId(v string) *CreateTransitUploadPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransitUploadPolicyResponseBody) SetSuccess(v bool) *CreateTransitUploadPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTransitUploadPolicyResponseBody) SetTransitId(v string) *CreateTransitUploadPolicyResponseBody {
	s.TransitId = &v
	return s
}

func (s *CreateTransitUploadPolicyResponseBody) Validate() error {
	if s.PolicyInfo != nil {
		if err := s.PolicyInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTransitUploadPolicyResponseBodyPolicyInfo struct {
	// The `OSSAccessKeyId` field in the PostObject form. Protect this value together with the entire `PolicyInfo`.
	//
	// example:
	//
	// <REDACTED>
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// The object storage key. The value is the same as the top-level `FilePath`.
	//
	// example:
	//
	// skill-bundle/tenant-demo/user-demo/20260904120000_0123456789abcdef0123456789abcdef.zip
	Dir *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	// The target URL to which the client sends the PostObject request.
	//
	// example:
	//
	// https://upload.example.invalid
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The Base64-encoded PostObject upload policy. Protect this value together with the entire `PolicyInfo`.
	//
	// example:
	//
	// <REDACTED>
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The `x-oss-security-token` field in the PostObject form when STS credentials are used. This field may be empty when STS is not used. This field contains sensitive authorization information.
	//
	// example:
	//
	// <REDACTED>
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The signature field in the PostObject form. This field contains sensitive authorization information.
	//
	// example:
	//
	// <REDACTED>
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s CreateTransitUploadPolicyResponseBodyPolicyInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitUploadPolicyResponseBodyPolicyInfo) GoString() string {
	return s.String()
}

func (s *CreateTransitUploadPolicyResponseBodyPolicyInfo) GetAccessId() *string {
	return s.AccessId
}

func (s *CreateTransitUploadPolicyResponseBodyPolicyInfo) GetDir() *string {
	return s.Dir
}

func (s *CreateTransitUploadPolicyResponseBodyPolicyInfo) GetHost() *string {
	return s.Host
}

func (s *CreateTransitUploadPolicyResponseBodyPolicyInfo) GetPolicy() *string {
	return s.Policy
}

func (s *CreateTransitUploadPolicyResponseBodyPolicyInfo) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateTransitUploadPolicyResponseBodyPolicyInfo) GetSignature() *string {
	return s.Signature
}

func (s *CreateTransitUploadPolicyResponseBodyPolicyInfo) SetAccessId(v string) *CreateTransitUploadPolicyResponseBodyPolicyInfo {
	s.AccessId = &v
	return s
}

func (s *CreateTransitUploadPolicyResponseBodyPolicyInfo) SetDir(v string) *CreateTransitUploadPolicyResponseBodyPolicyInfo {
	s.Dir = &v
	return s
}

func (s *CreateTransitUploadPolicyResponseBodyPolicyInfo) SetHost(v string) *CreateTransitUploadPolicyResponseBodyPolicyInfo {
	s.Host = &v
	return s
}

func (s *CreateTransitUploadPolicyResponseBodyPolicyInfo) SetPolicy(v string) *CreateTransitUploadPolicyResponseBodyPolicyInfo {
	s.Policy = &v
	return s
}

func (s *CreateTransitUploadPolicyResponseBodyPolicyInfo) SetSecurityToken(v string) *CreateTransitUploadPolicyResponseBodyPolicyInfo {
	s.SecurityToken = &v
	return s
}

func (s *CreateTransitUploadPolicyResponseBodyPolicyInfo) SetSignature(v string) *CreateTransitUploadPolicyResponseBodyPolicyInfo {
	s.Signature = &v
	return s
}

func (s *CreateTransitUploadPolicyResponseBodyPolicyInfo) Validate() error {
	return dara.Validate(s)
}
