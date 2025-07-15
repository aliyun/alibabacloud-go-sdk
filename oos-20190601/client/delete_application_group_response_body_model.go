// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteApplicationGroupResponseBody
	GetRequestId() *string
}

type DeleteApplicationGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9E011F98-4EE5-5E3A-8FA3-D38F2C531C1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApplicationGroupResponseBody) SetRequestId(v string) *DeleteApplicationGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApplicationGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
