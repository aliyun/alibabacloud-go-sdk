// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSQLReviewOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateSQLReviewOrderRequest
	GetComment() *string
	SetParam(v *CreateSQLReviewOrderRequestParam) *CreateSQLReviewOrderRequest
	GetParam() *CreateSQLReviewOrderRequestParam
	SetRelatedUserList(v []*int64) *CreateSQLReviewOrderRequest
	GetRelatedUserList() []*int64
	SetTid(v int64) *CreateSQLReviewOrderRequest
	GetTid() *int64
}

type CreateSQLReviewOrderRequest struct {
	// The purpose or objective of the SQL review. This reduces unnecessary communication.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The parameters of the ticket.
	//
	// This parameter is required.
	Param *CreateSQLReviewOrderRequestParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	// The stakeholders involved in this operation. All the specified stakeholders can view the ticket details and take part in the approval process. Irrelevant users other than DMS administrators and database administrators (DBAs) are not allowed to view the ticket details.
	RelatedUserList []*int64 `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty" type:"Repeated"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 1
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateSQLReviewOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSQLReviewOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateSQLReviewOrderRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateSQLReviewOrderRequest) GetParam() *CreateSQLReviewOrderRequestParam {
	return s.Param
}

func (s *CreateSQLReviewOrderRequest) GetRelatedUserList() []*int64 {
	return s.RelatedUserList
}

func (s *CreateSQLReviewOrderRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateSQLReviewOrderRequest) SetComment(v string) *CreateSQLReviewOrderRequest {
	s.Comment = &v
	return s
}

func (s *CreateSQLReviewOrderRequest) SetParam(v *CreateSQLReviewOrderRequestParam) *CreateSQLReviewOrderRequest {
	s.Param = v
	return s
}

func (s *CreateSQLReviewOrderRequest) SetRelatedUserList(v []*int64) *CreateSQLReviewOrderRequest {
	s.RelatedUserList = v
	return s
}

func (s *CreateSQLReviewOrderRequest) SetTid(v int64) *CreateSQLReviewOrderRequest {
	s.Tid = &v
	return s
}

func (s *CreateSQLReviewOrderRequest) Validate() error {
	if s.Param != nil {
		if err := s.Param.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSQLReviewOrderRequestParam struct {
	// The files to be reviewed. Multiple files can be reviewed at a time.
	//
	// This parameter is required.
	AttachmentKeyList []*string `json:"AttachmentKeyList,omitempty" xml:"AttachmentKeyList,omitempty" type:"Repeated"`
	// The ID of the database. You can call the [SearchDatabases](https://help.aliyun.com/document_detail/141876.html) operation to query the ID of the database.
	//
	// >  You can call this operation to query only physical databases. This operation is unavailable to query logical databases.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123321
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The name of the project.
	//
	// This parameter is required.
	//
	// example:
	//
	// SQL review for xxx
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s CreateSQLReviewOrderRequestParam) String() string {
	return dara.Prettify(s)
}

func (s CreateSQLReviewOrderRequestParam) GoString() string {
	return s.String()
}

func (s *CreateSQLReviewOrderRequestParam) GetAttachmentKeyList() []*string {
	return s.AttachmentKeyList
}

func (s *CreateSQLReviewOrderRequestParam) GetDbId() *int64 {
	return s.DbId
}

func (s *CreateSQLReviewOrderRequestParam) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateSQLReviewOrderRequestParam) SetAttachmentKeyList(v []*string) *CreateSQLReviewOrderRequestParam {
	s.AttachmentKeyList = v
	return s
}

func (s *CreateSQLReviewOrderRequestParam) SetDbId(v int64) *CreateSQLReviewOrderRequestParam {
	s.DbId = &v
	return s
}

func (s *CreateSQLReviewOrderRequestParam) SetProjectName(v string) *CreateSQLReviewOrderRequestParam {
	s.ProjectName = &v
	return s
}

func (s *CreateSQLReviewOrderRequestParam) Validate() error {
	return dara.Validate(s)
}
