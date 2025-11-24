// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCensResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v []*string) *DescribeCensResponseBody
	GetClusters() []*string
	SetRequestId(v string) *DescribeCensResponseBody
	GetRequestId() *string
}

type DescribeCensResponseBody struct {
	// The list of Kubernetes clusters that are added to the same ASM instance but are in different VPCs and are not connected by using a Cloud Enterprise Network (CEN) instance.
	Clusters []*string `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCensResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCensResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCensResponseBody) GetClusters() []*string {
	return s.Clusters
}

func (s *DescribeCensResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCensResponseBody) SetClusters(v []*string) *DescribeCensResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeCensResponseBody) SetRequestId(v string) *DescribeCensResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCensResponseBody) Validate() error {
	return dara.Validate(s)
}
