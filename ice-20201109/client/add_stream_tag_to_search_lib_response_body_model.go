// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddStreamTagToSearchLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddStreamTagToSearchLibResponseBody
	GetCode() *string
	SetMediaId(v string) *AddStreamTagToSearchLibResponseBody
	GetMediaId() *string
	SetRequestId(v string) *AddStreamTagToSearchLibResponseBody
	GetRequestId() *string
	SetSuccess(v string) *AddStreamTagToSearchLibResponseBody
	GetSuccess() *string
}

type AddStreamTagToSearchLibResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// *****ACB-44F2-5F2D-88D7-1283E70*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddStreamTagToSearchLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddStreamTagToSearchLibResponseBody) GoString() string {
	return s.String()
}

func (s *AddStreamTagToSearchLibResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddStreamTagToSearchLibResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *AddStreamTagToSearchLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddStreamTagToSearchLibResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *AddStreamTagToSearchLibResponseBody) SetCode(v string) *AddStreamTagToSearchLibResponseBody {
	s.Code = &v
	return s
}

func (s *AddStreamTagToSearchLibResponseBody) SetMediaId(v string) *AddStreamTagToSearchLibResponseBody {
	s.MediaId = &v
	return s
}

func (s *AddStreamTagToSearchLibResponseBody) SetRequestId(v string) *AddStreamTagToSearchLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddStreamTagToSearchLibResponseBody) SetSuccess(v string) *AddStreamTagToSearchLibResponseBody {
	s.Success = &v
	return s
}

func (s *AddStreamTagToSearchLibResponseBody) Validate() error {
	return dara.Validate(s)
}
