// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDqlResult interface {
	dara.Model
	String() string
	GoString() string
	SetStatementIndex(v int32) *DqlResult
	GetStatementIndex() *int32
	SetSubmitPreviewResult(v *SubmitPreviewResult) *DqlResult
	GetSubmitPreviewResult() *SubmitPreviewResult
	SetTableResults(v []*TableResult) *DqlResult
	GetTableResults() []*TableResult
}

type DqlResult struct {
	StatementIndex      *int32               `json:"statementIndex,omitempty" xml:"statementIndex,omitempty"`
	SubmitPreviewResult *SubmitPreviewResult `json:"submitPreviewResult,omitempty" xml:"submitPreviewResult,omitempty"`
	TableResults        []*TableResult       `json:"tableResults,omitempty" xml:"tableResults,omitempty" type:"Repeated"`
}

func (s DqlResult) String() string {
	return dara.Prettify(s)
}

func (s DqlResult) GoString() string {
	return s.String()
}

func (s *DqlResult) GetStatementIndex() *int32 {
	return s.StatementIndex
}

func (s *DqlResult) GetSubmitPreviewResult() *SubmitPreviewResult {
	return s.SubmitPreviewResult
}

func (s *DqlResult) GetTableResults() []*TableResult {
	return s.TableResults
}

func (s *DqlResult) SetStatementIndex(v int32) *DqlResult {
	s.StatementIndex = &v
	return s
}

func (s *DqlResult) SetSubmitPreviewResult(v *SubmitPreviewResult) *DqlResult {
	s.SubmitPreviewResult = v
	return s
}

func (s *DqlResult) SetTableResults(v []*TableResult) *DqlResult {
	s.TableResults = v
	return s
}

func (s *DqlResult) Validate() error {
	if s.SubmitPreviewResult != nil {
		if err := s.SubmitPreviewResult.Validate(); err != nil {
			return err
		}
	}
	if s.TableResults != nil {
		for _, item := range s.TableResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
