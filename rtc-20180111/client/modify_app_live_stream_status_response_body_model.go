// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppLiveStreamStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAppLiveStreamStatusResponseBody
	GetRequestId() *string
}

type ModifyAppLiveStreamStatusResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppLiveStreamStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppLiveStreamStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppLiveStreamStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppLiveStreamStatusResponseBody) SetRequestId(v string) *ModifyAppLiveStreamStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppLiveStreamStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
