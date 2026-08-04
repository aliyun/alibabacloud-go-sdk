// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadIdentityRegistrationDocConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUploadIdentityRegistrationDocConfigResponseBody
	GetCode() *string
	SetData(v *GetUploadIdentityRegistrationDocConfigResponseBodyData) *GetUploadIdentityRegistrationDocConfigResponseBody
	GetData() *GetUploadIdentityRegistrationDocConfigResponseBodyData
	SetMessage(v string) *GetUploadIdentityRegistrationDocConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUploadIdentityRegistrationDocConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUploadIdentityRegistrationDocConfigResponseBody
	GetSuccess() *bool
}

type GetUploadIdentityRegistrationDocConfigResponseBody struct {
	Code      *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetUploadIdentityRegistrationDocConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUploadIdentityRegistrationDocConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUploadIdentityRegistrationDocConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadIdentityRegistrationDocConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUploadIdentityRegistrationDocConfigResponseBody) GetData() *GetUploadIdentityRegistrationDocConfigResponseBodyData {
	return s.Data
}

func (s *GetUploadIdentityRegistrationDocConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUploadIdentityRegistrationDocConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUploadIdentityRegistrationDocConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUploadIdentityRegistrationDocConfigResponseBody) SetCode(v string) *GetUploadIdentityRegistrationDocConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetUploadIdentityRegistrationDocConfigResponseBody) SetData(v *GetUploadIdentityRegistrationDocConfigResponseBodyData) *GetUploadIdentityRegistrationDocConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetUploadIdentityRegistrationDocConfigResponseBody) SetMessage(v string) *GetUploadIdentityRegistrationDocConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetUploadIdentityRegistrationDocConfigResponseBody) SetRequestId(v string) *GetUploadIdentityRegistrationDocConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUploadIdentityRegistrationDocConfigResponseBody) SetSuccess(v bool) *GetUploadIdentityRegistrationDocConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetUploadIdentityRegistrationDocConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUploadIdentityRegistrationDocConfigResponseBodyData struct {
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	FileName   *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
}

func (s GetUploadIdentityRegistrationDocConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUploadIdentityRegistrationDocConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUploadIdentityRegistrationDocConfigResponseBodyData) GetBucketName() *string {
	return s.BucketName
}

func (s *GetUploadIdentityRegistrationDocConfigResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *GetUploadIdentityRegistrationDocConfigResponseBodyData) SetBucketName(v string) *GetUploadIdentityRegistrationDocConfigResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *GetUploadIdentityRegistrationDocConfigResponseBodyData) SetFileName(v string) *GetUploadIdentityRegistrationDocConfigResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetUploadIdentityRegistrationDocConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
