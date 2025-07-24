// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *RunClusterResponseBody
	GetClusterId() *string
	SetOperationId(v string) *RunClusterResponseBody
	GetOperationId() *string
	SetRequestId(v string) *RunClusterResponseBody
	GetRequestId() *string
}

type RunClusterResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// c-b933c5aac7f7***
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The operation ID.
	//
	// example:
	//
	// op-13c37a77c505****
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunClusterResponseBody) GoString() string {
	return s.String()
}

func (s *RunClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *RunClusterResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *RunClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunClusterResponseBody) SetClusterId(v string) *RunClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *RunClusterResponseBody) SetOperationId(v string) *RunClusterResponseBody {
	s.OperationId = &v
	return s
}

func (s *RunClusterResponseBody) SetRequestId(v string) *RunClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
