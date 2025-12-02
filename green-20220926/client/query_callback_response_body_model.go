// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCallbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCryptType(v string) *QueryCallbackResponseBody
	GetCryptType() *string
	SetExistsOssCheckTask(v bool) *QueryCallbackResponseBody
	GetExistsOssCheckTask() *bool
	SetGmtCreate(v string) *QueryCallbackResponseBody
	GetGmtCreate() *string
	SetGmtModified(v string) *QueryCallbackResponseBody
	GetGmtModified() *string
	SetId(v int64) *QueryCallbackResponseBody
	GetId() *int64
	SetName(v string) *QueryCallbackResponseBody
	GetName() *string
	SetRequestId(v string) *QueryCallbackResponseBody
	GetRequestId() *string
	SetScope(v string) *QueryCallbackResponseBody
	GetScope() *string
	SetSeed(v string) *QueryCallbackResponseBody
	GetSeed() *string
	SetUid(v string) *QueryCallbackResponseBody
	GetUid() *string
	SetUrl(v string) *QueryCallbackResponseBody
	GetUrl() *string
}

type QueryCallbackResponseBody struct {
	// Encryption algorithm.
	//
	// example:
	//
	// SHA256
	CryptType *string `json:"CryptType,omitempty" xml:"CryptType,omitempty"`
	// Whether there is an OSS detection task.
	//
	// example:
	//
	// false
	ExistsOssCheckTask *bool `json:"ExistsOssCheckTask,omitempty" xml:"ExistsOssCheckTask,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 2022-11-30 16:30:29
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
	// ID assigned by the backend, used to uniquely identify a request. Can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// cb2MysbJTAAIf6gB3u4vpIEU-1ySnnf
	Seed *string `json:"Seed,omitempty" xml:"Seed,omitempty"`
	// UID.
	//
	// example:
	//
	// 19964*****086772
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// Callback URL.
	//
	// example:
	//
	// https://www.aliyuncs.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QueryCallbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCallbackResponseBody) GetCryptType() *string {
	return s.CryptType
}

func (s *QueryCallbackResponseBody) GetExistsOssCheckTask() *bool {
	return s.ExistsOssCheckTask
}

func (s *QueryCallbackResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryCallbackResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryCallbackResponseBody) GetId() *int64 {
	return s.Id
}

func (s *QueryCallbackResponseBody) GetName() *string {
	return s.Name
}

func (s *QueryCallbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCallbackResponseBody) GetScope() *string {
	return s.Scope
}

func (s *QueryCallbackResponseBody) GetSeed() *string {
	return s.Seed
}

func (s *QueryCallbackResponseBody) GetUid() *string {
	return s.Uid
}

func (s *QueryCallbackResponseBody) GetUrl() *string {
	return s.Url
}

func (s *QueryCallbackResponseBody) SetCryptType(v string) *QueryCallbackResponseBody {
	s.CryptType = &v
	return s
}

func (s *QueryCallbackResponseBody) SetExistsOssCheckTask(v bool) *QueryCallbackResponseBody {
	s.ExistsOssCheckTask = &v
	return s
}

func (s *QueryCallbackResponseBody) SetGmtCreate(v string) *QueryCallbackResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *QueryCallbackResponseBody) SetGmtModified(v string) *QueryCallbackResponseBody {
	s.GmtModified = &v
	return s
}

func (s *QueryCallbackResponseBody) SetId(v int64) *QueryCallbackResponseBody {
	s.Id = &v
	return s
}

func (s *QueryCallbackResponseBody) SetName(v string) *QueryCallbackResponseBody {
	s.Name = &v
	return s
}

func (s *QueryCallbackResponseBody) SetRequestId(v string) *QueryCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCallbackResponseBody) SetScope(v string) *QueryCallbackResponseBody {
	s.Scope = &v
	return s
}

func (s *QueryCallbackResponseBody) SetSeed(v string) *QueryCallbackResponseBody {
	s.Seed = &v
	return s
}

func (s *QueryCallbackResponseBody) SetUid(v string) *QueryCallbackResponseBody {
	s.Uid = &v
	return s
}

func (s *QueryCallbackResponseBody) SetUrl(v string) *QueryCallbackResponseBody {
	s.Url = &v
	return s
}

func (s *QueryCallbackResponseBody) Validate() error {
	return dara.Validate(s)
}
