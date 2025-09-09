// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAcrImageTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAcrImageTagsResponseBody
	GetRequestId() *string
}

type DeleteAcrImageTagsResponseBody struct {
	// example:
	//
	// E73F09DC-6C13-5CB1-A10F-7A4E125ABD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAcrImageTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAcrImageTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAcrImageTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAcrImageTagsResponseBody) SetRequestId(v string) *DeleteAcrImageTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAcrImageTagsResponseBody) Validate() error {
	return dara.Validate(s)
}
