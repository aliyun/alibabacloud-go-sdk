// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnUserConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDcdnUserConfigResponseBody
	GetRequestId() *string
}

type DeleteDcdnUserConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5CC228B4-7A67-4016-9C9F-4A4133494A91
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDcdnUserConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnUserConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDcdnUserConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDcdnUserConfigResponseBody) SetRequestId(v string) *DeleteDcdnUserConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDcdnUserConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
