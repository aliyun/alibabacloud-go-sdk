// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHiveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHiveResponseBody
	GetRequestId() *string
}

type DeleteHiveResponseBody struct {
	// example:
	//
	// xxxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHiveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHiveResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHiveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHiveResponseBody) SetRequestId(v string) *DeleteHiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHiveResponseBody) Validate() error {
	return dara.Validate(s)
}
