// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveAppRecordConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLiveAppRecordConfigResponseBody
	GetRequestId() *string
}

type AddLiveAppRecordConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLiveAppRecordConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLiveAppRecordConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveAppRecordConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLiveAppRecordConfigResponseBody) SetRequestId(v string) *AddLiveAppRecordConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLiveAppRecordConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
