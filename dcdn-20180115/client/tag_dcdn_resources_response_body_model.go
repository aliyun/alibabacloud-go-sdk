// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagDcdnResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TagDcdnResourcesResponseBody
	GetRequestId() *string
}

type TagDcdnResourcesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 97C68796-EB7F-4D41-9D5B-12B909D76508
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagDcdnResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TagDcdnResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagDcdnResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TagDcdnResourcesResponseBody) SetRequestId(v string) *TagDcdnResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagDcdnResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
