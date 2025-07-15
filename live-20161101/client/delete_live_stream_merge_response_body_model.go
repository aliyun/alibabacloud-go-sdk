// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamMergeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveStreamMergeResponseBody
	GetRequestId() *string
}

type DeleteLiveStreamMergeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB9*********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveStreamMergeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamMergeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamMergeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveStreamMergeResponseBody) SetRequestId(v string) *DeleteLiveStreamMergeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveStreamMergeResponseBody) Validate() error {
	return dara.Validate(s)
}
