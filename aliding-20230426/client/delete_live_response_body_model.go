// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteLiveResponseBody
	GetSuccess() *bool
}

type DeleteLiveResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteLiveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteLiveResponseBody) SetRequestId(v string) *DeleteLiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveResponseBody) SetSuccess(v bool) *DeleteLiveResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteLiveResponseBody) Validate() error {
	return dara.Validate(s)
}
