// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBindAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListBindAccountResponseBodyData) *ListBindAccountResponseBody
	GetData() []*ListBindAccountResponseBodyData
	SetRequestId(v string) *ListBindAccountResponseBody
	GetRequestId() *string
}

type ListBindAccountResponseBody struct {
	// The data returned.
	Data []*ListBindAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBindAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBindAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ListBindAccountResponseBody) GetData() []*ListBindAccountResponseBodyData {
	return s.Data
}

func (s *ListBindAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBindAccountResponseBody) SetData(v []*ListBindAccountResponseBodyData) *ListBindAccountResponseBody {
	s.Data = v
	return s
}

func (s *ListBindAccountResponseBody) SetRequestId(v string) *ListBindAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBindAccountResponseBody) Validate() error {
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

type ListBindAccountResponseBodyData struct {
	// The AccessKey ID of the cloud account.
	//
	// example:
	//
	// ABCXXXXXXXX
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// The ID of the cloud account.
	//
	// example:
	//
	// 123xxxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The username of the cloud account.
	//
	// example:
	//
	// sas_account_xxx
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The ID that is generated when the cloud account is added.
	//
	// example:
	//
	// 123xxxxxxx
	BindId *int64 `json:"BindId,omitempty" xml:"BindId,omitempty"`
	// The code of the cloud service provider. Valid values:
	//
	// 	- qcloud: Tencent Cloud
	//
	// 	- aliyun: Alibaba Cloud
	//
	// 	- hcloud: Huawei Cloud
	//
	// example:
	//
	// hcloud
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The ID of the account that is used to add the cloud account.
	//
	// example:
	//
	// 123xxxxxxx
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The number of data sources that are added to the threat analysis feature within the cloud account.
	//
	// example:
	//
	// 2
	DataSourceCount *int64 `json:"DataSourceCount,omitempty" xml:"DataSourceCount,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2023-11-10 12:20:35
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s ListBindAccountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBindAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBindAccountResponseBodyData) GetAccessId() *string {
	return s.AccessId
}

func (s *ListBindAccountResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *ListBindAccountResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *ListBindAccountResponseBodyData) GetBindId() *int64 {
	return s.BindId
}

func (s *ListBindAccountResponseBodyData) GetCloudCode() *string {
	return s.CloudCode
}

func (s *ListBindAccountResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListBindAccountResponseBodyData) GetDataSourceCount() *int64 {
	return s.DataSourceCount
}

func (s *ListBindAccountResponseBodyData) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListBindAccountResponseBodyData) SetAccessId(v string) *ListBindAccountResponseBodyData {
	s.AccessId = &v
	return s
}

func (s *ListBindAccountResponseBodyData) SetAccountId(v string) *ListBindAccountResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *ListBindAccountResponseBodyData) SetAccountName(v string) *ListBindAccountResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *ListBindAccountResponseBodyData) SetBindId(v int64) *ListBindAccountResponseBodyData {
	s.BindId = &v
	return s
}

func (s *ListBindAccountResponseBodyData) SetCloudCode(v string) *ListBindAccountResponseBodyData {
	s.CloudCode = &v
	return s
}

func (s *ListBindAccountResponseBodyData) SetCreateUser(v string) *ListBindAccountResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListBindAccountResponseBodyData) SetDataSourceCount(v int64) *ListBindAccountResponseBodyData {
	s.DataSourceCount = &v
	return s
}

func (s *ListBindAccountResponseBodyData) SetModifyTime(v string) *ListBindAccountResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *ListBindAccountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
