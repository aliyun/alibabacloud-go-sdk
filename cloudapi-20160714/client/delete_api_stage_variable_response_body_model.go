// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiStageVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteApiStageVariableResponseBody
	GetRequestId() *string
}

type DeleteApiStageVariableResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 03442A3D-3B7D-434C-8A95-A5FEB969B529
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApiStageVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiStageVariableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApiStageVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApiStageVariableResponseBody) SetRequestId(v string) *DeleteApiStageVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApiStageVariableResponseBody) Validate() error {
	return dara.Validate(s)
}
