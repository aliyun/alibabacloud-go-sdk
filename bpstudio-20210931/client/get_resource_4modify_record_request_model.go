// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResource4ModifyRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *GetResource4ModifyRecordRequest
	GetApplicationId() *string
	SetMaxResults(v int64) *GetResource4ModifyRecordRequest
	GetMaxResults() *int64
	SetNextToken(v int64) *GetResource4ModifyRecordRequest
	GetNextToken() *int64
}

type GetResource4ModifyRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// BE68D71ZY5YYIU9R
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 1
	NextToken *int64 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s GetResource4ModifyRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResource4ModifyRecordRequest) GoString() string {
	return s.String()
}

func (s *GetResource4ModifyRecordRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetResource4ModifyRecordRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *GetResource4ModifyRecordRequest) GetNextToken() *int64 {
	return s.NextToken
}

func (s *GetResource4ModifyRecordRequest) SetApplicationId(v string) *GetResource4ModifyRecordRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetResource4ModifyRecordRequest) SetMaxResults(v int64) *GetResource4ModifyRecordRequest {
	s.MaxResults = &v
	return s
}

func (s *GetResource4ModifyRecordRequest) SetNextToken(v int64) *GetResource4ModifyRecordRequest {
	s.NextToken = &v
	return s
}

func (s *GetResource4ModifyRecordRequest) Validate() error {
	return dara.Validate(s)
}
