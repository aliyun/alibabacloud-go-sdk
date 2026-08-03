// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelInspectionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelInspectionTaskResponseBody
	GetRequestId() *string
}

type CancelInspectionTaskResponseBody struct {
	// example:
	//
	// 2BE6E619-A657-42E3-AD2D-18F8428A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelInspectionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelInspectionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelInspectionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelInspectionTaskResponseBody) SetRequestId(v string) *CancelInspectionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelInspectionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
