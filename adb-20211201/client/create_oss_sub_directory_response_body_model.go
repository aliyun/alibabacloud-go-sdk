// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOssSubDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateOssSubDirectoryResponseBodyData) *CreateOssSubDirectoryResponseBody
	GetData() *CreateOssSubDirectoryResponseBodyData
	SetHttpStatusCode(v int64) *CreateOssSubDirectoryResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *CreateOssSubDirectoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateOssSubDirectoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateOssSubDirectoryResponseBody
	GetSuccess() *bool
}

type CreateOssSubDirectoryResponseBody struct {
	// The returned data.
	Data *CreateOssSubDirectoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// 	- If the request was successful, a **success*	- message is returned.
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3A0DE2E0-A37B-5EE4-9136-C4C473714802
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateOssSubDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOssSubDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOssSubDirectoryResponseBody) GetData() *CreateOssSubDirectoryResponseBodyData {
	return s.Data
}

func (s *CreateOssSubDirectoryResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *CreateOssSubDirectoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateOssSubDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOssSubDirectoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateOssSubDirectoryResponseBody) SetData(v *CreateOssSubDirectoryResponseBodyData) *CreateOssSubDirectoryResponseBody {
	s.Data = v
	return s
}

func (s *CreateOssSubDirectoryResponseBody) SetHttpStatusCode(v int64) *CreateOssSubDirectoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateOssSubDirectoryResponseBody) SetMessage(v string) *CreateOssSubDirectoryResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOssSubDirectoryResponseBody) SetRequestId(v string) *CreateOssSubDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOssSubDirectoryResponseBody) SetSuccess(v bool) *CreateOssSubDirectoryResponseBody {
	s.Success = &v
	return s
}

func (s *CreateOssSubDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateOssSubDirectoryResponseBodyData struct {
	// The cyclic redundancy check (CRC) value on the client.
	//
	// example:
	//
	// 1
	ClientCRC *int64 `json:"ClientCRC,omitempty" xml:"ClientCRC,omitempty"`
	// The tag of the OSS path.
	//
	// example:
	//
	// 1
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3A0DE2E0-A37B-5EE4-9136-C4C473714802
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The CRC-64 value on the OSS bucket.
	//
	// example:
	//
	// 1
	ServerCRC *int64 `json:"ServerCRC,omitempty" xml:"ServerCRC,omitempty"`
}

func (s CreateOssSubDirectoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateOssSubDirectoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateOssSubDirectoryResponseBodyData) GetClientCRC() *int64 {
	return s.ClientCRC
}

func (s *CreateOssSubDirectoryResponseBodyData) GetETag() *string {
	return s.ETag
}

func (s *CreateOssSubDirectoryResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOssSubDirectoryResponseBodyData) GetServerCRC() *int64 {
	return s.ServerCRC
}

func (s *CreateOssSubDirectoryResponseBodyData) SetClientCRC(v int64) *CreateOssSubDirectoryResponseBodyData {
	s.ClientCRC = &v
	return s
}

func (s *CreateOssSubDirectoryResponseBodyData) SetETag(v string) *CreateOssSubDirectoryResponseBodyData {
	s.ETag = &v
	return s
}

func (s *CreateOssSubDirectoryResponseBodyData) SetRequestId(v string) *CreateOssSubDirectoryResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *CreateOssSubDirectoryResponseBodyData) SetServerCRC(v int64) *CreateOssSubDirectoryResponseBodyData {
	s.ServerCRC = &v
	return s
}

func (s *CreateOssSubDirectoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
