// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualificationUploadPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessid(v string) *GetQualificationUploadPolicyResponseBody
	GetAccessid() *string
	SetDir(v string) *GetQualificationUploadPolicyResponseBody
	GetDir() *string
	SetExpire(v string) *GetQualificationUploadPolicyResponseBody
	GetExpire() *string
	SetHost(v string) *GetQualificationUploadPolicyResponseBody
	GetHost() *string
	SetPolicy(v string) *GetQualificationUploadPolicyResponseBody
	GetPolicy() *string
	SetPrefix(v string) *GetQualificationUploadPolicyResponseBody
	GetPrefix() *string
	SetRequestId(v string) *GetQualificationUploadPolicyResponseBody
	GetRequestId() *string
	SetSignature(v string) *GetQualificationUploadPolicyResponseBody
	GetSignature() *string
}

type GetQualificationUploadPolicyResponseBody struct {
	// example:
	//
	// hObpgEXoca42****
	Accessid *string `json:"Accessid,omitempty" xml:"Accessid,omitempty"`
	// example:
	//
	// 20211220/131953297274****_4de3db85-4f98-488d-845b-d75bf035b13d
	Dir *string `json:"Dir,omitempty" xml:"Dir,omitempty"`
	// example:
	//
	// 1593688811881
	Expire *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// example:
	//
	// https://********-review.oss-cn-********.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAaMC0wNy0wMlQxKToyMDoxMS44ODRaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsNTI0Mjg4MDBdLFsic3RhcnRzLXdpdGgiLCIka2V5IiwiMTIxOTU0MTE2MTIxMzA1Ny9PRkZMSU5FX1RSQU5TRkVSLzE1OTM2ODg1MTE4ODMi****
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// 20211220/131953297274****_4de3db85-4f98-488d-845b-d75bf035b13d_${filename}
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// example:
	//
	// 9DFCF6F8-243C-****-8035-4B12FEFD7D48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// pNVECGkyL0tl4bKXekV5ErZ****
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GetQualificationUploadPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualificationUploadPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualificationUploadPolicyResponseBody) GetAccessid() *string {
	return s.Accessid
}

func (s *GetQualificationUploadPolicyResponseBody) GetDir() *string {
	return s.Dir
}

func (s *GetQualificationUploadPolicyResponseBody) GetExpire() *string {
	return s.Expire
}

func (s *GetQualificationUploadPolicyResponseBody) GetHost() *string {
	return s.Host
}

func (s *GetQualificationUploadPolicyResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *GetQualificationUploadPolicyResponseBody) GetPrefix() *string {
	return s.Prefix
}

func (s *GetQualificationUploadPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualificationUploadPolicyResponseBody) GetSignature() *string {
	return s.Signature
}

func (s *GetQualificationUploadPolicyResponseBody) SetAccessid(v string) *GetQualificationUploadPolicyResponseBody {
	s.Accessid = &v
	return s
}

func (s *GetQualificationUploadPolicyResponseBody) SetDir(v string) *GetQualificationUploadPolicyResponseBody {
	s.Dir = &v
	return s
}

func (s *GetQualificationUploadPolicyResponseBody) SetExpire(v string) *GetQualificationUploadPolicyResponseBody {
	s.Expire = &v
	return s
}

func (s *GetQualificationUploadPolicyResponseBody) SetHost(v string) *GetQualificationUploadPolicyResponseBody {
	s.Host = &v
	return s
}

func (s *GetQualificationUploadPolicyResponseBody) SetPolicy(v string) *GetQualificationUploadPolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *GetQualificationUploadPolicyResponseBody) SetPrefix(v string) *GetQualificationUploadPolicyResponseBody {
	s.Prefix = &v
	return s
}

func (s *GetQualificationUploadPolicyResponseBody) SetRequestId(v string) *GetQualificationUploadPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualificationUploadPolicyResponseBody) SetSignature(v string) *GetQualificationUploadPolicyResponseBody {
	s.Signature = &v
	return s
}

func (s *GetQualificationUploadPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
