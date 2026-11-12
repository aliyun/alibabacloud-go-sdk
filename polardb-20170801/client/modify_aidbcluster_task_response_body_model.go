// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAIDBClusterTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyAIDBClusterTaskResponseBody
	GetDBClusterId() *string
	SetRelativeDBClusterId(v string) *ModifyAIDBClusterTaskResponseBody
	GetRelativeDBClusterId() *string
	SetRequestId(v string) *ModifyAIDBClusterTaskResponseBody
	GetRequestId() *string
}

type ModifyAIDBClusterTaskResponseBody struct {
	// The task ID.
	//
	// example:
	//
	// pm-2ze9***
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The instance ID of the PolarDB database for the AI model service.
	//
	// example:
	//
	// pc-bp10ze***
	RelativeDBClusterId *string `json:"RelativeDBClusterId,omitempty" xml:"RelativeDBClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6BD9CDE4-5E7B-4BF3-9BB8-83C73E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAIDBClusterTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAIDBClusterTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAIDBClusterTaskResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyAIDBClusterTaskResponseBody) GetRelativeDBClusterId() *string {
	return s.RelativeDBClusterId
}

func (s *ModifyAIDBClusterTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAIDBClusterTaskResponseBody) SetDBClusterId(v string) *ModifyAIDBClusterTaskResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAIDBClusterTaskResponseBody) SetRelativeDBClusterId(v string) *ModifyAIDBClusterTaskResponseBody {
	s.RelativeDBClusterId = &v
	return s
}

func (s *ModifyAIDBClusterTaskResponseBody) SetRequestId(v string) *ModifyAIDBClusterTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAIDBClusterTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
