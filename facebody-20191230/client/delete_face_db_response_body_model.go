// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceDbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFaceDbResponseBody
	GetRequestId() *string
}

type DeleteFaceDbResponseBody struct {
	// example:
	//
	// 203AE658-CFE1-507D-B105-F88CD1F1F71B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFaceDbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceDbResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceDbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFaceDbResponseBody) SetRequestId(v string) *DeleteFaceDbResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceDbResponseBody) Validate() error {
	return dara.Validate(s)
}
