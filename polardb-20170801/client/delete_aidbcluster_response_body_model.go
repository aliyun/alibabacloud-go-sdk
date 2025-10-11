// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIDBClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteAIDBClusterResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *DeleteAIDBClusterResponseBody
	GetRequestId() *string
}

type DeleteAIDBClusterResponseBody struct {
	// example:
	//
	// pc-*****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3F9E6A3B-C13E-4064-A010-18582A******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAIDBClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAIDBClusterResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteAIDBClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAIDBClusterResponseBody) SetDBClusterId(v string) *DeleteAIDBClusterResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAIDBClusterResponseBody) SetRequestId(v string) *DeleteAIDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAIDBClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
