// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGeneralConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteGeneralConfigResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteGeneralConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteGeneralConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGeneralConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteGeneralConfigResponseBody
	GetSuccess() *bool
}

type DeleteGeneralConfigResponseBody struct {
	// example:
	//
	// NoData
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGeneralConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGeneralConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGeneralConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteGeneralConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteGeneralConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGeneralConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGeneralConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGeneralConfigResponseBody) SetCode(v string) *DeleteGeneralConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGeneralConfigResponseBody) SetHttpStatusCode(v int32) *DeleteGeneralConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteGeneralConfigResponseBody) SetMessage(v string) *DeleteGeneralConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGeneralConfigResponseBody) SetRequestId(v string) *DeleteGeneralConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGeneralConfigResponseBody) SetSuccess(v bool) *DeleteGeneralConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGeneralConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
