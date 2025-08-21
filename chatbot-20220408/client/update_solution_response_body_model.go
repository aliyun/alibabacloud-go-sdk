// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSolutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSolutionResponseBody
	GetRequestId() *string
}

type UpdateSolutionResponseBody struct {
	// example:
	//
	// 8B8F098D-A338-54DD-B19C-24BBBCBD8498
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSolutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSolutionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSolutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSolutionResponseBody) SetRequestId(v string) *UpdateSolutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSolutionResponseBody) Validate() error {
	return dara.Validate(s)
}
