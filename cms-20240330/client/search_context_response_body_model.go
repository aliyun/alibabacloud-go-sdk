// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchContextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SearchContextResponseBody
	GetRequestId() *string
	SetResults(v []map[string]interface{}) *SearchContextResponseBody
	GetResults() []map[string]interface{}
}

type SearchContextResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 123-0F43-23423-AC43-34234
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return value
	Results []map[string]interface{} `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s SearchContextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchContextResponseBody) GoString() string {
	return s.String()
}

func (s *SearchContextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchContextResponseBody) GetResults() []map[string]interface{} {
	return s.Results
}

func (s *SearchContextResponseBody) SetRequestId(v string) *SearchContextResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchContextResponseBody) SetResults(v []map[string]interface{}) *SearchContextResponseBody {
	s.Results = v
	return s
}

func (s *SearchContextResponseBody) Validate() error {
	return dara.Validate(s)
}
