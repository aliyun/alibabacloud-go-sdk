// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOnlineConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListOnlineConfigsResponseBody
	GetRequestId() *string
	SetResult(v []*ListOnlineConfigsResponseBodyResult) *ListOnlineConfigsResponseBody
	GetResult() []*ListOnlineConfigsResponseBodyResult
}

type ListOnlineConfigsResponseBody struct {
	// id of request
	//
	// example:
	//
	// E45380E8-994A-5402-9806-F114B3295FCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// List
	Result []*ListOnlineConfigsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListOnlineConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOnlineConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOnlineConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOnlineConfigsResponseBody) GetResult() []*ListOnlineConfigsResponseBodyResult {
	return s.Result
}

func (s *ListOnlineConfigsResponseBody) SetRequestId(v string) *ListOnlineConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOnlineConfigsResponseBody) SetResult(v []*ListOnlineConfigsResponseBodyResult) *ListOnlineConfigsResponseBody {
	s.Result = v
	return s
}

func (s *ListOnlineConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListOnlineConfigsResponseBodyResult struct {
	// The configuration information
	//
	// example:
	//
	// {\\"specItems\\":[{\\"specKey\\":\\"YQ_KEYWORD_NUMBER_PLUS\\",\\"value\\":\\"1\\"}]}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// The name of the index
	//
	// example:
	//
	// generation
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
}

func (s ListOnlineConfigsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListOnlineConfigsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListOnlineConfigsResponseBodyResult) GetConfig() *string {
	return s.Config
}

func (s *ListOnlineConfigsResponseBodyResult) GetIndexName() *string {
	return s.IndexName
}

func (s *ListOnlineConfigsResponseBodyResult) SetConfig(v string) *ListOnlineConfigsResponseBodyResult {
	s.Config = &v
	return s
}

func (s *ListOnlineConfigsResponseBodyResult) SetIndexName(v string) *ListOnlineConfigsResponseBodyResult {
	s.IndexName = &v
	return s
}

func (s *ListOnlineConfigsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
