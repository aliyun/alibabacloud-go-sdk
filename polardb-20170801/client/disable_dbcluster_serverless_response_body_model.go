// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDBClusterServerlessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DisableDBClusterServerlessResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *DisableDBClusterServerlessResponseBody
	GetRequestId() *string
}

type DisableDBClusterServerlessResponseBody struct {
	// The ID of the serverless cluster.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CD3FA5F3-FAF3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableDBClusterServerlessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableDBClusterServerlessResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDBClusterServerlessResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DisableDBClusterServerlessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableDBClusterServerlessResponseBody) SetDBClusterId(v string) *DisableDBClusterServerlessResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DisableDBClusterServerlessResponseBody) SetRequestId(v string) *DisableDBClusterServerlessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableDBClusterServerlessResponseBody) Validate() error {
	return dara.Validate(s)
}
