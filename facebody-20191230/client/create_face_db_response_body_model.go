// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFaceDbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateFaceDbResponseBody
	GetRequestId() *string
}

type CreateFaceDbResponseBody struct {
	// example:
	//
	// 818D24A6-E368-46B3-99C5-3CF36D98CCA3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFaceDbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFaceDbResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFaceDbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFaceDbResponseBody) SetRequestId(v string) *CreateFaceDbResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFaceDbResponseBody) Validate() error {
	return dara.Validate(s)
}
