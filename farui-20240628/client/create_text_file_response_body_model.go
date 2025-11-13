// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTextFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTextFileResponseBody
	GetCode() *string
	SetData(v *CreateTextFileResponseBodyData) *CreateTextFileResponseBody
	GetData() *CreateTextFileResponseBodyData
	SetHttpStatusCode(v int64) *CreateTextFileResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *CreateTextFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTextFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTextFileResponseBody
	GetSuccess() *bool
}

type CreateTextFileResponseBody struct {
	// example:
	//
	// Request.Signature.Error
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateTextFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 81E6F6D2-8ACB-5BDA-9C7C-4D6268CD9652
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTextFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTextFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTextFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTextFileResponseBody) GetData() *CreateTextFileResponseBodyData {
	return s.Data
}

func (s *CreateTextFileResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *CreateTextFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTextFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTextFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTextFileResponseBody) SetCode(v string) *CreateTextFileResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTextFileResponseBody) SetData(v *CreateTextFileResponseBodyData) *CreateTextFileResponseBody {
	s.Data = v
	return s
}

func (s *CreateTextFileResponseBody) SetHttpStatusCode(v int64) *CreateTextFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTextFileResponseBody) SetMessage(v string) *CreateTextFileResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTextFileResponseBody) SetRequestId(v string) *CreateTextFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTextFileResponseBody) SetSuccess(v bool) *CreateTextFileResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTextFileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTextFileResponseBodyData struct {
	ContractId *string `json:"ContractId,omitempty" xml:"ContractId,omitempty"`
	// example:
	//
	// 36d6447d277c4a1c9fd0def1d16341f1
	TextFileId   *string `json:"TextFileId,omitempty" xml:"TextFileId,omitempty"`
	TextFileName *string `json:"TextFileName,omitempty" xml:"TextFileName,omitempty"`
	TextFileUrl  *string `json:"TextFileUrl,omitempty" xml:"TextFileUrl,omitempty"`
}

func (s CreateTextFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateTextFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTextFileResponseBodyData) GetContractId() *string {
	return s.ContractId
}

func (s *CreateTextFileResponseBodyData) GetTextFileId() *string {
	return s.TextFileId
}

func (s *CreateTextFileResponseBodyData) GetTextFileName() *string {
	return s.TextFileName
}

func (s *CreateTextFileResponseBodyData) GetTextFileUrl() *string {
	return s.TextFileUrl
}

func (s *CreateTextFileResponseBodyData) SetContractId(v string) *CreateTextFileResponseBodyData {
	s.ContractId = &v
	return s
}

func (s *CreateTextFileResponseBodyData) SetTextFileId(v string) *CreateTextFileResponseBodyData {
	s.TextFileId = &v
	return s
}

func (s *CreateTextFileResponseBodyData) SetTextFileName(v string) *CreateTextFileResponseBodyData {
	s.TextFileName = &v
	return s
}

func (s *CreateTextFileResponseBodyData) SetTextFileUrl(v string) *CreateTextFileResponseBodyData {
	s.TextFileUrl = &v
	return s
}

func (s *CreateTextFileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
