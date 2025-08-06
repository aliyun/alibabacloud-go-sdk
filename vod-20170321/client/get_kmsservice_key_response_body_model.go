// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKMSServiceKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKmsKeyInfoList(v []*GetKMSServiceKeyResponseBodyKmsKeyInfoList) *GetKMSServiceKeyResponseBody
	GetKmsKeyInfoList() []*GetKMSServiceKeyResponseBodyKmsKeyInfoList
	SetRequestId(v string) *GetKMSServiceKeyResponseBody
	GetRequestId() *string
}

type GetKMSServiceKeyResponseBody struct {
	KmsKeyInfoList []*GetKMSServiceKeyResponseBodyKmsKeyInfoList `json:"KmsKeyInfoList,omitempty" xml:"KmsKeyInfoList,omitempty" type:"Repeated"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetKMSServiceKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKMSServiceKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetKMSServiceKeyResponseBody) GetKmsKeyInfoList() []*GetKMSServiceKeyResponseBodyKmsKeyInfoList {
	return s.KmsKeyInfoList
}

func (s *GetKMSServiceKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKMSServiceKeyResponseBody) SetKmsKeyInfoList(v []*GetKMSServiceKeyResponseBodyKmsKeyInfoList) *GetKMSServiceKeyResponseBody {
	s.KmsKeyInfoList = v
	return s
}

func (s *GetKMSServiceKeyResponseBody) SetRequestId(v string) *GetKMSServiceKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKMSServiceKeyResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetKMSServiceKeyResponseBodyKmsKeyInfoList struct {
	Arn          *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	Creator      *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DeleteDate   *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	KeyState     *string `json:"KeyState,omitempty" xml:"KeyState,omitempty"`
	KeyUsage     *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty"`
}

func (s GetKMSServiceKeyResponseBodyKmsKeyInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetKMSServiceKeyResponseBodyKmsKeyInfoList) GoString() string {
	return s.String()
}

func (s *GetKMSServiceKeyResponseBodyKmsKeyInfoList) GetArn() *string {
	return s.Arn
}

func (s *GetKMSServiceKeyResponseBodyKmsKeyInfoList) GetCreationDate() *string {
	return s.CreationDate
}

func (s *GetKMSServiceKeyResponseBodyKmsKeyInfoList) GetCreator() *string {
	return s.Creator
}

func (s *GetKMSServiceKeyResponseBodyKmsKeyInfoList) GetDeleteDate() *string {
	return s.DeleteDate
}

func (s *GetKMSServiceKeyResponseBodyKmsKeyInfoList) GetDescription() *string {
	return s.Description
}

func (s *GetKMSServiceKeyResponseBodyKmsKeyInfoList) GetKeyId() *string {
	return s.KeyId
}

func (s *GetKMSServiceKeyResponseBodyKmsKeyInfoList) GetKeyState() *string {
	return s.KeyState
}

func (s *GetKMSServiceKeyResponseBodyKmsKeyInfoList) GetKeyUsage() *string {
	return s.KeyUsage
}

func (s *GetKMSServiceKeyResponseBodyKmsKeyInfoList) SetArn(v string) *GetKMSServiceKeyResponseBodyKmsKeyInfoList {
	s.Arn = &v
	return s
}

func (s *GetKMSServiceKeyResponseBodyKmsKeyInfoList) SetCreationDate(v string) *GetKMSServiceKeyResponseBodyKmsKeyInfoList {
	s.CreationDate = &v
	return s
}

func (s *GetKMSServiceKeyResponseBodyKmsKeyInfoList) SetCreator(v string) *GetKMSServiceKeyResponseBodyKmsKeyInfoList {
	s.Creator = &v
	return s
}

func (s *GetKMSServiceKeyResponseBodyKmsKeyInfoList) SetDeleteDate(v string) *GetKMSServiceKeyResponseBodyKmsKeyInfoList {
	s.DeleteDate = &v
	return s
}

func (s *GetKMSServiceKeyResponseBodyKmsKeyInfoList) SetDescription(v string) *GetKMSServiceKeyResponseBodyKmsKeyInfoList {
	s.Description = &v
	return s
}

func (s *GetKMSServiceKeyResponseBodyKmsKeyInfoList) SetKeyId(v string) *GetKMSServiceKeyResponseBodyKmsKeyInfoList {
	s.KeyId = &v
	return s
}

func (s *GetKMSServiceKeyResponseBodyKmsKeyInfoList) SetKeyState(v string) *GetKMSServiceKeyResponseBodyKmsKeyInfoList {
	s.KeyState = &v
	return s
}

func (s *GetKMSServiceKeyResponseBodyKmsKeyInfoList) SetKeyUsage(v string) *GetKMSServiceKeyResponseBodyKmsKeyInfoList {
	s.KeyUsage = &v
	return s
}

func (s *GetKMSServiceKeyResponseBodyKmsKeyInfoList) Validate() error {
	return dara.Validate(s)
}
