// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterAndNodesParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterAndNodesParametersResponseBody
	GetRequestId() *string
}

type ModifyDBClusterAndNodesParametersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B7BFB11-C077-4FE3-B051-F69CEB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterAndNodesParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterAndNodesParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAndNodesParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterAndNodesParametersResponseBody) SetRequestId(v string) *ModifyDBClusterAndNodesParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterAndNodesParametersResponseBody) Validate() error {
	return dara.Validate(s)
}
