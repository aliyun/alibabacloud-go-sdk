// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFeatureConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteFeatureConfigResponseBody
	GetCode() *int32
	SetData(v bool) *DeleteFeatureConfigResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *DeleteFeatureConfigResponseBody
	GetHttpStatusCode() *int32
	SetMsg(v string) *DeleteFeatureConfigResponseBody
	GetMsg() *string
	SetRequestId(v string) *DeleteFeatureConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteFeatureConfigResponseBody
	GetSuccess() *bool
}

type DeleteFeatureConfigResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return result.
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
	// Response message of this request.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// ID assigned by the backend, used to uniquely identify a request. Can be used for troubleshooting.
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

func (s DeleteFeatureConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFeatureConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFeatureConfigResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteFeatureConfigResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteFeatureConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteFeatureConfigResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *DeleteFeatureConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFeatureConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFeatureConfigResponseBody) SetCode(v int32) *DeleteFeatureConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFeatureConfigResponseBody) SetData(v bool) *DeleteFeatureConfigResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFeatureConfigResponseBody) SetHttpStatusCode(v int32) *DeleteFeatureConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteFeatureConfigResponseBody) SetMsg(v string) *DeleteFeatureConfigResponseBody {
	s.Msg = &v
	return s
}

func (s *DeleteFeatureConfigResponseBody) SetRequestId(v string) *DeleteFeatureConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFeatureConfigResponseBody) SetSuccess(v bool) *DeleteFeatureConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFeatureConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
