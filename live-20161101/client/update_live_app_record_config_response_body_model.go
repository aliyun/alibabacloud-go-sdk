// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveAppRecordConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveAppRecordConfigResponseBody
	GetRequestId() *string
}

type UpdateLiveAppRecordConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveAppRecordConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveAppRecordConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveAppRecordConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveAppRecordConfigResponseBody) SetRequestId(v string) *UpdateLiveAppRecordConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveAppRecordConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
