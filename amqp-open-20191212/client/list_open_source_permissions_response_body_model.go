// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpenSourcePermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListOpenSourcePermissionsResponseBody
	GetCode() *int32
	SetData(v []*ListOpenSourcePermissionsResponseBodyData) *ListOpenSourcePermissionsResponseBody
	GetData() []*ListOpenSourcePermissionsResponseBodyData
	SetMaxResults(v int32) *ListOpenSourcePermissionsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListOpenSourcePermissionsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListOpenSourcePermissionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListOpenSourcePermissionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListOpenSourcePermissionsResponseBody
	GetSuccess() *bool
}

type ListOpenSourcePermissionsResponseBody struct {
	// The return code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*ListOpenSourcePermissionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The maximum number of results to return.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token for the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FEBA5E0C-50D0-4FA6-A794-4901E5465***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListOpenSourcePermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOpenSourcePermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOpenSourcePermissionsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListOpenSourcePermissionsResponseBody) GetData() []*ListOpenSourcePermissionsResponseBodyData {
	return s.Data
}

func (s *ListOpenSourcePermissionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOpenSourcePermissionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListOpenSourcePermissionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOpenSourcePermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOpenSourcePermissionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListOpenSourcePermissionsResponseBody) SetCode(v int32) *ListOpenSourcePermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListOpenSourcePermissionsResponseBody) SetData(v []*ListOpenSourcePermissionsResponseBodyData) *ListOpenSourcePermissionsResponseBody {
	s.Data = v
	return s
}

func (s *ListOpenSourcePermissionsResponseBody) SetMaxResults(v int32) *ListOpenSourcePermissionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListOpenSourcePermissionsResponseBody) SetMessage(v string) *ListOpenSourcePermissionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListOpenSourcePermissionsResponseBody) SetNextToken(v string) *ListOpenSourcePermissionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOpenSourcePermissionsResponseBody) SetRequestId(v string) *ListOpenSourcePermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOpenSourcePermissionsResponseBody) SetSuccess(v bool) *ListOpenSourcePermissionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListOpenSourcePermissionsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOpenSourcePermissionsResponseBodyData struct {
	// The Alibaba Cloud UID.
	//
	// example:
	//
	// 1234567890123456
	AliyunUserId *int64 `json:"AliyunUserId,omitempty" xml:"AliyunUserId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// amqp-cn-v0h1kb9nu***
	CInstanceId *string `json:"CInstanceId,omitempty" xml:"CInstanceId,omitempty"`
	// The regular expression for configure permissions.
	//
	// example:
	//
	// ^$
	Configure *string `json:"Configure,omitempty" xml:"Configure,omitempty"`
	// The username.
	//
	// example:
	//
	// shhtzn
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The regular expression for read permissions.
	//
	// example:
	//
	// ^$
	Read *string `json:"Read,omitempty" xml:"Read,omitempty"`
	// The vhost name.
	//
	// example:
	//
	// production
	Vhost *string `json:"Vhost,omitempty" xml:"Vhost,omitempty"`
	// The regular expression for write permissions.
	//
	// example:
	//
	// order_exchange
	Write *string `json:"Write,omitempty" xml:"Write,omitempty"`
}

func (s ListOpenSourcePermissionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListOpenSourcePermissionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOpenSourcePermissionsResponseBodyData) GetAliyunUserId() *int64 {
	return s.AliyunUserId
}

func (s *ListOpenSourcePermissionsResponseBodyData) GetCInstanceId() *string {
	return s.CInstanceId
}

func (s *ListOpenSourcePermissionsResponseBodyData) GetConfigure() *string {
	return s.Configure
}

func (s *ListOpenSourcePermissionsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListOpenSourcePermissionsResponseBodyData) GetRead() *string {
	return s.Read
}

func (s *ListOpenSourcePermissionsResponseBodyData) GetVhost() *string {
	return s.Vhost
}

func (s *ListOpenSourcePermissionsResponseBodyData) GetWrite() *string {
	return s.Write
}

func (s *ListOpenSourcePermissionsResponseBodyData) SetAliyunUserId(v int64) *ListOpenSourcePermissionsResponseBodyData {
	s.AliyunUserId = &v
	return s
}

func (s *ListOpenSourcePermissionsResponseBodyData) SetCInstanceId(v string) *ListOpenSourcePermissionsResponseBodyData {
	s.CInstanceId = &v
	return s
}

func (s *ListOpenSourcePermissionsResponseBodyData) SetConfigure(v string) *ListOpenSourcePermissionsResponseBodyData {
	s.Configure = &v
	return s
}

func (s *ListOpenSourcePermissionsResponseBodyData) SetName(v string) *ListOpenSourcePermissionsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListOpenSourcePermissionsResponseBodyData) SetRead(v string) *ListOpenSourcePermissionsResponseBodyData {
	s.Read = &v
	return s
}

func (s *ListOpenSourcePermissionsResponseBodyData) SetVhost(v string) *ListOpenSourcePermissionsResponseBodyData {
	s.Vhost = &v
	return s
}

func (s *ListOpenSourcePermissionsResponseBodyData) SetWrite(v string) *ListOpenSourcePermissionsResponseBodyData {
	s.Write = &v
	return s
}

func (s *ListOpenSourcePermissionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
