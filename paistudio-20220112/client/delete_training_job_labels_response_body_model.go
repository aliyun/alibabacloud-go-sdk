// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrainingJobLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTrainingJobLabelsResponseBody
	GetRequestId() *string
}

type DeleteTrainingJobLabelsResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrainingJobLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrainingJobLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrainingJobLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTrainingJobLabelsResponseBody) SetRequestId(v string) *DeleteTrainingJobLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTrainingJobLabelsResponseBody) Validate() error {
	return dara.Validate(s)
}
