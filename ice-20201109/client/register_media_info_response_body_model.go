// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterMediaInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *RegisterMediaInfoResponseBody
	GetMediaId() *string
	SetRequestId(v string) *RegisterMediaInfoResponseBody
	GetRequestId() *string
}

type RegisterMediaInfoResponseBody struct {
	// The ID of the media asset in IMS.
	//
	// example:
	//
	// ******b48fb04483915d4f2cd8******
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******5A-CAAC-4850-A3AF-B74606******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterMediaInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterMediaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterMediaInfoResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *RegisterMediaInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterMediaInfoResponseBody) SetMediaId(v string) *RegisterMediaInfoResponseBody {
	s.MediaId = &v
	return s
}

func (s *RegisterMediaInfoResponseBody) SetRequestId(v string) *RegisterMediaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterMediaInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
