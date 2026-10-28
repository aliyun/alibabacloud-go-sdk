// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileDetectUploadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateFileDetectUploadUrlResponseBody
	GetRequestId() *string
	SetUploadUrlList(v []*CreateFileDetectUploadUrlResponseBodyUploadUrlList) *CreateFileDetectUploadUrlResponseBody
	GetUploadUrlList() []*CreateFileDetectUploadUrlResponseBodyUploadUrlList
}

type CreateFileDetectUploadUrlResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 09969D2C-4FAD-429E-BFBF-9A60DEF8BF6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array consisting of the parameters that are required to upload a file.
	UploadUrlList []*CreateFileDetectUploadUrlResponseBodyUploadUrlList `json:"UploadUrlList,omitempty" xml:"UploadUrlList,omitempty" type:"Repeated"`
}

func (s CreateFileDetectUploadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFileDetectUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileDetectUploadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFileDetectUploadUrlResponseBody) GetUploadUrlList() []*CreateFileDetectUploadUrlResponseBodyUploadUrlList {
	return s.UploadUrlList
}

func (s *CreateFileDetectUploadUrlResponseBody) SetRequestId(v string) *CreateFileDetectUploadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFileDetectUploadUrlResponseBody) SetUploadUrlList(v []*CreateFileDetectUploadUrlResponseBodyUploadUrlList) *CreateFileDetectUploadUrlResponseBody {
	s.UploadUrlList = v
	return s
}

func (s *CreateFileDetectUploadUrlResponseBody) Validate() error {
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

type CreateFileDetectUploadUrlResponseBodyUploadUrlList struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The signature information.
	Context *CreateFileDetectUploadUrlResponseBodyUploadUrlListContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
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
	// The identifier of the file.
	//
	// example:
	//
	// 0a212417e65c26ff133cfff28f6c****
	HashKey *string `json:"HashKey,omitempty" xml:"HashKey,omitempty"`
	// The internal endpoint of the URL to which the file is uploaded.
	//
	// example:
	//
	// http://example.com
	InternalUrl *string `json:"InternalUrl,omitempty" xml:"InternalUrl,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The public endpoint of the URL to which the file is uploaded.
	//
	// example:
	//
	// http://example.com
	PublicUrl *string `json:"PublicUrl,omitempty" xml:"PublicUrl,omitempty"`
}

func (s CreateFileDetectUploadUrlResponseBodyUploadUrlList) String() string {
	return dara.Prettify(s)
}

func (s CreateFileDetectUploadUrlResponseBodyUploadUrlList) GoString() string {
	return s.String()
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlList) GetCode() *string {
	return s.Code
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlList) GetContext() *CreateFileDetectUploadUrlResponseBodyUploadUrlListContext {
	return s.Context
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlList) GetExpire() *string {
	return s.Expire
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlList) GetFileExist() *bool {
	return s.FileExist
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlList) GetHashKey() *string {
	return s.HashKey
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlList) GetInternalUrl() *string {
	return s.InternalUrl
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlList) GetMessage() *string {
	return s.Message
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlList) GetPublicUrl() *string {
	return s.PublicUrl
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlList) SetCode(v string) *CreateFileDetectUploadUrlResponseBodyUploadUrlList {
	s.Code = &v
	return s
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlList) SetContext(v *CreateFileDetectUploadUrlResponseBodyUploadUrlListContext) *CreateFileDetectUploadUrlResponseBodyUploadUrlList {
	s.Context = v
	return s
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlList) SetExpire(v string) *CreateFileDetectUploadUrlResponseBodyUploadUrlList {
	s.Expire = &v
	return s
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlList) SetFileExist(v bool) *CreateFileDetectUploadUrlResponseBodyUploadUrlList {
	s.FileExist = &v
	return s
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlList) SetHashKey(v string) *CreateFileDetectUploadUrlResponseBodyUploadUrlList {
	s.HashKey = &v
	return s
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlList) SetInternalUrl(v string) *CreateFileDetectUploadUrlResponseBodyUploadUrlList {
	s.InternalUrl = &v
	return s
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlList) SetMessage(v string) *CreateFileDetectUploadUrlResponseBodyUploadUrlList {
	s.Message = &v
	return s
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlList) SetPublicUrl(v string) *CreateFileDetectUploadUrlResponseBodyUploadUrlList {
	s.PublicUrl = &v
	return s
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlList) Validate() error {
	if s.Context != nil {
		if err := s.Context.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateFileDetectUploadUrlResponseBodyUploadUrlListContext struct {
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
	// eyJleHBpcmF0aW9uIjoiMjAyMi0wNy0yM1QxMDo1ODoxMC43NTNaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsMjA5NzE1MjBdLFsiZXEiLCIka2V5IiwiMS8yMDIyLzA2LzIzLzE4LzU4LzE2NTU5ODE4OTA3NTM4NTc2MjFkNS1kN2E1LTQ5YzAtOGJjZi0yMTMyY2JiYTdmYzMi****
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The signature that is used to upload the file.
	//
	// example:
	//
	// wDhPgVdnY/bkKFYcYFl+4crl****
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s CreateFileDetectUploadUrlResponseBodyUploadUrlListContext) String() string {
	return dara.Prettify(s)
}

func (s CreateFileDetectUploadUrlResponseBodyUploadUrlListContext) GoString() string {
	return s.String()
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlListContext) GetAccessId() *string {
	return s.AccessId
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlListContext) GetOssKey() *string {
	return s.OssKey
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlListContext) GetPolicy() *string {
	return s.Policy
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlListContext) GetSignature() *string {
	return s.Signature
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlListContext) SetAccessId(v string) *CreateFileDetectUploadUrlResponseBodyUploadUrlListContext {
	s.AccessId = &v
	return s
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlListContext) SetOssKey(v string) *CreateFileDetectUploadUrlResponseBodyUploadUrlListContext {
	s.OssKey = &v
	return s
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlListContext) SetPolicy(v string) *CreateFileDetectUploadUrlResponseBodyUploadUrlListContext {
	s.Policy = &v
	return s
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlListContext) SetSignature(v string) *CreateFileDetectUploadUrlResponseBodyUploadUrlListContext {
	s.Signature = &v
	return s
}

func (s *CreateFileDetectUploadUrlResponseBodyUploadUrlListContext) Validate() error {
	return dara.Validate(s)
}
