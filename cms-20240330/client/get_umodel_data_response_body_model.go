// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUmodelDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrors(v []*GetUmodelDataResponseBodyErrors) *GetUmodelDataResponseBody
	GetErrors() []*GetUmodelDataResponseBodyErrors
	SetLinks(v []interface{}) *GetUmodelDataResponseBody
	GetLinks() []interface{}
	SetNodes(v []interface{}) *GetUmodelDataResponseBody
	GetNodes() []interface{}
	SetRequestId(v string) *GetUmodelDataResponseBody
	GetRequestId() *string
	SetTotalLinksCount(v int32) *GetUmodelDataResponseBody
	GetTotalLinksCount() *int32
	SetTotalNodesCount(v int32) *GetUmodelDataResponseBody
	GetTotalNodesCount() *int32
}

type GetUmodelDataResponseBody struct {
	// Error information
	Errors []*GetUmodelDataResponseBodyErrors `json:"errors,omitempty" xml:"errors,omitempty" type:"Repeated"`
	// List of node link relationships
	Links []interface{} `json:"links,omitempty" xml:"links,omitempty" type:"Repeated"`
	// List of nodes
	Nodes []interface{} `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	// Request ID
	//
	// example:
	//
	// 123-123-234-345-123
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Total number of node links
	//
	// example:
	//
	// 0
	TotalLinksCount *int32 `json:"totalLinksCount,omitempty" xml:"totalLinksCount,omitempty"`
	// Total number of nodes
	//
	// example:
	//
	// 0
	TotalNodesCount *int32 `json:"totalNodesCount,omitempty" xml:"totalNodesCount,omitempty"`
}

func (s GetUmodelDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUmodelDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetUmodelDataResponseBody) GetErrors() []*GetUmodelDataResponseBodyErrors {
	return s.Errors
}

func (s *GetUmodelDataResponseBody) GetLinks() []interface{} {
	return s.Links
}

func (s *GetUmodelDataResponseBody) GetNodes() []interface{} {
	return s.Nodes
}

func (s *GetUmodelDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUmodelDataResponseBody) GetTotalLinksCount() *int32 {
	return s.TotalLinksCount
}

func (s *GetUmodelDataResponseBody) GetTotalNodesCount() *int32 {
	return s.TotalNodesCount
}

func (s *GetUmodelDataResponseBody) SetErrors(v []*GetUmodelDataResponseBodyErrors) *GetUmodelDataResponseBody {
	s.Errors = v
	return s
}

func (s *GetUmodelDataResponseBody) SetLinks(v []interface{}) *GetUmodelDataResponseBody {
	s.Links = v
	return s
}

func (s *GetUmodelDataResponseBody) SetNodes(v []interface{}) *GetUmodelDataResponseBody {
	s.Nodes = v
	return s
}

func (s *GetUmodelDataResponseBody) SetRequestId(v string) *GetUmodelDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUmodelDataResponseBody) SetTotalLinksCount(v int32) *GetUmodelDataResponseBody {
	s.TotalLinksCount = &v
	return s
}

func (s *GetUmodelDataResponseBody) SetTotalNodesCount(v int32) *GetUmodelDataResponseBody {
	s.TotalNodesCount = &v
	return s
}

func (s *GetUmodelDataResponseBody) Validate() error {
	if s.Errors != nil {
		for _, item := range s.Errors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetUmodelDataResponseBodyErrors struct {
	// Details.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Error type
	//
	// example:
	//
	// external
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetUmodelDataResponseBodyErrors) String() string {
	return dara.Prettify(s)
}

func (s GetUmodelDataResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *GetUmodelDataResponseBodyErrors) GetMessage() *string {
	return s.Message
}

func (s *GetUmodelDataResponseBodyErrors) GetType() *string {
	return s.Type
}

func (s *GetUmodelDataResponseBodyErrors) SetMessage(v string) *GetUmodelDataResponseBodyErrors {
	s.Message = &v
	return s
}

func (s *GetUmodelDataResponseBodyErrors) SetType(v string) *GetUmodelDataResponseBodyErrors {
	s.Type = &v
	return s
}

func (s *GetUmodelDataResponseBodyErrors) Validate() error {
	return dara.Validate(s)
}
