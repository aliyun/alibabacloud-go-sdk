// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStudioLayoutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteStudioLayoutResponseBody
	GetRequestId() *string
}

type DeleteStudioLayoutResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0d-f228-4a64-af62-20e9*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStudioLayoutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStudioLayoutResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStudioLayoutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStudioLayoutResponseBody) SetRequestId(v string) *DeleteStudioLayoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStudioLayoutResponseBody) Validate() error {
	return dara.Validate(s)
}
