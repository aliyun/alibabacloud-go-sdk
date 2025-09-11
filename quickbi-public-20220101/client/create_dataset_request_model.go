// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDsId(v string) *CreateDatasetRequest
	GetDsId() *string
	SetTableName(v string) *CreateDatasetRequest
	GetTableName() *string
	SetTargetDirectoryId(v string) *CreateDatasetRequest
	GetTargetDirectoryId() *string
	SetUserDefineCubeName(v string) *CreateDatasetRequest
	GetUserDefineCubeName() *string
	SetUserId(v string) *CreateDatasetRequest
	GetUserId() *string
	SetWorkspaceId(v string) *CreateDatasetRequest
	GetWorkspaceId() *string
}

type CreateDatasetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 7AAB95D-*****-****-*4FC0C976
	DsId *string `json:"DsId,omitempty" xml:"DsId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// st_trd_user_purchase_day_inc
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// asah-fsdfsf*****92342
	TargetDirectoryId *string `json:"TargetDirectoryId,omitempty" xml:"TargetDirectoryId,omitempty"`
	// This parameter is required.
	UserDefineCubeName *string `json:"UserDefineCubeName,omitempty" xml:"UserDefineCubeName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 726bee5a-****-43e1-9a8e-b550f0120f35
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 726bee5a-****-43e1-9a8e-b550f0120f35
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetRequest) GetDsId() *string {
	return s.DsId
}

func (s *CreateDatasetRequest) GetTableName() *string {
	return s.TableName
}

func (s *CreateDatasetRequest) GetTargetDirectoryId() *string {
	return s.TargetDirectoryId
}

func (s *CreateDatasetRequest) GetUserDefineCubeName() *string {
	return s.UserDefineCubeName
}

func (s *CreateDatasetRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateDatasetRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDatasetRequest) SetDsId(v string) *CreateDatasetRequest {
	s.DsId = &v
	return s
}

func (s *CreateDatasetRequest) SetTableName(v string) *CreateDatasetRequest {
	s.TableName = &v
	return s
}

func (s *CreateDatasetRequest) SetTargetDirectoryId(v string) *CreateDatasetRequest {
	s.TargetDirectoryId = &v
	return s
}

func (s *CreateDatasetRequest) SetUserDefineCubeName(v string) *CreateDatasetRequest {
	s.UserDefineCubeName = &v
	return s
}

func (s *CreateDatasetRequest) SetUserId(v string) *CreateDatasetRequest {
	s.UserId = &v
	return s
}

func (s *CreateDatasetRequest) SetWorkspaceId(v string) *CreateDatasetRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDatasetRequest) Validate() error {
	return dara.Validate(s)
}
