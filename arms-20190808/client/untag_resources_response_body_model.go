// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UntagResourcesResponseBody
	GetData() *string
	SetRequestId(v string) *UntagResourcesResponseBody
	GetRequestId() *string
}

type UntagResourcesResponseBody struct {
	// The returned struct.
	//
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FC183B5E-C9AD-5E9E-937F-*******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) GetData() *string {
	return s.Data
}

func (s *UntagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UntagResourcesResponseBody) SetData(v string) *UntagResourcesResponseBody {
	s.Data = &v
	return s
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
