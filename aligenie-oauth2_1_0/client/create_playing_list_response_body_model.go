// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePlayingListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreatePlayingListResponseBody
	GetRequestId() *string
}

type CreatePlayingListResponseBody struct {
	// example:
	//
	// 10002398812
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePlayingListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePlayingListResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePlayingListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePlayingListResponseBody) SetRequestId(v string) *CreatePlayingListResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePlayingListResponseBody) Validate() error {
	return dara.Validate(s)
}
