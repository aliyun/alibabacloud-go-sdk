// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagDcdnResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UntagDcdnResourcesResponseBody
	GetRequestId() *string
}

type UntagDcdnResourcesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 97C68796-EB7F-4D41-9D5B-12B909D76508
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagDcdnResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UntagDcdnResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagDcdnResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UntagDcdnResourcesResponseBody) SetRequestId(v string) *UntagDcdnResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagDcdnResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
