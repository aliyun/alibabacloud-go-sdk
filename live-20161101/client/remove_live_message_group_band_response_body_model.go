// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveLiveMessageGroupBandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveLiveMessageGroupBandResponseBody
	GetRequestId() *string
}

type RemoveLiveMessageGroupBandResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 84AF36BF-0B39-1F8A-A416-FAC7C484****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveLiveMessageGroupBandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveLiveMessageGroupBandResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveLiveMessageGroupBandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveLiveMessageGroupBandResponseBody) SetRequestId(v string) *RemoveLiveMessageGroupBandResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveLiveMessageGroupBandResponseBody) Validate() error {
	return dara.Validate(s)
}
