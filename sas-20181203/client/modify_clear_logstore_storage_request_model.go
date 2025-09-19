// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClearLogstoreStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *ModifyClearLogstoreStorageRequest
	GetFrom() *string
	SetLang(v string) *ModifyClearLogstoreStorageRequest
	GetLang() *string
	SetUserLogStore(v string) *ModifyClearLogstoreStorageRequest
	GetUserLogStore() *string
	SetUserProject(v string) *ModifyClearLogstoreStorageRequest
	GetUserProject() *string
}

type ModifyClearLogstoreStorageRequest struct {
	// The ID of the request source. Set the value to **sas**.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the Logstore that stores logs.
	//
	// example:
	//
	// sas_sls_storage
	UserLogStore *string `json:"UserLogStore,omitempty" xml:"UserLogStore,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// sas-log-1234(uid)-cn-hangzhou
	UserProject *string `json:"UserProject,omitempty" xml:"UserProject,omitempty"`
}

func (s ModifyClearLogstoreStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClearLogstoreStorageRequest) GoString() string {
	return s.String()
}

func (s *ModifyClearLogstoreStorageRequest) GetFrom() *string {
	return s.From
}

func (s *ModifyClearLogstoreStorageRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyClearLogstoreStorageRequest) GetUserLogStore() *string {
	return s.UserLogStore
}

func (s *ModifyClearLogstoreStorageRequest) GetUserProject() *string {
	return s.UserProject
}

func (s *ModifyClearLogstoreStorageRequest) SetFrom(v string) *ModifyClearLogstoreStorageRequest {
	s.From = &v
	return s
}

func (s *ModifyClearLogstoreStorageRequest) SetLang(v string) *ModifyClearLogstoreStorageRequest {
	s.Lang = &v
	return s
}

func (s *ModifyClearLogstoreStorageRequest) SetUserLogStore(v string) *ModifyClearLogstoreStorageRequest {
	s.UserLogStore = &v
	return s
}

func (s *ModifyClearLogstoreStorageRequest) SetUserProject(v string) *ModifyClearLogstoreStorageRequest {
	s.UserProject = &v
	return s
}

func (s *ModifyClearLogstoreStorageRequest) Validate() error {
	return dara.Validate(s)
}
