// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAliasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAliasesResponseBody
	GetRequestId() *string
	SetResult(v []*ListAliasesResponseBodyResult) *ListAliasesResponseBody
	GetResult() []*ListAliasesResponseBodyResult
}

type ListAliasesResponseBody struct {
	// id of request
	//
	// example:
	//
	// 10D5E615-69F7-5F49-B850-00169ADE513C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// List
	Result []*ListAliasesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListAliasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAliasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAliasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAliasesResponseBody) GetResult() []*ListAliasesResponseBodyResult {
	return s.Result
}

func (s *ListAliasesResponseBody) SetRequestId(v string) *ListAliasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAliasesResponseBody) SetResult(v []*ListAliasesResponseBodyResult) *ListAliasesResponseBody {
	s.Result = v
	return s
}

func (s *ListAliasesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAliasesResponseBodyResult struct {
	// alias name
	//
	// example:
	//
	// test
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// index name
	//
	// example:
	//
	// index
	Index *string `json:"index,omitempty" xml:"index,omitempty"`
}

func (s ListAliasesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListAliasesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAliasesResponseBodyResult) GetAlias() *string {
	return s.Alias
}

func (s *ListAliasesResponseBodyResult) GetIndex() *string {
	return s.Index
}

func (s *ListAliasesResponseBodyResult) SetAlias(v string) *ListAliasesResponseBodyResult {
	s.Alias = &v
	return s
}

func (s *ListAliasesResponseBodyResult) SetIndex(v string) *ListAliasesResponseBodyResult {
	s.Index = &v
	return s
}

func (s *ListAliasesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
