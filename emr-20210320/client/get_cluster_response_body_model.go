// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCluster(v *Cluster) *GetClusterResponseBody
	GetCluster() *Cluster
	SetRequestId(v string) *GetClusterResponseBody
	GetRequestId() *string
}

type GetClusterResponseBody struct {
	// The details of the cluster.
	Cluster *Cluster `json:"Cluster,omitempty" xml:"Cluster,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBody) GetCluster() *Cluster {
	return s.Cluster
}

func (s *GetClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClusterResponseBody) SetCluster(v *Cluster) *GetClusterResponseBody {
	s.Cluster = v
	return s
}

func (s *GetClusterResponseBody) SetRequestId(v string) *GetClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
