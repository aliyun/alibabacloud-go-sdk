// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrainingJobLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTrainingJobLabelsResponseBody
	GetRequestId() *string
}

type UpdateTrainingJobLabelsResponseBody struct {
	// example:
	//
	// F082BD0D-21E1-5F9B-81A0-AB07485B03CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTrainingJobLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrainingJobLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTrainingJobLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTrainingJobLabelsResponseBody) SetRequestId(v string) *UpdateTrainingJobLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTrainingJobLabelsResponseBody) Validate() error {
	return dara.Validate(s)
}
