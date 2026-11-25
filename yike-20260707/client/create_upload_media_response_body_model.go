// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadMediaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileURL(v string) *CreateUploadMediaResponseBody
	GetFileURL() *string
	SetMediaId(v string) *CreateUploadMediaResponseBody
	GetMediaId() *string
	SetObjectKey(v string) *CreateUploadMediaResponseBody
	GetObjectKey() *string
	SetRequestId(v string) *CreateUploadMediaResponseBody
	GetRequestId() *string
	SetStorageLocation(v string) *CreateUploadMediaResponseBody
	GetStorageLocation() *string
	SetTicketExpiration(v string) *CreateUploadMediaResponseBody
	GetTicketExpiration() *string
	SetUploadCredentials(v *CreateUploadMediaResponseBodyUploadCredentials) *CreateUploadMediaResponseBody
	GetUploadCredentials() *CreateUploadMediaResponseBodyUploadCredentials
	SetUploadTicket(v string) *CreateUploadMediaResponseBody
	GetUploadTicket() *string
}

type CreateUploadMediaResponseBody struct {
	// The unsigned OSS URL of the file in the `oss://bucket/objectKey` format.
	//
	// example:
	//
	// oss://yike-shared-ap-southeast-1-1/yike-v1/gi3dkn-gi3dkn/2026-11-18/08/video/b7e94fc0e1bf71f1bfa2f6e7c5486601aB3xYz01.mp4
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The media asset ID.
	//
	// example:
	//
	// b7e94fc0e*****6e7c5486601
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The OSS object key in the `yike-v1/uid segment-ws segment/date/hour/type/mediaId+random string+source suffix.file name extension` format.
	//
	// example:
	//
	// yike-v1/gi3dkn-gi3dkn/2026-11-18/08/video/b7e94fc0e1bf71f1bfa2f6e7c5486601aB3xYz01.mp4
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9B4C2E10-3F5A-4D6B-8C7D-1E2F3A4B5C6D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The storage location identifier in the `bucket.endpoint` format.
	//
	// example:
	//
	// yike-shared-ap-southeast-1-1.oss-ap-southeast-1.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The expiration time of the ticket in UTC and ISO 8601 format. This field is returned when `UPLOAD_TICKET` is used.
	//
	// example:
	//
	// 2026-11-18T12:00:00Z
	TicketExpiration *string `json:"TicketExpiration,omitempty" xml:"TicketExpiration,omitempty"`
	// The upload credentials. Some fields may be empty depending on the `CredentialType`.
	UploadCredentials *CreateUploadMediaResponseBodyUploadCredentials `json:"UploadCredentials,omitempty" xml:"UploadCredentials,omitempty" type:"Struct"`
	// The upload ticket. This field is returned when `UPLOAD_TICKET` is used. It is a 27-character string in the UT1+region code+random string+check digit format. A real credential is issued based on this ticket during a ticket exchange.
	//
	// example:
	//
	// UT1SGK7M2QX9R4T6V8W0Y3Z5ABC
	UploadTicket *string `json:"UploadTicket,omitempty" xml:"UploadTicket,omitempty"`
}

func (s CreateUploadMediaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadMediaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUploadMediaResponseBody) GetFileURL() *string {
	return s.FileURL
}

func (s *CreateUploadMediaResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *CreateUploadMediaResponseBody) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *CreateUploadMediaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUploadMediaResponseBody) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *CreateUploadMediaResponseBody) GetTicketExpiration() *string {
	return s.TicketExpiration
}

func (s *CreateUploadMediaResponseBody) GetUploadCredentials() *CreateUploadMediaResponseBodyUploadCredentials {
	return s.UploadCredentials
}

func (s *CreateUploadMediaResponseBody) GetUploadTicket() *string {
	return s.UploadTicket
}

func (s *CreateUploadMediaResponseBody) SetFileURL(v string) *CreateUploadMediaResponseBody {
	s.FileURL = &v
	return s
}

func (s *CreateUploadMediaResponseBody) SetMediaId(v string) *CreateUploadMediaResponseBody {
	s.MediaId = &v
	return s
}

func (s *CreateUploadMediaResponseBody) SetObjectKey(v string) *CreateUploadMediaResponseBody {
	s.ObjectKey = &v
	return s
}

func (s *CreateUploadMediaResponseBody) SetRequestId(v string) *CreateUploadMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUploadMediaResponseBody) SetStorageLocation(v string) *CreateUploadMediaResponseBody {
	s.StorageLocation = &v
	return s
}

func (s *CreateUploadMediaResponseBody) SetTicketExpiration(v string) *CreateUploadMediaResponseBody {
	s.TicketExpiration = &v
	return s
}

func (s *CreateUploadMediaResponseBody) SetUploadCredentials(v *CreateUploadMediaResponseBodyUploadCredentials) *CreateUploadMediaResponseBody {
	s.UploadCredentials = v
	return s
}

func (s *CreateUploadMediaResponseBody) SetUploadTicket(v string) *CreateUploadMediaResponseBody {
	s.UploadTicket = &v
	return s
}

func (s *CreateUploadMediaResponseBody) Validate() error {
	if s.UploadCredentials != nil {
		if err := s.UploadCredentials.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateUploadMediaResponseBodyUploadCredentials struct {
	// The temporary STS AccessKey ID. This field is returned when `STS_TOKEN` is used.
	//
	// example:
	//
	// STS.NUgYrLnoC63mLqWeRtYuIoPaSdF
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The temporary STS AccessKey secret. This field is returned when `STS_TOKEN` is used.
	//
	// example:
	//
	// CVwjCkNzTMupZ8nRtYeXwQaBcDfGhIjK
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// The OSS bucket. This field is returned when `STS_TOKEN`, `POST_POLICY`, or `PRESIGNED_URL` is used.
	//
	// example:
	//
	// yike-shared-ap-southeast-1-1
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The expiration time of the credential in UTC and ISO 8601 format.
	//
	// example:
	//
	// 2026-11-18T12:00:00Z
	Expiration *string `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The Base64-encoded PostObject policy. This field is returned when `POST_POLICY` is used.
	//
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyNi0xMS0xOFQxMjowMDowMFoiLCJjb25kaXRpb25zIjpbXX0=
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The AccessKey ID used for the signature. This field is returned when `POST_POLICY` is used.
	//
	// example:
	//
	// LTAI5tG8XkQmP2vR9wYzAbCd
	PostAccessKeyId *string `json:"PostAccessKeyId,omitempty" xml:"PostAccessKeyId,omitempty"`
	// The destination endpoint for the POST request. This field is returned when `POST_POLICY` is used.
	//
	// example:
	//
	// https://yike-shared-ap-southeast-1-1.oss-ap-southeast-1.aliyuncs.com
	PostEndpoint *string `json:"PostEndpoint,omitempty" xml:"PostEndpoint,omitempty"`
	// The OSS region. This field is returned in all three credential modes.
	//
	// example:
	//
	// ap-southeast-1
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The STS security token. This field is returned when `STS_TOKEN` is used.
	//
	// example:
	//
	// CAISjgJ1q6Ft5B2yfSjIr5DDEOTM2pZ6g7oIExampleToken
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The PostObject signature. This field is returned when `POST_POLICY` is used.
	//
	// example:
	//
	// vXkQ9dRt2mLpN8sYwZaBcDfEgHj=
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// The pre-signed PUT URL. This field is returned when `PRESIGNED_URL` is used.
	//
	// example:
	//
	// https://yike-shared-ap-southeast-1-1.oss-ap-southeast-1.aliyuncs.com/yike-v1/.../demo.mp4?Expires=1763467200&OSSAccessKeyId=STS.xxx&Signature=xxx
	SignedUrl *string `json:"SignedUrl,omitempty" xml:"SignedUrl,omitempty"`
}

func (s CreateUploadMediaResponseBodyUploadCredentials) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadMediaResponseBodyUploadCredentials) GoString() string {
	return s.String()
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) GetBucket() *string {
	return s.Bucket
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) GetExpiration() *string {
	return s.Expiration
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) GetPolicy() *string {
	return s.Policy
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) GetPostAccessKeyId() *string {
	return s.PostAccessKeyId
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) GetPostEndpoint() *string {
	return s.PostEndpoint
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) GetRegion() *string {
	return s.Region
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) GetSignature() *string {
	return s.Signature
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) GetSignedUrl() *string {
	return s.SignedUrl
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) SetAccessKeyId(v string) *CreateUploadMediaResponseBodyUploadCredentials {
	s.AccessKeyId = &v
	return s
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) SetAccessKeySecret(v string) *CreateUploadMediaResponseBodyUploadCredentials {
	s.AccessKeySecret = &v
	return s
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) SetBucket(v string) *CreateUploadMediaResponseBodyUploadCredentials {
	s.Bucket = &v
	return s
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) SetExpiration(v string) *CreateUploadMediaResponseBodyUploadCredentials {
	s.Expiration = &v
	return s
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) SetPolicy(v string) *CreateUploadMediaResponseBodyUploadCredentials {
	s.Policy = &v
	return s
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) SetPostAccessKeyId(v string) *CreateUploadMediaResponseBodyUploadCredentials {
	s.PostAccessKeyId = &v
	return s
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) SetPostEndpoint(v string) *CreateUploadMediaResponseBodyUploadCredentials {
	s.PostEndpoint = &v
	return s
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) SetRegion(v string) *CreateUploadMediaResponseBodyUploadCredentials {
	s.Region = &v
	return s
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) SetSecurityToken(v string) *CreateUploadMediaResponseBodyUploadCredentials {
	s.SecurityToken = &v
	return s
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) SetSignature(v string) *CreateUploadMediaResponseBodyUploadCredentials {
	s.Signature = &v
	return s
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) SetSignedUrl(v string) *CreateUploadMediaResponseBodyUploadCredentials {
	s.SignedUrl = &v
	return s
}

func (s *CreateUploadMediaResponseBodyUploadCredentials) Validate() error {
	return dara.Validate(s)
}
