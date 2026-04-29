// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecallManagementTableRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *QueryRecallManagementTableRecordsRequest
	GetInstanceId() *string
	SetPrimaryKeys(v []byte) *QueryRecallManagementTableRecordsRequest
	GetPrimaryKeys() []byte
	SetRecallManagementTableVersionId(v string) *QueryRecallManagementTableRecordsRequest
	GetRecallManagementTableVersionId() *string
}

type QueryRecallManagementTableRecordsRequest struct {
	// example:
	//
	// pairec-cn-test123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ["1001","1002"]
	PrimaryKeys []byte `json:"PrimaryKeys,omitempty" xml:"PrimaryKeys,omitempty"`
	// example:
	//
	// 1
	RecallManagementTableVersionId *string `json:"RecallManagementTableVersionId,omitempty" xml:"RecallManagementTableVersionId,omitempty"`
}

func (s QueryRecallManagementTableRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRecallManagementTableRecordsRequest) GoString() string {
	return s.String()
}

func (s *QueryRecallManagementTableRecordsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryRecallManagementTableRecordsRequest) GetPrimaryKeys() []byte {
	return s.PrimaryKeys
}

func (s *QueryRecallManagementTableRecordsRequest) GetRecallManagementTableVersionId() *string {
	return s.RecallManagementTableVersionId
}

func (s *QueryRecallManagementTableRecordsRequest) SetInstanceId(v string) *QueryRecallManagementTableRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryRecallManagementTableRecordsRequest) SetPrimaryKeys(v []byte) *QueryRecallManagementTableRecordsRequest {
	s.PrimaryKeys = v
	return s
}

func (s *QueryRecallManagementTableRecordsRequest) SetRecallManagementTableVersionId(v string) *QueryRecallManagementTableRecordsRequest {
	s.RecallManagementTableVersionId = &v
	return s
}

func (s *QueryRecallManagementTableRecordsRequest) Validate() error {
	return dara.Validate(s)
}
