// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchResult interface {
	dara.Model
	String() string
	GoString() string
	SetResultMessage(v string) *FetchResult
	GetResultMessage() *string
	SetResultType(v string) *FetchResult
	GetResultType() *string
	SetTableResults(v []*TableResult) *FetchResult
	GetTableResults() []*TableResult
}

type FetchResult struct {
	ResultMessage *string        `json:"resultMessage,omitempty" xml:"resultMessage,omitempty"`
	ResultType    *string        `json:"resultType,omitempty" xml:"resultType,omitempty"`
	TableResults  []*TableResult `json:"tableResults,omitempty" xml:"tableResults,omitempty" type:"Repeated"`
}

func (s FetchResult) String() string {
	return dara.Prettify(s)
}

func (s FetchResult) GoString() string {
	return s.String()
}

func (s *FetchResult) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *FetchResult) GetResultType() *string {
	return s.ResultType
}

func (s *FetchResult) GetTableResults() []*TableResult {
	return s.TableResults
}

func (s *FetchResult) SetResultMessage(v string) *FetchResult {
	s.ResultMessage = &v
	return s
}

func (s *FetchResult) SetResultType(v string) *FetchResult {
	s.ResultType = &v
	return s
}

func (s *FetchResult) SetTableResults(v []*TableResult) *FetchResult {
	s.TableResults = v
	return s
}

func (s *FetchResult) Validate() error {
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
