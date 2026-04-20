// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIDBClusterTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteAIDBClusterTaskResponseBody
	GetDBClusterId() *string
	SetRelativeDBClusterId(v string) *DeleteAIDBClusterTaskResponseBody
	GetRelativeDBClusterId() *string
	SetRequestId(v string) *DeleteAIDBClusterTaskResponseBody
	GetRequestId() *string
}

type DeleteAIDBClusterTaskResponseBody struct {
	// example:
	//
	// pm-2ze9***
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// pc-2zejpr***
	RelativeDBClusterId *string `json:"RelativeDBClusterId,omitempty" xml:"RelativeDBClusterId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E56531A4-E552-40BA-9C58-137B80******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAIDBClusterTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIDBClusterTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAIDBClusterTaskResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteAIDBClusterTaskResponseBody) GetRelativeDBClusterId() *string {
	return s.RelativeDBClusterId
}

func (s *DeleteAIDBClusterTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAIDBClusterTaskResponseBody) SetDBClusterId(v string) *DeleteAIDBClusterTaskResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAIDBClusterTaskResponseBody) SetRelativeDBClusterId(v string) *DeleteAIDBClusterTaskResponseBody {
	s.RelativeDBClusterId = &v
	return s
}

func (s *DeleteAIDBClusterTaskResponseBody) SetRequestId(v string) *DeleteAIDBClusterTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAIDBClusterTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
