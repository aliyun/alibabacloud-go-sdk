// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectWhiteListsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetProjectWhiteListsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetProjectWhiteListsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetProjectWhiteListsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetProjectWhiteListsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetProjectWhiteListsResponseBody
	GetSuccess() *bool
	SetWhiteLists(v []*GetProjectWhiteListsResponseBodyWhiteLists) *GetProjectWhiteListsResponseBody
	GetWhiteLists() []*GetProjectWhiteListsResponseBodyWhiteLists
}

type GetProjectWhiteListsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	WhiteLists []*GetProjectWhiteListsResponseBodyWhiteLists `json:"WhiteLists,omitempty" xml:"WhiteLists,omitempty" type:"Repeated"`
}

func (s GetProjectWhiteListsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProjectWhiteListsResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectWhiteListsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetProjectWhiteListsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetProjectWhiteListsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetProjectWhiteListsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProjectWhiteListsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetProjectWhiteListsResponseBody) GetWhiteLists() []*GetProjectWhiteListsResponseBodyWhiteLists {
	return s.WhiteLists
}

func (s *GetProjectWhiteListsResponseBody) SetCode(v string) *GetProjectWhiteListsResponseBody {
	s.Code = &v
	return s
}

func (s *GetProjectWhiteListsResponseBody) SetHttpStatusCode(v int32) *GetProjectWhiteListsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetProjectWhiteListsResponseBody) SetMessage(v string) *GetProjectWhiteListsResponseBody {
	s.Message = &v
	return s
}

func (s *GetProjectWhiteListsResponseBody) SetRequestId(v string) *GetProjectWhiteListsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectWhiteListsResponseBody) SetSuccess(v bool) *GetProjectWhiteListsResponseBody {
	s.Success = &v
	return s
}

func (s *GetProjectWhiteListsResponseBody) SetWhiteLists(v []*GetProjectWhiteListsResponseBodyWhiteLists) *GetProjectWhiteListsResponseBody {
	s.WhiteLists = v
	return s
}

func (s *GetProjectWhiteListsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetProjectWhiteListsResponseBodyWhiteLists struct {
	// example:
	//
	// 测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// ip
	//
	// example:
	//
	// 10.1.0.2
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// 5432
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s GetProjectWhiteListsResponseBodyWhiteLists) String() string {
	return dara.Prettify(s)
}

func (s GetProjectWhiteListsResponseBodyWhiteLists) GoString() string {
	return s.String()
}

func (s *GetProjectWhiteListsResponseBodyWhiteLists) GetDescription() *string {
	return s.Description
}

func (s *GetProjectWhiteListsResponseBodyWhiteLists) GetIp() *string {
	return s.Ip
}

func (s *GetProjectWhiteListsResponseBodyWhiteLists) GetPort() *string {
	return s.Port
}

func (s *GetProjectWhiteListsResponseBodyWhiteLists) SetDescription(v string) *GetProjectWhiteListsResponseBodyWhiteLists {
	s.Description = &v
	return s
}

func (s *GetProjectWhiteListsResponseBodyWhiteLists) SetIp(v string) *GetProjectWhiteListsResponseBodyWhiteLists {
	s.Ip = &v
	return s
}

func (s *GetProjectWhiteListsResponseBodyWhiteLists) SetPort(v string) *GetProjectWhiteListsResponseBodyWhiteLists {
	s.Port = &v
	return s
}

func (s *GetProjectWhiteListsResponseBodyWhiteLists) Validate() error {
	return dara.Validate(s)
}
