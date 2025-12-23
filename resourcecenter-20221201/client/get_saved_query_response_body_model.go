// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSavedQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSavedQueryResponseBody
	GetRequestId() *string
	SetSavedQuery(v *GetSavedQueryResponseBodySavedQuery) *GetSavedQueryResponseBody
	GetSavedQuery() *GetSavedQueryResponseBodySavedQuery
}

type GetSavedQueryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6D98D9B0-318D-56A4-910C-93B5F945AF2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the template.
	SavedQuery *GetSavedQueryResponseBodySavedQuery `json:"SavedQuery,omitempty" xml:"SavedQuery,omitempty" type:"Struct"`
}

func (s GetSavedQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSavedQueryResponseBody) GoString() string {
	return s.String()
}

func (s *GetSavedQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSavedQueryResponseBody) GetSavedQuery() *GetSavedQueryResponseBodySavedQuery {
	return s.SavedQuery
}

func (s *GetSavedQueryResponseBody) SetRequestId(v string) *GetSavedQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSavedQueryResponseBody) SetSavedQuery(v *GetSavedQueryResponseBodySavedQuery) *GetSavedQueryResponseBody {
	s.SavedQuery = v
	return s
}

func (s *GetSavedQueryResponseBody) Validate() error {
	if s.SavedQuery != nil {
		if err := s.SavedQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSavedQueryResponseBodySavedQuery struct {
	// The time when the template was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-10-30T01:43:16Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
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
	// SELECT 	- FROM resources;
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// Query of All Alibaba Cloud Resources
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The template ID.
	//
	// example:
	//
	// sq-GeAck****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
	// The time when the template was updated. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-10-30T01:43:16Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetSavedQueryResponseBodySavedQuery) String() string {
	return dara.Prettify(s)
}

func (s GetSavedQueryResponseBodySavedQuery) GoString() string {
	return s.String()
}

func (s *GetSavedQueryResponseBodySavedQuery) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSavedQueryResponseBodySavedQuery) GetDescription() *string {
	return s.Description
}

func (s *GetSavedQueryResponseBodySavedQuery) GetExpression() *string {
	return s.Expression
}

func (s *GetSavedQueryResponseBodySavedQuery) GetName() *string {
	return s.Name
}

func (s *GetSavedQueryResponseBodySavedQuery) GetQueryId() *string {
	return s.QueryId
}

func (s *GetSavedQueryResponseBodySavedQuery) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetSavedQueryResponseBodySavedQuery) SetCreateTime(v string) *GetSavedQueryResponseBodySavedQuery {
	s.CreateTime = &v
	return s
}

func (s *GetSavedQueryResponseBodySavedQuery) SetDescription(v string) *GetSavedQueryResponseBodySavedQuery {
	s.Description = &v
	return s
}

func (s *GetSavedQueryResponseBodySavedQuery) SetExpression(v string) *GetSavedQueryResponseBodySavedQuery {
	s.Expression = &v
	return s
}

func (s *GetSavedQueryResponseBodySavedQuery) SetName(v string) *GetSavedQueryResponseBodySavedQuery {
	s.Name = &v
	return s
}

func (s *GetSavedQueryResponseBodySavedQuery) SetQueryId(v string) *GetSavedQueryResponseBodySavedQuery {
	s.QueryId = &v
	return s
}

func (s *GetSavedQueryResponseBodySavedQuery) SetUpdateTime(v string) *GetSavedQueryResponseBodySavedQuery {
	s.UpdateTime = &v
	return s
}

func (s *GetSavedQueryResponseBodySavedQuery) Validate() error {
	return dara.Validate(s)
}
