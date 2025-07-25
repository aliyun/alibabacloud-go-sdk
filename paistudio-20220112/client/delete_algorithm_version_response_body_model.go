// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlgorithmVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAlgorithmVersionResponseBody
	GetRequestId() *string
}

type DeleteAlgorithmVersionResponseBody struct {
	// example:
	//
	// F082BD0D-21E1-5F9B-81A0-AB07485B03CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAlgorithmVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlgorithmVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlgorithmVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAlgorithmVersionResponseBody) SetRequestId(v string) *DeleteAlgorithmVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlgorithmVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
