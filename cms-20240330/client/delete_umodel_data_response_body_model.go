// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUmodelDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteUmodelDataResponseBody
	GetRequestId() *string
}

type DeleteUmodelDataResponseBody struct {
	// example:
	//
	// 111111-222-333-1111-33333
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteUmodelDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUmodelDataResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUmodelDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUmodelDataResponseBody) SetRequestId(v string) *DeleteUmodelDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUmodelDataResponseBody) Validate() error {
	return dara.Validate(s)
}
