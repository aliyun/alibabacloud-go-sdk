// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetJobConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDatasetJobConfigResponseBody
	GetRequestId() *string
}

type UpdateDatasetJobConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D5BFFEE3-6025-443F-8A03-02D619B5C4B9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDatasetJobConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetJobConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDatasetJobConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDatasetJobConfigResponseBody) SetRequestId(v string) *UpdateDatasetJobConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDatasetJobConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
