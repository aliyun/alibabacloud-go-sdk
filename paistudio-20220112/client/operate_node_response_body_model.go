// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *OperateNodeResponseBody
	GetNodeId() *string
	SetRequestId(v string) *OperateNodeResponseBody
	GetRequestId() *string
}

type OperateNodeResponseBody struct {
	// example:
	//
	// lingjunxxxx-mgxxx-xxxx
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 96496E6E-00B4-5F55-80F6-1844FA9E92DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OperateNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateNodeResponseBody) GoString() string {
	return s.String()
}

func (s *OperateNodeResponseBody) GetNodeId() *string {
	return s.NodeId
}

func (s *OperateNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateNodeResponseBody) SetNodeId(v string) *OperateNodeResponseBody {
	s.NodeId = &v
	return s
}

func (s *OperateNodeResponseBody) SetRequestId(v string) *OperateNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
