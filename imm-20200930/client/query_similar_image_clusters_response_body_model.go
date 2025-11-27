// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySimilarImageClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *QuerySimilarImageClustersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *QuerySimilarImageClustersResponseBody
	GetRequestId() *string
	SetSimilarImageClusters(v []*SimilarImageCluster) *QuerySimilarImageClustersResponseBody
	GetSimilarImageClusters() []*SimilarImageCluster
}

type QuerySimilarImageClustersResponseBody struct {
	// The pagination token. If the total number of clusters is greater than the value of MaxResults, this token can be used to retrieve the next page. This parameter has a value only if not all the clusters that meet the condition are returned.
	//
	// Pass this value as the value of NextToken in the next query to return the subsequent clusters.
	//
	// example:
	//
	// CAESEgoQCg4KClVwZGF0ZVRpbWUQARgBIs8ECgkAAJLUwUCAQ****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CA995EFD-083D-4F40-BE8A-BDF75FFF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of similar image clusters.
	SimilarImageClusters []*SimilarImageCluster `json:"SimilarImageClusters,omitempty" xml:"SimilarImageClusters,omitempty" type:"Repeated"`
}

func (s QuerySimilarImageClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySimilarImageClustersResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySimilarImageClustersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QuerySimilarImageClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySimilarImageClustersResponseBody) GetSimilarImageClusters() []*SimilarImageCluster {
	return s.SimilarImageClusters
}

func (s *QuerySimilarImageClustersResponseBody) SetNextToken(v string) *QuerySimilarImageClustersResponseBody {
	s.NextToken = &v
	return s
}

func (s *QuerySimilarImageClustersResponseBody) SetRequestId(v string) *QuerySimilarImageClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySimilarImageClustersResponseBody) SetSimilarImageClusters(v []*SimilarImageCluster) *QuerySimilarImageClustersResponseBody {
	s.SimilarImageClusters = v
	return s
}

func (s *QuerySimilarImageClustersResponseBody) Validate() error {
	if s.SimilarImageClusters != nil {
		for _, item := range s.SimilarImageClusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
