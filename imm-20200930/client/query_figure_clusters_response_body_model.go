// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFigureClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFigureClusters(v []*FigureCluster) *QueryFigureClustersResponseBody
	GetFigureClusters() []*FigureCluster
	SetNextToken(v string) *QueryFigureClustersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *QueryFigureClustersResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *QueryFigureClustersResponseBody
	GetTotalCount() *int64
}

type QueryFigureClustersResponseBody struct {
	// The face groups.
	FigureClusters []*FigureCluster `json:"FigureClusters,omitempty" xml:"FigureClusters,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 10
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CA995EFD-083D-4F40-BE8A-BDF75FFF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of face groups that matches the current query conditions.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryFigureClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFigureClustersResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFigureClustersResponseBody) GetFigureClusters() []*FigureCluster {
	return s.FigureClusters
}

func (s *QueryFigureClustersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryFigureClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFigureClustersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QueryFigureClustersResponseBody) SetFigureClusters(v []*FigureCluster) *QueryFigureClustersResponseBody {
	s.FigureClusters = v
	return s
}

func (s *QueryFigureClustersResponseBody) SetNextToken(v string) *QueryFigureClustersResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryFigureClustersResponseBody) SetRequestId(v string) *QueryFigureClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFigureClustersResponseBody) SetTotalCount(v int64) *QueryFigureClustersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryFigureClustersResponseBody) Validate() error {
	if s.FigureClusters != nil {
		for _, item := range s.FigureClusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
