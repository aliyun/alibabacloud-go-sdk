// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExampleQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExampleQuery(v *GetExampleQueryResponseBodyExampleQuery) *GetExampleQueryResponseBody
	GetExampleQuery() *GetExampleQueryResponseBodyExampleQuery
	SetRequestId(v string) *GetExampleQueryResponseBody
	GetRequestId() *string
}

type GetExampleQueryResponseBody struct {
	// The information about the sample query template.
	ExampleQuery *GetExampleQueryResponseBodyExampleQuery `json:"ExampleQuery,omitempty" xml:"ExampleQuery,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 36A3D9BE-B607-5993-B546-7E19EF65DC00
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetExampleQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExampleQueryResponseBody) GoString() string {
	return s.String()
}

func (s *GetExampleQueryResponseBody) GetExampleQuery() *GetExampleQueryResponseBodyExampleQuery {
	return s.ExampleQuery
}

func (s *GetExampleQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExampleQueryResponseBody) SetExampleQuery(v *GetExampleQueryResponseBodyExampleQuery) *GetExampleQueryResponseBody {
	s.ExampleQuery = v
	return s
}

func (s *GetExampleQueryResponseBody) SetRequestId(v string) *GetExampleQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExampleQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetExampleQueryResponseBodyExampleQuery struct {
	// The description of the template.
	//
	// example:
	//
	// Queries all resources on which you have permissions and sorts the resources by resource type and resource ID.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The query statement in the template.
	//
	// example:
	//
	// SELECT
	//
	//   resource_id,
	//
	//   resource_name,
	//
	//   region_id,
	//
	//   zone_id,
	//
	//   resource_type,
	//
	//   account_id,
	//
	//   create_time,
	//
	//   resource_group_id,
	//
	//   tags,
	//
	//   ip_addresses,
	//
	//   vpc_id,
	//
	//   v_switch_id
	//
	// FROM
	//
	//   resources
	//
	// ORDER BY
	//
	//   resource_type,
	//
	//   resource_id
	//
	// LIMIT
	//
	//   1000 OFFSET 0;
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// Query of All Alibaba Cloud Resources
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// sq-0PfKy****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s GetExampleQueryResponseBodyExampleQuery) String() string {
	return dara.Prettify(s)
}

func (s GetExampleQueryResponseBodyExampleQuery) GoString() string {
	return s.String()
}

func (s *GetExampleQueryResponseBodyExampleQuery) GetDescription() *string {
	return s.Description
}

func (s *GetExampleQueryResponseBodyExampleQuery) GetExpression() *string {
	return s.Expression
}

func (s *GetExampleQueryResponseBodyExampleQuery) GetName() *string {
	return s.Name
}

func (s *GetExampleQueryResponseBodyExampleQuery) GetQueryId() *string {
	return s.QueryId
}

func (s *GetExampleQueryResponseBodyExampleQuery) SetDescription(v string) *GetExampleQueryResponseBodyExampleQuery {
	s.Description = &v
	return s
}

func (s *GetExampleQueryResponseBodyExampleQuery) SetExpression(v string) *GetExampleQueryResponseBodyExampleQuery {
	s.Expression = &v
	return s
}

func (s *GetExampleQueryResponseBodyExampleQuery) SetName(v string) *GetExampleQueryResponseBodyExampleQuery {
	s.Name = &v
	return s
}

func (s *GetExampleQueryResponseBodyExampleQuery) SetQueryId(v string) *GetExampleQueryResponseBodyExampleQuery {
	s.QueryId = &v
	return s
}

func (s *GetExampleQueryResponseBodyExampleQuery) Validate() error {
	return dara.Validate(s)
}
