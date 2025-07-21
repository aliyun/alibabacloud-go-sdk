// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UntagResourceResponseBody
	GetRequestId() *string
}

type UntagResourceResponseBody struct {
	// example:
	//
	// 4162a6af-bc99-40b3-a552-89dcc8aaf7c8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UntagResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UntagResourceResponseBody) SetRequestId(v string) *UntagResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
