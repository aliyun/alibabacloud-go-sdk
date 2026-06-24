// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManualTriggerMaskingProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ManualTriggerMaskingProcessResponseBody
	GetRequestId() *string
}

type ManualTriggerMaskingProcessResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7C3AC882-E5A8-4855-BE77-B6837B695EF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ManualTriggerMaskingProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ManualTriggerMaskingProcessResponseBody) GoString() string {
	return s.String()
}

func (s *ManualTriggerMaskingProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ManualTriggerMaskingProcessResponseBody) SetRequestId(v string) *ManualTriggerMaskingProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *ManualTriggerMaskingProcessResponseBody) Validate() error {
	return dara.Validate(s)
}
