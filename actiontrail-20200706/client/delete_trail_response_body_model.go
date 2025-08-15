// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTrailResponseBody
	GetRequestId() *string
}

type DeleteTrailResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 145318BE-DEE1-4C57-AA7C-5BE7D34A6AE0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrailResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTrailResponseBody) SetRequestId(v string) *DeleteTrailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTrailResponseBody) Validate() error {
	return dara.Validate(s)
}
