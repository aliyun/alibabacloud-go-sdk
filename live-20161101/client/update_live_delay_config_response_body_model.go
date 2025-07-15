// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveDelayConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveDelayConfigResponseBody
	GetRequestId() *string
}

type UpdateLiveDelayConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveDelayConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveDelayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveDelayConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveDelayConfigResponseBody) SetRequestId(v string) *UpdateLiveDelayConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveDelayConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
