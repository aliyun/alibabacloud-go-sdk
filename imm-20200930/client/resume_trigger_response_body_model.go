// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeTriggerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResumeTriggerResponseBody
	GetRequestId() *string
}

type ResumeTriggerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FEDC9B1F-30F2-4C1F-8ED2-B7860187****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeTriggerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeTriggerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeTriggerResponseBody) SetRequestId(v string) *ResumeTriggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeTriggerResponseBody) Validate() error {
	return dara.Validate(s)
}
