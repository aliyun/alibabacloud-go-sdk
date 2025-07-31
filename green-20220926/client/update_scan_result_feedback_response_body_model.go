// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScanResultFeedbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateScanResultFeedbackResponseBody
	GetData() *bool
	SetRequestId(v string) *UpdateScanResultFeedbackResponseBody
	GetRequestId() *string
}

type UpdateScanResultFeedbackResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateScanResultFeedbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateScanResultFeedbackResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScanResultFeedbackResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateScanResultFeedbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateScanResultFeedbackResponseBody) SetData(v bool) *UpdateScanResultFeedbackResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateScanResultFeedbackResponseBody) SetRequestId(v string) *UpdateScanResultFeedbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateScanResultFeedbackResponseBody) Validate() error {
	return dara.Validate(s)
}
