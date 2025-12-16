// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateMergedTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GenerateMergedTableResponseBody
	GetRequestId() *string
	SetResult(v *GenerateMergedTableResponseBodyResult) *GenerateMergedTableResponseBody
	GetResult() *GenerateMergedTableResponseBodyResult
}

type GenerateMergedTableResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The response parameters.
	Result *GenerateMergedTableResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GenerateMergedTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateMergedTableResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateMergedTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateMergedTableResponseBody) GetResult() *GenerateMergedTableResponseBodyResult {
	return s.Result
}

func (s *GenerateMergedTableResponseBody) SetRequestId(v string) *GenerateMergedTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateMergedTableResponseBody) SetResult(v *GenerateMergedTableResponseBodyResult) *GenerateMergedTableResponseBody {
	s.Result = v
	return s
}

func (s *GenerateMergedTableResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateMergedTableResponseBodyResult struct {
	// The tables on which the JOIN operation is performed.
	//
	// example:
	//
	// -
	FromTable map[string]interface{} `json:"fromTable,omitempty" xml:"fromTable,omitempty"`
	// The wide table that is generated after the JOIN operation is performed on multiple tables.
	//
	// example:
	//
	// -
	MergeTable map[string]interface{} `json:"mergeTable,omitempty" xml:"mergeTable,omitempty"`
	// The primary key.
	//
	// example:
	//
	// -
	PrimaryKey *string `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
}

func (s GenerateMergedTableResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GenerateMergedTableResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GenerateMergedTableResponseBodyResult) GetFromTable() map[string]interface{} {
	return s.FromTable
}

func (s *GenerateMergedTableResponseBodyResult) GetMergeTable() map[string]interface{} {
	return s.MergeTable
}

func (s *GenerateMergedTableResponseBodyResult) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *GenerateMergedTableResponseBodyResult) SetFromTable(v map[string]interface{}) *GenerateMergedTableResponseBodyResult {
	s.FromTable = v
	return s
}

func (s *GenerateMergedTableResponseBodyResult) SetMergeTable(v map[string]interface{}) *GenerateMergedTableResponseBodyResult {
	s.MergeTable = v
	return s
}

func (s *GenerateMergedTableResponseBodyResult) SetPrimaryKey(v string) *GenerateMergedTableResponseBodyResult {
	s.PrimaryKey = &v
	return s
}

func (s *GenerateMergedTableResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
