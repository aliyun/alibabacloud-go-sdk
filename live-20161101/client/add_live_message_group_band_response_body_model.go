// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveMessageGroupBandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLiveMessageGroupBandResponseBody
	GetRequestId() *string
}

type AddLiveMessageGroupBandResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 84AF36BF-0B39-1F8A-A416-FAC7C484****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLiveMessageGroupBandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLiveMessageGroupBandResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveMessageGroupBandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLiveMessageGroupBandResponseBody) SetRequestId(v string) *AddLiveMessageGroupBandResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLiveMessageGroupBandResponseBody) Validate() error {
	return dara.Validate(s)
}
