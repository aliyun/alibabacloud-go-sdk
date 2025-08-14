// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecognitionSampleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRecognitionSampleResponseBody
	GetRequestId() *string
}

type DeleteRecognitionSampleResponseBody struct {
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRecognitionSampleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecognitionSampleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecognitionSampleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRecognitionSampleResponseBody) SetRequestId(v string) *DeleteRecognitionSampleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRecognitionSampleResponseBody) Validate() error {
	return dara.Validate(s)
}
