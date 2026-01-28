// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteApplicationResponseBody
	GetCode() *string
	SetData(v string) *DeleteApplicationResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteApplicationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteApplicationResponseBody
	GetRequestId() *string
}

type DeleteApplicationResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// a395011f-a247-400f-bc69-28796749fd52
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D771A1B6-3D5F-174A-BEE1-98CE1000D337
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteApplicationResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteApplicationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApplicationResponseBody) SetCode(v string) *DeleteApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetData(v string) *DeleteApplicationResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetHttpStatusCode(v int32) *DeleteApplicationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetMessage(v string) *DeleteApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetRequestId(v string) *DeleteApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
