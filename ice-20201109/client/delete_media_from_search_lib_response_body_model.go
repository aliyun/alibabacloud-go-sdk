// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaFromSearchLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMediaFromSearchLibResponseBody
	GetCode() *string
	SetMediaId(v string) *DeleteMediaFromSearchLibResponseBody
	GetMediaId() *string
	SetRequestId(v string) *DeleteMediaFromSearchLibResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteMediaFromSearchLibResponseBody
	GetSuccess() *string
}

type DeleteMediaFromSearchLibResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// ******b48fb04483915d4f2cd8******
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMediaFromSearchLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaFromSearchLibResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMediaFromSearchLibResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMediaFromSearchLibResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *DeleteMediaFromSearchLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMediaFromSearchLibResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteMediaFromSearchLibResponseBody) SetCode(v string) *DeleteMediaFromSearchLibResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMediaFromSearchLibResponseBody) SetMediaId(v string) *DeleteMediaFromSearchLibResponseBody {
	s.MediaId = &v
	return s
}

func (s *DeleteMediaFromSearchLibResponseBody) SetRequestId(v string) *DeleteMediaFromSearchLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMediaFromSearchLibResponseBody) SetSuccess(v string) *DeleteMediaFromSearchLibResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMediaFromSearchLibResponseBody) Validate() error {
	return dara.Validate(s)
}
