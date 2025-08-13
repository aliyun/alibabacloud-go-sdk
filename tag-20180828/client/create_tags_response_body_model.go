// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTagsResponseBody
	GetRequestId() *string
}

type CreateTagsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 94E16BB6-3FB6-1297-B5B2-ED2250F437CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTagsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTagsResponseBody) SetRequestId(v string) *CreateTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTagsResponseBody) Validate() error {
	return dara.Validate(s)
}
