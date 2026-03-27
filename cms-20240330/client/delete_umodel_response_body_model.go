// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUmodelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUmodelResponseBody
	GetRequestId() *string
}

type DeleteUmodelResponseBody struct {
	// request ID
	//
	// example:
	//
	// 123123-3213-345-9941-345345345
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteUmodelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUmodelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUmodelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUmodelResponseBody) SetRequestId(v string) *DeleteUmodelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUmodelResponseBody) Validate() error {
	return dara.Validate(s)
}
