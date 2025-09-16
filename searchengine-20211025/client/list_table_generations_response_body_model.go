// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableGenerationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTableGenerationsResponseBody
	GetRequestId() *string
	SetResult(v []*ListTableGenerationsResponseBodyResult) *ListTableGenerationsResponseBody
	GetResult() []*ListTableGenerationsResponseBodyResult
}

type ListTableGenerationsResponseBody struct {
	// requestId
	//
	// example:
	//
	// F6E3D968-529C-5C40-AFDD-133A8B8FD930
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	Result []*ListTableGenerationsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListTableGenerationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTableGenerationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTableGenerationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTableGenerationsResponseBody) GetResult() []*ListTableGenerationsResponseBodyResult {
	return s.Result
}

func (s *ListTableGenerationsResponseBody) SetRequestId(v string) *ListTableGenerationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTableGenerationsResponseBody) SetResult(v []*ListTableGenerationsResponseBodyResult) *ListTableGenerationsResponseBody {
	s.Result = v
	return s
}

func (s *ListTableGenerationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTableGenerationsResponseBodyResult struct {
	// The ID of the full index version.
	//
	// example:
	//
	// 1708674867
	GenerationId *int64 `json:"generationId,omitempty" xml:"generationId,omitempty"`
}

func (s ListTableGenerationsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListTableGenerationsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListTableGenerationsResponseBodyResult) GetGenerationId() *int64 {
	return s.GenerationId
}

func (s *ListTableGenerationsResponseBodyResult) SetGenerationId(v int64) *ListTableGenerationsResponseBodyResult {
	s.GenerationId = &v
	return s
}

func (s *ListTableGenerationsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
