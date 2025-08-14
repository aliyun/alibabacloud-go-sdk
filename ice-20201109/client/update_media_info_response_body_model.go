// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *UpdateMediaInfoResponseBody
	GetMediaId() *string
	SetRequestId(v string) *UpdateMediaInfoResponseBody
	GetRequestId() *string
}

type UpdateMediaInfoResponseBody struct {
	// The ID of the media asset in IMS.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMediaInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMediaInfoResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *UpdateMediaInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMediaInfoResponseBody) SetMediaId(v string) *UpdateMediaInfoResponseBody {
	s.MediaId = &v
	return s
}

func (s *UpdateMediaInfoResponseBody) SetRequestId(v string) *UpdateMediaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMediaInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
