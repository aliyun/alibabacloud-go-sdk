// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCallCenterProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCallCenterProviderResponseBody
	GetCode() *string
	SetData(v string) *DeleteCallCenterProviderResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteCallCenterProviderResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteCallCenterProviderResponseBody
	GetMessage() *string
	SetParams(v []*string) *DeleteCallCenterProviderResponseBody
	GetParams() []*string
	SetRequestId(v string) *DeleteCallCenterProviderResponseBody
	GetRequestId() *string
}

type DeleteCallCenterProviderResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// xxxxxxxxx
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 7401D698-0AAC-5909-B68E-88C68805FFB8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCallCenterProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCallCenterProviderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCallCenterProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCallCenterProviderResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteCallCenterProviderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteCallCenterProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCallCenterProviderResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DeleteCallCenterProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCallCenterProviderResponseBody) SetCode(v string) *DeleteCallCenterProviderResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCallCenterProviderResponseBody) SetData(v string) *DeleteCallCenterProviderResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCallCenterProviderResponseBody) SetHttpStatusCode(v int32) *DeleteCallCenterProviderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteCallCenterProviderResponseBody) SetMessage(v string) *DeleteCallCenterProviderResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCallCenterProviderResponseBody) SetParams(v []*string) *DeleteCallCenterProviderResponseBody {
	s.Params = v
	return s
}

func (s *DeleteCallCenterProviderResponseBody) SetRequestId(v string) *DeleteCallCenterProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCallCenterProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
