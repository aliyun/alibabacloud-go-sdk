// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterventionDictionaryRelatedEntitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListInterventionDictionaryRelatedEntitiesResponseBody
	GetRequestId() *string
	SetResult(v []map[string]interface{}) *ListInterventionDictionaryRelatedEntitiesResponseBody
	GetResult() []map[string]interface{}
}

type ListInterventionDictionaryRelatedEntitiesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
	Result []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListInterventionDictionaryRelatedEntitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInterventionDictionaryRelatedEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryRelatedEntitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInterventionDictionaryRelatedEntitiesResponseBody) GetResult() []map[string]interface{} {
	return s.Result
}

func (s *ListInterventionDictionaryRelatedEntitiesResponseBody) SetRequestId(v string) *ListInterventionDictionaryRelatedEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterventionDictionaryRelatedEntitiesResponseBody) SetResult(v []map[string]interface{}) *ListInterventionDictionaryRelatedEntitiesResponseBody {
	s.Result = v
	return s
}

func (s *ListInterventionDictionaryRelatedEntitiesResponseBody) Validate() error {
	return dara.Validate(s)
}
