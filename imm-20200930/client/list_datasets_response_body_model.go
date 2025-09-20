// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatasets(v []*Dataset) *ListDatasetsResponseBody
	GetDatasets() []*Dataset
	SetNextToken(v string) *ListDatasetsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDatasetsResponseBody
	GetRequestId() *string
}

type ListDatasetsResponseBody struct {
	// The list of datasets.
	Datasets []*Dataset `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// The pagination token. If the total number of datasets is greater than the value of MaxResults, you must specify this parameter. This parameter has a value only if not all the datasets that meet the conditions are returned.
	//
	// Pass this value as the value of NextToken in the next call to query subsequent datasets.
	//
	// example:
	//
	// 12345678:immtest:dataset002
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FEEDE356-C928-4A36-951A-6EB5A592****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDatasetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasetsResponseBody) GetDatasets() []*Dataset {
	return s.Datasets
}

func (s *ListDatasetsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDatasetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatasetsResponseBody) SetDatasets(v []*Dataset) *ListDatasetsResponseBody {
	s.Datasets = v
	return s
}

func (s *ListDatasetsResponseBody) SetNextToken(v string) *ListDatasetsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDatasetsResponseBody) SetRequestId(v string) *ListDatasetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatasetsResponseBody) Validate() error {
	return dara.Validate(s)
}
