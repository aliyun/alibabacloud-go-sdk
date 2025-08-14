// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertMediaToSearchLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InsertMediaToSearchLibResponseBody
	GetCode() *string
	SetMediaId(v string) *InsertMediaToSearchLibResponseBody
	GetMediaId() *string
	SetRequestId(v string) *InsertMediaToSearchLibResponseBody
	GetRequestId() *string
	SetSuccess(v string) *InsertMediaToSearchLibResponseBody
	GetSuccess() *string
}

type InsertMediaToSearchLibResponseBody struct {
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
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// *****ACB-44F2-5F2D-88D7-1283E70*****
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

func (s InsertMediaToSearchLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsertMediaToSearchLibResponseBody) GoString() string {
	return s.String()
}

func (s *InsertMediaToSearchLibResponseBody) GetCode() *string {
	return s.Code
}

func (s *InsertMediaToSearchLibResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *InsertMediaToSearchLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsertMediaToSearchLibResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *InsertMediaToSearchLibResponseBody) SetCode(v string) *InsertMediaToSearchLibResponseBody {
	s.Code = &v
	return s
}

func (s *InsertMediaToSearchLibResponseBody) SetMediaId(v string) *InsertMediaToSearchLibResponseBody {
	s.MediaId = &v
	return s
}

func (s *InsertMediaToSearchLibResponseBody) SetRequestId(v string) *InsertMediaToSearchLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertMediaToSearchLibResponseBody) SetSuccess(v string) *InsertMediaToSearchLibResponseBody {
	s.Success = &v
	return s
}

func (s *InsertMediaToSearchLibResponseBody) Validate() error {
	return dara.Validate(s)
}
