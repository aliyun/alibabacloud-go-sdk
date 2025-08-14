// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaToSearchLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMediaToSearchLibResponseBody
	GetCode() *string
	SetMediaId(v string) *UpdateMediaToSearchLibResponseBody
	GetMediaId() *string
	SetRequestId(v string) *UpdateMediaToSearchLibResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UpdateMediaToSearchLibResponseBody
	GetSuccess() *string
}

type UpdateMediaToSearchLibResponseBody struct {
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
	// 4E84BE44-58A7-****-****-FBEBEA16EF94
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

func (s UpdateMediaToSearchLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaToSearchLibResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMediaToSearchLibResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMediaToSearchLibResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *UpdateMediaToSearchLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMediaToSearchLibResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UpdateMediaToSearchLibResponseBody) SetCode(v string) *UpdateMediaToSearchLibResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMediaToSearchLibResponseBody) SetMediaId(v string) *UpdateMediaToSearchLibResponseBody {
	s.MediaId = &v
	return s
}

func (s *UpdateMediaToSearchLibResponseBody) SetRequestId(v string) *UpdateMediaToSearchLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMediaToSearchLibResponseBody) SetSuccess(v string) *UpdateMediaToSearchLibResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMediaToSearchLibResponseBody) Validate() error {
	return dara.Validate(s)
}
