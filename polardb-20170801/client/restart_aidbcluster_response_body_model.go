// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartAIDBClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestartAIDBClusterResponseBody
	GetRequestId() *string
}

type RestartAIDBClusterResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 24A1990B-4F6E-482B-B8CB-75C612******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartAIDBClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartAIDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *RestartAIDBClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartAIDBClusterResponseBody) SetRequestId(v string) *RestartAIDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartAIDBClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
