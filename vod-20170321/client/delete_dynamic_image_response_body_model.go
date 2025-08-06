// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDynamicImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDynamicImageResponseBody
	GetRequestId() *string
}

type DeleteDynamicImageResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0C8F0FDD-A99F-4188-B41934C97A54****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDynamicImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDynamicImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDynamicImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDynamicImageResponseBody) SetRequestId(v string) *DeleteDynamicImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDynamicImageResponseBody) Validate() error {
	return dara.Validate(s)
}
