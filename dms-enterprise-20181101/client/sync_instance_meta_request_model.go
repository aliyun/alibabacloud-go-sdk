// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncInstanceMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIgnoreTable(v bool) *SyncInstanceMetaRequest
	GetIgnoreTable() *bool
	SetInstanceId(v string) *SyncInstanceMetaRequest
	GetInstanceId() *string
	SetTid(v int64) *SyncInstanceMetaRequest
	GetTid() *int64
}

type SyncInstanceMetaRequest struct {
	// Specifies whether to skip synchronization for the metadata of table dictionaries. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IgnoreTable *bool `json:"IgnoreTable,omitempty" xml:"IgnoreTable,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see the [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html) section of the Manage DMS tenants topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s SyncInstanceMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncInstanceMetaRequest) GoString() string {
	return s.String()
}

func (s *SyncInstanceMetaRequest) GetIgnoreTable() *bool {
	return s.IgnoreTable
}

func (s *SyncInstanceMetaRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SyncInstanceMetaRequest) GetTid() *int64 {
	return s.Tid
}

func (s *SyncInstanceMetaRequest) SetIgnoreTable(v bool) *SyncInstanceMetaRequest {
	s.IgnoreTable = &v
	return s
}

func (s *SyncInstanceMetaRequest) SetInstanceId(v string) *SyncInstanceMetaRequest {
	s.InstanceId = &v
	return s
}

func (s *SyncInstanceMetaRequest) SetTid(v int64) *SyncInstanceMetaRequest {
	s.Tid = &v
	return s
}

func (s *SyncInstanceMetaRequest) Validate() error {
	return dara.Validate(s)
}
