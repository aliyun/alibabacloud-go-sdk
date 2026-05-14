// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDeleteClientResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkDeleteClientResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkDeleteClientResponseBody
	GetCode() *string
	SetData(v *ClinkDeleteClientResponseBodyData) *ClinkDeleteClientResponseBody
	GetData() *ClinkDeleteClientResponseBodyData
	SetMessage(v string) *ClinkDeleteClientResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkDeleteClientResponseBody
	GetRequestId() *string
}

type ClinkDeleteClientResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkDeleteClientResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkDeleteClientResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkDeleteClientResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkDeleteClientResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkDeleteClientResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkDeleteClientResponseBody) GetData() *ClinkDeleteClientResponseBodyData {
	return s.Data
}

func (s *ClinkDeleteClientResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkDeleteClientResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkDeleteClientResponseBody) SetAccessDeniedDetail(v string) *ClinkDeleteClientResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkDeleteClientResponseBody) SetCode(v string) *ClinkDeleteClientResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkDeleteClientResponseBody) SetData(v *ClinkDeleteClientResponseBodyData) *ClinkDeleteClientResponseBody {
	s.Data = v
	return s
}

func (s *ClinkDeleteClientResponseBody) SetMessage(v string) *ClinkDeleteClientResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkDeleteClientResponseBody) SetRequestId(v string) *ClinkDeleteClientResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkDeleteClientResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkDeleteClientResponseBodyData struct {
	// 请求 id
	//
	// example:
	//
	// xxx
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
	// 座席号
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
}

func (s ClinkDeleteClientResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkDeleteClientResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkDeleteClientResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkDeleteClientResponseBodyData) GetCno() *string {
	return s.Cno
}

func (s *ClinkDeleteClientResponseBodyData) SetClinkRequestId(v string) *ClinkDeleteClientResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkDeleteClientResponseBodyData) SetCno(v string) *ClinkDeleteClientResponseBodyData {
	s.Cno = &v
	return s
}

func (s *ClinkDeleteClientResponseBodyData) Validate() error {
	return dara.Validate(s)
}
