// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveMessageGroupBandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyLiveMessageGroupBandResponseBody
	GetRequestId() *string
}

type ModifyLiveMessageGroupBandResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 84AF36BF-0B39-1F8A-A416-FAC7C484****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLiveMessageGroupBandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveMessageGroupBandResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLiveMessageGroupBandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLiveMessageGroupBandResponseBody) SetRequestId(v string) *ModifyLiveMessageGroupBandResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLiveMessageGroupBandResponseBody) Validate() error {
	return dara.Validate(s)
}
