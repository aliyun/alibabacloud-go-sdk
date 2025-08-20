// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteDBClusterResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *DeleteDBClusterResponseBody
	GetRequestId() *string
}

type DeleteDBClusterResponseBody struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition (V3.0) cluster.
	//
	// example:
	//
	// amv-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteDBClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDBClusterResponseBody) SetDBClusterId(v string) *DeleteDBClusterResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBClusterResponseBody) SetRequestId(v string) *DeleteDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
