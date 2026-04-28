// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteModelApiResponseBody
	GetRequestId() *string
}

type DeleteModelApiResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6BD9CDE4-5E7B-4BF3-9BB8-83C73E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteModelApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelApiResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteModelApiResponseBody) SetRequestId(v string) *DeleteModelApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteModelApiResponseBody) Validate() error {
	return dara.Validate(s)
}
