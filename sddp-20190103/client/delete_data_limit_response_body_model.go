// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataLimitResponseBody
	GetRequestId() *string
}

type DeleteDataLimitResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7C3AC882-E5A8-4855-BE77-B6837B695EF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDataLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLimitResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataLimitResponseBody) SetRequestId(v string) *DeleteDataLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataLimitResponseBody) Validate() error {
	return dara.Validate(s)
}
