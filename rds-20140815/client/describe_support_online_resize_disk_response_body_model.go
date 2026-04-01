// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupportOnlineResizeDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSupportOnlineResizeDiskResponseBody
	GetCode() *string
	SetData(v string) *DescribeSupportOnlineResizeDiskResponseBody
	GetData() *string
	SetMessage(v string) *DescribeSupportOnlineResizeDiskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSupportOnlineResizeDiskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSupportOnlineResizeDiskResponseBody
	GetSuccess() *bool
}

type DescribeSupportOnlineResizeDiskResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// NotExists.InstanceId
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response result set.
	//
	// example:
	//
	// {"SupportOnlineResizeDisk":true,"DBInstanceName":"rm-uf6wjk5xxxxxxx"}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The response code.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8B993DA9-5272-5414-94E3-4CA8BA0146C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSupportOnlineResizeDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportOnlineResizeDiskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupportOnlineResizeDiskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSupportOnlineResizeDiskResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeSupportOnlineResizeDiskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSupportOnlineResizeDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSupportOnlineResizeDiskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSupportOnlineResizeDiskResponseBody) SetCode(v string) *DescribeSupportOnlineResizeDiskResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSupportOnlineResizeDiskResponseBody) SetData(v string) *DescribeSupportOnlineResizeDiskResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeSupportOnlineResizeDiskResponseBody) SetMessage(v string) *DescribeSupportOnlineResizeDiskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSupportOnlineResizeDiskResponseBody) SetRequestId(v string) *DescribeSupportOnlineResizeDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSupportOnlineResizeDiskResponseBody) SetSuccess(v bool) *DescribeSupportOnlineResizeDiskResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSupportOnlineResizeDiskResponseBody) Validate() error {
	return dara.Validate(s)
}
