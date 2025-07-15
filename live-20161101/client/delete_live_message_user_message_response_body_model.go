// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveMessageUserMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveMessageUserMessageResponseBody
	GetRequestId() *string
}

type DeleteLiveMessageUserMessageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6CFDE7AB-571A-14EA-B072-989FF753****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveMessageUserMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveMessageUserMessageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveMessageUserMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveMessageUserMessageResponseBody) SetRequestId(v string) *DeleteLiveMessageUserMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveMessageUserMessageResponseBody) Validate() error {
	return dara.Validate(s)
}
