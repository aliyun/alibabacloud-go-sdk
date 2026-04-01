// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTagsToResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddTagsToResourceResponseBody
	GetRequestId() *string
}

type AddTagsToResourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 224DB9F7-3100-4899-AB9C-C938BCCB43E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddTagsToResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTagsToResourceResponseBody) GoString() string {
	return s.String()
}

func (s *AddTagsToResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTagsToResourceResponseBody) SetRequestId(v string) *AddTagsToResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTagsToResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
