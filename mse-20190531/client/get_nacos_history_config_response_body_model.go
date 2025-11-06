// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNacosHistoryConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfiguration(v *GetNacosHistoryConfigResponseBodyConfiguration) *GetNacosHistoryConfigResponseBody
	GetConfiguration() *GetNacosHistoryConfigResponseBodyConfiguration
	SetErrorCode(v string) *GetNacosHistoryConfigResponseBody
	GetErrorCode() *string
	SetMessage(v string) *GetNacosHistoryConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetNacosHistoryConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetNacosHistoryConfigResponseBody
	GetSuccess() *bool
}

type GetNacosHistoryConfigResponseBody struct {
	// The configuration information.
	Configuration *GetNacosHistoryConfigResponseBodyConfiguration `json:"Configuration,omitempty" xml:"Configuration,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 69AD2AA7-DB47-449B-941B-B14409DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNacosHistoryConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNacosHistoryConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetNacosHistoryConfigResponseBody) GetConfiguration() *GetNacosHistoryConfigResponseBodyConfiguration {
	return s.Configuration
}

func (s *GetNacosHistoryConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetNacosHistoryConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetNacosHistoryConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNacosHistoryConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetNacosHistoryConfigResponseBody) SetConfiguration(v *GetNacosHistoryConfigResponseBodyConfiguration) *GetNacosHistoryConfigResponseBody {
	s.Configuration = v
	return s
}

func (s *GetNacosHistoryConfigResponseBody) SetErrorCode(v string) *GetNacosHistoryConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetNacosHistoryConfigResponseBody) SetMessage(v string) *GetNacosHistoryConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetNacosHistoryConfigResponseBody) SetRequestId(v string) *GetNacosHistoryConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNacosHistoryConfigResponseBody) SetSuccess(v bool) *GetNacosHistoryConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetNacosHistoryConfigResponseBody) Validate() error {
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNacosHistoryConfigResponseBodyConfiguration struct {
	// The name of the application.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The content of the configuration.
	//
	// example:
	//
	// test
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the configuration.
	//
	// example:
	//
	// test.yaml
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// The encryption key.
	//
	// example:
	//
	// 23fds****
	EncryptedDataKey *string `json:"EncryptedDataKey,omitempty" xml:"EncryptedDataKey,omitempty"`
	// The name of the configuration group.
	//
	// example:
	//
	// public
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The MD5 value of the configuration.
	//
	// example:
	//
	// 23sdf32f****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The configuration type.
	//
	// example:
	//
	// text
	OpType *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
}

func (s GetNacosHistoryConfigResponseBodyConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetNacosHistoryConfigResponseBodyConfiguration) GoString() string {
	return s.String()
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) GetAppName() *string {
	return s.AppName
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) GetContent() *string {
	return s.Content
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) GetDataId() *string {
	return s.DataId
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) GetEncryptedDataKey() *string {
	return s.EncryptedDataKey
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) GetGroup() *string {
	return s.Group
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) GetMd5() *string {
	return s.Md5
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) GetOpType() *string {
	return s.OpType
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) SetAppName(v string) *GetNacosHistoryConfigResponseBodyConfiguration {
	s.AppName = &v
	return s
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) SetContent(v string) *GetNacosHistoryConfigResponseBodyConfiguration {
	s.Content = &v
	return s
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) SetDataId(v string) *GetNacosHistoryConfigResponseBodyConfiguration {
	s.DataId = &v
	return s
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) SetEncryptedDataKey(v string) *GetNacosHistoryConfigResponseBodyConfiguration {
	s.EncryptedDataKey = &v
	return s
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) SetGroup(v string) *GetNacosHistoryConfigResponseBodyConfiguration {
	s.Group = &v
	return s
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) SetMd5(v string) *GetNacosHistoryConfigResponseBodyConfiguration {
	s.Md5 = &v
	return s
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) SetOpType(v string) *GetNacosHistoryConfigResponseBodyConfiguration {
	s.OpType = &v
	return s
}

func (s *GetNacosHistoryConfigResponseBodyConfiguration) Validate() error {
	return dara.Validate(s)
}
