// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveAppRecordConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveAppRecordConfigResponseBody
	GetRequestId() *string
}

type DeleteLiveAppRecordConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6EBD1AC4-C34D-4AE1-963E-B688A228BE31
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveAppRecordConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveAppRecordConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveAppRecordConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveAppRecordConfigResponseBody) SetRequestId(v string) *DeleteLiveAppRecordConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveAppRecordConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
