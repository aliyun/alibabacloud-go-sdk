// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableGenerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTableGenerationResponseBody
	GetRequestId() *string
	SetResult(v *GetTableGenerationResponseBodyResult) *GetTableGenerationResponseBody
	GetResult() *GetTableGenerationResponseBodyResult
}

type GetTableGenerationResponseBody struct {
	// requestId
	//
	// example:
	//
	// E7B7D598-B080-5C8E-AA35-D43EC0D5F886
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned.
	Result *GetTableGenerationResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetTableGenerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableGenerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableGenerationResponseBody) GetResult() *GetTableGenerationResponseBodyResult {
	return s.Result
}

func (s *GetTableGenerationResponseBody) SetRequestId(v string) *GetTableGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableGenerationResponseBody) SetResult(v *GetTableGenerationResponseBodyResult) *GetTableGenerationResponseBody {
	s.Result = v
	return s
}

func (s *GetTableGenerationResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTableGenerationResponseBodyResult struct {
	// generationId
	//
	// example:
	//
	// 1708674867
	GenerationId *int64 `json:"generationId,omitempty" xml:"generationId,omitempty"`
	// starting, building, ready, stopped, failed
	//
	// example:
	//
	// ready
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetTableGenerationResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetTableGenerationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetTableGenerationResponseBodyResult) GetGenerationId() *int64 {
	return s.GenerationId
}

func (s *GetTableGenerationResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetTableGenerationResponseBodyResult) SetGenerationId(v int64) *GetTableGenerationResponseBodyResult {
	s.GenerationId = &v
	return s
}

func (s *GetTableGenerationResponseBodyResult) SetStatus(v string) *GetTableGenerationResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetTableGenerationResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
