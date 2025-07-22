// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateInstanceLabelsResponseBody
	GetRequestId() *string
}

type UpdateInstanceLabelsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceLabelsResponseBody) SetRequestId(v string) *UpdateInstanceLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceLabelsResponseBody) Validate() error {
	return dara.Validate(s)
}
