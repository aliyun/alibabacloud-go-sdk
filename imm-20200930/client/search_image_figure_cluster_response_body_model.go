// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchImageFigureClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v []*SearchImageFigureClusterResponseBodyClusters) *SearchImageFigureClusterResponseBody
	GetClusters() []*SearchImageFigureClusterResponseBodyClusters
	SetRequestId(v string) *SearchImageFigureClusterResponseBody
	GetRequestId() *string
}

type SearchImageFigureClusterResponseBody struct {
	// The face clusters.
	Clusters []*SearchImageFigureClusterResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C2734912-E6D5-052C-AC67-6A9FD02*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchImageFigureClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchImageFigureClusterResponseBody) GoString() string {
	return s.String()
}

func (s *SearchImageFigureClusterResponseBody) GetClusters() []*SearchImageFigureClusterResponseBodyClusters {
	return s.Clusters
}

func (s *SearchImageFigureClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchImageFigureClusterResponseBody) SetClusters(v []*SearchImageFigureClusterResponseBodyClusters) *SearchImageFigureClusterResponseBody {
	s.Clusters = v
	return s
}

func (s *SearchImageFigureClusterResponseBody) SetRequestId(v string) *SearchImageFigureClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchImageFigureClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchImageFigureClusterResponseBodyClusters struct {
	// The bounding box of the face.
	Boundary *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
	// The ID of the face cluster that contains faces similar to the face within the bounding box.
	//
	// example:
	//
	// Cluster-ca730577-06b1-42c7-a25b-8f2c7******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The similarity between the face within the bounding box and the face cluster. Valid value: 0 to 1.
	//
	// example:
	//
	// 0.87413794
	Similarity *float32 `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
}

func (s SearchImageFigureClusterResponseBodyClusters) String() string {
	return dara.Prettify(s)
}

func (s SearchImageFigureClusterResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *SearchImageFigureClusterResponseBodyClusters) GetBoundary() *Boundary {
	return s.Boundary
}

func (s *SearchImageFigureClusterResponseBodyClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *SearchImageFigureClusterResponseBodyClusters) GetSimilarity() *float32 {
	return s.Similarity
}

func (s *SearchImageFigureClusterResponseBodyClusters) SetBoundary(v *Boundary) *SearchImageFigureClusterResponseBodyClusters {
	s.Boundary = v
	return s
}

func (s *SearchImageFigureClusterResponseBodyClusters) SetClusterId(v string) *SearchImageFigureClusterResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *SearchImageFigureClusterResponseBodyClusters) SetSimilarity(v float32) *SearchImageFigureClusterResponseBodyClusters {
	s.Similarity = &v
	return s
}

func (s *SearchImageFigureClusterResponseBodyClusters) Validate() error {
	return dara.Validate(s)
}
