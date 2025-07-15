// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKickLiveMessageGroupUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *KickLiveMessageGroupUserResponseBody
	GetRequestId() *string
}

type KickLiveMessageGroupUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 97168878-5288-10CE-AE56-E2D1627FB5F4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s KickLiveMessageGroupUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s KickLiveMessageGroupUserResponseBody) GoString() string {
	return s.String()
}

func (s *KickLiveMessageGroupUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *KickLiveMessageGroupUserResponseBody) SetRequestId(v string) *KickLiveMessageGroupUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *KickLiveMessageGroupUserResponseBody) Validate() error {
	return dara.Validate(s)
}
