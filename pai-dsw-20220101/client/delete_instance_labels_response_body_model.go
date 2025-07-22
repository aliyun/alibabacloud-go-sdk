// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteInstanceLabelsResponseBody
	GetRequestId() *string
}

type DeleteInstanceLabelsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstanceLabelsResponseBody) SetRequestId(v string) *DeleteInstanceLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceLabelsResponseBody) Validate() error {
	return dara.Validate(s)
}
