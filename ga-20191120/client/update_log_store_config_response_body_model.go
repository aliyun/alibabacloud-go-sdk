// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLogStoreConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLogStoreConfigResponseBody
	GetRequestId() *string
}

type UpdateLogStoreConfigResponseBody struct {
	// example:
	//
	// DE77A7F3-3B74-41C0-A5BC-CAFD188C28B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLogStoreConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogStoreConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLogStoreConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLogStoreConfigResponseBody) SetRequestId(v string) *UpdateLogStoreConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLogStoreConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
