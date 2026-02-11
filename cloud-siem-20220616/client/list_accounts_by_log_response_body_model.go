// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountsByLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListAccountsByLogResponseBodyData) *ListAccountsByLogResponseBody
	GetData() []*ListAccountsByLogResponseBodyData
	SetRequestId(v string) *ListAccountsByLogResponseBody
	GetRequestId() *string
}

type ListAccountsByLogResponseBody struct {
	// The data returned.
	Data []*ListAccountsByLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAccountsByLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsByLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountsByLogResponseBody) GetData() []*ListAccountsByLogResponseBodyData {
	return s.Data
}

func (s *ListAccountsByLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccountsByLogResponseBody) SetData(v []*ListAccountsByLogResponseBodyData) *ListAccountsByLogResponseBody {
	s.Data = v
	return s
}

func (s *ListAccountsByLogResponseBody) SetRequestId(v string) *ListAccountsByLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccountsByLogResponseBody) Validate() error {
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

type ListAccountsByLogResponseBodyData struct {
	// The ID of the cloud account.
	//
	// example:
	//
	// 123xxxxxxx
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the cloud account.
	//
	// example:
	//
	// sas_account_xxx
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Indicates whether the account is added. Valid values: -1: yes -0: no
	//
	// example:
	//
	// 123xxxxxxx
	Imported *int32 `json:"Imported,omitempty" xml:"Imported,omitempty"`
	// The code of the log.
	//
	// example:
	//
	// cloud_siem_waf_xxxxx
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// The ID of the Alibaba Cloud account that is used to purchase the threat analysis feature.
	//
	// example:
	//
	// 123XXXXXXXXX
	MainUserId *int64 `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	// The code of the service.
	//
	// example:
	//
	// qcloud_waf
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// The ID of the Alibaba Cloud account for which the threat analysis feature is enabled.
	//
	// example:
	//
	// 123XXXXXXXX
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s ListAccountsByLogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsByLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAccountsByLogResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *ListAccountsByLogResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *ListAccountsByLogResponseBodyData) GetImported() *int32 {
	return s.Imported
}

func (s *ListAccountsByLogResponseBodyData) GetLogCode() *string {
	return s.LogCode
}

func (s *ListAccountsByLogResponseBodyData) GetMainUserId() *int64 {
	return s.MainUserId
}

func (s *ListAccountsByLogResponseBodyData) GetProdCode() *string {
	return s.ProdCode
}

func (s *ListAccountsByLogResponseBodyData) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *ListAccountsByLogResponseBodyData) SetAccountId(v string) *ListAccountsByLogResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *ListAccountsByLogResponseBodyData) SetAccountName(v string) *ListAccountsByLogResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *ListAccountsByLogResponseBodyData) SetImported(v int32) *ListAccountsByLogResponseBodyData {
	s.Imported = &v
	return s
}

func (s *ListAccountsByLogResponseBodyData) SetLogCode(v string) *ListAccountsByLogResponseBodyData {
	s.LogCode = &v
	return s
}

func (s *ListAccountsByLogResponseBodyData) SetMainUserId(v int64) *ListAccountsByLogResponseBodyData {
	s.MainUserId = &v
	return s
}

func (s *ListAccountsByLogResponseBodyData) SetProdCode(v string) *ListAccountsByLogResponseBodyData {
	s.ProdCode = &v
	return s
}

func (s *ListAccountsByLogResponseBodyData) SetSubUserId(v int64) *ListAccountsByLogResponseBodyData {
	s.SubUserId = &v
	return s
}

func (s *ListAccountsByLogResponseBodyData) Validate() error {
	return dara.Validate(s)
}
