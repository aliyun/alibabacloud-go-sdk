// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttachmentUploadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetAttachmentUploadUrlResponseBody
	GetCode() *int32
	SetData(v *GetAttachmentUploadUrlResponseBodyData) *GetAttachmentUploadUrlResponseBody
	GetData() *GetAttachmentUploadUrlResponseBodyData
	SetMessage(v string) *GetAttachmentUploadUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAttachmentUploadUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAttachmentUploadUrlResponseBody
	GetSuccess() *bool
}

type GetAttachmentUploadUrlResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned after the call succeeds.
	Data *GetAttachmentUploadUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. If success is set to false, the message is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ED195C2C-787F-511C-8204-714456781861
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates that the call is normal.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAttachmentUploadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAttachmentUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAttachmentUploadUrlResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetAttachmentUploadUrlResponseBody) GetData() *GetAttachmentUploadUrlResponseBodyData {
	return s.Data
}

func (s *GetAttachmentUploadUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAttachmentUploadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAttachmentUploadUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAttachmentUploadUrlResponseBody) SetCode(v int32) *GetAttachmentUploadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetAttachmentUploadUrlResponseBody) SetData(v *GetAttachmentUploadUrlResponseBodyData) *GetAttachmentUploadUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetAttachmentUploadUrlResponseBody) SetMessage(v string) *GetAttachmentUploadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetAttachmentUploadUrlResponseBody) SetRequestId(v string) *GetAttachmentUploadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAttachmentUploadUrlResponseBody) SetSuccess(v bool) *GetAttachmentUploadUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetAttachmentUploadUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAttachmentUploadUrlResponseBodyData struct {
	// Query the signed URL of an OSS object
	//
	// example:
	//
	// https://gts-workorder-pre.oss-cn-beijing.aliyuncs.com
	//
	// /20220314/cabb80c3-430b-4079-a9f2-d2a0d1c2a587.png?Expires=1647328689&OSSAccessKeyId=LTAI****************&Signature=AbSEh26G3oYrJ8ivr4B0xzF89zk%3D
	GetSignedUrl *string `json:"GetSignedUrl,omitempty" xml:"GetSignedUrl,omitempty"`
	// Uploaded file identifier
	//
	// example:
	//
	// cdb5d174-c282-4b2d-9048-e74ea2223127.jpg
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	// The signed URL used to upload the object to OSS.
	//
	// example:
	//
	// https://gts-workorder-pre.oss-cn-beijing.aliyuncs.com
	//
	// /20220314/cabb80c3-430b-4079-a9f2-d2a0d1c2a587.png?Expires=1647328689&OSSAccessKeyId=LTAI****************&Signature=AbSEh26G3oYrJ8ivr4B0xzF89zk%3D
	PutSignedUrl *string `json:"PutSignedUrl,omitempty" xml:"PutSignedUrl,omitempty"`
}

func (s GetAttachmentUploadUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAttachmentUploadUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAttachmentUploadUrlResponseBodyData) GetGetSignedUrl() *string {
	return s.GetSignedUrl
}

func (s *GetAttachmentUploadUrlResponseBodyData) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *GetAttachmentUploadUrlResponseBodyData) GetPutSignedUrl() *string {
	return s.PutSignedUrl
}

func (s *GetAttachmentUploadUrlResponseBodyData) SetGetSignedUrl(v string) *GetAttachmentUploadUrlResponseBodyData {
	s.GetSignedUrl = &v
	return s
}

func (s *GetAttachmentUploadUrlResponseBodyData) SetObjectKey(v string) *GetAttachmentUploadUrlResponseBodyData {
	s.ObjectKey = &v
	return s
}

func (s *GetAttachmentUploadUrlResponseBodyData) SetPutSignedUrl(v string) *GetAttachmentUploadUrlResponseBodyData {
	s.PutSignedUrl = &v
	return s
}

func (s *GetAttachmentUploadUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
