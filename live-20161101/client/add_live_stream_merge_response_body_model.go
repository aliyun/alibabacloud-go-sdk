// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveStreamMergeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *AddLiveStreamMergeResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddLiveStreamMergeResponseBody
	GetRequestId() *string
}

type AddLiveStreamMergeResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *AddLiveStreamMergeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddLiveStreamMergeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLiveStreamMergeResponseBody) SetMessage(v string) *AddLiveStreamMergeResponseBody {
	s.Message = &v
	return s
}

func (s *AddLiveStreamMergeResponseBody) SetRequestId(v string) *AddLiveStreamMergeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLiveStreamMergeResponseBody) Validate() error {
	return dara.Validate(s)
}
