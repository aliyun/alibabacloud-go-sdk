// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentificationSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetIdentificationSessionResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetIdentificationSessionResponseBody
	GetCode() *string
	SetData(v *GetIdentificationSessionResponseBodyData) *GetIdentificationSessionResponseBody
	GetData() *GetIdentificationSessionResponseBodyData
	SetMessage(v string) *GetIdentificationSessionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetIdentificationSessionResponseBody
	GetRequestId() *string
}

type GetIdentificationSessionResponseBody struct {
	// example:
	//
	// -
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetIdentificationSessionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 68A40250-50CD-034C-B728-0BD******177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIdentificationSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIdentificationSessionResponseBody) GoString() string {
	return s.String()
}

func (s *GetIdentificationSessionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetIdentificationSessionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetIdentificationSessionResponseBody) GetData() *GetIdentificationSessionResponseBodyData {
	return s.Data
}

func (s *GetIdentificationSessionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetIdentificationSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIdentificationSessionResponseBody) SetAccessDeniedDetail(v string) *GetIdentificationSessionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetIdentificationSessionResponseBody) SetCode(v string) *GetIdentificationSessionResponseBody {
	s.Code = &v
	return s
}

func (s *GetIdentificationSessionResponseBody) SetData(v *GetIdentificationSessionResponseBodyData) *GetIdentificationSessionResponseBody {
	s.Data = v
	return s
}

func (s *GetIdentificationSessionResponseBody) SetMessage(v string) *GetIdentificationSessionResponseBody {
	s.Message = &v
	return s
}

func (s *GetIdentificationSessionResponseBody) SetRequestId(v string) *GetIdentificationSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIdentificationSessionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentificationSessionResponseBodyData struct {
	// example:
	//
	// 示例值
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// example:
	//
	// 示例值示例值
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetIdentificationSessionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetIdentificationSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIdentificationSessionResponseBodyData) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *GetIdentificationSessionResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *GetIdentificationSessionResponseBodyData) SetExpireDate(v string) *GetIdentificationSessionResponseBodyData {
	s.ExpireDate = &v
	return s
}

func (s *GetIdentificationSessionResponseBodyData) SetSessionId(v string) *GetIdentificationSessionResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *GetIdentificationSessionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
