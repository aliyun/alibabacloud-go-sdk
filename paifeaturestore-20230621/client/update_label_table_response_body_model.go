// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLabelTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLabelTableResponseBody
	GetRequestId() *string
}

type UpdateLabelTableResponseBody struct {
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLabelTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLabelTableResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLabelTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLabelTableResponseBody) SetRequestId(v string) *UpdateLabelTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLabelTableResponseBody) Validate() error {
	return dara.Validate(s)
}
