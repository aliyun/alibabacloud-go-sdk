// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveStreamMergeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLiveStreamMergeResponseBody
	GetRequestId() *string
}

type AddLiveStreamMergeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLiveStreamMergeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLiveStreamMergeResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveStreamMergeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLiveStreamMergeResponseBody) SetRequestId(v string) *AddLiveStreamMergeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLiveStreamMergeResponseBody) Validate() error {
	return dara.Validate(s)
}
