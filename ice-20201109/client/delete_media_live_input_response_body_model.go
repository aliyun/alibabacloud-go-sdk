// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaLiveInputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMediaLiveInputResponseBody
	GetRequestId() *string
}

type DeleteMediaLiveInputResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMediaLiveInputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaLiveInputResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMediaLiveInputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMediaLiveInputResponseBody) SetRequestId(v string) *DeleteMediaLiveInputResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMediaLiveInputResponseBody) Validate() error {
	return dara.Validate(s)
}
