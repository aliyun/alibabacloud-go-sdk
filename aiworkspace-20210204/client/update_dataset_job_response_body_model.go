// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDatasetJobResponseBody
	GetRequestId() *string
}

type UpdateDatasetJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDatasetJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDatasetJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDatasetJobResponseBody) SetRequestId(v string) *UpdateDatasetJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDatasetJobResponseBody) Validate() error {
	return dara.Validate(s)
}
