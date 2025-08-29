// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSampleConsistencyJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSampleConsistencyJobResponseBody
	GetRequestId() *string
}

type DeleteSampleConsistencyJobResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSampleConsistencyJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSampleConsistencyJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSampleConsistencyJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSampleConsistencyJobResponseBody) SetRequestId(v string) *DeleteSampleConsistencyJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSampleConsistencyJobResponseBody) Validate() error {
	return dara.Validate(s)
}
