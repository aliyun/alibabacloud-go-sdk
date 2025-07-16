// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveId(v string) *CreateLiveResponseBody
	GetLiveId() *string
	SetRequestId(v string) *CreateLiveResponseBody
	GetRequestId() *string
}

type CreateLiveResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	LiveId *string `json:"liveId,omitempty" xml:"liveId,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateLiveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLiveResponseBody) GetLiveId() *string {
	return s.LiveId
}

func (s *CreateLiveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLiveResponseBody) SetLiveId(v string) *CreateLiveResponseBody {
	s.LiveId = &v
	return s
}

func (s *CreateLiveResponseBody) SetRequestId(v string) *CreateLiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLiveResponseBody) Validate() error {
	return dara.Validate(s)
}
