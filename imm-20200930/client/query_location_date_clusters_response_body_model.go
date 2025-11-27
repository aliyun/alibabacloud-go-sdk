// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLocationDateClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLocationDateClusters(v []*LocationDateCluster) *QueryLocationDateClustersResponseBody
	GetLocationDateClusters() []*LocationDateCluster
	SetNextToken(v string) *QueryLocationDateClustersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *QueryLocationDateClustersResponseBody
	GetRequestId() *string
}

type QueryLocationDateClustersResponseBody struct {
	// The list of spatiotemporal clusters.
	LocationDateClusters []*LocationDateCluster `json:"LocationDateClusters,omitempty" xml:"LocationDateClusters,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MzQNjmY2MzYxNhNjk2ZNjEu****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7055FCF7-4D7B-098E-BD4D-DD2932B0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryLocationDateClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryLocationDateClustersResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLocationDateClustersResponseBody) GetLocationDateClusters() []*LocationDateCluster {
	return s.LocationDateClusters
}

func (s *QueryLocationDateClustersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryLocationDateClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryLocationDateClustersResponseBody) SetLocationDateClusters(v []*LocationDateCluster) *QueryLocationDateClustersResponseBody {
	s.LocationDateClusters = v
	return s
}

func (s *QueryLocationDateClustersResponseBody) SetNextToken(v string) *QueryLocationDateClustersResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryLocationDateClustersResponseBody) SetRequestId(v string) *QueryLocationDateClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryLocationDateClustersResponseBody) Validate() error {
	if s.LocationDateClusters != nil {
		for _, item := range s.LocationDateClusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
