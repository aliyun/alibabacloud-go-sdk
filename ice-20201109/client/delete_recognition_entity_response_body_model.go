// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecognitionEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRecognitionEntityResponseBody
	GetRequestId() *string
}

type DeleteRecognitionEntityResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRecognitionEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecognitionEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecognitionEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRecognitionEntityResponseBody) SetRequestId(v string) *DeleteRecognitionEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRecognitionEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
