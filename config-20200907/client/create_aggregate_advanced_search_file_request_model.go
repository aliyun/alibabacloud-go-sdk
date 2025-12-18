// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggregateAdvancedSearchFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *CreateAggregateAdvancedSearchFileRequest
	GetAggregatorId() *string
	SetSql(v string) *CreateAggregateAdvancedSearchFileRequest
	GetSql() *string
}

type CreateAggregateAdvancedSearchFileRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-edd3626622af00b3****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The SQL statement used to query resources.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT 	- WHERE ResourceType = \\"ACS::ECS::Instance\\"
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
}

func (s CreateAggregateAdvancedSearchFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateAdvancedSearchFileRequest) GoString() string {
	return s.String()
}

func (s *CreateAggregateAdvancedSearchFileRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *CreateAggregateAdvancedSearchFileRequest) GetSql() *string {
	return s.Sql
}

func (s *CreateAggregateAdvancedSearchFileRequest) SetAggregatorId(v string) *CreateAggregateAdvancedSearchFileRequest {
	s.AggregatorId = &v
	return s
}

func (s *CreateAggregateAdvancedSearchFileRequest) SetSql(v string) *CreateAggregateAdvancedSearchFileRequest {
	s.Sql = &v
	return s
}

func (s *CreateAggregateAdvancedSearchFileRequest) Validate() error {
	return dara.Validate(s)
}
