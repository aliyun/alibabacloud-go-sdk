// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImagesFromLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteImagesFromLibResponseBody
	GetCode() *int32
	SetData(v bool) *DeleteImagesFromLibResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *DeleteImagesFromLibResponseBody
	GetHttpStatusCode() *int32
	SetMsg(v string) *DeleteImagesFromLibResponseBody
	GetMsg() *string
	SetRequestId(v string) *DeleteImagesFromLibResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteImagesFromLibResponseBody
	GetSuccess() *bool
}

type DeleteImagesFromLibResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success indicator.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteImagesFromLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteImagesFromLibResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImagesFromLibResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteImagesFromLibResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteImagesFromLibResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteImagesFromLibResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *DeleteImagesFromLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteImagesFromLibResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteImagesFromLibResponseBody) SetCode(v int32) *DeleteImagesFromLibResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteImagesFromLibResponseBody) SetData(v bool) *DeleteImagesFromLibResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteImagesFromLibResponseBody) SetHttpStatusCode(v int32) *DeleteImagesFromLibResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteImagesFromLibResponseBody) SetMsg(v string) *DeleteImagesFromLibResponseBody {
	s.Msg = &v
	return s
}

func (s *DeleteImagesFromLibResponseBody) SetRequestId(v string) *DeleteImagesFromLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteImagesFromLibResponseBody) SetSuccess(v bool) *DeleteImagesFromLibResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteImagesFromLibResponseBody) Validate() error {
	return dara.Validate(s)
}
