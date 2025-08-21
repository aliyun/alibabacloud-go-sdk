// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFileInfoResponseBody
	GetRequestId() *string
}

type UpdateFileInfoResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFileInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFileInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFileInfoResponseBody) SetRequestId(v string) *UpdateFileInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFileInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
