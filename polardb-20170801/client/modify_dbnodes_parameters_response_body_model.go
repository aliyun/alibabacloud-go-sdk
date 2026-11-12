// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodesParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBNodesParametersResponseBody
	GetRequestId() *string
}

type ModifyDBNodesParametersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EBEAA83D-1734-42E3-85E3-E25F6E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBNodesParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodesParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBNodesParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBNodesParametersResponseBody) SetRequestId(v string) *ModifyDBNodesParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBNodesParametersResponseBody) Validate() error {
	return dara.Validate(s)
}
