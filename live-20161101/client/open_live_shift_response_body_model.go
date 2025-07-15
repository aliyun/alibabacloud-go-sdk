// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenLiveShiftResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OpenLiveShiftResponseBody
	GetRequestId() *string
}

type OpenLiveShiftResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenLiveShiftResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenLiveShiftResponseBody) GoString() string {
	return s.String()
}

func (s *OpenLiveShiftResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenLiveShiftResponseBody) SetRequestId(v string) *OpenLiveShiftResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenLiveShiftResponseBody) Validate() error {
	return dara.Validate(s)
}
