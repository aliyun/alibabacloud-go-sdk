// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSampleDataSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckSampleDataSetResponseBody
	GetRequestId() *string
	SetStatus(v string) *CheckSampleDataSetResponseBody
	GetStatus() *string
}

type CheckSampleDataSetResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0CE655C3-C211-513D-A42F-D4AE2D1A867C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the built-in dataset. Valid values:
	//
	// 	- **SUCCEED**: The dataset is loaded.
	//
	// 	- **INIT**: The dataset is being loaded.
	//
	// 	- **FAILED**: The dataset failed to be loaded.
	//
	// 	- **UNINITIALIZED**: The dataset is not loaded.
	//
	// example:
	//
	// UNINITIALIZED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckSampleDataSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckSampleDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSampleDataSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckSampleDataSetResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CheckSampleDataSetResponseBody) SetRequestId(v string) *CheckSampleDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckSampleDataSetResponseBody) SetStatus(v string) *CheckSampleDataSetResponseBody {
	s.Status = &v
	return s
}

func (s *CheckSampleDataSetResponseBody) Validate() error {
	return dara.Validate(s)
}
