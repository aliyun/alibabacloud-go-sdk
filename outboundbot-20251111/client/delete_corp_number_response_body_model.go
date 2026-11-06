// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCorpNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCorpNumberResponseBody
	GetCode() *string
	SetData(v string) *DeleteCorpNumberResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteCorpNumberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteCorpNumberResponseBody
	GetMessage() *string
	SetParams(v []*string) *DeleteCorpNumberResponseBody
	GetParams() []*string
	SetRequestId(v string) *DeleteCorpNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCorpNumberResponseBody
	GetSuccess() *bool
}

type DeleteCorpNumberResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b37
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned by the API.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCorpNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCorpNumberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCorpNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCorpNumberResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteCorpNumberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteCorpNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCorpNumberResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DeleteCorpNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCorpNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCorpNumberResponseBody) SetCode(v string) *DeleteCorpNumberResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCorpNumberResponseBody) SetData(v string) *DeleteCorpNumberResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCorpNumberResponseBody) SetHttpStatusCode(v int32) *DeleteCorpNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteCorpNumberResponseBody) SetMessage(v string) *DeleteCorpNumberResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCorpNumberResponseBody) SetParams(v []*string) *DeleteCorpNumberResponseBody {
	s.Params = v
	return s
}

func (s *DeleteCorpNumberResponseBody) SetRequestId(v string) *DeleteCorpNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCorpNumberResponseBody) SetSuccess(v bool) *DeleteCorpNumberResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCorpNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
