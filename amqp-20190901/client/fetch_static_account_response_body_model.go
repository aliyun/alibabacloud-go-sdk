// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchStaticAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *FetchStaticAccountResponseBody
	GetCode() *int32
	SetData(v *FetchStaticAccountResponseBodyData) *FetchStaticAccountResponseBody
	GetData() *FetchStaticAccountResponseBodyData
	SetMessage(v string) *FetchStaticAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *FetchStaticAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FetchStaticAccountResponseBody
	GetSuccess() *bool
}

type FetchStaticAccountResponseBody struct {
	Code      *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *FetchStaticAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FetchStaticAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FetchStaticAccountResponseBody) GoString() string {
	return s.String()
}

func (s *FetchStaticAccountResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *FetchStaticAccountResponseBody) GetData() *FetchStaticAccountResponseBodyData {
	return s.Data
}

func (s *FetchStaticAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FetchStaticAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FetchStaticAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FetchStaticAccountResponseBody) SetCode(v int32) *FetchStaticAccountResponseBody {
	s.Code = &v
	return s
}

func (s *FetchStaticAccountResponseBody) SetData(v *FetchStaticAccountResponseBodyData) *FetchStaticAccountResponseBody {
	s.Data = v
	return s
}

func (s *FetchStaticAccountResponseBody) SetMessage(v string) *FetchStaticAccountResponseBody {
	s.Message = &v
	return s
}

func (s *FetchStaticAccountResponseBody) SetRequestId(v string) *FetchStaticAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchStaticAccountResponseBody) SetSuccess(v bool) *FetchStaticAccountResponseBody {
	s.Success = &v
	return s
}

func (s *FetchStaticAccountResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FetchStaticAccountResponseBodyData struct {
	// example:
	//
	// yourAccessKeyID
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// example:
	//
	// 1671175303522
	CreateTimeStamp *int64 `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	// example:
	//
	// amqp-cn-*********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1565***********01
	MasterUId *int64 `json:"MasterUId,omitempty" xml:"MasterUId,omitempty"`
	// example:
	//
	// OUYwQzM2QjZBRkUxNDRFM***************MzZCNzdDQzoxNjcxNDMwMzkyODI1
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// 备注示例
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// MjphbXFwLWNuLXVxbTJ6cjc2djAwMzpMVEFJNX*******ZNMWVSWnRFSjZ2Zm8=
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s FetchStaticAccountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FetchStaticAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *FetchStaticAccountResponseBodyData) GetAccessKey() *string {
	return s.AccessKey
}

func (s *FetchStaticAccountResponseBodyData) GetCreateTimeStamp() *int64 {
	return s.CreateTimeStamp
}

func (s *FetchStaticAccountResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *FetchStaticAccountResponseBodyData) GetMasterUId() *int64 {
	return s.MasterUId
}

func (s *FetchStaticAccountResponseBodyData) GetPassword() *string {
	return s.Password
}

func (s *FetchStaticAccountResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *FetchStaticAccountResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *FetchStaticAccountResponseBodyData) SetAccessKey(v string) *FetchStaticAccountResponseBodyData {
	s.AccessKey = &v
	return s
}

func (s *FetchStaticAccountResponseBodyData) SetCreateTimeStamp(v int64) *FetchStaticAccountResponseBodyData {
	s.CreateTimeStamp = &v
	return s
}

func (s *FetchStaticAccountResponseBodyData) SetInstanceId(v string) *FetchStaticAccountResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *FetchStaticAccountResponseBodyData) SetMasterUId(v int64) *FetchStaticAccountResponseBodyData {
	s.MasterUId = &v
	return s
}

func (s *FetchStaticAccountResponseBodyData) SetPassword(v string) *FetchStaticAccountResponseBodyData {
	s.Password = &v
	return s
}

func (s *FetchStaticAccountResponseBodyData) SetRemark(v string) *FetchStaticAccountResponseBodyData {
	s.Remark = &v
	return s
}

func (s *FetchStaticAccountResponseBodyData) SetUserName(v string) *FetchStaticAccountResponseBodyData {
	s.UserName = &v
	return s
}

func (s *FetchStaticAccountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
