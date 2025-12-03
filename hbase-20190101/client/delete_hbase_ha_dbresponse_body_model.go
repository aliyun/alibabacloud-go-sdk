// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHBaseHaDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHBaseHaDBResponseBody
	GetRequestId() *string
}

type DeleteHBaseHaDBResponseBody struct {
	// example:
	//
	// B409CF51-E01F-4551-BE40-123678FA9026
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHBaseHaDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHBaseHaDBResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHBaseHaDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHBaseHaDBResponseBody) SetRequestId(v string) *DeleteHBaseHaDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHBaseHaDBResponseBody) Validate() error {
	return dara.Validate(s)
}
