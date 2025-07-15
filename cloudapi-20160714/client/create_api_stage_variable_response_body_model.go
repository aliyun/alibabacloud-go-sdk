// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiStageVariableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateApiStageVariableResponseBody
	GetRequestId() *string
}

type CreateApiStageVariableResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 03442A3D-3B7D-434C-8A95-A5FEB999B529
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApiStageVariableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApiStageVariableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApiStageVariableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApiStageVariableResponseBody) SetRequestId(v string) *CreateApiStageVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApiStageVariableResponseBody) Validate() error {
	return dara.Validate(s)
}
