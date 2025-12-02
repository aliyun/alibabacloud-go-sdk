// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListCallbackResponseBodyData) *ListCallbackResponseBody
	GetData() []*ListCallbackResponseBodyData
	SetRequestId(v string) *ListCallbackResponseBody
	GetRequestId() *string
}

type ListCallbackResponseBody struct {
	// Returned data.
	Data []*ListCallbackResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Backend-assigned ID, used to uniquely identify a request. Can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCallbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *ListCallbackResponseBody) GetData() []*ListCallbackResponseBodyData {
	return s.Data
}

func (s *ListCallbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCallbackResponseBody) SetData(v []*ListCallbackResponseBodyData) *ListCallbackResponseBody {
	s.Data = v
	return s
}

func (s *ListCallbackResponseBody) SetRequestId(v string) *ListCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCallbackResponseBody) Validate() error {
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

type ListCallbackResponseBodyData struct {
	// Encryption algorithm.
	//
	// example:
	//
	// SHA256
	CryptType *string `json:"CryptType,omitempty" xml:"CryptType,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 2024-06-03 15:20:14
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 2024-06-03 15:20:14
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Primary key ID.
	//
	// example:
	//
	// 11234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Name.
	//
	// example:
	//
	// 回调通知
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Result scope.
	//
	// example:
	//
	// all
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// Seed.
	//
	// example:
	//
	// cbupVnpBjkgjFxfINMHKkrHS-1zZPUm
	Seed *string `json:"Seed,omitempty" xml:"Seed,omitempty"`
	// UID.
	//
	// example:
	//
	// 16537*****831937
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// Callback URL.
	//
	// example:
	//
	// https://console.aliyun.com/
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListCallbackResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCallbackResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCallbackResponseBodyData) GetCryptType() *string {
	return s.CryptType
}

func (s *ListCallbackResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListCallbackResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListCallbackResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListCallbackResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListCallbackResponseBodyData) GetScope() *string {
	return s.Scope
}

func (s *ListCallbackResponseBodyData) GetSeed() *string {
	return s.Seed
}

func (s *ListCallbackResponseBodyData) GetUid() *string {
	return s.Uid
}

func (s *ListCallbackResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *ListCallbackResponseBodyData) SetCryptType(v string) *ListCallbackResponseBodyData {
	s.CryptType = &v
	return s
}

func (s *ListCallbackResponseBodyData) SetGmtCreate(v string) *ListCallbackResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListCallbackResponseBodyData) SetGmtModified(v string) *ListCallbackResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListCallbackResponseBodyData) SetId(v int64) *ListCallbackResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListCallbackResponseBodyData) SetName(v string) *ListCallbackResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListCallbackResponseBodyData) SetScope(v string) *ListCallbackResponseBodyData {
	s.Scope = &v
	return s
}

func (s *ListCallbackResponseBodyData) SetSeed(v string) *ListCallbackResponseBodyData {
	s.Seed = &v
	return s
}

func (s *ListCallbackResponseBodyData) SetUid(v string) *ListCallbackResponseBodyData {
	s.Uid = &v
	return s
}

func (s *ListCallbackResponseBodyData) SetUrl(v string) *ListCallbackResponseBodyData {
	s.Url = &v
	return s
}

func (s *ListCallbackResponseBodyData) Validate() error {
	return dara.Validate(s)
}
