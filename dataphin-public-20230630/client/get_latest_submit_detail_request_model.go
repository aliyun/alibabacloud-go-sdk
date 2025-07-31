// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLatestSubmitDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetLatestSubmitDetailRequest
	GetOpTenantId() *int64
	SetSubmitDetailQuery(v *GetLatestSubmitDetailRequestSubmitDetailQuery) *GetLatestSubmitDetailRequest
	GetSubmitDetailQuery() *GetLatestSubmitDetailRequestSubmitDetailQuery
}

type GetLatestSubmitDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	SubmitDetailQuery *GetLatestSubmitDetailRequestSubmitDetailQuery `json:"SubmitDetailQuery,omitempty" xml:"SubmitDetailQuery,omitempty" type:"Struct"`
}

func (s GetLatestSubmitDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLatestSubmitDetailRequest) GoString() string {
	return s.String()
}

func (s *GetLatestSubmitDetailRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetLatestSubmitDetailRequest) GetSubmitDetailQuery() *GetLatestSubmitDetailRequestSubmitDetailQuery {
	return s.SubmitDetailQuery
}

func (s *GetLatestSubmitDetailRequest) SetOpTenantId(v int64) *GetLatestSubmitDetailRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetLatestSubmitDetailRequest) SetSubmitDetailQuery(v *GetLatestSubmitDetailRequestSubmitDetailQuery) *GetLatestSubmitDetailRequest {
	s.SubmitDetailQuery = v
	return s
}

func (s *GetLatestSubmitDetailRequest) Validate() error {
	return dara.Validate(s)
}

type GetLatestSubmitDetailRequestSubmitDetailQuery struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MAX_COMPUTE_SQL
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
}

func (s GetLatestSubmitDetailRequestSubmitDetailQuery) String() string {
	return dara.Prettify(s)
}

func (s GetLatestSubmitDetailRequestSubmitDetailQuery) GoString() string {
	return s.String()
}

func (s *GetLatestSubmitDetailRequestSubmitDetailQuery) GetObjectId() *string {
	return s.ObjectId
}

func (s *GetLatestSubmitDetailRequestSubmitDetailQuery) GetObjectType() *string {
	return s.ObjectType
}

func (s *GetLatestSubmitDetailRequestSubmitDetailQuery) SetObjectId(v string) *GetLatestSubmitDetailRequestSubmitDetailQuery {
	s.ObjectId = &v
	return s
}

func (s *GetLatestSubmitDetailRequestSubmitDetailQuery) SetObjectType(v string) *GetLatestSubmitDetailRequestSubmitDetailQuery {
	s.ObjectType = &v
	return s
}

func (s *GetLatestSubmitDetailRequestSubmitDetailQuery) Validate() error {
	return dara.Validate(s)
}
