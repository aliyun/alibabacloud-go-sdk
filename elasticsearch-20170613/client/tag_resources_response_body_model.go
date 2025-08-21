// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TagResourcesResponseBody
	GetRequestId() *string
	SetResult(v bool) *TagResourcesResponseBody
	GetResult() *bool
}

type TagResourcesResponseBody struct {
	// example:
	//
	// 3D8795D9-8FF5-46B2-86E6-E3B407*******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TagResourcesResponseBody) GetResult() *bool {
	return s.Result
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesResponseBody) SetResult(v bool) *TagResourcesResponseBody {
	s.Result = &v
	return s
}

func (s *TagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
