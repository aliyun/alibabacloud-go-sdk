// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadIoTDataToBlockchainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UploadIoTDataToBlockchainResponseBody
	GetCode() *string
	SetData(v string) *UploadIoTDataToBlockchainResponseBody
	GetData() *string
	SetMessage(v string) *UploadIoTDataToBlockchainResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadIoTDataToBlockchainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UploadIoTDataToBlockchainResponseBody
	GetSuccess() *bool
}

type UploadIoTDataToBlockchainResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadIoTDataToBlockchainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadIoTDataToBlockchainResponseBody) GoString() string {
	return s.String()
}

func (s *UploadIoTDataToBlockchainResponseBody) GetCode() *string {
	return s.Code
}

func (s *UploadIoTDataToBlockchainResponseBody) GetData() *string {
	return s.Data
}

func (s *UploadIoTDataToBlockchainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadIoTDataToBlockchainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadIoTDataToBlockchainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UploadIoTDataToBlockchainResponseBody) SetCode(v string) *UploadIoTDataToBlockchainResponseBody {
	s.Code = &v
	return s
}

func (s *UploadIoTDataToBlockchainResponseBody) SetData(v string) *UploadIoTDataToBlockchainResponseBody {
	s.Data = &v
	return s
}

func (s *UploadIoTDataToBlockchainResponseBody) SetMessage(v string) *UploadIoTDataToBlockchainResponseBody {
	s.Message = &v
	return s
}

func (s *UploadIoTDataToBlockchainResponseBody) SetRequestId(v string) *UploadIoTDataToBlockchainResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadIoTDataToBlockchainResponseBody) SetSuccess(v bool) *UploadIoTDataToBlockchainResponseBody {
	s.Success = &v
	return s
}

func (s *UploadIoTDataToBlockchainResponseBody) Validate() error {
	return dara.Validate(s)
}
