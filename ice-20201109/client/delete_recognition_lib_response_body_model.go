// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecognitionLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRecognitionLibResponseBody
	GetRequestId() *string
}

type DeleteRecognitionLibResponseBody struct {
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRecognitionLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecognitionLibResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecognitionLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRecognitionLibResponseBody) SetRequestId(v string) *DeleteRecognitionLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRecognitionLibResponseBody) Validate() error {
	return dara.Validate(s)
}
