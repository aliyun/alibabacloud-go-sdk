// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchSimilarImageClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextMarker(v string) *SearchSimilarImageClustersResponseBody
	GetNextMarker() *string
	SetSimilarImageClusters(v []*SearchSimilarImageClustersResponseBodySimilarImageClusters) *SearchSimilarImageClustersResponseBody
	GetSimilarImageClusters() []*SearchSimilarImageClustersResponseBodySimilarImageClusters
}

type SearchSimilarImageClustersResponseBody struct {
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0***
	NextMarker           *string                                                       `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
	SimilarImageClusters []*SearchSimilarImageClustersResponseBodySimilarImageClusters `json:"similar_image_clusters,omitempty" xml:"similar_image_clusters,omitempty" type:"Repeated"`
}

func (s SearchSimilarImageClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchSimilarImageClustersResponseBody) GoString() string {
	return s.String()
}

func (s *SearchSimilarImageClustersResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *SearchSimilarImageClustersResponseBody) GetSimilarImageClusters() []*SearchSimilarImageClustersResponseBodySimilarImageClusters {
	return s.SimilarImageClusters
}

func (s *SearchSimilarImageClustersResponseBody) SetNextMarker(v string) *SearchSimilarImageClustersResponseBody {
	s.NextMarker = &v
	return s
}

func (s *SearchSimilarImageClustersResponseBody) SetSimilarImageClusters(v []*SearchSimilarImageClustersResponseBodySimilarImageClusters) *SearchSimilarImageClustersResponseBody {
	s.SimilarImageClusters = v
	return s
}

func (s *SearchSimilarImageClustersResponseBody) Validate() error {
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

type SearchSimilarImageClustersResponseBodySimilarImageClusters struct {
	Files []*File `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
}

func (s SearchSimilarImageClustersResponseBodySimilarImageClusters) String() string {
	return dara.Prettify(s)
}

func (s SearchSimilarImageClustersResponseBodySimilarImageClusters) GoString() string {
	return s.String()
}

func (s *SearchSimilarImageClustersResponseBodySimilarImageClusters) GetFiles() []*File {
	return s.Files
}

func (s *SearchSimilarImageClustersResponseBodySimilarImageClusters) SetFiles(v []*File) *SearchSimilarImageClustersResponseBodySimilarImageClusters {
	s.Files = v
	return s
}

func (s *SearchSimilarImageClustersResponseBodySimilarImageClusters) Validate() error {
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
