// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchUploadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateBatchUploadUrlResponseBody
	GetRequestId() *string
	SetUploadUrlList(v []*CreateBatchUploadUrlResponseBodyUploadUrlList) *CreateBatchUploadUrlResponseBody
	GetUploadUrlList() []*CreateBatchUploadUrlResponseBodyUploadUrlList
}

type CreateBatchUploadUrlResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// DA62490B-7883-5EB4-8601-F2D1D9******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array consisting of the parameters that are required to upload a file.
	UploadUrlList []*CreateBatchUploadUrlResponseBodyUploadUrlList `json:"UploadUrlList,omitempty" xml:"UploadUrlList,omitempty" type:"Repeated"`
}

func (s CreateBatchUploadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBatchUploadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBatchUploadUrlResponseBody) GetUploadUrlList() []*CreateBatchUploadUrlResponseBodyUploadUrlList {
	return s.UploadUrlList
}

func (s *CreateBatchUploadUrlResponseBody) SetRequestId(v string) *CreateBatchUploadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBatchUploadUrlResponseBody) SetUploadUrlList(v []*CreateBatchUploadUrlResponseBodyUploadUrlList) *CreateBatchUploadUrlResponseBody {
	s.UploadUrlList = v
	return s
}

func (s *CreateBatchUploadUrlResponseBody) Validate() error {
	if s.UploadUrlList != nil {
		for _, item := range s.UploadUrlList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateBatchUploadUrlResponseBodyUploadUrlList struct {
	// The signature information.
	Context *CreateBatchUploadUrlResponseBodyUploadUrlListContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
	// The timestamp when the values of the parameters expire. Unit: milliseconds.
	//
	// example:
	//
	// 1658562101370
	Expire *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// Indicates whether the file exists in the cloud. Valid values:
	//
	// 	- **true**: The file exists in the cloud. You do not need to upload the file.
	//
	// 	- **false**: The file does not exist in the cloud. You must upload the file.
	//
	// example:
	//
	// false
	FileExist *bool `json:"FileExist,omitempty" xml:"FileExist,omitempty"`
	// The internal endpoint of the URL to which the file is uploaded.
	//
	// example:
	//
	// http://example.com
	InternalUrl *string `json:"InternalUrl,omitempty" xml:"InternalUrl,omitempty"`
	// The identifier of the file.
	//
	// example:
	//
	// 2f8dc248a0fbb96c69e45acad2******
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The public endpoint of the URL to which the file is uploaded.
	//
	// example:
	//
	// http://example.com
	PublicUrl *string `json:"PublicUrl,omitempty" xml:"PublicUrl,omitempty"`
}

func (s CreateBatchUploadUrlResponseBodyUploadUrlList) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchUploadUrlResponseBodyUploadUrlList) GoString() string {
	return s.String()
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlList) GetContext() *CreateBatchUploadUrlResponseBodyUploadUrlListContext {
	return s.Context
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlList) GetExpire() *string {
	return s.Expire
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlList) GetFileExist() *bool {
	return s.FileExist
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlList) GetInternalUrl() *string {
	return s.InternalUrl
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlList) GetMd5() *string {
	return s.Md5
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlList) GetPublicUrl() *string {
	return s.PublicUrl
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlList) SetContext(v *CreateBatchUploadUrlResponseBodyUploadUrlListContext) *CreateBatchUploadUrlResponseBodyUploadUrlList {
	s.Context = v
	return s
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlList) SetExpire(v string) *CreateBatchUploadUrlResponseBodyUploadUrlList {
	s.Expire = &v
	return s
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlList) SetFileExist(v bool) *CreateBatchUploadUrlResponseBodyUploadUrlList {
	s.FileExist = &v
	return s
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlList) SetInternalUrl(v string) *CreateBatchUploadUrlResponseBodyUploadUrlList {
	s.InternalUrl = &v
	return s
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlList) SetMd5(v string) *CreateBatchUploadUrlResponseBodyUploadUrlList {
	s.Md5 = &v
	return s
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlList) SetPublicUrl(v string) *CreateBatchUploadUrlResponseBodyUploadUrlList {
	s.PublicUrl = &v
	return s
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlList) Validate() error {
	if s.Context != nil {
		if err := s.Context.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBatchUploadUrlResponseBodyUploadUrlListContext struct {
	// The AccessKey ID that is used to access the OSS bucket.
	//
	// example:
	//
	// yourAccessKeyID
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// The key of the file that is used after the file is uploaded to the OSS bucket.
	//
	// example:
	//
	// 1/2022/06/23/15/41/16559701077444693a0c6-33b2-4cc2-a99f-9f38b8b8****
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// The policy that poses limits on file upload. For example, the policy can limit the size of the file.
	//
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyMi0wNy0yM1QxMDo1ODoxMC****
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The signature that is used to upload the file.
	//
	// example:
	//
	// wDhPgVdnY/bkKFYcYFl+4crl****
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s CreateBatchUploadUrlResponseBodyUploadUrlListContext) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchUploadUrlResponseBodyUploadUrlListContext) GoString() string {
	return s.String()
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlListContext) GetAccessId() *string {
	return s.AccessId
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlListContext) GetOssKey() *string {
	return s.OssKey
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlListContext) GetPolicy() *string {
	return s.Policy
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlListContext) GetSignature() *string {
	return s.Signature
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlListContext) SetAccessId(v string) *CreateBatchUploadUrlResponseBodyUploadUrlListContext {
	s.AccessId = &v
	return s
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlListContext) SetOssKey(v string) *CreateBatchUploadUrlResponseBodyUploadUrlListContext {
	s.OssKey = &v
	return s
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlListContext) SetPolicy(v string) *CreateBatchUploadUrlResponseBodyUploadUrlListContext {
	s.Policy = &v
	return s
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlListContext) SetSignature(v string) *CreateBatchUploadUrlResponseBodyUploadUrlListContext {
	s.Signature = &v
	return s
}

func (s *CreateBatchUploadUrlResponseBodyUploadUrlListContext) Validate() error {
	return dara.Validate(s)
}
